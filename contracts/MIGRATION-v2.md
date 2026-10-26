# Core API v2 Migration

Core/open `v2.0.0` removes the deprecated core endpoints introduced before the
core/EE boundary migration. The v1 contract remains published as
`openapi-v1.yaml` for reference during migration.

## Removed core endpoints

| Removed | Replacement |
| --- | --- |
| `GET /v1/auth/oidc/authorize-url` | EE login (`/v1/auth/oidc/authorize-url` on the EE service) |
| `POST /v1/auth/oidc/token` | EE login (`/v1/auth/oidc/token` on the EE service) |
| `POST /v1/federation/peers` | `POST /v1/enterprise/federation/peers` |
| `GET /v1/federation/peers` | `GET /v1/enterprise/federation/peers` |
| `POST /v1/federation/shifts` | `POST /v1/enterprise/federation/shifts` |
| `POST /v1/federation/shifts/ingest` | `POST /v1/enterprise/federation/shifts/ingest` |
| `GET /v1/federation/shifts` | `GET /v1/enterprise/federation/shifts` |
| `POST /v1/spatial/views` | `POST /v1/enterprise/spatial/views` |
| `GET /v1/spatial/views` | `GET /v1/enterprise/spatial/views` |
| `GET /v1/spatial/views/{entity_ref}` | `GET /v1/enterprise/spatial/views/{entity_ref}` |
| `POST/GET/DELETE /v1/spatial/safety-zones` | `POST/GET/DELETE /v1/enterprise/spatial/safety-zones` |
| `GET /v1/spatial/safety-zones/{id}/violations` | `GET /v1/enterprise/spatial/safety-zones/{id}/violations` |

Retention policy configuration (`SYMBIVELA_RETENTION_POLICY`) and the core
retention scheduler are removed. EE owns policy; core exposes a generic
`PurgeDecisions` executor through the public facade.

## Kept in core

- Governance kernel: goals, attention, decisions, handoffs, outcomes, evidence.
- Approvals and interventions, including the COORDIVELA EE command bridges.
- Notification storage/delivery.
- Attention policy primitive.
- Basic security: trusted headers, workspace roles, service tokens.
- WorkProjection storage/reads.
- Generic retention execution (`PurgeDecisions`).

## Data migration: federation and spatial

Core migrations `0030`/`0031` drop the core federation and spatial tables. EE
migration `0011_core_data_backfill` copies their rows into the EE tables:

| Core table | EE table |
| --- | --- |
| `federation_peers` | `enterprise_federation_peers` |
| `federation_shifts` | `enterprise_federation_shifts` |
| `spatial_views` | `enterprise_spatial_views` |
| `safety_zones` | `enterprise_safety_zones` |
| `safety_violations` | `enterprise_safety_violations` |

**Ordering:** apply EE migrations through `0011` *before* core `0030`/`0031`.
The backfill is idempotent, skips missing core tables, and never overwrites an
existing EE row. For a fresh install the order does not matter.

If core `0030`/`0031` already ran, the source tables are gone. Restore them
from a pre-upgrade backup before applying EE `0011`, or insert the equivalent
rows directly into the EE tables. EE `0011` has no destructive down migration;
rollback requires restoring the pre-upgrade backup.

When core and EE use separate databases, `0011` cannot see the core tables.
Export the five tables from the core database and load them into the EE
database (or a temporary staging schema) before running `0011`.

## Configuration changes

- Set `SYMBIVELA_OIDC_TRUSTED_PROXIES` on EE when login runs behind a trusted
  proxy; otherwise the OIDC token endpoint rate-limits by proxy address.
- Federation and spatial producers now authenticate with the EE service token
  (`X-SYMBIVELA-Service`), not a core workspace role.
- EE exposes `GET /v1/enterprise/metrics` (service token) for enterprise
  counters such as `federation_shift_delivered_total`.

## Migration sequence

1. Back up the database and apply EE migrations through `0011`.
2. Move OIDC/SAML login to EE.
3. Move federation calls to EE `/v1/enterprise/federation/*`.
4. Move spatial calls to EE `/v1/enterprise/spatial/*`.
5. Apply core `0030`/`0031` to drop the core compatibility tables.
6. Replace `SYMBIVELA_RETENTION_POLICY` with EE retention configuration.
7. Remove any SDK methods for the removed endpoints.

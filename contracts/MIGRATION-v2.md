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

## Migration sequence

1. Move OIDC/SAML login to EE.
2. Move federation calls to EE `/v1/enterprise/federation/*`.
3. Move spatial calls to EE `/v1/enterprise/spatial/*`.
4. Replace `SYMBIVELA_RETENTION_POLICY` with EE retention configuration.
5. Remove any SDK methods for the removed endpoints.

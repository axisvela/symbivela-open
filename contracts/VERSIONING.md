# SYMBIVELA Contract Versioning

**Status:** Governance baseline

## Principle

The public API and event contracts are versioned with **semantic versioning** and
**additive-first evolution**. Consumers can rely on backward compatibility within a
major version.

## Rules

- **Additive** changes (new endpoints, new optional fields, new enum values) are
  backward-compatible and released within the same major version.
- **Breaking** changes (removed/renamed endpoints, removed or required-field changes,
  enum removals, semantic changes) require a **major version bump**, migration
  guidance, and a deprecation window.
- `openapi.yaml` declares the current major under `info.version`. Contract schemas
  carry the major in their `$id` (e.g. `.../governance-kernel/v1`).
- Each release pins the contract version it implements; the core `symbivela` repo
  records the supported contract version in its release notes.

## Change Process

1. Additive change: update `contracts/openapi.yaml` and schemas, add SDK coverage.
2. Deprecation: mark the endpoint/schema deprecated; keep it working for at least
   one minor release.
3. Breaking change: create a `vN+1` schema/`openapi.vN+1.yaml`, keep `vN` operational
   during the transition, and document the migration path.

## Scope

Applies to every `contracts/*.schema.json`, `contracts/openapi.yaml`, and the SDKs
generated/validated against them. Internal endpoints (health, ready, metrics) are
operational and excluded from the compatibility contract.

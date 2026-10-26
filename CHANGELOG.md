# Changelog

All notable changes to SYMBIVELA-open are recorded here.
This repository follows Keep a Changelog and Semantic Versioning.

## [Unreleased]

### Changed

- `contracts/MIGRATION-v2.md` documents the EE `0011_core_data_backfill`
  ordering, table mapping, and restore path for core federation/spatial data.

## [v2.0.0] - 2026-10-26

### Removed

- Core OpenAPI paths for `/v1/auth/oidc/*`, `/v1/federation/*`, and
  `/v1/spatial/*`.
- SDK methods for the removed core endpoints.

### Changed

- OpenAPI `info.version` and `x-product-version` are `2.0.0`.
- `contracts/openapi-v1.yaml` preserves the v1 contract.
- `contracts/MIGRATION-v2.md` documents EE replacements and the retention
  executor change.

## [v1.4.0] - 2026-10-26

### Changed

- Deprecation release for the core/EE boundary migration. Core OIDC/SAML,
  federation, spatial, and retention endpoints remain available for the
  compatibility window and are replaced by EE endpoints or facade integration.
- `x-product-version` is `1.4.0`; the OpenAPI contract remains v1.

## [v1.3.0] - 2026-10-25

### Added

- OpenAPI paths for the COORDIVELA EE approval decision
  (`POST /v1/approvals/{id}/ee-decision`), approval external references
  (`POST /v1/approvals/{id}/external-refs`), and intervention command
  (`POST /v1/interventions/{id}/ee-command`).
- `x-product-version: 1.3.0` on the OpenAPI document.
- WorkProjection v1 and WorkProjection request contracts.
- Go and TypeScript SDK `3.1.0` releases.
- Core `v1.3.0` binary release manifest.

### Changed

- Product version now tracks core `1.3.0`.
- The existing `v3.0.0` tag is documented as the `/v3` module-line baseline,
  not a product release.

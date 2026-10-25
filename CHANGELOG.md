# Changelog

All notable changes to SYMBIVELA-open are recorded here.
This repository follows Keep a Changelog and Semantic Versioning.

## [Unreleased]

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

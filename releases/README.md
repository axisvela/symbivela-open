# Release Manifests

Published binary releases include checksums, the originating Core version, and
all applicable license notices.

## v1.1.1

- Manifest: [`v1.1.1/`](v1.1.1/)
- Binaries: `symbivela`, `symbivela-migrate`, `symbivela-audit`, and
  `symbivela-bench` for linux/amd64, darwin/arm64, and windows/amd64.
- Source: [SYMBIVELA core](https://github.com/axisrobo/SYMBIVELA) at commit
  `7d0ccdf` (`v1.1.1`). Binaries are AGPL-3.0-or-later.
- Note: `v1.1.1` moves core services into the reserved port range `1926–1935`
  (core API `:1926` with `--addr`/`SYMBIVELA_ADDR`, frontend dev `1927`).

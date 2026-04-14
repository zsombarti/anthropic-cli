# Changelog

## 1.2.0 (2026-04-14)

Full Changelog: [v1.1.0...v1.2.0](https://github.com/anthropics/anthropic-cli/compare/v1.1.0...v1.2.0)

### Features

* **api:** manual updates ([07273ef](https://github.com/anthropics/anthropic-cli/commit/07273ef2e27993e452db24cfdb59088989349c9f))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([991b8e9](https://github.com/anthropics/anthropic-cli/commit/991b8e972802e2ec3ca5663ab0c6fb31ead8a4df))


### Bug Fixes

* **cli:** fix incompatible Go types for flag generated as array of maps ([ced5845](https://github.com/anthropics/anthropic-cli/commit/ced58459c9d668fdde293adeb4ed676e5c73b800))
* fix for failing to drop invalid module replace in link script ([ad79ded](https://github.com/anthropics/anthropic-cli/commit/ad79ded899364b5e8cb288d90597fd4b7984e538))
* use correct multipart array format ([326a8b5](https://github.com/anthropics/anthropic-cli/commit/326a8b5ae00259c439cf0ea613d57fd41babc602))


### Chores

* add documentation for ./scripts/link ([d1a18e2](https://github.com/anthropics/anthropic-cli/commit/d1a18e23681a821cd3d626bc73d9ad2750e465ab))
* **ci:** remove release-doctor workflow ([2c92e20](https://github.com/anthropics/anthropic-cli/commit/2c92e20fdd01bb42f6051c668cdb7be544ade2d7))
* **cli:** additional test cases for `ShowJSONIterator` ([9c94055](https://github.com/anthropics/anthropic-cli/commit/9c94055e3e651cc383e1022ab3cc1c5474d46167))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([cd58bd2](https://github.com/anthropics/anthropic-cli/commit/cd58bd23c08c9716aa7c73d789b3cbe1662ed9cf))
* **internal:** codegen related update ([48aff04](https://github.com/anthropics/anthropic-cli/commit/48aff040e5a6b166e7f4d0f9073e15dbab875a3d))


### Documentation

* update examples ([3213488](https://github.com/anthropics/anthropic-cli/commit/3213488ea69fab6b47e2cef8c807b26961d857ee))

## 1.1.0 (2026-04-09)

Full Changelog: [v1.0.0...v1.1.0](https://github.com/anthropics/anthropic-cli/compare/v1.0.0...v1.1.0)

### Features

* **api:** manual updates ([0563971](https://github.com/anthropics/anthropic-cli/commit/0563971f7ecbb7a0abe9c7ad4131ce0ec7891b2b))


### Chores

* **cli:** let `--format raw` be used in conjunction with `--transform` ([4748f25](https://github.com/anthropics/anthropic-cli/commit/4748f255fd1e151019115e8e2ed37e0c7a56a607))

## 1.0.0 (2026-04-08)

Full Changelog: [v0.0.1-alpha.0...v1.0.0](https://github.com/anthropics/anthropic-cli/compare/v0.0.1-alpha.0...v1.0.0)

### Features

- Initial release of the `ant` CLI.

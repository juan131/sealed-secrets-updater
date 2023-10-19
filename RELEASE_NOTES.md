# Release Notes

Latest release:

[![](https://img.shields.io/github/release/juan131/sealed-secrets-updater.svg)](https://github.com/juan131/sealed-secrets-updater/releases/latest)

## v0.4.1

- Fix: filter conditions when `--only-secrets` flag is not provided

## v0.4.0

- Feat: add support for updating only provided secrets
- Feat: add support for version subcommand

## v0.3.0

- Feat: add support for skipping certain secrets
- Feat: add support for secrets inputs in CSV

## v0.2.0

- Chore: (deps) bump golang.org/x/net
- Feat: use JSON schema for config validation
- Feat: support both YAML & JSON formats on input files

## v0.1.2

- Fix: creates the output file's parent directory if it's missing.

## v0.1.1

- Fix: avoid encoding secret data as base64 twice

## v0.1.0

- Initial release

# New Relic Infrastructure Integrations builder

You can use this command-line tool to create and scaffold a new integration in
Golang for [New Relic Integration Agent](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure).

## Usage

At this moment, there is only one command available: `init`. This is an example
of the usage:

```bash
$ nr-integrations-builder init mysql
```

You can check the available options for this command using the `--help` option:

```bash
$ nr-integrations-builder init --help
```

Usage with options:

```bash
$ nr-integrations-builder init \
  --company_prefix "nr" \
  --company_name "newrelic" \
  --destination_path "~/Gocode/nr-integrations/" \
  mysql
```

## Install

To install, use `go get`:

```bash
$ go get github.com/newrelic/nr-integrations-builder
```

## Contribution

1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `make test` command and confirm that it passes
1. Create a new Pull Request

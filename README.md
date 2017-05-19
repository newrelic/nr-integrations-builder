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

## Contributing Code

We welcome code contributions (in the form of pull requests) from our user
community.  Before submitting a pull request please review
[these guidelines](https://github.com/newrelic/nr-integrations-builder/blob/master/CONTRIBUTING.md).

Following these helps us efficiently review and incorporate your contribution
and avoid breaking your code with future changes to the agent.

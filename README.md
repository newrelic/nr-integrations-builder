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
  --company-prefix "nr" \
  --company-name "newrelic" \
  --destination-path "~/Gocode/nr-integrations/" \
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

## Support

You can find more detailed documentation [on our website](http://newrelic.com/docs),
and specifically in the [Infrastructure category](https://docs.newrelic.com/docs/infrastructure).

If you can't find what you're looking for there, reach out to us on our [support
site](http://support.newrelic.com/) or our [community forum](http://forum.newrelic.com)
and we'll be happy to help you.

Find a bug? Contact us via [support.newrelic.com](http://support.newrelic.com/),
or email support@newrelic.com.

New Relic, Inc.

# New Relic Infrastructure Integrations builder

You can use this command-line tool to create and scaffold a new integration in
Golang for [New Relic Integration Agent](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure).

## Getting started

### Prerequisites

Before starting to write Go code, we suggest taking a look at Golang's
documentation to setup the environment and familiarize yourself with Golang
language.

It is required to install [the Vendor Tool for Go](https://github.com/kardianos/govendor), which is used for managing dependencies. The project that you will create must be within a `$GOPATH/src`, otherwise the Vendor Tool won't work properly.

#### Windows

You need a bash environment (e.g. [Git bash for windows](https://git-scm.com/download/win)), with `make` and `awk` commands installed.

### Installation

To install the command-line tool, use `go get`:

```bash
$ go get github.com/newrelic/nr-integrations-builder
```

After this step, the binary of `nr-integrations-builder` should be installed in
`$GOPATH/bin` or `$GOBIN` if this last one is set. In a common Golang workspace,
it's recommended having `$GOPATH/bin` or `$GOBIN` included in your `$PATH`. That way, you
would have `nr-integrations-builder` available globally.

## Usage

At this moment, there is only one command available: `init`. You can check the available options for this command using the `--help` option:

```bash
$ nr-integrations-builder init --help
```
It's obligatory to specify `company-name` and `company-prefix` flags. Otherwise, the `nr-integrations-builder` will not initialize the integration.

This is an example of the usage:

```bash
$ nr-integrations-builder init \
  --company-name "myorganization" \
  --company-prefix "myorg" \
  --destination-path "$GOPATH/src/myorg-integrations/" \
  mysql
```

If you don't specify `destination-path` flag the current directory will be used.
It is also possible to specify the flags in short form:
```bash
nr-integrations-builder init mysql -n "myorganization" -c "myorg"
```

## Contributing Code

We welcome code contributions (in the form of pull requests) from our user
community.  Before submitting a pull request please review
[these guidelines](https://github.com/newrelic/nr-integrations-builder/blob/master/CONTRIBUTING.md).

Following these helps us efficiently review and incorporate your contribution
and avoid breaking your code with future changes to the agent.

IMPORTANT: Never edit manually the `scaffold/bindata.go file`. This must be generated using `make generate`

## Support

You can find more detailed documentation [on our website](http://newrelic.com/docs),
and specifically in the [Infrastructure category](https://docs.newrelic.com/docs/infrastructure).

If you can't find what you're looking for there, reach out to us on our [support
site](http://support.newrelic.com/) or our [community forum](http://forum.newrelic.com)
and we'll be happy to help you.

Find a bug? Contact us via [support.newrelic.com](http://support.newrelic.com/),
or email support@newrelic.com.

New Relic, Inc.

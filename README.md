# New Relic Infrastructure Integrations builder

You can use this command-line tool to create and scaffold a new integration in
Golang for [New Relic Integration Agent](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure).

The generated code uses the [New Relic Infrastructure Integrations Go SDK](https://github.com/newrelic/infra-integrations-sdk).
Please visit the [Go SDK documentation](https://github.com/newrelic/infra-integrations-sdk/blob/master/docs/README.md)
for more information about its API and structure.

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

The available options are:

  * `-n` or `--company-name` (string): company name (required).
  * `-c` or `--company-prefix` (string): company prefix identifier (required).
  * `-p` or `--destination-path` (string): the destination path were your integration will be generated as a subfolder
    (default: current directory)
  * `-e` or `--entity-type` (string): type of entity to generate (`remote` (default) or `local`). To know the differences
    between `remote` or `local` entities, please refer to the [Entity definition section in the GoSDK v3 documents](https://github.com/newrelic/infra-integrations-sdk/blob/master/docs/entity-definition.md). 

`company-name` and `company-prefix` flags must be specified. Otherwise, the `nr-integrations-builder` will not initialize the integration.

Following command shows an usage example:

```bash
$ nr-integrations-builder init \
  --company-name myorganization \
  --company-prefix myorg \
  --destination-path $GOPATH/src/myorg-integrations/ \
  mysql
```

It is also possible to specify the flags in short form:
```bash
nr-integrations-builder init mysql -n myorganization -c myorg -p $GOPATH/src/myorg-integrations/
```

The above command would generate the next file structure in the `$GOPATH/src/myorg-integrations/` folder:

```
myorg-integrations
└── mysql
    ├── CHANGELOG.md
    ├── LICENSE
    ├── Makefile
    ├── README.md
    ├── myorg-mysql-config.yml
    ├── myorg-mysql-config.yml.template
    ├── myorg-mysql-definition.yml
    ├── src
    │   ├── mysql.go
    │   └── mysql_test.go
    └── vendor
        ├── (...)
        ├── (vendored third-party dependencies)
        ├── (...)
        └── vendor.json
```

The generated project will contain:

* A `Makefile` with the basic commands to verify, build and test the integration code. Run `make all` to
  generate an integration binary in the `bin/` subfolder.
* An `src/` folder with the basic integration main and test file. The generated main file creates a dummy integration
  with fake data.
* Basic config and definition `.yml` files. Please refer to the [Integrations SDK documentation](https://docs.newrelic.com/docs/integrations/integrations-sdk/file-specifications)
  for more information about configuration and description files.
* `vendor` folder and `vendor/vendor.json` file with all the vendorized third-party libraries. Please refer to
  [the Vendor Tool for Go](https://github.com/kardianos/govendor) for more information about how to install and use
  Go Vendor.
* Project description files: `README.md`, `CHANGELOG.md` and `LICENSE`

### Installing your integration

Taking as example the previously generated `mysql` sample, you must:

1. Compile and build your integration binary file: `make all`
2. In the Infrastructure Agent's host:
    * Copy the `myorg-mysql-config.yml` into one of the following folders:
        - Linux: `/etc/newrelic-infra/integrations.d`
        - Windows: `C:\Program Files\New Relic\newrelic-infra\integrations.d`
    * Copy the `myorg-mysql-definition.yml` file and the generated `bin/` folder (which contains the compiled
       `myorg-mysql` integration binary) into one of the following folders:
        - Linux: `/var/db/newrelic-infra/custom-integrations`
        - Windows: `C:\Program Files\New Relic\newrelic-infra\custom-integrations`
3. Restart the New Relic Infrastructure Agent service in the host where you installed the integration.       

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

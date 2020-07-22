[![Community Project header](https://github.com/newrelic/opensource-website/raw/master/src/images/categories/Community_Project.png)](https://opensource.newrelic.com/oss-category/#community-project)

# New Relic infrastructure integrations builder

You can use this command-line tool to create and scaffold a new integration in Golang for the [New Relic infrastructure agent](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure). 

The generated code uses the [New Relic infrastructure integrations go SDK](https://github.com/newrelic/infra-integrations-sdk). Visit the [go SDK documentation](https://github.com/newrelic/infra-integrations-sdk/blob/master/docs/README.md) for more information about its API and structure.

Before starting to write Go code, we suggest taking a look at Golang's documentation to set up the environment and familiarize yourself with the Golang language. 

## Installing

### Prerequisites

- [Vendor Tool for Go](https://github.com/kardianos/govendor), which is used for managing dependencies. The project that you create must be within a `$GOPATH/src`, otherwise the Vendor Tool won't work properly.

#### Windows

You need a Bash environment (for example, [Git Bash for Windows](https://git-scm.com/download/win)), with `make` and `awk` commands installed. 

### Installation

To install the command-line tool, use `go get`:

```bash
$ go get github.com/newrelic/nr-integrations-builder
```

After this step, the `nr-integrations-builder` binary should be installed in `$GOPATH/bin` or `$GOBIN` (if this last one is set). In a common Golang workspace,
it's recommended having `$GOPATH/bin` or `$GOBIN` included in your `$PATH`; that way, you would have `nr-integrations-builder` available globally.

## Usage

Currently, there is only one command available: `init`. You can check out the available options for this command using the `--help` option:

```bash
$ nr-integrations-builder init --help
```

Available options are:

  * `-n` or `--company-name` (string): company name (required)
  * `-c` or `--company-prefix` (string): company prefix identifier (required)
  * `-p` or `--destination-path` (string): the destination path where your integration will be generated as a subfolder
    (default: current directory)
  * `-e` or `--entity-type` (string): type of entity to generate (`remote` (default) or `local`). To know the differences
    between `remote` or `local` entities, refer to the [Entity definition section in the GoSDK v3 documents](https://github.com/newrelic/infra-integrations-sdk/blob/master/docs/entity-definition.md).

`company-name` and `company-prefix` flags must be specified. Otherwise, the `nr-integrations-builder` will not initialize the integration.

Here's an example of how to use the tool:

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

The above command would generate the following file structure in the `$GOPATH/src/myorg-integrations/` folder:

```
myorg-integrations
└── mysql
    ├── CHANGELOG.md
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

* A `Makefile` with the basic commands to verify, build and test the integration code. Run `make all` to generate an integration binary in the `bin/` subfolder.
* An `src/` folder with the basic integration main and test file. The generated main file creates a dummy integration with fake data.
* Basic config and definition `.yml` files. For more information about configuration and description files, refer to the [Integrations SDK documentation](https://docs.newrelic.com/docs/integrations/integrations-sdk/file-specifications).
* `vendor` folder and `vendor/vendor.json` file with all the vendorized third-party libraries. For more information about how to install and use
  Go Vendor, refer to [the Vendor Tool for Go docs](https://github.com/kardianos/govendor).
* Project description files: `README.md` and `CHANGELOG.md`

### Installing your integration

Continuing with the previously generated `mysql` example, you must:

1. Compile and build your integration binary file: `make all`
2. In the Infrastructure Agent's host:
    * Copy the `myorg-mysql-config.yml` into one of the following folders:
        - Linux: `/etc/newrelic-infra/integrations.d`
        - Windows: `C:\Program Files\New Relic\newrelic-infra\integrations.d`
    * Copy the `myorg-mysql-definition.yml` file and the generated `bin/` folder (which contains the compiled `myorg-mysql` integration binary) into one of the following folders:
        - Linux: `/var/db/newrelic-infra/custom-integrations`
        - Windows: `C:\Program Files\New Relic\newrelic-infra\custom-integrations`
3. Restart the New Relic Infrastructure Agent service in the host where you installed the integration.       

## Testing

The command-line tool includes a suite of unit tests with each package which should be used to verify your changes don't break existing functionality.

### Running Tests

To run the tests, invoke:

```bash
$ make test
```

### Writing tests

For most contributions it is strongly recommended to add additional tests which exercise your changes. This helps us efficiently incorporate your changes into our mainline codebase and provides a safeguard that your change won't be broken by future development.

There are some rare cases where code changes do not result in changed functionality (for example, a performance optimization) and new tests are not required.
In general, including tests with your pull request dramatically increases the chances that it will be accepted.

## Support

New Relic hosts and moderates an online forum where customers can interact with New Relic employees as well as other customers to get help and share best practices. Like all official New Relic open source projects, there's a related Community topic in the New Relic Explorers Hub. You can find this project's topic/threads here:

https://discuss.newrelic.com/c/support-products-agents/new-relic-infrastructure

## Contributing
We encourage your contributions to improve [project name]! Keep in mind when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project.
If you have any questions, or to execute our corporate CLA, required if your contribution is on behalf of a company,  please drop us an email at opensource@newrelic.com.

IMPORTANT: Never edit manually the `scaffold/bindata.go file`. This must be generated using `make generate`

## License
New Relic infrastructure integrations builder is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.

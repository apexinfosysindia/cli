# ApexOS CLI

Command line interface to facilitate interaction with the ApexOS Supervisor.

## Usage

- `apex help`
- `apex <subcommand> <action> [<options>]`

E.g.:

- `apex core info --raw-json`

### Modifiers

#### Global

```text
      --api-token string   ApexOS Supervisor API token
      --config string      Optional config file (default is $HOME/.apexos.yaml)
      --endpoint string    Endpoint for ApexOS Supervisor (default is 'supervisor')
  -h, --help               help for apex
      --log-level string   Log level (defaults to Warn)
      --no-progress        Disable the progress spinner
      --raw-json           Output raw JSON from the API
```

All options are also available as `SUPERVISOR_` prefixed environment variables like `SUPERVISOR_LOG_LEVEL`

#### Subcommands

Available commands:

```text
  apps           Install, update, remove and configure ApexOS apps
  audio          Audio device handling.
  authentication Authentication for ApexOS users.
  cli            Get information, update or configure the ApexOS cli backend
  core           Provides control of the ApexOS Core
  dns            Get information, update or configure the ApexOS DNS server
  docker         Docker backend specific for info and OCI configuration
  hardware       Provides hardware information about your system
  help           Help about any command
  host           Control the host/system that ApexOS is running on
  info           Provides a general ApexOS information overview
  multicast      Get information, update or configure the ApexOS Multicast
  network        Network specific for updating, info and configuration imports
  observer       Get information, update or configure the ApexOS observer
  os             Operating System specific for updating, info and configuration imports
  resolution     Resolution center of Supervisor, show issues and suggest solutions
  backups        Create, restore and remove backups
  supervisor     Monitor, control and configure the ApexOS Supervisor
```

## Installation

The CLI is provided by the CLI container on ApexOS systems and is
available on the device terminal when using the ApexOS Operating System.

The CLI is automatically updated on those systems.

Furthermore, the SSH app (available in the app store) provides
access to this tool and several community apps provide it as well (e.g.,
the Visual Studio Code app).

## Developing & contributing

### Prerequisites

The CLI can interact remotely with the ApexOS Supervisor using the
`remote_api` app from the [developer app repository](https://github.com/apexinfosysindia/addons-development).

After installing and starting the app, a token is shown in the `remote_api`
app log, which is needed for further development.

### Get the source code

Fork ([https://github.com/apexinfosysindia/cli/fork](https://github.com/apexinfosysindia/cli/fork)) or clone this repository.

### Using it in development

```shell
export SUPERVISOR_ENDPOINT=http://192.168.1.2
export SUPERVISOR_API_TOKEN=replace_this_with_remote_api_token
go run main.go info
```

**Note**: Replace the `192.168.1.2` with the IP address of your ApexOS
instance running the `remote_api` app and use the token provided.

### Building

We use go modules; an example build below:

```bash
CGO_ENABLED=0 go build -ldflags="-s -w" -o "apex"
```

For details how we build cross for different architectures,
please see our [build action file](https://github.com/apexinfosysindia/cli/blob/master/.github/workflows/build.yml).

### Contributing a change

1. Create a feature branch on your fork/clone of the git repository.
2. Commit your changes.
3. Rebase your local changes against the `master` branch.
4. Run test suite with the `go test ./...` command and confirm that it passes.
5. Run `gofmt -s` to ensure your code is formatted properly.
6. Create a new Pull Request.

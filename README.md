# hava CLI
hava is a Go CLI tool to interact with app.hava.io platform. 


## Overview

This tool is a CLI driven binary to interface with the Hava API/SaaS platform. This will allow practitioners to download a single binary that can be and used as part of CI/CD pipelines in tools like Github Actions, Gitlab CICD and Azure Devops, as well as part of local users commands or scripts.

This tool utilises [Hava's Go SDK](https://pkg.go.dev/github.com/teamhava/hava-sdk-go) current avaialable endpoints and with ongoing develop will support more features available via the [Hava UI](https://app.hava.io/). 

Check out our [current roadmap and drivers](./ROADMAP.md) for Hava CLI.

- [hava CLI](#hava-cli)
  - [Overview](#overview)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Print hava CLI help](#print-hava-cli-help)
    - [Create and List Sources](#create-and-list-sources)
  - [Pre-Requisites](#pre-requisites)
    - [Environment Variable](#environment-variable)
    - [Configuration File](#configuration-file)
    - [Precedence](#precedence)
    - [Running in Automation/CICD](#running-in-automationcicd)
  - [Output Formats](#output-formats)
    - [Table Output](#table-output)
    - [Json Output](#json-output)
- [Source Commands](#source-commands)
  - [Source List all](#source-list-all)
  - [Source Create AWS (Using Access Keys)](#source-create-aws-using-access-keys)
- [Build Local Binary](#build-local-binary)
  - [Testing Locally](#testing-locally)

## Installation

Binaries are created as part of a release, check out the [Release Page](https://github.com/teamhava/hava-ui-cli/releases) for the latest version.

**MacOs Installation Homebrew**
```sh
brew tap teamhava/hava
brew install hava
```


**Linux Installation**
```sh
version="x.x.x"
arch="arm64"
curl -L -o hava.zip "https://github.com/teamhava/hava-ui-cli/releases/download/${version}/hava_Linux_${arch}.zip"
unzip hava.zip
```

Specific [OS (Linux|OSX|Windows|Docker) installation here](./docs/installation.md). 


## Usage

### Print hava CLI help

![print help](./docs/images/hava_help.cast.gif)


`hava -h`

```bash
A CLI to interface with the Hava platform.

Hava CLI empowers engineers the ability to automate and interact
with the Hava platform.

Usage:
  hava [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  source      Create/List/Delete/Update sources to Hava

Flags:
      --autoapprove   Auto approve the hava command. --autoapprove, false by default
  -h, --help          help for hava
  -v, --version       version for hava

Use "hava [command] --help" for more information about a command.
```

### Create and List Sources

![hava_create_list](./docs/images/hava_create_list_sources.cast.gif)

More [Usage examples](./docs/usage.md).

## Pre-Requisites

### Environment Variable

`HAVA_TOKEN` (Required) and `HAVA_ENDPOINT` (Optional) to be set as an environment variable. This is the preferred method when working with the hava cli tool. 

### Configuration File

The following default file paths `$HOME/.hava.yaml`, `./.hava.yaml` are yaml config files that can be set as follows:

```yaml
---
hava_token: <token from hava website> ## Required from https://app.hava.io/<account>/tokens
hava_endpoint: https://havaapi.company.com ## Optional if using self hosted Hava platform, default https://api.hava.io
```

`hava` CLI by default will look in the default file paths for the configuration file.
The config file can also be specified by using the `--config <filename>` flag as follows:

```bash
hava --config <location-to-config-file> source list
```

If no environment variable or config file found AND not running in automation, user will be met with an [interactive configuration file setup](./docs/usage.md#cli-interactive-config-setup). setup.

### Precedence

`hava` CLI will use the following precedence when determining which variables to utilise:

- environment variable
- config file
- default

### Running in Automation/CICD

When running `hava` CLI in automation or a CICD pipeline, we recommend export/setting `AUTOMATION=1` as an environment variable or `automation: 1` in the config file.

Some commands do require human inputs or can be bypassed with a flag (eg `--autoapprove`). 

Checkout our [Github CLI Test workflow](./.github/workflows/cli-test.yml) and other [CICD examples](./docs/cicd-examples/README.md) for other platforms like [Gitlab-CI](./docs/cicd-examples/gitlab/.gitlab-ci.yml) and [Azure DevOps](./docs/cicd-examples/azuredevops/azure-pipelines.yml). 


## Output Formats

Table is the default format, however there are output formats of JSON (--json), CSV (--csv) ,  Markdown (--markdown) and HTML (--html)
See [output formats for more information](./docs/outputs-format.md).

### Table Output
A table is the default format when outputting information about hava sources

```bash
hava source list                                                                                     
╭───┬───────────────┬──────────────────────────────────────┬──────────────────────┬───────────────┬────────┬─────────────────────────────────────────╮ 
│   │ DISPLAYNAME   │ ID                                   │ INFO                 │ NAME          │ STATE  │ TYPE                                    │
├───┼───────────────┼──────────────────────────────────────┼──────────────────────┼───────────────┼────────┼─────────────────────────────────────────┤
│ 1 │ dev           │ 4f14c115-3b0d-40ea-b075-6df9b2fb81c9 │ AKIAIOSFODNN7EXAMPLE │ dev           │ active │ Sources::AWS::Keys                      │
│ 2 │ GCPDevChange3 │ f2a26440-10bf-43d1-9742-8361de30590f │ credentials.json     │ GCPDevChange3 │ active │ Sources::GCP::ServiceAccountCredentials │
╰───┴───────────────┴──────────────────────────────────────┴──────────────────────┴───────────────┴────────┴─────────────────────────────────────────╯
```

### Json Output


```bash
hava source list --json | jq 
[
  {
    "DisplayName": "dev",
    "Id": "4f14c115-3b0d-40ea-b075-6df9b2fb81c9",
    "Info": "AKIAIOSFODNN7EXAMPLE",
    "Name": "dev",
    "State": "active",
    "Type": "Sources::AWS::Keys"
  },
  {
    "DisplayName": "GCPDevChange3",
    "Id": "f2a26440-10bf-43d1-9742-8361de30590f",
    "Info": "credentials.json",
    "Name": "GCPDevChange3",
    "State": "active",
    "Type": "Sources::GCP::ServiceAccountCredentials"
  }
]
```

# Source Commands

`hava source -h`

```bash
Create/List/Delete/Update sources to Hava for different providers (AWS/Azure/Google Cloud)

Usage:
  hava source [flags]
  hava source [command]

Available Commands:
  create      Create sources to Hava
  delete      Delete sources to Hava
  list        List sources of Hava
  sync        Sync sources to Hava
  update      Update sources to Hava

Flags:
      --all               hava source --all
  -h, --help              help for source
      --sourceId string   sourceId of AWS|Azure|GCP source
```

## Source List all

`hava source list`

```
╭───┬─────────────┬──────────────────────────────────────┬──────────────────────┬────────────┬────────┬────────────────────╮
│   │ DISPLAYNAME │ ID                                   │ INFO                 │ NAME       │ STATE  │ TYPE               │
├───┼─────────────┼──────────────────────────────────────┼──────────────────────┼────────────┼────────┼────────────────────┤
│ 1 │ devTestAWS  │ 8eb192e2-9beb-466b-ae14-c05fc8403cf4 │ AKIAIOSFODNN7EXAMPLE │ devTestAWS │ active │ Sources::AWS::Keys │
╰───┴─────────────┴──────────────────────────────────────┴──────────────────────┴────────────┴────────┴────────────────────╯
```

## Source Create AWS (Using Access Keys)

`hava source create aws --name dev --access-key $AWS_ACCESS_KEY_ID --secret-key $AWS_SECRET_ACCESS_KEY`

```bash
[INFO]   Created AWS Source for the following source:

╭───┬─────────────┬──────────────────────────────────────┬──────────────────────┬────────────┬────────┬────────────────────╮
│   │ DISPLAYNAME │ ID                                   │ INFO                 │ NAME       │ STATE  │ TYPE               │
├───┼─────────────┼──────────────────────────────────────┼──────────────────────┼────────────┼────────┼────────────────────┤
│ 1 │ devTestAWS  │ 040d5b5e-03b9-4343-9393-8ad0794512f4 │ AKIAIOSFODNN7EXAMPLE │ devTestAWS │ queued │ Sources::AWS::Keys │
╰───┴─────────────┴──────────────────────────────────────┴──────────────────────┴────────────┴────────┴────────────────────╯
```

More [`hava source` commands found here](./docs/source_cmds.md).



# Build Local Binary
`make local-build` will build a local binary 


## Testing Locally

```bash
go run main.go source list
╭───┬───────────────┬──────────────────────────────────────┬──────────────────────┬───────────────┬────────┬─────────────────────────────────────────╮
│   │ DISPLAYNAME   │ ID                                   │ INFO                 │ NAME          │ STATE  │ TYPE                                    │
├───┼───────────────┼──────────────────────────────────────┼──────────────────────┼───────────────┼────────┼─────────────────────────────────────────┤
│ 1 │ dev           │ 4f14c115-3b0d-40ea-b075-6df9b2fb81c9 │ AKIAIOSFODNN7EXAMPLE │ dev           │ active │ Sources::AWS::Keys                      │
│ 2 │ GCPDevChange3 │ f2a26440-10bf-43d1-9742-8361de30590f │ credentials.json     │ GCPDevChange3 │ active │ Sources::GCP::ServiceAccountCredentials │
╰───┴───────────────┴──────────────────────────────────────┴──────────────────────┴───────────────┴────────┴─────────────────────────────────────────╯
```
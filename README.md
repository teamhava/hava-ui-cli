# hava CLI
hava is a Go CLI tool to interact with app.hava.io platform. 




## Overview

This tool is a CLI driven binary to interface with the Hava API/SaaS platform. This will allow practitioners to download a single binary that can be and used as part of CI/CD pipelines in tools like Github Actions, Gitlab CICD and Azure Devops, as well as part of local users commands or scripts.

- [hava CLI](#hava-cli)
  - [Overview](#overview)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Print hava CLI help](#print-hava-cli-help)
    - [Create and List Sources](#create-and-list-sources)
    - [CLI JSON Output](#cli-json-output)
    - [Sync Sources](#sync-sources)
    - [Delete Sources](#delete-sources)
    - [CLI Completion Setup](#cli-completion-setup)
    - [CLI Interactive Config Setup](#cli-interactive-config-setup)
  - [Pre-Requisites](#pre-requisites)
    - [Environment Variable](#environment-variable)
    - [Configuration File](#configuration-file)
    - [Precedence](#precedence)
    - [Running in Automation/CICD](#running-in-automationcicd)
- [Source Commands](#source-commands)
  - [Source List all](#source-list-all)
  - [Source List SourceID](#source-list-sourceid)
  - [Source Sync SourceID](#source-sync-sourceid)
  - [Source Create AWS (Using Access Keys)](#source-create-aws-using-access-keys)
  - [Source Create AWS (Using Cross Account Role)](#source-create-aws-using-cross-account-role)
  - [Source Create GCP](#source-create-gcp)
  - [Source Create Azure](#source-create-azure)
  - [Source Delete](#source-delete)
  - [Source Update AWS|AZURE|GCP](#source-update-awsazuregcp)
  - [Output Formats](#output-formats)
    - [Table Output](#table-output)
    - [Json Output](#json-output)
    - [CSV Output](#csv-output)
    - [Markdown Output](#markdown-output)
- [Build Local Binary](#build-local-binary)
  - [Testing Locally](#testing-locally)

## Installation

Binaries are created as part of a release, check out the [Release Page](https://github.com/teamhava/hava-ui-cli/releases) for the latest version.

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

### CLI JSON Output

![hava_json_out](./docs/images/hava_CLI_json_output.cast.gif)


### Sync Sources

![hava_print_source](./docs/images/hava_source_sync.cast.gif)

### Delete Sources

![hava_delete_sources](./docs/images/hava_source_delete.cast.gif)

### CLI Completion Setup

![hava_cli_completion](./docs/images/hava_cli_completion.cast.gif)

### CLI Interactive Config Setup

![hava_config_setup](./docs/images/hava_config_setup.cast.gif)

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

### Precedence

`hava` CLI will use the following precedence when determining which variables to utilise:

- environment variable
- config file
- default

### Running in Automation/CICD

When running `hava` CLI in automation or a CICD pipeline, we recommend export/setting `AUTOMATION=1` as an environment variable or `automation: 1` in the config file.

Some commands do require human inputs or can be bypassed with a flag (eg `--autoapprove`). 

Checkout our [Github CLI Test workflow](./.github/workflows/cli-test.yml). 

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

## Source List SourceID

`hava source list --source-id a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38`


## Source Sync SourceID

`hava source sync a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38`


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

## Source Create AWS (Using Cross Account Role)

`hava source create aws --name devCAR --role-arn $AWS_CROSS_ACCOUNT_ROLE`

```bash
[INFO]   Created AWS Source for the following source:

╭───┬─────────────┬──────────────────────────────────────┬────────────────────────────────────────┬────────┬────────┬────────────────────────────────╮
│   │ DISPLAYNAME │ ID                                   │ INFO                                   │ NAME   │ STATE  │ TYPE                           │
├───┼─────────────┼──────────────────────────────────────┼────────────────────────────────────────┼────────┼────────┼────────────────────────────────┤
│ 1 │ devCAR      │ 31f9b6a6-4e48-400a-8ced-5bfc2a1aacc2 │ arn:aws:iam::123456789012:role/HavaCAR │ devCAR │ queued │ Sources::AWS::CrossAccountRole │
╰───┴─────────────┴──────────────────────────────────────┴────────────────────────────────────────┴────────┴────────┴────────────────────────────────╯
```

## Source Create GCP

`hava source create gcp  --name GCPDev --configFile $GCP_ENCODED_FILE`

```bash
[INFO]   Created GCP Source for the following source:

╭───┬─────────────┬──────────────────────────────────────┬──────────────────┬────────┬────────┬─────────────────────────────────────────╮
│   │ DISPLAYNAME │ ID                                   │ INFO             │ NAME   │ STATE  │ TYPE                                    │
├───┼─────────────┼──────────────────────────────────────┼──────────────────┼────────┼────────┼─────────────────────────────────────────┤
│ 1 │ GCPDev      │ 8818d864-46e9-4b8d-acdd-a31d8936e62f │ credentials.json │ GCPDev │ queued │ Sources::GCP::ServiceAccountCredentials │
╰───┴─────────────┴──────────────────────────────────────┴──────────────────┴────────┴────────┴─────────────────────────────────────────╯
```

## Source Create Azure

`hava create source azure --name AzureDev --client-id $ARM_CLIENT_ID --tenant-id $ARM_TENANT_ID --subscription-id $ARM_SUBSCRIPTION_ID`


## Source Delete

`hava source delete 22872411-20e8-4b6e-aa46-41866c9c1897`


## Source Update AWS|AZURE|GCP

`hava source update gcp --name GCPDevChange --source-id f2a26440-10bf-43d1-9742-8361de30590f`




## Output Formats

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

### CSV Output

```bash
hava source list --csv       
,DisplayName,Id,Info,Name,State,Type
1,dev,4f14c115-3b0d-40ea-b075-6df9b2fb81c9,AKIAIOSFODNN7EXAMPLE,dev,active,Sources::AWS::Keys
2,GCPDevChange3,f2a26440-10bf-43d1-9742-8361de30590f,credentials.json,GCPDevChange3,active,Sources::GCP::ServiceAccountCredentials
```


### Markdown Output

```bash
hava source list --markdown
| | DisplayName | Id | Info | Name | State | Type |
| ---:| --- | --- | --- | --- | --- | --- |
| 1 | dev | 4f14c115-3b0d-40ea-b075-6df9b2fb81c9 | AKIAIOSFODNN7EXAMPLE | dev | active | Sources::AWS::Keys |
| 2 | GCPDevChange3 | f2a26440-10bf-43d1-9742-8361de30590f | credentials.json | GCPDevChange3 | active | Sources::GCP::ServiceAccountCredentials |
```

# Build Local Binary
`make local-build`


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
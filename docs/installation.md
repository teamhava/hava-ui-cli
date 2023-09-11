## Installation

Binaries are created as part of a release, check out the [Release Page](https://github.com/teamhava/hava-ui-cli/releases) for the latest version.

- [Installation](#installation)
- [Pre-Requisites](#pre-requisites)
  - [Environment Variable](#environment-variable)
  - [Configuration File](#configuration-file)
  - [Precedence](#precedence)
  - [Running in Automation/CICD](#running-in-automationcicd)



**MacOs Installation amd64 OR arm64**
```sh
version="x.x.x"
curl -L -o hava.zip "https://github.com/teamhava/hava-ui-cli/releases/download/${version}/hava_Darwin_all.zip"
ditto -x -k hava.zip ./
```

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
Currently arch=["arm64, "x86_64"]

**Windows Installation**
```sh
version="x.x.x"
arch="x86_64"
curl -L -o hava.exe "https://github.com/teamhava/hava-ui-cli/releases/download/${version}/hava_Windows_${arch}.zip"
unzip hava.zip
```
Currently arch=["arm64, "x86_64"]


**Docker Container Installation**

```sh
docker pull hava/cli
docker run --env HAVA_TOKEN hava/cli hava source list
```


## Pre-Requisites

### Environment Variable

`HAVA_TOKEN` (Required) and `HAVA_ENDPOINT` (Optional) to be set as an environment variable. This is the preferred method when working with the hava cli too. 

```bash
export HAVA_TOKEN="HAVATOKEN"
export HAVA_ENDPOINT="https://api.selfhosted.io ## Optional as default is https://api.hava.io
```

### Configuration File

The following default file paths `$HOME/.hava.yaml`, `./.hava.yaml` are yaml config files that can be set as follows:

```yaml
---
hava_token: <token from hava website> ## Required
hava_endpoint: https://havaapi.company.com ## Optional if using self hosted Hava platform
```

`hava` CLI by default will look in the default locations for the configuration file.
The config file can also be specified by using the `--config <filename>` flag as follows:

```bash
hava --config <location-to-config-file> source list
```

### Precedence

`hava` CLI will use the following precedence when determining which item to take

- environment variable
- config file
- default

### Running in Automation/CICD

When running `hava` CLI in automation or a CICD pipeline, export/set `AUTOMATION=1` as an environment variable or `automation: 1` in the config file.
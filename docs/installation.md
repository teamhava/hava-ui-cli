## Installation

Binaries are created as part of a release, check out the [Release Page](https://github.com/teamhava/hava-ui-cli/releases) for the latest version.


**MacOs Installation amd64**
```sh
version="x.x.x"
curl -L -o hava "https://github.com/teamhava/hava-ui-cli/releases/download/${version}/hava-darwin-amd64"
chmod +x hava
```

**MacOs Installation arm64**
```sh
version="x.x.x"
curl -L -o hava "https://github.com/teamhava/hava-ui-cli/releases/download/${version}/hava-darwin-arm64"
chmod +x hava
```

!!! note ""
    Note: `hava` CLI is currently not developer-signed or notorised and you will run into an initial issue where `hava` is not allowed to run. Please follow "[safely open apps on your mac](https://support.apple.com/en-au/HT202491#:~:text=View%20the%20app%20security%20settings%20on%20your%20Mac&text=In%20System%20Preferences%2C%20click%20Security,%E2%80%9CAllow%20apps%20downloaded%20from.%E2%80%9D)" to allow `hava` to run

**Linux Installation**
```sh
version="x.x.x"
curl -L -o hava "https://github.com/teamhava/hava-ui-cli/releases/download/${version}/hava-linux-amd64"
chmod +x hava
```

**Windows Installation**
```sh
version="x.x.x"
curl -L -o hava.exe "https://github.com/teamhava/hava-ui-cli/releases/download/${version}/hava-windows-amd64"
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

- environmen variable
- config file
- default

### Running in Automation/CICD

When running `hava` CLI in automation or a CICD pipeline, export/set `AUTOMATION=1` as an environment variable or `automation: 1` in the config file.
## Usage
- [Usage](#usage)
  - [Print hava CLI help](#print-hava-cli-help)
  - [Create and List Sources](#create-and-list-sources)
  - [CLI JSON Output](#cli-json-output)
  - [Sync Sources](#sync-sources)
  - [Delete Sources](#delete-sources)
  - [CLI Completion Setup](#cli-completion-setup)
  - [CLI Interactive Config Setup](#cli-interactive-config-setup)

### Print hava CLI help

![print help](./images/hava_help.cast.gif)


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

![hava_create_list](./images/hava_create_list_sources.cast.gif)

### CLI JSON Output

![hava_json_out](./images/hava_CLI_json_output.cast.gif)


### Sync Sources

![hava_print_source](./images/hava_source_sync.cast.gif)

### Delete Sources

![hava_delete_sources](./images/hava_source_delete.cast.gif)

### CLI Completion Setup

![hava_cli_completion](./images/hava_cli_completion.cast.gif)

### CLI Interactive Config Setup

The following will trigger when no `HAVA_TOKEN` found as an environment variable AND Config file found AND `AUTOMATION=1` is not set. 

![hava_config_setup](./images/hava_config_setup.cast.gif)
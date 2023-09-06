# Hava CLI Roadmap & Vision

## Why Hava CLI Exists ?

- Our customer engineers are wanting an automated way to interact with the Hava Platform instead of using the UI
- Our customers engineers utilise various automation tools (eg terraform, ansible)
- We do not expect engineers to learn or interace directly with our APIs



## Hava CLI Vision

Create a single binary CLI that can be downloaded and used on the command line, either as a management tool, through scripts, or as part of an automated CI/CD pipeline.



## Hava CLI Initial Features:

- Utilise [Hava's Go SDK](https://pkg.go.dev/github.com/teamhava/hava-sdk-go) current avaialable endpoints
- [Source API Endpoint integration](https://github.com/teamhava/hava-sdk-go#documentation-for-api-endpoints)
  - `hava source list`
  - `hava source create [aws|azure|gcp]`
  - `hava source delete`
  - `hava source sync`
  - `hava source update`
- Authentication via HAVA token via a CLI config file or envrionment variable (`HAVA_TOKEN`)
  - If no credentials found, Hava CLI will go through an interactive way for user to setup credentials
- Provide engineers the ability to output results in JSON, CSV, HTML
- Code structure setup to build on future features





## Future Work and Possible Features

- Enable more [Hava API endpoints](https://app.swaggerhub.com/apis-docs/H252/hava) via [Hava's Go SDK](https://pkg.go.dev/github.com/teamhava/hava-sdk-go)
- Ability to search for cloud resources via CLI (`hava search <object>`)
- Adminstration of Hava Teams via CLI (`hava teams *`)
- Automation of Hava reports (`hava reports *`)

Any new feature suggestions or improvements, [click here to discuss](https://github.com/teamhava/hava-ui-cli/issues/new/choose).
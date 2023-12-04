# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## 0.1.2 - 2023-12-04

### Changelog

* 4bf3b0d Bump docker/login-action from 2 to 3
* 2c3c0e7 Bump docker/setup-qemu-action from 2 to 3
* fa58a50 Bump github.com/jedib0t/go-pretty/v6 from 6.4.7 to 6.4.9
* 26965e4 Bump github.com/teamhava/hava-sdk-go from 0.2.1 to 1.1.3
* 77d1af4 Bump golang.org/x/term from 0.12.0 to 0.15.0
* 9f1be5a Bump goreleaser/goreleaser-action from 4 to 5
* fc1e77c Bump stefanzweifel/git-auto-commit-action from 4 to 5
* 15762ad Merge branch 'main' into dependabot/go_modules/github.com/teamhava/hava-sdk-go-1.1.3
* 249215c Merge pull request #15 from teamhava/dependabot/github_actions/goreleaser/goreleaser-action-5
* 6e53f2b Merge pull request #16 from teamhava/dependabot/github_actions/docker/setup-qemu-action-3
* 20bf9dc Merge pull request #17 from teamhava/dependabot/github_actions/docker/login-action-3
* 9d0cc30 Merge pull request #21 from teamhava/dependabot/github_actions/stefanzweifel/git-auto-commit-action-5
* c95b338 Merge pull request #23 from teamhava/dependabot/go_modules/github.com/teamhava/hava-sdk-go-1.1.3
* 2e3a54a Merge pull request #24 from teamhava/dependabot/go_modules/github.com/jedib0t/go-pretty/v6-6.4.9
* e3a836e Merge pull request #27 from teamhava/dependabot/go_modules/golang.org/x/term-0.15.0

## 0.0.26-beta - 2023-09-05

### Changelog

- 7e02cae Merge branch 'main' into repo-standards
- 31a6d29 Merge pull request #9 from teamhava/repo-standards
- 997b97c Release and contributing guides
- a62f217 Remove second "setup go" step
- 6e8926c Update contributing guide

## 0.0.25-beta - 2023-09-05

### Changelog

- 7e02cae Merge branch 'main' into repo-standards
- 31a6d29 Merge pull request #9 from teamhava/repo-standards
- 997b97c Release and contributing guides
- a62f217 Remove second "setup go" step
- 6e8926c Update contributing guide

## 0.0.24-pre-alpha - 2023-09-04

### Changelog

- 5b43908 Update CHANGELOG
- 78f0649 Update issue templates
- 8004cf0 Update release job name
- 726cd76 move brew to macos stage

## 0.0.23-pre-alpha - 2023-09-04

### Changelog

- 524b89f - Build/release, Sign Darwin binary, followed by windows/linux - 2 goreleaser configs
- 6e98bc1 - Get rid of gon.hcl - Update trigger for main workflow
- e303b60 - change example folder to cicd-example - update github actions workflow
- c18f4bf - github actions workflow - gitlab ci example
- a6f691f Add azureDevOps pipeline and readme files
- e8869be Merge branch 'main' into cicd-examples
- 66679a4 Merge pull request #6 from teamhava/cicd-examples
- 8dd9331 Remove AWS access IDs and typos
- 9fe1c53 Update CHANGELOG
- abf670f Update README
- 715d692 Update README with correct links to release and instructions for install
- f2ee63a Update github actions cicd demo
- 5b9ff59 github actions examples

## 0.0.22-pre-alpha - 2023-09-01

### Changelog

- bb60d37 Bump github.com/jedib0t/go-pretty/v6 from 6.4.6 to 6.4.7
- 567cc38 Enable brew publishing
- 993501a Merge pull request #4 from teamhava/dependabot/go_modules/github.com/jedib0t/go-pretty/v6-6.4.7

## 0.0.5-pre-alpha - 2023-08-24

- Cleaned up sync|delete commands

## [Unreleased]

- kubernetes source create|update

## 0.0.4-pre-alpha - 2023-08-23

### Added

- hava source [create|delete|list|sync|update] [aws|azure|gcp]
- config file generator

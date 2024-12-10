# UDS Package Kit

## Overview

UDS Package Kit is a tool designed to assist in developing, maintaining, and publishing UDS Packages.

## Features

- Automated release and tag creation in GitLab and GitHub
- Customizable release configuration file

## Installation

Download the latest UDS Package Kit binaries from the [GitHub Releases](https://github.com/defenseunicorns/uds-pk/releases) page.

## Usage

After installation, you can use uds-pk via the command line:

> [!TIP]
> To view available commands run `uds-pk help`

## Release Example

Pseudo flow for CI/CD:

```bash
uds-pk release check <flavor>

uds-pk release update-yaml <flavor>

# publish the package #

uds-pk release <platform> <flavor>
```

### Gitlab

When running `uds-pk release gitlab <flavor>` you are expected to have an environment variable set to a GitLab token that has write permissions for your current project. This defaults to `GITLAB_RELEASE_TOKEN` but can be changed with the `--token-var-name` flag.

### GitHub

When running `uds-pk release github <flavor>` you are expected to have an environment variable set to a GitHub token that has write permissions for your current project. This defaults to `GITHUB_TOKEN` but can be changed with the `--token-var-name` flag.

### Release Configuration

UDS Package Kit release commands can be configured using a YAML file named releaser.yaml in your project's root directory.

```yaml
flavors:
  - name: upstream
    version: "1.0.0-uds.0"
  - name: registry1
    version: "2.0.0-uds.0"
  - name: unicorn
    version: "1.0.0-uds.0"
```

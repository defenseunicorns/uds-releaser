# UDS-Releaser

## Overview

UDS-Releaser is a tool designed to assist in publishing releases for UDS Packages. It automates the process of creating releases and tagging versions for each flavor.

## Features

- Automated release and tag creation
- GitLab Integration (More to come)
- Customizable release configuration file

## Installation

Download the latest UDS-Releaser binaries from the [GitHub Releases](https://github.com/defenseunicorns/uds-releaser/releases) page.

## Usage

After installation, you can use uds-releaser via the command line:

### Commands

- `check`: Check if release is necessary for a given flavor
- `release`: Create a new release on the specified platform for a given flavor
- `show`: Print the current version of the specified flavor
- `update-yaml`: Update the version fields in the zarf and uds-cli yaml files based on flavor

### Example

Pseudo flow for CI/CD:

```bash
uds-releaser check <flavor>

uds-releaser update-yaml <flavor>

# publish the package #

uds-releaser release <flavor>
```

## Configuration

UDS-Releaser can be configured using a YAML file named uds-releaser.yml in your project's root directory.

```yaml
flavors:
  - name: upstream
    version: "1.0.0-uds.0"
  - name: registry1
    version: "2.0.0-uds.0"
  - name: unicorn
    version: "1.0.0-uds.0"
```

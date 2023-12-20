 # Pokta - Okta interactions utilities <!-- omit in toc -->
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/piaverous/pokta)](https://goreportcard.com/report/github.com/piaverous/pokta)
![ci-status](https://github.com/piaverous/pokta/actions/workflows/ci.yml/badge.svg)

Interacting with Okta as a developer is not always the most straightforward.
`pokta` aims to bring simple features to your CLI, to make your life easier.

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Currently supported](#currently-supported)


## Installation

Follow any of the instructions below.

### Mac OS <!-- omit in toc -->

```bash
brew install piaverous/tap/pokta
```
This formula is maintained by me. It is updated automatically after every release.

### General Linux <!-- omit in toc -->

```bash
curl -L https://raw.githubusercontent.com/piaverous/pokta/main/install.sh | bash
```
This script installs the latest release by default.

## Getting Started

Setup a config file in `~/.pokta/config.yml` like so : 

```yaml
okta:
  api_key: <your_client_api_key>
  aud: https://<your_okta_domain_base_url>/token
  client_id: <your_client_id>
  kid: <your_key_id>
  private_key: <your_base64_encoded_p12_key>
  scope: <scopes_you_wish_to_request_in_your_token>
```

Then simply run :

```bash
pokta auth pkjwt
```

## Currently supported

Currently, this CLI only supports Private Key JWT login.


> [!WARNING]
> My understanding of the topic is not very exhaustive, and this CLI is essentially used as a debuging tool. Therefore : 
> - Some things may be inexact and not work properly
> - No proper error management has been implemented yet

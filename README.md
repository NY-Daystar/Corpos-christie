# Income tax calculator

[![ci](https://github.com/LucasNoga/corpos-christie/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/LucasNoga/corpos-christie/actions)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Version](https://img.shields.io/github/tag/LucasNoga/corpos-christie.svg)](https://github.com/LucasNoga/corpos-christie/releases)
[![Total views](https://img.shields.io/sourcegraph/rrc/github.com/LucasNoga/corpos-christie.svg)](https://sourcegraph.com/github.com/LucasNoga/corpos-christie)
[![Godoc](https://godoc.org/github.com/LucasNoga/corpos-christie?status.svg)](https://godoc.org/github.com/LucasNoga/corpos-christie)

This project has been developped in Golang allows to calculate your taxes in the current year.

The government has created an explanatory sheet to understand the calculation of the tax rate but this calculation is relatively complex and we want to create a simpler interface to calculate things.  
Here's the sheet: https://www.economie.gouv.fr/particuliers/tranches-imposition-impot-revenu#etapescalculir

## Table of contents

- [Requirements](#requirements)
- [How to launch program](#how-to-launch-program)
- [For Developpers](#for-developpers)
- [Docker](#docker)
- [Installing and Setup Golang](#installing-and-setup-golang)
- [Suggestions](#suggestions)
- [Credits](#credits)

## Requirements

- [Golang](https://golang.org/dl/) >= 1.18.3

## How to launch program

1. Get program  
    1.1 Linux

   ```bash
   $ wget https://github.com/LucasNoga/corpos-christie/releases/download/v1.1.0/linux-corpos-christie-1.1.0.zip -O corpos-christie.zip
   ```

   1.2 Windows

   ```bash
   $ wget https://github.com/LucasNoga/corpos-christie/releases/download/v1.1.0/windows-corpos-christie-1.1.0.zip -O corpos-christie.zip
   ```

   1.3 Mac

   ```bash
   $ wget https://github.com/LucasNoga/corpos-christie/releases/download/v1.1.0/mac-corpos-christie-1.1.0.zip -O corpos-christie.zip
   ```

2. Unzip it

```bash
$ unzip corpos-christie.zip -d corpos-christie
```

3. Start program

```bash
$ cd corpos-christie
$ ./corpos-christie
```

## For Developpers

Clone th repository

```bash
$ git clone https://github.com/lucasnoga/corpos-christie.git
```

Launch program basically

```bash
$ go run main.go
```

or

```
$ go build
$ ./corpos-christie
```

Import module for an other project

```bash
go get github.com/LucasNoga/corpos-christie
```

Launch console mode

```bash
$ make runconsole
```

To build program

```bash
$ make
```

To launch tests

```bash
$ make test
```

To import modules

```bash
$ go mod init corpos-christie
$ go mod download
```

To see go doc (ex: tax package)

```bash
$ go doc github.com/LucasNoga/corpos-christie/tax
```

See [Project dependencies](https://deps.dev/go/github.com/lucasnoga/corpos-christie) To watch go project used in this program

## Docker

The official image of this projet: https://hub.docker.com/r/lucasnoga/corpos-christie

To use this image

```bash
$ docker pull lucasnoga/corpos-christie
$ docker run -it --rm --name corpos-christie lucasnoga/corpos-christie
```

To build a new docker image

```bash
$ make docker-build
```

To run this image

```bash
$ make docker-run
```

## Installing and Setup Golang

To install golang

```bash
$ wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
$ tar -xvf go1.16.5.linux-amd64.tar.gz
$ sudo mv go /usr/lib
$ go version
```

## Suggestions

- To make a pull request: https://github.com/LucasNoga/corpos-christie/pulls
- To summon an issue: https://github.com/LucasNoga/corpos-christie/issues
- For any specific demand by mail: luc4snoga@gmail.com

## Credits

Made by Lucas Noga.  
Licensed under GPLv3.

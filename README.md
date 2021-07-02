# Income tax calculator

[![ci](https://github.com/LucasNoga/corpos-christie/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/LucasNoga/corpos-christie/actions)
[![Total views](https://img.shields.io/sourcegraph/rrc/github.com/LucasNoga/corpos-christie.svg)](https://sourcegraph.com/github.com/LucasNoga/corpos-christie)
[![Godoc](https://godoc.org/github.com/LucasNoga/corpos-christie?status.svg)](https://godoc.org/github.com/LucasNoga/corpos-christie)

This project developped in Golang allows to calculate your taxes in the current year.  

The government has created an explanatory sheet to understand the calculation of the tax rate but this calculation is relatively complex and we want to create a simpler interface to calculate things.  
Here's the sheet: https://www.economie.gouv.fr/particuliers/tranches-imposition-impot-revenu#etapescalculir  

Installation package
```bash
$ go get go get github.com/LucasNoga/corpos-christie
```


## Table of contents
- [Requirements](#requirements)
- [Get started](#get-started)
- [Configuration file](#configuration-file-configjson)
- [Installing and Setup Golang](#installing-and-setup-golang)
- [Suggestions](#suggestions)
- [Credits](#credits)

### Version 0.0.8

## How to launch program
1. Get program
```bash
$ wget https://github.com/LucasNoga/corpos-christie/releases/download/v0.0.7/corpos-christie-0.0.8.zip
```

2. Unzip it
```bash
$ unzip corpos-christie-0.0.8.zip -d corpos-christie
```

3. Start program
```bash
$ cd corpos-christie
$ ./corpos-christie
```

## Requirements
- [Golang](https://golang.org/dl/) >= 1.16.4

## Developpers Get started
Launch program basically
```bash
$ go run main.go
```
or
```
$ go build
$ ./corpos-christie
```

Launch console mode
```bash
$ make runconsole
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

## Configuration file (config.json)
```js
{
    "tax": [ // Tax options
        {
            "year": 2021, // Year of the tax specifications
            "tranches": [
            // Tranches list see this document to understand https://www.economie.gouv.fr/particuliers/tranches-imposition-impot-revenu#etapescalculir      
                {
                    "min": 0, // Minimun in euros to get in the tranche
                    "max": 10084, // Maximum in euros to get in the tranche
                    "percentage": 0 // Percentage taxable in euros in this tranche
                },
                {
                    "min": 10085,
                    "max": 25710,
                    "percentage": 11
                },
                {
                    "min": 25711,
                    "max": 73516,
                    "percentage": 30
                }
            ]
        }
    ]
}
```

## Installing and Setup Golang
To install golang
```bash
$ wget https://golang.org/dl/go1.16.4.linux-amd64.tar.gz
$ tar -xvf go1.16.4.linux-amd64.tar.gz
$ sudo mv go /usr/lib
$ go version
```

## Suggestions
- To make a pull request: https://github.com/LucasNoga/corpos-christie/pulls
- To summon an issue: https://github.com/LucasNoga/corpos-christie/issues
- For any specific demand by mail: luc4snoga@gmail.com

## TODO
- ~~Starting project~~ - `done`
- ~~Tax calculator v1~~ - `done`
- ~~Tax calculator v2~~ - `done`
- ~~Display tax tranches~~ - `done`
- Add GoDoc
- Portability tests
- Docker
- Tax calculator v3
- Tax calculator v4
- Features command line management
- GUI


## Credits
Made by Lucas Noga.  
Licensed under GPLv3.
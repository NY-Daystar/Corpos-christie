# Income tax calculator

[![ci](https://github.com/LucasNoga/tablewriter/corpos-christie/ci/badge.svg?branch=master)](https://github.com/LucasNoga/corpos-christie/actions)
[![Total views](https://img.shields.io/sourcegraph/rrc/github.com/LucasNoga/corpos-christie.svg)](https://sourcegraph.com/github.com/LucasNoga/corpos-christie)
[![Godoc](https://godoc.org/github.com/LucasNoga/corpos-christie?status.svg)](https://godoc.org/github.com/LucasNoga/corpos-christie)

This project developped in Golang allows to calculate your taxes in the current year.  

The government has created an explanatory sheet to understand the calculation of the tax rate but this calculation is relatively complex and we want to create a simpler interface to calculate things.  
Here's the sheet: https://www.economie.gouv.fr/particuliers/tranches-imposition-impot-revenu#etapescalculir

## Table of contents
- [Requirements](#requirements)
- [Get started](#get-started)
- [Configuration file](#configuration-file-configjson)
- [Installing and Setup Golang](#installing-and-setup-golang)
- [Suggestions](#suggestions)
- [Credits](#credits)

### Version 0.0.5

## Requirements
- [Golang](https://golang.org/dl/) >= 1.16.4

## Get started
```bash
$ go run main.go
```
or
```
$ go build
$ ./corpos-christie
```

To launch tests
```bash
$ go test ./...
```

To import modules
```bash
$ go mod init corpos-christie
$ go mod download
```

## Configuration file (config.json)
```js
{
    "name": "Corpos-Christie",          // Name of the project
    "version": "X.X.X",                 // Version of the project
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
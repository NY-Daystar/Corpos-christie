# Income tax calculator

This project developped in Golang allows to calculate your taxes in the current year.  

The government has created an explanatory sheet to understand the calculation of the tax rate but this calculation is relatively complex and we want to create a simpler interface to calculate things.  
Here's the sheet: https://www.economie.gouv.fr/particuliers/tranches-imposition-impot-revenu#etapescalculir

### Version 1.0.0

## Golang version >= 1.16.4

## Get started
```bash
$ go run main.go
```
or
```
$ go build
# ./corpos-christie
```

To import the module
```bash
go mod init corpos-christie
go mod download
```

## Install Golang

To install golang
```bash
$ wget https://golang.org/dl/go1.16.4.linux-amd64.tar.gz
$ tar -xvf go1.16.4.linux-amd64.tar.gz
ls
$ sudo mv go /usr/lib
$ go version
```
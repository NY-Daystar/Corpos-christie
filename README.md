# Income tax calculator

This project developped in Golang allows to calculate your taxes in the current year.  

The government has created an explanatory sheet to understand the calculation of the tax rate but this calculation is relatively complex and we want to create a simpler interface to calculate things.  
Here's the sheet: https://www.economie.gouv.fr/particuliers/tranches-imposition-impot-revenu#etapescalculir

### Version 0.0.2

## Requirements
- [Golang](https://golang.org/dl/) >= 1.16.4

## Get started
```bash
$ go run main.go
```
or
```
$ go build
# ./corpos-christie
```

To import modules
```bash
go mod init corpos-christie
go mod download
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
{
  "cron": "5min", 
  "blockType": ["block", "uncle", "stale"], // list of different block
  "blockConfirmations": 10, // Range of block to not processing, therefore we wait consensus
  "defaultGasUsed": 21000, // Default Gas for ETH transaction before the RPC transactionReceipt was created
  "nodes": [
    {
      "coin": "ETC", // coin of the node
      "enabled": true, // active so we collect stats
      "cronBlock": true, // active the cron to fetch blocks data
      "cronTX": true, // active the cron to fetch transactions data
      "host": "51.38.180.231", // ip of the node
      "port": 8546, // port of the node
      "wallet": "0x8ccfe15255cddcd20fc667fc508936c74e91e5c1", // wallet to check block and uncle mined
      "era": 5000000,
      "baseReward": 5, // base reward for etc block
      "uncleRatioReward": "3.125%", // 3.125% of uncle value compare to the block reward
      "fullTxSync": true // go find all transactions since block 0
    },
    {
      "coin": "eth",
      "enabled": true,
      "host": "51.38.180.231",
      "port": 8545,
      "wallet": "0x249bdb4499bd7c683664C149276C1D86108E2137",
      "uncleRatioReward": "12.5%",
      "fullTxSync": true,
      // reward list with different fork
      "rewards":[
          {
              "blockno": 4370000, // Base Reward 5 until block 4370000, then the Byzantium fork in 2017
              "reward": 5
          },
          {
              "blockno": 7280000, // Base Reward 3 until block 7280000, then Constantinople fork in 2019
              "reward": 3
          },
          {
              "blockno": "Unknown", // Base Reward 2 until now
              "reward": 2
          },
      ],
    },
    {
      "coin": "rvn",
      "enabled": true,
      "cronBlock": true,
      "cronTX": false,
      "host": "51.38.180.231",
      "port": 8766,
      "baseReward": 5000, // Base reward for ravencoin
      "rpcuser": "ravencoin_beta", // RPC authentication user to get the access of the data
      "rpcpassword": "vHX11g8eJ2", // RPC auth password to get the access of the data
      "wallet": "RRH8mYJUdA2U26wsRCqYby4wCtUbXfBkEC" // Cruxpool's mining wallet
    }
  ],

  "uncleRatio": {
    "node": "etc", // node to get Uncle ratio of the network
    "active": true, // if true it will disable the explorer to get only the uncle ratio
    "blocksNumber": 500 // last N block to check
  },

  "detectionBlock": {
    "active": true // if true it will disable the explorer to get only the detectionBlock
  },

// Maria DB Instance
  "db": {
    "host": "127.0.0.1",
    "user": "user_stats",
    "db": "stats_db",
    "password": "user_password",
    "port": 3316
  }
}
```


## Installing and Setup Golang
To install golang
```bash
$ wget https://golang.org/dl/go1.16.4.linux-amd64.tar.gz
$ tar -xvf go1.16.4.linux-amd64.tar.gz
ls
$ sudo mv go /usr/lib
$ go version
```
# Caching proxy for `eth_getBlockByNumber`

## Content

- [Description](#description)
- [Configuration](#configuration)
- [Usage](#usage)
    - [Docker](#run-docker-compose)
    - [Build](#build)
    - [Tests](#tests)
    - [Linter](#linter)
    - [Clean](#clean)
- [Endpoints](#endpoints)
- [Returned errors](#returned-http-error-codes)
- [TODOs](#todos)


## Description
`eth-caching-proxy` is a package that provides LRU caching proxy for API method `eth_getBlockByNumber` using Cloudflare Ethereum Gateway.
- It uses simple [golang-lru](https://github.com/hashicorp/golang-lru) as a cache library with support of concurrent reads/writes
- Configuration via `yaml` and `.env` files using [Viper](https://github.com/spf13/viper)
- Latest 20 blocks are not cached
- Supports docker-compose
- Supports Linter
- Made with controller -> service -> repository pattern
- Uses [Gin](https://github.com/gin-gonic/gin) as a REST API framework

## Configuration

You should modify `config/config.yaml`

```
server:
  runmode: "debug" # or "release"
  httpport: 8000
  readtimeout: 60
  writetimeout: 60

cache:
  maxblocks: 128

cloudflare:
  url: "https://cloudflare-eth.com"
```

## Usage
### Run docker-compose
```bash
make docker
```

### Build
```bash
make build
```

### Tests
```bash
make test
```

### Linter
```bash
make lint
```

### Clean
```bash
make clean
```


# Endpoints

#### /block/:blockNumber
* `GET` : Get a block [has either "latest" or a number as a param.]

*Response:*

```js
HTTP 200 OK

{
    "number": 13593280, 
    "header": {
        "parentHash": "0x855637ad8b0c7ed477460118f5db1be30f6d15518a6e5c8d44d6db49aff31591",
        "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
        "miner": "0xc93452a74e596e81e4f73ca1acff532089ad4c62",
        "stateRoot": "0xc70b69c817b8761ad9c564a45c8b9281e8527a2565c3360d68e7be6a722541ac",
        "transactionsRoot": "0x786d6d4f1236db5dd9dbcd249853f251530ec52dc0ad79b097bb7c47d0f4b33a",
        "receiptsRoot": "0x83fcf0376a111b51881f347a23f5a97fcbd837637f95ba5e0c7ce66d90ef4461",
        "logsBloom": "0x75b2597752a9dfdf77ecb7de9edd3b7f89e2f1b79e47daad1f87e7777f96d61f948fd74f4cfce68edd5f59c2f3a9c517efd794d77c7effb9ff7d9d2c5ffff8f3e2d496d8ffffbf29dfc272febaf5d7e7b2fff0cfed6f905325713543eec7d557fffeb71d3fe7ea5b7ffcfbf540fd7bf9dffbf5f49ed7b572fa3edfb3e93ceb359eb5f7ebf7e2ff3ed8fe0cf33bda06355ca09d8d6fb7ae7a7fb5f85fdfb745681f7dfc7b13efed732fadf6dbfd830ff3febdf279f4d6ed66abebfffe3a6e58dee15e331fdb6d5c73afb74fd777dddb1e6f34b61f1d2ee4be527f69f7b12fff21773d272fbc2fbcbbaf65ae6efd2cdecf65f9407fed73bdde6ffdfb926f7e7637",
        "difficulty": "0x2779a24f1434e5",
        "number": "0xcf6ac0",
        "gasLimit": "0x1ca35d3",
        "gasUsed": "0x1c9ec32",
        "timestamp": "0x618cb514",
        "extraData": "0x657531",
        "mixHash": "0xb02a45eaea019e160f333b55fa98685d82db63f30b8510572b8f7bf85f9aca07",
        "nonce": "0xe47c9b3308220bc6",
        "baseFeePerGas": "0x183a551c7b",
        "hash": "0xc38f4490705bb7c2de60d4f0aaf0f2c3dd3ad0069d0ea0c68774250c16e219cb"
    },
    "bloom": "0x75b2597752a9dfdf77ecb7de9edd3b7f89e2f1b79e47daad1f87e7777f96d61f948fd74f4cfce68edd5f59c2f3a9c517efd794d77c7effb9ff7d9d2c5ffff8f3e2d496d8ffffbf29dfc272febaf5d7e7b2fff0cfed6f905325713543eec7d557fffeb71d3fe7ea5b7ffcfbf540fd7bf9dffbf5f49ed7b572fa3edfb3e93ceb359eb5f7ebf7e2ff3ed8fe0cf33bda06355ca09d8d6fb7ae7a7fb5f85fdfb745681f7dfc7b13efed732fadf6dbfd830ff3febdf279f4d6ed66abebfffe3a6e58dee15e331fdb6d5c73afb74fd777dddb1e6f34b61f1d2ee4be527f69f7b12fff21773d272fbc2fbcbbaf65ae6efd2cdecf65f9407fed73bdde6ffdfb926f7e7637",
    "receivedAt": "0001-01-01T00:00:00Z",
    "receivedFrom": null,
    "baseFee": 104057871483,
    "coinbase": "0xc93452a74e596e81e4f73ca1acff532089ad4c62",
    "difficulty": 11111262110102757,
    "extra": "ZXUx",
    "gasLimit": 30029267,
    "gasUsed": 30010418,
    "hash": "0xc38f4490705bb7c2de60d4f0aaf0f2c3dd3ad0069d0ea0c68774250c16e219cb",
    "mixDigest": "0xb02a45eaea019e160f333b55fa98685d82db63f30b8510572b8f7bf85f9aca07",
    "nonce": 16464204981241777094,
    "parentHash": "0x855637ad8b0c7ed477460118f5db1be30f6d15518a6e5c8d44d6db49aff31591",
    "receiptHash": "0x83fcf0376a111b51881f347a23f5a97fcbd837637f95ba5e0c7ce66d90ef4461",
    "root": "0xc70b69c817b8761ad9c564a45c8b9281e8527a2565c3360d68e7be6a722541ac",
    "size": 115840,
    "time": 1636611348,
    "uncles": [],
}
```

#### /block/:blockNumber/txs/:txNumOrHash
* `GET` : Get a transaction of a block [has either index or tx hash as a param.]

*Response:*

```js
HTTP 200 OK

{
    "type": "0x2",
    "nonce": "0x5232",
    "gasPrice": null,
    "maxPriorityFeePerGas": "0x3b9aca00",
    "maxFeePerGas": "0x1bd7ebf0f1",
    "gas": "0xf4240",
    "value": "0x0",
    "input": "0x000000520000000000000000000000000000000000000000000000000000000000cf5871000000000000000000000000000000000000000000000000320880c6ada9f111000000000000000000000000c18360217d8f7ab5e7c516566761ea12ce7f9d72000000000000000000000000b87b65dacc6171b3ca8c4a934601d0fcb6c61049000000000000000000000000a1181481beb2dc5de0daf2c85392d81c704bf75d00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001",
    "v": "0x0",
    "r": "0x3c5795728b1a870f17da8046249916d27afbb2d1429fc06f39dfb8c1beff6220",
    "s": "0x1727a2c49a42cf2a5ca2b4de81cf50859aff6096aae4938651b64f7d3cbda250",
    "to": "0x58418d6c83efab01ed78b0ac42e55af01ee77dba",
    "chainId": "0x1",
    "accessList": [],
    "hash": "0x856ed296eb7f87619393143cc28ae5705e22866709ff79cd93af71d8132037c0"
}
```

## Returned http error codes

| HTTP code | description |
| :-: | :-: |
| 400 | bad request |
| 404 | not found |
| 500 | Internal server error |


## TODOs
- Increase test coverage
- Isolate `ethclient` domain models in `repository/cloudflare` package (now only `types.Block` is isolated, but not `types.Transaction` and `types.Bloom`)
- Add data transfer object for transaction response
- Add separate cache for last 20 blocks with block hash checks on request
- Add separate cache for transactions only (now if you request the transaction the whole block is caching)

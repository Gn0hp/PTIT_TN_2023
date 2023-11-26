# Online Election

[![Go](https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667)](https://go.dev//)
[Gin-Gonic](https://gin-gonic.com) [Gorm](https://gorm.io/index.html) [Vue2](https://v2.vuejs.org/) [BSC](https://www.bnbchain.org/en/smartChain)
[MySQL](https://www.mysql.com/) [Redis](https://redis.io/)

## Introduction

## Installation

### Requirements
    - node >= ^20.9.0
    - npm >= ^10.1.0
    - go >= ^1.20.0
Then make sure MySQL, redis-server and RabbitMQ are installed and running on your machine. 
### Dependencies
- Client
```bash
$ cd ./client
$ npm install
```
If there is any error as version conflict, try to use yarn with stable version.
```bash
$ npm install -g yarn
$ yarn set version stable
$ yarn install
$ cd ..
```
- Server:
```bash
$ go mod download
```
### Running the app
+ Client 
```bash
$ cd ./client
# build
$ npm run build

# watch mode
$ npm run dev
```
+ Server
```bash
# Migrate database
$ go build ./cli

# Run server
$ go build -o ./bin/server ./cmd/servers

# Run microservice
$ go build -o ./bin/server ./cmd/servers
```
### Test
```bash
# unit tests
$ npm run test

# e2e tests
$ npm run test:e2e

# test coverage
$ npm run test:cov
```
## Support

For more information, please [contact here](https://github.com/Gn0hp) or mail to [gn0hp289@gmail.com]() .



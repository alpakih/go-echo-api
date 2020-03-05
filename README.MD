[![Coverage Status](https://coveralls.io/repos/github/huzairuje/go-echo-api/badge.svg?branch=master)](https://coveralls.io/github/huzairuje/go-echo-api?branch=master)
[![Build Status](https://travis-ci.org/huzairuje/go-echo-api.svg?branch=master)](https://travis-ci.org/huzairuje/go-echo-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# REST API Go (Echo HTTP Framework)

## Introduction
here's making REST API with echo Framework (Go Programming Language), 
using module based frame thinking

## Installation
#### 1. Requirements

##### a. Golang Language SDK minimal 1.13 https://golang.org/dl/
##### b. Dep Golang Package Management https://golang.github.io/dep/docs/installation.html
##### c. Database (PostgreSQL)

#### 2.Clone The Projects
Go to your src/ folder inside go project directory
```$xslt
    cd ~YOUR/GO/FOLDER/DIRECTORY/src/
``` 
then clone the project
```$xslt
    git clone https://github.com/alpakih/go-echo-api.git
```
#### 3. install all the dependencies
a. Go to the project folder
```$xslt
    cd ~YOUR/GO/FOLDER/DIRECTORY/src/go-echo-api/
```
b. then run `dep ensure` command.
## Database
Firstly you need a Database, and add new database name with
`go-echo-api`

## .Env (environment variable)
Make sure `dep ensure` is successful, and then make .env file with command
```$xslt
    cp .env.example .env
```

Update the content of `.env` value, like the database name
```$xslt
    DB_DRIVER=postgres
    DB_NAME=echo-api
    DB_PORT=5432
    DB_HOST=localhost
    DB_USERNAME=postgres
    DB_PASSWORD=postgres
    DB_SSL=disable
```
## Run
run the project with
```$xslt
    go run main.go
```

## Build
#### 1. Linux
```$xslt
    #build for linux (arch amd64)
    env GOOS=linux GOARCH=amd64 go build -o go-echo-api-linux-amd64  && mv ./go-echo-api-linux-amd64 ./build/
```

#### 2. MacOs (Darwin / BSD)
```$xslt
    #build for MacOs (arch amd64)
    env GOOS=darwin GOARCH=amd64 go build -o go-echo-api-darwin-amd64  && mv ./go-echo-api-darwin-amd64 ./build/
```

#### 3. Windows (DOS)
```$xslt
    #build for Windows (arch amd64)
    env GOOS=windows GOARCH=amd64 go build -o go-echo-api-windows-amd64.exe  && mv ./go-echo-api-windows-amd64.exe ./build/
```

## License
This package is open-sourced software licensed under the MIT license.
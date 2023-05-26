# Getting started with Testcontainers for Go

This is sample code for [Getting started with Testcontainers for Go](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go) guide.

## 1. Setup Environment
Make sure you have Go 1.18+ and a [compatible Docker environment](https://golang.testcontainers.org/system_requirements/docker/) installed.

For example:

```shell
$ go version
go version go1.19.3 darwin/arm64

$ docker version
...
Server: Docker Desktop 4.19.0 (106363)
 Engine:
  Version:          23.0.5
  API version:      1.42 (minimum version 1.12)
  Go version:       go1.19.8
...
```

## 2. Setup Project

* Clone the repository

```shell
git clone https://github.com/testcontainers/tc-guide-getting-started-with-testcontainers-for-go.git
cd tc-guide-getting-started-with-testcontainers-for-go
```

* Open the **tc-guide-getting-started-with-testcontainers-for-go** project in your favorite IDE.

## 3. Run Tests

Run the command to run the tests.

```shell
$ go test -v ./...
```

The tests should pass.

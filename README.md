# Logur adapter for TEMPLATE

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/logur/adapter-apex/CI?style=flat-square)](https://github.com/logur/adapter-apex/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/logur/adapter-apex?style=flat-square)](https://codecov.io/gh/logur/adapter-apex)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/adapter/apex?style=flat-square)](https://goreportcard.com/report/logur.dev/adapter/apex)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/logur.dev/adapter/apex)


## Installation

```bash
go get logur.dev/adapter/apex
```


## Usage

```go
package main

import (
	apexadapter "logur.dev/adapter/apex"
)

func main() {
	logger := apexadapter.New(/*logger*/)
}
```


## Development

When all coding and testing is done, please run the test suite:

```bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

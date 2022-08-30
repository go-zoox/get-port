# GetPort - Get a free and open tcp port that is ready to use

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/get-port)](https://pkg.go.dev/github.com/go-zoox/get-port)
[![Build Status](https://github.com/go-zoox/get-port/actions/workflows/lint.yml/badge.svg?branch=master)](https://github.com/go-zoox/get-port/actions/workflows/lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/get-port)](https://goreportcard.com/report/github.com/go-zoox/get-port)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/get-port/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/get-port?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/get-port.svg)](https://github.com/go-zoox/get-port/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/get-port.svg?label=Release)](https://github.com/go-zoox/get-port/tags)


## Installation
To install the package, run:
```bash
go get -u github.com/go-zoox/get-port
```

## Quick Start

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/go-zoox/get-port"
)

func main() {
	fmt.Println("Get a free port:", getport.MustGetPort())
}

```

## Inspiration
* [phayes/freeport](https://github.com/phayes/freeport) - Get a free and open tcp port that is ready to use.

## License
GoZoox is released under the [MIT License](./LICENSE).

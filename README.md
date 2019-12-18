# go-opengraph
OpenGraph Protocol fetcher for Golang

[![Build Status](https://travis-ci.com/masa0221/go-opengraph.svg?branch=master)](https://travis-ci.com/masa0221/go-opengraph)

## Installation

```
$ go get github.com/masa0221/go-opengraph
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/masa0221/go-opengraph"
)

func main() {
	og, err := opengraph.Fetch("http://ogp.me")
	if (err != nil) {
		fmt.Println(err)
	}

	fmt.Println(og)
}
```

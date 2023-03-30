# Telegram Bot API

[Telegram Bot API](https://core.telegram.org/bots/api) [Golang](https://golang.org) implementation

[![Build Status](https://travis-ci.org/nasermirzaei89/telegram.svg?branch=master)](https://travis-ci.org/nasermirzaei89/telegram)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/telegram)](https://goreportcard.com/report/github.com/nasermirzaei89/telegram)
[![GolangCI](https://golangci.com/badges/github.com/nasermirzaei89/telegram.svg)](https://golangci.com/r/github.com/nasermirzaei89/telegram)
[![Codecov](https://codecov.io/gh/nasermirzaei89/telegram/branch/master/graph/badge.svg)](https://codecov.io/gh/nasermirzaei89/telegram)
[![GoDoc](https://godoc.org/github.com/nasermirzaei89/telegram?status.svg)](https://godoc.org/github.com/nasermirzaei89/telegram)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://raw.githubusercontent.com/nasermirzaei89/telegram/master/LICENSE)

## Status

Bot API 6.0 (https://core.telegram.org/bots/api-changelog#april-16-2022)

## Install

```sh
go get github.com/nasermirzaei89/telegram
```

## Example

```go
package main

import (
	"context"
	"github.com/nasermirzaei89/telegram"
	"log"
	"os"
)

func main() {
	bot := telegram.New(os.Getenv("TOKEN"))

	res, err := bot.GetUpdates(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	if res.IsOK() {
		log.Printf("%+v", res.GetUpdates())
	} else {
		log.Printf("%d: %s", res.GetErrorCode(), res.GetDescription())
	}
}
```

## License

MIT License

## Contributing

You can submit a [new issue](https://github.com/nasermirzaei89/telegram/issues/new) in GitHub [issues](https://github.com/nasermirzaei89/telegram/issues).
Or you can [create a fork](https://help.github.com/articles/fork-a-repo), hack on your fork and when you're done create a [pull request](https://help.github.com/articles/fork-a-repo#pull-requests), so that the code contribution can get merged into the main package.

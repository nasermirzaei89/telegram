# Telegram Bot API

[Telegram Bot API](https://core.telegram.org/bots/api) [Golang](https://golang.org) implementation

![Build Status](https://github.com/nasermirzaei89/telegram/actions/workflows/build.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/telegram)](https://goreportcard.com/report/github.com/nasermirzaei89/telegram)
[![Codecov](https://codecov.io/gh/nasermirzaei89/telegram/branch/master/graph/badge.svg)](https://codecov.io/gh/nasermirzaei89/telegram)
[![Go Reference](https://pkg.go.dev/badge/github.com/nasermirzaei89/telegram.svg)](https://pkg.go.dev/github.com/nasermirzaei89/telegram)
[![License](https://img.shields.io/github/license/nasermirzaei89/telegram)](https://raw.githubusercontent.com/nasermirzaei89/telegram/master/LICENSE)

## Status

Bot API 5.5 (https://core.telegram.org/bots/api#december-7-2021)

## Install

```sh
go get github.com/nasermirzaei89/telegram
```

## Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/nasermirzaei89/telegram"
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
Or you can [create a fork](https://help.github.com/articles/fork-a-repo),
hack on your fork and when you're done create a [pull request](https://help.github.com/articles/fork-a-repo#pull-requests),
so that the code contribution can get merged into the main package.

# Telegram Bot API

[telegram bot api](https://core.telegram.org/bots/api) [golang](https://golang.org) implementation

[![Build Status](https://travis-ci.org/nasermirzaei89/telegram.svg?branch=master)](https://travis-ci.org/nasermirzaei89/telegram)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/telegram)](https://goreportcard.com/report/github.com/nasermirzaei89/telegram)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://raw.githubusercontent.com/nasermirzaei89/telegram/master/LICENSE)

## Install

```sh
go get github.com/nasermirzaei89/telegram
```

## Example

```go
package main

import (
	"github.com/nasermirzaei89/telegram"
	"log"
	"os"
)

func main() {
	bot := telegram.New(os.Getenv("TOKEN"))
	res, err := bot.GetUpdates()
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("%v", res)
}
```

## Status

Bot API 4.3 (https://core.telegram.org/bots/api#may-31-2019)


## License

MIT License

## Contributing

You can submit a [new issue](https://github.com/nasermirzaei89/telegram/issues/new) in github [issues](https://github.com/nasermirzaei89/telegram/issues).
Or you can [create a fork](https://help.github.com/articles/fork-a-repo), hack on your fork and when you're done create a [pull request](https://help.github.com/articles/fork-a-repo#pull-requests), so that the code contribution can get merged into the main package.

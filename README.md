# Telegram Bot API
[telegram bot api](https://core.telegram.org/bots/api) [golang](https://golang.org) implementation

[![Build Status](https://travis-ci.org/lujem/telegram.svg?branch=master)](https://travis-ci.org/lujem/telegram)
[![Go Report Card](https://goreportcard.com/badge/github.com/lujem/telegram)](https://goreportcard.com/report/github.com/lujem/telegram)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://raw.githubusercontent.com/lujem/telegram/master/LICENSE)

## Install
```sh
go get github.com/lujem/telegram
```

## Example
no example yet!

## Status
- [x] Making requests
- [ ] Getting updates
    - [X] Update
    - [X] getUpdates
    - [ ] setWebhook
    - [ ] deleteWebhook
    - [ ] getWebhookInfo
    - [ ] WebhookInfo
- [X] Available types
    - [X] User
    - [X] Chat
    - [X] Message
    - [X] MessageEntity
    - [X] PhotoSize
    - [X] Audio
    - [X] Document
    - [X] Sticker
    - [X] Video
    - [X] Voice
    - [X] Contact
    - [X] Location
    - [X] Venue
    - [X] UserProfilePhotos
    - [X] File
    - [X] ReplyKeyboardMarkup
    - [X] KeyboardButton
    - [X] ReplyKeyboardRemove
    - [X] InlineKeyboardMarkup
    - [X] InlineKeyboardButton
    - [X] CallbackQuery
    - [X] ForceReply
    - [X] ChatMember
    - [X] ResponseParameters
    - [X] InputFile
- [X] Available methods
    - [X] getMe
    - [X] sendMessage
    - [X] forwardMessage
    - [X] sendPhoto
    - [X] sendAudio
    - [X] sendDocument
    - [X] sendSticker
    - [X] sendVideo
    - [X] sendVoice
    - [X] sendLocation
    - [X] sendVenue
    - [X] sendContact
    - [X] sendChatAction
    - [X] getUserProfilePhotos
    - [X] getFile
    - [X] kickChatMember
    - [X] leaveChat
    - [X] unbanChatMember
    - [X] getChat
    - [X] getChatAdministrators
    - [X] getChatMembersCount
    - [X] getChatMember
    - [X] answerCallbackQuery
- [ ] Updating messages
    - [ ] editMessageText
    - [ ] editMessageCaption
    - [ ] editMessageReplyMarkup
- [ ] Inline mode
    - [X] InlineQuery
    - [X] answerInlineQuery
    - [X] InlineQueryResult
    - [X] InlineQueryResultArticle
    - [X] InlineQueryResultPhoto
    - [X] InlineQueryResultGif
    - [X] InlineQueryResultMpeg4Gif
    - [X] InlineQueryResultVideo
    - [X] InlineQueryResultAudio
    - [X] InlineQueryResultVoice
    - [X] InlineQueryResultDocument
    - [X] InlineQueryResultLocation
    - [X] InlineQueryResultVenue
    - [X] InlineQueryResultContact
    - [X] InlineQueryResultGame
    - [X] InlineQueryResultCachedPhoto
    - [X] InlineQueryResultCachedGif
    - [X] InlineQueryResultCachedMpeg4Gif
    - [X] InlineQueryResultCachedSticker
    - [X] InlineQueryResultCachedDocument
    - [X] InlineQueryResultCachedVideo
    - [X] InlineQueryResultCachedVoice
    - [X] InlineQueryResultCachedAudio
    - [X] InputMessageContent
    - [X] InputTextMessageContent
    - [X] InputLocationMessageContent
    - [X] InputVenueMessageContent
    - [X] InputContactMessageContent
    - [X] ChosenInlineResult
- [ ] Games
    - [ ] sendGame
    - [X] Game
    - [X] Animation
    - [ ] CallbackGame
    - [ ] setGameScore
    - [ ] getGameHighScores
    - [X] GameHighScore


## License
MIT License

## Contributing
You can submit a [new issue](https://github.com/lujem/telegram/issues/new) in github [issues](https://github.com/lujem/telegram/issues).
Or you can [create a fork](https://help.github.com/articles/fork-a-repo), hack on your fork and when you're done create a [pull request](https://help.github.com/articles/fork-a-repo#pull-requests), so that the code contribution can get merged into the main package.

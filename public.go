package telegram

// GetMe is a simple method for testing your bot's auth token. Requires no parameters. Returns basic information about the bot in form of a User object.
func (obj *API) GetMe() (*User, error) {
	return obj.getMe()
}

// GetUpdates is a method to receive incoming updates using long polling
func (obj *API) GetUpdates() ([]Update, error) {
	updates, err := obj.getUpdates(obj.offset, 100, 0, []string{})
	if err != nil {
		return nil, err
	}

	if len(updates) > 0 {
		obj.offset = updates[0].UpdateID + 1
	}

	return updates, nil
}

// SendMessage is a method to send text messages
func (obj *API) SendMessage(chatID interface{}, text string, parseMode *string, disableWebPagePreview, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendMessage(chatID, text, parseMode, disableWebPagePreview, disableNotification, replyToMessageID, replyMarkup)
}

// ForwardMessage is a method to forward messages of any kind
func (obj *API) ForwardMessage(chatID, fromChatID interface{}, disableNotification *bool, messageID int64) (*Message, error) {
	return obj.forwardMessage(chatID, fromChatID, disableNotification, messageID)
}

// SendPhoto is a method to send photos
func (obj *API) SendPhoto(chatID interface{}, photo interface{}, caption *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendPhoto(chatID, photo, caption, disableNotification, replyToMessageID, replyMarkup)
}

// SendAudio is a method to send audio files, if you want Telegram clients to display them in the music player
func (obj *API) SendAudio(chatID interface{}, audio interface{}, caption *string, duration *int64, performer *string, title *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendAudio(chatID, audio, caption, duration, performer, title, disableNotification, replyToMessageID, replyMarkup)
}

// SendDocument is a method to send general files
func (obj *API) SendDocument(chatID interface{}, document interface{}, caption *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendDocument(chatID, document, caption, disableNotification, replyToMessageID, replyMarkup)
}

// SendSticker is a method to send .webp stickers
func (obj *API) SendSticker(chatID interface{}, sticker interface{}, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendSticker(chatID, sticker, disableNotification, replyToMessageID, replyMarkup)
}

// SendVideo is a method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document)
func (obj *API) SendVideo(chatID interface{}, video interface{}, duration, width, height *int64, caption *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendVideo(chatID, video, duration, width, height, caption, disableNotification, replyToMessageID, replyMarkup)
}

// SendVoice is a method to send audio files, if you want Telegram clients to display the file as a playable voice message
func (obj *API) SendVoice(chatID interface{}, voice interface{}, caption *string, duration *int64, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendVoice(chatID, voice, caption, duration, disableNotification, replyToMessageID, replyMarkup)
}

// SendLocation is a method to send point on the map
func (obj *API) SendLocation(chatID interface{}, latitude, longitude float64, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendLocation(chatID, latitude, longitude, disableNotification, replyToMessageID, replyMarkup)
}

// SendVenue is a method to send information about a venue
func (obj *API) SendVenue(chatID interface{}, latitude, longitude float64, title, address string, foursquareID *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendVenue(chatID, latitude, longitude, title, address, foursquareID, disableNotification, replyToMessageID, replyMarkup)
}

// SendContact is a method to send phone contacts
func (obj *API) SendContact(chatID interface{}, phoneNumber string, firstName string, lastName *string, disableNotification *bool, replyToMessageID *int64, replyMarkup interface{}) (*Message, error) {
	return obj.sendContact(chatID, phoneNumber, firstName, lastName, disableNotification, replyToMessageID, replyMarkup)
}

// SendChatAction is a method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status)
func (obj *API) SendChatAction(chatID interface{}, action string) (*bool, error) {
	return obj.sendChatAction(chatID, action)
}

// GetUserProfilePhotos is a method to get a list of profile pictures for a user
func (obj *API) GetUserProfilePhotos(userID int64, offset, limit *int64) (*UserProfilePhotos, error) {
	return obj.getUserProfilePhotos(userID, offset, limit)
}

// GetFile is a method to get basic info about a file and prepare it for downloading
func (obj *API) GetFile(fileID string) (*File, error) {
	return obj.getFile(fileID)
}

// KickChatMember is a method to kick a user from a group or a supergroup. In the case of supergroups, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the group for this to work.
func (obj *API) KickChatMember(chatID interface{}, userID int64) (*bool, error) {
	return obj.kickChatMember(chatID, userID)
}

// LeaveChat is a method for your bot to leave a group, supergroup or channel
func (obj *API) LeaveChat(chatID interface{}) (*bool, error) {
	return obj.leaveChat(chatID)
}

// UnbanChatMember is a method to unban a previously kicked user in a supergroup. The user will not return to the group automatically, but will be able to join via link, etc. The bot must be an administrator in the group for this to work.
func (obj *API) UnbanChatMember(chatID interface{}, userID int64) (*bool, error) {
	return obj.unbanChatMember(chatID, userID)
}

// GetChat is a method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.)
func (obj *API) GetChat(chatID interface{}) (*Chat, error) {
	return obj.getChat(chatID)
}

// GetChatAdministrators is a method to get a list of administrators in a chat
func (obj *API) GetChatAdministrators(chatID interface{}) (*[]ChatMember, error) {
	return obj.getChatAdministrators(chatID)
}

// GetChatMembersCount is a method to get the number of members in a chat
func (obj *API) GetChatMembersCount(chatID interface{}) (*int64, error) {
	return obj.getChatMembersCount(chatID)
}

// GetChatMember is a method to get information about a member of a chat
func (obj *API) GetChatMember(chatID interface{}, userID int64) (*ChatMember, error) {
	return obj.getChatMember(chatID, userID)
}

// AnswerCallbackQuery is a method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
func (obj *API) AnswerCallbackQuery(callbackQueryID string, text *string, showAlert *bool, url *string, cacheTime *int64) (*bool, error) {
	return obj.answerCallbackQuery(callbackQueryID, text, showAlert, url, cacheTime)
}

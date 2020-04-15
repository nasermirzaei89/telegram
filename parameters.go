package telegram

// Option type
type Option func(*request)

// SetOffset option function
func SetOffset(v int) Option {
	return func(r *request) {
		r.params["offset"] = v
	}
}

// SetLimit option function
func SetLimit(v int) Option {
	return func(r *request) {
		r.params["limit"] = v
	}
}

// SetTimeout option function
func SetTimeout(v int) Option {
	return func(r *request) {
		r.params["timeout"] = v
	}
}

// SetAllowedUpdates option function
func SetAllowedUpdates(v ...string) Option {
	return func(r *request) {
		r.params["allowed_updates"] = v
	}
}

// SetURL option function
func SetURL(v string) Option {
	return func(r *request) {
		r.params["url"] = v
	}
}

// SetCertificate option function
func SetCertificate(v InputFile) Option {
	return func(r *request) {
		r.params["certificate"] = v
	}
}

// SetMaxConnections option function
func SetMaxConnections(v int) Option {
	return func(r *request) {
		r.params["max_connections"] = v
	}
}

// SetChatID option function
func SetChatID(v int) Option {
	return func(r *request) {
		r.params["chat_id"] = v
	}
}

// SetChatUsername option function
func SetChatUsername(v string) Option {
	return func(r *request) {
		r.params["chat_id"] = v
	}
}

// SetText option function
func SetText(v string) Option {
	return func(r *request) {
		r.params["text"] = v
	}
}

// SetParseMode option function
func SetParseMode(v string) Option {
	return func(r *request) {
		r.params["parse_mode"] = v
	}
}

// DisableWebPagePreview option function
func DisableWebPagePreview() Option {
	return func(r *request) {
		r.params["disable_web_page_preview"] = true
	}
}

// DisableNotification option function
func DisableNotification() Option {
	return func(r *request) {
		r.params["disable_notification"] = true
	}
}

// SetReplyToMessageID option function
func SetReplyToMessageID(v int) Option {
	return func(r *request) {
		r.params["reply_to_message_id"] = v
	}
}

// SetReplyMarkup option function
func SetReplyMarkup(v interface{}) Option {
	return func(r *request) {
		r.params["reply_markup"] = v
	}
}

// SetFromChatID option function
func SetFromChatID(v int) Option {
	return func(r *request) {
		r.params["from_chat_id"] = v
	}
}

// SetFromChatUsername option function
func SetFromChatUsername(v string) Option {
	return func(r *request) {
		r.params["from_chat_id"] = v
	}
}

// SetMessageID option function
func SetMessageID(v int) Option {
	return func(r *request) {
		r.params["message_id"] = v
	}
}

// SetPhotoFromFileID option function
func SetPhotoFromFileID(v string) Option {
	return func(r *request) {
		r.params["photo"] = v
	}
}

// SetPhotoFromURL option function
func SetPhotoFromURL(v string) Option {
	return func(r *request) {
		r.params["photo"] = v
	}
}

// SetPhoto option function
func SetPhoto(v InputFile) Option {
	return func(r *request) {
		r.params["photo"] = v
	}
}

// SetCaption option function
func SetCaption(v string) Option {
	return func(r *request) {
		r.params["caption"] = v
	}
}

// SetAudioFromFileID option function
func SetAudioFromFileID(v string) Option {
	return func(r *request) {
		r.params["audio"] = v
	}
}

// SetAudioFromURL option function
func SetAudioFromURL(v string) Option {
	return func(r *request) {
		r.params["audio"] = v
	}
}

// SetAudio option function
func SetAudio(v InputFile) Option {
	return func(r *request) {
		r.params["audio"] = v
	}
}

// SetDuration option function
func SetDuration(v int) Option {
	return func(r *request) {
		r.params["duration"] = v
	}
}

// SetPerformer option function
func SetPerformer(v string) Option {
	return func(r *request) {
		r.params["performer"] = v
	}
}

// SetTitle option function
func SetTitle(v string) Option {
	return func(r *request) {
		r.params["title"] = v
	}
}

// TODO: add Thumb string option

// SetThumb option function
func SetThumb(v InputFile) Option {
	return func(r *request) {
		r.params["thumb"] = v
	}
}

// SetDocumentFromFileID option function
func SetDocumentFromFileID(v string) Option {
	return func(r *request) {
		r.params["document"] = v
	}
}

// SetDocumentFromURL option function
func SetDocumentFromURL(v string) Option {
	return func(r *request) {
		r.params["document"] = v
	}
}

// SetDocument option function
func SetDocument(v InputFile) Option {
	return func(r *request) {
		r.params["document"] = v
	}
}

// SetVideoFromFileID option function
func SetVideoFromFileID(v string) Option {
	return func(r *request) {
		r.params["video"] = v
	}
}

// SetVideoFromURL option function
func SetVideoFromURL(v string) Option {
	return func(r *request) {
		r.params["video"] = v
	}
}

// SetVideo option function
func SetVideo(v InputFile) Option {
	return func(r *request) {
		r.params["video"] = v
	}
}

// SetWidth option function
func SetWidth(v string) Option {
	return func(r *request) {
		r.params["width"] = v
	}
}

// SetHeight option function
func SetHeight(v string) Option {
	return func(r *request) {
		r.params["height"] = v
	}
}

// SupportsStreaming option function
func SupportsStreaming() Option {
	return func(r *request) {
		r.params["supports_streaming"] = true
	}
}

// SetAnimationFromFileID option function
func SetAnimationFromFileID(v string) Option {
	return func(r *request) {
		r.params["animation"] = v
	}
}

// SetAnimationFromURL option function
func SetAnimationFromURL(v string) Option {
	return func(r *request) {
		r.params["animation"] = v
	}
}

// SetAnimation option function
func SetAnimation(v InputFile) Option {
	return func(r *request) {
		r.params["animation"] = v
	}
}

// SetVoiceFromFileID option function
func SetVoiceFromFileID(v string) Option {
	return func(r *request) {
		r.params["voice"] = v
	}
}

// SetVoiceFromURL option function
func SetVoiceFromURL(v string) Option {
	return func(r *request) {
		r.params["voice"] = v
	}
}

// SetVoice option function
func SetVoice(v InputFile) Option {
	return func(r *request) {
		r.params["voice"] = v
	}
}

// SetVideoNoteFromFileID option function
func SetVideoNoteFromFileID(v string) Option {
	return func(r *request) {
		r.params["video_note"] = v
	}
}

// SetVideoNoteFromURL option function
func SetVideoNoteFromURL(v string) Option {
	return func(r *request) {
		r.params["video_note"] = v
	}
}

// SetVideoNote option function
func SetVideoNote(v InputFile) Option {
	return func(r *request) {
		r.params["video_note"] = v
	}
}

// SetMedia option function
func SetMedia(v ...InputMedia) Option {
	return func(r *request) {
		r.params["media"] = v
	}
}

// SetLatitude option function
func SetLatitude(v float32) Option {
	return func(r *request) {
		r.params["latitude"] = v
	}
}

// SetLongitude option function
func SetLongitude(v float32) Option {
	return func(r *request) {
		r.params["longitude"] = v
	}
}

// SetLivePeriod option function
func SetLivePeriod(v int) Option {
	return func(r *request) {
		r.params["live_period"] = v
	}
}

// SetInlineMessageID option function
func SetInlineMessageID(v string) Option {
	return func(r *request) {
		r.params["inline_message_id"] = v
	}
}

// SetInlineQueryID option function
func SetInlineQueryID(v string) Option {
	return func(r *request) {
		r.params["inline_query_id"] = v
	}
}

// SetResults option function
func SetResults(v []InlineQueryResult) Option {
	return func(r *request) {
		r.params["results"] = v
	}
}

// SetAddress option function
func SetAddress(v string) Option {
	return func(r *request) {
		r.params["address"] = v
	}
}

// SetFoursquareID option function
func SetFoursquareID(v string) Option {
	return func(r *request) {
		r.params["foursquare_id"] = v
	}
}

// SetFoursquareType option function
func SetFoursquareType(v string) Option {
	return func(r *request) {
		r.params["foursquare_type"] = v
	}
}

// SetPhoneNumber option function
func SetPhoneNumber(v string) Option {
	return func(r *request) {
		r.params["phone_number"] = v
	}
}

// SetFirstName option function
func SetFirstName(v string) Option {
	return func(r *request) {
		r.params["first_name"] = v
	}
}

// SetLastName option function
func SetLastName(v string) Option {
	return func(r *request) {
		r.params["last_name"] = v
	}
}

// SetVCard option function
func SetVCard(v string) Option {
	return func(r *request) {
		r.params["vcard"] = v
	}
}

// SetQuestion option function
func SetQuestion(v string) Option {
	return func(r *request) {
		r.params["question"] = v
	}
}

// SetOptions option function
func SetOptions(v ...string) Option {
	return func(r *request) {
		r.params["options"] = v
	}
}

// SetCommands option function
func SetCommands(v ...BotCommand) Option {
	return func(r *request) {
		r.params["commands"] = v
	}
}

// SetAction option function
func SetAction(v string) Option {
	return func(r *request) {
		r.params["action"] = v
	}
}

// SetFileID option function
func SetFileID(v string) Option {
	return func(r *request) {
		r.params["file_id"] = v
	}
}

// SetUserID option function
func SetUserID(v int) Option {
	return func(r *request) {
		r.params["user_id"] = v
	}
}

// SetUntilDate option function
func SetUntilDate(v int) Option {
	return func(r *request) {
		r.params["until_date"] = v
	}
}

// CanSendMessages option function
func CanSendMessages() Option {
	return func(r *request) {
		r.params["can_send_messages"] = true
	}
}

// CanSendMediaMessages option function
func CanSendMediaMessages() Option {
	return func(r *request) {
		r.params["can_send_media_messages"] = true
	}
}

// CanSendOtherMessages option function
func CanSendOtherMessages() Option {
	return func(r *request) {
		r.params["can_send_other_messages"] = true
	}
}

// CanAddWebPagePreviews option function
func CanAddWebPagePreviews() Option {
	return func(r *request) {
		r.params["can_add_web_page_previews"] = true
	}
}

// CanChangeInfo option function
func CanChangeInfo() Option {
	return func(r *request) {
		r.params["can_change_info"] = true
	}
}

// CanPostMessages option function
func CanPostMessages() Option {
	return func(r *request) {
		r.params["can_post_messages"] = true
	}
}

// CanEditMessages option function
func CanEditMessages() Option {
	return func(r *request) {
		r.params["can_edit_messages"] = true
	}
}

// CanDeleteMessages option function
func CanDeleteMessages() Option {
	return func(r *request) {
		r.params["can_delete_messages"] = true
	}
}

// CanInviteUsers option function
func CanInviteUsers() Option {
	return func(r *request) {
		r.params["can_invite_users"] = true
	}
}

// CanRestrictMemebers option function
func CanRestrictMemebers() Option {
	return func(r *request) {
		r.params["can_restrict_members"] = true
	}
}

// CanPinMessages option function
func CanPinMessages() Option {
	return func(r *request) {
		r.params["can_pin_messages"] = true
	}
}

// CanPromoteMemebers option function
func CanPromoteMemebers() Option {
	return func(r *request) {
		r.params["can_promote_members"] = true
	}
}

// SetDescription option function
func SetDescription(v string) Option {
	return func(r *request) {
		r.params["description"] = v
	}
}

// SetStickerSetName option function
func SetStickerSetName(v string) Option {
	return func(r *request) {
		r.params["sticker_set_name"] = v
	}
}

// SetCallbackQueryID option function
func SetCallbackQueryID(v string) Option {
	return func(r *request) {
		r.params["callback_query_id"] = v
	}
}

// ShowAlert option function
func ShowAlert() Option {
	return func(r *request) {
		r.params["show_alert"] = true
	}
}

// SetCacheTime option function
func SetCacheTime(v int) Option {
	return func(r *request) {
		r.params["cache_time"] = v
	}
}

// SetNextOffset option function
func SetNextOffset(v string) Option {
	return func(r *request) {
		r.params["next_offset"] = v
	}
}

// SetSwitchPMText option function
func SetSwitchPMText(v string) Option {
	return func(r *request) {
		r.params["switch_pm_text"] = v
	}
}

// SetSwitchPMParameter option function
func SetSwitchPMParameter(v string) Option {
	return func(r *request) {
		r.params["switch_pm_parameter"] = v
	}
}

// IsPersonal option function
func IsPersonal() Option {
	return func(r *request) {
		r.params["is_personal"] = true
	}
}

// SetPermissions option function
func SetPermissions(v ChatPermissions) Option {
	return func(r *request) {
		r.params["permissions"] = v
	}
}

// SetCustomTitle option function
func SetCustomTitle(v string) Option {
	return func(r *request) {
		r.params["custom_title"] = v
	}
}

// IsAnonymous option function
func IsAnonymous() Option {
	return func(r *request) {
		r.params["is_anonymous"] = true
	}
}

// SetType option function
func SetType(v string) Option {
	return func(r *request) {
		r.params["type"] = v
	}
}

// AllowsMultipleAnswers option function
func AllowsMultipleAnswers() Option {
	return func(r *request) {
		r.params["allows_multiple_answers"] = true
	}
}

// SetCorrectOptionID option function
func SetCorrectOptionID(v int) Option {
	return func(r *request) {
		r.params["correct_option_id"] = v
	}
}

// IsClosed option function
func IsClosed() Option {
	return func(r *request) {
		r.params["is_closed"] = true
	}
}

// SetPNGStickerFromFileID option function
func SetPNGStickerFromFileID(v string) Option {
	return func(r *request) {
		r.params["png_sticker"] = v
	}
}

// SetPNGStickerFromURL option function
func SetPNGStickerFromURL(v string) Option {
	return func(r *request) {
		r.params["png_sticker"] = v
	}
}

// SetPNGSticker option function
func SetPNGSticker(v InputFile) Option {
	return func(r *request) {
		r.params["png_sticker"] = v
	}
}

// SetTGSSticker option function
func SetTGSSticker(v InputFile) Option {
	return func(r *request) {
		r.params["tgs_sticker"] = v
	}
}

// SetEmojis option function
func SetEmojis(v string) Option {
	return func(r *request) {
		r.params["emojis"] = v
	}
}

// ContainsMasks option function
func ContainsMasks() Option {
	return func(r *request) {
		r.params["contains_masks"] = true
	}
}

// SetMaskPosition option function
func SetMaskPosition(v MaskPosition) Option {
	return func(r *request) {
		r.params["mask_position"] = v
	}
}

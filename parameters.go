package telegram

// MethodOption type
type MethodOption func(*request)

// SetOffset option function
func SetOffset(v int) MethodOption {
	return func(r *request) {
		r.params["offset"] = v
	}
}

// SetLimit option function
func SetLimit(v int) MethodOption {
	return func(r *request) {
		r.params["limit"] = v
	}
}

// SetTimeout option function
func SetTimeout(v int) MethodOption {
	return func(r *request) {
		r.params["timeout"] = v
	}
}

// SetAllowedUpdates option function
func SetAllowedUpdates(v ...string) MethodOption {
	return func(r *request) {
		r.params["allowed_updates"] = v
	}
}

// SetURL option function
func SetURL(v string) MethodOption {
	return func(r *request) {
		r.params["url"] = v
	}
}

// SetCertificate option function
func SetCertificate(v InputFile) MethodOption {
	return func(r *request) {
		r.params["certificate"] = v
	}
}

// SetIPAddress option function
func SetIPAddress(v string) MethodOption {
	return func(r *request) {
		r.params["ip_address"] = v
	}
}

// DropPendingUpdates option function
func DropPendingUpdates() MethodOption {
	return func(r *request) {
		r.params["drop_pending_updates"] = true
	}
}

// OnlyIfBanned option function
func OnlyIfBanned() MethodOption {
	return func(r *request) {
		r.params["only_if_banned"] = true
	}
}

// DisableContentTypeDetection option function
func DisableContentTypeDetection() MethodOption {
	return func(r *request) {
		r.params["disable_content_type_detection"] = true
	}
}

// SetMaxConnections option function
func SetMaxConnections(v int) MethodOption {
	return func(r *request) {
		r.params["max_connections"] = v
	}
}

// SetChatID option function
func SetChatID(v int) MethodOption {
	return func(r *request) {
		r.params["chat_id"] = v
	}
}

// SetChatUsername option function
func SetChatUsername(v string) MethodOption {
	return func(r *request) {
		r.params["chat_id"] = v
	}
}

// SetText option function
func SetText(v string) MethodOption {
	return func(r *request) {
		r.params["text"] = v
	}
}

// SetParseMode option function
func SetParseMode(v string) MethodOption {
	return func(r *request) {
		r.params["parse_mode"] = v
	}
}

// DisableWebPagePreview option function
func DisableWebPagePreview() MethodOption {
	return func(r *request) {
		r.params["disable_web_page_preview"] = true
	}
}

// DisableNotification option function
func DisableNotification() MethodOption {
	return func(r *request) {
		r.params["disable_notification"] = true
	}
}

// SetReplyToMessageID option function
func SetReplyToMessageID(v int) MethodOption {
	return func(r *request) {
		r.params["reply_to_message_id"] = v
	}
}

// SetReplyMarkup option function
func SetReplyMarkup(v interface{}) MethodOption {
	return func(r *request) {
		r.params["reply_markup"] = v
	}
}

// SetFromChatID option function
func SetFromChatID(v int) MethodOption {
	return func(r *request) {
		r.params["from_chat_id"] = v
	}
}

// SetFromChatUsername option function
func SetFromChatUsername(v string) MethodOption {
	return func(r *request) {
		r.params["from_chat_id"] = v
	}
}

// SetMessageID option function
func SetMessageID(v int) MethodOption {
	return func(r *request) {
		r.params["message_id"] = v
	}
}

// SetPhotoFromFileID option function
func SetPhotoFromFileID(v string) MethodOption {
	return func(r *request) {
		r.params["photo"] = v
	}
}

// SetPhotoFromURL option function
func SetPhotoFromURL(v string) MethodOption {
	return func(r *request) {
		r.params["photo"] = v
	}
}

// SetPhoto option function
func SetPhoto(v InputFile) MethodOption {
	return func(r *request) {
		r.params["photo"] = v
	}
}

// SetCaption option function
func SetCaption(v string) MethodOption {
	return func(r *request) {
		r.params["caption"] = v
	}
}

// SetAudioFromFileID option function
func SetAudioFromFileID(v string) MethodOption {
	return func(r *request) {
		r.params["audio"] = v
	}
}

// SetAudioFromURL option function
func SetAudioFromURL(v string) MethodOption {
	return func(r *request) {
		r.params["audio"] = v
	}
}

// SetAudio option function
func SetAudio(v InputFile) MethodOption {
	return func(r *request) {
		r.params["audio"] = v
	}
}

// SetDuration option function
func SetDuration(v int) MethodOption {
	return func(r *request) {
		r.params["duration"] = v
	}
}

// SetPerformer option function
func SetPerformer(v string) MethodOption {
	return func(r *request) {
		r.params["performer"] = v
	}
}

// SetTitle option function
func SetTitle(v string) MethodOption {
	return func(r *request) {
		r.params["title"] = v
	}
}

// TODO: add Thumb string option

// SetThumb option function
func SetThumb(v InputFile) MethodOption {
	return func(r *request) {
		r.params["thumb"] = v
	}
}

// SetDocumentFromFileID option function
func SetDocumentFromFileID(v string) MethodOption {
	return func(r *request) {
		r.params["document"] = v
	}
}

// SetDocumentFromURL option function
func SetDocumentFromURL(v string) MethodOption {
	return func(r *request) {
		r.params["document"] = v
	}
}

// SetDocument option function
func SetDocument(v InputFile) MethodOption {
	return func(r *request) {
		r.params["document"] = v
	}
}

// SetVideoFromFileID option function
func SetVideoFromFileID(v string) MethodOption {
	return func(r *request) {
		r.params["video"] = v
	}
}

// SetVideoFromURL option function
func SetVideoFromURL(v string) MethodOption {
	return func(r *request) {
		r.params["video"] = v
	}
}

// SetVideo option function
func SetVideo(v InputFile) MethodOption {
	return func(r *request) {
		r.params["video"] = v
	}
}

// SetWidth option function
func SetWidth(v string) MethodOption {
	return func(r *request) {
		r.params["width"] = v
	}
}

// SetHeight option function
func SetHeight(v string) MethodOption {
	return func(r *request) {
		r.params["height"] = v
	}
}

// SupportsStreaming option function
func SupportsStreaming() MethodOption {
	return func(r *request) {
		r.params["supports_streaming"] = true
	}
}

// SetAnimationFromFileID option function
func SetAnimationFromFileID(v string) MethodOption {
	return func(r *request) {
		r.params["animation"] = v
	}
}

// SetAnimationFromURL option function
func SetAnimationFromURL(v string) MethodOption {
	return func(r *request) {
		r.params["animation"] = v
	}
}

// SetAnimation option function
func SetAnimation(v InputFile) MethodOption {
	return func(r *request) {
		r.params["animation"] = v
	}
}

// SetVoiceFromFileID option function
func SetVoiceFromFileID(v string) MethodOption {
	return func(r *request) {
		r.params["voice"] = v
	}
}

// SetVoiceFromURL option function
func SetVoiceFromURL(v string) MethodOption {
	return func(r *request) {
		r.params["voice"] = v
	}
}

// SetVoice option function
func SetVoice(v InputFile) MethodOption {
	return func(r *request) {
		r.params["voice"] = v
	}
}

// SetVideoNoteFromFileID option function
func SetVideoNoteFromFileID(v string) MethodOption {
	return func(r *request) {
		r.params["video_note"] = v
	}
}

// SetVideoNoteFromURL option function
func SetVideoNoteFromURL(v string) MethodOption {
	return func(r *request) {
		r.params["video_note"] = v
	}
}

// SetVideoNote option function
func SetVideoNote(v InputFile) MethodOption {
	return func(r *request) {
		r.params["video_note"] = v
	}
}

// SetMedia option function
func SetMedia(v ...InputMedia) MethodOption {
	return func(r *request) {
		r.params["media"] = v
	}
}

// SetLatitude option function
func SetLatitude(v float32) MethodOption {
	return func(r *request) {
		r.params["latitude"] = v
	}
}

// SetLongitude option function
func SetLongitude(v float32) MethodOption {
	return func(r *request) {
		r.params["longitude"] = v
	}
}

// SetHorizontalAccuracy option function
func SetHorizontalAccuracy(v float32) MethodOption {
	return func(r *request) {
		r.params["horizontal_accuracy"] = v
	}
}

// SetLivePeriod option function
func SetLivePeriod(v int) MethodOption {
	return func(r *request) {
		r.params["live_period"] = v
	}
}

// SetHeading option function
func SetHeading(v int) MethodOption {
	return func(r *request) {
		r.params["heading"] = v
	}
}

// SetProximityAlertRadius option function
func SetProximityAlertRadius(v int) MethodOption {
	return func(r *request) {
		r.params["proximity_alert_radius"] = v
	}
}

// SetInlineMessageID option function
func SetInlineMessageID(v string) MethodOption {
	return func(r *request) {
		r.params["inline_message_id"] = v
	}
}

// SetInlineQueryID option function
func SetInlineQueryID(v string) MethodOption {
	return func(r *request) {
		r.params["inline_query_id"] = v
	}
}

// SetResults option function
func SetResults(v []InlineQueryResult) MethodOption {
	return func(r *request) {
		r.params["results"] = v
	}
}

// SetAddress option function
func SetAddress(v string) MethodOption {
	return func(r *request) {
		r.params["address"] = v
	}
}

// SetFoursquareID option function
func SetFoursquareID(v string) MethodOption {
	return func(r *request) {
		r.params["foursquare_id"] = v
	}
}

// SetFoursquareType option function
func SetFoursquareType(v string) MethodOption {
	return func(r *request) {
		r.params["foursquare_type"] = v
	}
}

// SetGooglePlaceID option function
func SetGooglePlaceID(v string) MethodOption {
	return func(r *request) {
		r.params["google_place_id"] = v
	}
}

// SetGooglePlaceType option function
func SetGooglePlaceType(v string) MethodOption {
	return func(r *request) {
		r.params["google_place_type"] = v
	}
}

// SetPhoneNumber option function
func SetPhoneNumber(v string) MethodOption {
	return func(r *request) {
		r.params["phone_number"] = v
	}
}

// SetFirstName option function
func SetFirstName(v string) MethodOption {
	return func(r *request) {
		r.params["first_name"] = v
	}
}

// SetLastName option function
func SetLastName(v string) MethodOption {
	return func(r *request) {
		r.params["last_name"] = v
	}
}

// SetVCard option function
func SetVCard(v string) MethodOption {
	return func(r *request) {
		r.params["vcard"] = v
	}
}

// SetQuestion option function
func SetQuestion(v string) MethodOption {
	return func(r *request) {
		r.params["question"] = v
	}
}

// SetOptions option function
func SetOptions(v ...string) MethodOption {
	return func(r *request) {
		r.params["options"] = v
	}
}

// SetCommands option function
func SetCommands(v ...BotCommand) MethodOption {
	return func(r *request) {
		r.params["commands"] = v
	}
}

type Action string

const (
	ActionTyping          Action = "typing"
	ActionUploadPhoto     Action = "upload_photo"
	ActionRecordVideo     Action = "record_video"
	ActionUploadVideo     Action = "upload_video"
	ActionRecordVoice     Action = "record_voice"
	ActionUploadVoice     Action = "upload_voice"
	ActionUploadDocument  Action = "upload_document"
	ActionChooseSticker   Action = "choose_sticker"
	ActionFindLocation    Action = "find_location"
	ActionRecordVideoNote Action = "record_video_note"
	ActionUploadVideoNote Action = "upload_video_note"
)

// SetAction option function
func SetAction(v Action) MethodOption {
	return func(r *request) {
		r.params["action"] = v
	}
}

// SetFileID option function
func SetFileID(v string) MethodOption {
	return func(r *request) {
		r.params["file_id"] = v
	}
}

// SetUserID option function
func SetUserID(v int) MethodOption {
	return func(r *request) {
		r.params["user_id"] = v
	}
}

// SetUntilDate option function
func SetUntilDate(v int) MethodOption {
	return func(r *request) {
		r.params["until_date"] = v
	}
}

// AllowSendingWithoutReply option function
func AllowSendingWithoutReply() MethodOption {
	return func(r *request) {
		r.params["allow_sending_without_reply"] = true
	}
}

// CanSendMessages option function
func CanSendMessages() MethodOption {
	return func(r *request) {
		r.params["can_send_messages"] = true
	}
}

// CanSendMediaMessages option function
func CanSendMediaMessages() MethodOption {
	return func(r *request) {
		r.params["can_send_media_messages"] = true
	}
}

// CanSendOtherMessages option function
func CanSendOtherMessages() MethodOption {
	return func(r *request) {
		r.params["can_send_other_messages"] = true
	}
}

// CanAddWebPagePreviews option function
func CanAddWebPagePreviews() MethodOption {
	return func(r *request) {
		r.params["can_add_web_page_previews"] = true
	}
}

// CanChangeInfo option function
func CanChangeInfo() MethodOption {
	return func(r *request) {
		r.params["can_change_info"] = true
	}
}

// CanPostMessages option function
func CanPostMessages() MethodOption {
	return func(r *request) {
		r.params["can_post_messages"] = true
	}
}

// CanEditMessages option function
func CanEditMessages() MethodOption {
	return func(r *request) {
		r.params["can_edit_messages"] = true
	}
}

// CanDeleteMessages option function
func CanDeleteMessages() MethodOption {
	return func(r *request) {
		r.params["can_delete_messages"] = true
	}
}

// CanInviteUsers option function
func CanInviteUsers() MethodOption {
	return func(r *request) {
		r.params["can_invite_users"] = true
	}
}

// CanRestrictMembers option function
func CanRestrictMembers() MethodOption {
	return func(r *request) {
		r.params["can_restrict_members"] = true
	}
}

// CanPinMessages option function
func CanPinMessages() MethodOption {
	return func(r *request) {
		r.params["can_pin_messages"] = true
	}
}

// CanPromoteMembers option function
func CanPromoteMembers() MethodOption {
	return func(r *request) {
		r.params["can_promote_members"] = true
	}
}

// SetDescription option function
func SetDescription(v string) MethodOption {
	return func(r *request) {
		r.params["description"] = v
	}
}

// SetStickerSetName option function
func SetStickerSetName(v string) MethodOption {
	return func(r *request) {
		r.params["sticker_set_name"] = v
	}
}

// SetCallbackQueryID option function
func SetCallbackQueryID(v string) MethodOption {
	return func(r *request) {
		r.params["callback_query_id"] = v
	}
}

// ShowAlert option function
func ShowAlert() MethodOption {
	return func(r *request) {
		r.params["show_alert"] = true
	}
}

// SetCacheTime option function
func SetCacheTime(v int) MethodOption {
	return func(r *request) {
		r.params["cache_time"] = v
	}
}

// SetNextOffset option function
func SetNextOffset(v string) MethodOption {
	return func(r *request) {
		r.params["next_offset"] = v
	}
}

// SetSwitchPMText option function
func SetSwitchPMText(v string) MethodOption {
	return func(r *request) {
		r.params["switch_pm_text"] = v
	}
}

// SetSwitchPMParameter option function
func SetSwitchPMParameter(v string) MethodOption {
	return func(r *request) {
		r.params["switch_pm_parameter"] = v
	}
}

// IsPersonal option function
func IsPersonal() MethodOption {
	return func(r *request) {
		r.params["is_personal"] = true
	}
}

// SetPermissions option function
func SetPermissions(v ChatPermissions) MethodOption {
	return func(r *request) {
		r.params["permissions"] = v
	}
}

// SetCustomTitle option function
func SetCustomTitle(v string) MethodOption {
	return func(r *request) {
		r.params["custom_title"] = v
	}
}

// IsAnonymous option function
func IsAnonymous() MethodOption {
	return func(r *request) {
		r.params["is_anonymous"] = true
	}
}

// SetType option function
func SetType(v string) MethodOption {
	return func(r *request) {
		r.params["type"] = v
	}
}

// AllowsMultipleAnswers option function
func AllowsMultipleAnswers() MethodOption {
	return func(r *request) {
		r.params["allows_multiple_answers"] = true
	}
}

// SetCorrectOptionID option function
func SetCorrectOptionID(v int) MethodOption {
	return func(r *request) {
		r.params["correct_option_id"] = v
	}
}

// IsClosed option function
func IsClosed() MethodOption {
	return func(r *request) {
		r.params["is_closed"] = true
	}
}

// SetPNGStickerFromFileID option function
func SetPNGStickerFromFileID(v string) MethodOption {
	return func(r *request) {
		r.params["png_sticker"] = v
	}
}

// SetPNGStickerFromURL option function
func SetPNGStickerFromURL(v string) MethodOption {
	return func(r *request) {
		r.params["png_sticker"] = v
	}
}

// SetPNGSticker option function
func SetPNGSticker(v InputFile) MethodOption {
	return func(r *request) {
		r.params["png_sticker"] = v
	}
}

// SetTGSSticker option function
func SetTGSSticker(v InputFile) MethodOption {
	return func(r *request) {
		r.params["tgs_sticker"] = v
	}
}

// SetEmojis option function
func SetEmojis(v string) MethodOption {
	return func(r *request) {
		r.params["emojis"] = v
	}
}

// SetEmoji option function
func SetEmoji(v string) MethodOption {
	return func(r *request) {
		r.params["emoji"] = v
	}
}

// ContainsMasks option function
func ContainsMasks() MethodOption {
	return func(r *request) {
		r.params["contains_masks"] = true
	}
}

// SetMaskPosition option function
func SetMaskPosition(v MaskPosition) MethodOption {
	return func(r *request) {
		r.params["mask_position"] = v
	}
}

// SetExplanation option function
func SetExplanation(v string) MethodOption {
	return func(r *request) {
		r.params["explanation"] = v
	}
}

// SetExplanationParseMode option function
func SetExplanationParseMode(v string) MethodOption {
	return func(r *request) {
		r.params["explanation_parse_mode"] = v
	}
}

// SetOpenPeriod option function
func SetOpenPeriod(v int) MethodOption {
	return func(r *request) {
		r.params["open_period"] = v
	}
}

// SetCloseDate option function
func SetCloseDate(v int) MethodOption {
	return func(r *request) {
		r.params["close_date"] = v
	}
}

// SetExpireDate option function
func SetExpireDate(v int) MethodOption {
	return func(r *request) {
		r.params["expire_date"] = v
	}
}

// SetMemberLimit option function
func SetMemberLimit(v int) MethodOption {
	return func(r *request) {
		r.params["member_limit"] = v
	}
}

// SetInviteLink option function
func SetInviteLink(v string) MethodOption {
	return func(r *request) {
		r.params["invite_link"] = v
	}
}

// CanManageChat option function
func CanManageChat() MethodOption {
	return func(r *request) {
		r.params["can_manage_chat"] = true
	}
}

// CanManageVoiceChats option function
func CanManageVoiceChats() MethodOption {
	return func(r *request) {
		r.params["can_manage_voice_chats"] = true
	}
}

// RevokeMessages option function
func RevokeMessages() MethodOption {
	return func(r *request) {
		r.params["revoke_messages"] = true
	}
}

// SetPayload option function
func SetPayload(v string) MethodOption {
	return func(r *request) {
		r.params["payload"] = v
	}
}

// SetProviderToken option function
func SetProviderToken(v string) MethodOption {
	return func(r *request) {
		r.params["provider_token"] = v
	}
}

// SetCurrency option function
func SetCurrency(v Currency) MethodOption {
	return func(r *request) {
		r.params["currency"] = v
	}
}

// SetPrices option function
func SetPrices(v ...LabeledPrice) MethodOption {
	return func(r *request) {
		r.params["prices"] = v
	}
}

// SetMaxTipAmount option function
func SetMaxTipAmount(v int) MethodOption {
	return func(r *request) {
		r.params["max_tip_amount"] = v
	}
}

// SetSuggestedTipAmounts option function
func SetSuggestedTipAmounts(v ...int) MethodOption {
	return func(r *request) {
		r.params["suggested_tip_amounts"] = v
	}
}

// SetStartParameter option function
func SetStartParameter(v string) MethodOption {
	return func(r *request) {
		r.params["start_parameter"] = v
	}
}

// SetProviderData option function
func SetProviderData(v string) MethodOption {
	return func(r *request) {
		r.params["provider_data"] = v
	}
}

// SetPhotoURL option function
func SetPhotoURL(v string) MethodOption {
	return func(r *request) {
		r.params["photo_url"] = v
	}
}

// SetPhotoSize option function
func SetPhotoSize(v int) MethodOption {
	return func(r *request) {
		r.params["photo_size"] = v
	}
}

// SetPhotoWidth option function
func SetPhotoWidth(v int) MethodOption {
	return func(r *request) {
		r.params["photo_width"] = v
	}
}

// SetPhotoHeight option function
func SetPhotoHeight(v int) MethodOption {
	return func(r *request) {
		r.params["photo_height"] = v
	}
}

// NeedName option function
func NeedName() MethodOption {
	return func(r *request) {
		r.params["need_name"] = true
	}
}

// NeedPhoneNumber option function
func NeedPhoneNumber() MethodOption {
	return func(r *request) {
		r.params["need_phone_number"] = true
	}
}

// NeedEmail option function
func NeedEmail() MethodOption {
	return func(r *request) {
		r.params["need_email"] = true
	}
}

// NeedShippingAddress option function
func NeedShippingAddress() MethodOption {
	return func(r *request) {
		r.params["need_shipping_address"] = true
	}
}

// SendPhoneNumberToProvider option function
func SendPhoneNumberToProvider() MethodOption {
	return func(r *request) {
		r.params["send_phone_number_to_provider"] = true
	}
}

// SendEmailToProvider option function
func SendEmailToProvider() MethodOption {
	return func(r *request) {
		r.params["send_email_to_provider"] = true
	}
}

// IsFlexible option function
func IsFlexible() MethodOption {
	return func(r *request) {
		r.params["is_flexible"] = true
	}
}

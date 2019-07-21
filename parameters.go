package telegram

import (
	"encoding/json"
	"regexp"
	"strings"
)

// Option type
type Option func(*request)

func (r *request) setParam(k string, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		r.err = err
		return
	}
	s := string(b)
	if strings.Trim(s, "\"") != s {
		s = strings.Trim(s, "\"")
		s = regexp.MustCompile(`\\n`).ReplaceAllString(s, "\n")
	}
	err = r.writer.WriteField(k, s)
	if err != nil {
		r.err = err
	}
}

// SetOffset option function
func SetOffset(v int) Option {
	return func(r *request) {
		r.setParam("offset", v)
	}
}

// SetLimit option function
func SetLimit(v int) Option {
	return func(r *request) {
		r.setParam("limit", v)
	}
}

// SetTimeout option function
func SetTimeout(v int) Option {
	return func(r *request) {
		r.setParam("timeout", v)
	}
}

// SetAllowedUpdates option function
func SetAllowedUpdates(v ...string) Option {
	return func(r *request) {
		r.setParam("allowed_updates", v)
	}
}

// SetURL option function
func SetURL(v string) Option {
	return func(r *request) {
		r.setParam("url", v)
	}
}

// SetCertificate option function
func SetCertificate(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

// SetMaxConnections option function
func SetMaxConnections(v string) Option {
	return func(r *request) {
		r.setParam("max_connections", v)
	}
}

// SetChatID option function
func SetChatID(v int) Option {
	return func(r *request) {
		r.setParam("chat_id", v)
	}
}

// SetChatUsername option function
func SetChatUsername(v string) Option {
	return func(r *request) {
		r.setParam("chat_id", v)
	}
}

// SetText option function
func SetText(v string) Option {
	return func(r *request) {
		r.setParam("text", v)
	}
}

// SetParseMode option function
func SetParseMode(v string) Option {
	return func(r *request) {
		r.setParam("parse_mode", v)
	}
}

// DisableWebPagePreview option function
func DisableWebPagePreview() Option {
	return func(r *request) {
		r.setParam("disable_web_page_preview", "true")
	}
}

// DisableNotification option function
func DisableNotification() Option {
	return func(r *request) {
		r.setParam("disable_notification", "true")
	}
}

// SetReplyToMessageID option function
func SetReplyToMessageID(v int) Option {
	return func(r *request) {
		r.setParam("reply_to_message_id", v)
	}
}

// SetReplyMarkup option function
func SetReplyMarkup(v interface{}) Option {
	return func(r *request) {
		r.setParam("reply_markup", v)
	}
}

// SetFromChatID option function
func SetFromChatID(v int) Option {
	return func(r *request) {
		r.setParam("from_chat_id", v)
	}
}

// SetFromChatUsername option function
func SetFromChatUsername(v string) Option {
	return func(r *request) {
		r.setParam("from_chat_id", v)
	}
}

// SetMessageID option function
func SetMessageID(v int) Option {
	return func(r *request) {
		r.setParam("message_id", v)
	}
}

// SetPhotoFromFileID option function
func SetPhotoFromFileID(v string) Option {
	return func(r *request) {
		r.setParam("photo", v)
	}
}

// SetPhotoFromURL option function
func SetPhotoFromURL(v string) Option {
	return func(r *request) {
		r.setParam("photo", v)
	}
}

// SetPhoto option function
func SetPhoto(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

// SetCaption option function
func SetCaption(v string) Option {
	return func(r *request) {
		r.setParam("caption", v)
	}
}

// SetAudioFromFileID option function
func SetAudioFromFileID(v string) Option {
	return func(r *request) {
		r.setParam("audio", v)
	}
}

// SetAudioFromURL option function
func SetAudioFromURL(v string) Option {
	return func(r *request) {
		r.setParam("audio", v)
	}
}

// SetAudio option function
func SetAudio(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

// SetDuration option function
func SetDuration(v int) Option {
	return func(r *request) {
		r.setParam("duration", v)
	}
}

// SetPerformer option function
func SetPerformer(v string) Option {
	return func(r *request) {
		r.setParam("performer", v)
	}
}

// SetTitle option function
func SetTitle(v string) Option {
	return func(r *request) {
		r.setParam("title", v)
	}
}

// SetThumb option function
func SetThumb(v interface{}) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

// SetDocumentFromFileID option function
func SetDocumentFromFileID(v string) Option {
	return func(r *request) {
		r.setParam("document", v)
	}
}

// SetDocumentFromURL option function
func SetDocumentFromURL(v string) Option {
	return func(r *request) {
		r.setParam("document", v)
	}
}

// SetDocument option function
func SetDocument(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

// SetVideoFromFileID option function
func SetVideoFromFileID(v string) Option {
	return func(r *request) {
		r.setParam("video", v)
	}
}

// SetVideoFromURL option function
func SetVideoFromURL(v string) Option {
	return func(r *request) {
		r.setParam("video", v)
	}
}

// SetVideo option function
func SetVideo(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

// SetWidth option function
func SetWidth(v string) Option {
	return func(r *request) {
		r.setParam("width", v)
	}
}

// SetHeight option function
func SetHeight(v string) Option {
	return func(r *request) {
		r.setParam("height", v)
	}
}

// SupportsStreaming option function
func SupportsStreaming() Option {
	return func(r *request) {
		r.setParam("supports_streaming", "true")
	}
}

// SetAnimationFromFileID option function
func SetAnimationFromFileID(v string) Option {
	return func(r *request) {
		r.setParam("animation", v)
	}
}

// SetAnimationFromURL option function
func SetAnimationFromURL(v string) Option {
	return func(r *request) {
		r.setParam("animation", v)
	}
}

// SetAnimation option function
func SetAnimation(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

// SetVoiceFromFileID option function
func SetVoiceFromFileID(v string) Option {
	return func(r *request) {
		r.setParam("voice", v)
	}
}

// SetVoiceFromURL option function
func SetVoiceFromURL(v string) Option {
	return func(r *request) {
		r.setParam("voice", v)
	}
}

// SetVoice option function
func SetVoice(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

// SetVideoNoteFromFileID option function
func SetVideoNoteFromFileID(v string) Option {
	return func(r *request) {
		r.setParam("video_note", v)
	}
}

// SetVideoNoteFromURL option function
func SetVideoNoteFromURL(v string) Option {
	return func(r *request) {
		r.setParam("video_note", v)
	}
}

// SetVideoNote option function
func SetVideoNote(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

// SetMedia option function
func SetMedia(v ...InputMedia) Option {
	return func(r *request) {
		r.setParam("media", v)
	}
}

// SetLatitude option function
func SetLatitude(v float32) Option {
	return func(r *request) {
		r.setParam("latitude", v)
	}
}

// SetLongitude option function
func SetLongitude(v float32) Option {
	return func(r *request) {
		r.setParam("longitude", v)
	}
}

// SetLivePeriod option function
func SetLivePeriod(v int) Option {
	return func(r *request) {
		r.setParam("live_period", v)
	}
}

// SetInlineMessageID option function
func SetInlineMessageID(v string) Option {
	return func(r *request) {
		r.setParam("inline_message_id", v)
	}
}

// SetInlineQueryID option function
func SetInlineQueryID(v string) Option {
	return func(r *request) {
		r.setParam("inline_query_id", v)
	}
}

// SetResults option function
func SetResults(v []InlineQueryResult) Option {
	return func(r *request) {
		r.setParam("results", v)
	}
}

// SetAddress option function
func SetAddress(v string) Option {
	return func(r *request) {
		r.setParam("address", v)
	}
}

// SetFoursquareID option function
func SetFoursquareID(v string) Option {
	return func(r *request) {
		r.setParam("foursquare_id", v)
	}
}

// SetFoursquareType option function
func SetFoursquareType(v string) Option {
	return func(r *request) {
		r.setParam("foursquare_type", v)
	}
}

// SetPhoneNumber option function
func SetPhoneNumber(v string) Option {
	return func(r *request) {
		r.setParam("phone_number", v)
	}
}

// SetFirstName option function
func SetFirstName(v string) Option {
	return func(r *request) {
		r.setParam("first_name", v)
	}
}

// SetLastName option function
func SetLastName(v string) Option {
	return func(r *request) {
		r.setParam("last_name", v)
	}
}

// SetVCard option function
func SetVCard(v string) Option {
	return func(r *request) {
		r.setParam("vcard", v)
	}
}

// SetQuestion option function
func SetQuestion(v string) Option {
	return func(r *request) {
		r.setParam("question", v)
	}
}

// SetOptions option function
func SetOptions(v ...string) Option {
	return func(r *request) {
		r.setParam("options", v)
	}
}

// SetAction option function
func SetAction(v string) Option {
	return func(r *request) {
		r.setParam("action", v)
	}
}

// SetFileID option function
func SetFileID(v string) Option {
	return func(r *request) {
		r.setParam("file_id", v)
	}
}

// SetUserID option function
func SetUserID(v int) Option {
	return func(r *request) {
		r.setParam("user_id", v)
	}
}

// SetUntilDate option function
func SetUntilDate(v int) Option {
	return func(r *request) {
		r.setParam("until_date", v)
	}
}

// CanSendMessages option function
func CanSendMessages() Option {
	return func(r *request) {
		r.setParam("can_send_messages", "true")
	}
}

// CanSendMediaMessages option function
func CanSendMediaMessages() Option {
	return func(r *request) {
		r.setParam("can_send_media_messages", "true")
	}
}

// CanSendOtherMessages option function
func CanSendOtherMessages() Option {
	return func(r *request) {
		r.setParam("can_send_other_messages", "true")
	}
}

// CanAddWebPagePreviews option function
func CanAddWebPagePreviews() Option {
	return func(r *request) {
		r.setParam("can_add_web_page_previews", "true")
	}
}

// CanChangeInfo option function
func CanChangeInfo() Option {
	return func(r *request) {
		r.setParam("can_change_info", "true")
	}
}

// CanPostMessages option function
func CanPostMessages() Option {
	return func(r *request) {
		r.setParam("can_post_messages", "true")
	}
}

// CanEditMessages option function
func CanEditMessages() Option {
	return func(r *request) {
		r.setParam("can_edit_messages", "true")
	}
}

// CanDeleteMessages option function
func CanDeleteMessages() Option {
	return func(r *request) {
		r.setParam("can_delete_messages", "true")
	}
}

// CanInviteUsers option function
func CanInviteUsers() Option {
	return func(r *request) {
		r.setParam("can_invite_users", "true")
	}
}

// CanRestrictMemebers option function
func CanRestrictMemebers() Option {
	return func(r *request) {
		r.setParam("can_restrict_members", "true")
	}
}

// CanPinMessages option function
func CanPinMessages() Option {
	return func(r *request) {
		r.setParam("can_pin_messages", "true")
	}
}

// CanPromoteMemebers option function
func CanPromoteMemebers() Option {
	return func(r *request) {
		r.setParam("can_promote_members", "true")
	}
}

// SetDescription option function
func SetDescription(v string) Option {
	return func(r *request) {
		r.setParam("description", v)
	}
}

// SetStickerSetName option function
func SetStickerSetName(v string) Option {
	return func(r *request) {
		r.setParam("sticker_set_name", v)
	}
}

// SetCallbackQueryID option function
func SetCallbackQueryID(v string) Option {
	return func(r *request) {
		r.setParam("callback_query_id", v)
	}
}

// ShowAlert option function
func ShowAlert() Option {
	return func(r *request) {
		r.setParam("show_alert", "true")
	}
}

// SetCacheTime option function
func SetCacheTime(v int) Option {
	return func(r *request) {
		r.setParam("cache_time", v)
	}
}

// SetNextOffset option function
func SetNextOffset(v string) Option {
	return func(r *request) {
		r.setParam("next_offset", v)
	}
}

// SetSwitchPMText option function
func SetSwitchPMText(v string) Option {
	return func(r *request) {
		r.setParam("switch_pm_text", v)
	}
}

// SetSwitchPMParameter option function
func SetSwitchPMParameter(v string) Option {
	return func(r *request) {
		r.setParam("switch_pm_parameter", v)
	}
}

// IsPersonal option function
func IsPersonal() Option {
	return func(r *request) {
		r.setParam("is_personal", true)
	}
}

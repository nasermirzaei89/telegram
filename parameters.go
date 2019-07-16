package telegram

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Option type
type Option func(*request)

func (r *request) setString(k, v string) {
	err := r.writer.WriteField(k, v)
	if err != nil {
		r.err = err
	}
}

func (r *request) setInt(k string, v int) {
	err := r.writer.WriteField(k, fmt.Sprintf("%d", v))
	if err != nil {
		r.err = err
	}
}

func (r *request) setFloat(k string, v float32) {
	err := r.writer.WriteField(k, fmt.Sprintf("%f", v))
	if err != nil {
		r.err = err
	}
}

func (r *request) setStrings(k string, v ...string) {
	str := "[]"
	if len(v) == 0 {
		str = fmt.Sprintf(`["%s"]`, strings.Join(v, `","`))
	}
	err := r.writer.WriteField(k, str)
	if err != nil {
		r.err = err
	}
}

func (r *request) setObject(k string, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		r.err = err
		return
	}
	err = r.writer.WriteField(k, string(b))
	if err != nil {
		r.err = err
	}
}

// SetOffset option function
func SetOffset(v int) Option {
	return func(r *request) {
		r.setInt("offset", v)
	}
}

// SetLimit option function
func SetLimit(v int) Option {
	return func(r *request) {
		r.setInt("limit", v)
	}
}

// SetTimeout option function
func SetTimeout(v int) Option {
	return func(r *request) {
		r.setInt("timeout", v)
	}
}

// SetAllowedUpdates option function
func SetAllowedUpdates(v ...string) Option {
	return func(r *request) {
		r.setStrings("allowed_updates", v...)
	}
}

// SetURL option function
func SetURL(v string) Option {
	return func(r *request) {
		r.setString("url", v)
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
		r.setString("max_connections", v)
	}
}

// SetChatID option function
func SetChatID(v int) Option {
	return func(r *request) {
		r.setInt("chat_id", v)
	}
}

// SetChatUsername option function
func SetChatUsername(v string) Option {
	return func(r *request) {
		r.setString("chat_id", v)
	}
}

// SetText option function
func SetText(v string) Option {
	return func(r *request) {
		r.setString("text", v)
	}
}

// SetParseMode option function
func SetParseMode(v string) Option {
	return func(r *request) {
		r.setString("parse_mode", v)
	}
}

// DisableWebPagePreview option function
func DisableWebPagePreview() Option {
	return func(r *request) {
		r.setString("disable_web_page_preview", "true")
	}
}

// DisableNotification option function
func DisableNotification() Option {
	return func(r *request) {
		r.setString("disable_notification", "true")
	}
}

// SetReplyToMessageID option function
func SetReplyToMessageID(v int) Option {
	return func(r *request) {
		r.setInt("reply_to_message_id", v)
	}
}

// SetReplyMarkup option function
func SetReplyMarkup(v interface{}) Option {
	return func(r *request) {
		r.setObject("reply_markup", v)
	}
}

// SetFromChatID option function
func SetFromChatID(v int) Option {
	return func(r *request) {
		r.setInt("from_chat_id", v)
	}
}

// SetFromChatUsername option function
func SetFromChatUsername(v string) Option {
	return func(r *request) {
		r.setString("from_chat_id", v)
	}
}

// SetMessageID option function
func SetMessageID(v int) Option {
	return func(r *request) {
		r.setInt("message_id", v)
	}
}

// SetPhotoFromFileID option function
func SetPhotoFromFileID(v string) Option {
	return func(r *request) {
		r.setString("photo", v)
	}
}

// SetPhotoFromURL option function
func SetPhotoFromURL(v string) Option {
	return func(r *request) {
		r.setString("photo", v)
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
		r.setString("caption", v)
	}
}

// SetAudioFromFileID option function
func SetAudioFromFileID(v string) Option {
	return func(r *request) {
		r.setString("audio", v)
	}
}

// SetAudioFromURL option function
func SetAudioFromURL(v string) Option {
	return func(r *request) {
		r.setString("audio", v)
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
		r.setInt("duration", v)
	}
}

// SetPerformer option function
func SetPerformer(v string) Option {
	return func(r *request) {
		r.setString("performer", v)
	}
}

// SetTitle option function
func SetTitle(v string) Option {
	return func(r *request) {
		r.setString("title", v)
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
		r.setString("document", v)
	}
}

// SetDocumentFromURL option function
func SetDocumentFromURL(v string) Option {
	return func(r *request) {
		r.setString("document", v)
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
		r.setString("video", v)
	}
}

// SetVideoFromURL option function
func SetVideoFromURL(v string) Option {
	return func(r *request) {
		r.setString("video", v)
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
		r.setString("width", v)
	}
}

// SetHeight option function
func SetHeight(v string) Option {
	return func(r *request) {
		r.setString("height", v)
	}
}

// SupportsStreaming option function
func SupportsStreaming() Option {
	return func(r *request) {
		r.setString("supports_streaming", "true")
	}
}

// SetAnimationFromFileID option function
func SetAnimationFromFileID(v string) Option {
	return func(r *request) {
		r.setString("animation", v)
	}
}

// SetAnimationFromURL option function
func SetAnimationFromURL(v string) Option {
	return func(r *request) {
		r.setString("animation", v)
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
		r.setString("voice", v)
	}
}

// SetVoiceFromURL option function
func SetVoiceFromURL(v string) Option {
	return func(r *request) {
		r.setString("voice", v)
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
		r.setString("video_note", v)
	}
}

// SetVideoNoteFromURL option function
func SetVideoNoteFromURL(v string) Option {
	return func(r *request) {
		r.setString("video_note", v)
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
		r.setObject("media", v)
	}
}

// SetLatitude option function
func SetLatitude(v float32) Option {
	return func(r *request) {
		r.setFloat("latitude", v)
	}
}

// SetLongitude option function
func SetLongitude(v float32) Option {
	return func(r *request) {
		r.setFloat("longitude", v)
	}
}

// SetLivePeriod option function
func SetLivePeriod(v int) Option {
	return func(r *request) {
		r.setInt("live_period", v)
	}
}

// SetInlineMessageID option function
func SetInlineMessageID(v string) Option {
	return func(r *request) {
		r.setString("inline_message_id", v)
	}
}

// SetAddress option function
func SetAddress(v string) Option {
	return func(r *request) {
		r.setString("address", v)
	}
}

// SetFoursquareID option function
func SetFoursquareID(v string) Option {
	return func(r *request) {
		r.setString("foursquare_id", v)
	}
}

// SetFoursquareType option function
func SetFoursquareType(v string) Option {
	return func(r *request) {
		r.setString("foursquare_type", v)
	}
}

// SetPhoneNumber option function
func SetPhoneNumber(v string) Option {
	return func(r *request) {
		r.setString("phone_number", v)
	}
}

// SetFirstName option function
func SetFirstName(v string) Option {
	return func(r *request) {
		r.setString("first_name", v)
	}
}

// SetLastName option function
func SetLastName(v string) Option {
	return func(r *request) {
		r.setString("last_name", v)
	}
}

// SetVCard option function
func SetVCard(v string) Option {
	return func(r *request) {
		r.setString("vcard", v)
	}
}

// SetQuestion option function
func SetQuestion(v string) Option {
	return func(r *request) {
		r.setString("question", v)
	}
}

// SetOptions option function
func SetOptions(v ...string) Option {
	return func(r *request) {
		r.setStrings("options", v...)
	}
}

// SetAction option function
func SetAction(v string) Option {
	return func(r *request) {
		r.setString("action", v)
	}
}

// SetFileID option function
func SetFileID(v string) Option {
	return func(r *request) {
		r.setString("file_id", v)
	}
}

// SetUserID option function
func SetUserID(v int) Option {
	return func(r *request) {
		r.setInt("user_id", v)
	}
}

// SetUntilDate option function
func SetUntilDate(v int) Option {
	return func(r *request) {
		r.setInt("until_date", v)
	}
}

// CanSendMessages option function
func CanSendMessages() Option {
	return func(r *request) {
		r.setString("can_send_messages", "true")
	}
}

// CanSendMediaMessages option function
func CanSendMediaMessages() Option {
	return func(r *request) {
		r.setString("can_send_media_messages", "true")
	}
}

// CanSendOtherMessages option function
func CanSendOtherMessages() Option {
	return func(r *request) {
		r.setString("can_send_other_messages", "true")
	}
}

// CanAddWebPagePreviews option function
func CanAddWebPagePreviews() Option {
	return func(r *request) {
		r.setString("can_add_web_page_previews", "true")
	}
}

// CanChangeInfo option function
func CanChangeInfo() Option {
	return func(r *request) {
		r.setString("can_change_info", "true")
	}
}

// CanPostMessages option function
func CanPostMessages() Option {
	return func(r *request) {
		r.setString("can_post_messages", "true")
	}
}

// CanEditMessages option function
func CanEditMessages() Option {
	return func(r *request) {
		r.setString("can_edit_messages", "true")
	}
}

// CanDeleteMessages option function
func CanDeleteMessages() Option {
	return func(r *request) {
		r.setString("can_delete_messages", "true")
	}
}

// CanInviteUsers option function
func CanInviteUsers() Option {
	return func(r *request) {
		r.setString("can_invite_users", "true")
	}
}

// CanRestrictMemebers option function
func CanRestrictMemebers() Option {
	return func(r *request) {
		r.setString("can_restrict_members", "true")
	}
}

// CanPinMessages option function
func CanPinMessages() Option {
	return func(r *request) {
		r.setString("can_pin_messages", "true")
	}
}

// CanPromoteMemebers option function
func CanPromoteMemebers() Option {
	return func(r *request) {
		r.setString("can_promote_members", "true")
	}
}

// SetDescription option function
func SetDescription(v string) Option {
	return func(r *request) {
		r.setString("description", v)
	}
}

// SetStickerSetName option function
func SetStickerSetName(v string) Option {
	return func(r *request) {
		r.setString("sticker_set_name", v)
	}
}

// SetCallbackQueryID option function
func SetCallbackQueryID(v string) Option {
	return func(r *request) {
		r.setString("callback_query_id", v)
	}
}

// ShowAlert option function
func ShowAlert() Option {
	return func(r *request) {
		r.setString("show_alert", "true")
	}
}

// SetCacheTime option function
func SetCacheTime(v int) Option {
	return func(r *request) {
		r.setInt("cache_time", v)
	}
}

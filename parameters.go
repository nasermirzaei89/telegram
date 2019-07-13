package telegram

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func SetOffset(v int) Option {
	return func(r *request) {
		r.setInt("offset", v)
	}
}

func SetLimit(v int) Option {
	return func(r *request) {
		r.setInt("limit", v)
	}
}

func SetTimeout(v int) Option {
	return func(r *request) {
		r.setInt("timeout", v)
	}
}

func SetAllowedUpdates(v ...string) Option {
	return func(r *request) {
		r.setStrings("allowed_updates", v...)
	}
}

func SetURL(v string) Option {
	return func(r *request) {
		r.setString("url", v)
	}
}

func SetCertificate(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

func SetMaxConnections(v string) Option {
	return func(r *request) {
		r.setString("max_connections", v)
	}
}

func SetChatID(v int) Option {
	return func(r *request) {
		r.setInt("chat_id", v)
	}
}

func SetChatUsername(v string) Option {
	return func(r *request) {
		r.setString("chat_id", v)
	}
}

func SetText(v string) Option {
	return func(r *request) {
		r.setString("text", v)
	}
}

func SetParseMode(v string) Option {
	return func(r *request) {
		r.setString("parse_mode", v)
	}
}

func DisableWebPagePreview() Option {
	return func(r *request) {
		r.setString("disable_web_page_preview", "true")
	}
}

func DisableNotification() Option {
	return func(r *request) {
		r.setString("disable_notification", "true")
	}
}

func SetReplyToMessageID(v int) Option {
	return func(r *request) {
		r.setInt("reply_to_message_id", v)
	}
}

func SetReplyMarkup(v interface{}) Option {
	return func(r *request) {
		r.setObject("reply_markup", v)
	}
}

func SetFromChatID(v int) Option {
	return func(r *request) {
		r.setInt("from_chat_id", v)
	}
}

func SetFromChatUsername(v string) Option {
	return func(r *request) {
		r.setString("from_chat_id", v)
	}
}

func SetMessageID(v int) Option {
	return func(r *request) {
		r.setInt("message_id", v)
	}
}

func SetPhotoFromFileID(v string) Option {
	return func(r *request) {
		r.setString("photo", v)
	}
}

func SetPhotoFromURL(v string) Option {
	return func(r *request) {
		r.setString("photo", v)
	}
}

func SetPhoto(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

func SetCaption(v string) Option {
	return func(r *request) {
		r.setString("caption", v)
	}
}

func SetAudioFromFileID(v string) Option {
	return func(r *request) {
		r.setString("audio", v)
	}
}

func SetAudioFromURL(v string) Option {
	return func(r *request) {
		r.setString("audio", v)
	}
}

func SetAudio(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

func SetDuration(v int) Option {
	return func(r *request) {
		r.setInt("duration", v)
	}
}

func SetPerformer(v string) Option {
	return func(r *request) {
		r.setString("performer", v)
	}
}

func SetTitle(v string) Option {
	return func(r *request) {
		r.setString("title", v)
	}
}

func SetThumb(v interface{}) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

func SetDocumentFromFileID(v string) Option {
	return func(r *request) {
		r.setString("document", v)
	}
}

func SetDocumentFromURL(v string) Option {
	return func(r *request) {
		r.setString("document", v)
	}
}

func SetDocument(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

func SetVideoFromFileID(v string) Option {
	return func(r *request) {
		r.setString("video", v)
	}
}

func SetVideoFromURL(v string) Option {
	return func(r *request) {
		r.setString("video", v)
	}
}

func SetVideo(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

func SetWidth(v string) Option {
	return func(r *request) {
		r.setString("width", v)
	}
}

func SetHeight(v string) Option {
	return func(r *request) {
		r.setString("height", v)
	}
}

func SupportsStreaming() Option {
	return func(r *request) {
		r.setString("supports_streaming", "true")
	}
}

func SetAnimationFromFileID(v string) Option {
	return func(r *request) {
		r.setString("animation", v)
	}
}

func SetAnimationFromURL(v string) Option {
	return func(r *request) {
		r.setString("animation", v)
	}
}

func SetAnimation(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

func SetVoiceFromFileID(v string) Option {
	return func(r *request) {
		r.setString("voice", v)
	}
}

func SetVoiceFromURL(v string) Option {
	return func(r *request) {
		r.setString("voice", v)
	}
}

func SetVoice(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

func SetVideoNoteFromFileID(v string) Option {
	return func(r *request) {
		r.setString("video_note", v)
	}
}

func SetVideoNoteFromURL(v string) Option {
	return func(r *request) {
		r.setString("video_note", v)
	}
}

func SetVideoNote(v InputFile) Option {
	return func(r *request) {
		// TODO: implement me
		panic("implement me")
	}
}

func SetMedia(v ...InputMedia) Option {
	return func(r *request) {
		r.setObject("media", v)
	}
}

func SetLatitude(v float32) Option {
	return func(r *request) {
		r.setFloat("latitude", v)
	}
}

func SetLongitude(v float32) Option {
	return func(r *request) {
		r.setFloat("longitude", v)
	}
}

func SetLivePeriod(v int) Option {
	return func(r *request) {
		r.setInt("live_period", v)
	}
}

func SetInlineMessageID(v string) Option {
	return func(r *request) {
		r.setString("inline_message_id", v)
	}
}

func SetAddress(v string) Option {
	return func(r *request) {
		r.setString("address", v)
	}
}

func SetFoursquareID(v string) Option {
	return func(r *request) {
		r.setString("foursquare_id", v)
	}
}

func SetFoursquareType(v string) Option {
	return func(r *request) {
		r.setString("foursquare_type", v)
	}
}

func SetPhoneNumber(v string) Option {
	return func(r *request) {
		r.setString("phone_number", v)
	}
}

func SetFirstName(v string) Option {
	return func(r *request) {
		r.setString("first_name", v)
	}
}

func SetLastName(v string) Option {
	return func(r *request) {
		r.setString("last_name", v)
	}
}

func SetVCard(v string) Option {
	return func(r *request) {
		r.setString("vcard", v)
	}
}

func SetQuestion(v string) Option {
	return func(r *request) {
		r.setString("question", v)
	}
}

func SetOptions(v ...string) Option {
	return func(r *request) {
		r.setStrings("options", v...)
	}
}

func SetAction(v string) Option {
	return func(r *request) {
		r.setString("action", v)
	}
}

func SetFileID(v string) Option {
	return func(r *request) {
		r.setString("file_id", v)
	}
}

func SetUserID(v int) Option {
	return func(r *request) {
		r.setInt("user_id", v)
	}
}

func SetUntilDate(v int) Option {
	return func(r *request) {
		r.setInt("until_date", v)
	}
}

func CanSendMessages() Option {
	return func(r *request) {
		r.setString("can_send_messages", "true")
	}
}

func CanSendMediaMessages() Option {
	return func(r *request) {
		r.setString("can_send_media_messages", "true")
	}
}

func CanSendOtherMessages() Option {
	return func(r *request) {
		r.setString("can_send_other_messages", "true")
	}
}

func CanAddWebPagePreviews() Option {
	return func(r *request) {
		r.setString("can_add_web_page_previews", "true")
	}
}

func CanChangeInfo() Option {
	return func(r *request) {
		r.setString("can_change_info", "true")
	}
}

func CanPostMessages() Option {
	return func(r *request) {
		r.setString("can_post_messages", "true")
	}
}

func CanEditMessages() Option {
	return func(r *request) {
		r.setString("can_edit_messages", "true")
	}
}

func CanDeleteMessages() Option {
	return func(r *request) {
		r.setString("can_delete_messages", "true")
	}
}

func CanInviteUsers() Option {
	return func(r *request) {
		r.setString("can_invite_users", "true")
	}
}

func CanRestrictMemebers() Option {
	return func(r *request) {
		r.setString("can_restrict_members", "true")
	}
}

func CanPinMessages() Option {
	return func(r *request) {
		r.setString("can_pin_messages", "true")
	}
}

func CanPromoteMemebers() Option {
	return func(r *request) {
		r.setString("can_promote_members", "true")
	}
}

func SetDescription(v string) Option {
	return func(r *request) {
		r.setString("description", v)
	}
}

func SetStickerSetName(v string) Option {
	return func(r *request) {
		r.setString("sticker_set_name", v)
	}
}

func SetCallbackQueryID(v string) Option {
	return func(r *request) {
		r.setString("callback_query_id", v)
	}
}

func ShowAlert() Option {
	return func(r *request) {
		r.setString("show_alert", "true")
	}
}

func SetCacheTime(v int) Option {
	return func(r *request) {
		r.setInt("cache_time", v)
	}
}

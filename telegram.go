package telegram

// Bot interface
type Bot interface {
	GetToken() string

	// getting updates
	GetUpdates(...Option) (GetUpdatesResponse, error)
	SetWebhook(...Option) (SetWebhookResponse, error)
	DeleteWebhook() (DeleteWebhookResponse, error)
	GetWebhookInfo() (GetWebhookInfoResponse, error)

	// available methods
	GetMe() (GetMeResponse, error)
	SendMessage(...Option) (SendMessageResponse, error)
	ForwardMessage(...Option) (ForwardMessageResponse, error)
	SendPhoto(...Option) (SendPhotoResponse, error)
	SendAudio(...Option) (SendAudioResponse, error)
	SendDocument(...Option) (SendDocumentResponse, error)
	SendVideo(...Option) (SendVideoResponse, error)
	SendAnimation(...Option) (SendAnimationResponse, error)
	SendVoice(...Option) (SendVoiceResponse, error)
	SendVideoNote(...Option) (SendVideoNoteResponse, error)
	SendMediaGroup(...Option) (SendMediaGroupResponse, error)
	SendLocation(...Option) (SendLocationResponse, error)
	EditMessageLiveLocation(...Option) (EditMessageLiveLocationResponse, error)
	StopMessageLiveLocation(...Option) (StopMessageLiveLocationResponse, error)
	SendVenue(...Option) (SendVenueResponse, error)
	SendContact(...Option) (SendContactResponse, error)
	SendPoll(...Option) (SendPollResponse, error)
	SendChatAction(...Option) (SendChatActionResponse, error)
	GetUserProfilePhotos(...Option) (GetUserProfilePhotosResponse, error)
	GetFile(...Option) (GetFileResponse, error)
	KickChatMember(...Option) (KickChatMemberResponse, error)
	UnbanChatMember(...Option) (UnbanChatMemberResponse, error)
	RestrictChatMember(...Option) (RestrictChatMemberResponse, error)
	PromoteChatMember(...Option) (PromoteChatMemberResponse, error)
	ExportChatInviteLink(...Option) (ExportChatInviteLinkResponse, error)
	SetChatPhoto(...Option) (SetChatPhotoResponse, error)
	DeleteChatPhoto(...Option) (DeleteChatPhotoResponse, error)
	SetChatTitle(...Option) (SetChatTitleResponse, error)
	SetChatDescription(...Option) (SetChatDescriptionResponse, error)
	PinChatMessage(...Option) (PinChatMessageResponse, error)
	UnpinChatMessage(...Option) (UnpinChatMessageResponse, error)
	LeaveChat(...Option) (LeaveChatResponse, error)
	GetChat(...Option) (GetChatResponse, error)
	GetChatAdministrators(...Option) (GetChatAdministratorsResponse, error)
	GetChatMembersCount(...Option) (GetChatMembersCountResponse, error)
	GetChatMember(...Option) (GetChatMemberResponse, error)
	SetChatStickerSet(...Option) (SetChatStickerSetResponse, error)
	DeleteChatStickerSet(...Option) (DeleteChatStickerSetResponse, error)
	AnswerCallbackQuery(...Option) (AnswerCallbackQueryResponse, error)

	// updating messages
	EditMessageText(...Option) (EditMessageTextResponse, error)
	EditMessageCaption(...Option) (EditMessageCaptionResponse, error)
	EditMessageMedia(...Option) (EditMessageMediaResponse, error)
	EditMessageReplyMarkup(...Option) (EditMessageReplyMarkupResponse, error)
	StopPoll(...Option) (StopPollResponse, error)
	DeleteMessage(...Option) (DeleteMessageResponse, error)

	// stickers
	SendSticker(...Option) (SendStickerResponse, error)
	GetStickerSet(...Option) (GetStickerSetResponse, error)
	UploadStickerFile(...Option) (UploadStickerFileResponse, error)
	CreateNewStickerSet(...Option) (CreateNewStickerSetResponse, error)
	AddStickerToSet(...Option) (AddStickerToSetResponse, error)
	SetStickerPositionInSet(...Option) (SetStickerPositionInSetResponse, error)
	DeleteStickerFromSet(...Option) (DeleteStickerFromSetResponse, error)

	// inline mode
	AnswerInlineQuery(...Option) (AnswerInlineQueryResponse, error)

	// payments
	SendInvoice(...Option) (SendInvoiceResponse, error)
	AnswerShippingQuery(...Option) (AnswerShippingQueryResponse, error)
	AnswerPreCheckoutQuery(...Option) (AnswerPreCheckoutQueryResponse, error)

	// telegram passport
	SetPassportDataErrors(...Option) (SetPassportDataErrorsResponse, error)

	// games
	SendGame(...Option) (SendGameResponse, error)
	SetGameScore(...Option) (SetGameScoreResponse, error)
	GetGameHighScores(...Option) (GetGameHighScoresResponse, error)
}

type bot struct {
	token string
}

// New return a telegram bot instance
func New(token string) Bot {
	return &bot{token: token}
}

func (b *bot) GetToken() string {
	return b.token
}

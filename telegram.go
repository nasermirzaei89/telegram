package telegram

// Bot interface
type Bot interface {
	GetToken() string

	// getting updates
	GetUpdates(...MethodOption) (GetUpdatesResponse, error)
	SetWebhook(...MethodOption) (SetWebhookResponse, error)
	DeleteWebhook() (DeleteWebhookResponse, error)
	GetWebhookInfo() (GetWebhookInfoResponse, error)

	// available methods
	GetMe() (GetMeResponse, error)
	SendMessage(...MethodOption) (SendMessageResponse, error)
	ForwardMessage(...MethodOption) (ForwardMessageResponse, error)
	SendPhoto(...MethodOption) (SendPhotoResponse, error)
	SendAudio(...MethodOption) (SendAudioResponse, error)
	SendDocument(...MethodOption) (SendDocumentResponse, error)
	SendVideo(...MethodOption) (SendVideoResponse, error)
	SendAnimation(...MethodOption) (SendAnimationResponse, error)
	SendVoice(...MethodOption) (SendVoiceResponse, error)
	SendVideoNote(...MethodOption) (SendVideoNoteResponse, error)
	SendMediaGroup(...MethodOption) (SendMediaGroupResponse, error)
	SendLocation(...MethodOption) (SendLocationResponse, error)
	EditMessageLiveLocation(...MethodOption) (EditMessageLiveLocationResponse, error)
	StopMessageLiveLocation(...MethodOption) (StopMessageLiveLocationResponse, error)
	SendVenue(...MethodOption) (SendVenueResponse, error)
	SendContact(...MethodOption) (SendContactResponse, error)
	SendPoll(...MethodOption) (SendPollResponse, error)
	SendChatAction(...MethodOption) (SendChatActionResponse, error)
	GetUserProfilePhotos(...MethodOption) (GetUserProfilePhotosResponse, error)
	GetFile(...MethodOption) (GetFileResponse, error)
	KickChatMember(...MethodOption) (KickChatMemberResponse, error)
	UnbanChatMember(...MethodOption) (UnbanChatMemberResponse, error)
	RestrictChatMember(...MethodOption) (RestrictChatMemberResponse, error)
	PromoteChatMember(...MethodOption) (PromoteChatMemberResponse, error)
	ExportChatInviteLink(...MethodOption) (ExportChatInviteLinkResponse, error)
	SetChatPhoto(...MethodOption) (SetChatPhotoResponse, error)
	DeleteChatPhoto(...MethodOption) (DeleteChatPhotoResponse, error)
	SetChatTitle(...MethodOption) (SetChatTitleResponse, error)
	SetChatDescription(...MethodOption) (SetChatDescriptionResponse, error)
	PinChatMessage(...MethodOption) (PinChatMessageResponse, error)
	UnpinChatMessage(...MethodOption) (UnpinChatMessageResponse, error)
	LeaveChat(...MethodOption) (LeaveChatResponse, error)
	GetChat(...MethodOption) (GetChatResponse, error)
	GetChatAdministrators(...MethodOption) (GetChatAdministratorsResponse, error)
	GetChatMembersCount(...MethodOption) (GetChatMembersCountResponse, error)
	GetChatMember(...MethodOption) (GetChatMemberResponse, error)
	SetChatStickerSet(...MethodOption) (SetChatStickerSetResponse, error)
	DeleteChatStickerSet(...MethodOption) (DeleteChatStickerSetResponse, error)
	AnswerCallbackQuery(...MethodOption) (AnswerCallbackQueryResponse, error)

	// updating messages
	EditMessageText(...MethodOption) (EditMessageTextResponse, error)
	EditMessageCaption(...MethodOption) (EditMessageCaptionResponse, error)
	EditMessageMedia(...MethodOption) (EditMessageMediaResponse, error)
	EditMessageReplyMarkup(...MethodOption) (EditMessageReplyMarkupResponse, error)
	StopPoll(...MethodOption) (StopPollResponse, error)
	DeleteMessage(...MethodOption) (DeleteMessageResponse, error)

	// stickers
	SendSticker(...MethodOption) (SendStickerResponse, error)
	GetStickerSet(...MethodOption) (GetStickerSetResponse, error)
	UploadStickerFile(...MethodOption) (UploadStickerFileResponse, error)
	CreateNewStickerSet(...MethodOption) (CreateNewStickerSetResponse, error)
	AddStickerToSet(...MethodOption) (AddStickerToSetResponse, error)
	SetStickerPositionInSet(...MethodOption) (SetStickerPositionInSetResponse, error)
	DeleteStickerFromSet(...MethodOption) (DeleteStickerFromSetResponse, error)

	// inline mode
	AnswerInlineQuery(...MethodOption) (AnswerInlineQueryResponse, error)

	// payments
	SendInvoice(...MethodOption) (SendInvoiceResponse, error)
	AnswerShippingQuery(...MethodOption) (AnswerShippingQueryResponse, error)
	AnswerPreCheckoutQuery(...MethodOption) (AnswerPreCheckoutQueryResponse, error)

	// telegram passport
	SetPassportDataErrors(...MethodOption) (SetPassportDataErrorsResponse, error)

	// games
	SendGame(...MethodOption) (SendGameResponse, error)
	SetGameScore(...MethodOption) (SetGameScoreResponse, error)
	GetGameHighScores(...MethodOption) (GetGameHighScoresResponse, error)
}

type BotOption func(*bot)

type bot struct {
	token   string
	baseURL string
}

// New return a telegram bot instance
func New(token string, options ...BotOption) Bot {
	b := bot{
		token:   token,
		baseURL: "https://api.telegram.org",
	}

	for i := range options {
		options[i](&b)
	}

	return &b
}

func (b *bot) GetToken() string {
	return b.token
}

func SetBaseURL(v string) BotOption {
	return func(b *bot) {
		b.baseURL = v
	}
}

func (b *bot) GetBaseURL() string {
	return b.baseURL
}

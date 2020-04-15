package telegram

// Bot interface
type Bot interface {
	GetToken() string

	// getting updates
	GetUpdates(options ...MethodOption) (res GetUpdatesResponse, err error)
	SetWebhook(options ...MethodOption) (res SetWebhookResponse, err error)
	DeleteWebhook() (res DeleteWebhookResponse, err error)
	GetWebhookInfo() (res GetWebhookInfoResponse, err error)

	// available methods
	GetMe() (res GetMeResponse, err error)
	SendMessage(options ...MethodOption) (res SendMessageResponse, err error)
	ForwardMessage(options ...MethodOption) (res ForwardMessageResponse, err error)
	SendPhoto(options ...MethodOption) (res SendPhotoResponse, err error)
	SendAudio(options ...MethodOption) (res SendAudioResponse, err error)
	SendDocument(options ...MethodOption) (res SendDocumentResponse, err error)
	SendVideo(options ...MethodOption) (res SendVideoResponse, err error)
	SendAnimation(options ...MethodOption) (res SendAnimationResponse, err error)
	SendVoice(options ...MethodOption) (res SendVoiceResponse, err error)
	SendVideoNote(options ...MethodOption) (res SendVideoNoteResponse, err error)
	SendMediaGroup(options ...MethodOption) (res SendMediaGroupResponse, err error)
	SendLocation(options ...MethodOption) (res SendLocationResponse, err error)
	EditMessageLiveLocation(options ...MethodOption) (res EditMessageLiveLocationResponse, err error)
	StopMessageLiveLocation(options ...MethodOption) (res StopMessageLiveLocationResponse, err error)
	SendVenue(options ...MethodOption) (res SendVenueResponse, err error)
	SendContact(options ...MethodOption) (res SendContactResponse, err error)
	SendPoll(options ...MethodOption) (res SendPollResponse, err error)
	SendDice(options ...MethodOption) (res SendDiceResponse, err error)
	SendChatAction(options ...MethodOption) (res SendChatActionResponse, err error)
	GetUserProfilePhotos(options ...MethodOption) (res GetUserProfilePhotosResponse, err error)
	GetFile(options ...MethodOption) (res GetFileResponse, err error)
	KickChatMember(options ...MethodOption) (res KickChatMemberResponse, err error)
	UnbanChatMember(options ...MethodOption) (res UnbanChatMemberResponse, err error)
	RestrictChatMember(options ...MethodOption) (res RestrictChatMemberResponse, err error)
	PromoteChatMember(options ...MethodOption) (res PromoteChatMemberResponse, err error)
	SetChatAdministratorCustomTitle(options ...MethodOption) (res SetChatAdministratorCustomTitleResponse, err error)
	ExportChatInviteLink(options ...MethodOption) (res ExportChatInviteLinkResponse, err error)
	SetChatPermissions(options ...MethodOption) (res SetChatPermissionsResponse, err error)
	SetChatPhoto(options ...MethodOption) (res SetChatPhotoResponse, err error)
	DeleteChatPhoto(options ...MethodOption) (res DeleteChatPhotoResponse, err error)
	SetChatTitle(options ...MethodOption) (res SetChatTitleResponse, err error)
	SetChatDescription(options ...MethodOption) (res SetChatDescriptionResponse, err error)
	PinChatMessage(options ...MethodOption) (res PinChatMessageResponse, err error)
	UnpinChatMessage(options ...MethodOption) (res UnpinChatMessageResponse, err error)
	LeaveChat(options ...MethodOption) (res LeaveChatResponse, err error)
	GetChat(options ...MethodOption) (res GetChatResponse, err error)
	GetChatAdministrators(options ...MethodOption) (res GetChatAdministratorsResponse, err error)
	GetChatMembersCount(options ...MethodOption) (res GetChatMembersCountResponse, err error)
	GetChatMember(options ...MethodOption) (res GetChatMemberResponse, err error)
	SetChatStickerSet(options ...MethodOption) (res SetChatStickerSetResponse, err error)
	DeleteChatStickerSet(options ...MethodOption) (res DeleteChatStickerSetResponse, err error)
	AnswerCallbackQuery(options ...MethodOption) (res AnswerCallbackQueryResponse, err error)
	SetMyCommands(options ...MethodOption) (res SetMyCommandsResponse, err error)
	GetMyCommands() (res GetMyCommandsResponse, err error)

	// updating messages
	EditMessageText(options ...MethodOption) (res EditMessageTextResponse, err error)
	EditMessageCaption(options ...MethodOption) (res EditMessageCaptionResponse, err error)
	EditMessageMedia(options ...MethodOption) (res EditMessageMediaResponse, err error)
	EditMessageReplyMarkup(options ...MethodOption) (res EditMessageReplyMarkupResponse, err error)
	StopPoll(options ...MethodOption) (res StopPollResponse, err error)
	DeleteMessage(options ...MethodOption) (res DeleteMessageResponse, err error)

	// stickers
	SendSticker(options ...MethodOption) (res SendStickerResponse, err error)
	GetStickerSet(options ...MethodOption) (res GetStickerSetResponse, err error)
	UploadStickerFile(options ...MethodOption) (res UploadStickerFileResponse, err error)
	CreateNewStickerSet(options ...MethodOption) (res CreateNewStickerSetResponse, err error)
	AddStickerToSet(options ...MethodOption) (res AddStickerToSetResponse, err error)
	SetStickerPositionInSet(options ...MethodOption) (res SetStickerPositionInSetResponse, err error)
	DeleteStickerFromSet(options ...MethodOption) (res DeleteStickerFromSetResponse, err error)
	SetStickerSetThumb(options ...MethodOption) (res SetStickerSetThumbResponse, err error)

	// inline mode
	AnswerInlineQuery(options ...MethodOption) (res AnswerInlineQueryResponse, err error)

	// payments
	SendInvoice(options ...MethodOption) (res SendInvoiceResponse, err error)
	AnswerShippingQuery(options ...MethodOption) (res AnswerShippingQueryResponse, err error)
	AnswerPreCheckoutQuery(options ...MethodOption) (res AnswerPreCheckoutQueryResponse, err error)

	// telegram passport
	SetPassportDataErrors(options ...MethodOption) (res SetPassportDataErrorsResponse, err error)

	// games
	SendGame(options ...MethodOption) (res SendGameResponse, err error)
	SetGameScore(options ...MethodOption) (res SetGameScoreResponse, err error)
	GetGameHighScores(options ...MethodOption) (res GetGameHighScoresResponse, err error)
}

// BotOption type
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

// SetBaseURL method
func SetBaseURL(v string) BotOption {
	return func(b *bot) {
		b.baseURL = v
	}
}

func (b *bot) GetBaseURL() string {
	return b.baseURL
}

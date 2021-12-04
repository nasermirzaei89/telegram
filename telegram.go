package telegram

import "context"

// Bot interface
type Bot interface {
	GetToken() string

	// getting updates

	GetUpdates(ctx context.Context, options ...MethodOption) (res GetUpdatesResponse, err error)
	SetWebhook(ctx context.Context, options ...MethodOption) (res SetWebhookResponse, err error)
	DeleteWebhook(ctx context.Context) (res DeleteWebhookResponse, err error)
	GetWebhookInfo(ctx context.Context) (res GetWebhookInfoResponse, err error)

	// available methods

	GetMe(ctx context.Context) (res GetMeResponse, err error)
	LogOut(ctx context.Context) (res LogOutResponse, err error)
	Close(ctx context.Context) (res CloseResponse, err error)
	SendMessage(ctx context.Context, options ...MethodOption) (res SendMessageResponse, err error)
	ForwardMessage(ctx context.Context, options ...MethodOption) (res ForwardMessageResponse, err error)
	CopyMessage(ctx context.Context, options ...MethodOption) (CopyMessageResponse, error)
	SendPhoto(ctx context.Context, options ...MethodOption) (res SendPhotoResponse, err error)
	SendAudio(ctx context.Context, options ...MethodOption) (res SendAudioResponse, err error)
	SendDocument(ctx context.Context, options ...MethodOption) (res SendDocumentResponse, err error)
	SendVideo(ctx context.Context, options ...MethodOption) (res SendVideoResponse, err error)
	SendAnimation(ctx context.Context, options ...MethodOption) (res SendAnimationResponse, err error)
	SendVoice(ctx context.Context, options ...MethodOption) (res SendVoiceResponse, err error)
	SendVideoNote(ctx context.Context, options ...MethodOption) (res SendVideoNoteResponse, err error)
	SendMediaGroup(ctx context.Context, options ...MethodOption) (res SendMediaGroupResponse, err error)
	SendLocation(ctx context.Context, options ...MethodOption) (res SendLocationResponse, err error)
	EditMessageLiveLocation(ctx context.Context, options ...MethodOption) (res EditMessageLiveLocationResponse, err error)
	StopMessageLiveLocation(ctx context.Context, options ...MethodOption) (res StopMessageLiveLocationResponse, err error)
	SendVenue(ctx context.Context, options ...MethodOption) (res SendVenueResponse, err error)
	SendContact(ctx context.Context, options ...MethodOption) (res SendContactResponse, err error)
	SendPoll(ctx context.Context, options ...MethodOption) (res SendPollResponse, err error)
	SendDice(ctx context.Context, options ...MethodOption) (res SendDiceResponse, err error)
	SendChatAction(ctx context.Context, options ...MethodOption) (res SendChatActionResponse, err error)
	GetUserProfilePhotos(ctx context.Context, options ...MethodOption) (res GetUserProfilePhotosResponse, err error)
	GetFile(ctx context.Context, options ...MethodOption) (res GetFileResponse, err error)
	BanChatMember(ctx context.Context, options ...MethodOption) (res BanChatMemberResponse, err error)
	UnbanChatMember(ctx context.Context, options ...MethodOption) (res UnbanChatMemberResponse, err error)
	RestrictChatMember(ctx context.Context, options ...MethodOption) (res RestrictChatMemberResponse, err error)
	PromoteChatMember(ctx context.Context, options ...MethodOption) (res PromoteChatMemberResponse, err error)
	SetChatAdministratorCustomTitle(ctx context.Context, options ...MethodOption) (res SetChatAdministratorCustomTitleResponse, err error)
	ExportChatInviteLink(ctx context.Context, options ...MethodOption) (res ExportChatInviteLinkResponse, err error)
	CreateChatInviteLink(ctx context.Context, options ...MethodOption) (CreateChatInviteLinkResponse, error)
	EditChatInviteLink(ctx context.Context, options ...MethodOption) (EditChatInviteLinkResponse, error)
	RevokeChatInviteLink(ctx context.Context, options ...MethodOption) (RevokeChatInviteLinkResponse, error)
	SetChatPermissions(ctx context.Context, options ...MethodOption) (res SetChatPermissionsResponse, err error)
	SetChatPhoto(ctx context.Context, options ...MethodOption) (res SetChatPhotoResponse, err error)
	DeleteChatPhoto(ctx context.Context, options ...MethodOption) (res DeleteChatPhotoResponse, err error)
	SetChatTitle(ctx context.Context, options ...MethodOption) (res SetChatTitleResponse, err error)
	SetChatDescription(ctx context.Context, options ...MethodOption) (res SetChatDescriptionResponse, err error)
	PinChatMessage(ctx context.Context, options ...MethodOption) (res PinChatMessageResponse, err error)
	UnpinChatMessage(ctx context.Context, options ...MethodOption) (res UnpinChatMessageResponse, err error)
	UnpinAllChatMessages(ctx context.Context, options ...MethodOption) (UnpinAllChatMessagesResponse, error)
	LeaveChat(ctx context.Context, options ...MethodOption) (res LeaveChatResponse, err error)
	GetChat(ctx context.Context, options ...MethodOption) (res GetChatResponse, err error)
	GetChatAdministrators(ctx context.Context, options ...MethodOption) (res GetChatAdministratorsResponse, err error)
	GetChatMemberCount(ctx context.Context, options ...MethodOption) (res GetChatMemberCountResponse, err error)
	GetChatMember(ctx context.Context, options ...MethodOption) (res GetChatMemberResponse, err error)
	SetChatStickerSet(ctx context.Context, options ...MethodOption) (res SetChatStickerSetResponse, err error)
	DeleteChatStickerSet(ctx context.Context, options ...MethodOption) (res DeleteChatStickerSetResponse, err error)
	AnswerCallbackQuery(ctx context.Context, options ...MethodOption) (res AnswerCallbackQueryResponse, err error)
	SetMyCommands(ctx context.Context, options ...MethodOption) (res SetMyCommandsResponse, err error)
	DeleteMyCommands(ctx context.Context, options ...MethodOption) (res DeleteMyCommandsResponse, err error)
	GetMyCommands(ctx context.Context) (res GetMyCommandsResponse, err error)

	// updating messages

	EditMessageText(ctx context.Context, options ...MethodOption) (res EditMessageTextResponse, err error)
	EditMessageCaption(ctx context.Context, options ...MethodOption) (res EditMessageCaptionResponse, err error)
	EditMessageMedia(ctx context.Context, options ...MethodOption) (res EditMessageMediaResponse, err error)
	EditMessageReplyMarkup(ctx context.Context, options ...MethodOption) (res EditMessageReplyMarkupResponse, err error)
	StopPoll(ctx context.Context, options ...MethodOption) (res StopPollResponse, err error)
	DeleteMessage(ctx context.Context, options ...MethodOption) (res DeleteMessageResponse, err error)

	// stickers

	SendSticker(ctx context.Context, options ...MethodOption) (res SendStickerResponse, err error)
	GetStickerSet(ctx context.Context, options ...MethodOption) (res GetStickerSetResponse, err error)
	UploadStickerFile(ctx context.Context, options ...MethodOption) (res UploadStickerFileResponse, err error)
	CreateNewStickerSet(ctx context.Context, options ...MethodOption) (res CreateNewStickerSetResponse, err error)
	AddStickerToSet(ctx context.Context, options ...MethodOption) (res AddStickerToSetResponse, err error)
	SetStickerPositionInSet(ctx context.Context, options ...MethodOption) (res SetStickerPositionInSetResponse, err error)
	DeleteStickerFromSet(ctx context.Context, options ...MethodOption) (res DeleteStickerFromSetResponse, err error)
	SetStickerSetThumb(ctx context.Context, options ...MethodOption) (res SetStickerSetThumbResponse, err error)

	// inline mode

	AnswerInlineQuery(ctx context.Context, options ...MethodOption) (res AnswerInlineQueryResponse, err error)

	// payments

	SendInvoice(ctx context.Context, options ...MethodOption) (res SendInvoiceResponse, err error)
	AnswerShippingQuery(ctx context.Context, options ...MethodOption) (res AnswerShippingQueryResponse, err error)
	AnswerPreCheckoutQuery(ctx context.Context, options ...MethodOption) (res AnswerPreCheckoutQueryResponse, err error)

	// telegram passport

	SetPassportDataErrors(ctx context.Context, options ...MethodOption) (res SetPassportDataErrorsResponse, err error)

	// games

	SendGame(ctx context.Context, options ...MethodOption) (res SendGameResponse, err error)
	SetGameScore(ctx context.Context, options ...MethodOption) (res SetGameScoreResponse, err error)
	GetGameHighScores(ctx context.Context, options ...MethodOption) (res GetGameHighScoresResponse, err error)
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

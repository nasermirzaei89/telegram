package telegram

type Bot interface {
	// getting updates
	GetUpdates(...GetUpdatesOption) (*GetUpdatesResponse, error)
	SetWebhook(...SetWebhookOption) (*SetWebhookResponse, error)
	DeleteWebhook() (*DeleteWebhookResponse, error)
	GetWebhookInfo() (*GetWebhookInfoResponse, error)

	// available methods
	GetMe() (*GetMeResponse, error)
	SendMessage() SendMessageRequest
	ForwardMessage() ForwardMessageRequest
	SendPhoto() SendPhotoRequest
	SendAudio() SendAudioRequest
	SendDocument() SendDocumentRequest
	SendVideo() SendVideoRequest
	SendAnimation() SendAnimationRequest
	SendVoice() SendVoiceRequest
	SendVideoNote() SendVideoNoteRequest
	SendMediaGroup() SendMediaGroupRequest
	SendLocation() SendLocationRequest
	EditMessageLiveLocation() EditMessageLiveLocationRequest
	StopMessageLiveLocation() StopMessageLiveLocationRequest
	SendVenue() SendVenueRequest
	SendContact() SendContactRequest
	SendPoll() SendPollRequest
	SendChatAction() SendChatActionRequest
	GetUserProfilePhotos() GetUserProfilePhotosRequest
	GetFile() GetFileRequest
	KickChatMember() KickChatMemberRequest
	UnbanChatMember() UnbanChatMemberRequest
	RestrictChatMember() RestrictChatMemberRequest
	PromoteChatMember() PromoteChatMemberRequest
	ExportChatInviteLink() ExportChatInviteLinkRequest
	SetChatPhoto() SetChatPhotoRequest
	DeleteChatPhoto() DeleteChatPhotoRequest
	SetChatTitle() SetChatTitleRequest
	SetChatDescription() SetChatDescriptionRequest
	PinChatMessage() PinChatMessageRequest
	UnpinChatMessage() UnpinChatMessageRequest
	LeaveChat() LeaveChatRequest
	GetChat() GetChatRequest
	GetChatAdministrators() GetChatAdministratorsRequest
	GetChatMembersCount() GetChatMembersCountRequest
	GetChatMember() GetChatMemberRequest
	SetChatStickerSet() SetChatStickerSetRequest
	DeleteChatStickerSet() DeleteChatStickerSetRequest
	AnswerCallbackQuery() AnswerCallbackQueryRequest

	// updating messages
	EditMessageText() EditMessageTextRequest
	EditMessageCaption() EditMessageCaptionRequest
	EditMessageMedia() EditMessageMediaRequest
	EditMessageReplyMarkup() EditMessageReplyMarkupRequest
	StopPoll() StopPollRequest
	DeleteMessage() DeleteMessageRequest

	// stickers
	SendSticker() SendStickerRequest
	GetStickerSet() GetStickerSetRequest
	UploadStickerFile() UploadStickerFileRequest
	CreateNewStickerSet() CreateNewStickerSetRequest
	AddStickerToSet() AddStickerToSetRequest
	SetStickerPositionInSet() SetStickerPositionInSetRequest
	DeleteStickerFromSet() DeleteStickerFromSetRequest

	// inline mode
	AnswerInlineQuery() AnswerInlineQueryRequest

	// payments
	SendInvoice() SendInvoiceRequest
	AnswerShippingQuery() AnswerShippingQueryRequest
	AnswerPreCheckoutQuery() AnswerPreCheckoutQueryRequest

	// telegram passport
	SetPassportDataErrors() SetPassportDataErrorsRequest

	// games
	SendGame() SendGameRequest
	SetGameScore() SetGameScoreRequest
	GetGameHighScores() GetGameHighScoresRequest
}

type bot struct {
	Token string
}

func New(token string) Bot {
	return &bot{Token: token}
}

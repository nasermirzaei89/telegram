package telegram

type GetMeRequest interface {
	Do() (*User, error)
}

func (b *bot) GetMe() GetMeRequest {
	panic("implement me")
}

type SendMessageRequest interface {
	ChatID(int) SendMessageRequest
	ChatUsername(string) SendMessageRequest
	Text(string) SendMessageRequest
	ParseMode(string) SendMessageRequest
	DisableWebPagePreview() SendMessageRequest
	DisableNotification() SendMessageRequest
	ReplyToMessageID(int) SendMessageRequest
	ReplyMarkup(interface{}) SendMessageRequest
	Do() (*Message, error)
}

func (b *bot) SendMessage() SendMessageRequest {
	panic("implement me")
}

type ForwardMessageRequest interface {
	ChatID(int) ForwardMessageRequest
	ChatUsername(string) ForwardMessageRequest
	FromChatID(int) ForwardMessageRequest
	FromChatUsername(string) ForwardMessageRequest
	DisableNotification() ForwardMessageRequest
	MessageID(int) ForwardMessageRequest
	Do() (*Message, error)
}

func (b *bot) ForwardMessage() ForwardMessageRequest {
	panic("implement me")
}

type SendPhotoRequest interface {
	ChatID(int) SendPhotoRequest
	ChatUsername(string) SendPhotoRequest
	Photo(InputFile) SendPhotoRequest
	PhotoFileID(string) SendPhotoRequest
	PhotoURL(string) SendPhotoRequest
	Caption(string) SendPhotoRequest
	ParseMode(string) SendPhotoRequest
	DisableNotification() SendPhotoRequest
	ReplyToMessageID(int) SendPhotoRequest
	ReplyMarkup(interface{}) SendPhotoRequest
	Do() (*Message, error)
}

func (b *bot) SendPhoto() SendPhotoRequest {
	panic("implement me")
}

type SendAudioRequest interface {
	ChatID(int) SendAudioRequest
	ChatUsername(string) SendAudioRequest
	Audio(InputFile) SendAudioRequest
	AudioFileID(string) SendAudioRequest
	AudioURL(string) SendAudioRequest
	Caption(string) SendAudioRequest
	ParseMode(string) SendAudioRequest
	Duration(int) SendAudioRequest
	Performer(string) SendAudioRequest
	Title(string) SendAudioRequest
	Thumb(InputFile) SendAudioRequest
	ThumbFileName(string) SendAudioRequest
	DisableNotification() SendAudioRequest
	ReplyToMessageID(int) SendAudioRequest
	ReplyMarkup(interface{}) SendAudioRequest
	Do() (*Message, error)
}

func (b *bot) SendAudio() SendAudioRequest {
	panic("implement me")
}

type SendDocumentRequest interface {
	ChatID(int) SendDocumentRequest
	ChatUsername(string) SendDocumentRequest
	Document(InputFile) SendDocumentRequest
	DocumentFileID(string) SendDocumentRequest
	DocumentURL(string) SendDocumentRequest
	Thumb(InputFile) SendDocumentRequest
	ThumbFileName(string) SendDocumentRequest
	Caption(string) SendDocumentRequest
	ParseMode(string) SendDocumentRequest
	DisableNotification() SendDocumentRequest
	ReplyToMessageID(int) SendDocumentRequest
	ReplyMarkup(interface{}) SendDocumentRequest
	Do() (*Message, error)
}

func (b *bot) SendDocument() SendDocumentRequest {
	panic("implement me")
}

type SendVideoRequest interface {
	ChatID(int) SendVideoRequest
	ChatUsername(string) SendVideoRequest
	Video(InputFile) SendVideoRequest
	VideoFileID(string) SendVideoRequest
	VideoURL(string) SendVideoRequest
	Duration(int) SendVideoRequest
	Width(int) SendVideoRequest
	Height(int) SendVideoRequest
	Thumb(InputFile) SendVideoRequest
	ThumbFileName(string) SendVideoRequest
	Caption(string) SendVideoRequest
	ParseMode(string) SendVideoRequest
	SupportsStreaming() SendVideoRequest
	DisableNotification() SendVideoRequest
	ReplyToMessageID(int) SendVideoRequest
	ReplyMarkup(interface{}) SendVideoRequest
	Do() (*Message, error)
}

func (b *bot) SendVideo() SendVideoRequest {
	panic("implement me")
}

type SendAnimationRequest interface {
	ChatID(int) SendAnimationRequest
	ChatUsername(string) SendAnimationRequest
	Animation(InputFile) SendAnimationRequest
	AnimationFileID(string) SendAnimationRequest
	AnimationURL(string) SendAnimationRequest
	Duration(int) SendAnimationRequest
	Width(int) SendAnimationRequest
	Height(int) SendAnimationRequest
	Thumb(InputFile) SendAnimationRequest
	ThumbFileName(string) SendAnimationRequest
	Caption(string) SendAnimationRequest
	ParseMode(string) SendAnimationRequest
	DisableNotification() SendAnimationRequest
	ReplyToMessageID(int) SendAnimationRequest
	ReplyMarkup(interface{}) SendAnimationRequest
	Do() (*Message, error)
}

func (b *bot) SendAnimation() SendAnimationRequest {
	panic("implement me")
}

type SendVoiceRequest interface {
	ChatID(int) SendVoiceRequest
	ChatUsername(string) SendVoiceRequest
	Voice(InputFile) SendVoiceRequest
	VoiceFileID(string) SendVoiceRequest
	VoiceURL(string) SendVoiceRequest
	Caption(string) SendVoiceRequest
	ParseMode(string) SendVoiceRequest
	Duration(int) SendVoiceRequest
	DisableNotification() SendVoiceRequest
	ReplyToMessageID(int) SendVoiceRequest
	ReplyMarkup(interface{}) SendVoiceRequest
	Do() (*Message, error)
}

func (b *bot) SendVoice() SendVoiceRequest {
	panic("implement me")
}

type SendVideoNoteRequest interface {
	ChatID(int) SendVideoNoteRequest
	ChatUsername(string) SendVideoNoteRequest
	VideoNote(InputFile) SendVideoNoteRequest
	VideoNoteFileID(string) SendVideoNoteRequest
	VideoNoteURL(string) SendVideoNoteRequest
	Duration(int) SendVideoNoteRequest
	Length(int) SendVideoNoteRequest
	Thumb(InputFile) SendVideoNoteRequest
	ThumbFileName(string) SendVideoNoteRequest
	DisableNotification() SendVideoNoteRequest
	ReplyToMessageID(int) SendVideoNoteRequest
	ReplyMarkup(interface{}) SendVideoNoteRequest
	Do() (*Message, error)
}

func (b *bot) SendVideoNote() SendVideoNoteRequest {
	panic("implement me")
}

type SendMediaGroupRequest interface {
	ChatID(int) SendMediaGroupRequest
	ChatUsername(string) SendMediaGroupRequest
	Media([]InputMedia) SendMediaGroupRequest
	DisableNotification() SendMediaGroupRequest
	ReplyToMessageID(int) SendMediaGroupRequest
	Do() (*Message, error)
}

func (b *bot) SendMediaGroup() SendMediaGroupRequest {
	panic("implement me")
}

type SendLocationRequest interface {
	ChatID(int) SendLocationRequest
	ChatUsername(string) SendLocationRequest
	Latitude(float32) SendLocationRequest
	Longitude(float32) SendLocationRequest
	LivePeriod(int) SendLocationRequest
	DisableNotification() SendLocationRequest
	ReplyToMessageID(int) SendLocationRequest
	ReplyMarkup(interface{}) SendLocationRequest
	Do() (*Message, error)
}

func (b *bot) SendLocation() SendLocationRequest {
	panic("implement me")
}

type EditMessageLiveLocationRequest interface {
	ChatID(int) EditMessageLiveLocationRequest
	ChatUsername(string) EditMessageLiveLocationRequest
	MessageID(int) EditMessageLiveLocationRequest
	InlineMessageID(string) EditMessageLiveLocationRequest
	Latitude(float32) EditMessageLiveLocationRequest
	Longitude(float32) EditMessageLiveLocationRequest
	ReplyMarkup(interface{}) EditMessageLiveLocationRequest
	Do() (*Message, error)
}

func (b *bot) EditMessageLiveLocation() EditMessageLiveLocationRequest {
	panic("implement me")
}

type StopMessageLiveLocationRequest interface {
	ChatID(int) StopMessageLiveLocationRequest
	ChatUsername(string) StopMessageLiveLocationRequest
	MessageID(int) StopMessageLiveLocationRequest
	InlineMessageID(string) StopMessageLiveLocationRequest
	ReplyMarkup(interface{}) StopMessageLiveLocationRequest
	Do() (*Message, error)
}

func (b *bot) StopMessageLiveLocation() StopMessageLiveLocationRequest {
	panic("implement me")
}

type SendVenueRequest interface {
	ChatID(int) SendVenueRequest
	ChatUsername(string) SendVenueRequest
	Latitude(float32) SendVenueRequest
	Longitude(float32) SendVenueRequest
	Title(string) SendVenueRequest
	Address(string) SendVenueRequest
	FoursquareID(string) SendVenueRequest
	FoursquareType(string) SendVenueRequest
	DisableNotification() SendVenueRequest
	ReplyToMessageID(int) SendVenueRequest
	ReplyMarkup(interface{}) SendVenueRequest
	Do() (*Message, error)
}

func (b *bot) SendVenue() SendVenueRequest {
	panic("implement me")
}

type SendContactRequest interface {
	ChatID(int) SendContactRequest
	ChatUsername(string) SendContactRequest
	PhoneNumber(string) SendContactRequest
	FirstName(string) SendContactRequest
	LastName(string) SendContactRequest
	VCard(string) SendContactRequest
	DisableNotification() SendContactRequest
	ReplyToMessageID(int) SendContactRequest
	ReplyMarkup(interface{}) SendContactRequest
	Do() (*Message, error)
}

func (b *bot) SendContact() SendContactRequest {
	panic("implement me")
}

type SendPollRequest interface {
	ChatID(int) SendPollRequest
	ChatUsername(string) SendPollRequest
	Question(string) SendPollRequest
	Options(...string) SendPollRequest
	DisableNotification() SendPollRequest
	ReplyToMessageID(int) SendPollRequest
	ReplyMarkup(interface{}) SendPollRequest
	Do() (*Message, error)
}

func (b *bot) SendPoll() SendPollRequest {
	panic("implement me")
}

type SendChatActionRequest interface {
	ChatID(int) SendChatActionRequest
	ChatUsername(string) SendChatActionRequest
	Action(string) SendChatActionRequest
	Do() (*Message, error)
}

func (b *bot) SendChatAction() SendChatActionRequest {
	panic("implement me")
}

type GetUserProfilePhotosRequest interface {
	UserID(int) GetUserProfilePhotosRequest
	Offset(int) GetUserProfilePhotosRequest
	Limit(int) GetUserProfilePhotosRequest
	Do() (UserProfilePhotos, error)
}

func (b *bot) GetUserProfilePhotos() GetUserProfilePhotosRequest {
	panic("implement me")
}

type GetFileRequest interface {
	FileID(string) GetFileRequest
	Do() (File, error)
}

func (b *bot) GetFile() GetFileRequest {
	panic("implement me")
}

type KickChatMemberRequest interface {
	ChatID(int) KickChatMemberRequest
	ChatUsername(string) KickChatMemberRequest
	UserID(int) KickChatMemberRequest
	UntilDate(int) KickChatMemberRequest
	Do() (bool, error)
}

func (b *bot) KickChatMember() KickChatMemberRequest {
	panic("implement me")
}

type UnbanChatMemberRequest interface {
	ChatID(int) UnbanChatMemberRequest
	ChatUsername(string) UnbanChatMemberRequest
	UserID(int) UnbanChatMemberRequest
	Do() (bool, error)
}

func (b *bot) UnbanChatMember() UnbanChatMemberRequest {
	panic("implement me")
}

type RestrictChatMemberRequest interface {
	ChatID(int) RestrictChatMemberRequest
	ChatUsername(string) RestrictChatMemberRequest
	UserID(int) RestrictChatMemberRequest
	UntilDate(int) RestrictChatMemberRequest
	CanSendMessages() RestrictChatMemberRequest
	CanSendMediaMessages() RestrictChatMemberRequest
	CanSendOtherMessages() RestrictChatMemberRequest
	CanAddWebPagePreviews() RestrictChatMemberRequest
	Do() (bool, error)
}

func (b *bot) RestrictChatMember() RestrictChatMemberRequest {
	panic("implement me")
}

type PromoteChatMemberRequest interface {
	ChatID(int) PromoteChatMemberRequest
	ChatUsername(string) PromoteChatMemberRequest
	UserID(int) PromoteChatMemberRequest
	CanChangeInfo() PromoteChatMemberRequest
	CanSendMessages() PromoteChatMemberRequest
	CanEditMessages() PromoteChatMemberRequest
	CanDeleteMessages() PromoteChatMemberRequest
	CanInviteUsers() PromoteChatMemberRequest
	CanRestrictMembers() PromoteChatMemberRequest
	CanPinMessages() PromoteChatMemberRequest
	CanPromoteMembers() PromoteChatMemberRequest
	Do() (bool, error)
}

func (b *bot) PromoteChatMember() PromoteChatMemberRequest {
	panic("implement me")
}

type ExportChatInviteLinkRequest interface {
	ChatID(int) ExportChatInviteLinkRequest
	ChatUsername(string) ExportChatInviteLinkRequest
	Do() (bool, error)
}

func (b *bot) ExportChatInviteLink() ExportChatInviteLinkRequest {
	panic("implement me")
}

type SetChatPhotoRequest interface {
	ChatID(int) SetChatPhotoRequest
	ChatUsername(string) SetChatPhotoRequest
	Photo(InputFile) SetChatPhotoRequest
	Do() (bool, error)
}

func (b *bot) SetChatPhoto() SetChatPhotoRequest {
	panic("implement me")
}

type DeleteChatPhotoRequest interface {
	ChatID(int) DeleteChatPhotoRequest
	ChatUsername(string) DeleteChatPhotoRequest
	Do() (bool, error)
}

func (b *bot) DeleteChatPhoto() DeleteChatPhotoRequest {
	panic("implement me")
}

type SetChatTitleRequest interface {
	ChatID(int) SetChatTitleRequest
	ChatUsername(string) SetChatTitleRequest
	Title(string) SetChatTitleRequest
	Do() (bool, error)
}

func (b *bot) SetChatTitle() SetChatTitleRequest {
	panic("implement me")
}

type SetChatDescriptionRequest interface {
	ChatID(int) SetChatDescriptionRequest
	ChatUsername(string) SetChatDescriptionRequest
	Description(string) SetChatDescriptionRequest
	Do() (bool, error)
}

func (b *bot) SetChatDescription() SetChatDescriptionRequest {
	panic("implement me")
}

type PinChatMessageRequest interface {
	ChatID(int) PinChatMessageRequest
	ChatUsername(string) PinChatMessageRequest
	MessageID(int) PinChatMessageRequest
	DisableNotification() PinChatMessageRequest
	Do() (bool, error)
}

func (b *bot) PinChatMessage() PinChatMessageRequest {
	panic("implement me")
}

type UnpinChatMessageRequest interface {
	ChatID(int) UnpinChatMessageRequest
	ChatUsername(string) UnpinChatMessageRequest
	Do() (bool, error)
}

func (b *bot) UnpinChatMessage() UnpinChatMessageRequest {
	panic("implement me")
}

type LeaveChatRequest interface {
	ChatID(int) LeaveChatRequest
	ChatUsername(string) LeaveChatRequest
	Do() (bool, error)
}

func (b *bot) LeaveChat() LeaveChatRequest {
	panic("implement me")
}

type GetChatRequest interface {
	ChatID(int) GetChatRequest
	ChatUsername(string) GetChatRequest
	Do() (*Chat, error)
}

func (b *bot) GetChat() GetChatRequest {
	panic("implement me")
}

type GetChatAdministratorsRequest interface {
	ChatID(int) GetChatAdministratorsRequest
	ChatUsername(string) GetChatAdministratorsRequest
	Do() ([]ChatMember, error)
}

func (b *bot) GetChatAdministrators() GetChatAdministratorsRequest {
	panic("implement me")
}

type GetChatMembersCountRequest interface {
	ChatID(int) GetChatMembersCountRequest
	ChatUsername(string) GetChatMembersCountRequest
	Do() (int, error)
}

func (b *bot) GetChatMembersCount() GetChatMembersCountRequest {
	panic("implement me")
}

type GetChatMemberRequest interface {
	ChatID(int) GetChatMemberRequest
	ChatUsername(string) GetChatMemberRequest
	UserID(int) GetChatMemberRequest
	Do() (*ChatMember, error)
}

func (b *bot) GetChatMember() GetChatMemberRequest {
	panic("implement me")
}

type SetChatStickerSetRequest interface {
	ChatID(int) SetChatStickerSetRequest
	ChatUsername(string) SetChatStickerSetRequest
	StickerSetName(string) SetChatStickerSetRequest
	Do() (bool, error)
}

func (b *bot) SetChatStickerSet() SetChatStickerSetRequest {
	panic("implement me")
}

type DeleteChatStickerSetRequest interface {
	ChatID(int) DeleteChatStickerSetRequest
	ChatUsername(string) DeleteChatStickerSetRequest
	Do() (bool, error)
}

func (b *bot) DeleteChatStickerSet() DeleteChatStickerSetRequest {
	panic("implement me")
}

type AnswerCallbackQueryRequest interface {
	CallbackQueryID(string) AnswerCallbackQueryRequest
	Text(string) AnswerCallbackQueryRequest
	ShowAlert() AnswerCallbackQueryRequest
	URL(string) AnswerCallbackQueryRequest
	CacheTime(int) AnswerCallbackQueryRequest
	Do() (bool, error)
}

func (b *bot) AnswerCallbackQuery() AnswerCallbackQueryRequest {
	panic("implement me")
}

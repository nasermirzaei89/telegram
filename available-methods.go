package telegram

type GetMeResponse interface {
	Response
	GetUser() *User
}

type getMeResponse struct {
	response
	Result *User `json:"result,omitempty"`
}

func (r *getMeResponse) GetUser() *User {
	return r.Result
}

func (b *bot) GetMe() (GetMeResponse, error) {
	var res getMeResponse
	err := doRequest(b.Token, "getMe", &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type SendMessageResponse interface {
}

func (b *bot) SendMessage(options ...Option) (SendMessageResponse, error) {
	panic("implement me")
}

type ForwardMessageResponse interface {
}

func (b *bot) ForwardMessage(options ...Option) (ForwardMessageResponse, error) {
	panic("implement me")
}

type SendPhotoResponse interface {
}

func (b *bot) SendPhoto(options ...Option) (SendPhotoResponse , error){
	panic("implement me")
}

type SendAudioResponse interface {
}

func (b *bot) SendAudio(options ...Option) (SendAudioResponse, error) {
	panic("implement me")
}

type SendDocumentResponse interface {
}

func (b *bot) SendDocument(options ...Option) (SendDocumentResponse, error) {
	panic("implement me")
}

type SendVideoResponse interface {
}

func (b *bot) SendVideo(options ...Option) (SendVideoResponse, error) {
	panic("implement me")
}

type SendAnimationResponse interface {
}

func (b *bot) SendAnimation(options ...Option) (SendAnimationResponse, error) {
	panic("implement me")
}

type SendVoiceResponse interface {
}

func (b *bot) SendVoice(options ...Option) (SendVoiceResponse, error) {
	panic("implement me")
}

type SendVideoNoteResponse interface {
}

func (b *bot) SendVideoNote(options ...Option) (SendVideoNoteResponse, error) {
	panic("implement me")
}

type SendMediaGroupResponse interface {
}

func (b *bot) SendMediaGroup(options ...Option) (SendMediaGroupResponse, error) {
	panic("implement me")
}

type SendLocationResponse interface {
}

func (b *bot) SendLocation(options ...Option) (SendLocationResponse, error) {
	panic("implement me")
}

type EditMessageLiveLocationResponse interface {
}

func (b *bot) EditMessageLiveLocation(options ...Option) (EditMessageLiveLocationResponse, error) {
	panic("implement me")
}

type StopMessageLiveLocationResponse interface {
}

func (b *bot) StopMessageLiveLocation(options ...Option) (StopMessageLiveLocationResponse, error) {
	panic("implement me")
}

type SendVenueResponse interface {
}

func (b *bot) SendVenue(options ...Option) (SendVenueResponse, error) {
	panic("implement me")
}

type SendContactResponse interface {
}

func (b *bot) SendContact(options ...Option) (SendContactResponse, error) {
	panic("implement me")
}

type SendPollResponse interface {
}

func (b *bot) SendPoll(options ...Option) (SendPollResponse, error) {
	panic("implement me")
}

type SendChatActionResponse interface {
}

func (b *bot) SendChatAction(options ...Option) (SendChatActionResponse, error) {
	panic("implement me")
}

type GetUserProfilePhotosResponse interface {
}

func (b *bot) GetUserProfilePhotos(options ...Option) (GetUserProfilePhotosResponse, error) {
	panic("implement me")
}

type GetFileResponse interface {
}

func (b *bot) GetFile(options ...Option) (GetFileResponse, error) {
	panic("implement me")
}

type KickChatMemberResponse interface {
}

func (b *bot) KickChatMember(options ...Option) (KickChatMemberResponse, error) {
	panic("implement me")
}

type UnbanChatMemberResponse interface {
}

func (b *bot) UnbanChatMember(options ...Option) (UnbanChatMemberResponse, error) {
	panic("implement me")
}

type RestrictChatMemberResponse interface {
}

func (b *bot) RestrictChatMember(options ...Option) (RestrictChatMemberResponse, error) {
	panic("implement me")
}

type PromoteChatMemberResponse interface {
}

func (b *bot) PromoteChatMember(options ...Option) (PromoteChatMemberResponse, error) {
	panic("implement me")
}

type ExportChatInviteLinkResponse interface {
}

func (b *bot) ExportChatInviteLink(options ...Option) (ExportChatInviteLinkResponse, error) {
	panic("implement me")
}

type SetChatPhotoResponse interface {
}

func (b *bot) SetChatPhoto(options ...Option) (SetChatPhotoResponse, error) {
	panic("implement me")
}

type DeleteChatPhotoResponse interface {
}

func (b *bot) DeleteChatPhoto(options ...Option) (DeleteChatPhotoResponse, error) {
	panic("implement me")
}

type SetChatTitleResponse interface {
}

func (b *bot) SetChatTitle(options ...Option) (SetChatTitleResponse, error) {
	panic("implement me")
}

type SetChatDescriptionResponse interface {
}

func (b *bot) SetChatDescription(options ...Option) (SetChatDescriptionResponse, error) {
	panic("implement me")
}

type PinChatMessageResponse interface {
}

func (b *bot) PinChatMessage(options ...Option) (PinChatMessageResponse, error) {
	panic("implement me")
}

type UnpinChatMessageResponse interface {
}

func (b *bot) UnpinChatMessage(options ...Option) (UnpinChatMessageResponse, error) {
	panic("implement me")
}

type LeaveChatResponse interface {
}

func (b *bot) LeaveChat(options ...Option) (LeaveChatResponse, error) {
	panic("implement me")
}

type GetChatResponse interface {
}

func (b *bot) GetChat(options ...Option) (GetChatResponse, error) {
	panic("implement me")
}

type GetChatAdministratorsResponse interface {
}

func (b *bot) GetChatAdministrators(options ...Option) (GetChatAdministratorsResponse, error) {
	panic("implement me")
}

type GetChatMembersCountResponse interface {
}

func (b *bot) GetChatMembersCount(options ...Option) (GetChatMembersCountResponse, error) {
	panic("implement me")
}

type GetChatMemberResponse interface {
}

func (b *bot) GetChatMember(options ...Option) (GetChatMemberResponse, error) {
	panic("implement me")
}

type SetChatStickerSetResponse interface {
}

func (b *bot) SetChatStickerSet(options ...Option) (SetChatStickerSetResponse, error) {
	panic("implement me")
}

type DeleteChatStickerSetResponse interface {
}

func (b *bot) DeleteChatStickerSet(options ...Option) (DeleteChatStickerSetResponse, error) {
	panic("implement me")
}

type AnswerCallbackQueryResponse interface {
}

func (b *bot) AnswerCallbackQuery(options ...Option) (AnswerCallbackQueryResponse, error) {
	panic("implement me")
}

package telegram

import "context"

// GetMeResponse interface.
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

func (b *bot) GetMe(ctx context.Context) (GetMeResponse, error) {
	var res getMeResponse

	err := b.doRequest(ctx, "getMe", &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// LogOutResponse interface.
type LogOutResponse interface {
	Response
}

type logOutResponse struct {
	response
}

func (b *bot) LogOut(ctx context.Context) (LogOutResponse, error) {
	var res logOutResponse

	err := b.doRequest(ctx, "logOut", &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CloseResponse interface.
type CloseResponse interface {
	Response
}

type closeResponse struct {
	response
}

func (b *bot) Close(ctx context.Context) (CloseResponse, error) {
	var res closeResponse

	err := b.doRequest(ctx, "close", &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendMessageResponse interface.
type SendMessageResponse interface {
	Response
	GetMessage() *Message
}

type sendMessageResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendMessageResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendMessage(ctx context.Context, options ...MethodOption) (SendMessageResponse, error) {
	var res sendMessageResponse

	err := b.doRequest(ctx, "sendMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ForwardMessageResponse interface.
type ForwardMessageResponse interface {
	Response
	GetMessage() *Message
}

type forwardMessageResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *forwardMessageResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) ForwardMessage(ctx context.Context, options ...MethodOption) (ForwardMessageResponse, error) {
	var res forwardMessageResponse

	err := b.doRequest(ctx, "forwardMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CopyMessageResponse interface.
type CopyMessageResponse interface {
	Response
	GetMessageID() *MessageID
}

type copyMessageResponse struct {
	response
	Result *MessageID `json:"result,omitempty"`
}

func (r *copyMessageResponse) GetMessageID() *MessageID {
	return r.Result
}

func (b *bot) CopyMessage(ctx context.Context, options ...MethodOption) (CopyMessageResponse, error) {
	var res copyMessageResponse

	err := b.doRequest(ctx, "copyMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendPhotoResponse interface
type SendPhotoResponse interface {
	Response
	GetMessage() *Message
}

type sendPhotoResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendPhotoResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendPhoto(ctx context.Context, options ...MethodOption) (SendPhotoResponse, error) {
	var res sendPhotoResponse

	err := b.doRequest(ctx, "sendPhoto", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendAudioResponse interface
type SendAudioResponse interface {
	Response
	GetMessage() *Message
}

type sendAudioResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendAudioResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendAudio(ctx context.Context, options ...MethodOption) (SendAudioResponse, error) {
	var res sendAudioResponse

	err := b.doRequest(ctx, "sendAudio", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendDocumentResponse interface
type SendDocumentResponse interface {
	Response
	GetMessage() *Message
}

type sendDocumentResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendDocumentResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendDocument(ctx context.Context, options ...MethodOption) (SendDocumentResponse, error) {
	var res sendDocumentResponse

	err := b.doRequest(ctx, "sendDocument", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendVideoResponse interface
type SendVideoResponse interface {
	Response
	GetMessage() *Message
}

type sendVideoResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendVideoResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendVideo(ctx context.Context, options ...MethodOption) (SendVideoResponse, error) {
	var res sendVideoResponse

	err := b.doRequest(ctx, "sendVideo", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendAnimationResponse interface
type SendAnimationResponse interface {
	Response
	GetMessage() *Message
}

type sendAnimationResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendAnimationResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendAnimation(ctx context.Context, options ...MethodOption) (SendAnimationResponse, error) {
	var res sendAnimationResponse

	err := b.doRequest(ctx, "sendAnimation", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendVoiceResponse interface
type SendVoiceResponse interface {
	Response
	GetMessage() *Message
}

type sendVoiceResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendVoiceResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendVoice(ctx context.Context, options ...MethodOption) (SendVoiceResponse, error) {
	var res sendVoiceResponse

	err := b.doRequest(ctx, "sendVoice", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendVideoNoteResponse interface
type SendVideoNoteResponse interface {
	Response
	GetMessage() *Message
}

type sendVideoNoteResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendVideoNoteResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendVideoNote(ctx context.Context, options ...MethodOption) (SendVideoNoteResponse, error) {
	var res sendVideoNoteResponse

	err := b.doRequest(ctx, "sendVideoNote", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendMediaGroupResponse interface
type SendMediaGroupResponse interface {
	Response
	GetMessage() *Message
}

type sendMediaGroupResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendMediaGroupResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendMediaGroup(ctx context.Context, options ...MethodOption) (SendMediaGroupResponse, error) {
	var res sendMediaGroupResponse

	err := b.doRequest(ctx, "sendMediaGroup", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendLocationResponse interface
type SendLocationResponse interface {
	Response
	GetMessage() *Message
}

type sendLocationResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendLocationResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendLocation(ctx context.Context, options ...MethodOption) (SendLocationResponse, error) {
	var res sendLocationResponse

	err := b.doRequest(ctx, "sendLocation", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// EditMessageLiveLocationResponse interface
type EditMessageLiveLocationResponse interface {
	Response
	GetEditedMessage() *Message
}

type editMessageLiveLocationResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *editMessageLiveLocationResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) EditMessageLiveLocation(ctx context.Context, options ...MethodOption) (EditMessageLiveLocationResponse, error) {
	var res editMessageLiveLocationResponse

	err := b.doRequest(ctx, "editMessageLiveLocation", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// StopMessageLiveLocationResponse interface
type StopMessageLiveLocationResponse interface {
	Response
	GetMessage() *Message
}

type stopMessageLiveLocationResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *stopMessageLiveLocationResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) StopMessageLiveLocation(ctx context.Context, options ...MethodOption) (StopMessageLiveLocationResponse, error) {
	var res stopMessageLiveLocationResponse

	err := b.doRequest(ctx, "stopMessageLiveLocation", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendVenueResponse interface
type SendVenueResponse interface {
	Response
	GetMessage() *Message
}

type sendVenueResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendVenueResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendVenue(ctx context.Context, options ...MethodOption) (SendVenueResponse, error) {
	var res sendVenueResponse

	err := b.doRequest(ctx, "sendVenue", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendContactResponse interface
type SendContactResponse interface {
	Response
	GetMessage() *Message
}

type sendContactResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendContactResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendContact(ctx context.Context, options ...MethodOption) (SendContactResponse, error) {
	var res sendContactResponse

	err := b.doRequest(ctx, "sendContact", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendPollResponse interface
type SendPollResponse interface {
	Response
	GetMessage() *Message
}

type sendPollResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendPollResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendPoll(ctx context.Context, options ...MethodOption) (SendPollResponse, error) {
	var res sendPollResponse

	err := b.doRequest(ctx, "sendPoll", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendDiceResponse interface
type SendDiceResponse interface {
	Response
}

type sendDiceResponse struct {
	response
}

func (b *bot) SendDice(ctx context.Context, options ...MethodOption) (SendDiceResponse, error) {
	var res sendDiceResponse

	err := b.doRequest(ctx, "sendDice", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendChatActionResponse interface
type SendChatActionResponse interface {
	Response
}

type sendChatActionResponse struct {
	response
}

func (b *bot) SendChatAction(ctx context.Context, options ...MethodOption) (SendChatActionResponse, error) {
	var res sendChatActionResponse

	err := b.doRequest(ctx, "sendChatAction", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetUserProfilePhotosResponse interface
type GetUserProfilePhotosResponse interface {
	Response
	GetUserProfilePhotos() *UserProfilePhotos
}

type getUserProfilePhotosResponse struct {
	response
	Result *UserProfilePhotos `json:"result,omitempty"`
}

func (r *getUserProfilePhotosResponse) GetUserProfilePhotos() *UserProfilePhotos {
	return r.Result
}

func (b *bot) GetUserProfilePhotos(ctx context.Context, options ...MethodOption) (GetUserProfilePhotosResponse, error) {
	var res getUserProfilePhotosResponse

	err := b.doRequest(ctx, "getUserProfilePhotos", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetFileResponse interface
type GetFileResponse interface {
	Response
	GetFile() *File
}

type getFileResponse struct {
	response
	Result *File `json:"result,omitempty"`
}

func (r *getFileResponse) GetFile() *File {
	return r.Result
}

func (b *bot) GetFile(ctx context.Context, options ...MethodOption) (GetFileResponse, error) {
	var res getFileResponse

	err := b.doRequest(ctx, "getFile", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// BanChatMemberResponse interface
type BanChatMemberResponse interface {
	Response
}

type banChatMemberResponse struct {
	response
}

func (b *bot) BanChatMember(ctx context.Context, options ...MethodOption) (BanChatMemberResponse, error) {
	var res banChatMemberResponse

	err := b.doRequest(ctx, "banChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UnbanChatMemberResponse interface
type UnbanChatMemberResponse interface {
	Response
}

type unbanChatMemberResponse struct {
	response
}

func (b *bot) UnbanChatMember(ctx context.Context, options ...MethodOption) (UnbanChatMemberResponse, error) {
	var res unbanChatMemberResponse

	err := b.doRequest(ctx, "unbanChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// RestrictChatMemberResponse interface
type RestrictChatMemberResponse interface {
	Response
}

type restrictChatMemberResponse struct {
	response
}

func (b *bot) RestrictChatMember(ctx context.Context, options ...MethodOption) (RestrictChatMemberResponse, error) {
	var res restrictChatMemberResponse

	err := b.doRequest(ctx, "restrictChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// PromoteChatMemberResponse interface
type PromoteChatMemberResponse interface {
	Response
}

type promoteChatMemberResponse struct {
	response
}

func (b *bot) PromoteChatMember(ctx context.Context, options ...MethodOption) (PromoteChatMemberResponse, error) {
	var res promoteChatMemberResponse

	err := b.doRequest(ctx, "promoteChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetChatAdministratorCustomTitleResponse interface
type SetChatAdministratorCustomTitleResponse interface {
	Response
}

type setChatAdministratorCustomTitleResponse struct {
	response
}

func (b *bot) SetChatAdministratorCustomTitle(ctx context.Context, options ...MethodOption) (SetChatAdministratorCustomTitleResponse, error) {
	var res setChatAdministratorCustomTitleResponse

	err := b.doRequest(ctx, "setChatAdministratorCustomTitle", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetChatPermissionsResponse interface
type SetChatPermissionsResponse interface {
	Response
}

type setChatPermissionsResponse struct {
	response
}

func (b *bot) SetChatPermissions(ctx context.Context, options ...MethodOption) (SetChatPermissionsResponse, error) {
	var res setChatPermissionsResponse

	err := b.doRequest(ctx, "setChatPermissions", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ExportChatInviteLinkResponse interface
type ExportChatInviteLinkResponse interface {
	Response
	GetNewInviteLink() string
}

type exportChatInviteLinkResponse struct {
	response
	Result string `json:"result,omitempty"`
}

func (r *exportChatInviteLinkResponse) GetNewInviteLink() string {
	return r.Result
}

func (b *bot) ExportChatInviteLink(ctx context.Context, options ...MethodOption) (ExportChatInviteLinkResponse, error) {
	var res exportChatInviteLinkResponse

	err := b.doRequest(ctx, "exportChatInviteLink", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateChatInviteLinkResponse interface
type CreateChatInviteLinkResponse interface {
	Response
	GetNewInviteLink() ChatInviteLink
}

type createChatInviteLinkResponse struct {
	response
	Result ChatInviteLink `json:"result,omitempty"`
}

func (r *createChatInviteLinkResponse) GetNewInviteLink() ChatInviteLink {
	return r.Result
}

func (b *bot) CreateChatInviteLink(ctx context.Context, options ...MethodOption) (CreateChatInviteLinkResponse, error) {
	var res createChatInviteLinkResponse

	err := b.doRequest(ctx, "createChatInviteLink", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// EditChatInviteLinkResponse interface
type EditChatInviteLinkResponse interface {
	Response
	GetEditedInviteLink() ChatInviteLink
}

type editChatInviteLinkResponse struct {
	response
	Result ChatInviteLink `json:"result,omitempty"`
}

func (r *editChatInviteLinkResponse) GetEditedInviteLink() ChatInviteLink {
	return r.Result
}

func (b *bot) EditChatInviteLink(ctx context.Context, options ...MethodOption) (EditChatInviteLinkResponse, error) {
	var res editChatInviteLinkResponse

	err := b.doRequest(ctx, "editChatInviteLink", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// RevokeChatInviteLinkResponse interface
type RevokeChatInviteLinkResponse interface {
	Response
	GetRevokedInviteLink() ChatInviteLink
}

type revokeChatInviteLinkResponse struct {
	response
	Result ChatInviteLink `json:"result,omitempty"`
}

func (r *revokeChatInviteLinkResponse) GetRevokedInviteLink() ChatInviteLink {
	return r.Result
}

func (b *bot) RevokeChatInviteLink(ctx context.Context, options ...MethodOption) (RevokeChatInviteLinkResponse, error) {
	var res revokeChatInviteLinkResponse

	err := b.doRequest(ctx, "revokeChatInviteLink", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ApproveChatJoinRequestResponse interface
type ApproveChatJoinRequestResponse interface {
	Response
}

type approveChatJoinRequestResponse struct {
	response
}

func (b *bot) ApproveChatJoinRequest(ctx context.Context, options ...MethodOption) (ApproveChatJoinRequestResponse, error) {
	var res approveChatJoinRequestResponse

	err := b.doRequest(ctx, "approveChatJoinRequest", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeclineChatJoinRequestResponse interface
type DeclineChatJoinRequestResponse interface {
	Response
}

type declineChatJoinRequestResponse struct {
	response
}

func (b *bot) DeclineChatJoinRequest(ctx context.Context, options ...MethodOption) (DeclineChatJoinRequestResponse, error) {
	var res declineChatJoinRequestResponse

	err := b.doRequest(ctx, "declineChatJoinRequest", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetChatPhotoResponse interface
type SetChatPhotoResponse interface {
	Response
}

type setChatPhotoResponse struct {
	response
}

func (b *bot) SetChatPhoto(ctx context.Context, options ...MethodOption) (SetChatPhotoResponse, error) {
	var res setChatPhotoResponse

	err := b.doRequest(ctx, "setChatPhoto", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteChatPhotoResponse interface
type DeleteChatPhotoResponse interface {
	Response
}

type deleteChatPhotoResponse struct {
	response
}

func (b *bot) DeleteChatPhoto(ctx context.Context, options ...MethodOption) (DeleteChatPhotoResponse, error) {
	var res deleteChatPhotoResponse

	err := b.doRequest(ctx, "deleteChatPhoto", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetChatTitleResponse interface
type SetChatTitleResponse interface {
	Response
}

type setChatTitleResponse struct {
	response
}

func (b *bot) SetChatTitle(ctx context.Context, options ...MethodOption) (SetChatTitleResponse, error) {
	var res setChatTitleResponse

	err := b.doRequest(ctx, "setChatTitle", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetChatDescriptionResponse interface
type SetChatDescriptionResponse interface {
	Response
}

type setChatDescriptionResponse struct {
	response
}

func (b *bot) SetChatDescription(ctx context.Context, options ...MethodOption) (SetChatDescriptionResponse, error) {
	var res setChatDescriptionResponse

	err := b.doRequest(ctx, "setChatDescription", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// PinChatMessageResponse interface
type PinChatMessageResponse interface {
	Response
}

type pinChatMessageResponse struct {
	response
}

func (b *bot) PinChatMessage(ctx context.Context, options ...MethodOption) (PinChatMessageResponse, error) {
	var res pinChatMessageResponse

	err := b.doRequest(ctx, "pinChatMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UnpinChatMessageResponse interface
type UnpinChatMessageResponse interface {
	Response
}

type unpinChatMessageResponse struct {
	response
}

func (b *bot) UnpinChatMessage(ctx context.Context, options ...MethodOption) (UnpinChatMessageResponse, error) {
	var res unpinChatMessageResponse

	err := b.doRequest(ctx, "unpinChatMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UnpinAllChatMessagesResponse interface
type UnpinAllChatMessagesResponse interface {
	Response
}

type unpinAllChatMessagesResponse struct {
	response
}

func (b *bot) UnpinAllChatMessages(ctx context.Context, options ...MethodOption) (UnpinAllChatMessagesResponse, error) {
	var res unpinAllChatMessagesResponse

	err := b.doRequest(ctx, "unpinAllChatMessages", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// LeaveChatResponse interface
type LeaveChatResponse interface {
	Response
}

type leaveChatResponse struct {
	response
}

func (b *bot) LeaveChat(ctx context.Context, options ...MethodOption) (LeaveChatResponse, error) {
	var res leaveChatResponse

	err := b.doRequest(ctx, "leaveChat", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetChatResponse interface
type GetChatResponse interface {
	Response
	GetChat() *Chat
}

type getChatResponse struct {
	response
	Result *Chat `json:"result,omitempty"`
}

func (r *getChatResponse) GetChat() *Chat {
	return r.Result
}

func (b *bot) GetChat(ctx context.Context, options ...MethodOption) (GetChatResponse, error) {
	var res getChatResponse

	err := b.doRequest(ctx, "getChat", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetChatAdministratorsResponse interface
type GetChatAdministratorsResponse interface {
	Response
	GetChatAdministrators() []ChatMember
}

type getChatAdministratorsResponse struct {
	response
	Result []ChatMember `json:"result,omitempty"`
}

func (r *getChatAdministratorsResponse) GetChatAdministrators() []ChatMember {
	return r.Result
}

func (b *bot) GetChatAdministrators(ctx context.Context, options ...MethodOption) (GetChatAdministratorsResponse, error) {
	var res getChatAdministratorsResponse

	err := b.doRequest(ctx, "getChatAdministrators", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetChatMemberCountResponse interface
type GetChatMemberCountResponse interface {
	Response
	GetChatMemberCount() int
}

type getChatMemberCountResponse struct {
	response
	Result int `json:"result,omitempty"`
}

func (r *getChatMemberCountResponse) GetChatMemberCount() int {
	return r.Result
}

func (b *bot) GetChatMemberCount(ctx context.Context, options ...MethodOption) (GetChatMemberCountResponse, error) {
	var res getChatMemberCountResponse

	err := b.doRequest(ctx, "getChatMemberCount", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetChatMemberResponse interface
type GetChatMemberResponse interface {
	Response
	GetChatMember() *ChatMember
}

type getChatMemberResponse struct {
	response
	Result *ChatMember `json:"result,omitempty"`
}

func (r *getChatMemberResponse) GetChatMember() *ChatMember {
	return r.Result
}

func (b *bot) GetChatMember(ctx context.Context, options ...MethodOption) (GetChatMemberResponse, error) {
	var res getChatMemberResponse

	err := b.doRequest(ctx, "getChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetChatStickerSetResponse interface
type SetChatStickerSetResponse interface {
	Response
}

type setChatStickerSetResponse struct {
	response
}

func (b *bot) SetChatStickerSet(ctx context.Context, options ...MethodOption) (SetChatStickerSetResponse, error) {
	var res setChatStickerSetResponse

	err := b.doRequest(ctx, "setChatStickerSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteChatStickerSetResponse interface
type DeleteChatStickerSetResponse interface {
	Response
}

type deleteChatStickerSetResponse struct {
	response
}

func (b *bot) DeleteChatStickerSet(ctx context.Context, options ...MethodOption) (DeleteChatStickerSetResponse, error) {
	var res deleteChatStickerSetResponse

	err := b.doRequest(ctx, "deleteChatStickerSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// AnswerCallbackQueryResponse interface
type AnswerCallbackQueryResponse interface {
	Response
}

type answerCallbackQueryResponse struct {
	response
}

func (b *bot) AnswerCallbackQuery(ctx context.Context, options ...MethodOption) (AnswerCallbackQueryResponse, error) {
	var res answerCallbackQueryResponse

	err := b.doRequest(ctx, "answerCallbackQuery", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetMyCommandsResponse interface
type SetMyCommandsResponse interface {
	Response
}

type setMyCommandsResponse struct {
	response
}

func (b *bot) SetMyCommands(ctx context.Context, options ...MethodOption) (SetMyCommandsResponse, error) {
	var res setMyCommandsResponse

	err := b.doRequest(ctx, "setMyCommands", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteMyCommandsResponse interface
type DeleteMyCommandsResponse interface {
	Response
}

type deleteMyCommandsResponse struct {
	response
}

func (b *bot) DeleteMyCommands(ctx context.Context, options ...MethodOption) (DeleteMyCommandsResponse, error) {
	var res deleteMyCommandsResponse

	err := b.doRequest(ctx, "deleteMyCommands", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetMyCommandsResponse interface
type GetMyCommandsResponse interface {
	Response
	GetCommands() []BotCommand
}

type getMyCommandsResponse struct {
	response
	Result []BotCommand `json:"result,omitempty"`
}

func (r *getMyCommandsResponse) GetCommands() []BotCommand {
	return r.Result
}

func (b *bot) GetMyCommands(ctx context.Context) (GetMyCommandsResponse, error) {
	var res getMyCommandsResponse

	err := b.doRequest(ctx, "getMyCommands", &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

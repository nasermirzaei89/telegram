package telegram

// GetMeResponse interface
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
	err := doRequest(b.token, "getMe", &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SendMessageResponse interface
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

func (b *bot) SendMessage(options ...Option) (SendMessageResponse, error) {
	var res sendMessageResponse
	err := doRequest(b.token, "sendMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ForwardMessageResponse interface
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

func (b *bot) ForwardMessage(options ...Option) (ForwardMessageResponse, error) {
	var res forwardMessageResponse
	err := doRequest(b.token, "forwardMessage", &res, options...)
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

func (b *bot) SendPhoto(options ...Option) (SendPhotoResponse, error) {
	var res sendPhotoResponse
	err := doRequest(b.token, "sendPhoto", &res, options...)
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

func (b *bot) SendAudio(options ...Option) (SendAudioResponse, error) {
	var res sendAudioResponse
	err := doRequest(b.token, "sendAudio", &res, options...)
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

func (b *bot) SendDocument(options ...Option) (SendDocumentResponse, error) {
	var res sendDocumentResponse
	err := doRequest(b.token, "sendDocument", &res, options...)
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

func (b *bot) SendVideo(options ...Option) (SendVideoResponse, error) {
	var res sendVideoResponse
	err := doRequest(b.token, "sendVideo", &res, options...)
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

func (b *bot) SendAnimation(options ...Option) (SendAnimationResponse, error) {
	var res sendAnimationResponse
	err := doRequest(b.token, "sendAnimation", &res, options...)
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

func (b *bot) SendVoice(options ...Option) (SendVoiceResponse, error) {
	var res sendVoiceResponse
	err := doRequest(b.token, "sendVoice", &res, options...)
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

func (b *bot) SendVideoNote(options ...Option) (SendVideoNoteResponse, error) {
	var res sendVideoNoteResponse
	err := doRequest(b.token, "sendVideoNote", &res, options...)
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

func (b *bot) SendMediaGroup(options ...Option) (SendMediaGroupResponse, error) {
	var res sendMediaGroupResponse
	err := doRequest(b.token, "sendMediaGroup", &res, options...)
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

func (b *bot) SendLocation(options ...Option) (SendLocationResponse, error) {
	var res sendLocationResponse
	err := doRequest(b.token, "sendLocation", &res, options...)
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

func (b *bot) EditMessageLiveLocation(options ...Option) (EditMessageLiveLocationResponse, error) {
	var res editMessageLiveLocationResponse
	err := doRequest(b.token, "editMessageLiveLocation", &res, options...)
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

func (b *bot) StopMessageLiveLocation(options ...Option) (StopMessageLiveLocationResponse, error) {
	var res stopMessageLiveLocationResponse
	err := doRequest(b.token, "stopMessageLiveLocation", &res, options...)
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

func (b *bot) SendVenue(options ...Option) (SendVenueResponse, error) {
	var res sendVenueResponse
	err := doRequest(b.token, "sendVenue", &res, options...)
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

func (b *bot) SendContact(options ...Option) (SendContactResponse, error) {
	var res sendContactResponse
	err := doRequest(b.token, "sendContact", &res, options...)
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

func (b *bot) SendPoll(options ...Option) (SendPollResponse, error) {
	var res sendPollResponse
	err := doRequest(b.token, "sendPoll", &res, options...)
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

func (b *bot) SendChatAction(options ...Option) (SendChatActionResponse, error) {
	var res sendChatActionResponse
	err := doRequest(b.token, "sendChatAction", &res, options...)
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

func (b *bot) GetUserProfilePhotos(options ...Option) (GetUserProfilePhotosResponse, error) {
	var res getUserProfilePhotosResponse
	err := doRequest(b.token, "getUserProfilePhotos", &res, options...)
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

func (b *bot) GetFile(options ...Option) (GetFileResponse, error) {
	var res getFileResponse
	err := doRequest(b.token, "getFile", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// KickChatMemberResponse interface
type KickChatMemberResponse interface {
	Response
}

type kickChatMemberResponse struct {
	response
}

func (b *bot) KickChatMember(options ...Option) (KickChatMemberResponse, error) {
	var res kickChatMemberResponse
	err := doRequest(b.token, "kickChatMember", &res, options...)
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

func (b *bot) UnbanChatMember(options ...Option) (UnbanChatMemberResponse, error) {
	var res unbanChatMemberResponse
	err := doRequest(b.token, "unbanChatMember", &res, options...)
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

func (b *bot) RestrictChatMember(options ...Option) (RestrictChatMemberResponse, error) {
	var res restrictChatMemberResponse
	err := doRequest(b.token, "restrictChatMember", &res, options...)
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

func (b *bot) PromoteChatMember(options ...Option) (PromoteChatMemberResponse, error) {
	var res promoteChatMemberResponse
	err := doRequest(b.token, "promoteChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ExportChatInviteLinkResponse interface
type ExportChatInviteLinkResponse interface {
	Response
	GetInviteLink() string
}

type exportChatInviteLinkResponse struct {
	response
	Result string `json:"result,omitempty"`
}

func (r *exportChatInviteLinkResponse) GetInviteLink() string {
	return r.Result
}

func (b *bot) ExportChatInviteLink(options ...Option) (ExportChatInviteLinkResponse, error) {
	var res exportChatInviteLinkResponse
	err := doRequest(b.token, "exportChatInviteLink", &res, options...)
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

func (b *bot) SetChatPhoto(options ...Option) (SetChatPhotoResponse, error) {
	var res setChatPhotoResponse
	err := doRequest(b.token, "setChatPhoto", &res, options...)
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

func (b *bot) DeleteChatPhoto(options ...Option) (DeleteChatPhotoResponse, error) {
	var res deleteChatPhotoResponse
	err := doRequest(b.token, "deleteChatPhoto", &res, options...)
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

func (b *bot) SetChatTitle(options ...Option) (SetChatTitleResponse, error) {
	var res setChatTitleResponse
	err := doRequest(b.token, "setChatTitle", &res, options...)
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

func (b *bot) SetChatDescription(options ...Option) (SetChatDescriptionResponse, error) {
	var res setChatDescriptionResponse
	err := doRequest(b.token, "setChatDescription", &res, options...)
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

func (b *bot) PinChatMessage(options ...Option) (PinChatMessageResponse, error) {
	var res pinChatMessageResponse
	err := doRequest(b.token, "pinChatMessage", &res, options...)
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

func (b *bot) UnpinChatMessage(options ...Option) (UnpinChatMessageResponse, error) {
	var res unpinChatMessageResponse
	err := doRequest(b.token, "unpinChatMessage", &res, options...)
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

func (b *bot) LeaveChat(options ...Option) (LeaveChatResponse, error) {
	var res leaveChatResponse
	err := doRequest(b.token, "leaveChat", &res, options...)
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

func (b *bot) GetChat(options ...Option) (GetChatResponse, error) {
	var res getChatResponse
	err := doRequest(b.token, "getChat", &res, options...)
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

func (b *bot) GetChatAdministrators(options ...Option) (GetChatAdministratorsResponse, error) {
	var res getChatAdministratorsResponse
	err := doRequest(b.token, "getChatAdministrators", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetChatMembersCountResponse interface
type GetChatMembersCountResponse interface {
	Response
	GetChatMembersCount() int
}

type getChatMembersCountResponse struct {
	response
	Result int `json:"result,omitempty"`
}

func (r *getChatMembersCountResponse) GetChatMembersCount() int {
	return r.Result
}

func (b *bot) GetChatMembersCount(options ...Option) (GetChatMembersCountResponse, error) {
	var res getChatMembersCountResponse
	err := doRequest(b.token, "getChatMembersCount", &res, options...)
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

func (b *bot) GetChatMember(options ...Option) (GetChatMemberResponse, error) {
	var res getChatMemberResponse
	err := doRequest(b.token, "getChatMember", &res, options...)
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

func (b *bot) SetChatStickerSet(options ...Option) (SetChatStickerSetResponse, error) {
	var res setChatStickerSetResponse
	err := doRequest(b.token, "setChatStickerSet", &res, options...)
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

func (b *bot) DeleteChatStickerSet(options ...Option) (DeleteChatStickerSetResponse, error) {
	var res deleteChatStickerSetResponse
	err := doRequest(b.token, "deleteChatStickerSet", &res, options...)
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

func (b *bot) AnswerCallbackQuery(options ...Option) (AnswerCallbackQueryResponse, error) {
	var res answerCallbackQueryResponse
	err := doRequest(b.token, "answerCallbackQuery", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

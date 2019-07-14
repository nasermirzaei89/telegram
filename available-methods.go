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
	err := doRequest(b.Token, "sendMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "forwardMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendPhoto", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendAudio", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendDocument", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendVideo", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendAnimation", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendVoice", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendVideoNote", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendMediaGroup", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendLocation", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "editMessageLiveLocation", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "stopMessageLiveLocation", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendVenue", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendContact", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "sendPoll", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type SendChatActionResponse interface {
	Response
}

type sendChatActionResponse struct {
	response
}

func (b *bot) SendChatAction(options ...Option) (SendChatActionResponse, error) {
	var res sendChatActionResponse
	err := doRequest(b.Token, "sendChatAction", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "getUserProfilePhotos", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "getFile", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type KickChatMemberResponse interface {
	Response
}

type kickChatMemberResponse struct {
	response
}

func (b *bot) KickChatMember(options ...Option) (KickChatMemberResponse, error) {
	var res kickChatMemberResponse
	err := doRequest(b.Token, "kickChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type UnbanChatMemberResponse interface {
	Response
}

type unbanChatMemberResponse struct {
	response
}

func (b *bot) UnbanChatMember(options ...Option) (UnbanChatMemberResponse, error) {
	var res unbanChatMemberResponse
	err := doRequest(b.Token, "unbanChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type RestrictChatMemberResponse interface {
	Response
}

type restrictChatMemberResponse struct {
	response
}

func (b *bot) RestrictChatMember(options ...Option) (RestrictChatMemberResponse, error) {
	var res restrictChatMemberResponse
	err := doRequest(b.Token, "restrictChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type PromoteChatMemberResponse interface {
	Response
}

type promoteChatMemberResponse struct {
	response
}

func (b *bot) PromoteChatMember(options ...Option) (PromoteChatMemberResponse, error) {
	var res promoteChatMemberResponse
	err := doRequest(b.Token, "promoteChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "exportChatInviteLink", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type SetChatPhotoResponse interface {
	Response
}

type setChatPhotoResponse struct {
	response
}

func (b *bot) SetChatPhoto(options ...Option) (SetChatPhotoResponse, error) {
	var res setChatPhotoResponse
	err := doRequest(b.Token, "setChatPhoto", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type DeleteChatPhotoResponse interface {
	Response
}

type deleteChatPhotoResponse struct {
	response
}

func (b *bot) DeleteChatPhoto(options ...Option) (DeleteChatPhotoResponse, error) {
	var res deleteChatPhotoResponse
	err := doRequest(b.Token, "deleteChatPhoto", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type SetChatTitleResponse interface {
	Response
}

type setChatTitleResponse struct {
	response
}

func (b *bot) SetChatTitle(options ...Option) (SetChatTitleResponse, error) {
	var res setChatTitleResponse
	err := doRequest(b.Token, "setChatTitle", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type SetChatDescriptionResponse interface {
	Response
}

type setChatDescriptionResponse struct {
	response
}

func (b *bot) SetChatDescription(options ...Option) (SetChatDescriptionResponse, error) {
	var res setChatDescriptionResponse
	err := doRequest(b.Token, "setChatDescription", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type PinChatMessageResponse interface {
	Response
}

type pinChatMessageResponse struct {
	response
}

func (b *bot) PinChatMessage(options ...Option) (PinChatMessageResponse, error) {
	var res pinChatMessageResponse
	err := doRequest(b.Token, "pinChatMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type UnpinChatMessageResponse interface {
	Response
}

type unpinChatMessageResponse struct {
	response
}

func (b *bot) UnpinChatMessage(options ...Option) (UnpinChatMessageResponse, error) {
	var res unpinChatMessageResponse
	err := doRequest(b.Token, "unpinChatMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type LeaveChatResponse interface {
	Response
}

type leaveChatResponse struct {
	response
}

func (b *bot) LeaveChat(options ...Option) (LeaveChatResponse, error) {
	var res leaveChatResponse
	err := doRequest(b.Token, "leaveChat", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "getChat", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "getChatAdministrators", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "getChatMembersCount", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

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
	err := doRequest(b.Token, "getChatMember", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type SetChatStickerSetResponse interface {
	Response
}

type setChatStickerSetResponse struct {
	response
}

func (b *bot) SetChatStickerSet(options ...Option) (SetChatStickerSetResponse, error) {
	var res setChatStickerSetResponse
	err := doRequest(b.Token, "setChatStickerSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type DeleteChatStickerSetResponse interface {
	Response
}

type deleteChatStickerSetResponse struct {
	response
}

func (b *bot) DeleteChatStickerSet(options ...Option) (DeleteChatStickerSetResponse, error) {
	var res deleteChatStickerSetResponse
	err := doRequest(b.Token, "deleteChatStickerSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type AnswerCallbackQueryResponse interface {
	Response
}

type answerCallbackQueryResponse struct {
	response
}

func (b *bot) AnswerCallbackQuery(options ...Option) (AnswerCallbackQueryResponse, error) {
	var res answerCallbackQueryResponse
	err := doRequest(b.Token, "answerCallbackQuery", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

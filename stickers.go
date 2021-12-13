package telegram

import "context"

// Sticker struct.
type Sticker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        *string       `json:"emoji,omitempty"`
	SetName      *string       `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     *int          `json:"file_size,omitempty"`
}

// StickerSet struct.
type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []Sticker  `json:"stickers"`
	Thumb         *PhotoSize `json:"thumb,omitempty"`
}

// MaskPosition struct.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// SendStickerResponse interface.
type SendStickerResponse interface {
	Response
	GetMessage() *Message
}

type sendStickerResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendStickerResponse) GetMessage() *Message {
	return r.Result
}

func (b *Bot) SendSticker(ctx context.Context, options ...MethodOption) (SendStickerResponse, error) {
	var res sendStickerResponse

	err := b.doRequest(ctx, "sendSticker", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetStickerSetResponse interface.
type GetStickerSetResponse interface {
	Response
	GetStickerSet() *StickerSet
}

type getStickerSetResponse struct {
	response
	Result *StickerSet `json:"result,omitempty"`
}

func (r *getStickerSetResponse) GetStickerSet() *StickerSet {
	return r.Result
}

func (b *Bot) GetStickerSet(ctx context.Context, options ...MethodOption) (GetStickerSetResponse, error) {
	var res getStickerSetResponse

	err := b.doRequest(ctx, "getStickerSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UploadStickerFileResponse interface.
type UploadStickerFileResponse interface {
	Response
	GetUploadedFile() *File
}

type uploadStickerFileResponse struct {
	response
	Result *File `json:"result,omitempty"`
}

func (r *uploadStickerFileResponse) GetUploadedFile() *File {
	return r.Result
}

func (b *Bot) UploadStickerFile(ctx context.Context, options ...MethodOption) (UploadStickerFileResponse, error) {
	var res uploadStickerFileResponse

	err := b.doRequest(ctx, "uploadStickerFile", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateNewStickerSetResponse interface.
type CreateNewStickerSetResponse interface {
	Response
}

type createNewStickerSetResponse struct {
	response
}

func (b *Bot) CreateNewStickerSet(ctx context.Context, options ...MethodOption) (CreateNewStickerSetResponse, error) {
	var res createNewStickerSetResponse

	err := b.doRequest(ctx, "createNewStickerSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// AddStickerToSetResponse interface.
type AddStickerToSetResponse interface {
	Response
}

type addStickerToSetResponse struct {
	response
}

func (b *Bot) AddStickerToSet(ctx context.Context, options ...MethodOption) (AddStickerToSetResponse, error) {
	var res addStickerToSetResponse

	err := b.doRequest(ctx, "addStickerToSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetStickerPositionInSetResponse interface.
type SetStickerPositionInSetResponse interface {
	Response
}

type setStickerPositionInSetResponse struct {
	response
}

func (b *Bot) SetStickerPositionInSet(ctx context.Context, options ...MethodOption) (SetStickerPositionInSetResponse, error) {
	var res setStickerPositionInSetResponse

	err := b.doRequest(ctx, "setStickerPositionInSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteStickerFromSetResponse interface.
type DeleteStickerFromSetResponse interface {
	Response
}

type deleteStickerFromSetResponse struct {
	response
}

func (b *Bot) DeleteStickerFromSet(ctx context.Context, options ...MethodOption) (DeleteStickerFromSetResponse, error) {
	var res deleteStickerFromSetResponse

	err := b.doRequest(ctx, "deleteStickerFromSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetStickerSetThumbResponse interface.
type SetStickerSetThumbResponse interface {
	Response
}

type setStickerSetThumbResponse struct {
	response
}

func (b *Bot) SetStickerSetThumb(ctx context.Context, options ...MethodOption) (SetStickerSetThumbResponse, error) {
	var res setStickerSetThumbResponse

	err := b.doRequest(ctx, "setStickerSetThumb", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

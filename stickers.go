package telegram

type Sticker struct {
	FileID       string        `json:"file_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        *string       `json:"emoji,omitempty"`
	SetName      *string       `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     *int          `json:"file_size,omitempty"`
}

type StickerSet struct {
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	ContainsMasks bool      `json:"contains_masks"`
	Stickers      []Sticker `json:"stickers"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

type SendStickerResponse interface {
}

func (b *bot) SendSticker(options ...Option) (SendStickerResponse, error) {
	panic("implement me")
}

type GetStickerSetResponse interface {
	Name(string) GetStickerSetResponse
	Do() (*StickerSet, error)
}

func (b *bot) GetStickerSet(options ...Option) (GetStickerSetResponse, error) {
	panic("implement me")
}

type UploadStickerFileResponse interface {
	UserID(int) UploadStickerFileResponse
	PNGSticker(InputFile) UploadStickerFileResponse
	Do() (*File, error)
}

func (b *bot) UploadStickerFile(options ...Option) (UploadStickerFileResponse, error) {
	panic("implement me")
}

type CreateNewStickerSetResponse interface {
}

func (b *bot) CreateNewStickerSet(options ...Option) (CreateNewStickerSetResponse, error) {
	panic("implement me")
}

type AddStickerToSetResponse interface {
}

func (b *bot) AddStickerToSet(options ...Option) (AddStickerToSetResponse, error) {
	panic("implement me")
}

type SetStickerPositionInSetResponse interface {
}

func (b *bot) SetStickerPositionInSet(options ...Option) (SetStickerPositionInSetResponse, error) {
	panic("implement me")
}

type DeleteStickerFromSetResponse interface {
}

func (b *bot) DeleteStickerFromSet(options ...Option) (DeleteStickerFromSetResponse, error) {
	panic("implement me")
}

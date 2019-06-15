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

type SendStickerRequest interface {
	ChatID(int) SendStickerRequest
	ChatUsername(string) SendStickerRequest
	Sticker(InputFile) SendStickerRequest
	StickerFileID(string) SendStickerRequest
	StickerURL(string) SendStickerRequest
	DisableNotification() SendStickerRequest
	ReplyToMessageID(int) SendStickerRequest
	ReplyMarkup(interface{}) SendStickerRequest
	Do() (*Message, error)
}

func (b *bot) SendSticker() SendStickerRequest {
	panic("implement me")
}

type GetStickerSetRequest interface {
	Name(string) GetStickerSetRequest
	Do() (*StickerSet, error)
}

func (b *bot) GetStickerSet() GetStickerSetRequest {
	panic("implement me")
}

type UploadStickerFileRequest interface {
	UserID(int) UploadStickerFileRequest
	PNGSticker(InputFile) UploadStickerFileRequest
	Do() (*File, error)
}

func (b *bot) UploadStickerFile() UploadStickerFileRequest {
	panic("implement me")
}

type CreateNewStickerSetRequest interface {
	UserID(int) CreateNewStickerSetRequest
	Name(string) CreateNewStickerSetRequest
	Title(string) CreateNewStickerSetRequest
	PNGSticker(InputFile) CreateNewStickerSetRequest
	PNGStickerFileID(string) CreateNewStickerSetRequest
	PNGStickerURL(string) CreateNewStickerSetRequest
	Emojis(string) CreateNewStickerSetRequest
	ContainsMasks() CreateNewStickerSetRequest
	MaskPosition(MaskPosition) CreateNewStickerSetRequest
	Do() (bool, error)
}

func (b *bot) CreateNewStickerSet() CreateNewStickerSetRequest {
	panic("implement me")
}

type AddStickerToSetRequest interface {
	UserID(int) AddStickerToSetRequest
	Name(string) AddStickerToSetRequest
	PNGSticker(InputFile) AddStickerToSetRequest
	PNGStickerFileID(string) AddStickerToSetRequest
	PNGStickerURL(string) AddStickerToSetRequest
	Emojis(string) AddStickerToSetRequest
	MaskPosition(MaskPosition) AddStickerToSetRequest
	Do() (bool, error)
}

func (b *bot) AddStickerToSet() AddStickerToSetRequest {
	panic("implement me")
}

type SetStickerPositionInSetRequest interface {
	Sticker(string) SetStickerPositionInSetRequest
	Position(int) SetStickerPositionInSetRequest
	Do() (bool, error)
}

func (b *bot) SetStickerPositionInSet() SetStickerPositionInSetRequest {
	panic("implement me")
}

type DeleteStickerFromSetRequest interface {
	Sticker(string) DeleteStickerFromSetRequest
	Do() (bool, error)
}

func (b *bot) DeleteStickerFromSet() DeleteStickerFromSetRequest {
	panic("implement me")
}

package telegram

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

type PassportFile struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FileDate int    `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        *string        `json:"data,omitempty"`
	PhoneNumber *string        `json:"phone_number,omitempty"`
	Email       *string        `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type SetPassportDataErrorsRequest interface {
	UserID(int) SetPassportDataErrorsRequest
	Errors([]PassportElementError) SetPassportDataErrorsRequest
	Do() (bool, error)
}

func (b *bot) SetPassportDataErrors() SetPassportDataErrorsRequest {
	panic("implement me")
}

type PassportElementError interface{}

type PassportElementErrorDataField struct {
	Source    string `json:"source"`
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorSelfie struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

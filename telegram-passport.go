package telegram

// PassportData struct
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

// PassportFile struct
type PassportFile struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FileDate int    `json:"file_date"`
}

// EncryptedPassportElement struct
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

// EncryptedCredentials struct
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// SetPassportDataErrorsResponse interface
type SetPassportDataErrorsResponse interface {
	Response
}

type setPassportDataErrorsResponse struct {
	response
}

func (b *bot) SetPassportDataErrors(options ...Option) (SetPassportDataErrorsResponse, error) {
	var res setPassportDataErrorsResponse
	err := doRequest(b.Token, "setPassportDataErrors", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// PassportElementError interface
type PassportElementError interface{}

// PassportElementErrorDataField struct
type PassportElementErrorDataField struct {
	Source    string `json:"source"`
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

// PassportElementErrorFrontSide struct
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorReverseSide struct
type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorSelfie struct
type PassportElementErrorSelfie struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorFile struct
type PassportElementErrorFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorFiles struct
type PassportElementErrorFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

// PassportElementErrorTranslationFile struct
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorTranslationFiles struct
type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

// PassportElementErrorUnspecified struct
type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

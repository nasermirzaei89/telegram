package telegram

import "io"

type User struct {
	ID           int     `json:"id"`
	IsBot        bool    `json:"is_bot"`
	FirstName    string  `json:"first_name"`
	LastName     *string `json:"last_name,omitempty"`
	Username     *string `json:"username,omitempty"`
	LanguageCode *string `json:"language_code,omitempty"`
}

type Chat struct {
	ID                          int        `json:"id"`
	Type                        string     `json:"type"`
	Title                       *string    `json:"title,omitempty"`
	Username                    *string    `json:"username,omitempty"`
	FirstName                   *string    `json:"first_name,omitempty"`
	LastName                    *string    `json:"last_name,omitempty"`
	AllMembersAreAdministrators *bool      `json:"all_members_are_administrators,omitempty"`
	Photo                       *ChatPhoto `json:"photo,omitempty"`
	Description                 *string    `json:"description,omitempty"`
	InviteLink                  *string    `json:"invite_link,omitempty"`
	PinnedMessage               *Message   `json:"pinned_message,omitempty"`
	StickerSetName              *string    `json:"sticker_set_name,omitempty"`
	CanSetStickerSet            *bool      `json:"can_set_sticker_set,omitempty"`
}

type Message struct {
	MessageID             int                   `json:"message_id"`
	From                  *User                 `json:"from,omitempty"`
	Date                  int                   `json:"date"`
	Chat                  Chat                  `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat                 `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  *int                  `json:"forward_from_message_id,omitempty"`
	ForwardSignature      *string               `json:"forward_signature,omitempty"`
	ForwardSenderName     *string               `json:"forward_sender_name,omitempty"`
	ForwardDate           *int                  `json:"forward_date,omitempty"`
	ReplyToMessage        *Message              `json:"reply_to_message,omitempty"`
	EditDate              *int                  `json:"edit_date,omitempty"`
	MediaGroupID          *string               `json:"media_group_id,omitempty"`
	AuthorSignature       *string               `json:"author_signature,omitempty"`
	Text                  *string               `json:"text,omitempty"`
	Entities              []MessageEntity       `json:"entities,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	Audio                 *Audio                `json:"audio,omitempty"`
	Document              *Document             `json:"document,omitempty"`
	Animation             *Animation            `json:"animation,omitempty"`
	Game                  *Game                 `json:"game,omitempty"`
	Photo                 []PhotoSize           `json:"photo,omitempty"`
	Sticker               *Sticker              `json:"sticker,omitempty"`
	Video                 *Video                `json:"video,omitempty"`
	Voice                 *Voice                `json:"voice,omitempty"`
	VideoNote             *VideoNote            `json:"video_note,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	Contact               *Contact              `json:"contact,omitempty"`
	Location              *Location             `json:"location,omitempty"`
	Venue                 *Venue                `json:"venue,omitempty"`
	Poll                  *Poll                 `json:"poll,omitempty"`
	NewChatMembers        []User                `json:"new_chat_members,omitempty"`
	LeftChatMember        *User                 `json:"left_chat_member,omitempty"`
	NewChatTitle          *string               `json:"new_chat_title,omitempty"`
	NewChatPhoto          []PhotoSize           `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       *bool                 `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      *bool                 `json:"group_chat_created,omitempty"`
	SupergroupChatCreated *bool                 `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    *bool                 `json:"channel_chat_created,omitempty"`
	MigrateToChatID       *int                  `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     *int                  `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message              `json:"pinned_message,omitempty"`
	Invoice               *Invoice              `json:"invoice,omitempty"`
	SuccessfulPayment     *SuccessfulPayment    `json:"successful_payment,omitempty"`
	ConnectedWebsite      *string               `json:"connected_website,omitempty"`
	PassportData          *PassportData         `json:"passport_data,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

//// GetCommand present a more flexible way to communicate with your bot.
//func (m *Message) GetCommand() (*string, []string) {
//	if m.Text != nil {
//		if (*m.Text)[0] == '/' {
//			ss := strings.Split(*m.Text, " ")
//			command := ss[0][1:]
//			return &command, ss[1:]
//		}
//	}
//
//	return nil, nil
//}
//
type MessageEntity struct {
	Type   string  `json:"type"`
	Offset int     `json:"offset"`
	Length int     `json:"length"`
	URL    *string `json:"url,omitempty"`
	User   *User   `json:"user,omitempty"`
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize *int   `json:"file_size,omitempty"`
}

type Audio struct {
	FileID    string     `json:"file_id"`
	Duration  int        `json:"duration"`
	Performer *string    `json:"performer,omitempty"`
	Title     *string    `json:"title,omitempty"`
	MimeType  *string    `json:"mime_type,omitempty"`
	FileSize  *int       `json:"file_size,omitempty"`
	Thumb     *PhotoSize `json:"thumb,omitempty"`
}

type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName *string    `json:"file_name,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize *int       `json:"file_size,omitempty"`
}

type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize *int       `json:"file_size,omitempty"`
}

type Animation struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	Filename *string    `json:"filename,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize *int       `json:"file_size,omitempty"`
}

type Voice struct {
	FileID   string  `json:"file_id"`
	Duration int     `json:"duration"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize *int    `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileID   string     `json:"file_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumbs   *PhotoSize `json:"thumbs,omitempty"`
	FileSize *int       `json:"file_size,omitempty"`
}

type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int    `json:"user_id,omitempty"`
	VCard       *string `json:"vcard,omitempty"`
}

type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type Venue struct {
	Location       Location `json:"location"`
	Title          string   `json:"title"`
	Address        string   `json:"address"`
	FoursquareID   *string  `json:"foursquare_id,omitempty"`
	FoursquareType *string  `json:"foursquare_type,omitempty"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type Poll struct {
	ID       string       `json:"id"`
	Question string       `json:"question"`
	Options  []PollOption `json:"options"`
	IsClosed bool         `json:"is_closed"`
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type File struct {
	FileID   string  `json:"file_id"`
	FileSize *int    `json:"file_size,omitempty"`
	FilePath *string `json:"file_path,omitempty"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  *bool              `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard *bool              `json:"one_time_keyboard,omitempty"`
	Selective       *bool              `json:"selective,omitempty"`
}

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  *bool  `json:"request_contact,omitempty"`
	RequestLocation *bool  `json:"request_location,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          *string       `json:"url,omitempty"`
	LoginURL                     *LoginURL     `json:"login_url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          *bool         `json:"pay,omitempty"`
}

type LoginURL struct {
	URL                string  `json:"url"`
	ForwardText        *string `json:"forward_text,omitempty"`
	BotUsername        *string `json:"bot_username,omitempty"`
	RequestWriteAccess *bool   `json:"request_write_access,omitempty"`
}

type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID *string  `json:"inline_message_id,omitempty"`
	ChatInstance    *string  `json:"chat_instance,omitempty"`
	Data            *string  `json:"data,omitempty"`
	GameShortName   *string  `json:"game_short_name,omitempty"`
}

type ForceReply struct {
	ForceReply bool  `json:"force_reply"`
	Selective  *bool `json:"selective,omitempty"`
}

type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

type ChatMember struct {
	User                  User   `json:"user"`
	Status                string `json:"status"`
	UntilDate             *int   `json:"until_date,omitempty"`
	CanBeEdited           *bool  `json:"can_be_edited,omitempty"`
	CanChangeInfo         *bool  `json:"can_change_info,omitempty"`
	CanPostMessages       *bool  `json:"can_post_messages,omitempty"`
	CanEditMessages       *bool  `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     *bool  `json:"can_delete_messages,omitempty"`
	CanInviteUsers        *bool  `json:"can_invite_users,omitempty"`
	CanRestrictMembers    *bool  `json:"can_restrict_members,omitempty"`
	CanPinMessages        *bool  `json:"can_pin_messages,omitempty"`
	CanPromoteMembers     *bool  `json:"can_promote_members,omitempty"`
	IsMember              *bool  `json:"is_member,omitempty"`
	CanSendMessages       *bool  `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  *bool  `json:"can_send_media_messages,omitempty"`
	CanSendOtherMessages  *bool  `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool  `json:"can_add_web_page_previews,omitempty"`
}

type ResponseParameters struct {
	MigrateToChatID *int `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int `json:"retry_after,omitempty"`
}

type InputMedia interface {
}

type InputMediaPhoto struct {
	Type      string  `json:"type"`
	Media     string  `json:"media"`
	Caption   *string `json:"caption,omitempty"`
	ParseMode *string `json:"parse_mode,omitempty"`
}

type InputMediaVideo struct {
	Type              string      `json:"type"`
	Media             string      `json:"media"`
	Thumb             interface{} `json:"thumb,omitempty"`
	Caption           *string     `json:"caption,omitempty"`
	ParseMode         *string     `json:"parse_mode,omitempty"`
	Width             *int        `json:"width,omitempty"`
	Height            *int        `json:"height,omitempty"`
	Duration          *int        `json:"duration,omitempty"`
	SupportsStreaming *bool       `json:"supports_streaming,omitempty"`
}

type InputMediaAnimation struct {
	Type      string      `json:"type"`
	Media     string      `json:"media"`
	Thumb     interface{} `json:"thumb,omitempty"`
	Caption   *string     `json:"caption,omitempty"`
	ParseMode *string     `json:"parse_mode,omitempty"`
	Width     *int        `json:"width,omitempty"`
	Height    *int        `json:"height,omitempty"`
	Duration  *int        `json:"duration,omitempty"`
}

type InputMediaAudio struct {
	Type      string      `json:"type"`
	Media     string      `json:"media"`
	Thumb     interface{} `json:"thumb,omitempty"`
	Caption   *string     `json:"caption,omitempty"`
	ParseMode *string     `json:"parse_mode,omitempty"`
	Duration  *int        `json:"duration,omitempty"`
	Performer *string     `json:"performer,omitempty"`
	Title     *string     `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Type      string      `json:"type"`
	Media     string      `json:"media"`
	Thumb     interface{} `json:"thumb,omitempty"`
	Caption   *string     `json:"caption,omitempty"`
	ParseMode *string     `json:"parse_mode,omitempty"`
}

type InputFile interface {
	io.Reader
}

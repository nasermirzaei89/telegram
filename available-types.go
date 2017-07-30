package telegram

import (
	"io"
	"strings"
)

// User represents a Telegram user or bot.
type User struct {
	ID           int64   `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     *string `json:"last_name,omitempty"`
	Username     *string `json:"username,omitempty"`
	LanguageCode *string `json:"language_code,omitempty"`
}

// Chat represents a chat.
type Chat struct {
	ID                          int64   `json:"id"`
	Type                        string  `json:"type"`
	Title                       *string `json:"title,omitempty"`
	Username                    *string `json:"username,omitempty"`
	FirstName                   *string `json:"first_name,omitempty"`
	LastName                    *string `json:"last_name,omitempty"`
	AllMembersAreAdministrators *bool   `json:"all_members_are_administrators,omitempty"`
}

// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
const (
	ChatTypePrivate    = "private"
	ChatTypeGroup      = "group"
	ChatTypeSuperGroup = "supergroup"
	ChatTypeChannel    = "channel"
)

// Message represents a message.
type Message struct {
	MessageID             int64            `json:"message_id"`
	From                  *User            `json:"from,omitempty"`
	Date                  int64            `json:"date"`
	Chat                  Chat             `json:"chat"`
	ForwardFrom           *User            `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat            `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  *int64           `json:"forward_from_message_id,omitempty"`
	ForwardDate           *int64           `json:"forward_date,omitempty"`
	ReplyToMessage        *Message         `json:"reply_to_message,omitempty"`
	EditDate              *int64           `json:"edit_date,omitempty"`
	Text                  *string          `json:"text,omitempty"`
	Entities              *[]MessageEntity `json:"entities,omitempty"`
	Audio                 *Audio           `json:"audio,omitempty"`
	Document              *Document        `json:"document,omitempty"`
	Game                  *Game            `json:"game,omitempty"`
	Photo                 *[]PhotoSize     `json:"photo,omitempty"`
	Sticker               *Sticker         `json:"sticker,omitempty"`
	Video                 *Video           `json:"video,omitempty"`
	Voice                 *Voice           `json:"voice,omitempty"`
	Caption               *string          `json:"caption,omitempty"`
	Contact               *Contact         `json:"contact,omitempty"`
	Location              *Location        `json:"location,omitempty"`
	Venue                 *Venue           `json:"venue,omitempty"`
	NewChatMember         *User            `json:"new_chat_member,omitempty"`
	LeftChatMember        *User            `json:"left_chat_member,omitempty"`
	NewChatTitle          *string          `json:"new_chat_title,omitempty"`
	NewChatPhoto          *[]PhotoSize     `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       *bool            `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      *bool            `json:"group_chat_created,omitempty"`
	SupergroupChatCreated *bool            `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    *bool            `json:"channel_chat_created,omitempty"`
	MigrateToChatID       *int64           `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     *int64           `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`
}

// GetCommand present a more flexible way to communicate with your bot.
func (m *Message) GetCommand() (*string, []string) {
	if m.Text != nil {
		if (*m.Text)[0] == '/' {
			ss := strings.Split(*m.Text, " ")
			command := ss[0][1:]
			return &command, ss[1:]
		}
	}

	return nil, nil
}

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type   string  `json:"type"`
	Offset int64   `json:"offset"`
	Length int64   `json:"length"`
	URL    *string `json:"url,omitempty"`
	User   *User   `json:"user,omitempty"`
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	FileSize *int64 `json:"file_size,omitempty"`
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID    string  `json:"file_id"`
	Duration  int64   `json:"duration"`
	Performer *string `json:"performer,omitempty"`
	Title     *string `json:"title,omitempty"`
	MimeType  *string `json:"mime_type,omitempty"`
	FileSize  *int64  `json:"file_size,omitempty"`
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName *string    `json:"file_name,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize *int64     `json:"file_size,omitempty"`
}

// Sticker represents a sticker.
type Sticker struct {
	FileID   string     `json:"file_id"`
	Width    int64      `json:"width"`
	Height   int64      `json:"height"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	Emoji    *string    `json:"emoji,omitempty"`
	FileSize *int64     `json:"file_size,omitempty"`
}

// Video represents a video file.
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int64      `json:"width"`
	Height   int64      `json:"height"`
	Duration int64      `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize *int64     `json:"file_size,omitempty"`
}

// Voice represents a voice note.
type Voice struct {
	FileID   string  `json:"file_id"`
	Duration int64   `json:"duration"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize *string `json:"file_size,omitempty"`
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int64  `json:"user_id,omitempty"`
}

// Location represents a point on the map.
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Venue represents a venue.
type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareID *string  `json:"foursquare_id,omitempty"`
}

// UserProfilePhotos represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int64         `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// File represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
type File struct {
	FileID   string  `json:"file_id"`
	FileSize *int64  `json:"file_size,omitempty"`
	FilePath *string `json:"file_path,omitempty"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  *bool              `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard *bool              `json:"one_time_keyboard,omitempty"`
	Selective       *bool              `json:"selective,omitempty"`
}

// KeyboardButton represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields are mutually exclusive.
type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  *bool  `json:"request_contact,omitempty"`
	RequestLocation *bool  `json:"request_location,omitempty"`
}

// ReplyKeyboardRemove ...
// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          *string       `json:"url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID *string  `json:"inline_message_id,omitempty"`
	ChatInstance    *string  `json:"chat_instance,omitempty"`
	Data            *string  `json:"data,omitempty"`
	GameShortName   *string  `json:"game_short_name,omitempty"`
}

// ForceReply ...
// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply bool  `json:"force_reply"`
	Selective  *bool `json:"selective,omitempty"`
}

// ChatMember contains information about one member of the chat.
type ChatMember struct {
	User   User   `json:"user"`
	Status string `json:"status"`
}

// ResponseParameters Contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatID *int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int64 `json:"retry_after,omitempty"`
}

// InputFile represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile interface {
	io.Reader
}

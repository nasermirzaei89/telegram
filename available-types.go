package telegram

import "io"

type MessageEntityType string

const (
	MessageEntityTypeMention       MessageEntityType = "mention"
	MessageEntityTypeHashtag       MessageEntityType = "hashtag"
	MessageEntityTypeCashtag       MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand    MessageEntityType = "bot_command"
	MessageEntityTypeURL           MessageEntityType = "url"
	MessageEntityTypeEmail         MessageEntityType = "email"
	MessageEntityTypePhoneNumber   MessageEntityType = "phone_number"
	MessageEntityTypeBold          MessageEntityType = "bold"
	MessageEntityTypeItalic        MessageEntityType = "italic"
	MessageEntityTypeUnderline     MessageEntityType = "underline"
	MessageEntityTypeStrikethrough MessageEntityType = "strikethrough"
	MessageEntityTypeCode          MessageEntityType = "code"
	MessageEntityTypePre           MessageEntityType = "pre"
	MessageEntityTypeTextLink      MessageEntityType = "text_link"
	MessageEntityTypeTextMention   MessageEntityType = "text_mention"
)

// MessageEntity struct.
type MessageEntity struct {
	Type     MessageEntityType `json:"type"`
	Offset   int               `json:"offset"`
	Length   int               `json:"length"`
	URL      *string           `json:"url,omitempty"`
	User     *User             `json:"user,omitempty"`
	Language string            `json:"language,omitempty"`
}

// PhotoSize struct.
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     *int   `json:"file_size,omitempty"`
}

// Animation struct.
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	Filename     *string    `json:"filename,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
}

// Audio struct.
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    *string    `json:"performer,omitempty"`
	Title        *string    `json:"title,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
}

// Document struct.
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
}

// Video struct.
type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
}

// Voice struct.
type Voice struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	Duration     int     `json:"duration"`
	MimeType     *string `json:"mime_type,omitempty"`
	FileSize     *int    `json:"file_size,omitempty"`
}

// VideoNote struct.
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbs       *PhotoSize `json:"thumbs,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
}

// Contact struct.
type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int64  `json:"user_id,omitempty"`
	VCard       *string `json:"vcard,omitempty"`
}

// Dice struct.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// PollOption struct.
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// PollAnswer struct.
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      User   `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

// Poll struct.
type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       *int            `json:"correct_option_id,omitempty"`
	Explanation           *string         `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            *int            `json:"open_period,omitempty"`
	CloseDate             *int            `json:"close_date,omitempty"`
}

// Location struct.
type Location struct {
	Longitude            float32  `json:"longitude"`
	Latitude             float32  `json:"latitude"`
	HorizontalAccuracy   *float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int     `json:"live_period,omitempty"`
	Heading              *int     `json:"heading,omitempty"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"`
}

// Venue struct.
type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareID    *string  `json:"foursquare_id,omitempty"`
	FoursquareType  *string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string  `json:"google_place_id,omitempty"`
	GooglePlaceType *string  `json:"google_place_type,omitempty"`
}

// ProximityAlertTriggered struct.
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

// MessageAutoDeleteTimerChanged struct.
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// VoiceChatScheduled represents a service message about a voice chat scheduled in the chat.
type VoiceChatScheduled struct {
	StartDate int `json:"start_date"` // Point in time (Unix timestamp) when the voice chat is supposed to be started by a chat administrator
}

// VoiceChatStarted struct.
type VoiceChatStarted struct{}

// VoiceChatEnded struct.
type VoiceChatEnded struct {
	Duration int `json:"duration"` // Voice chat duration; in seconds
}

// VoiceChatParticipantsInvited struct.
type VoiceChatParticipantsInvited struct {
	Users *[]User `json:"users,omitempty"`
}

// UserProfilePhotos struct.
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// File struct.
type File struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	FileSize     *int    `json:"file_size,omitempty"`
	FilePath     *string `json:"file_path,omitempty"`
}

// ReplyKeyboardMarkup struct.
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`                          // Array of button rows, each represented by an Array of KeyboardButton objects
	ResizeKeyboard        *bool              `json:"resize_keyboard,omitempty"`         // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard       *bool              `json:"one_time_keyboard,omitempty"`       // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	InputFieldPlaceholder *string            `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	Selective             *bool              `json:"selective,omitempty"`               // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
}

// KeyboardButton struct.
type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  *bool                   `json:"request_contact,omitempty"`
	RequestLocation *bool                   `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

// KeyboardButtonPollType struct.
type KeyboardButtonPollType struct {
	Type *string `json:"type,omitempty"`
}

// ReplyKeyboardRemove struct.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup struct.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton struct.
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

// LoginURL struct.
type LoginURL struct {
	URL                string  `json:"url"`
	ForwardText        *string `json:"forward_text,omitempty"`
	BotUsername        *string `json:"bot_username,omitempty"`
	RequestWriteAccess *bool   `json:"request_write_access,omitempty"`
}

// CallbackQuery struct.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID *string  `json:"inline_message_id,omitempty"`
	ChatInstance    *string  `json:"chat_instance,omitempty"`
	Data            *string  `json:"data,omitempty"`
	GameShortName   *string  `json:"game_short_name,omitempty"`
}

// ForceReply struct.
type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`                       // Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	Selective             *bool   `json:"selective,omitempty"`               // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// ChatPhoto struct.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// ChatInviteLink represents an invite link for a chat.
type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`
	Creator                 User    `json:"creator"`
	CreatesJoinRequest      User    `json:"creates_join_request"`
	IsPrimary               bool    `json:"is_primary"`
	IsRevoked               bool    `json:"is_revoked"`
	Name                    *string `json:"name"`
	ExpireDate              *int    `json:"expire_date,omitempty"`
	MemberLimit             *int    `json:"member_limit,omitempty"`
	PendingJoinRequestCount *int    `json:"pending_join_request_count,omitempty"`
}

// ChatMember contains information about one member of a chat. Currently, the following 6 types of chat members are supported:
// * ChatMemberOwner
// * ChatMemberAdministrator
// * ChatMemberMember
// * ChatMemberRestricted
// * ChatMemberLeft
// * ChatMemberBanned.
type ChatMember interface{}

// ChatMemberOwner represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	Status      string  `json:"status"`
	User        User    `json:"user"`
	IsAnonymous bool    `json:"is_anonymous,omitempty"`
	CustomTitle *string `json:"custom_title,omitempty"`
}

// ChatMemberAdministrator represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	Status              string  `json:"status"`
	User                User    `json:"user"`
	CanBeEdited         bool    `json:"can_be_edited,omitempty"`
	IsAnonymous         bool    `json:"is_anonymous,omitempty"`
	CanManageChat       bool    `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   bool    `json:"can_delete_messages,omitempty"`
	CanManageVoiceChats bool    `json:"can_manage_voice_chats,omitempty"`
	CanRestrictMembers  bool    `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool    `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool    `json:"can_change_info,omitempty"`
	CanInviteUsers      bool    `json:"can_invite_users,omitempty"`
	CanPostMessages     *bool   `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool   `json:"can_pin_messages,omitempty"`
	CustomTitle         *string `json:"custom_title,omitempty"`
}

// ChatMemberMember represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

// ChatMemberRestricted represents a chat member that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
	UntilDate             int    `json:"until_date,omitempty"`
}

// ChatMemberLeft represents a chat member that isn't currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

// ChatMemberBanned represents a chat member that was banned in the chat and can't return to the chat or view chat messages.
type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate int    `json:"until_date,omitempty"`
}

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat          Chat            `json:"chat"`
	From          User            `json:"from"`
	Date          int             `json:"date"`
	OldChatMember ChatMember      `json:"old_chat_member"`
	NewChatMember ChatMember      `json:"new_chat_member"`
	InviteLink    *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatJoinRequest represents a join request sent to a chat.
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
	Date       int             `json:"date"`
	Bio        *string         `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatPermissions struct.
type ChatPermissions struct {
	CanSendMessages       *bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  *bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
}

// ChatLocation struct.
type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

// BotCommand struct.
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// BotCommandScope represents the scope to which bot commands are applied. Currently, the following 7 scopes are supported:
// * BotCommandScopeDefault
// * BotCommandScopeAllPrivateChats
// * BotCommandScopeAllGroupChats
// * BotCommandScopeAllChatAdministrators
// * BotCommandScopeChat
// * BotCommandScopeChatAdministrators
// * BotCommandScopeChatMember.
type BotCommandScope interface{}

// BotCommandScopeDefault represents the default scope of bot commands. Default commands are used if no commands with a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
	Type string `json:"type"` // Scope type, must be default
}

// BotCommandScopeAllPrivateChats represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"` // Scope type, must be all_private_chats
}

// BotCommandScopeAllGroupChats represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"` // Scope type, must be all_group_chats
}

// BotCommandScopeAllChatAdministrators represents the scope of bot commands, covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"` // Scope type, must be all_chat_administrators
}

// BotCommandScopeChat represents the scope of bot commands, covering specific chat.
type BotCommandScopeChat struct {
	Type   string      `json:"type"`    // Scope type, must be chat
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// BotCommandScopeChatAdministrators represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	Type   string      `json:"type"`    // Scope type, must be chat_administrators
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// BotCommandScopeChatMember represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	Type   string      `json:"type"`    // Scope type, must be chat_member
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID int         `json:"user_id"` // Unique identifier of the target user
}

// ResponseParameters struct.
type ResponseParameters struct {
	MigrateToChatID *int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int   `json:"retry_after,omitempty"`
}

// InputMedia interface.
type InputMedia interface{}

// InputMediaPhoto struct.
type InputMediaPhoto struct {
	Type      string  `json:"type"`
	Media     string  `json:"media"`
	Caption   *string `json:"caption,omitempty"`
	ParseMode *string `json:"parse_mode,omitempty"`
}

// InputMediaVideo struct.
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

// InputMediaAnimation struct.
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

// InputMediaAudio struct.
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

// InputMediaDocument struct.
type InputMediaDocument struct {
	Type                        string      `json:"type"`
	Media                       string      `json:"media"`
	Thumb                       interface{} `json:"thumb,omitempty"`
	Caption                     *string     `json:"caption,omitempty"`
	ParseMode                   *string     `json:"parse_mode,omitempty"`
	DisableContentTypeDetection *bool       `json:"disable_content_type_detection,omitempty"`
}

// InputFile interface.
type InputFile interface {
	io.Reader
}

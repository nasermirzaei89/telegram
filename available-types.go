package telegram

import "io"

// User struct.
type User struct {
	ID                      int     `json:"id"`
	IsBot                   bool    `json:"is_bot"`
	FirstName               string  `json:"first_name"`
	LastName                *string `json:"last_name,omitempty"`
	Username                *string `json:"username,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`
}

// Chat struct.
type Chat struct {
	ID               int              `json:"id"`
	Type             ChatType         `json:"type"`
	Title            *string          `json:"title,omitempty"`
	Username         *string          `json:"username,omitempty"`
	FirstName        *string          `json:"first_name,omitempty"`
	LastName         *string          `json:"last_name,omitempty"`
	Photo            *ChatPhoto       `json:"photo,omitempty"`
	Bio              *string          `json:"bio,omitempty"`
	Description      *string          `json:"description,omitempty"`
	InviteLink       *string          `json:"invite_link,omitempty"`
	PinnedMessage    *Message         `json:"pinned_message,omitempty"`
	Permissions      *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay    *int             `json:"slow_mode_delay,omitempty"`
	StickerSetName   *string          `json:"sticker_set_name,omitempty"`
	CanSetStickerSet *bool            `json:"can_set_sticker_set,omitempty"`
	LinkedChatID     *int             `json:"linked_chat_id,omitempty"`
	Location         *ChatLocation    `json:"location,omitempty"`
}

// Message struct.
type Message struct {
	MessageID                     int                            `json:"message_id"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	Date                          int                            `json:"date"`
	Chat                          Chat                           `json:"chat"`
	ForwardFrom                   *User                          `json:"forward_from,omitempty"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID          *int                           `json:"forward_from_message_id,omitempty"`
	ForwardSignature              *string                        `json:"forward_signature,omitempty"`
	ForwardSenderName             *string                        `json:"forward_sender_name,omitempty"`
	ForwardDate                   *int                           `json:"forward_date,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      *int                           `json:"edit_date,omitempty"`
	MediaGroupID                  *string                        `json:"media_group_id,omitempty"`
	AuthorSignature               *string                        `json:"author_signature,omitempty"`
	Text                          *string                        `json:"text,omitempty"`
	Entities                      []MessageEntity                `json:"entities,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	Photo                         []PhotoSize                    `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       *string                        `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  *string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               *bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              *bool                          `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         *bool                          `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            *bool                          `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               *int                           `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             *int                           `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	ConnectedWebsite              *string                        `json:"connected_website,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	VoiceChatScheduled            *VoiceChatScheduled            `json:"voice_chat_scheduled,omitempty"`
	VoiceChatStarted              *VoiceChatStarted              `json:"voice_chat_started,omitempty"`
	VoiceChatEnded                *VoiceChatEnded                `json:"voice_chat_ended,omitempty"`
	VoiceChatParticipantsInvited  *VoiceChatParticipantsInvited  `json:"voice_chat_participants_invited,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

// MessageID struct.
type MessageID struct {
	MessageID int `json:"message_id"`
}

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

// Document struct
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
}

// Video struct
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

// Voice struct
type Voice struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	Duration     int     `json:"duration"`
	MimeType     *string `json:"mime_type,omitempty"`
	FileSize     *int    `json:"file_size,omitempty"`
}

// VideoNote struct
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbs       *PhotoSize `json:"thumbs,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
}

// Contact struct
type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int    `json:"user_id,omitempty"`
	VCard       *string `json:"vcard,omitempty"`
}

// Dice struct
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// PollOption struct
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// PollAnswer struct
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      User   `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

// Poll struct
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

// Location struct
type Location struct {
	Longitude            float32  `json:"longitude"`
	Latitude             float32  `json:"latitude"`
	HorizontalAccuracy   *float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int     `json:"live_period,omitempty"`
	Heading              *int     `json:"heading,omitempty"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"`
}

// Venue struct
type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareID    *string  `json:"foursquare_id,omitempty"`
	FoursquareType  *string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string  `json:"google_place_id,omitempty"`
	GooglePlaceType *string  `json:"google_place_type,omitempty"`
}

// ProximityAlertTriggered struct
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

// MessageAutoDeleteTimerChanged struct
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// VoiceChatScheduled represents a service message about a voice chat scheduled in the chat.
type VoiceChatScheduled struct {
	StartDate int `json:"start_date"` // Point in time (Unix timestamp) when the voice chat is supposed to be started by a chat administrator
}

// VoiceChatStarted struct
type VoiceChatStarted struct{}

// VoiceChatEnded struct
type VoiceChatEnded struct {
	Duration int `json:"duration"` // Voice chat duration; in seconds
}

// VoiceChatParticipantsInvited struct
type VoiceChatParticipantsInvited struct {
	Users *[]User `json:"users,omitempty"`
}

// UserProfilePhotos struct
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// File struct
type File struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	FileSize     *int    `json:"file_size,omitempty"`
	FilePath     *string `json:"file_path,omitempty"`
}

// ReplyKeyboardMarkup struct
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  *bool              `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard *bool              `json:"one_time_keyboard,omitempty"`
	Selective       *bool              `json:"selective,omitempty"`
}

// KeyboardButton struct
type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  *bool                   `json:"request_contact,omitempty"`
	RequestLocation *bool                   `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

// KeyboardButtonPollType struct
type KeyboardButtonPollType struct {
	Type *string `json:"type,omitempty"`
}

// ReplyKeyboardRemove struct
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup struct
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton struct
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

// LoginURL struct
type LoginURL struct {
	URL                string  `json:"url"`
	ForwardText        *string `json:"forward_text,omitempty"`
	BotUsername        *string `json:"bot_username,omitempty"`
	RequestWriteAccess *bool   `json:"request_write_access,omitempty"`
}

// CallbackQuery struct
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID *string  `json:"inline_message_id,omitempty"`
	ChatInstance    *string  `json:"chat_instance,omitempty"`
	Data            *string  `json:"data,omitempty"`
	GameShortName   *string  `json:"game_short_name,omitempty"`
}

// ForceReply struct
type ForceReply struct {
	ForceReply bool  `json:"force_reply"`
	Selective  *bool `json:"selective,omitempty"`
}

// ChatPhoto struct
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// ChatInviteLink struct
type ChatInviteLink struct {
	InviteLink  string `json:"invite_link"`
	Creator     User   `json:"creator"`
	IsPrimary   bool   `json:"is_primary"`
	IsRevoked   bool   `json:"is_revoked"`
	ExpireDate  *int   `json:"expire_date,omitempty"`
	MemberLimit *int   `json:"member_limit,omitempty"`
}

// ChatMember struct
type ChatMember struct {
	User                  User    `json:"user"`
	Status                string  `json:"status"`
	CustomTitle           *string `json:"custom_title,omitempty"`
	IsAnonymous           *bool   `json:"is_anonymous,omitempty"`
	UntilDate             *int    `json:"until_date,omitempty"`
	CanBeEdited           *bool   `json:"can_be_edited,omitempty"`
	CanManageChat         *bool   `json:"can_manage_chat,omitempty"`
	CanPostMessages       *bool   `json:"can_post_messages,omitempty"`
	CanEditMessages       *bool   `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     *bool   `json:"can_delete_messages,omitempty"`
	CanManageVoiceChats   *bool   `json:"can_manage_voice_chats,omitempty"`
	CanRestrictMembers    *bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     *bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo         *bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool   `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool   `json:"can_pin_messages,omitempty"`
	IsMember              *bool   `json:"is_member,omitempty"`
	CanSendMessages       *bool   `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  *bool   `json:"can_send_media_messages,omitempty"`
	CanSendPolls          *bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool   `json:"can_add_web_page_previews,omitempty"`
}

// ChatMemberUpdated struct
type ChatMemberUpdated struct {
	Chat          Chat            `json:"chat"`
	From          User            `json:"from"`
	Date          int             `json:"date"`
	OldChatMember ChatMember      `json:"old_chat_member"`
	NewChatMember ChatMember      `json:"new_chat_member"`
	InviteLink    *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatPermissions struct
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

// ChatLocation struct
type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

// BotCommand struct
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// ResponseParameters struct
type ResponseParameters struct {
	MigrateToChatID *int `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int `json:"retry_after,omitempty"`
}

// InputMedia interface
type InputMedia interface{}

// InputMediaPhoto struct
type InputMediaPhoto struct {
	Type      string  `json:"type"`
	Media     string  `json:"media"`
	Caption   *string `json:"caption,omitempty"`
	ParseMode *string `json:"parse_mode,omitempty"`
}

// InputMediaVideo struct
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

// InputMediaAnimation struct
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

// InputMediaAudio struct
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

// InputMediaDocument struct
type InputMediaDocument struct {
	Type                        string      `json:"type"`
	Media                       string      `json:"media"`
	Thumb                       interface{} `json:"thumb,omitempty"`
	Caption                     *string     `json:"caption,omitempty"`
	ParseMode                   *string     `json:"parse_mode,omitempty"`
	DisableContentTypeDetection *bool       `json:"disable_content_type_detection,omitempty"`
}

// InputFile interface
type InputFile interface {
	io.Reader
}

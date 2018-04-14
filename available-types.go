package telegram

import (
	"io"
	"strings"
)

// All types used in the Bot API responses are represented as JSON-objects.
//
// It is safe to use 32-bit signed integers for storing all Integer fields unless otherwise noted.
//
// Optional fields may be not returned when irrelevant.

// User represents a Telegram user or bot.
type User struct {
	ID           int32   `json:"id"`                      // Unique identifier for this user or bot
	IsBot        bool    `json:"is_bot"`                  // True, if this user is a bot
	FirstName    string  `json:"first_name"`              // User‘s or bot’s first name
	LastName     *string `json:"last_name,omitempty"`     // Optional. User‘s or bot’s last name
	Username     *string `json:"username,omitempty"`      // Optional. User‘s or bot’s username
	LanguageCode *string `json:"language_code,omitempty"` // Optional. IETF language tag of the user's language
}

// Chat represents a chat.
type Chat struct {
	ID                          int64      `json:"id"`                                       // Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	Type                        string     `json:"type"`                                     // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title                       *string    `json:"title,omitempty"`                          // Optional. Title, for supergroups, channels and group chats
	Username                    *string    `json:"username,omitempty"`                       // Optional. Username, for private chats, supergroups and channels if available
	FirstName                   *string    `json:"first_name,omitempty"`                     // Optional. First name of the other party in a private chat
	LastName                    *string    `json:"last_name,omitempty"`                      // Optional. Last name of the other party in a private chat
	AllMembersAreAdministrators *bool      `json:"all_members_are_administrators,omitempty"` // Optional. True if a group has ‘All Members Are Admins’ enabled.
	Photo                       *ChatPhoto `json:"photo,omitempty"`                          // Optional. Chat photo. Returned only in getChat.
	Description                 *string    `json:"description,omitempty"`                    // Optional. Description, for supergroups and channel chats. Returned only in getChat.
	InviteLink                  *string    `json:"invite_link,omitempty"`                    // Optional. Chat invite link, for supergroups and channel chats. Returned only in getChat.
	PinnedMessage               *Message   `json:"pinned_message,omitempty"`                 // Optional. Pinned message, for supergroups and channel chats. Returned only in getChat.
	StickerSetName              *string    `json:"sticker_set_name,omitempty"`               // Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet            *bool      `json:"can_set_sticker_set,omitempty"`            // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
}

// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
// TODO: Move this
const (
	ChatTypePrivate    = "private"
	ChatTypeGroup      = "group"
	ChatTypeSuperGroup = "supergroup"
	ChatTypeChannel    = "channel"
)

// Message represents a message.
type Message struct {
	MessageID             int32              `json:"message_id"`                        // Unique message identifier inside this chat
	From                  *User              `json:"from,omitempty"`                    // Optional. Sender, empty for messages sent to channels
	Date                  int32              `json:"date"`                              // Date the message was sent in Unix time
	Chat                  Chat               `json:"chat"`                              // Conversation the message belongs to
	ForwardFrom           *User              `json:"forward_from,omitempty"`            // Optional. For forwarded messages, sender of the original message
	ForwardFromChat       *Chat              `json:"forward_from_chat,omitempty"`       // Optional. For messages forwarded from channels, information about the original channel
	ForwardFromMessageID  *int32             `json:"forward_from_message_id,omitempty"` // Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardSignature      *string            `json:"forward_signature,omitempty"`       // Optional. For messages forwarded from channels, signature of the post author if present
	ForwardDate           *int32             `json:"forward_date,omitempty"`            // Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage        *Message           `json:"reply_to_message,omitempty"`        // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	EditDate              *int32             `json:"edit_date,omitempty"`               // Optional. Date the message was last edited in Unix time
	MediaGroupID          *string            `json:"media_group_id,omitempty"`          // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature       *string            `json:"author_signature,omitempty"`        // Optional. Signature of the post author for messages in channels
	Text                  *string            `json:"text,omitempty"`                    // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Entities              *[]MessageEntity   `json:"entities,omitempty"`                // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	CaptionEntities       *[]MessageEntity   `json:"caption_entities,omitempty"`        // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	Audio                 *Audio             `json:"audio,omitempty"`                   // Optional. Message is an audio file, information about the file
	Document              *Document          `json:"document,omitempty"`                // Optional. Message is a general file, information about the file
	Game                  *Game              `json:"game,omitempty"`                    // Optional. Message is a game, information about the game
	Photo                 *[]PhotoSize       `json:"photo,omitempty"`                   // Optional. Message is a photo, available sizes of the photo
	Sticker               *Sticker           `json:"sticker,omitempty"`                 // Optional. Message is a sticker, information about the sticker
	Video                 *Video             `json:"video,omitempty"`                   // Optional. Message is a video, information about the video
	Voice                 *Voice             `json:"voice,omitempty"`                   // Optional. Message is a voice message, information about the file
	VideoNote             *VideoNote         `json:"video_note,omitempty"`              // Optional. Message is a video note, information about the video message
	Caption               *string            `json:"caption,omitempty"`                 // Optional. Caption for the audio, document, photo, video or voice, 0-200 characters
	Contact               *Contact           `json:"contact,omitempty"`                 // Optional. Message is a shared contact, information about the contact
	Location              *Location          `json:"location,omitempty"`                // Optional. Message is a shared location, information about the location
	Venue                 *Venue             `json:"venue,omitempty"`                   // Optional. Message is a venue, information about the venue
	NewChatMembers        *[]User            `json:"new_chat_members,omitempty"`        // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember        *User              `json:"left_chat_member,omitempty"`        // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle          *string            `json:"new_chat_title,omitempty"`          // Optional. A chat title was changed to this value
	NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo,omitempty"`          // Optional. A chat photo was change to this value
	DeleteChatPhoto       *bool              `json:"delete_chat_photo,omitempty"`       // Optional. Service message: the chat photo was deleted
	GroupChatCreated      *bool              `json:"group_chat_created,omitempty"`      // Optional. Service message: the group has been created
	SupergroupChatCreated *bool              `json:"supergroup_chat_created,omitempty"` // Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated    *bool              `json:"channel_chat_created,omitempty"`    // Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MigrateToChatID       *int64             `json:"migrate_to_chat_id,omitempty"`      // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatID     *int64             `json:"migrate_from_chat_id,omitempty"`    // Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	PinnedMessage         *Message           `json:"pinned_message,omitempty"`          // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	Invoice               *Invoice           `json:"invoice,omitempty"`                 // Optional. Message is an invoice for a payment, information about the invoice
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment,omitempty"`      // Optional. Message is a service message about a successful payment, information about the payment
	ConnectedWebsite      *string            `json:"connected_website,omitempty"`       // Optional. The domain name of the website on which the user has logged in
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
	Type   string  `json:"type"`           // Type of the entity. Can be mention (@username), hashtag, bot_command, url, email, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
	Offset int32   `json:"offset"`         // Offset in UTF-16 code units to the start of the entity
	Length int32   `json:"length"`         // Length of the entity in UTF-16 code units
	URL    *string `json:"url,omitempty"`  // Optional. For “text_link” only, url that will be opened after user taps on the text
	User   *User   `json:"user,omitempty"` // Optional. For “text_mention” only, the mentioned user
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID   string `json:"file_id"`             // Unique identifier for this file
	Width    int32  `json:"width"`               // Photo width
	Height   int32  `json:"height"`              // Photo height
	FileSize *int32 `json:"file_size,omitempty"` // Optional. File size
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID    string  `json:"file_id"`             // Unique identifier for this file
	Duration  int32   `json:"duration"`            // Duration of the audio in seconds as defined by sender
	Performer *string `json:"performer,omitempty"` // Optional. Performer of the audio as defined by sender or by audio tags
	Title     *string `json:"title,omitempty"`     // Optional. Title of the audio as defined by sender or by audio tags
	MimeType  *string `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize  *int32  `json:"file_size,omitempty"` // Optional. File size
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID   string     `json:"file_id"`             // Unique file identifier
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Document thumbnail as defined by sender
	FileName *string    `json:"file_name,omitempty"` // Optional. Original filename as defined by sender
	MimeType *string    `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize *int32     `json:"file_size,omitempty"` // Optional. File size
}

// Video represents a video file.
type Video struct {
	FileID   string     `json:"file_id"`             // Unique identifier for this file
	Width    int32      `json:"width"`               // Video width as defined by sender
	Height   int32      `json:"height"`              // Video height as defined by sender
	Duration int32      `json:"duration"`            // Duration of the video in seconds as defined by sender
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Video thumbnail
	MimeType *string    `json:"mime_type,omitempty"` // Optional. Mime type of a file as defined by sender
	FileSize *int32     `json:"file_size,omitempty"` // Optional. File size
}

// Voice represents a voice note.
type Voice struct {
	FileID   string  `json:"file_id"`             // Unique identifier for this file
	Duration int32   `json:"duration"`            // Duration of the audio in seconds as defined by sender
	MimeType *string `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize *int32  `json:"file_size,omitempty"` // Optional. File size
}

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileID   string     `json:"file_id"`             // Unique identifier for this file
	Length   int32      `json:"length"`              // Video width and height as defined by sender
	Duration int32      `json:"duration"`            // Duration of the video in seconds as defined by sender
	Thumbs   *PhotoSize `json:"thumbs,omitempty"`    // Optional. Video thumbnail
	FileSize *int32     `json:"file_size,omitempty"` // Optional. File size
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string  `json:"phone_number"`        // Contact's phone number
	FirstName   string  `json:"first_name"`          // Contact's first name
	LastName    *string `json:"last_name,omitempty"` // Optional. Contact's last name
	UserID      *int32  `json:"user_id,omitempty"`   // Optional. Contact's user identifier in Telegram
}

// Location represents a point on the map.
type Location struct {
	Longitude float32 `json:"longitude"` // Longitude as defined by sender
	Latitude  float32 `json:"latitude"`  // Latitude as defined by sender
}

// Venue represents a venue.
type Venue struct {
	Location     Location `json:"location"`                // Venue location
	Title        string   `json:"title"`                   // Name of the venue
	Address      string   `json:"address"`                 // Address of the venue
	FoursquareID *string  `json:"foursquare_id,omitempty"` // Optional. Foursquare identifier of the venue
}

// UserProfilePhotos represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int32         `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}

// File represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<Token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
// Maximum file size to download is 20 MB
type File struct {
	FileID   string  `json:"file_id"`             // Unique identifier for this file
	FileSize *int32  `json:"file_size,omitempty"` // Optional. File size, if known
	FilePath *string `json:"file_path,omitempty"` // Optional. File path. Use https://api.telegram.org/file/bot<Token>/<file_path> to get the file.
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`                    // Array of button rows, each represented by an Array of KeyboardButton objects
	ResizeKeyboard  *bool              `json:"resize_keyboard,omitempty"`   // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard *bool              `json:"one_time_keyboard,omitempty"` // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	Selective       *bool              `json:"selective,omitempty"`         // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// KeyboardButton represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields are mutually exclusive.
// Note: request_contact and request_location options will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type KeyboardButton struct {
	Text            string `json:"text"`                       // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestContact  *bool  `json:"request_contact,omitempty"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestLocation *bool  `json:"request_location,omitempty"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
}

// ReplyKeyboardRemove ...
// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`     // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      *bool `json:"selective,omitempty"` // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	// Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will display unsupported message.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

// InlineKeyboardButton represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text              string  `json:"text"`                          // Label text on the button
	URL               *string `json:"url,omitempty"`                 // Optional. HTTP url to be opened when button is pressed
	CallbackData      *string `json:"callback_data,omitempty"`       // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"` // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.
	// Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"` // Optional. If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.
	// This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"` // Optional. Description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.
	Pay *bool `json:"pay,omitempty"` // Optional. Specify True, to send a Pay button.
	// NOTE: This type of button must always be the first button in the first row.
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
// NOTE: After the user presses a callback button, Telegram clients will display a progress bar until you call answerCallbackQuery. It is, therefore, necessary to react by calling answerCallbackQuery even if no notification to the user is needed (e.g., without specifying any of the optional parameters).
type CallbackQuery struct {
	ID              string   `json:"id"`                          // Unique identifier for this query
	From            User     `json:"from"`                        // Sender
	Message         *Message `json:"message,omitempty"`           // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	InlineMessageID *string  `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    *string  `json:"chat_instance,omitempty"`     // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            *string  `json:"data,omitempty"`              // Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
	GameShortName   *string  `json:"game_short_name,omitempty"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

// ForceReply ...
// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
//
// Example: A poll bot for groups runs in privacy mode (only receives commands, replies to its messages and mentions). There could be two ways to create a new poll:
//
// Explain the user how to send a command with parameters (e.g. /newpoll question answer1 answer2). May be appealing for hardcore users but lacks modern day polish.
// Guide the user through a step-by-step process. ‘Please send me your question’, ‘Cool, now let’s add the first answer option‘, ’Great. Keep adding answer options, then send /done when you‘re ready’.
// The last option is definitely more attractive. And if you use ForceReply in your bot‘s questions, it will receive the user’s answers even if it only receives replies, commands and mentions — without any extra work for the user.
type ForceReply struct {
	ForceReply bool  `json:"force_reply"`         // Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	Selective  *bool `json:"selective,omitempty"` // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"` // Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download.
	BigFileID   string `json:"big_file_id"`   // Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download.
}

// ChatMember contains information about one member of the chat.
type ChatMember struct {
	User                  User   `json:"user"`                                // Information about the user
	Status                string `json:"status"`                              // The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
	UntilDate             *int32 `json:"until_date,omitempty"`                // Optional. Restricted and kicked only. Date when restrictions will be lifted for this user, unix time
	CanBeEdited           *bool  `json:"can_be_edited,omitempty"`             // Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
	CanChangeInfo         *bool  `json:"can_change_info,omitempty"`           // Optional. Administrators only. True, if the administrator can change the chat title, photo and other settings
	CanPostMessages       *bool  `json:"can_post_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can post in the channel, channels only
	CanEditMessages       *bool  `json:"can_edit_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMessages     *bool  `json:"can_delete_messages,omitempty"`       // Optional. Administrators only. True, if the administrator can delete messages of other users
	CanInviteUsers        *bool  `json:"can_invite_users,omitempty"`          // Optional. Administrators only. True, if the administrator can invite new users to the chat
	CanRestrictMembers    *bool  `json:"can_restrict_members,omitempty"`      // Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
	CanPinMessages        *bool  `json:"can_pin_messages,omitempty"`          // Optional. Administrators only. True, if the administrator can pin messages, supergroups only
	CanPromoteMembers     *bool  `json:"can_promote_members,omitempty"`       // Optional. Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanSendMessages       *bool  `json:"can_send_messages,omitempty"`         // Optional. Restricted only. True, if the user can send text messages, contacts, locations and venues
	CanSendMediaMessages  *bool  `json:"can_send_media_messages,omitempty"`   // Optional. Restricted only. True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendOtherMessages  *bool  `json:"can_send_other_messages,omitempty"`   // Optional. Restricted only. True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePreviews *bool  `json:"can_add_web_page_previews,omitempty"` // Optional. Restricted only. True, if user may add web page previews to his messages, implies can_send_media_messages
}

// ResponseParameters Contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatID *int32 `json:"migrate_to_chat_id,omitempty"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	RetryAfter      *int32 `json:"retry_after,omitempty"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

// InputMedia represents the content of a media message to be sent. It should be one of
// InputMediaPhoto
// InputMediaVideo
type InputMedia interface {
}

// InputMediaPhoto represents a photo to be sent.
type InputMediaPhoto struct {
	Type      string  `json:"type"`                 // Type of the result, must be photo
	Media     string  `json:"media"`                // File to send.Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name.More info on Sending Files »
	Caption   *string `json:"caption,omitempty"`    // Optional.Caption of the photo to be sent, 0-200 characters
	ParseMode *string `json:"parse_mode,omitempty"` // Optional.Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
}

// InputMediaVideo represents a video to be sent.
type InputMediaVideo struct {
	Type              string  `json:"type"`                         // Type of the result, must be video
	Media             string  `json:"media"`                        // File to send.Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name.More info on Sending Files »
	Caption           *string `json:"caption,omitempty"`            // Optional.Caption of the video to be sent, 0-200 characters
	ParseMode         *string `json:"parse_mode,omitempty"`         // Optional.Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Width             *int32  `json:"width,omitempty"`              // Optional.Video width
	Height            *int32  `json:"height,omitempty"`             // Optional.Video height
	Duration          *int32  `json:"duration,omitempty"`           // Optional.Video duration
	SupportsStreaming *bool   `json:"supports_streaming,omitempty"` // Optional.Pass True, if the uploaded video is suitable for streaming
}

// InputFile represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
// Sending files
// There are three ways to send files (photos, stickers, audio, media, etc.):
//
// If the file is already stored somewhere on the Telegram servers, you don't need to reupload it: each file object has a file_id field, simply pass this file_id as a parameter instead of uploading. There are no limits for files sent this way.
// Provide Telegram with an HTTP URL for the file to be sent. Telegram will download and send the file. 5 MB max size for photos and 20 MB max for other types of content.
// Post the file using multipart/form-data in the usual way that files are uploaded via the browser. 10 MB max size for photos, 50 MB for other files.
// Sending by file_id
//
// It is not possible to change the file type when resending by file_id. I.e. a video can't be sent as a photo, a photo can't be sent as a document, etc.
// It is not possible to resend thumbnails.
// Resending a photo by file_id will send all of its sizes.
// file_id is unique for each individual bot and can't be transferred from one bot to another.
// Sending by URL
//
// When sending by URL the target file must have the correct MIME type (e.g., audio/mpeg for sendAudio, etc.).
// In sendDocument, sending by URL will currently only work for gif, pdf and zip files.
// To use sendVoice, the file must have the type audio/ogg and be no more than 1MB in size. 1–20MB voice notes will be sent as files.
// Other configurations may work but we can't guarantee that they will.
type InputFile interface {
	io.Reader
}

// Inline mode objects
// Objects and methods used in the inline mode are described in the Inline mode section.

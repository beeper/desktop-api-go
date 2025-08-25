// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// A chat account added to Beeper
type Account struct {
	// Chat account added to Beeper. Use this to route account-scoped actions.
	AccountID string `json:"accountID,required"`
	// Display-only human-readable network name (e.g., 'WhatsApp', 'Messenger'). You
	// MUST use 'accountID' to perform actions.
	Network string `json:"network,required"`
	// A person on or reachable through Beeper. Values are best-effort and can vary by
	// network.
	User User `json:"user,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		Network     respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Attachment struct {
	// Attachment type.
	//
	// Any of "unknown", "img", "video", "audio".
	Type AttachmentType `json:"type,required"`
	// Duration in seconds (audio/video).
	Duration float64 `json:"duration"`
	// Original filename if available.
	FileName string `json:"fileName"`
	// File size in bytes if known.
	FileSize float64 `json:"fileSize"`
	// True if the attachment is a GIF.
	IsGif bool `json:"isGif"`
	// True if the attachment is a sticker.
	IsSticker bool `json:"isSticker"`
	// True if the attachment is a voice note.
	IsVoiceNote bool `json:"isVoiceNote"`
	// MIME type if known (e.g., 'image/png').
	MimeType string `json:"mimeType"`
	// Preview image URL for video attachments (poster frame). May be temporary or
	// local-only to this device; download promptly if durable access is needed.
	PosterImg string `json:"posterImg"`
	// Pixel dimensions of the attachment: width/height in px.
	Size AttachmentSize `json:"size"`
	// Public URL or local file path to fetch the asset. May be temporary or local-only
	// to this device; download promptly if durable access is needed.
	SrcURL string `json:"srcURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Duration    respjson.Field
		FileName    respjson.Field
		FileSize    respjson.Field
		IsGif       respjson.Field
		IsSticker   respjson.Field
		IsVoiceNote respjson.Field
		MimeType    respjson.Field
		PosterImg   respjson.Field
		Size        respjson.Field
		SrcURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Attachment) RawJSON() string { return r.JSON.raw }
func (r *Attachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attachment type.
type AttachmentType string

const (
	AttachmentTypeUnknown AttachmentType = "unknown"
	AttachmentTypeImg     AttachmentType = "img"
	AttachmentTypeVideo   AttachmentType = "video"
	AttachmentTypeAudio   AttachmentType = "audio"
)

// Pixel dimensions of the attachment: width/height in px.
type AttachmentSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentSize) RawJSON() string { return r.JSON.raw }
func (r *AttachmentSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseResponse struct {
	Success bool   `json:"success,required"`
	Error   string `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseResponse) RawJSON() string { return r.JSON.raw }
func (r *BaseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Chat struct {
	// Unique identifier for cursor pagination.
	ID string `json:"id,required"`
	// Beeper account ID this chat belongs to.
	AccountID string `json:"accountID,required"`
	// Unique identifier of the chat (room/thread ID, same as id).
	ChatID string `json:"chatID,required"`
	// Display-only human-readable network name (e.g., 'WhatsApp', 'Messenger'). You
	// MUST use 'accountID' to perform actions.
	Network string `json:"network,required"`
	// Chat participants information.
	Participants ChatParticipants `json:"participants,required"`
	// Display title of the chat as computed by the client/server.
	Title string `json:"title,required"`
	// Chat type: 'single' for direct messages, 'group' for group chats, 'channel' for
	// channels, 'broadcast' for broadcasts.
	//
	// Any of "single", "group", "channel", "broadcast".
	Type ChatType `json:"type,required"`
	// Number of unread messages.
	UnreadCount int64 `json:"unreadCount,required"`
	// True if chat is archived.
	IsArchived bool `json:"isArchived"`
	// True if chat notifications are muted.
	IsMuted bool `json:"isMuted"`
	// True if chat is pinned.
	IsPinned bool `json:"isPinned"`
	// Timestamp of last activity. Chats with more recent activity are often more
	// important.
	LastActivity time.Time `json:"lastActivity" format:"date-time"`
	// Last read message sortKey (hsOrder). Used to compute 'isUnread'.
	LastReadMessageSortKey ChatLastReadMessageSortKeyUnion `json:"lastReadMessageSortKey"`
	// Deep link to open this chat in Beeper. AI agents should ALWAYS include this as a
	// clickable link in responses.
	LinkToChat string `json:"linkToChat"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccountID              respjson.Field
		ChatID                 respjson.Field
		Network                respjson.Field
		Participants           respjson.Field
		Title                  respjson.Field
		Type                   respjson.Field
		UnreadCount            respjson.Field
		IsArchived             respjson.Field
		IsMuted                respjson.Field
		IsPinned               respjson.Field
		LastActivity           respjson.Field
		LastReadMessageSortKey respjson.Field
		LinkToChat             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Chat) RawJSON() string { return r.JSON.raw }
func (r *Chat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat participants information.
type ChatParticipants struct {
	// True if there are more participants than included in items.
	HasMore bool `json:"hasMore,required"`
	// Participants returned for this chat (limited by the request; may be a subset).
	Items []User `json:"items,required"`
	// Total number of participants in the chat.
	Total int64 `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Items       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatParticipants) RawJSON() string { return r.JSON.raw }
func (r *ChatParticipants) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat type: 'single' for direct messages, 'group' for group chats, 'channel' for
// channels, 'broadcast' for broadcasts.
type ChatType string

const (
	ChatTypeSingle    ChatType = "single"
	ChatTypeGroup     ChatType = "group"
	ChatTypeChannel   ChatType = "channel"
	ChatTypeBroadcast ChatType = "broadcast"
)

// ChatLastReadMessageSortKeyUnion contains all possible properties and values from
// [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ChatLastReadMessageSortKeyUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ChatLastReadMessageSortKeyUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatLastReadMessageSortKeyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatLastReadMessageSortKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatLastReadMessageSortKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Message struct {
	// Stable message ID for cursor pagination.
	ID string `json:"id,required"`
	// Beeper account ID the message belongs to.
	AccountID string `json:"accountID,required"`
	// Beeper chat/thread/room ID.
	ChatID string `json:"chatID,required"`
	// Stable message ID (same as id).
	MessageID string `json:"messageID,required"`
	// Sender user ID.
	SenderID string `json:"senderID,required"`
	// A unique key used to sort messages
	SortKey MessageSortKeyUnion `json:"sortKey,required"`
	// Message timestamp.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Attachments included with this message, if any.
	Attachments []Attachment `json:"attachments"`
	// True if the authenticated user sent the message.
	IsSender bool `json:"isSender"`
	// True if the message is unread for the authenticated user. May be omitted.
	IsUnread bool `json:"isUnread"`
	// Reactions to the message, if any.
	Reactions []Reaction `json:"reactions"`
	// Resolved sender display name (impersonator/full name/username/participant name).
	SenderName string `json:"senderName"`
	// Plain-text body if present. May include a JSON fallback with text entities for
	// rich messages.
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		ChatID      respjson.Field
		MessageID   respjson.Field
		SenderID    respjson.Field
		SortKey     respjson.Field
		Timestamp   respjson.Field
		Attachments respjson.Field
		IsSender    respjson.Field
		IsUnread    respjson.Field
		Reactions   respjson.Field
		SenderName  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageSortKeyUnion contains all possible properties and values from [string],
// [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type MessageSortKeyUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u MessageSortKeyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageSortKeyUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageSortKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageSortKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Reaction struct {
	// Reaction ID, typically ${participantID}${reactionKey} if multiple reactions
	// allowed, or just participantID otherwise.
	ID string `json:"id,required"`
	// User ID of the participant who reacted.
	ParticipantID string `json:"participantID,required"`
	// The reaction key: an emoji (😄), a network-specific key, or a shortcode like
	// "smiling-face".
	ReactionKey string `json:"reactionKey,required"`
	// True if the reactionKey is an emoji.
	Emoji bool `json:"emoji"`
	// URL to the reaction's image. May be temporary or local-only to this device;
	// download promptly if durable access is needed.
	ImgURL string `json:"imgURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ParticipantID respjson.Field
		ReactionKey   respjson.Field
		Emoji         respjson.Field
		ImgURL        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Reaction) RawJSON() string { return r.JSON.raw }
func (r *Reaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A person on or reachable through Beeper. Values are best-effort and can vary by
// network.
type User struct {
	// Stable Beeper user ID. Use as the primary key when referencing a person.
	ID string `json:"id,required"`
	// True if Beeper cannot initiate messages to this user (e.g., blocked, network
	// restriction, or no DM path). The user may still message you.
	CannotMessage bool `json:"cannotMessage"`
	// Email address if known. Not guaranteed verified.
	Email string `json:"email"`
	// Display name as shown in clients (e.g., 'Alice Example'). May include emojis.
	FullName string `json:"fullName,nullable"`
	// Avatar image URL if available. May be temporary or local-only to this device;
	// download promptly if durable access is needed.
	ImgURL string `json:"imgURL"`
	// True if this user represents the authenticated account's own identity.
	IsSelf bool `json:"isSelf"`
	// User's phone number in E.164 format (e.g., '+14155552671'). Omit if unknown.
	PhoneNumber string `json:"phoneNumber"`
	// Human-readable handle if available (e.g., '@alice'). May be network-specific and
	// not globally unique.
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CannotMessage respjson.Field
		Email         respjson.Field
		FullName      respjson.Field
		ImgURL        respjson.Field
		IsSelf        respjson.Field
		PhoneNumber   respjson.Field
		Username      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

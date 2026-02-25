// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Attachment struct {
	// Attachment type.
	//
	// Any of "unknown", "img", "video", "audio".
	Type AttachmentType `json:"type" api:"required"`
	// Attachment identifier (typically an mxc:// URL). Use with /v1/assets/download to
	// get a local file path.
	ID string `json:"id"`
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
		ID          respjson.Field
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

type Message struct {
	// Message ID.
	ID string `json:"id" api:"required"`
	// Beeper account ID the message belongs to.
	AccountID string `json:"accountID" api:"required"`
	// Unique identifier of the chat.
	ChatID string `json:"chatID" api:"required"`
	// Sender user ID.
	SenderID string `json:"senderID" api:"required"`
	// A unique, sortable key used to sort messages.
	SortKey string `json:"sortKey" api:"required"`
	// Message timestamp.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Attachments included with this message, if any.
	Attachments []Attachment `json:"attachments"`
	// True if the authenticated user sent the message.
	IsSender bool `json:"isSender"`
	// True if the message is unread for the authenticated user. May be omitted.
	IsUnread bool `json:"isUnread"`
	// ID of the message this is a reply to, if any.
	LinkedMessageID string `json:"linkedMessageID"`
	// Reactions to the message, if any.
	Reactions []Reaction `json:"reactions"`
	// Resolved sender display name (impersonator/full name/username/participant name).
	SenderName string `json:"senderName"`
	// Plain-text body if present. May include a JSON fallback with text entities for
	// rich messages.
	Text string `json:"text"`
	// Message content type. Useful for distinguishing reactions, media messages, and
	// state events from regular text messages.
	//
	// Any of "TEXT", "NOTICE", "IMAGE", "VIDEO", "VOICE", "AUDIO", "FILE", "STICKER",
	// "LOCATION", "REACTION".
	Type MessageType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		ChatID          respjson.Field
		SenderID        respjson.Field
		SortKey         respjson.Field
		Timestamp       respjson.Field
		Attachments     respjson.Field
		IsSender        respjson.Field
		IsUnread        respjson.Field
		LinkedMessageID respjson.Field
		Reactions       respjson.Field
		SenderName      respjson.Field
		Text            respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message content type. Useful for distinguishing reactions, media messages, and
// state events from regular text messages.
type MessageType string

const (
	MessageTypeText     MessageType = "TEXT"
	MessageTypeNotice   MessageType = "NOTICE"
	MessageTypeImage    MessageType = "IMAGE"
	MessageTypeVideo    MessageType = "VIDEO"
	MessageTypeVoice    MessageType = "VOICE"
	MessageTypeAudio    MessageType = "AUDIO"
	MessageTypeFile     MessageType = "FILE"
	MessageTypeSticker  MessageType = "STICKER"
	MessageTypeLocation MessageType = "LOCATION"
	MessageTypeReaction MessageType = "REACTION"
)

type Reaction struct {
	// Reaction ID, typically ${participantID}${reactionKey} if multiple reactions
	// allowed, or just participantID otherwise.
	ID string `json:"id" api:"required"`
	// User ID of the participant who reacted.
	ParticipantID string `json:"participantID" api:"required"`
	// The reaction key: an emoji (😄), a network-specific key, or a shortcode like
	// "smiling-face".
	ReactionKey string `json:"reactionKey" api:"required"`
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

// User the account belongs to.
type User struct {
	// Stable Beeper user ID. Use as the primary key when referencing a person.
	ID string `json:"id" api:"required"`
	// True if Beeper cannot initiate messages to this user (e.g., blocked, network
	// restriction, or no DM path). The user may still message you.
	CannotMessage bool `json:"cannotMessage"`
	// Email address if known. Not guaranteed verified.
	Email string `json:"email"`
	// Display name as shown in clients (e.g., 'Alice Example'). May include emojis.
	FullName string `json:"fullName"`
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

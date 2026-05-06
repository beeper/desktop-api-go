// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/packages/param"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
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
	// Attachment identifier (typically an mxc:// URL). Use the download file endpoint
	// to get a local file path.
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
	// Public URL or local file path to fetch the file. May be temporary or local-only
	// to this device; download promptly if durable access is needed.
	SrcURL string `json:"srcURL"`
	// Attachment transcription if available.
	Transcription AttachmentTranscription `json:"transcription"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		ID            respjson.Field
		Duration      respjson.Field
		FileName      respjson.Field
		FileSize      respjson.Field
		IsGif         respjson.Field
		IsSticker     respjson.Field
		IsVoiceNote   respjson.Field
		MimeType      respjson.Field
		PosterImg     respjson.Field
		Size          respjson.Field
		SrcURL        respjson.Field
		Transcription respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
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

// Attachment transcription if available.
type AttachmentTranscription struct {
	// Transcription engine.
	Engine string `json:"engine" api:"required"`
	// Transcribed text.
	Transcription string `json:"transcription" api:"required"`
	// Detected or selected language.
	Language string `json:"language"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Engine        respjson.Field
		Transcription respjson.Field
		Language      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentTranscription) RawJSON() string { return r.JSON.raw }
func (r *AttachmentTranscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Message struct {
	// Message ID.
	ID string `json:"id" api:"required"`
	// Beeper account ID the message belongs to.
	AccountID string `json:"accountID" api:"required"`
	// Chat ID. Input routes also accept the local chat ID from this Beeper Desktop
	// installation when available.
	ChatID string `json:"chatID" api:"required"`
	// Matrix-style fully-qualified sender user ID, usually including a bridge prefix
	// and homeserver.
	SenderID string `json:"senderID" api:"required"`
	// A unique, sortable key used to sort messages.
	SortKey string `json:"sortKey" api:"required"`
	// Message timestamp.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Attachments included with this message, if any.
	Attachments []Attachment `json:"attachments"`
	// Timestamp when the message was edited, if known.
	EditedTimestamp time.Time `json:"editedTimestamp" format:"date-time"`
	// True if the message has been deleted.
	IsDeleted bool `json:"isDeleted"`
	// True if the message is hidden from normal display.
	IsHidden bool `json:"isHidden"`
	// True if the authenticated user sent the message.
	IsSender bool `json:"isSender"`
	// True if the message is unread for the authenticated user. May be omitted.
	IsUnread bool `json:"isUnread"`
	// ID of the message this is a reply to, if any.
	LinkedMessageID string `json:"linkedMessageID"`
	// Link previews included with this message, if any.
	Links []MessageLink `json:"links"`
	// Mentioned user IDs, @room, or null for legacy messages that require text
	// scanning.
	Mentions []string `json:"mentions" api:"nullable"`
	// Reactions to the message, if any.
	Reactions []Reaction `json:"reactions"`
	// Read receipt state for this message, when available.
	Seen MessageSeenUnion `json:"seen" format:"date-time"`
	// Resolved sender display name (impersonator/full name/username/participant name).
	SenderName string `json:"senderName"`
	// Message send status for this message, when reported by the bridge.
	SendStatus MessageSendStatus `json:"sendStatus"`
	// Matrix HTML body if present.
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
		EditedTimestamp respjson.Field
		IsDeleted       respjson.Field
		IsHidden        respjson.Field
		IsSender        respjson.Field
		IsUnread        respjson.Field
		LinkedMessageID respjson.Field
		Links           respjson.Field
		Mentions        respjson.Field
		Reactions       respjson.Field
		Seen            respjson.Field
		SenderName      respjson.Field
		SendStatus      respjson.Field
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

// Link preview included with a message.
type MessageLink struct {
	// Link preview title.
	Title string `json:"title" api:"required"`
	// Resolved link URL.
	URL string `json:"url" api:"required"`
	// Favicon URL if available. May be temporary or local-only to this device;
	// download promptly if durable access is needed.
	Favicon string `json:"favicon"`
	// Preview image URL if available. May be temporary or local-only to this device;
	// download promptly if durable access is needed.
	Img string `json:"img"`
	// Preview image dimensions.
	ImgSize MessageLinkImgSize `json:"imgSize"`
	// Original URL when the displayed URL is shortened or redirected.
	OriginalURL string `json:"originalURL"`
	// Link preview summary.
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Title       respjson.Field
		URL         respjson.Field
		Favicon     respjson.Field
		Img         respjson.Field
		ImgSize     respjson.Field
		OriginalURL respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageLink) RawJSON() string { return r.JSON.raw }
func (r *MessageLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Preview image dimensions.
type MessageLinkImgSize struct {
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
func (r MessageLinkImgSize) RawJSON() string { return r.JSON.raw }
func (r *MessageLinkImgSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageSeenUnion contains all possible properties and values from [bool],
// [time.Time], [map[string]MessageSeenByParticipantItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBoolean OfTimestamp]
type MessageSeenUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBoolean bool `json:",inline"`
	// This field will be present if the value is a [time.Time] instead of an object.
	OfTimestamp time.Time `json:",inline"`
	JSON        struct {
		OfBoolean   respjson.Field
		OfTimestamp respjson.Field
		raw         string
	} `json:"-"`
}

func (u MessageSeenUnion) AsBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageSeenUnion) AsTimestamp() (v time.Time) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageSeenUnion) AsByParticipant() (v map[string]MessageSeenByParticipantItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageSeenUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageSeenUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageSeenByParticipantItemUnion contains all possible properties and values
// from [bool], [time.Time].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBoolean OfTimestamp]
type MessageSeenByParticipantItemUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBoolean bool `json:",inline"`
	// This field will be present if the value is a [time.Time] instead of an object.
	OfTimestamp time.Time `json:",inline"`
	JSON        struct {
		OfBoolean   respjson.Field
		OfTimestamp respjson.Field
		raw         string
	} `json:"-"`
}

func (u MessageSeenByParticipantItemUnion) AsBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageSeenByParticipantItemUnion) AsTimestamp() (v time.Time) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageSeenByParticipantItemUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageSeenByParticipantItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message send status for this message, when reported by the bridge.
type MessageSendStatus struct {
	// Current status of the message send attempt.
	//
	// Any of "SUCCESS", "PENDING", "FAIL_RETRIABLE", "FAIL_PERMANENT".
	Status string `json:"status" api:"required"`
	// Timestamp for the send status event.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// User IDs the message was delivered to, when reported by the network.
	DeliveredToUsers []string `json:"deliveredToUsers"`
	// Internal bridge error detail. Intended for diagnostics, not end-user display.
	InternalError string `json:"internalError"`
	// Human-readable send status or failure message.
	Message string `json:"message"`
	// Machine-readable failure reason. Present when the send status is a failure.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status           respjson.Field
		Timestamp        respjson.Field
		DeliveredToUsers respjson.Field
		InternalError    respjson.Field
		Message          respjson.Field
		Reason           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendStatus) RawJSON() string { return r.JSON.raw }
func (r *MessageSendStatus) UnmarshalJSON(data []byte) error {
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
	// Reaction ID. When a participant can react more than once, the ID is the
	// participant ID concatenated with the reaction key; otherwise it equals the
	// participant ID.
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
	// Avatar image URL if available. This may be a remote URL, Matrix media URL, data
	// URL, or local filesystem URL depending on source and endpoint. May be temporary
	// or local-only to this device; download promptly if durable access is needed.
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

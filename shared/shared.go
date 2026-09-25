// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/beeper/desktop-api-go/v6/internal/apijson"
	"github.com/beeper/desktop-api-go/v6/packages/param"
	"github.com/beeper/desktop-api-go/v6/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type APIError struct {
	Code    string         `json:"code" api:"required"`
	Message string         `json:"message" api:"required"`
	Details map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIError) RawJSON() string { return r.JSON.raw }
func (r *APIError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Attachment struct {
	// Attachment type.
	//
	// Any of "unknown", "img", "video", "audio".
	Type AttachmentType `json:"type" api:"required"`
	// Attachment identifier, typically an mxc:// URL. Use the download file endpoint
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
	// available only on this device; download promptly if durable access is needed.
	PosterImg string `json:"posterImg"`
	// Pixel dimensions of the attachment: width/height in px.
	Size AttachmentSize `json:"size"`
	// Public URL or local file path to fetch the file. May be temporary or available
	// only on this device; download promptly if durable access is needed.
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

// Capabilities for one attachment message type.
type AttachmentCapabilities struct {
	// Supported MIME types or MIME patterns for this file message type. Missing MIME
	// types should be treated as rejected.
	//
	// Any of -2, -1, 0, 1, 2.
	MimeTypes map[string]int64 `json:"mimeTypes" api:"required"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Caption int64 `json:"caption"`
	// Maximum caption length when captions are supported.
	MaxCaptionLength int64 `json:"maxCaptionLength"`
	// Maximum audio or video duration in seconds.
	MaxDuration int64 `json:"maxDuration"`
	// Maximum image or video height in pixels.
	MaxHeight int64 `json:"maxHeight"`
	// Maximum file size in bytes.
	MaxSize int64 `json:"maxSize"`
	// Maximum image or video width in pixels.
	MaxWidth int64 `json:"maxWidth"`
	// True if this file type can be sent as view-once media.
	ViewOnce bool `json:"viewOnce"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MimeTypes        respjson.Field
		Caption          respjson.Field
		MaxCaptionLength respjson.Field
		MaxDuration      respjson.Field
		MaxHeight        respjson.Field
		MaxSize          respjson.Field
		MaxWidth         respjson.Field
		ViewOnce         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentCapabilities) RawJSON() string { return r.JSON.raw }
func (r *AttachmentCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat capabilities reported by the platform.
type ChatCapabilities struct {
	// Allowed Unicode reactions. Omitted means all emoji reactions are allowed.
	AllowedReactions []string `json:"allowedReactions"`
	// True if archive/unarchive is supported.
	Archive bool `json:"archive"`
	// Supported attachment message types and their per-type constraints, keyed by
	// Matrix msgtype or pseudo-msgtype (for example m.image, m.video,
	// org.matrix.msc3245.voice). Missing message types should be treated as rejected.
	Attachments map[string]AttachmentCapabilities `json:"attachments"`
	// True if custom emoji reactions are supported.
	CustomEmojiReactions bool `json:"customEmojiReactions"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Delete int64 `json:"delete"`
	// True if deleting chats for the authenticated user is supported.
	DeleteChat bool `json:"deleteChat"`
	// True if deleting chats for everyone is supported.
	DeleteChatForEveryone bool `json:"deleteChatForEveryone"`
	// True if deleting messages only for the authenticated user is supported.
	DeleteForMe bool `json:"deleteForMe"`
	// Maximum message age for delete-for-everyone, in seconds.
	DeleteMaxAge int64 `json:"deleteMaxAge"`
	// Disappearing-message timer capabilities.
	DisappearingTimer ChatCapabilitiesDisappearingTimer `json:"disappearingTimer"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Edit int64 `json:"edit"`
	// Maximum message age for edits, in seconds.
	EditMaxAge int64 `json:"editMaxAge"`
	// Maximum number of edits allowed for one message.
	EditMaxCount int64 `json:"editMaxCount"`
	// Supported rich-text formatting features keyed by feature name (for example bold,
	// inline_code, code_block.syntax_highlighting). Omitted means no formatting
	// support is advertised.
	//
	// Any of -2, -1, 0, 1, 2.
	Formatting map[string]int64 `json:"formatting"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	LocationMessage int64 `json:"locationMessage"`
	// True if marking chats unread is supported.
	MarkAsUnread bool `json:"markAsUnread"`
	// Maximum length of normal text messages.
	MaxTextLength int64 `json:"maxTextLength"`
	// Message request capabilities.
	MessageRequest ChatCapabilitiesMessageRequest `json:"messageRequest"`
	// Participant management capabilities.
	ParticipantActions ChatCapabilitiesParticipantActions `json:"participantActions"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Poll int64 `json:"poll"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Reaction int64 `json:"reaction"`
	// Maximum number of reactions allowed on a single message.
	ReactionCount int64 `json:"reactionCount"`
	// True if read receipts are supported.
	ReadReceipts bool `json:"readReceipts"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Reply int64 `json:"reply"`
	// Chat state update capabilities.
	State ChatStateCapabilities `json:"state"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Thread int64 `json:"thread"`
	// True if typing notifications are supported.
	TypingNotifications bool `json:"typingNotifications"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedReactions      respjson.Field
		Archive               respjson.Field
		Attachments           respjson.Field
		CustomEmojiReactions  respjson.Field
		Delete                respjson.Field
		DeleteChat            respjson.Field
		DeleteChatForEveryone respjson.Field
		DeleteForMe           respjson.Field
		DeleteMaxAge          respjson.Field
		DisappearingTimer     respjson.Field
		Edit                  respjson.Field
		EditMaxAge            respjson.Field
		EditMaxCount          respjson.Field
		Formatting            respjson.Field
		LocationMessage       respjson.Field
		MarkAsUnread          respjson.Field
		MaxTextLength         respjson.Field
		MessageRequest        respjson.Field
		ParticipantActions    respjson.Field
		Poll                  respjson.Field
		Reaction              respjson.Field
		ReactionCount         respjson.Field
		ReadReceipts          respjson.Field
		Reply                 respjson.Field
		State                 respjson.Field
		Thread                respjson.Field
		TypingNotifications   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCapabilities) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Disappearing-message timer capabilities.
type ChatCapabilitiesDisappearingTimer struct {
	// True if empty timer objects should be omitted from message content.
	OmitEmptyTimer bool `json:"omitEmptyTimer"`
	// Allowed disappearing timer values in milliseconds. Omitted means any timer is
	// allowed.
	Timers []int64 `json:"timers"`
	// Supported disappearing timer types.
	//
	// Any of "afterRead", "afterReadByRecipient", "afterSend".
	Types []string `json:"types"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OmitEmptyTimer respjson.Field
		Timers         respjson.Field
		Types          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCapabilitiesDisappearingTimer) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilitiesDisappearingTimer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message request capabilities.
type ChatCapabilitiesMessageRequest struct {
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	AcceptWithButton int64 `json:"acceptWithButton"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	AcceptWithMessage int64 `json:"acceptWithMessage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptWithButton  respjson.Field
		AcceptWithMessage respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCapabilitiesMessageRequest) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilitiesMessageRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Participant management capabilities.
type ChatCapabilitiesParticipantActions struct {
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Ban int64 `json:"ban"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Invite int64 `json:"invite"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Kick int64 `json:"kick"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Leave int64 `json:"leave"`
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	RevokeInvite int64 `json:"revokeInvite"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ban          respjson.Field
		Invite       respjson.Field
		Kick         respjson.Field
		Leave        respjson.Field
		RevokeInvite respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCapabilitiesParticipantActions) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilitiesParticipantActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current draft object for this chat, or null when no draft is set.
type ChatDraft struct {
	// Rich-text draft body as returned by Beeper.
	Text string `json:"text" api:"required"`
	// Draft attachments keyed by attachment ID.
	Attachments map[string]DraftAttachment `json:"attachments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Attachments respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatDraft) RawJSON() string { return r.JSON.raw }
func (r *ChatDraft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat state update capabilities.
type ChatStateCapabilities struct {
	// Chat avatar state capability.
	Avatar ChatStateCapabilitiesAvatar `json:"avatar"`
	// Chat description/topic state capability.
	Description ChatStateCapabilitiesDescription `json:"description"`
	// Disappearing-message timer state capability.
	DisappearingTimer ChatStateCapabilitiesDisappearingTimer `json:"disappearingTimer"`
	// Chat title state capability.
	Title ChatStateCapabilitiesTitle `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Avatar            respjson.Field
		Description       respjson.Field
		DisappearingTimer respjson.Field
		Title             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatStateCapabilities) RawJSON() string { return r.JSON.raw }
func (r *ChatStateCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat avatar state capability.
type ChatStateCapabilitiesAvatar struct {
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Level int64 `json:"level" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatStateCapabilitiesAvatar) RawJSON() string { return r.JSON.raw }
func (r *ChatStateCapabilitiesAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat description/topic state capability.
type ChatStateCapabilitiesDescription struct {
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Level int64 `json:"level" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatStateCapabilitiesDescription) RawJSON() string { return r.JSON.raw }
func (r *ChatStateCapabilitiesDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Disappearing-message timer state capability.
type ChatStateCapabilitiesDisappearingTimer struct {
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Level int64 `json:"level" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatStateCapabilitiesDisappearingTimer) RawJSON() string { return r.JSON.raw }
func (r *ChatStateCapabilitiesDisappearingTimer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat title state capability.
type ChatStateCapabilitiesTitle struct {
	// -2: rejected, -1: dropped, 0: unsupported, 1: partially supported, 2: fully
	// supported.
	//
	// Any of -2, -1, 0, 1, 2.
	Level int64 `json:"level" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatStateCapabilitiesTitle) RawJSON() string { return r.JSON.raw }
func (r *ChatStateCapabilitiesTitle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DraftAttachment struct {
	// Draft attachment identifier.
	ID string `json:"id" api:"required"`
	// Draft attachment type. GIF and recorded audio are mutually exclusive types.
	//
	// Any of "file", "gif", "recorded_audio".
	Type DraftAttachmentType `json:"type" api:"required"`
	// Audio duration in seconds if known.
	AudioDurationSeconds float64 `json:"audioDurationSeconds"`
	// Original filename if available.
	FileName string `json:"fileName"`
	// Local filesystem path for the draft attachment.
	FilePath string `json:"filePath"`
	// File size in bytes if known.
	FileSize float64 `json:"fileSize"`
	// MIME type if known.
	MimeType string `json:"mimeType"`
	// Pixel dimensions of the attachment.
	Size DraftAttachmentSize `json:"size"`
	// Sticker identifier if the draft attachment is a sticker.
	StickerID string `json:"stickerID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Type                 respjson.Field
		AudioDurationSeconds respjson.Field
		FileName             respjson.Field
		FilePath             respjson.Field
		FileSize             respjson.Field
		MimeType             respjson.Field
		Size                 respjson.Field
		StickerID            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DraftAttachment) RawJSON() string { return r.JSON.raw }
func (r *DraftAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Draft attachment type. GIF and recorded audio are mutually exclusive types.
type DraftAttachmentType string

const (
	DraftAttachmentTypeFile          DraftAttachmentType = "file"
	DraftAttachmentTypeGif           DraftAttachmentType = "gif"
	DraftAttachmentTypeRecordedAudio DraftAttachmentType = "recorded_audio"
)

// Pixel dimensions of the attachment.
type DraftAttachmentSize struct {
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
func (r DraftAttachmentSize) RawJSON() string { return r.JSON.raw }
func (r *DraftAttachmentSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Link preview included with a message.
type LinkPreview struct {
	// Link preview title.
	Title string `json:"title" api:"required"`
	// Resolved link URL.
	URL string `json:"url" api:"required"`
	// Favicon URL if available. May be temporary or available only on this device;
	// download promptly if durable access is needed.
	Favicon string `json:"favicon"`
	// Preview image URL if available. May be temporary or available only on this
	// device; download promptly if durable access is needed.
	Img string `json:"img"`
	// Preview image dimensions.
	ImgSize LinkPreviewImgSize `json:"imgSize"`
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
func (r LinkPreview) RawJSON() string { return r.JSON.raw }
func (r *LinkPreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Preview image dimensions.
type LinkPreviewImgSize struct {
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
func (r LinkPreviewImgSize) RawJSON() string { return r.JSON.raw }
func (r *LinkPreviewImgSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Message struct {
	// Message ID.
	ID string `json:"id" api:"required"`
	// Beeper account ID the message belongs to.
	AccountID string `json:"accountID" api:"required"`
	// Chat ID. Input routes also accept the local chat ID from this installation when
	// available.
	ChatID string `json:"chatID" api:"required"`
	// Fully qualified sender user ID. Network-backed IDs usually include the network
	// prefix and homeserver.
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
	Links []LinkPreview `json:"links"`
	// Mentioned user IDs, @room, or null for legacy messages that require text
	// scanning.
	Mentions []string `json:"mentions" api:"nullable"`
	// Reactions to the message, if any.
	Reactions []Reaction `json:"reactions"`
	// Read receipt state for this message, when available.
	Seen MessageSeenUnion `json:"seen" format:"date-time"`
	// Resolved sender display name.
	SenderName string `json:"senderName"`
	// Message send status for this message, when reported by the bridge.
	SendStatus SendStatus `json:"sendStatus"`
	// Rich-text message body if present.
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
	// URL to the reaction's image. May be temporary or available only on this device;
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

// Message send status for this message, when reported by the bridge.
type SendStatus struct {
	// Current status of the message send attempt.
	//
	// Any of "SUCCESS", "PENDING", "FAIL_RETRIABLE", "FAIL_PERMANENT".
	Status SendStatusStatus `json:"status" api:"required"`
	// Timestamp for the send status event.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// User IDs the message was delivered to, when reported by the network.
	DeliveredToUsers []string `json:"deliveredToUsers"`
	// Diagnostic error detail from the messaging network adapter. Do not show directly
	// to users.
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
func (r SendStatus) RawJSON() string { return r.JSON.raw }
func (r *SendStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the message send attempt.
type SendStatusStatus string

const (
	SendStatusStatusSuccess       SendStatusStatus = "SUCCESS"
	SendStatusStatusPending       SendStatusStatus = "PENDING"
	SendStatusStatusFailRetriable SendStatusStatus = "FAIL_RETRIABLE"
	SendStatusStatusFailPermanent SendStatusStatus = "FAIL_PERMANENT"
)

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
	// Avatar image URL if available. This may be a remote URL, media URL, data URL, or
	// local file URL depending on the source. May be temporary or available only on
	// this device; download promptly if durable access is needed.
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

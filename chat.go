// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/apiquery"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/pagination"
	"github.com/beeper/desktop-api-go/v5/packages/param"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
	"github.com/beeper/desktop-api-go/v5/shared"
)

// Manage chats
//
// ChatService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options []option.RequestOption
	// Manage reminders for chats
	Reminders ChatReminderService
	// Manage chat messages
	Messages ChatMessageService
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.Options = opts
	r.Reminders = NewChatReminderService(opts...)
	r.Messages = NewChatMessageService(opts...)
	return
}

// Create a direct or group chat from participant IDs. Returns the created chat.
func (r *ChatService) New(ctx context.Context, body ChatNewParams, opts ...option.RequestOption) (res *ChatNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve chat details including metadata, participants, and latest message
func (r *ChatService) Get(ctx context.Context, chatID string, query ChatGetParams, opts ...option.RequestOption) (res *Chat, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/chats/%s", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update supported chat fields. Non-empty draft objects are accepted only when the
// current draft is empty. Send draft=null to clear the draft before setting new
// draft text or attachments.
func (r *ChatService) Update(ctx context.Context, chatID string, body ChatUpdateParams, opts ...option.RequestOption) (res *Chat, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/chats/%s", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all chats sorted by last activity (most recent first). Combines all
// accounts into a single paginated list.
func (r *ChatService) List(ctx context.Context, query ChatListParams, opts ...option.RequestOption) (res *pagination.CursorNoLimit[ChatListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/chats"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all chats sorted by last activity (most recent first). Combines all
// accounts into a single paginated list.
func (r *ChatService) ListAutoPaging(ctx context.Context, query ChatListParams, opts ...option.RequestOption) *pagination.CursorNoLimitAutoPager[ChatListResponse] {
	return pagination.NewCursorNoLimitAutoPager(r.List(ctx, query, opts...))
}

// Archive or unarchive a chat. Set archived=true to move to archive,
// archived=false to move back to inbox
func (r *ChatService) Archive(ctx context.Context, chatID string, body ChatArchiveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return err
	}
	path := fmt.Sprintf("v1/chats/%s/archive", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Mark a chat as read, optionally through a specific message ID.
func (r *ChatService) MarkRead(ctx context.Context, chatID string, body ChatMarkReadParams, opts ...option.RequestOption) (res *Chat, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/chats/%s/read", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Mark a chat as unread, optionally from a specific message ID.
func (r *ChatService) MarkUnread(ctx context.Context, chatID string, body ChatMarkUnreadParams, opts ...option.RequestOption) (res *Chat, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/chats/%s/unread", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Force a delivery notification when supported by the underlying network.
// Currently intended for iMessage on macOS; unsupported networks return an error.
func (r *ChatService) NotifyAnyway(ctx context.Context, chatID string, body ChatNotifyAnywayParams, opts ...option.RequestOption) (res *Chat, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/chats/%s/notify-anyway", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Search chats by title, network, or participant names.
func (r *ChatService) Search(ctx context.Context, query ChatSearchParams, opts ...option.RequestOption) (res *pagination.CursorSearch[Chat], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/chats/search"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Search chats by title, network, or participant names.
func (r *ChatService) SearchAutoPaging(ctx context.Context, query ChatSearchParams, opts ...option.RequestOption) *pagination.CursorSearchAutoPager[Chat] {
	return pagination.NewCursorSearchAutoPager(r.Search(ctx, query, opts...))
}

// Resolve a user/contact and open a direct chat. Reuses and returns an existing
// direct chat when one is found. Available in Beeper Desktop v4.2.808+.
func (r *ChatService) Start(ctx context.Context, body ChatStartParams, opts ...option.RequestOption) (res *ChatStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chats/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type Chat struct {
	// Unique identifier of the chat across Beeper.
	ID string `json:"id" api:"required"`
	// Account ID this chat belongs to.
	AccountID string `json:"accountID" api:"required"`
	// Display-only human-readable account/network name.
	Network string `json:"network" api:"required"`
	// Chat participants information.
	Participants ChatParticipants `json:"participants" api:"required"`
	// Display title of the chat as computed by the client/server.
	Title string `json:"title" api:"required"`
	// Chat type: 'single' for direct messages, 'group' for group chats.
	//
	// Any of "single", "group".
	Type ChatType `json:"type" api:"required"`
	// Number of unread messages.
	UnreadCount int64 `json:"unreadCount" api:"required"`
	// Chat capabilities reported by the platform.
	Capabilities ChatCapabilities `json:"capabilities"`
	// Group chat description/topic when available.
	Description string `json:"description" api:"nullable"`
	// Current draft object for this chat, or null when no draft is set.
	Draft ChatDraft `json:"draft" api:"nullable"`
	// Local filesystem path to the chat avatar image when available.
	ImgURL string `json:"imgURL" api:"nullable"`
	// True if chat is archived.
	IsArchived bool `json:"isArchived"`
	// True if chat is marked low priority.
	IsLowPriority bool `json:"isLowPriority"`
	// True if the chat was explicitly marked unread by the authenticated user.
	IsMarkedUnread bool `json:"isMarkedUnread"`
	// True if chat notifications are muted.
	IsMuted bool `json:"isMuted"`
	// True if chat is pinned.
	IsPinned bool `json:"isPinned"`
	// True if messages cannot be sent in this chat.
	IsReadOnly bool `json:"isReadOnly"`
	// Timestamp of last activity.
	LastActivity time.Time `json:"lastActivity" format:"date-time"`
	// Last read message sortKey.
	LastReadMessageSortKey string `json:"lastReadMessageSortKey"`
	// Local chat ID specific to this Beeper Desktop installation.
	LocalChatID string `json:"localChatID" api:"nullable"`
	// Disappearing-message timer in seconds when available.
	MessageExpirySeconds int64 `json:"messageExpirySeconds" api:"nullable"`
	// Current reminder for this chat, or null when no reminder is set.
	Reminder ChatReminder `json:"reminder" api:"nullable"`
	// Current snooze state for this chat, or null when no snooze is set.
	Snooze ChatSnooze `json:"snooze" api:"nullable"`
	// Number of unread messages that mention the authenticated user or @room.
	UnreadMentionsCount int64 `json:"unreadMentionsCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccountID              respjson.Field
		Network                respjson.Field
		Participants           respjson.Field
		Title                  respjson.Field
		Type                   respjson.Field
		UnreadCount            respjson.Field
		Capabilities           respjson.Field
		Description            respjson.Field
		Draft                  respjson.Field
		ImgURL                 respjson.Field
		IsArchived             respjson.Field
		IsLowPriority          respjson.Field
		IsMarkedUnread         respjson.Field
		IsMuted                respjson.Field
		IsPinned               respjson.Field
		IsReadOnly             respjson.Field
		LastActivity           respjson.Field
		LastReadMessageSortKey respjson.Field
		LocalChatID            respjson.Field
		MessageExpirySeconds   respjson.Field
		Reminder               respjson.Field
		Snooze                 respjson.Field
		UnreadMentionsCount    respjson.Field
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
	HasMore bool `json:"hasMore" api:"required"`
	// Participants returned for this chat (limited by the request; may be a subset).
	Items []ChatParticipantsItem `json:"items" api:"required"`
	// Total number of participants in the chat.
	Total int64 `json:"total" api:"required"`
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

// A chat participant. Extends User with chat membership metadata.
type ChatParticipantsItem struct {
	// True if this participant has admin privileges in the chat.
	IsAdmin bool `json:"isAdmin"`
	// True if this participant represents a network or bridge bot.
	IsNetworkBot bool `json:"isNetworkBot"`
	// True if this participant has been invited but has not joined yet.
	IsPending bool `json:"isPending"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsAdmin      respjson.Field
		IsNetworkBot respjson.Field
		IsPending    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	shared.User
}

// Returns the unmodified JSON received from the API
func (r ChatParticipantsItem) RawJSON() string { return r.JSON.raw }
func (r *ChatParticipantsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat type: 'single' for direct messages, 'group' for group chats.
type ChatType string

const (
	ChatTypeSingle ChatType = "single"
	ChatTypeGroup  ChatType = "group"
)

// Chat capabilities reported by the platform.
type ChatCapabilities struct {
	// Allowed Unicode reactions. Omitted means all emoji reactions are allowed.
	AllowedReactions []string `json:"allowedReactions"`
	// True if archive/unarchive is supported.
	Archive bool `json:"archive"`
	// Supported attachment message types and their per-type constraints, keyed by
	// Matrix msgtype or pseudo-msgtype (for example m.image, m.video,
	// org.matrix.msc3245.voice). Missing message types should be treated as rejected.
	Attachments map[string]ChatCapabilitiesAttachment `json:"attachments"`
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
	State ChatCapabilitiesState `json:"state"`
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

// Capabilities for one attachment message type.
type ChatCapabilitiesAttachment struct {
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
func (r ChatCapabilitiesAttachment) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilitiesAttachment) UnmarshalJSON(data []byte) error {
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
	// Any of "afterRead", "afterSend".
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

// Chat state update capabilities.
type ChatCapabilitiesState struct {
	// Chat avatar state capability.
	Avatar ChatCapabilitiesStateAvatar `json:"avatar"`
	// Chat description/topic state capability.
	Description ChatCapabilitiesStateDescription `json:"description"`
	// Disappearing-message timer state capability.
	DisappearingTimer ChatCapabilitiesStateDisappearingTimer `json:"disappearingTimer"`
	// Chat title state capability.
	Title ChatCapabilitiesStateTitle `json:"title"`
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
func (r ChatCapabilitiesState) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilitiesState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat avatar state capability.
type ChatCapabilitiesStateAvatar struct {
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
func (r ChatCapabilitiesStateAvatar) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilitiesStateAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat description/topic state capability.
type ChatCapabilitiesStateDescription struct {
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
func (r ChatCapabilitiesStateDescription) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilitiesStateDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Disappearing-message timer state capability.
type ChatCapabilitiesStateDisappearingTimer struct {
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
func (r ChatCapabilitiesStateDisappearingTimer) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilitiesStateDisappearingTimer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat title state capability.
type ChatCapabilitiesStateTitle struct {
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
func (r ChatCapabilitiesStateTitle) RawJSON() string { return r.JSON.raw }
func (r *ChatCapabilitiesStateTitle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current draft object for this chat, or null when no draft is set.
type ChatDraft struct {
	// Matrix HTML draft body.
	Text string `json:"text" api:"required"`
	// Draft attachments keyed by attachment ID.
	Attachments map[string]ChatDraftAttachment `json:"attachments"`
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

type ChatDraftAttachment struct {
	// Draft attachment identifier.
	ID string `json:"id" api:"required"`
	// Draft attachment type. GIF and recorded audio are mutually exclusive types.
	//
	// Any of "file", "gif", "recorded_audio".
	Type string `json:"type" api:"required"`
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
	Size ChatDraftAttachmentSize `json:"size"`
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
func (r ChatDraftAttachment) RawJSON() string { return r.JSON.raw }
func (r *ChatDraftAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pixel dimensions of the attachment.
type ChatDraftAttachmentSize struct {
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
func (r ChatDraftAttachmentSize) RawJSON() string { return r.JSON.raw }
func (r *ChatDraftAttachmentSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current reminder for this chat, or null when no reminder is set.
type ChatReminder struct {
	// Cancel reminder if someone messages in the chat.
	DismissOnIncomingMessage bool `json:"dismissOnIncomingMessage"`
	// Timestamp when the reminder should trigger.
	RemindAt time.Time `json:"remindAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DismissOnIncomingMessage respjson.Field
		RemindAt                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatReminder) RawJSON() string { return r.JSON.raw }
func (r *ChatReminder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current snooze state for this chat, or null when no snooze is set.
type ChatSnooze struct {
	// Timestamp when the snooze expires.
	SnoozeUntil time.Time `json:"snoozeUntil" format:"date-time"`
	// Timestamp when the user set the snooze.
	UserSnoozedAt time.Time `json:"userSnoozedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SnoozeUntil   respjson.Field
		UserSnoozedAt respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatSnooze) RawJSON() string { return r.JSON.raw }
func (r *ChatSnooze) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewResponse struct {
	// DEPRECATED - use id instead. Compatibility alias for older clients.
	//
	// Deprecated: deprecated
	ChatID string `json:"chatID" api:"required"`
	// DEPRECATED - legacy start-chat status for older clients. New clients should
	// inspect the returned Chat instead.
	//
	// Any of "existing", "created".
	//
	// Deprecated: deprecated
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Chat
}

// Returns the unmodified JSON received from the API
func (r ChatNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat with optional last message preview.
type ChatListResponse struct {
	// Last message preview for this chat, if available.
	Preview shared.Message `json:"preview"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Preview     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Chat
}

// Returns the unmodified JSON received from the API
func (r ChatListResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatStartResponse struct {
	// DEPRECATED - use id instead. Compatibility alias for older clients.
	//
	// Deprecated: deprecated
	ChatID string `json:"chatID" api:"required"`
	// DEPRECATED - legacy start-chat status for older clients. New clients should
	// inspect the returned Chat instead.
	//
	// Any of "existing", "created".
	//
	// Deprecated: deprecated
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Chat
}

// Returns the unmodified JSON received from the API
func (r ChatStartResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewParams struct {
	// Account to create or start the chat on.
	AccountID string `json:"accountID" api:"required"`
	// User IDs to include in the new chat.
	ParticipantIDs []string `json:"participantIDs,omitzero" api:"required"`
	// 'single' requires exactly one participantID; 'group' supports multiple
	// participants and optional title.
	//
	// Any of "single", "group".
	Type ChatNewParamsType `json:"type,omitzero" api:"required"`
	// Optional first message content if the platform requires it to create the chat.
	MessageText param.Opt[string] `json:"messageText,omitzero"`
	// Optional title for group chats; ignored for single chats on most networks.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r ChatNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 'single' requires exactly one participantID; 'group' supports multiple
// participants and optional title.
type ChatNewParamsType string

const (
	ChatNewParamsTypeSingle ChatNewParamsType = "single"
	ChatNewParamsTypeGroup  ChatNewParamsType = "group"
)

type ChatGetParams struct {
	// Maximum number of participants to return. Use -1 for all; otherwise 0-500.
	// Defaults to 100. List and search endpoints return up to 20 participants per
	// chat.
	MaxParticipantCount param.Opt[int64] `query:"maxParticipantCount,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChatGetParams]'s query parameters as `url.Values`.
func (r ChatGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChatUpdateParams struct {
	// Group chat description/topic. Support depends on the chat account and chat
	// permissions.
	Description param.Opt[string] `json:"description,omitzero"`
	// Local filesystem path to a group chat avatar image. Support depends on the chat
	// account and chat permissions.
	ImgURL param.Opt[string] `json:"imgURL,omitzero"`
	// Disappearing-message timer in seconds, or null to clear when supported.
	MessageExpirySeconds param.Opt[int64] `json:"messageExpirySeconds,omitzero"`
	// Custom chat title. Support depends on the chat account and chat permissions.
	Title param.Opt[string] `json:"title,omitzero"`
	// Archive or unarchive the chat.
	IsArchived param.Opt[bool] `json:"isArchived,omitzero"`
	// Mark or unmark the chat as low priority when supported by the account.
	IsLowPriority param.Opt[bool] `json:"isLowPriority,omitzero"`
	// Mute or unmute the chat.
	IsMuted param.Opt[bool] `json:"isMuted,omitzero"`
	// Pin or unpin the chat when supported by the account.
	IsPinned param.Opt[bool] `json:"isPinned,omitzero"`
	// Draft object to set or clear. Non-empty drafts are only accepted when the
	// current draft is empty. Send draft=null to clear text and attachments together
	// before setting a new draft.
	Draft ChatUpdateParamsDraft `json:"draft,omitzero"`
	paramObj
}

func (r ChatUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Draft object to set or clear. Non-empty drafts are only accepted when the
// current draft is empty. Send draft=null to clear text and attachments together
// before setting a new draft.
//
// The property Text is required.
type ChatUpdateParamsDraft struct {
	// Draft text. Plain text and Markdown are converted to Matrix HTML with the same
	// rules used by send and edit.
	Text string `json:"text" api:"required"`
	// Draft attachments keyed by attachment ID. Each attachment must reference an
	// uploadID returned by the upload file endpoint.
	Attachments map[string]ChatUpdateParamsDraftAttachment `json:"attachments,omitzero"`
	paramObj
}

func (r ChatUpdateParamsDraft) MarshalJSON() (data []byte, err error) {
	type shadow ChatUpdateParamsDraft
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatUpdateParamsDraft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property UploadID is required.
type ChatUpdateParamsDraftAttachment struct {
	// Upload ID from uploadAsset endpoint. Required to reference uploaded files.
	UploadID string `json:"uploadID" api:"required"`
	// Optional draft attachment identifier. If omitted, a new identifier is generated.
	ID param.Opt[string] `json:"id,omitzero"`
	// Duration in seconds (optional override of cached value)
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Filename (optional override of cached value)
	FileName param.Opt[string] `json:"fileName,omitzero"`
	// MIME type (optional override of cached value)
	MimeType param.Opt[string] `json:"mimeType,omitzero"`
	// Dimensions (optional override of cached value)
	Size ChatUpdateParamsDraftAttachmentSize `json:"size,omitzero"`
	// Attachment type hint (image, video, audio, file, gif, voice-note, sticker). If
	// omitted, auto-detected from mimeType
	//
	// Any of "image", "video", "audio", "file", "gif", "voice-note", "sticker".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatUpdateParamsDraftAttachment) MarshalJSON() (data []byte, err error) {
	type shadow ChatUpdateParamsDraftAttachment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatUpdateParamsDraftAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatUpdateParamsDraftAttachment](
		"type", "image", "video", "audio", "file", "gif", "voice-note", "sticker",
	)
}

// Dimensions (optional override of cached value)
//
// The properties Height, Width are required.
type ChatUpdateParamsDraftAttachmentSize struct {
	Height float64 `json:"height" api:"required"`
	Width  float64 `json:"width" api:"required"`
	paramObj
}

func (r ChatUpdateParamsDraftAttachmentSize) MarshalJSON() (data []byte, err error) {
	type shadow ChatUpdateParamsDraftAttachmentSize
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatUpdateParamsDraftAttachmentSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatListParams struct {
	// Opaque pagination cursor; do not inspect. Use together with 'direction'.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Limit to specific account IDs. If omitted, fetches from all accounts.
	AccountIDs []string `query:"accountIDs,omitzero" json:"-"`
	// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
	// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
	//
	// Any of "after", "before".
	Direction ChatListParamsDirection `query:"direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChatListParams]'s query parameters as `url.Values`.
func (r ChatListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
type ChatListParamsDirection string

const (
	ChatListParamsDirectionAfter  ChatListParamsDirection = "after"
	ChatListParamsDirectionBefore ChatListParamsDirection = "before"
)

type ChatArchiveParams struct {
	// True to archive, false to unarchive
	Archived param.Opt[bool] `json:"archived,omitzero"`
	paramObj
}

func (r ChatArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMarkReadParams struct {
	// Optional message ID to mark read through.
	MessageID param.Opt[string] `json:"messageID,omitzero"`
	paramObj
}

func (r ChatMarkReadParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatMarkReadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMarkReadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMarkUnreadParams struct {
	// Optional message ID to mark unread from.
	MessageID param.Opt[string] `json:"messageID,omitzero"`
	paramObj
}

func (r ChatMarkUnreadParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatMarkUnreadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMarkUnreadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNotifyAnywayParams struct {
	paramObj
}

func (r ChatNotifyAnywayParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatNotifyAnywayParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNotifyAnywayParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatSearchParams struct {
	// Include chats marked as Muted by the user, which are usually less important.
	// Default: true. Set to false if the user wants a more refined search.
	IncludeMuted param.Opt[bool] `query:"includeMuted,omitzero" json:"-"`
	// Set to true to only retrieve chats that have unread messages
	UnreadOnly param.Opt[bool] `query:"unreadOnly,omitzero" json:"-"`
	// Opaque pagination cursor; do not inspect. Use together with 'direction'.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Provide an ISO datetime string to only retrieve chats with last activity after
	// this time
	LastActivityAfter param.Opt[time.Time] `query:"lastActivityAfter,omitzero" format:"date-time" json:"-"`
	// Provide an ISO datetime string to only retrieve chats with last activity before
	// this time
	LastActivityBefore param.Opt[time.Time] `query:"lastActivityBefore,omitzero" format:"date-time" json:"-"`
	// Set the maximum number of chats to retrieve. Valid range: 1-200, default is 50
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Literal token search (non-semantic). Use single words users type (e.g.,
	// "dinner"). When multiple words provided, ALL must match. Case-insensitive.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Provide an array of account IDs to filter chats from specific messaging accounts
	// only
	AccountIDs []string `query:"accountIDs,omitzero" json:"-"`
	// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
	// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
	//
	// Any of "after", "before".
	Direction ChatSearchParamsDirection `query:"direction,omitzero" json:"-"`
	// Filter by inbox type: "primary" (non-archived, non-low-priority),
	// "low-priority", or "archive". If not specified, shows all chats.
	//
	// Any of "primary", "low-priority", "archive".
	Inbox ChatSearchParamsInbox `query:"inbox,omitzero" json:"-"`
	// Search scope: 'titles' matches title + network; 'participants' matches
	// participant names.
	//
	// Any of "titles", "participants".
	Scope ChatSearchParamsScope `query:"scope,omitzero" json:"-"`
	// Specify the type of chats to retrieve: use "single" for direct messages, "group"
	// for group chats, or "any" to get all types
	//
	// Any of "single", "group", "any".
	Type ChatSearchParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChatSearchParams]'s query parameters as `url.Values`.
func (r ChatSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
type ChatSearchParamsDirection string

const (
	ChatSearchParamsDirectionAfter  ChatSearchParamsDirection = "after"
	ChatSearchParamsDirectionBefore ChatSearchParamsDirection = "before"
)

// Filter by inbox type: "primary" (non-archived, non-low-priority),
// "low-priority", or "archive". If not specified, shows all chats.
type ChatSearchParamsInbox string

const (
	ChatSearchParamsInboxPrimary     ChatSearchParamsInbox = "primary"
	ChatSearchParamsInboxLowPriority ChatSearchParamsInbox = "low-priority"
	ChatSearchParamsInboxArchive     ChatSearchParamsInbox = "archive"
)

// Search scope: 'titles' matches title + network; 'participants' matches
// participant names.
type ChatSearchParamsScope string

const (
	ChatSearchParamsScopeTitles       ChatSearchParamsScope = "titles"
	ChatSearchParamsScopeParticipants ChatSearchParamsScope = "participants"
)

// Specify the type of chats to retrieve: use "single" for direct messages, "group"
// for group chats, or "any" to get all types
type ChatSearchParamsType string

const (
	ChatSearchParamsTypeSingle ChatSearchParamsType = "single"
	ChatSearchParamsTypeGroup  ChatSearchParamsType = "group"
	ChatSearchParamsTypeAny    ChatSearchParamsType = "any"
)

type ChatStartParams struct {
	// Account to create or start the chat on.
	AccountID string `json:"accountID" api:"required"`
	// Merged user-like contact payload used to resolve the best identifier.
	User ChatStartParamsUser `json:"user,omitzero" api:"required"`
	// Whether invite-based DM creation is allowed when required by the platform.
	AllowInvite param.Opt[bool] `json:"allowInvite,omitzero"`
	// Optional first message content if the platform requires it to create the chat.
	MessageText param.Opt[string] `json:"messageText,omitzero"`
	paramObj
}

func (r ChatStartParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Merged user-like contact payload used to resolve the best identifier.
type ChatStartParamsUser struct {
	// Known user ID when available.
	ID param.Opt[string] `json:"id,omitzero"`
	// Email candidate.
	Email param.Opt[string] `json:"email,omitzero"`
	// Display name hint used for ranking only.
	FullName param.Opt[string] `json:"fullName,omitzero"`
	// Phone number candidate (E.164 preferred).
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// Username/handle candidate.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r ChatStartParamsUser) MarshalJSON() (data []byte, err error) {
	type shadow ChatStartParamsUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatStartParamsUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperbeeperdesktopapigo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/beeper/beeper-desktop-api-go/internal/apijson"
	"github.com/beeper/beeper-desktop-api-go/internal/apiquery"
	"github.com/beeper/beeper-desktop-api-go/internal/requestconfig"
	"github.com/beeper/beeper-desktop-api-go/option"
	"github.com/beeper/beeper-desktop-api-go/packages/pagination"
	"github.com/beeper/beeper-desktop-api-go/packages/param"
	"github.com/beeper/beeper-desktop-api-go/packages/respjson"
	"github.com/beeper/beeper-desktop-api-go/shared"
)

// Manage chats, conversations, and threads
//
// ChatService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options []option.RequestOption
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.Options = opts
	return
}

// Retrieve chat details: metadata, participants (limited), and latest message.
//
//   - When to use: fetch a complete view of a chat beyond what search returns.
//   - Constraints: not available for iMessage chats ('imsg##'). Participants limited
//     by 'maxParticipantCount' (default 20, max 500). Returns: chat details.Agents:
//     ALWAYS use linkToChat to make clickable links in your response
func (r *ChatService) Get(ctx context.Context, query ChatGetParams, opts ...option.RequestOption) (res *GetChatResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/get-chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Archive or unarchive a chat. Set archived=true to move to archive,
// archived=false to move back to inbox
func (r *ChatService) Archive(ctx context.Context, body ChatArchiveParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/archive-chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search and filter conversations across all messaging accounts.
//
//   - When to use: browse chats by inbox (primary/low-priority/archive), type,
//     unread status, or search terms.
//   - Pagination: use cursor + direction for pagination.
//   - Performance: provide accountIDs when known for faster filtering. Returns:
//     matching chats with pagination. Agents: ALWAYS use linkToChat to make
//     clickable links in your response
func (r *ChatService) Find(ctx context.Context, query ChatFindParams, opts ...option.RequestOption) (res *pagination.CursorID[Chat], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v0/find-chats"
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

// Search and filter conversations across all messaging accounts.
//
//   - When to use: browse chats by inbox (primary/low-priority/archive), type,
//     unread status, or search terms.
//   - Pagination: use cursor + direction for pagination.
//   - Performance: provide accountIDs when known for faster filtering. Returns:
//     matching chats with pagination. Agents: ALWAYS use linkToChat to make
//     clickable links in your response
func (r *ChatService) FindAutoPaging(ctx context.Context, query ChatFindParams, opts ...option.RequestOption) *pagination.CursorIDAutoPager[Chat] {
	return pagination.NewCursorIDAutoPager(r.Find(ctx, query, opts...))
}

// Generate a deep link to a specific chat or message. This link can be used to
// open the chat directly in the Beeper app.
func (r *ChatService) GetLink(ctx context.Context, body ChatGetLinkParams, opts ...option.RequestOption) (res *LinkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/get-link-to-chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
	Items []shared.User `json:"items,required"`
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

type FindChatsResponse struct {
	// Chats matching the filters.
	Data []Chat `json:"data,required"`
	// Whether there are more items available after this set.
	HasMore bool `json:"has_more,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindChatsResponse) RawJSON() string { return r.JSON.raw }
func (r *FindChatsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetChatResponse struct {
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
	Participants GetChatResponseParticipants `json:"participants,required"`
	// Display title of the chat as computed by the client/server.
	Title string `json:"title,required"`
	// Chat type: 'single' for direct messages, 'group' for group chats, 'channel' for
	// channels, 'broadcast' for broadcasts.
	//
	// Any of "single", "group", "channel", "broadcast".
	Type GetChatResponseType `json:"type,required"`
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
	LastReadMessageSortKey GetChatResponseLastReadMessageSortKeyUnion `json:"lastReadMessageSortKey"`
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
func (r GetChatResponse) RawJSON() string { return r.JSON.raw }
func (r *GetChatResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat participants information.
type GetChatResponseParticipants struct {
	// True if there are more participants than included in items.
	HasMore bool `json:"hasMore,required"`
	// Participants returned for this chat (limited by the request; may be a subset).
	Items []shared.User `json:"items,required"`
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
func (r GetChatResponseParticipants) RawJSON() string { return r.JSON.raw }
func (r *GetChatResponseParticipants) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat type: 'single' for direct messages, 'group' for group chats, 'channel' for
// channels, 'broadcast' for broadcasts.
type GetChatResponseType string

const (
	GetChatResponseTypeSingle    GetChatResponseType = "single"
	GetChatResponseTypeGroup     GetChatResponseType = "group"
	GetChatResponseTypeChannel   GetChatResponseType = "channel"
	GetChatResponseTypeBroadcast GetChatResponseType = "broadcast"
)

// GetChatResponseLastReadMessageSortKeyUnion contains all possible properties and
// values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type GetChatResponseLastReadMessageSortKeyUnion struct {
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

func (u GetChatResponseLastReadMessageSortKeyUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetChatResponseLastReadMessageSortKeyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GetChatResponseLastReadMessageSortKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *GetChatResponseLastReadMessageSortKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URL to open a specific chat or message.
type LinkResponse struct {
	// Deep link URL to the specified chat or message.
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkResponse) RawJSON() string { return r.JSON.raw }
func (r *LinkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetParams struct {
	// Unique identifier of the chat to retrieve. Not available for iMessage chats.
	// Participants are limited by 'maxParticipantCount'.
	ChatID string `query:"chatID,required" json:"-"`
	// Maximum number of participants to return. Use -1 for all; otherwise 0–500.
	// Defaults to 20.
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

type ChatArchiveParams struct {
	// The identifier of the chat to archive or unarchive
	ChatID string `json:"chatID,required"`
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

type ChatFindParams struct {
	// A cursor for use in pagination. ending_before is an object ID that defines your
	// place in the list. For instance, if you make a list request and receive 100
	// objects, starting with obj_bar, your subsequent call can include
	// ending_before=obj_bar in order to fetch the previous page of the list.
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" json:"-"`
	// Include chats marked as Muted by the user, which are usually less important.
	// Default: true. Set to false if the user wants a more refined search.
	IncludeMuted param.Opt[bool] `query:"includeMuted,omitzero" json:"-"`
	// Provide an ISO datetime string to only retrieve chats with last activity after
	// this time
	LastActivityAfter param.Opt[time.Time] `query:"lastActivityAfter,omitzero" format:"date-time" json:"-"`
	// Provide an ISO datetime string to only retrieve chats with last activity before
	// this time
	LastActivityBefore param.Opt[time.Time] `query:"lastActivityBefore,omitzero" format:"date-time" json:"-"`
	// Set the maximum number of chats to retrieve. Valid range: 1-200, default is 50
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search string to filter chats by participant names. When multiple words
	// provided, ALL words must match. Searches in username, displayName, and fullName
	// fields.
	ParticipantQuery param.Opt[string] `query:"participantQuery,omitzero" json:"-"`
	// Search string to filter chats by title. When multiple words provided, ALL words
	// must match. Matches are case-insensitive substrings.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// A cursor for use in pagination. starting_after is an object ID that defines your
	// place in the list. For instance, if you make a list request and receive 100
	// objects, ending with obj_foo, your subsequent call can include
	// starting_after=obj_foo in order to fetch the next page of the list.
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" json:"-"`
	// Set to true to only retrieve chats that have unread messages
	UnreadOnly param.Opt[bool] `query:"unreadOnly,omitzero" json:"-"`
	// Provide an array of account IDs to filter chats from specific messaging accounts
	// only
	AccountIDs []string `query:"accountIDs,omitzero" json:"-"`
	// Filter by inbox type: "primary" (non-archived, non-low-priority),
	// "low-priority", or "archive". If not specified, shows all chats.
	//
	// Any of "primary", "low-priority", "archive".
	Inbox ChatFindParamsInbox `query:"inbox,omitzero" json:"-"`
	// Specify the type of chats to retrieve: use "single" for direct messages, "group"
	// for group chats, "channel" for channels, or "any" to get all types
	//
	// Any of "single", "group", "channel", "any".
	Type ChatFindParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChatFindParams]'s query parameters as `url.Values`.
func (r ChatFindParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by inbox type: "primary" (non-archived, non-low-priority),
// "low-priority", or "archive". If not specified, shows all chats.
type ChatFindParamsInbox string

const (
	ChatFindParamsInboxPrimary     ChatFindParamsInbox = "primary"
	ChatFindParamsInboxLowPriority ChatFindParamsInbox = "low-priority"
	ChatFindParamsInboxArchive     ChatFindParamsInbox = "archive"
)

// Specify the type of chats to retrieve: use "single" for direct messages, "group"
// for group chats, "channel" for channels, or "any" to get all types
type ChatFindParamsType string

const (
	ChatFindParamsTypeSingle  ChatFindParamsType = "single"
	ChatFindParamsTypeGroup   ChatFindParamsType = "group"
	ChatFindParamsTypeChannel ChatFindParamsType = "channel"
	ChatFindParamsTypeAny     ChatFindParamsType = "any"
)

type ChatGetLinkParams struct {
	// The ID of the chat to link to.
	ChatID string `json:"chatID,required"`
	// Optional message sort key. Jumps to that message in the chat.
	MessageSortKey param.Opt[string] `json:"messageSortKey,omitzero"`
	paramObj
}

func (r ChatGetLinkParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatGetLinkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatGetLinkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

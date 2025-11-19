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

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/apiquery"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/pagination"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
	"github.com/beeper/desktop-api-go/shared"
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
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.Options = opts
	r.Reminders = NewChatReminderService(opts...)
	return
}

// Create a single or group chat on a specific account using participant IDs and
// optional title.
func (r *ChatService) New(ctx context.Context, body ChatNewParams, opts ...option.RequestOption) (res *ChatNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve chat details including metadata, participants, and latest message
func (r *ChatService) Get(ctx context.Context, chatID string, query ChatGetParams, opts ...option.RequestOption) (res *Chat, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return
	}
	path := fmt.Sprintf("v1/chats/%s", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
		return
	}
	path := fmt.Sprintf("v1/chats/%s/archive", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Search chats by title/network or participants using Beeper Desktop's renderer
// algorithm.
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

// Search chats by title/network or participants using Beeper Desktop's renderer
// algorithm.
func (r *ChatService) SearchAutoPaging(ctx context.Context, query ChatSearchParams, opts ...option.RequestOption) *pagination.CursorSearchAutoPager[Chat] {
	return pagination.NewCursorSearchAutoPager(r.Search(ctx, query, opts...))
}

type Chat struct {
	// Unique identifier of the chat across Beeper.
	ID string `json:"id,required"`
	// Account ID this chat belongs to.
	AccountID string `json:"accountID,required"`
	// Display-only human-readable network name (e.g., 'WhatsApp', 'Messenger').
	//
	// Deprecated: deprecated
	Network string `json:"network,required"`
	// Chat participants information.
	Participants ChatParticipants `json:"participants,required"`
	// Chat type: 'single' for direct messages, 'group' for group chats.
	//
	// Any of "single", "group".
	Type ChatType `json:"type,required"`
	// Number of unread messages.
	UnreadCount int64 `json:"unreadCount,required"`
	// Description of the chat.
	Description string `json:"description,nullable"`
	// True if chat is archived.
	IsArchived bool `json:"isArchived"`
	// True if the chat is muted.
	IsMuted bool `json:"isMuted"`
	// True if the chat is pinned.
	IsPinned bool `json:"isPinned"`
	// Timestamp of last activity.
	LastActivity time.Time `json:"lastActivity" format:"date-time"`
	// Last read message sortKey.
	LastReadMessageSortKey string `json:"lastReadMessageSortKey"`
	// Local chat ID specific to this Beeper Desktop installation.
	LocalChatID string `json:"localChatID,nullable"`
	// Display title of the chat.
	Title string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccountID              respjson.Field
		Network                respjson.Field
		Participants           respjson.Field
		Type                   respjson.Field
		UnreadCount            respjson.Field
		Description            respjson.Field
		IsArchived             respjson.Field
		IsMuted                respjson.Field
		IsPinned               respjson.Field
		LastActivity           respjson.Field
		LastReadMessageSortKey respjson.Field
		LocalChatID            respjson.Field
		Title                  respjson.Field
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

// Chat type: 'single' for direct messages, 'group' for group chats.
type ChatType string

const (
	ChatTypeSingle ChatType = "single"
	ChatTypeGroup  ChatType = "group"
)

type ChatNewResponse struct {
	// Newly created chat ID.
	ChatID string `json:"chatID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type ChatNewParams struct {
	// Account to create the chat on.
	AccountID string `json:"accountID,required"`
	// User IDs to include in the new chat.
	ParticipantIDs []string `json:"participantIDs,omitzero,required"`
	// Chat type to create: 'single' requires exactly one participantID; 'group'
	// supports multiple participants and optional title.
	//
	// Any of "single", "group".
	Type ChatNewParamsType `json:"type,omitzero,required"`
	// Optional first message content if the platform requires it to create the chat.
	MessageText param.Opt[string] `json:"messageText,omitzero"`
	// Optional title for group chats; ignored for single chats on most platforms.
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

// Chat type to create: 'single' requires exactly one participantID; 'group'
// supports multiple participants and optional title.
type ChatNewParamsType string

const (
	ChatNewParamsTypeSingle ChatNewParamsType = "single"
	ChatNewParamsTypeGroup  ChatNewParamsType = "group"
)

type ChatGetParams struct {
	// Maximum number of participants to return. Use -1 for all; otherwise 0–500.
	// Defaults to all (-1).
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

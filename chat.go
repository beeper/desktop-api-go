// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
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

// Chats operations
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

// Archive or unarchive a chat. Set archived=true to move to archive,
// archived=false to move back to inbox
func (r *ChatService) Archive(ctx context.Context, body ChatArchiveParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/archive-chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve chat details including metadata, participants, and latest message
func (r *ChatService) Get(ctx context.Context, query ChatGetParams, opts ...option.RequestOption) (res *ChatGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/get-chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Search and filter conversations across all messaging accounts
func (r *ChatService) Search(ctx context.Context, query ChatSearchParams, opts ...option.RequestOption) (res *pagination.CursorID[shared.Chat], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v0/search-chats"
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

// Search and filter conversations across all messaging accounts
func (r *ChatService) SearchAutoPaging(ctx context.Context, query ChatSearchParams, opts ...option.RequestOption) *pagination.CursorIDAutoPager[shared.Chat] {
	return pagination.NewCursorIDAutoPager(r.Search(ctx, query, opts...))
}

type ChatGetResponse struct {
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
	Participants ChatGetResponseParticipants `json:"participants,required"`
	// Display title of the chat as computed by the client/server.
	Title string `json:"title,required"`
	// Chat type: 'single' for direct messages, 'group' for group chats, 'channel' for
	// channels, 'broadcast' for broadcasts.
	//
	// Any of "single", "group", "channel", "broadcast".
	Type ChatGetResponseType `json:"type,required"`
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
	LastReadMessageSortKey ChatGetResponseLastReadMessageSortKeyUnion `json:"lastReadMessageSortKey"`
	// Local numeric chat ID specific to this Beeper Desktop installation. null for
	// iMessage chats.
	LocalChatID string `json:"localChatID,nullable"`
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
		LocalChatID            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat participants information.
type ChatGetResponseParticipants struct {
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
func (r ChatGetResponseParticipants) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseParticipants) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat type: 'single' for direct messages, 'group' for group chats, 'channel' for
// channels, 'broadcast' for broadcasts.
type ChatGetResponseType string

const (
	ChatGetResponseTypeSingle    ChatGetResponseType = "single"
	ChatGetResponseTypeGroup     ChatGetResponseType = "group"
	ChatGetResponseTypeChannel   ChatGetResponseType = "channel"
	ChatGetResponseTypeBroadcast ChatGetResponseType = "broadcast"
)

// ChatGetResponseLastReadMessageSortKeyUnion contains all possible properties and
// values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ChatGetResponseLastReadMessageSortKeyUnion struct {
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

func (u ChatGetResponseLastReadMessageSortKeyUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatGetResponseLastReadMessageSortKeyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatGetResponseLastReadMessageSortKeyUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatGetResponseLastReadMessageSortKeyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type ChatSearchParams struct {
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
	Inbox ChatSearchParamsInbox `query:"inbox,omitzero" json:"-"`
	// Specify the type of chats to retrieve: use "single" for direct messages, "group"
	// for group chats, "channel" for channels, or "any" to get all types
	//
	// Any of "single", "group", "channel", "any".
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

// Filter by inbox type: "primary" (non-archived, non-low-priority),
// "low-priority", or "archive". If not specified, shows all chats.
type ChatSearchParamsInbox string

const (
	ChatSearchParamsInboxPrimary     ChatSearchParamsInbox = "primary"
	ChatSearchParamsInboxLowPriority ChatSearchParamsInbox = "low-priority"
	ChatSearchParamsInboxArchive     ChatSearchParamsInbox = "archive"
)

// Specify the type of chats to retrieve: use "single" for direct messages, "group"
// for group chats, "channel" for channels, or "any" to get all types
type ChatSearchParamsType string

const (
	ChatSearchParamsTypeSingle  ChatSearchParamsType = "single"
	ChatSearchParamsTypeGroup   ChatSearchParamsType = "group"
	ChatSearchParamsTypeChannel ChatSearchParamsType = "channel"
	ChatSearchParamsTypeAny     ChatSearchParamsType = "any"
)

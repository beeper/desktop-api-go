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

// Beeper Desktop API v0
//
// V0Service contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0Service] method instead.
type V0Service struct {
	Options []option.RequestOption
}

// NewV0Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV0Service(opts ...option.RequestOption) (r V0Service) {
	r = V0Service{}
	r.Options = opts
	return
}

// Archive or unarchive a chat. Set archived=true to move to archive,
// archived=false to move back to inbox
func (r *V0Service) ArchiveChat(ctx context.Context, body V0ArchiveChatParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/archive-chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Clear an existing reminder from a chat
func (r *V0Service) ClearChatReminder(ctx context.Context, body V0ClearChatReminderParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/clear-chat-reminder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Draft a message in a specific chat. This will be placed in the message input
// field without sending
func (r *V0Service) DraftMessage(ctx context.Context, body V0DraftMessageParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/draft-message"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search and filter conversations across all messaging accounts
func (r *V0Service) FindChats(ctx context.Context, query V0FindChatsParams, opts ...option.RequestOption) (res *pagination.CursorID[shared.Chat], err error) {
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

// Search and filter conversations across all messaging accounts
func (r *V0Service) FindChatsAutoPaging(ctx context.Context, query V0FindChatsParams, opts ...option.RequestOption) *pagination.CursorIDAutoPager[shared.Chat] {
	return pagination.NewCursorIDAutoPager(r.FindChats(ctx, query, opts...))
}

// Bring Beeper Desktop to the foreground on this device
func (r *V0Service) FocusApp(ctx context.Context, body V0FocusAppParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/focus-app"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List connected Beeper accounts available on this device
func (r *V0Service) GetAccounts(ctx context.Context, opts ...option.RequestOption) (res *AccountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/get-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve chat details including metadata, participants, and latest message
func (r *V0Service) GetChat(ctx context.Context, query V0GetChatParams, opts ...option.RequestOption) (res *GetChatResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/get-chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Generate a deep link to a specific chat or message. This link can be used to
// open the chat directly in the Beeper app.
func (r *V0Service) GetLinkToChat(ctx context.Context, body V0GetLinkToChatParams, opts ...option.RequestOption) (res *LinkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/get-link-to-chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search messages across chats using Beeper's message index
func (r *V0Service) SearchMessages(ctx context.Context, query V0SearchMessagesParams, opts ...option.RequestOption) (res *pagination.CursorID[shared.Message], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v0/search-messages"
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

// Search messages across chats using Beeper's message index
func (r *V0Service) SearchMessagesAutoPaging(ctx context.Context, query V0SearchMessagesParams, opts ...option.RequestOption) *pagination.CursorIDAutoPager[shared.Message] {
	return pagination.NewCursorIDAutoPager(r.SearchMessages(ctx, query, opts...))
}

// Send a text message to a specific chat. Supports replying to existing messages.
// Returns the sent message ID and a deeplink to the chat
func (r *V0Service) SendMessage(ctx context.Context, body V0SendMessageParams, opts ...option.RequestOption) (res *SendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/send-message"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Set a reminder for a chat at a specific time
func (r *V0Service) SetChatReminder(ctx context.Context, body V0SetChatReminderParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/set-chat-reminder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response payload for listing connected Beeper accounts.
type AccountsResponse struct {
	// Connected accounts the user can act through. Includes accountID, network, and
	// user identity.
	Accounts []shared.Account `json:"accounts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindChatsResponse struct {
	// Chats matching the filters.
	Data []shared.Chat `json:"data,required"`
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

type SearchResponse struct {
	// Map of chatID -> chat details for chats referenced in data.
	Chats map[string]shared.Chat `json:"chats,required"`
	// Messages matching the query and filters.
	Data []shared.Message `json:"data,required"`
	// Whether there are more items available after this set.
	HasMore bool `json:"has_more,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chats       respjson.Field
		Data        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendResponse struct {
	// Link to the chat where the message was sent. This should always be shown to the
	// user.
	Deeplink string `json:"deeplink,required"`
	// Stable message ID.
	MessageID string `json:"messageID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deeplink    respjson.Field
		MessageID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r SendResponse) RawJSON() string { return r.JSON.raw }
func (r *SendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0ArchiveChatParams struct {
	// The identifier of the chat to archive or unarchive
	ChatID string `json:"chatID,required"`
	// True to archive, false to unarchive
	Archived param.Opt[bool] `json:"archived,omitzero"`
	paramObj
}

func (r V0ArchiveChatParams) MarshalJSON() (data []byte, err error) {
	type shadow V0ArchiveChatParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0ArchiveChatParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0ClearChatReminderParams struct {
	// The identifier of the chat to clear reminder from
	ChatID string `json:"chatID,required"`
	paramObj
}

func (r V0ClearChatReminderParams) MarshalJSON() (data []byte, err error) {
	type shadow V0ClearChatReminderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0ClearChatReminderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0DraftMessageParams struct {
	// Provide the unique identifier of the chat where you want to draft a message
	ChatID string `json:"chatID,required"`
	// Set to true to bring Beeper application to the foreground, or false to draft
	// silently in background
	FocusApp param.Opt[bool] `json:"focusApp,omitzero"`
	// Provide the text content you want to draft. This will be placed in the message
	// input field without sending
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r V0DraftMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow V0DraftMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0DraftMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0FindChatsParams struct {
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
	Inbox V0FindChatsParamsInbox `query:"inbox,omitzero" json:"-"`
	// Specify the type of chats to retrieve: use "single" for direct messages, "group"
	// for group chats, "channel" for channels, or "any" to get all types
	//
	// Any of "single", "group", "channel", "any".
	Type V0FindChatsParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V0FindChatsParams]'s query parameters as `url.Values`.
func (r V0FindChatsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by inbox type: "primary" (non-archived, non-low-priority),
// "low-priority", or "archive". If not specified, shows all chats.
type V0FindChatsParamsInbox string

const (
	V0FindChatsParamsInboxPrimary     V0FindChatsParamsInbox = "primary"
	V0FindChatsParamsInboxLowPriority V0FindChatsParamsInbox = "low-priority"
	V0FindChatsParamsInboxArchive     V0FindChatsParamsInbox = "archive"
)

// Specify the type of chats to retrieve: use "single" for direct messages, "group"
// for group chats, "channel" for channels, or "any" to get all types
type V0FindChatsParamsType string

const (
	V0FindChatsParamsTypeSingle  V0FindChatsParamsType = "single"
	V0FindChatsParamsTypeGroup   V0FindChatsParamsType = "group"
	V0FindChatsParamsTypeChannel V0FindChatsParamsType = "channel"
	V0FindChatsParamsTypeAny     V0FindChatsParamsType = "any"
)

type V0FocusAppParams struct {
	// Optional Beeper chat ID to focus after bringing the app to foreground. If
	// omitted, only foregrounds the app. Required if messageSortKey is present. No-op
	// in headless environments.
	ChatID param.Opt[string] `json:"chatID,omitzero"`
	// Optional message sort key. Jumps to that message in the chat when foregrounding.
	MessageSortKey param.Opt[string] `json:"messageSortKey,omitzero"`
	paramObj
}

func (r V0FocusAppParams) MarshalJSON() (data []byte, err error) {
	type shadow V0FocusAppParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0FocusAppParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0GetChatParams struct {
	// Unique identifier of the chat to retrieve. Not available for iMessage chats.
	// Participants are limited by 'maxParticipantCount'.
	ChatID string `query:"chatID,required" json:"-"`
	// Maximum number of participants to return. Use -1 for all; otherwise 0–500.
	// Defaults to 20.
	MaxParticipantCount param.Opt[int64] `query:"maxParticipantCount,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V0GetChatParams]'s query parameters as `url.Values`.
func (r V0GetChatParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V0GetLinkToChatParams struct {
	// The ID of the chat to link to.
	ChatID string `json:"chatID,required"`
	// Optional message sort key. Jumps to that message in the chat.
	MessageSortKey param.Opt[string] `json:"messageSortKey,omitzero"`
	paramObj
}

func (r V0GetLinkToChatParams) MarshalJSON() (data []byte, err error) {
	type shadow V0GetLinkToChatParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0GetLinkToChatParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0SearchMessagesParams struct {
	// Only include messages with timestamp strictly after this ISO 8601 datetime
	// (e.g., '2024-07-01T00:00:00Z' or '2024-07-01T00:00:00+02:00').
	DateAfter param.Opt[time.Time] `query:"dateAfter,omitzero" format:"date-time" json:"-"`
	// Only include messages with timestamp strictly before this ISO 8601 datetime
	// (e.g., '2024-07-31T23:59:59Z' or '2024-07-31T23:59:59+02:00').
	DateBefore param.Opt[time.Time] `query:"dateBefore,omitzero" format:"date-time" json:"-"`
	// A cursor for use in pagination. ending_before is an object ID that defines your
	// place in the list. For instance, if you make a list request and receive 100
	// objects, starting with obj_bar, your subsequent call can include
	// ending_before=obj_bar in order to fetch the previous page of the list.
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" json:"-"`
	// Exclude messages marked Low Priority by the user. Default: true. Set to false to
	// include all.
	ExcludeLowPriority param.Opt[bool] `query:"excludeLowPriority,omitzero" json:"-"`
	// Include messages in chats marked as Muted by the user, which are usually less
	// important. Default: true. Set to false if the user wants a more refined search.
	IncludeMuted param.Opt[bool] `query:"includeMuted,omitzero" json:"-"`
	// Maximum number of messages to return (1–500). Defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return messages that contain file attachments.
	OnlyWithFile param.Opt[bool] `query:"onlyWithFile,omitzero" json:"-"`
	// Only return messages that contain image attachments.
	OnlyWithImage param.Opt[bool] `query:"onlyWithImage,omitzero" json:"-"`
	// Only return messages that contain link attachments.
	OnlyWithLink param.Opt[bool] `query:"onlyWithLink,omitzero" json:"-"`
	// Only return messages that contain any type of media attachment.
	OnlyWithMedia param.Opt[bool] `query:"onlyWithMedia,omitzero" json:"-"`
	// Only return messages that contain video attachments.
	OnlyWithVideo param.Opt[bool] `query:"onlyWithVideo,omitzero" json:"-"`
	// Literal word search (NOT semantic). Finds messages containing these EXACT words
	// in any order. Use single words users actually type, not concepts or phrases.
	// Example: use "dinner" not "dinner plans", use "sick" not "health issues". If
	// omitted, returns results filtered only by other parameters.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// A cursor for use in pagination. starting_after is an object ID that defines your
	// place in the list. For instance, if you make a list request and receive 100
	// objects, ending with obj_foo, your subsequent call can include
	// starting_after=obj_foo in order to fetch the next page of the list.
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" json:"-"`
	// Limit search to specific Beeper account IDs (bridge instances).
	AccountIDs []string `query:"accountIDs,omitzero" json:"-"`
	// Limit search to specific Beeper chat IDs.
	ChatIDs []string `query:"chatIDs,omitzero" json:"-"`
	// Filter by chat type: 'group' for group chats, 'single' for 1:1 chats.
	//
	// Any of "group", "single".
	ChatType V0SearchMessagesParamsChatType `query:"chatType,omitzero" json:"-"`
	// Filter by sender: 'me' (messages sent by the authenticated user), 'others'
	// (messages sent by others), or a specific user ID string (user.id).
	Sender V0SearchMessagesParamsSender `query:"sender,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V0SearchMessagesParams]'s query parameters as `url.Values`.
func (r V0SearchMessagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by chat type: 'group' for group chats, 'single' for 1:1 chats.
type V0SearchMessagesParamsChatType string

const (
	V0SearchMessagesParamsChatTypeGroup  V0SearchMessagesParamsChatType = "group"
	V0SearchMessagesParamsChatTypeSingle V0SearchMessagesParamsChatType = "single"
)

// Filter by sender: 'me' (messages sent by the authenticated user), 'others'
// (messages sent by others), or a specific user ID string (user.id).
type V0SearchMessagesParamsSender string

const (
	V0SearchMessagesParamsSenderMe     V0SearchMessagesParamsSender = "me"
	V0SearchMessagesParamsSenderOthers V0SearchMessagesParamsSender = "others"
)

type V0SendMessageParams struct {
	// The identifier of the chat where the message will send
	ChatID string `json:"chatID,required"`
	// Provide a message ID to send this as a reply to an existing message
	ReplyToMessageID param.Opt[string] `json:"replyToMessageID,omitzero"`
	// Text content of the message you want to send. You may use markdown.
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r V0SendMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow V0SendMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0SendMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0SetChatReminderParams struct {
	// The identifier of the chat to set reminder for
	ChatID string `json:"chatID,required"`
	// Reminder configuration
	Reminder V0SetChatReminderParamsReminder `json:"reminder,omitzero,required"`
	paramObj
}

func (r V0SetChatReminderParams) MarshalJSON() (data []byte, err error) {
	type shadow V0SetChatReminderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0SetChatReminderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reminder configuration
//
// The property RemindAtMs is required.
type V0SetChatReminderParamsReminder struct {
	// Unix timestamp in milliseconds when reminder should trigger
	RemindAtMs float64 `json:"remindAtMs,required"`
	// Cancel reminder if someone messages in the chat
	DismissOnIncomingMessage param.Opt[bool] `json:"dismissOnIncomingMessage,omitzero"`
	paramObj
}

func (r V0SetChatReminderParamsReminder) MarshalJSON() (data []byte, err error) {
	type shadow V0SetChatReminderParamsReminder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0SetChatReminderParamsReminder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

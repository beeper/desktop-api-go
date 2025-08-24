// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperbeeperdesktopapigo

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

// Send, draft, and search messages across all chat networks
//
// MessageService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	return
}

// Draft a message in a specific chat. This will be placed in the message input
// field without sending
func (r *MessageService) Draft(ctx context.Context, body MessageDraftParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/draft-message"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search messages across chats using Beeper's message index.
//
//   - When to use: find messages by text and/or filters (chatIDs, accountIDs,
//     chatType, media type filters, sender, date ranges).
//   - CRITICAL: Query is LITERAL WORD MATCHING, NOT semantic search! Only finds
//     messages containing these EXACT words. • ✅ RIGHT: query="dinner" or
//     query="sick" or query="error" (single words users type) • ❌ WRONG:
//     query="dinner plans tonight" or query="health issues" (phrases/concepts) • The
//     query matches ALL words provided (in any order). Example: query="flight
//     booking" finds messages with both "flight" AND "booking".
//   - Media filters: Use onlyWithMedia for any media, or specific filters like
//     onlyWithVideo, onlyWithImage, onlyWithLink, onlyWithFile for specific types.
//   - Pagination: use 'oldestCursor' + direction='before' for older;
//     'newestCursor' + direction='after' for newer.
//   - Performance: provide chatIDs/accountIDs when known. Omitted 'query' returns
//     results based on filters only. Partial matches enabled; 'excludeLowPriority'
//     defaults to true.
//   - Workflow tip: To search messages in specific conversations: 1) Use find-chats
//     to get chatIDs, 2) Use search-messages with those chatIDs.
//   - IMPORTANT: Chat names vary widely. ASK the user for clarification: • "Which
//     chat do you mean by family?" (could be "The Smiths", "Mom Dad Kids", etc.) •
//     "What's the name of your work chat?" (could be "Team", company name, project
//     name) • "Who are the participants?" (use participantQuery in find-chats)
//     Returns: matching messages and referenced chats.
func (r *MessageService) Search(ctx context.Context, query MessageSearchParams, opts ...option.RequestOption) (res *pagination.CursorID[Message], err error) {
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

// Search messages across chats using Beeper's message index.
//
//   - When to use: find messages by text and/or filters (chatIDs, accountIDs,
//     chatType, media type filters, sender, date ranges).
//   - CRITICAL: Query is LITERAL WORD MATCHING, NOT semantic search! Only finds
//     messages containing these EXACT words. • ✅ RIGHT: query="dinner" or
//     query="sick" or query="error" (single words users type) • ❌ WRONG:
//     query="dinner plans tonight" or query="health issues" (phrases/concepts) • The
//     query matches ALL words provided (in any order). Example: query="flight
//     booking" finds messages with both "flight" AND "booking".
//   - Media filters: Use onlyWithMedia for any media, or specific filters like
//     onlyWithVideo, onlyWithImage, onlyWithLink, onlyWithFile for specific types.
//   - Pagination: use 'oldestCursor' + direction='before' for older;
//     'newestCursor' + direction='after' for newer.
//   - Performance: provide chatIDs/accountIDs when known. Omitted 'query' returns
//     results based on filters only. Partial matches enabled; 'excludeLowPriority'
//     defaults to true.
//   - Workflow tip: To search messages in specific conversations: 1) Use find-chats
//     to get chatIDs, 2) Use search-messages with those chatIDs.
//   - IMPORTANT: Chat names vary widely. ASK the user for clarification: • "Which
//     chat do you mean by family?" (could be "The Smiths", "Mom Dad Kids", etc.) •
//     "What's the name of your work chat?" (could be "Team", company name, project
//     name) • "Who are the participants?" (use participantQuery in find-chats)
//     Returns: matching messages and referenced chats.
func (r *MessageService) SearchAutoPaging(ctx context.Context, query MessageSearchParams, opts ...option.RequestOption) *pagination.CursorIDAutoPager[Message] {
	return pagination.NewCursorIDAutoPager(r.Search(ctx, query, opts...))
}

// Send a text message to a specific chat. Supports replying to existing messages.
// Returns the sent message ID and a deeplink to the chat
func (r *MessageService) Send(ctx context.Context, body MessageSendParams, opts ...option.RequestOption) (res *SendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/send-message"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
	Attachments []shared.Attachment `json:"attachments"`
	// True if the authenticated user sent the message.
	IsSender bool `json:"isSender"`
	// True if the message is unread for the authenticated user. May be omitted.
	IsUnread bool `json:"isUnread"`
	// Reactions to the message, if any.
	Reactions []shared.Reaction `json:"reactions"`
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

type SearchResponse struct {
	// Map of chatID -> chat details for chats referenced in data.
	Chats map[string]Chat `json:"chats,required"`
	// Messages matching the query and filters.
	Data []Message `json:"data,required"`
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

type MessageDraftParams struct {
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

func (r MessageDraftParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageDraftParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageDraftParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSearchParams struct {
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
	ChatType MessageSearchParamsChatType `query:"chatType,omitzero" json:"-"`
	// Filter by sender: 'me' (messages sent by the authenticated user), 'others'
	// (messages sent by others), or a specific user ID string (user.id).
	Sender MessageSearchParamsSender `query:"sender,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageSearchParams]'s query parameters as `url.Values`.
func (r MessageSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by chat type: 'group' for group chats, 'single' for 1:1 chats.
type MessageSearchParamsChatType string

const (
	MessageSearchParamsChatTypeGroup  MessageSearchParamsChatType = "group"
	MessageSearchParamsChatTypeSingle MessageSearchParamsChatType = "single"
)

// Filter by sender: 'me' (messages sent by the authenticated user), 'others'
// (messages sent by others), or a specific user ID string (user.id).
type MessageSearchParamsSender string

const (
	MessageSearchParamsSenderMe     MessageSearchParamsSender = "me"
	MessageSearchParamsSenderOthers MessageSearchParamsSender = "others"
)

type MessageSendParams struct {
	// The identifier of the chat where the message will send
	ChatID string `json:"chatID,required"`
	// Provide a message ID to send this as a reply to an existing message
	ReplyToMessageID param.Opt[string] `json:"replyToMessageID,omitzero"`
	// Text content of the message you want to send. You may use markdown.
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r MessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

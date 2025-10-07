// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
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

// Messages operations
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

// List all messages in a chat with cursor-based pagination. Sorted by timestamp.
func (r *MessageService) List(ctx context.Context, query MessageListParams, opts ...option.RequestOption) (res *pagination.Cursor[shared.Message], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/messages"
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

// List all messages in a chat with cursor-based pagination. Sorted by timestamp.
func (r *MessageService) ListAutoPaging(ctx context.Context, query MessageListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[shared.Message] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Search messages across chats using Beeper's message index
func (r *MessageService) Search(ctx context.Context, query MessageSearchParams, opts ...option.RequestOption) (res *MessageSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Send a text message to a specific chat. Supports replying to existing messages.
// Returns the sent message ID.
func (r *MessageService) Send(ctx context.Context, body MessageSendParams, opts ...option.RequestOption) (res *MessageSendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MessageSearchResponse struct {
	// Map of chatID -> chat details for chats referenced in items.
	Chats map[string]Chat `json:"chats,required"`
	// True if additional results can be fetched using the provided cursors.
	HasMore bool `json:"hasMore,required"`
	// Messages matching the query and filters.
	Items []shared.Message `json:"items,required"`
	// Cursor for fetching newer results (use with direction='after'). Opaque string;
	// do not inspect.
	NewestCursor string `json:"newestCursor,required"`
	// Cursor for fetching older results (use with direction='before'). Opaque string;
	// do not inspect.
	OldestCursor string `json:"oldestCursor,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chats        respjson.Field
		HasMore      respjson.Field
		Items        respjson.Field
		NewestCursor respjson.Field
		OldestCursor respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendResponse struct {
	// Unique identifier of the chat.
	ChatID string `json:"chatID,required"`
	// Pending message ID
	PendingMessageID string `json:"pendingMessageID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID           respjson.Field
		PendingMessageID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r MessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageListParams struct {
	// The chat ID to list messages from
	ChatID string `query:"chatID,required" json:"-"`
	// Message cursor for pagination. Use with direction to navigate results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of messages to return (1–500). Defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination direction used with 'cursor': 'before' fetches older messages,
	// 'after' fetches newer messages. Defaults to 'before' when only 'cursor' is
	// provided.
	//
	// Any of "after", "before".
	Direction MessageListParamsDirection `query:"direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageListParams]'s query parameters as `url.Values`.
func (r MessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Pagination direction used with 'cursor': 'before' fetches older messages,
// 'after' fetches newer messages. Defaults to 'before' when only 'cursor' is
// provided.
type MessageListParamsDirection string

const (
	MessageListParamsDirectionAfter  MessageListParamsDirection = "after"
	MessageListParamsDirectionBefore MessageListParamsDirection = "before"
)

type MessageSearchParams struct {
	// Exclude messages marked Low Priority by the user. Default: true. Set to false to
	// include all.
	ExcludeLowPriority param.Opt[bool] `query:"excludeLowPriority,omitzero" json:"-"`
	// Include messages in chats marked as Muted by the user, which are usually less
	// important. Default: true. Set to false if the user wants a more refined search.
	IncludeMuted param.Opt[bool] `query:"includeMuted,omitzero" json:"-"`
	// Opaque pagination cursor; do not inspect. Use together with 'direction'.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only include messages with timestamp strictly after this ISO 8601 datetime
	// (e.g., '2024-07-01T00:00:00Z' or '2024-07-01T00:00:00+02:00').
	DateAfter param.Opt[time.Time] `query:"dateAfter,omitzero" format:"date-time" json:"-"`
	// Only include messages with timestamp strictly before this ISO 8601 datetime
	// (e.g., '2024-07-31T23:59:59Z' or '2024-07-31T23:59:59+02:00').
	DateBefore param.Opt[time.Time] `query:"dateBefore,omitzero" format:"date-time" json:"-"`
	// Maximum number of messages to return (1–500). Defaults to 20. The current
	// implementation caps each page at 20 items even if a higher limit is requested.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Literal word search (NOT semantic). Finds messages containing these EXACT words
	// in any order. Use single words users actually type, not concepts or phrases.
	// Example: use "dinner" not "dinner plans", use "sick" not "health issues". If
	// omitted, returns results filtered only by other parameters.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Limit search to specific account IDs.
	AccountIDs []string `query:"accountIDs,omitzero" json:"-"`
	// Limit search to specific chat IDs.
	ChatIDs []string `query:"chatIDs,omitzero" json:"-"`
	// Filter by chat type: 'group' for group chats, 'single' for 1:1 chats.
	//
	// Any of "group", "single".
	ChatType MessageSearchParamsChatType `query:"chatType,omitzero" json:"-"`
	// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
	// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
	//
	// Any of "after", "before".
	Direction MessageSearchParamsDirection `query:"direction,omitzero" json:"-"`
	// Filter messages by media types. Use ['any'] for any media type, or specify exact
	// types like ['video', 'image']. Omit for no media filtering.
	//
	// Any of "any", "video", "image", "link", "file".
	MediaTypes []string `query:"mediaTypes,omitzero" json:"-"`
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

// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
type MessageSearchParamsDirection string

const (
	MessageSearchParamsDirectionAfter  MessageSearchParamsDirection = "after"
	MessageSearchParamsDirectionBefore MessageSearchParamsDirection = "before"
)

// Filter by sender: 'me' (messages sent by the authenticated user), 'others'
// (messages sent by others), or a specific user ID string (user.id).
type MessageSearchParamsSender string

const (
	MessageSearchParamsSenderMe     MessageSearchParamsSender = "me"
	MessageSearchParamsSenderOthers MessageSearchParamsSender = "others"
)

type MessageSendParams struct {
	// Unique identifier of the chat.
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

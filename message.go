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

// Manage messages in chats
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

// Retrieve a message by final message ID, pendingMessageID, or Matrix event ID.
// Chat ID may be a Beeper chat ID or local chat ID.
func (r *MessageService) Get(ctx context.Context, messageID string, query MessageGetParams, opts ...option.RequestOption) (res *shared.Message, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ChatID == "" {
		err = errors.New("missing required chatID parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/chats/%s/messages/%s", query.ChatID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Edit the text content of an existing message. Messages with attachments cannot
// be edited.
func (r *MessageService) Update(ctx context.Context, messageID string, params MessageUpdateParams, opts ...option.RequestOption) (res *MessageUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ChatID == "" {
		err = errors.New("missing required chatID parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/chats/%s/messages/%s", params.ChatID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all messages in a chat with cursor-based pagination. Sorted by timestamp.
func (r *MessageService) List(ctx context.Context, chatID string, query MessageListParams, opts ...option.RequestOption) (res *pagination.CursorNoLimit[shared.Message], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/chats/%s/messages", chatID)
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
func (r *MessageService) ListAutoPaging(ctx context.Context, chatID string, query MessageListParams, opts ...option.RequestOption) *pagination.CursorNoLimitAutoPager[shared.Message] {
	return pagination.NewCursorNoLimitAutoPager(r.List(ctx, chatID, query, opts...))
}

// Delete a message by final message ID. Pending message IDs are not accepted
// because messages cannot be deleted while sending.
func (r *MessageService) Delete(ctx context.Context, messageID string, params MessageDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ChatID == "" {
		err = errors.New("missing required chatID parameter")
		return err
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return err
	}
	path := fmt.Sprintf("v1/chats/%s/messages/%s", params.ChatID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Search messages across chats.
func (r *MessageService) Search(ctx context.Context, query MessageSearchParams, opts ...option.RequestOption) (res *pagination.CursorSearch[shared.Message], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/messages/search"
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

// Search messages across chats.
func (r *MessageService) SearchAutoPaging(ctx context.Context, query MessageSearchParams, opts ...option.RequestOption) *pagination.CursorSearchAutoPager[shared.Message] {
	return pagination.NewCursorSearchAutoPager(r.Search(ctx, query, opts...))
}

// Send a text message to a specific chat. Supports replying to existing messages.
// Returns a pending message ID.
func (r *MessageService) Send(ctx context.Context, chatID string, body MessageSendParams, opts ...option.RequestOption) (res *MessageSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/chats/%s/messages", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MessageUpdateResponse struct {
	// DEPRECATED - use id instead. Compatibility alias for older clients.
	//
	// Deprecated: deprecated
	MessageID string `json:"messageID" api:"required"`
	// DEPRECATED - compatibility field. Successful responses are already represented
	// by the 200 status code.
	//
	// Any of true.
	//
	// Deprecated: deprecated
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.Message
}

// Returns the unmodified JSON received from the API
func (r MessageUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendResponse struct {
	// Chat ID. Input routes also accept the local chat ID from this Beeper Desktop
	// installation when available.
	ChatID string `json:"chatID" api:"required"`
	// Pending ID assigned to the message before the network confirms the send. Pass it
	// to GET /v1/chats/{chatID}/messages/{messageID} to resolve, or wait for the
	// matching message.upserted over the WebSocket.
	PendingMessageID string `json:"pendingMessageID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID           respjson.Field
		PendingMessageID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetParams struct {
	// Chat ID. Input routes also accept the local chat ID from this Beeper Desktop
	// installation when available.
	ChatID string `path:"chatID" api:"required" json:"-"`
	paramObj
}

type MessageUpdateParams struct {
	// Chat ID. Input routes also accept the local chat ID from this Beeper Desktop
	// installation when available.
	ChatID string `path:"chatID" api:"required" json:"-"`
	// New text content for the message
	Text string `json:"text" api:"required"`
	paramObj
}

func (r MessageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageListParams struct {
	// Opaque pagination cursor; do not inspect. Use together with 'direction'.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
	// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
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

// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
type MessageListParamsDirection string

const (
	MessageListParamsDirectionAfter  MessageListParamsDirection = "after"
	MessageListParamsDirectionBefore MessageListParamsDirection = "before"
)

type MessageDeleteParams struct {
	// Chat ID. Input routes also accept the local chat ID from this Beeper Desktop
	// installation when available.
	ChatID string `path:"chatID" api:"required" json:"-"`
	// True to request deletion for everyone when the network supports it; false to
	// delete only for the authenticated user when supported.
	ForEveryone param.Opt[bool] `query:"forEveryone,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageDeleteParams]'s query parameters as `url.Values`.
func (r MessageDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
	// Maximum number of messages to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Literal word search (non-semantic). Finds messages containing these EXACT words
	// in any order. Use single words users actually type, not concepts or phrases.
	// Example: use "dinner" not "dinner plans", use "sick" not "health issues". If
	// omitted, returns results filtered only by other parameters.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Filter by sender: 'me' (messages sent by the authenticated user), 'others'
	// (messages sent by others), or a specific user ID string (user.id).
	Sender param.Opt[string] `query:"sender,omitzero" json:"-"`
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

type MessageSendParams struct {
	// Provide a message ID to send this as a reply to an existing message
	ReplyToMessageID param.Opt[string] `json:"replyToMessageID,omitzero"`
	// Draft text. Plain text and Markdown are converted to Matrix HTML with the same
	// rules used by send and edit.
	Text param.Opt[string] `json:"text,omitzero"`
	// Single attachment to send with the message
	Attachment MessageSendParamsAttachment `json:"attachment,omitzero"`
	paramObj
}

func (r MessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single attachment to send with the message
//
// The property UploadID is required.
type MessageSendParamsAttachment struct {
	// Upload ID from uploadAsset endpoint. Required to reference uploaded files.
	UploadID string `json:"uploadID" api:"required"`
	// Duration in seconds (optional override of cached value)
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Filename (optional override of cached value)
	FileName param.Opt[string] `json:"fileName,omitzero"`
	// MIME type (optional override of cached value)
	MimeType param.Opt[string] `json:"mimeType,omitzero"`
	// Dimensions (optional override of cached value)
	Size MessageSendParamsAttachmentSize `json:"size,omitzero"`
	// Attachment type hint (image, video, audio, file, gif, voice-note, sticker). If
	// omitted, auto-detected from mimeType
	//
	// Any of "image", "video", "audio", "file", "gif", "voice-note", "sticker".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r MessageSendParamsAttachment) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParamsAttachment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParamsAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageSendParamsAttachment](
		"type", "image", "video", "audio", "file", "gif", "voice-note", "sticker",
	)
}

// Dimensions (optional override of cached value)
//
// The properties Height, Width are required.
type MessageSendParamsAttachmentSize struct {
	Height float64 `json:"height" api:"required"`
	Width  float64 `json:"width" api:"required"`
	paramObj
}

func (r MessageSendParamsAttachmentSize) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParamsAttachmentSize
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParamsAttachmentSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

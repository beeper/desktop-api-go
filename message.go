// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo

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

// Download an attachment from a message and return the local file path
func (r *MessageService) GetAttachment(ctx context.Context, body MessageGetAttachmentParams, opts ...option.RequestOption) (res *MessageGetAttachmentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/get-attachment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search messages across chats using Beeper's message index
func (r *MessageService) Search(ctx context.Context, query MessageSearchParams, opts ...option.RequestOption) (res *pagination.CursorID[shared.Message], err error) {
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
func (r *MessageService) SearchAutoPaging(ctx context.Context, query MessageSearchParams, opts ...option.RequestOption) *pagination.CursorIDAutoPager[shared.Message] {
	return pagination.NewCursorIDAutoPager(r.Search(ctx, query, opts...))
}

// Send a text message to a specific chat. Supports replying to existing messages.
// Returns the sent message ID and a deeplink to the chat
func (r *MessageService) Send(ctx context.Context, body MessageSendParams, opts ...option.RequestOption) (res *MessageSendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/send-message"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MessageGetAttachmentResponse struct {
	// Whether the attachment was successfully downloaded.
	Success bool `json:"success,required"`
	// Error message if the download failed.
	Error string `json:"error"`
	// Local file system path to the downloaded attachment.
	FilePath string `json:"filePath"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		Error       respjson.Field
		FilePath    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetAttachmentResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetAttachmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendResponse struct {
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
func (r MessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetAttachmentParams struct {
	// Unique identifier of the chat (supports both chatID and localChatID).
	ChatID string `json:"chatID,required"`
	// The message ID (eventID) containing the attachment.
	MessageID string `json:"messageID,required"`
	paramObj
}

func (r MessageGetAttachmentParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageGetAttachmentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageGetAttachmentParams) UnmarshalJSON(data []byte) error {
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

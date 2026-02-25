// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/apiquery"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
)

// Manage message reactions
//
// ChatMessageReactionService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatMessageReactionService] method instead.
type ChatMessageReactionService struct {
	Options []option.RequestOption
}

// NewChatMessageReactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewChatMessageReactionService(opts ...option.RequestOption) (r ChatMessageReactionService) {
	r = ChatMessageReactionService{}
	r.Options = opts
	return
}

// Remove the authenticated user's reaction from an existing message.
func (r *ChatMessageReactionService) Delete(ctx context.Context, messageID string, params ChatMessageReactionDeleteParams, opts ...option.RequestOption) (res *ChatMessageReactionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ChatID == "" {
		err = errors.New("missing required chatID parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	path := fmt.Sprintf("v1/chats/%s/messages/%s/reactions", params.ChatID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Add a reaction to an existing message.
func (r *ChatMessageReactionService) Add(ctx context.Context, messageID string, params ChatMessageReactionAddParams, opts ...option.RequestOption) (res *ChatMessageReactionAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ChatID == "" {
		err = errors.New("missing required chatID parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	path := fmt.Sprintf("v1/chats/%s/messages/%s/reactions", params.ChatID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ChatMessageReactionDeleteResponse struct {
	// Unique identifier of the chat.
	ChatID string `json:"chatID" api:"required"`
	// Message ID.
	MessageID string `json:"messageID" api:"required"`
	// Reaction key that was removed
	ReactionKey string `json:"reactionKey" api:"required"`
	// Whether the reaction was successfully removed
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID      respjson.Field
		MessageID   respjson.Field
		ReactionKey respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageReactionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageReactionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageReactionAddResponse struct {
	// Unique identifier of the chat.
	ChatID string `json:"chatID" api:"required"`
	// Message ID.
	MessageID string `json:"messageID" api:"required"`
	// Reaction key that was added
	ReactionKey string `json:"reactionKey" api:"required"`
	// Whether the reaction was successfully added
	Success bool `json:"success" api:"required"`
	// Transaction ID used for the reaction event
	TransactionID string `json:"transactionID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChatID        respjson.Field
		MessageID     respjson.Field
		ReactionKey   respjson.Field
		Success       respjson.Field
		TransactionID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageReactionAddResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageReactionAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageReactionDeleteParams struct {
	// Unique identifier of the chat.
	ChatID string `path:"chatID" api:"required" json:"-"`
	// Reaction key to remove
	ReactionKey string `query:"reactionKey" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ChatMessageReactionDeleteParams]'s query parameters as
// `url.Values`.
func (r ChatMessageReactionDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChatMessageReactionAddParams struct {
	// Unique identifier of the chat.
	ChatID string `path:"chatID" api:"required" json:"-"`
	// Reaction key to add (emoji, shortcode, or custom emoji key)
	ReactionKey string `json:"reactionKey" api:"required"`
	// Optional transaction ID for deduplication and local echo tracking
	TransactionID param.Opt[string] `json:"transactionID,omitzero"`
	paramObj
}

func (r ChatMessageReactionAddParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageReactionAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageReactionAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

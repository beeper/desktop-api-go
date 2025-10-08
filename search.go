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

// SearchService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.Options = opts
	return
}

// Search chats by title/network or participants using Beeper Desktop's renderer
// algorithm.
func (r *SearchService) Chats(ctx context.Context, query SearchChatsParams, opts ...option.RequestOption) (res *pagination.CursorSearch[Chat], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/search/chats"
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
func (r *SearchService) ChatsAutoPaging(ctx context.Context, query SearchChatsParams, opts ...option.RequestOption) *pagination.CursorSearchAutoPager[Chat] {
	return pagination.NewCursorSearchAutoPager(r.Chats(ctx, query, opts...))
}

// Search contacts across on a specific account using the network's search API.
// Only use for creating new chats.
func (r *SearchService) Contacts(ctx context.Context, accountID string, query SearchContactsParams, opts ...option.RequestOption) (res *SearchContactsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountID parameter")
		return
	}
	path := fmt.Sprintf("v1/search/contacts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SearchContactsResponse struct {
	Items []shared.User `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchContactsResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchContactsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchChatsParams struct {
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
	Direction SearchChatsParamsDirection `query:"direction,omitzero" json:"-"`
	// Filter by inbox type: "primary" (non-archived, non-low-priority),
	// "low-priority", or "archive". If not specified, shows all chats.
	//
	// Any of "primary", "low-priority", "archive".
	Inbox SearchChatsParamsInbox `query:"inbox,omitzero" json:"-"`
	// Search scope: 'titles' matches title + network; 'participants' matches
	// participant names.
	//
	// Any of "titles", "participants".
	Scope SearchChatsParamsScope `query:"scope,omitzero" json:"-"`
	// Specify the type of chats to retrieve: use "single" for direct messages, "group"
	// for group chats, or "any" to get all types
	//
	// Any of "single", "group", "any".
	Type SearchChatsParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchChatsParams]'s query parameters as `url.Values`.
func (r SearchChatsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
type SearchChatsParamsDirection string

const (
	SearchChatsParamsDirectionAfter  SearchChatsParamsDirection = "after"
	SearchChatsParamsDirectionBefore SearchChatsParamsDirection = "before"
)

// Filter by inbox type: "primary" (non-archived, non-low-priority),
// "low-priority", or "archive". If not specified, shows all chats.
type SearchChatsParamsInbox string

const (
	SearchChatsParamsInboxPrimary     SearchChatsParamsInbox = "primary"
	SearchChatsParamsInboxLowPriority SearchChatsParamsInbox = "low-priority"
	SearchChatsParamsInboxArchive     SearchChatsParamsInbox = "archive"
)

// Search scope: 'titles' matches title + network; 'participants' matches
// participant names.
type SearchChatsParamsScope string

const (
	SearchChatsParamsScopeTitles       SearchChatsParamsScope = "titles"
	SearchChatsParamsScopeParticipants SearchChatsParamsScope = "participants"
)

// Specify the type of chats to retrieve: use "single" for direct messages, "group"
// for group chats, or "any" to get all types
type SearchChatsParamsType string

const (
	SearchChatsParamsTypeSingle SearchChatsParamsType = "single"
	SearchChatsParamsTypeGroup  SearchChatsParamsType = "group"
	SearchChatsParamsTypeAny    SearchChatsParamsType = "any"
)

type SearchContactsParams struct {
	// Text to search users by. Network-specific behavior.
	Query string `query:"query,required" json:"-"`
	paramObj
}

// URLQuery serializes [SearchContactsParams]'s query parameters as `url.Values`.
func (r SearchContactsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

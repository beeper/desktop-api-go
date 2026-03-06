// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"net/url"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/apiquery"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
	"github.com/beeper/desktop-api-go/shared"
)

// Response indicating successful app focus action.
type FocusResponse struct {
	// Whether the app was successfully opened/focused.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FocusResponse) RawJSON() string { return r.JSON.raw }
func (r *FocusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchResponse struct {
	Results SearchResponseResults `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchResponseResults struct {
	// Top chat results.
	Chats []Chat `json:"chats" api:"required"`
	// Top group results by participant matches.
	InGroups []Chat                        `json:"in_groups" api:"required"`
	Messages SearchResponseResultsMessages `json:"messages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chats       respjson.Field
		InGroups    respjson.Field
		Messages    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseResults) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchResponseResultsMessages struct {
	// Map of chatID -> chat details for chats referenced in items.
	Chats map[string]Chat `json:"chats" api:"required"`
	// True if additional results can be fetched using the provided cursors.
	HasMore bool `json:"hasMore" api:"required"`
	// Messages matching the query and filters.
	Items []shared.Message `json:"items" api:"required"`
	// Cursor for fetching newer results (use with direction='after'). Opaque string;
	// do not inspect.
	NewestCursor string `json:"newestCursor" api:"required"`
	// Cursor for fetching older results (use with direction='before'). Opaque string;
	// do not inspect.
	OldestCursor string `json:"oldestCursor" api:"required"`
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
func (r SearchResponseResultsMessages) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseResultsMessages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FocusParams struct {
	// Optional Beeper chat ID (or local chat ID) to focus after opening the app. If
	// omitted, only opens/focuses the app.
	ChatID param.Opt[string] `json:"chatID,omitzero"`
	// Optional draft attachment path to populate in the message input field.
	DraftAttachmentPath param.Opt[string] `json:"draftAttachmentPath,omitzero"`
	// Optional draft text to populate in the message input field.
	DraftText param.Opt[string] `json:"draftText,omitzero"`
	// Optional message ID. Jumps to that message in the chat when opening.
	MessageID param.Opt[string] `json:"messageID,omitzero"`
	paramObj
}

func (r FocusParams) MarshalJSON() (data []byte, err error) {
	type shadow FocusParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FocusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchParams struct {
	// User-typed search text. Literal word matching (non-semantic).
	Query string `query:"query" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [SearchParams]'s query parameters as `url.Values`.
func (r SearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

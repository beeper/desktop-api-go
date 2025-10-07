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

type DownloadAssetResponse struct {
	// Error message if the download failed.
	Error string `json:"error"`
	// Local file URL to the downloaded asset.
	SrcURL string `json:"srcURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		SrcURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DownloadAssetResponse) RawJSON() string { return r.JSON.raw }
func (r *DownloadAssetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response indicating successful app opening.
type OpenResponse struct {
	// Whether the app was successfully opened/focused.
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenResponse) RawJSON() string { return r.JSON.raw }
func (r *OpenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchResponse struct {
	Results SearchResponseResults `json:"results,required"`
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
	Chats []Chat `json:"chats,required"`
	// Top group results by participant matches.
	InGroups []Chat                        `json:"in_groups,required"`
	Messages SearchResponseResultsMessages `json:"messages,required"`
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
func (r SearchResponseResultsMessages) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseResultsMessages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DownloadAssetParams struct {
	// Matrix content URL (mxc:// or localmxc://) for the asset to download.
	URL string `json:"url,required"`
	paramObj
}

func (r DownloadAssetParams) MarshalJSON() (data []byte, err error) {
	type shadow DownloadAssetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DownloadAssetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenParams struct {
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

func (r OpenParams) MarshalJSON() (data []byte, err error) {
	type shadow OpenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchParams struct {
	// User-typed search text. Literal word matching (NOT semantic).
	Query string `query:"query,required" json:"-"`
	paramObj
}

// URLQuery serializes [SearchParams]'s query parameters as `url.Values`.
func (r SearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

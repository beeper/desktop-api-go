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
	"github.com/beeper/desktop-api-go/packages/pagination"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
	"github.com/beeper/desktop-api-go/shared"
)

// Manage contacts on a specific account
//
// AccountContactService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountContactService] method instead.
type AccountContactService struct {
	Options []option.RequestOption
}

// NewAccountContactService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountContactService(opts ...option.RequestOption) (r AccountContactService) {
	r = AccountContactService{}
	r.Options = opts
	return
}

// List merged contacts for a specific account with cursor-based pagination.
func (r *AccountContactService) List(ctx context.Context, accountID string, query AccountContactListParams, opts ...option.RequestOption) (res *pagination.CursorSearch[shared.User], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountID parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s/contacts/list", accountID)
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

// List merged contacts for a specific account with cursor-based pagination.
func (r *AccountContactService) ListAutoPaging(ctx context.Context, accountID string, query AccountContactListParams, opts ...option.RequestOption) *pagination.CursorSearchAutoPager[shared.User] {
	return pagination.NewCursorSearchAutoPager(r.List(ctx, accountID, query, opts...))
}

// Search contacts on a specific account using merged account contacts, network
// search, and exact identifier lookup.
func (r *AccountContactService) Search(ctx context.Context, accountID string, query AccountContactSearchParams, opts ...option.RequestOption) (res *AccountContactSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountID parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/%s/contacts", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountContactSearchResponse struct {
	Items []shared.User `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountContactSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountContactSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountContactListParams struct {
	// Opaque pagination cursor; do not inspect. Use together with 'direction'.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum contacts to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional search query for blended contact lookup.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
	// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
	//
	// Any of "after", "before".
	Direction AccountContactListParamsDirection `query:"direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountContactListParams]'s query parameters as
// `url.Values`.
func (r AccountContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Pagination direction used with 'cursor': 'before' fetches older results, 'after'
// fetches newer results. Defaults to 'before' when only 'cursor' is provided.
type AccountContactListParamsDirection string

const (
	AccountContactListParamsDirectionAfter  AccountContactListParamsDirection = "after"
	AccountContactListParamsDirectionBefore AccountContactListParamsDirection = "before"
)

type AccountContactSearchParams struct {
	// Text to search users by. Network-specific behavior.
	Query string `query:"query,required" json:"-"`
	paramObj
}

// URLQuery serializes [AccountContactSearchParams]'s query parameters as
// `url.Values`.
func (r AccountContactSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

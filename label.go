// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v6/internal/apijson"
	"github.com/beeper/desktop-api-go/v6/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v6/option"
	"github.com/beeper/desktop-api-go/v6/packages/respjson"
)

// User-created labels that organize chats
//
// LabelService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLabelService] method instead.
type LabelService struct {
	Options []option.RequestOption
}

// NewLabelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLabelService(opts ...option.RequestOption) (r LabelService) {
	r = LabelService{}
	r.Options = opts
	return
}

// List the labels the user has created for organizing chats. Filter chats by label
// via GET /v1/chats/search with labelID.
func (r *LabelService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Label, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/labels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A user-created label that organizes chats across accounts.
type Label struct {
	// Unique identifier of the label.
	ID string `json:"id" api:"required"`
	// Display name of the label.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Label) RawJSON() string { return r.JSON.raw }
func (r *Label) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

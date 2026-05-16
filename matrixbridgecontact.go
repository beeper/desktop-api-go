// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/apiquery"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/param"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// Matrix-compatible APIs for accounts and connected network bridges.
//
// MatrixBridgeContactService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixBridgeContactService] method instead.
type MatrixBridgeContactService struct {
	Options []option.RequestOption
}

// NewMatrixBridgeContactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatrixBridgeContactService(opts ...option.RequestOption) (r MatrixBridgeContactService) {
	r = MatrixBridgeContactService{}
	r.Options = opts
	return
}

// Get a list of contacts.
func (r *MatrixBridgeContactService) List(ctx context.Context, bridgeID string, query MatrixBridgeContactListParams, opts ...option.RequestOption) (res *MatrixBridgeContactListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/contacts", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MatrixBridgeContactListResponse struct {
	Contacts []MatrixBridgeContactListResponseContact `json:"contacts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contacts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeContactListResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeContactListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A successfully resolved identifier.
type MatrixBridgeContactListResponseContact struct {
	// The internal user ID of the resolved user.
	ID string `json:"id" api:"required"`
	// The avatar of the user on the remote network.
	AvatarURL string `json:"avatar_url" format:"mxc"`
	// The Matrix room ID of the direct chat with the user.
	DmRoomMxid string `json:"dm_room_mxid" format:"matrix_room_id"`
	// A list of identifiers for the user on the remote network.
	Identifiers []string `json:"identifiers" format:"uri"`
	// The Matrix user ID of the ghost representing the user.
	Mxid string `json:"mxid" format:"matrix_user_id"`
	// The name of the user on the remote network.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AvatarURL   respjson.Field
		DmRoomMxid  respjson.Field
		Identifiers respjson.Field
		Mxid        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeContactListResponseContact) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeContactListResponseContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixBridgeContactListParams struct {
	// An optional explicit login ID to do the action through.
	LoginID param.Opt[string] `query:"login_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MatrixBridgeContactListParams]'s query parameters as
// `url.Values`.
func (r MatrixBridgeContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
// MatrixBridgeUserService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixBridgeUserService] method instead.
type MatrixBridgeUserService struct {
	Options []option.RequestOption
}

// NewMatrixBridgeUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatrixBridgeUserService(opts ...option.RequestOption) (r MatrixBridgeUserService) {
	r = MatrixBridgeUserService{}
	r.Options = opts
	return
}

// Resolve an identifier to a user on the remote network.
func (r *MatrixBridgeUserService) Resolve(ctx context.Context, identifier string, params MatrixBridgeUserResolveParams, opts ...option.RequestOption) (res *MatrixBridgeUserResolveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/resolve_identifier/%s", params.BridgeID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Search for users on the remote network
func (r *MatrixBridgeUserService) Search(ctx context.Context, bridgeID string, params MatrixBridgeUserSearchParams, opts ...option.RequestOption) (res *MatrixBridgeUserSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/search_users", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// A successfully resolved identifier.
type MatrixBridgeUserResolveResponse struct {
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
func (r MatrixBridgeUserResolveResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeUserResolveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixBridgeUserSearchResponse struct {
	Results []MatrixBridgeUserSearchResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeUserSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeUserSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A successfully resolved identifier.
type MatrixBridgeUserSearchResponseResult struct {
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
func (r MatrixBridgeUserSearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeUserSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixBridgeUserResolveParams struct {
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	// An optional explicit login ID to do the action through.
	LoginID param.Opt[string] `query:"login_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MatrixBridgeUserResolveParams]'s query parameters as
// `url.Values`.
func (r MatrixBridgeUserResolveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MatrixBridgeUserSearchParams struct {
	// An optional explicit login ID to do the action through.
	LoginID param.Opt[string] `query:"login_id,omitzero" json:"-"`
	// The search query to send to the remote network
	Query param.Opt[string] `json:"query,omitzero"`
	paramObj
}

func (r MatrixBridgeUserSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow MatrixBridgeUserSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixBridgeUserSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MatrixBridgeUserSearchParams]'s query parameters as
// `url.Values`.
func (r MatrixBridgeUserSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

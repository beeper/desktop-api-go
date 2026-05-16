// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	shimjson "github.com/beeper/desktop-api-go/v5/internal/encoding/json"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
)

// MatrixRoomAccountDataService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixRoomAccountDataService] method instead.
type MatrixRoomAccountDataService struct {
	Options []option.RequestOption
}

// NewMatrixRoomAccountDataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatrixRoomAccountDataService(opts ...option.RequestOption) (r MatrixRoomAccountDataService) {
	r = MatrixRoomAccountDataService{}
	r.Options = opts
	return
}

// Get some account data for the client on a given room. This config is only
// visible to the user that set the account data.
func (r *MatrixRoomAccountDataService) Get(ctx context.Context, type_ string, query MatrixRoomAccountDataGetParams, opts ...option.RequestOption) (res *MatrixRoomAccountDataGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.UserID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	if query.RoomID == "" {
		err = errors.New("missing required roomId parameter")
		return nil, err
	}
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/v3/user/%s/rooms/%s/account_data/%s", query.UserID, query.RoomID, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set some account data for the client on a given room. This config is only
// visible to the user that set the account data. The config will be delivered to
// clients in the per-room entries via
// [/sync](https://spec.matrix.org/v1.18/client-server-api/#get_matrixclientv3sync).
func (r *MatrixRoomAccountDataService) Update(ctx context.Context, type_ string, params MatrixRoomAccountDataUpdateParams, opts ...option.RequestOption) (res *MatrixRoomAccountDataUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.UserID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	if params.RoomID == "" {
		err = errors.New("missing required roomId parameter")
		return nil, err
	}
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/v3/user/%s/rooms/%s/account_data/%s", params.UserID, params.RoomID, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type MatrixRoomAccountDataGetResponse = any

type MatrixRoomAccountDataUpdateResponse = any

type MatrixRoomAccountDataGetParams struct {
	UserID string `path:"userId" api:"required" json:"-"`
	RoomID string `path:"roomId" api:"required" json:"-"`
	paramObj
}

type MatrixRoomAccountDataUpdateParams struct {
	UserID string `path:"userId" api:"required" json:"-"`
	RoomID string `path:"roomId" api:"required" json:"-"`
	Body   any
	paramObj
}

func (r MatrixRoomAccountDataUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *MatrixRoomAccountDataUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

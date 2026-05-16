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

// MatrixUserAccountDataService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixUserAccountDataService] method instead.
type MatrixUserAccountDataService struct {
	Options []option.RequestOption
}

// NewMatrixUserAccountDataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatrixUserAccountDataService(opts ...option.RequestOption) (r MatrixUserAccountDataService) {
	r = MatrixUserAccountDataService{}
	r.Options = opts
	return
}

// Get some account data for the client. This config is only visible to the user
// that set the account data.
func (r *MatrixUserAccountDataService) Get(ctx context.Context, type_ string, query MatrixUserAccountDataGetParams, opts ...option.RequestOption) (res *MatrixUserAccountDataGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.UserID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/v3/user/%s/account_data/%s", query.UserID, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set some account data for the client. This config is only visible to the user
// that set the account data. The config will be available to clients through the
// top-level `account_data` field in the homeserver response to
// [/sync](https://spec.matrix.org/v1.18/client-server-api/#get_matrixclientv3sync).
func (r *MatrixUserAccountDataService) Update(ctx context.Context, type_ string, params MatrixUserAccountDataUpdateParams, opts ...option.RequestOption) (res *MatrixUserAccountDataUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.UserID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/v3/user/%s/account_data/%s", params.UserID, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type MatrixUserAccountDataGetResponse = any

type MatrixUserAccountDataUpdateResponse = any

type MatrixUserAccountDataGetParams struct {
	UserID string `path:"userId" api:"required" json:"-"`
	paramObj
}

type MatrixUserAccountDataUpdateParams struct {
	UserID string `path:"userId" api:"required" json:"-"`
	Body   any
	paramObj
}

func (r MatrixUserAccountDataUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *MatrixUserAccountDataUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

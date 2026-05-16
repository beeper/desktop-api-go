// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// MatrixUserService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixUserService] method instead.
type MatrixUserService struct {
	Options     []option.RequestOption
	AccountData MatrixUserAccountDataService
}

// NewMatrixUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatrixUserService(opts ...option.RequestOption) (r MatrixUserService) {
	r = MatrixUserService{}
	r.Options = opts
	r.AccountData = NewMatrixUserAccountDataService(opts...)
	return
}

// Get the complete profile for a user.
func (r *MatrixUserService) GetProfile(ctx context.Context, userID string, opts ...option.RequestOption) (res *MatrixUserGetProfileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/v3/profile/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MatrixUserGetProfileResponse struct {
	// The user's avatar URL if they have set one, otherwise not present.
	AvatarURL string `json:"avatar_url" format:"mx-mxc-uri"`
	// The user's display name if they have set one, otherwise not present.
	Displayname string `json:"displayname"`
	// The user's time zone.
	MTz         string         `json:"m.tz"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvatarURL   respjson.Field
		Displayname respjson.Field
		MTz         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixUserGetProfileResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixUserGetProfileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

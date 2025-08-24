// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo

import (
	"context"
	"net/http"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/shared"
)

// Control the Beeper Desktop application
//
// AppService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r AppService) {
	r = AppService{}
	r.Options = opts
	return
}

// Bring Beeper Desktop to the foreground on this device
func (r *AppService) Focus(ctx context.Context, body AppFocusParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/focus-app"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AppFocusParams struct {
	// Optional Beeper chat ID to focus after bringing the app to foreground. If
	// omitted, only foregrounds the app. Required if messageSortKey is present. No-op
	// in headless environments.
	ChatID param.Opt[string] `json:"chatID,omitzero"`
	// Optional message sort key. Jumps to that message in the chat when foregrounding.
	MessageSortKey param.Opt[string] `json:"messageSortKey,omitzero"`
	paramObj
}

func (r AppFocusParams) MarshalJSON() (data []byte, err error) {
	type shadow AppFocusParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppFocusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

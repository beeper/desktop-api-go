// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo

import (
	"context"
	"net/http"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
)

// App operations
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

// Open Beeper Desktop and optionally navigate to a specific chat, message, or
// pre-fill draft text
func (r *AppService) Focus(ctx context.Context, body AppFocusParams, opts ...option.RequestOption) (res *AppFocusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/open-app"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response indicating successful app opening.
type AppFocusResponse struct {
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
func (r AppFocusResponse) RawJSON() string { return r.JSON.raw }
func (r *AppFocusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppFocusParams struct {
	// Optional Beeper chat ID to focus after opening the app. If omitted, only
	// opens/focuses the app.
	ChatID param.Opt[string] `json:"chatID,omitzero"`
	// Optional draft text to populate in the message input field.
	DraftText param.Opt[string] `json:"draftText,omitzero"`
	// Optional message sort key. Jumps to that message in the chat when opening.
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

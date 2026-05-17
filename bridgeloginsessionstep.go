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
	"github.com/beeper/desktop-api-go/v5/packages/param"
)

// Available bridges, bridge logins, login sessions for connect and reconnect
// flows, and advanced network capabilities.
//
// BridgeLoginSessionStepService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeLoginSessionStepService] method instead.
type BridgeLoginSessionStepService struct {
	Options []option.RequestOption
}

// NewBridgeLoginSessionStepService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBridgeLoginSessionStepService(opts ...option.RequestOption) (r BridgeLoginSessionStepService) {
	r = BridgeLoginSessionStepService{}
	r.Options = opts
	return
}

// Submit input for the current step of a bridge login session.
func (r *BridgeLoginSessionStepService) Submit(ctx context.Context, stepID string, params BridgeLoginSessionStepSubmitParams, opts ...option.RequestOption) (res *LoginSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if params.LoginSessionID == "" {
		err = errors.New("missing required loginSessionID parameter")
		return nil, err
	}
	if stepID == "" {
		err = errors.New("missing required stepID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s/login-sessions/%s/steps/%s", params.BridgeID, params.LoginSessionID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type BridgeLoginSessionStepSubmitParams struct {
	// Bridge ID.
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	// Temporary bridge login session ID.
	LoginSessionID string `path:"loginSessionID" api:"required" json:"-"`
	// Any of "user_input", "cookies", "display_and_wait".
	Type BridgeLoginSessionStepSubmitParamsType `json:"type,omitzero" api:"required"`
	// Last browser URL reached during a cookies step, if available.
	LastURL param.Opt[string] `json:"lastURL,omitzero"`
	// Field values keyed by the field IDs from the current step.
	Fields map[string]string `json:"fields,omitzero"`
	// How the step was completed. Omit unless the client needs to distinguish an
	// embedded webview or browser extension.
	//
	// Any of "api", "webview", "browser_extension".
	Source BridgeLoginSessionStepSubmitParamsSource `json:"source,omitzero"`
	paramObj
}

func (r BridgeLoginSessionStepSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLoginSessionStepSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLoginSessionStepSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLoginSessionStepSubmitParamsType string

const (
	BridgeLoginSessionStepSubmitParamsTypeUserInput      BridgeLoginSessionStepSubmitParamsType = "user_input"
	BridgeLoginSessionStepSubmitParamsTypeCookies        BridgeLoginSessionStepSubmitParamsType = "cookies"
	BridgeLoginSessionStepSubmitParamsTypeDisplayAndWait BridgeLoginSessionStepSubmitParamsType = "display_and_wait"
)

// How the step was completed. Omit unless the client needs to distinguish an
// embedded webview or browser extension.
type BridgeLoginSessionStepSubmitParamsSource string

const (
	BridgeLoginSessionStepSubmitParamsSourceAPI              BridgeLoginSessionStepSubmitParamsSource = "api"
	BridgeLoginSessionStepSubmitParamsSourceWebview          BridgeLoginSessionStepSubmitParamsSource = "webview"
	BridgeLoginSessionStepSubmitParamsSourceBrowserExtension BridgeLoginSessionStepSubmitParamsSource = "browser_extension"
)

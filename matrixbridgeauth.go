// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/apiquery"
	shimjson "github.com/beeper/desktop-api-go/v5/internal/encoding/json"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/param"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
	"github.com/beeper/desktop-api-go/v5/shared/constant"
)

// Matrix-compatible APIs for accounts and connected network bridges.
//
// MatrixBridgeAuthService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixBridgeAuthService] method instead.
type MatrixBridgeAuthService struct {
	Options []option.RequestOption
}

// NewMatrixBridgeAuthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatrixBridgeAuthService(opts ...option.RequestOption) (r MatrixBridgeAuthService) {
	r = MatrixBridgeAuthService{}
	r.Options = opts
	return
}

// Get the available login flows.
func (r *MatrixBridgeAuthService) ListFlows(ctx context.Context, bridgeID string, opts ...option.RequestOption) (res *MatrixBridgeAuthListFlowsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/login/flows", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the login IDs of the current user.
func (r *MatrixBridgeAuthService) ListLogins(ctx context.Context, bridgeID string, opts ...option.RequestOption) (res *MatrixBridgeAuthListLoginsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/logins", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Log out of an existing login.
func (r *MatrixBridgeAuthService) Logout(ctx context.Context, loginID string, body MatrixBridgeAuthLogoutParams, opts ...option.RequestOption) (res *MatrixBridgeAuthLogoutResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if loginID == "" {
		err = errors.New("missing required loginID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/logout/%s", body.BridgeID, loginID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// This endpoint starts a new login process, which is used to log into the bridge.
//
// The basic flow of the entire login, including calling this endpoint, is:
//
//  1. Call `GET /v3/login/flows` to get the list of available flows. If there's
//     more than one flow, ask the user to pick which one they want to use.
//  2. Call this endpoint with the chosen flow ID to start the login. The first
//     login step will be returned.
//  3. Render the information provided in the step.
//  4. Call the `/login/step/...` endpoint corresponding to the step type:
//     - For `user_input` and `cookies`, acquire the requested fields before calling
//     the endpoint.
//     - For `display_and_wait`, call the endpoint immediately (as there's nothing
//     to acquire on the client side).
//  5. Handle the data returned by the login step endpoint:
//     - If an error is returned, the login has failed and must be restarted (from
//     either step 1 or step 2) if the user wants to try again.
//     - If step type `complete` is returned, the login finished successfully.
//     - Otherwise, go to step 3 with the new data.
func (r *MatrixBridgeAuthService) StartLogin(ctx context.Context, flowID string, params MatrixBridgeAuthStartLoginParams, opts ...option.RequestOption) (res *MatrixBridgeAuthStartLoginResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if flowID == "" {
		err = errors.New("missing required flowID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/login/start/%s", params.BridgeID, flowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Submit extracted cookies in a login process.
func (r *MatrixBridgeAuthService) SubmitCookies(ctx context.Context, stepID string, params MatrixBridgeAuthSubmitCookiesParams, opts ...option.RequestOption) (res *MatrixBridgeAuthSubmitCookiesResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if params.LoginProcessID == "" {
		err = errors.New("missing required loginProcessID parameter")
		return nil, err
	}
	if stepID == "" {
		err = errors.New("missing required stepID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/login/step/%s/%s/cookies", params.BridgeID, params.LoginProcessID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Submit user input in a login process.
func (r *MatrixBridgeAuthService) SubmitUserInput(ctx context.Context, stepID string, params MatrixBridgeAuthSubmitUserInputParams, opts ...option.RequestOption) (res *MatrixBridgeAuthSubmitUserInputResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if params.LoginProcessID == "" {
		err = errors.New("missing required loginProcessID parameter")
		return nil, err
	}
	if stepID == "" {
		err = errors.New("missing required stepID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/login/step/%s/%s/user_input", params.BridgeID, params.LoginProcessID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Wait for the next step after displaying data to the user.
func (r *MatrixBridgeAuthService) WaitForStep(ctx context.Context, stepID string, body MatrixBridgeAuthWaitForStepParams, opts ...option.RequestOption) (res *MatrixBridgeAuthWaitForStepResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if body.LoginProcessID == "" {
		err = errors.New("missing required loginProcessID parameter")
		return nil, err
	}
	if stepID == "" {
		err = errors.New("missing required stepID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/login/step/%s/%s/display_and_wait", body.BridgeID, body.LoginProcessID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get all info that is useful for presenting this bridge in a manager interface.
//
//   - Server details: remote network details, available login flows, homeserver
//     name, bridge bot user ID, command prefix
//   - User details: management room ID, list of logins with current state and info
func (r *MatrixBridgeAuthService) Whoami(ctx context.Context, bridgeID string, opts ...option.RequestOption) (res *MatrixBridgeAuthWhoamiResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/whoami", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MatrixBridgeAuthListFlowsResponse struct {
	Flows []MatrixBridgeAuthListFlowsResponseFlow `json:"flows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flows       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthListFlowsResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthListFlowsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual login flow which can be used to sign into the remote network.
type MatrixBridgeAuthListFlowsResponseFlow struct {
	// An internal ID that is passed to the /login/start call to start a login with
	// this flow.
	ID string `json:"id" api:"required"`
	// A human-readable description of the login flow.
	Description string `json:"description" api:"required"`
	// A human-readable name for the login flow.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthListFlowsResponseFlow) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthListFlowsResponseFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixBridgeAuthListLoginsResponse struct {
	LoginIDs []string `json:"login_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LoginIDs    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthListLoginsResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthListLoginsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixBridgeAuthLogoutResponse = any

// MatrixBridgeAuthStartLoginResponseUnion contains all possible properties and
// values from [MatrixBridgeAuthStartLoginResponseObject],
// [MatrixBridgeAuthStartLoginResponseObject2],
// [MatrixBridgeAuthStartLoginResponseObject3],
// [MatrixBridgeAuthStartLoginResponseObject4].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MatrixBridgeAuthStartLoginResponseUnion struct {
	// This field is from variant [MatrixBridgeAuthStartLoginResponseObject].
	DisplayAndWait MatrixBridgeAuthStartLoginResponseObjectDisplayAndWait `json:"display_and_wait"`
	Type           string                                                 `json:"type"`
	Instructions   string                                                 `json:"instructions"`
	LoginID        string                                                 `json:"login_id"`
	StepID         string                                                 `json:"step_id"`
	// This field is from variant [MatrixBridgeAuthStartLoginResponseObject2].
	UserInput MatrixBridgeAuthStartLoginResponseObject2UserInput `json:"user_input"`
	// This field is from variant [MatrixBridgeAuthStartLoginResponseObject3].
	Cookies MatrixBridgeAuthStartLoginResponseObject3Cookies `json:"cookies"`
	// This field is from variant [MatrixBridgeAuthStartLoginResponseObject4].
	Complete MatrixBridgeAuthStartLoginResponseObject4Complete `json:"complete"`
	JSON     struct {
		DisplayAndWait respjson.Field
		Type           respjson.Field
		Instructions   respjson.Field
		LoginID        respjson.Field
		StepID         respjson.Field
		UserInput      respjson.Field
		Cookies        respjson.Field
		Complete       respjson.Field
		raw            string
	} `json:"-"`
}

func (u MatrixBridgeAuthStartLoginResponseUnion) AsMatrixBridgeAuthStartLoginResponseObject() (v MatrixBridgeAuthStartLoginResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthStartLoginResponseUnion) AsMatrixBridgeAuthStartLoginResponseObject2() (v MatrixBridgeAuthStartLoginResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthStartLoginResponseUnion) AsMatrixBridgeAuthStartLoginResponseObject3() (v MatrixBridgeAuthStartLoginResponseObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthStartLoginResponseUnion) AsMatrixBridgeAuthStartLoginResponseObject4() (v MatrixBridgeAuthStartLoginResponseObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MatrixBridgeAuthStartLoginResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *MatrixBridgeAuthStartLoginResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Display and wait login step
type MatrixBridgeAuthStartLoginResponseObject struct {
	// Parameters for the display and wait login step
	DisplayAndWait MatrixBridgeAuthStartLoginResponseObjectDisplayAndWait `json:"display_and_wait" api:"required"`
	Type           constant.DisplayAndWait                                `json:"type" default:"display_and_wait"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayAndWait respjson.Field
		Type           respjson.Field
		Instructions   respjson.Field
		LoginID        respjson.Field
		StepID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the display and wait login step
type MatrixBridgeAuthStartLoginResponseObjectDisplayAndWait struct {
	// The type of thing to display
	//
	// Any of "qr", "emoji", "code", "nothing".
	Type string `json:"type" api:"required"`
	// The thing to display (raw data for QR, unicode emoji for emoji, plain string for
	// code)
	Data string `json:"data"`
	// An image containing the thing to display. If present, this is recommended over
	// using data directly. For emojis, the URL to the canonical image representation
	// of the emoji
	ImageURL string `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObjectDisplayAndWait) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObjectDisplayAndWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User input login step
type MatrixBridgeAuthStartLoginResponseObject2 struct {
	Type constant.UserInput `json:"type" default:"user_input"`
	// Parameters for the user input login step
	UserInput MatrixBridgeAuthStartLoginResponseObject2UserInput `json:"user_input" api:"required"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type         respjson.Field
		UserInput    respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the user input login step
type MatrixBridgeAuthStartLoginResponseObject2UserInput struct {
	// The list of fields that the user is requested to fill.
	Fields []MatrixBridgeAuthStartLoginResponseObject2UserInputField `json:"fields" api:"required"`
	// A list of media attachments to show the user alongside the form fields.
	Attachments []MatrixBridgeAuthStartLoginResponseObject2UserInputAttachment `json:"attachments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields      respjson.Field
		Attachments respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject2UserInput) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObject2UserInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that the user can fill.
type MatrixBridgeAuthStartLoginResponseObject2UserInputField struct {
	// The internal ID of the field. This must be used as the key in the object when
	// submitting the data back to the bridge.
	ID string `json:"id" api:"required"`
	// The name of the field shown to the user.
	Name string `json:"name" api:"required"`
	// The type of field.
	//
	// Any of "username", "phone_number", "email", "password", "2fa_code", "token",
	// "url", "domain", "select".
	Type string `json:"type" api:"required"`
	// A default value that the client can pre-fill the field with.
	DefaultValue string `json:"default_value"`
	// A more detailed description of the field shown to the user.
	Description string `json:"description"`
	// For fields of type select, the valid options.
	Options []string `json:"options"`
	// A regular expression that the field value must match.
	Pattern string `json:"pattern" format:"regex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		DefaultValue respjson.Field
		Description  respjson.Field
		Options      respjson.Field
		Pattern      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject2UserInputField) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObject2UserInputField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media attachment to show the user.
type MatrixBridgeAuthStartLoginResponseObject2UserInputAttachment struct {
	// The raw file content for the attachment encoded in base64.
	Content string `json:"content" api:"required"`
	// The filename for the media attachment.
	Filename string `json:"filename" api:"required"`
	// The type of media attachment, using the same media type identifiers as Matrix
	// attachments. Only some are supported.
	//
	// Any of "m.image", "m.audio".
	Type string `json:"type" api:"required"`
	// Optional but recommended metadata for the attachment. Can generally be derived
	// from the raw content if omitted.
	Info MatrixBridgeAuthStartLoginResponseObject2UserInputAttachmentInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject2UserInputAttachment) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthStartLoginResponseObject2UserInputAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional but recommended metadata for the attachment. Can generally be derived
// from the raw content if omitted.
type MatrixBridgeAuthStartLoginResponseObject2UserInputAttachmentInfo struct {
	// The height of the media in pixels. Only applicable for images and videos.
	H float64 `json:"h"`
	// The MIME type for the media content.
	Mimetype string `json:"mimetype"`
	// The size of the media content in number of bytes. Strongly recommended to
	// include.
	Size float64 `json:"size"`
	// The width of the media in pixels. Only applicable for images and videos.
	W float64 `json:"w"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H           respjson.Field
		Mimetype    respjson.Field
		Size        respjson.Field
		W           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject2UserInputAttachmentInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthStartLoginResponseObject2UserInputAttachmentInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cookie login step
type MatrixBridgeAuthStartLoginResponseObject3 struct {
	// Parameters for the cookie login step
	Cookies MatrixBridgeAuthStartLoginResponseObject3Cookies `json:"cookies" api:"required"`
	Type    constant.Cookies                                 `json:"type" default:"cookies"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cookies      respjson.Field
		Type         respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject3) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the cookie login step
type MatrixBridgeAuthStartLoginResponseObject3Cookies struct {
	// The list of cookies or other stored data that must be extracted.
	Fields []MatrixBridgeAuthStartLoginResponseObject3CookiesField `json:"fields" api:"required"`
	// The URL to open when using a webview to extract cookies.
	URL string `json:"url" api:"required" format:"uri"`
	// A JavaScript snippet that can extract some or all of the fields. The snippet
	// will evaluate to a promise that resolves when the relevant fields are found.
	// Fields that are not present in the promise result must be extracted another way.
	ExtractJs string `json:"extract_js"`
	// An optional user agent that the webview should use.
	UserAgent string `json:"user_agent"`
	// A regex pattern that the URL should match before the client closes the webview.
	//
	// The client may submit the login if the user closes the webview after all cookies
	// are collected even if this URL is not reached, but it should only automatically
	// close the webview after both cookies and the URL match.
	WaitForURLPattern string `json:"wait_for_url_pattern"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields            respjson.Field
		URL               respjson.Field
		ExtractJs         respjson.Field
		UserAgent         respjson.Field
		WaitForURLPattern respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject3Cookies) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObject3Cookies) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual cookie or other stored data item that must be extracted.
type MatrixBridgeAuthStartLoginResponseObject3CookiesField struct {
	// The name of the item to extract.
	Name string `json:"name" api:"required"`
	// The type of data to extract.
	//
	// Any of "cookie", "local_storage", "request_header", "request_body", "special".
	Type string `json:"type" api:"required"`
	// For the `cookie` type, the domain of the cookie.
	CookieDomain string `json:"cookie_domain"`
	// For the `request_header` and `request_body` types, a regex that matches the URLs
	// from which the values can be extracted.
	RequestURLRegex string `json:"request_url_regex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Type            respjson.Field
		CookieDomain    respjson.Field
		RequestURLRegex respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject3CookiesField) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObject3CookiesField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Login complete
type MatrixBridgeAuthStartLoginResponseObject4 struct {
	// Information about the completed login
	Complete MatrixBridgeAuthStartLoginResponseObject4Complete `json:"complete" api:"required"`
	Type     constant.Complete                                 `json:"type" default:"complete"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete     respjson.Field
		Type         respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject4) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the completed login
type MatrixBridgeAuthStartLoginResponseObject4Complete struct {
	// The unique ID of a login. Defined by the network connector.
	UserLoginID string `json:"user_login_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserLoginID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthStartLoginResponseObject4Complete) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthStartLoginResponseObject4Complete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MatrixBridgeAuthSubmitCookiesResponseUnion contains all possible properties and
// values from [MatrixBridgeAuthSubmitCookiesResponseObject],
// [MatrixBridgeAuthSubmitCookiesResponseObject2],
// [MatrixBridgeAuthSubmitCookiesResponseObject3],
// [MatrixBridgeAuthSubmitCookiesResponseObject4].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MatrixBridgeAuthSubmitCookiesResponseUnion struct {
	// This field is from variant [MatrixBridgeAuthSubmitCookiesResponseObject].
	DisplayAndWait MatrixBridgeAuthSubmitCookiesResponseObjectDisplayAndWait `json:"display_and_wait"`
	Type           string                                                    `json:"type"`
	Instructions   string                                                    `json:"instructions"`
	LoginID        string                                                    `json:"login_id"`
	StepID         string                                                    `json:"step_id"`
	// This field is from variant [MatrixBridgeAuthSubmitCookiesResponseObject2].
	UserInput MatrixBridgeAuthSubmitCookiesResponseObject2UserInput `json:"user_input"`
	// This field is from variant [MatrixBridgeAuthSubmitCookiesResponseObject3].
	Cookies MatrixBridgeAuthSubmitCookiesResponseObject3Cookies `json:"cookies"`
	// This field is from variant [MatrixBridgeAuthSubmitCookiesResponseObject4].
	Complete MatrixBridgeAuthSubmitCookiesResponseObject4Complete `json:"complete"`
	JSON     struct {
		DisplayAndWait respjson.Field
		Type           respjson.Field
		Instructions   respjson.Field
		LoginID        respjson.Field
		StepID         respjson.Field
		UserInput      respjson.Field
		Cookies        respjson.Field
		Complete       respjson.Field
		raw            string
	} `json:"-"`
}

func (u MatrixBridgeAuthSubmitCookiesResponseUnion) AsMatrixBridgeAuthSubmitCookiesResponseObject() (v MatrixBridgeAuthSubmitCookiesResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthSubmitCookiesResponseUnion) AsMatrixBridgeAuthSubmitCookiesResponseObject2() (v MatrixBridgeAuthSubmitCookiesResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthSubmitCookiesResponseUnion) AsMatrixBridgeAuthSubmitCookiesResponseObject3() (v MatrixBridgeAuthSubmitCookiesResponseObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthSubmitCookiesResponseUnion) AsMatrixBridgeAuthSubmitCookiesResponseObject4() (v MatrixBridgeAuthSubmitCookiesResponseObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MatrixBridgeAuthSubmitCookiesResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *MatrixBridgeAuthSubmitCookiesResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Display and wait login step
type MatrixBridgeAuthSubmitCookiesResponseObject struct {
	// Parameters for the display and wait login step
	DisplayAndWait MatrixBridgeAuthSubmitCookiesResponseObjectDisplayAndWait `json:"display_and_wait" api:"required"`
	Type           constant.DisplayAndWait                                   `json:"type" default:"display_and_wait"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayAndWait respjson.Field
		Type           respjson.Field
		Instructions   respjson.Field
		LoginID        respjson.Field
		StepID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitCookiesResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the display and wait login step
type MatrixBridgeAuthSubmitCookiesResponseObjectDisplayAndWait struct {
	// The type of thing to display
	//
	// Any of "qr", "emoji", "code", "nothing".
	Type string `json:"type" api:"required"`
	// The thing to display (raw data for QR, unicode emoji for emoji, plain string for
	// code)
	Data string `json:"data"`
	// An image containing the thing to display. If present, this is recommended over
	// using data directly. For emojis, the URL to the canonical image representation
	// of the emoji
	ImageURL string `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObjectDisplayAndWait) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthSubmitCookiesResponseObjectDisplayAndWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User input login step
type MatrixBridgeAuthSubmitCookiesResponseObject2 struct {
	Type constant.UserInput `json:"type" default:"user_input"`
	// Parameters for the user input login step
	UserInput MatrixBridgeAuthSubmitCookiesResponseObject2UserInput `json:"user_input" api:"required"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type         respjson.Field
		UserInput    respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitCookiesResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the user input login step
type MatrixBridgeAuthSubmitCookiesResponseObject2UserInput struct {
	// The list of fields that the user is requested to fill.
	Fields []MatrixBridgeAuthSubmitCookiesResponseObject2UserInputField `json:"fields" api:"required"`
	// A list of media attachments to show the user alongside the form fields.
	Attachments []MatrixBridgeAuthSubmitCookiesResponseObject2UserInputAttachment `json:"attachments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields      respjson.Field
		Attachments respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject2UserInput) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitCookiesResponseObject2UserInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that the user can fill.
type MatrixBridgeAuthSubmitCookiesResponseObject2UserInputField struct {
	// The internal ID of the field. This must be used as the key in the object when
	// submitting the data back to the bridge.
	ID string `json:"id" api:"required"`
	// The name of the field shown to the user.
	Name string `json:"name" api:"required"`
	// The type of field.
	//
	// Any of "username", "phone_number", "email", "password", "2fa_code", "token",
	// "url", "domain", "select".
	Type string `json:"type" api:"required"`
	// A default value that the client can pre-fill the field with.
	DefaultValue string `json:"default_value"`
	// A more detailed description of the field shown to the user.
	Description string `json:"description"`
	// For fields of type select, the valid options.
	Options []string `json:"options"`
	// A regular expression that the field value must match.
	Pattern string `json:"pattern" format:"regex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		DefaultValue respjson.Field
		Description  respjson.Field
		Options      respjson.Field
		Pattern      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject2UserInputField) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthSubmitCookiesResponseObject2UserInputField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media attachment to show the user.
type MatrixBridgeAuthSubmitCookiesResponseObject2UserInputAttachment struct {
	// The raw file content for the attachment encoded in base64.
	Content string `json:"content" api:"required"`
	// The filename for the media attachment.
	Filename string `json:"filename" api:"required"`
	// The type of media attachment, using the same media type identifiers as Matrix
	// attachments. Only some are supported.
	//
	// Any of "m.image", "m.audio".
	Type string `json:"type" api:"required"`
	// Optional but recommended metadata for the attachment. Can generally be derived
	// from the raw content if omitted.
	Info MatrixBridgeAuthSubmitCookiesResponseObject2UserInputAttachmentInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject2UserInputAttachment) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthSubmitCookiesResponseObject2UserInputAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional but recommended metadata for the attachment. Can generally be derived
// from the raw content if omitted.
type MatrixBridgeAuthSubmitCookiesResponseObject2UserInputAttachmentInfo struct {
	// The height of the media in pixels. Only applicable for images and videos.
	H float64 `json:"h"`
	// The MIME type for the media content.
	Mimetype string `json:"mimetype"`
	// The size of the media content in number of bytes. Strongly recommended to
	// include.
	Size float64 `json:"size"`
	// The width of the media in pixels. Only applicable for images and videos.
	W float64 `json:"w"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H           respjson.Field
		Mimetype    respjson.Field
		Size        respjson.Field
		W           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject2UserInputAttachmentInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthSubmitCookiesResponseObject2UserInputAttachmentInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cookie login step
type MatrixBridgeAuthSubmitCookiesResponseObject3 struct {
	// Parameters for the cookie login step
	Cookies MatrixBridgeAuthSubmitCookiesResponseObject3Cookies `json:"cookies" api:"required"`
	Type    constant.Cookies                                    `json:"type" default:"cookies"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cookies      respjson.Field
		Type         respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject3) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitCookiesResponseObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the cookie login step
type MatrixBridgeAuthSubmitCookiesResponseObject3Cookies struct {
	// The list of cookies or other stored data that must be extracted.
	Fields []MatrixBridgeAuthSubmitCookiesResponseObject3CookiesField `json:"fields" api:"required"`
	// The URL to open when using a webview to extract cookies.
	URL string `json:"url" api:"required" format:"uri"`
	// A JavaScript snippet that can extract some or all of the fields. The snippet
	// will evaluate to a promise that resolves when the relevant fields are found.
	// Fields that are not present in the promise result must be extracted another way.
	ExtractJs string `json:"extract_js"`
	// An optional user agent that the webview should use.
	UserAgent string `json:"user_agent"`
	// A regex pattern that the URL should match before the client closes the webview.
	//
	// The client may submit the login if the user closes the webview after all cookies
	// are collected even if this URL is not reached, but it should only automatically
	// close the webview after both cookies and the URL match.
	WaitForURLPattern string `json:"wait_for_url_pattern"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields            respjson.Field
		URL               respjson.Field
		ExtractJs         respjson.Field
		UserAgent         respjson.Field
		WaitForURLPattern respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject3Cookies) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitCookiesResponseObject3Cookies) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual cookie or other stored data item that must be extracted.
type MatrixBridgeAuthSubmitCookiesResponseObject3CookiesField struct {
	// The name of the item to extract.
	Name string `json:"name" api:"required"`
	// The type of data to extract.
	//
	// Any of "cookie", "local_storage", "request_header", "request_body", "special".
	Type string `json:"type" api:"required"`
	// For the `cookie` type, the domain of the cookie.
	CookieDomain string `json:"cookie_domain"`
	// For the `request_header` and `request_body` types, a regex that matches the URLs
	// from which the values can be extracted.
	RequestURLRegex string `json:"request_url_regex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Type            respjson.Field
		CookieDomain    respjson.Field
		RequestURLRegex respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject3CookiesField) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitCookiesResponseObject3CookiesField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Login complete
type MatrixBridgeAuthSubmitCookiesResponseObject4 struct {
	// Information about the completed login
	Complete MatrixBridgeAuthSubmitCookiesResponseObject4Complete `json:"complete" api:"required"`
	Type     constant.Complete                                    `json:"type" default:"complete"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete     respjson.Field
		Type         respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject4) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitCookiesResponseObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the completed login
type MatrixBridgeAuthSubmitCookiesResponseObject4Complete struct {
	// The unique ID of a login. Defined by the network connector.
	UserLoginID string `json:"user_login_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserLoginID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitCookiesResponseObject4Complete) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitCookiesResponseObject4Complete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MatrixBridgeAuthSubmitUserInputResponseUnion contains all possible properties
// and values from [MatrixBridgeAuthSubmitUserInputResponseObject],
// [MatrixBridgeAuthSubmitUserInputResponseObject2],
// [MatrixBridgeAuthSubmitUserInputResponseObject3],
// [MatrixBridgeAuthSubmitUserInputResponseObject4].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MatrixBridgeAuthSubmitUserInputResponseUnion struct {
	// This field is from variant [MatrixBridgeAuthSubmitUserInputResponseObject].
	DisplayAndWait MatrixBridgeAuthSubmitUserInputResponseObjectDisplayAndWait `json:"display_and_wait"`
	Type           string                                                      `json:"type"`
	Instructions   string                                                      `json:"instructions"`
	LoginID        string                                                      `json:"login_id"`
	StepID         string                                                      `json:"step_id"`
	// This field is from variant [MatrixBridgeAuthSubmitUserInputResponseObject2].
	UserInput MatrixBridgeAuthSubmitUserInputResponseObject2UserInput `json:"user_input"`
	// This field is from variant [MatrixBridgeAuthSubmitUserInputResponseObject3].
	Cookies MatrixBridgeAuthSubmitUserInputResponseObject3Cookies `json:"cookies"`
	// This field is from variant [MatrixBridgeAuthSubmitUserInputResponseObject4].
	Complete MatrixBridgeAuthSubmitUserInputResponseObject4Complete `json:"complete"`
	JSON     struct {
		DisplayAndWait respjson.Field
		Type           respjson.Field
		Instructions   respjson.Field
		LoginID        respjson.Field
		StepID         respjson.Field
		UserInput      respjson.Field
		Cookies        respjson.Field
		Complete       respjson.Field
		raw            string
	} `json:"-"`
}

func (u MatrixBridgeAuthSubmitUserInputResponseUnion) AsMatrixBridgeAuthSubmitUserInputResponseObject() (v MatrixBridgeAuthSubmitUserInputResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthSubmitUserInputResponseUnion) AsMatrixBridgeAuthSubmitUserInputResponseObject2() (v MatrixBridgeAuthSubmitUserInputResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthSubmitUserInputResponseUnion) AsMatrixBridgeAuthSubmitUserInputResponseObject3() (v MatrixBridgeAuthSubmitUserInputResponseObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthSubmitUserInputResponseUnion) AsMatrixBridgeAuthSubmitUserInputResponseObject4() (v MatrixBridgeAuthSubmitUserInputResponseObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MatrixBridgeAuthSubmitUserInputResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *MatrixBridgeAuthSubmitUserInputResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Display and wait login step
type MatrixBridgeAuthSubmitUserInputResponseObject struct {
	// Parameters for the display and wait login step
	DisplayAndWait MatrixBridgeAuthSubmitUserInputResponseObjectDisplayAndWait `json:"display_and_wait" api:"required"`
	Type           constant.DisplayAndWait                                     `json:"type" default:"display_and_wait"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayAndWait respjson.Field
		Type           respjson.Field
		Instructions   respjson.Field
		LoginID        respjson.Field
		StepID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitUserInputResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the display and wait login step
type MatrixBridgeAuthSubmitUserInputResponseObjectDisplayAndWait struct {
	// The type of thing to display
	//
	// Any of "qr", "emoji", "code", "nothing".
	Type string `json:"type" api:"required"`
	// The thing to display (raw data for QR, unicode emoji for emoji, plain string for
	// code)
	Data string `json:"data"`
	// An image containing the thing to display. If present, this is recommended over
	// using data directly. For emojis, the URL to the canonical image representation
	// of the emoji
	ImageURL string `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObjectDisplayAndWait) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthSubmitUserInputResponseObjectDisplayAndWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User input login step
type MatrixBridgeAuthSubmitUserInputResponseObject2 struct {
	Type constant.UserInput `json:"type" default:"user_input"`
	// Parameters for the user input login step
	UserInput MatrixBridgeAuthSubmitUserInputResponseObject2UserInput `json:"user_input" api:"required"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type         respjson.Field
		UserInput    respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitUserInputResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the user input login step
type MatrixBridgeAuthSubmitUserInputResponseObject2UserInput struct {
	// The list of fields that the user is requested to fill.
	Fields []MatrixBridgeAuthSubmitUserInputResponseObject2UserInputField `json:"fields" api:"required"`
	// A list of media attachments to show the user alongside the form fields.
	Attachments []MatrixBridgeAuthSubmitUserInputResponseObject2UserInputAttachment `json:"attachments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields      respjson.Field
		Attachments respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject2UserInput) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitUserInputResponseObject2UserInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that the user can fill.
type MatrixBridgeAuthSubmitUserInputResponseObject2UserInputField struct {
	// The internal ID of the field. This must be used as the key in the object when
	// submitting the data back to the bridge.
	ID string `json:"id" api:"required"`
	// The name of the field shown to the user.
	Name string `json:"name" api:"required"`
	// The type of field.
	//
	// Any of "username", "phone_number", "email", "password", "2fa_code", "token",
	// "url", "domain", "select".
	Type string `json:"type" api:"required"`
	// A default value that the client can pre-fill the field with.
	DefaultValue string `json:"default_value"`
	// A more detailed description of the field shown to the user.
	Description string `json:"description"`
	// For fields of type select, the valid options.
	Options []string `json:"options"`
	// A regular expression that the field value must match.
	Pattern string `json:"pattern" format:"regex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		DefaultValue respjson.Field
		Description  respjson.Field
		Options      respjson.Field
		Pattern      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject2UserInputField) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthSubmitUserInputResponseObject2UserInputField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media attachment to show the user.
type MatrixBridgeAuthSubmitUserInputResponseObject2UserInputAttachment struct {
	// The raw file content for the attachment encoded in base64.
	Content string `json:"content" api:"required"`
	// The filename for the media attachment.
	Filename string `json:"filename" api:"required"`
	// The type of media attachment, using the same media type identifiers as Matrix
	// attachments. Only some are supported.
	//
	// Any of "m.image", "m.audio".
	Type string `json:"type" api:"required"`
	// Optional but recommended metadata for the attachment. Can generally be derived
	// from the raw content if omitted.
	Info MatrixBridgeAuthSubmitUserInputResponseObject2UserInputAttachmentInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject2UserInputAttachment) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthSubmitUserInputResponseObject2UserInputAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional but recommended metadata for the attachment. Can generally be derived
// from the raw content if omitted.
type MatrixBridgeAuthSubmitUserInputResponseObject2UserInputAttachmentInfo struct {
	// The height of the media in pixels. Only applicable for images and videos.
	H float64 `json:"h"`
	// The MIME type for the media content.
	Mimetype string `json:"mimetype"`
	// The size of the media content in number of bytes. Strongly recommended to
	// include.
	Size float64 `json:"size"`
	// The width of the media in pixels. Only applicable for images and videos.
	W float64 `json:"w"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H           respjson.Field
		Mimetype    respjson.Field
		Size        respjson.Field
		W           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject2UserInputAttachmentInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthSubmitUserInputResponseObject2UserInputAttachmentInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cookie login step
type MatrixBridgeAuthSubmitUserInputResponseObject3 struct {
	// Parameters for the cookie login step
	Cookies MatrixBridgeAuthSubmitUserInputResponseObject3Cookies `json:"cookies" api:"required"`
	Type    constant.Cookies                                      `json:"type" default:"cookies"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cookies      respjson.Field
		Type         respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject3) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitUserInputResponseObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the cookie login step
type MatrixBridgeAuthSubmitUserInputResponseObject3Cookies struct {
	// The list of cookies or other stored data that must be extracted.
	Fields []MatrixBridgeAuthSubmitUserInputResponseObject3CookiesField `json:"fields" api:"required"`
	// The URL to open when using a webview to extract cookies.
	URL string `json:"url" api:"required" format:"uri"`
	// A JavaScript snippet that can extract some or all of the fields. The snippet
	// will evaluate to a promise that resolves when the relevant fields are found.
	// Fields that are not present in the promise result must be extracted another way.
	ExtractJs string `json:"extract_js"`
	// An optional user agent that the webview should use.
	UserAgent string `json:"user_agent"`
	// A regex pattern that the URL should match before the client closes the webview.
	//
	// The client may submit the login if the user closes the webview after all cookies
	// are collected even if this URL is not reached, but it should only automatically
	// close the webview after both cookies and the URL match.
	WaitForURLPattern string `json:"wait_for_url_pattern"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields            respjson.Field
		URL               respjson.Field
		ExtractJs         respjson.Field
		UserAgent         respjson.Field
		WaitForURLPattern respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject3Cookies) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitUserInputResponseObject3Cookies) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual cookie or other stored data item that must be extracted.
type MatrixBridgeAuthSubmitUserInputResponseObject3CookiesField struct {
	// The name of the item to extract.
	Name string `json:"name" api:"required"`
	// The type of data to extract.
	//
	// Any of "cookie", "local_storage", "request_header", "request_body", "special".
	Type string `json:"type" api:"required"`
	// For the `cookie` type, the domain of the cookie.
	CookieDomain string `json:"cookie_domain"`
	// For the `request_header` and `request_body` types, a regex that matches the URLs
	// from which the values can be extracted.
	RequestURLRegex string `json:"request_url_regex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Type            respjson.Field
		CookieDomain    respjson.Field
		RequestURLRegex respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject3CookiesField) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthSubmitUserInputResponseObject3CookiesField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Login complete
type MatrixBridgeAuthSubmitUserInputResponseObject4 struct {
	// Information about the completed login
	Complete MatrixBridgeAuthSubmitUserInputResponseObject4Complete `json:"complete" api:"required"`
	Type     constant.Complete                                      `json:"type" default:"complete"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete     respjson.Field
		Type         respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject4) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitUserInputResponseObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the completed login
type MatrixBridgeAuthSubmitUserInputResponseObject4Complete struct {
	// The unique ID of a login. Defined by the network connector.
	UserLoginID string `json:"user_login_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserLoginID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthSubmitUserInputResponseObject4Complete) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthSubmitUserInputResponseObject4Complete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MatrixBridgeAuthWaitForStepResponseUnion contains all possible properties and
// values from [MatrixBridgeAuthWaitForStepResponseObject],
// [MatrixBridgeAuthWaitForStepResponseObject2],
// [MatrixBridgeAuthWaitForStepResponseObject3],
// [MatrixBridgeAuthWaitForStepResponseObject4].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MatrixBridgeAuthWaitForStepResponseUnion struct {
	// This field is from variant [MatrixBridgeAuthWaitForStepResponseObject].
	DisplayAndWait MatrixBridgeAuthWaitForStepResponseObjectDisplayAndWait `json:"display_and_wait"`
	Type           string                                                  `json:"type"`
	Instructions   string                                                  `json:"instructions"`
	LoginID        string                                                  `json:"login_id"`
	StepID         string                                                  `json:"step_id"`
	// This field is from variant [MatrixBridgeAuthWaitForStepResponseObject2].
	UserInput MatrixBridgeAuthWaitForStepResponseObject2UserInput `json:"user_input"`
	// This field is from variant [MatrixBridgeAuthWaitForStepResponseObject3].
	Cookies MatrixBridgeAuthWaitForStepResponseObject3Cookies `json:"cookies"`
	// This field is from variant [MatrixBridgeAuthWaitForStepResponseObject4].
	Complete MatrixBridgeAuthWaitForStepResponseObject4Complete `json:"complete"`
	JSON     struct {
		DisplayAndWait respjson.Field
		Type           respjson.Field
		Instructions   respjson.Field
		LoginID        respjson.Field
		StepID         respjson.Field
		UserInput      respjson.Field
		Cookies        respjson.Field
		Complete       respjson.Field
		raw            string
	} `json:"-"`
}

func (u MatrixBridgeAuthWaitForStepResponseUnion) AsMatrixBridgeAuthWaitForStepResponseObject() (v MatrixBridgeAuthWaitForStepResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthWaitForStepResponseUnion) AsMatrixBridgeAuthWaitForStepResponseObject2() (v MatrixBridgeAuthWaitForStepResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthWaitForStepResponseUnion) AsMatrixBridgeAuthWaitForStepResponseObject3() (v MatrixBridgeAuthWaitForStepResponseObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MatrixBridgeAuthWaitForStepResponseUnion) AsMatrixBridgeAuthWaitForStepResponseObject4() (v MatrixBridgeAuthWaitForStepResponseObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MatrixBridgeAuthWaitForStepResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *MatrixBridgeAuthWaitForStepResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Display and wait login step
type MatrixBridgeAuthWaitForStepResponseObject struct {
	// Parameters for the display and wait login step
	DisplayAndWait MatrixBridgeAuthWaitForStepResponseObjectDisplayAndWait `json:"display_and_wait" api:"required"`
	Type           constant.DisplayAndWait                                 `json:"type" default:"display_and_wait"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayAndWait respjson.Field
		Type           respjson.Field
		Instructions   respjson.Field
		LoginID        respjson.Field
		StepID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the display and wait login step
type MatrixBridgeAuthWaitForStepResponseObjectDisplayAndWait struct {
	// The type of thing to display
	//
	// Any of "qr", "emoji", "code", "nothing".
	Type string `json:"type" api:"required"`
	// The thing to display (raw data for QR, unicode emoji for emoji, plain string for
	// code)
	Data string `json:"data"`
	// An image containing the thing to display. If present, this is recommended over
	// using data directly. For emojis, the URL to the canonical image representation
	// of the emoji
	ImageURL string `json:"image_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObjectDisplayAndWait) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObjectDisplayAndWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User input login step
type MatrixBridgeAuthWaitForStepResponseObject2 struct {
	Type constant.UserInput `json:"type" default:"user_input"`
	// Parameters for the user input login step
	UserInput MatrixBridgeAuthWaitForStepResponseObject2UserInput `json:"user_input" api:"required"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type         respjson.Field
		UserInput    respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the user input login step
type MatrixBridgeAuthWaitForStepResponseObject2UserInput struct {
	// The list of fields that the user is requested to fill.
	Fields []MatrixBridgeAuthWaitForStepResponseObject2UserInputField `json:"fields" api:"required"`
	// A list of media attachments to show the user alongside the form fields.
	Attachments []MatrixBridgeAuthWaitForStepResponseObject2UserInputAttachment `json:"attachments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields      respjson.Field
		Attachments respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject2UserInput) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObject2UserInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field that the user can fill.
type MatrixBridgeAuthWaitForStepResponseObject2UserInputField struct {
	// The internal ID of the field. This must be used as the key in the object when
	// submitting the data back to the bridge.
	ID string `json:"id" api:"required"`
	// The name of the field shown to the user.
	Name string `json:"name" api:"required"`
	// The type of field.
	//
	// Any of "username", "phone_number", "email", "password", "2fa_code", "token",
	// "url", "domain", "select".
	Type string `json:"type" api:"required"`
	// A default value that the client can pre-fill the field with.
	DefaultValue string `json:"default_value"`
	// A more detailed description of the field shown to the user.
	Description string `json:"description"`
	// For fields of type select, the valid options.
	Options []string `json:"options"`
	// A regular expression that the field value must match.
	Pattern string `json:"pattern" format:"regex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		DefaultValue respjson.Field
		Description  respjson.Field
		Options      respjson.Field
		Pattern      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject2UserInputField) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObject2UserInputField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media attachment to show the user.
type MatrixBridgeAuthWaitForStepResponseObject2UserInputAttachment struct {
	// The raw file content for the attachment encoded in base64.
	Content string `json:"content" api:"required"`
	// The filename for the media attachment.
	Filename string `json:"filename" api:"required"`
	// The type of media attachment, using the same media type identifiers as Matrix
	// attachments. Only some are supported.
	//
	// Any of "m.image", "m.audio".
	Type string `json:"type" api:"required"`
	// Optional but recommended metadata for the attachment. Can generally be derived
	// from the raw content if omitted.
	Info MatrixBridgeAuthWaitForStepResponseObject2UserInputAttachmentInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Filename    respjson.Field
		Type        respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject2UserInputAttachment) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthWaitForStepResponseObject2UserInputAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional but recommended metadata for the attachment. Can generally be derived
// from the raw content if omitted.
type MatrixBridgeAuthWaitForStepResponseObject2UserInputAttachmentInfo struct {
	// The height of the media in pixels. Only applicable for images and videos.
	H float64 `json:"h"`
	// The MIME type for the media content.
	Mimetype string `json:"mimetype"`
	// The size of the media content in number of bytes. Strongly recommended to
	// include.
	Size float64 `json:"size"`
	// The width of the media in pixels. Only applicable for images and videos.
	W float64 `json:"w"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H           respjson.Field
		Mimetype    respjson.Field
		Size        respjson.Field
		W           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject2UserInputAttachmentInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *MatrixBridgeAuthWaitForStepResponseObject2UserInputAttachmentInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cookie login step
type MatrixBridgeAuthWaitForStepResponseObject3 struct {
	// Parameters for the cookie login step
	Cookies MatrixBridgeAuthWaitForStepResponseObject3Cookies `json:"cookies" api:"required"`
	Type    constant.Cookies                                  `json:"type" default:"cookies"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cookies      respjson.Field
		Type         respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject3) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the cookie login step
type MatrixBridgeAuthWaitForStepResponseObject3Cookies struct {
	// The list of cookies or other stored data that must be extracted.
	Fields []MatrixBridgeAuthWaitForStepResponseObject3CookiesField `json:"fields" api:"required"`
	// The URL to open when using a webview to extract cookies.
	URL string `json:"url" api:"required" format:"uri"`
	// A JavaScript snippet that can extract some or all of the fields. The snippet
	// will evaluate to a promise that resolves when the relevant fields are found.
	// Fields that are not present in the promise result must be extracted another way.
	ExtractJs string `json:"extract_js"`
	// An optional user agent that the webview should use.
	UserAgent string `json:"user_agent"`
	// A regex pattern that the URL should match before the client closes the webview.
	//
	// The client may submit the login if the user closes the webview after all cookies
	// are collected even if this URL is not reached, but it should only automatically
	// close the webview after both cookies and the URL match.
	WaitForURLPattern string `json:"wait_for_url_pattern"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields            respjson.Field
		URL               respjson.Field
		ExtractJs         respjson.Field
		UserAgent         respjson.Field
		WaitForURLPattern respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject3Cookies) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObject3Cookies) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual cookie or other stored data item that must be extracted.
type MatrixBridgeAuthWaitForStepResponseObject3CookiesField struct {
	// The name of the item to extract.
	Name string `json:"name" api:"required"`
	// The type of data to extract.
	//
	// Any of "cookie", "local_storage", "request_header", "request_body", "special".
	Type string `json:"type" api:"required"`
	// For the `cookie` type, the domain of the cookie.
	CookieDomain string `json:"cookie_domain"`
	// For the `request_header` and `request_body` types, a regex that matches the URLs
	// from which the values can be extracted.
	RequestURLRegex string `json:"request_url_regex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Type            respjson.Field
		CookieDomain    respjson.Field
		RequestURLRegex respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject3CookiesField) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObject3CookiesField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Login complete
type MatrixBridgeAuthWaitForStepResponseObject4 struct {
	// Information about the completed login
	Complete MatrixBridgeAuthWaitForStepResponseObject4Complete `json:"complete" api:"required"`
	Type     constant.Complete                                  `json:"type" default:"complete"`
	// Human-readable instructions for completing this login step.
	Instructions string `json:"instructions"`
	// An identifier for the current login process. Must be passed to execute more
	// steps of the login.
	LoginID string `json:"login_id"`
	// An unique ID identifying this step. This can be used to implement special
	// behavior in clients.
	StepID string `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complete     respjson.Field
		Type         respjson.Field
		Instructions respjson.Field
		LoginID      respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject4) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the completed login
type MatrixBridgeAuthWaitForStepResponseObject4Complete struct {
	// The unique ID of a login. Defined by the network connector.
	UserLoginID string `json:"user_login_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserLoginID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWaitForStepResponseObject4Complete) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWaitForStepResponseObject4Complete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Info about the bridge and user
type MatrixBridgeAuthWhoamiResponse struct {
	// The Matrix user ID of the bridge bot.
	BridgeBot string `json:"bridge_bot" api:"required" format:"matrix_user_id"`
	// The command prefix used by this bridge.
	CommandPrefix string `json:"command_prefix" api:"required"`
	// The server name the bridge is running on.
	Homeserver string `json:"homeserver" api:"required"`
	// The login flows that the bridge supports.
	LoginFlows []MatrixBridgeAuthWhoamiResponseLoginFlow `json:"login_flows" api:"required"`
	// The logins of the user who made the /whoami call
	Logins []MatrixBridgeAuthWhoamiResponseLogin `json:"logins" api:"required"`
	// Info about the network that the bridge is bridging to.
	Network MatrixBridgeAuthWhoamiResponseNetwork `json:"network" api:"required"`
	// The Matrix management room ID of the user who made the /whoami call.
	ManagementRoom string `json:"management_room" format:"matrix_room_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BridgeBot      respjson.Field
		CommandPrefix  respjson.Field
		Homeserver     respjson.Field
		LoginFlows     respjson.Field
		Logins         respjson.Field
		Network        respjson.Field
		ManagementRoom respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWhoamiResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWhoamiResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual login flow which can be used to sign into the remote network.
type MatrixBridgeAuthWhoamiResponseLoginFlow struct {
	// An internal ID that is passed to the /login/start call to start a login with
	// this flow.
	ID string `json:"id" api:"required"`
	// A human-readable description of the login flow.
	Description string `json:"description" api:"required"`
	// A human-readable name for the login flow.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWhoamiResponseLoginFlow) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWhoamiResponseLoginFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The info of an individual login
type MatrixBridgeAuthWhoamiResponseLogin struct {
	// The unique ID of a login. Defined by the network connector.
	ID string `json:"id" api:"required"`
	// A human-readable name for the login. Defined by the network connector.
	Name string `json:"name" api:"required"`
	// The profile info of the logged-in user on the remote network.
	Profile MatrixBridgeAuthWhoamiResponseLoginProfile `json:"profile" api:"required"`
	// The connection status of an individual login
	State MatrixBridgeAuthWhoamiResponseLoginState `json:"state" api:"required"`
	// The personal filtering space room ID for this login.
	SpaceRoom string `json:"space_room" format:"matrix_room_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Profile     respjson.Field
		State       respjson.Field
		SpaceRoom   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWhoamiResponseLogin) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWhoamiResponseLogin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The profile info of the logged-in user on the remote network.
type MatrixBridgeAuthWhoamiResponseLoginProfile struct {
	// The user's avatar
	Avatar string `json:"avatar" format:"mxc"`
	// The user's email address
	Email string `json:"email" format:"email"`
	// The user's displayname
	Name string `json:"name"`
	// The user's phone number
	Phone string `json:"phone" format:"phone"`
	// The user's username
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Avatar      respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		Phone       respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWhoamiResponseLoginProfile) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWhoamiResponseLoginProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The connection status of an individual login
type MatrixBridgeAuthWhoamiResponseLoginState struct {
	// The current state of this login.
	//
	// Any of "CONNECTING", "CONNECTED", "TRANSIENT_DISCONNECT", "BAD_CREDENTIALS",
	// "UNKNOWN_ERROR".
	StateEvent string `json:"state_event" api:"required"`
	// The time when the state was last updated.
	Timestamp float64 `json:"timestamp" api:"required" format:"unix milliseconds"`
	// An error code defined by the network connector.
	Error string `json:"error"`
	// Additional arbitrary info provided by the network connector.
	Info any `json:"info"`
	// A human-readable error message defined by the network connector.
	Message string `json:"message"`
	// A reason code for non-error states that aren't exactly successes either.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StateEvent  respjson.Field
		Timestamp   respjson.Field
		Error       respjson.Field
		Info        respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWhoamiResponseLoginState) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWhoamiResponseLoginState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Info about the network that the bridge is bridging to.
type MatrixBridgeAuthWhoamiResponseNetwork struct {
	// An identifier uniquely identifying the bridge software.
	BeeperBridgeType string `json:"beeper_bridge_type" api:"required"`
	// The displayname of the network.
	Displayname string `json:"displayname" api:"required"`
	// The icon of the network as a `mxc://` URI.
	NetworkIcon string `json:"network_icon" api:"required" format:"mxc"`
	// An identifier uniquely identifying the network.
	NetworkID string `json:"network_id" api:"required"`
	// The URL to the website of the network.
	NetworkURL string `json:"network_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BeeperBridgeType respjson.Field
		Displayname      respjson.Field
		NetworkIcon      respjson.Field
		NetworkID        respjson.Field
		NetworkURL       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeAuthWhoamiResponseNetwork) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeAuthWhoamiResponseNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixBridgeAuthLogoutParams struct {
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	paramObj
}

type MatrixBridgeAuthStartLoginParams struct {
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	// An existing login ID to re-login as. If this is specified and the user logs into
	// a different account, the provided ID will be logged out.
	LoginID param.Opt[string] `query:"login_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MatrixBridgeAuthStartLoginParams]'s query parameters as
// `url.Values`.
func (r MatrixBridgeAuthStartLoginParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MatrixBridgeAuthSubmitCookiesParams struct {
	BridgeID       string `path:"bridgeID" api:"required" json:"-"`
	LoginProcessID string `path:"loginProcessID" api:"required" json:"-"`
	Body           map[string]string
	paramObj
}

func (r MatrixBridgeAuthSubmitCookiesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *MatrixBridgeAuthSubmitCookiesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixBridgeAuthSubmitUserInputParams struct {
	BridgeID       string `path:"bridgeID" api:"required" json:"-"`
	LoginProcessID string `path:"loginProcessID" api:"required" json:"-"`
	Body           map[string]string
	paramObj
}

func (r MatrixBridgeAuthSubmitUserInputParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *MatrixBridgeAuthSubmitUserInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixBridgeAuthWaitForStepParams struct {
	BridgeID       string `path:"bridgeID" api:"required" json:"-"`
	LoginProcessID string `path:"loginProcessID" api:"required" json:"-"`
	paramObj
}

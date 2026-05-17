// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
	"github.com/beeper/desktop-api-go/v5/shared"
	"github.com/beeper/desktop-api-go/v5/shared/constant"
)

// Manage bridge-backed account types, connections, and login sessions
//
// BridgeService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeService] method instead.
type BridgeService struct {
	Options []option.RequestOption
	// Available bridges, bridge logins, login sessions for connect and reconnect
	// flows, and advanced network capabilities.
	LoginFlows  BridgeLoginFlowService
	Connections BridgeConnectionService
	// Available bridges, bridge logins, login sessions for connect and reconnect
	// flows, and advanced network capabilities.
	LoginSessions BridgeLoginSessionService
}

// NewBridgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBridgeService(opts ...option.RequestOption) (r BridgeService) {
	r = BridgeService{}
	r.Options = opts
	r.LoginFlows = NewBridgeLoginFlowService(opts...)
	r.Connections = NewBridgeConnectionService(opts...)
	r.LoginSessions = NewBridgeLoginSessionService(opts...)
	return
}

// Get one bridge, including the chat accounts connected through it.
func (r *BridgeService) Get(ctx context.Context, bridgeID string, opts ...option.RequestOption) (res *BridgeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List available bridges. A bridge is a chat-network connector that can connect or
// reconnect chat accounts. Connected accounts use the same Account schema as GET
// /v1/accounts.
func (r *BridgeService) List(ctx context.Context, opts ...option.RequestOption) (res *BridgeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/bridges"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get advanced network capabilities for a bridge. This endpoint is intended for
// clients that build custom connect or chat-creation flows.
func (r *BridgeService) GetCapabilities(ctx context.Context, bridgeID string, opts ...option.RequestOption) (res *ProvisioningCapabilities, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s/capabilities", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Available bridge that can connect or reconnect chat accounts.
type Bridge struct {
	// Bridge ID. Use with bridge endpoints.
	ID string `json:"id" api:"required"`
	// Connected accounts for this bridge. Uses the same Account schema as GET
	// /v1/accounts.
	Accounts []Account `json:"accounts" api:"required"`
	// Number of active accounts for this network on this device.
	ActiveAccountCount int64 `json:"activeAccountCount" api:"required"`
	// Human-friendly bridge name shown in Beeper.
	DisplayName string `json:"displayName" api:"required"`
	// Where accounts for this bridge run: on this device or in Beeper Cloud.
	//
	// Any of "cloud", "self-hosted", "local", "platform-sdk".
	Provider BridgeProvider `json:"provider" api:"required"`
	// Whether this bridge can currently be used to connect new accounts.
	//
	// Any of "available", "connected", "limit_reached", "temporarily_unavailable",
	// "disabled".
	Status BridgeStatus `json:"status" api:"required"`
	// Whether this bridge can have multiple active accounts for the same network.
	SupportsMultipleAccounts bool `json:"supportsMultipleAccounts" api:"required"`
	// Underlying bridge type, such as matrix, discordgo, slackgo, whatsapp, telegram,
	// or twitter.
	Type string `json:"type" api:"required"`
	// Network grouping used for account counts and limits.
	Network string `json:"network"`
	// Human-friendly status text matching Beeper account management language.
	StatusText string `json:"statusText"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Accounts                 respjson.Field
		ActiveAccountCount       respjson.Field
		DisplayName              respjson.Field
		Provider                 respjson.Field
		Status                   respjson.Field
		SupportsMultipleAccounts respjson.Field
		Type                     respjson.Field
		Network                  respjson.Field
		StatusText               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Bridge) RawJSON() string { return r.JSON.raw }
func (r *Bridge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where accounts for this bridge run: on this device or in Beeper Cloud.
type BridgeProvider string

const (
	BridgeProviderCloud       BridgeProvider = "cloud"
	BridgeProviderSelfHosted  BridgeProvider = "self-hosted"
	BridgeProviderLocal       BridgeProvider = "local"
	BridgeProviderPlatformSDK BridgeProvider = "platform-sdk"
)

// Whether this bridge can currently be used to connect new accounts.
type BridgeStatus string

const (
	BridgeStatusAvailable              BridgeStatus = "available"
	BridgeStatusConnected              BridgeStatus = "connected"
	BridgeStatusLimitReached           BridgeStatus = "limit_reached"
	BridgeStatusTemporarilyUnavailable BridgeStatus = "temporarily_unavailable"
	BridgeStatusDisabled               BridgeStatus = "disabled"
)

type CookieField struct {
	// Field ID to send back in the fields object.
	ID string `json:"id" api:"required"`
	// Cookie, header, or local storage key to collect.
	Name string `json:"name"`
	// Browser storage source for this value.
	//
	// Any of "cookie", "header", "local_storage".
	Type CookieFieldType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CookieField) RawJSON() string { return r.JSON.raw }
func (r *CookieField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser storage source for this value.
type CookieFieldType string

const (
	CookieFieldTypeCookie       CookieFieldType = "cookie"
	CookieFieldTypeHeader       CookieFieldType = "header"
	CookieFieldTypeLocalStorage CookieFieldType = "local_storage"
)

// Disappearing-message timer capability.
type DisappearingTimerCapability struct {
	// Any of "", "after_read", "after_send".
	Types []string `json:"types" api:"required"`
	// Any of true.
	OmitEmptyTimer bool    `json:"omit_empty_timer"`
	Timers         []int64 `json:"timers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Types          respjson.Field
		OmitEmptyTimer respjson.Field
		Timers         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisappearingTimerCapability) RawJSON() string { return r.JSON.raw }
func (r *DisappearingTimerCapability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Group creation field capability.
type GroupFieldCapability struct {
	Allowed   bool  `json:"allowed" api:"required"`
	MaxLength int64 `json:"max_length"`
	MinLength int64 `json:"min_length"`
	Required  bool  `json:"required"`
	// Disappearing-message timer capability.
	Settings DisappearingTimerCapability `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		MaxLength   respjson.Field
		MinLength   respjson.Field
		Required    respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupFieldCapability) RawJSON() string { return r.JSON.raw }
func (r *GroupFieldCapability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Group creation capabilities for one group type.
type GroupTypeCapabilities struct {
	TypeDescription string `json:"type_description" api:"required"`
	// Group creation field capability.
	Avatar GroupFieldCapability `json:"avatar"`
	// Group creation field capability.
	Disappear GroupFieldCapability `json:"disappear"`
	// Group creation field capability.
	Name GroupFieldCapability `json:"name"`
	// Group creation field capability.
	Parent GroupFieldCapability `json:"parent"`
	// Group creation field capability.
	Participants GroupFieldCapability `json:"participants"`
	// Group creation field capability.
	Topic GroupFieldCapability `json:"topic"`
	// Group creation field capability.
	Username GroupFieldCapability `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TypeDescription respjson.Field
		Avatar          respjson.Field
		Disappear       respjson.Field
		Name            respjson.Field
		Parent          respjson.Field
		Participants    respjson.Field
		Topic           respjson.Field
		Username        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupTypeCapabilities) RawJSON() string { return r.JSON.raw }
func (r *GroupTypeCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connect or reconnect flow option for a bridge.
type LoginFlow struct {
	// Flow ID to pass when creating a bridge login session.
	ID string `json:"id" api:"required"`
	// Short explanation for when to use this flow, when provided.
	Description string `json:"description"`
	// Display name for the flow, when provided.
	Name string `json:"name"`
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
func (r LoginFlow) RawJSON() string { return r.JSON.raw }
func (r *LoginFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginInputField struct {
	// Field ID to send back in the fields object.
	ID string `json:"id" api:"required"`
	// Initial field value, when provided by the network.
	InitialValue string `json:"initialValue"`
	// Field label to show to the user.
	Label string `json:"label"`
	// True if the user can leave this field empty.
	Optional bool `json:"optional"`
	// Placeholder text to show when the field is empty.
	Placeholder string `json:"placeholder"`
	// Suggested input type, such as text, password, or email.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		InitialValue respjson.Field
		Label        respjson.Field
		Optional     respjson.Field
		Placeholder  respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginInputField) RawJSON() string { return r.JSON.raw }
func (r *LoginInputField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginSession struct {
	// Bridge ID.
	BridgeID string `json:"bridgeID" api:"required"`
	// Temporary bridge login session ID.
	LoginSessionID string `json:"loginSessionID" api:"required"`
	// Any of "waiting_for_input", "waiting_for_cookies", "waiting_for_display",
	// "complete", "cancelled", "failed".
	Status LoginSessionStatus `json:"status" api:"required"`
	// A chat account added to Beeper.
	Account Account `json:"account"`
	// Chat account ID for reconnect flows, when known.
	AccountID string `json:"accountID"`
	// Step the client should show or complete next. Omitted when the session is
	// complete, cancelled, or failed.
	CurrentStep LoginSessionCurrentStepUnion `json:"currentStep"`
	Error       shared.APIError              `json:"error"`
	// Signed-in identity for a bridge. One bridge login can contain multiple chat
	// accounts.
	Login LoginSessionLogin `json:"login"`
	// Bridge login ID for reconnect flows, when known.
	LoginID string `json:"loginID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BridgeID       respjson.Field
		LoginSessionID respjson.Field
		Status         respjson.Field
		Account        respjson.Field
		AccountID      respjson.Field
		CurrentStep    respjson.Field
		Error          respjson.Field
		Login          respjson.Field
		LoginID        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSession) RawJSON() string { return r.JSON.raw }
func (r *LoginSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginSessionStatus string

const (
	LoginSessionStatusWaitingForInput   LoginSessionStatus = "waiting_for_input"
	LoginSessionStatusWaitingForCookies LoginSessionStatus = "waiting_for_cookies"
	LoginSessionStatusWaitingForDisplay LoginSessionStatus = "waiting_for_display"
	LoginSessionStatusComplete          LoginSessionStatus = "complete"
	LoginSessionStatusCancelled         LoginSessionStatus = "cancelled"
	LoginSessionStatusFailed            LoginSessionStatus = "failed"
)

// LoginSessionCurrentStepUnion contains all possible properties and values from
// [LoginSessionCurrentStepUserInput], [LoginSessionCurrentStepCookies],
// [LoginSessionCurrentStepDisplayAndWait], [LoginSessionCurrentStepComplete].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LoginSessionCurrentStepUnion struct {
	// This field is a union of [[]LoginInputField], [[]CookieField]
	Fields LoginSessionCurrentStepUnionFields `json:"fields"`
	StepID string                             `json:"stepID"`
	Type   string                             `json:"type"`
	// This field is from variant [LoginSessionCurrentStepUserInput].
	Attachments  []any  `json:"attachments"`
	Instructions string `json:"instructions"`
	// This field is from variant [LoginSessionCurrentStepCookies].
	URL string `json:"url"`
	// This field is from variant [LoginSessionCurrentStepCookies].
	ExpectedFinalURLRegex string `json:"expectedFinalURLRegex"`
	// This field is from variant [LoginSessionCurrentStepCookies].
	ExtractJs string `json:"extractJS"`
	// This field is from variant [LoginSessionCurrentStepCookies].
	UserAgent string `json:"userAgent"`
	// This field is from variant [LoginSessionCurrentStepDisplayAndWait].
	Display LoginSessionCurrentStepDisplayAndWaitDisplayUnion `json:"display"`
	// This field is from variant [LoginSessionCurrentStepComplete].
	Account Account `json:"account"`
	// This field is from variant [LoginSessionCurrentStepComplete].
	Login LoginSessionCurrentStepCompleteLogin `json:"login"`
	JSON  struct {
		Fields                respjson.Field
		StepID                respjson.Field
		Type                  respjson.Field
		Attachments           respjson.Field
		Instructions          respjson.Field
		URL                   respjson.Field
		ExpectedFinalURLRegex respjson.Field
		ExtractJs             respjson.Field
		UserAgent             respjson.Field
		Display               respjson.Field
		Account               respjson.Field
		Login                 respjson.Field
		raw                   string
	} `json:"-"`
}

func (u LoginSessionCurrentStepUnion) AsUserInput() (v LoginSessionCurrentStepUserInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LoginSessionCurrentStepUnion) AsCookies() (v LoginSessionCurrentStepCookies) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LoginSessionCurrentStepUnion) AsDisplayAndWait() (v LoginSessionCurrentStepDisplayAndWait) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LoginSessionCurrentStepUnion) AsComplete() (v LoginSessionCurrentStepComplete) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LoginSessionCurrentStepUnion) RawJSON() string { return u.JSON.raw }

func (r *LoginSessionCurrentStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LoginSessionCurrentStepUnionFields is an implicit subunion of
// [LoginSessionCurrentStepUnion]. LoginSessionCurrentStepUnionFields provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LoginSessionCurrentStepUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfLoginInputFieldArray OfCookieFieldArray]
type LoginSessionCurrentStepUnionFields struct {
	// This field will be present if the value is a [[]LoginInputField] instead of an
	// object.
	OfLoginInputFieldArray []LoginInputField `json:",inline"`
	// This field will be present if the value is a [[]CookieField] instead of an
	// object.
	OfCookieFieldArray []CookieField `json:",inline"`
	JSON               struct {
		OfLoginInputFieldArray respjson.Field
		OfCookieFieldArray     respjson.Field
		raw                    string
	} `json:"-"`
}

func (r *LoginSessionCurrentStepUnionFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginSessionCurrentStepUserInput struct {
	Fields      []LoginInputField  `json:"fields" api:"required"`
	StepID      string             `json:"stepID" api:"required"`
	Type        constant.UserInput `json:"type" default:"user_input"`
	Attachments []any              `json:"attachments"`
	// User-facing instructions for this step.
	Instructions string `json:"instructions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields       respjson.Field
		StepID       respjson.Field
		Type         respjson.Field
		Attachments  respjson.Field
		Instructions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSessionCurrentStepUserInput) RawJSON() string { return r.JSON.raw }
func (r *LoginSessionCurrentStepUserInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginSessionCurrentStepCookies struct {
	Fields []CookieField    `json:"fields" api:"required"`
	StepID string           `json:"stepID" api:"required"`
	Type   constant.Cookies `json:"type" default:"cookies"`
	// URL to open for the user.
	URL string `json:"url" api:"required"`
	// Regular expression that identifies the final URL after sign-in.
	ExpectedFinalURLRegex string `json:"expectedFinalURLRegex"`
	// Optional extraction script for browser-based sign-in helpers. Treat as an opaque
	// helper value.
	ExtractJs string `json:"extractJS"`
	// User-facing instructions for this browser step.
	Instructions string `json:"instructions"`
	// Suggested user agent for the browser session.
	UserAgent string `json:"userAgent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields                respjson.Field
		StepID                respjson.Field
		Type                  respjson.Field
		URL                   respjson.Field
		ExpectedFinalURLRegex respjson.Field
		ExtractJs             respjson.Field
		Instructions          respjson.Field
		UserAgent             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSessionCurrentStepCookies) RawJSON() string { return r.JSON.raw }
func (r *LoginSessionCurrentStepCookies) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginSessionCurrentStepDisplayAndWait struct {
	Display LoginSessionCurrentStepDisplayAndWaitDisplayUnion `json:"display" api:"required"`
	StepID  string                                            `json:"stepID" api:"required"`
	Type    constant.DisplayAndWait                           `json:"type" default:"display_and_wait"`
	// User-facing instructions for this step.
	Instructions string `json:"instructions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Display      respjson.Field
		StepID       respjson.Field
		Type         respjson.Field
		Instructions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSessionCurrentStepDisplayAndWait) RawJSON() string { return r.JSON.raw }
func (r *LoginSessionCurrentStepDisplayAndWait) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LoginSessionCurrentStepDisplayAndWaitDisplayUnion contains all possible
// properties and values from [LoginSessionCurrentStepDisplayAndWaitDisplayQrCode],
// [LoginSessionCurrentStepDisplayAndWaitDisplayEmoji],
// [LoginSessionCurrentStepDisplayAndWaitDisplayEmpty].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LoginSessionCurrentStepDisplayAndWaitDisplayUnion struct {
	// This field is from variant [LoginSessionCurrentStepDisplayAndWaitDisplayQrCode].
	Data string `json:"data"`
	Type string `json:"type"`
	// This field is from variant [LoginSessionCurrentStepDisplayAndWaitDisplayEmoji].
	ImageURL string `json:"imageURL"`
	JSON     struct {
		Data     respjson.Field
		Type     respjson.Field
		ImageURL respjson.Field
		raw      string
	} `json:"-"`
}

func (u LoginSessionCurrentStepDisplayAndWaitDisplayUnion) AsQrCode() (v LoginSessionCurrentStepDisplayAndWaitDisplayQrCode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LoginSessionCurrentStepDisplayAndWaitDisplayUnion) AsEmoji() (v LoginSessionCurrentStepDisplayAndWaitDisplayEmoji) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LoginSessionCurrentStepDisplayAndWaitDisplayUnion) AsEmpty() (v LoginSessionCurrentStepDisplayAndWaitDisplayEmpty) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LoginSessionCurrentStepDisplayAndWaitDisplayUnion) RawJSON() string { return u.JSON.raw }

func (r *LoginSessionCurrentStepDisplayAndWaitDisplayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginSessionCurrentStepDisplayAndWaitDisplayQrCode struct {
	Data string      `json:"data" api:"required"`
	Type constant.Qr `json:"type" default:"qr"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSessionCurrentStepDisplayAndWaitDisplayQrCode) RawJSON() string { return r.JSON.raw }
func (r *LoginSessionCurrentStepDisplayAndWaitDisplayQrCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginSessionCurrentStepDisplayAndWaitDisplayEmoji struct {
	ImageURL string         `json:"imageURL" api:"required"`
	Type     constant.Emoji `json:"type" default:"emoji"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSessionCurrentStepDisplayAndWaitDisplayEmoji) RawJSON() string { return r.JSON.raw }
func (r *LoginSessionCurrentStepDisplayAndWaitDisplayEmoji) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginSessionCurrentStepDisplayAndWaitDisplayEmpty struct {
	Type constant.Nothing `json:"type" default:"nothing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSessionCurrentStepDisplayAndWaitDisplayEmpty) RawJSON() string { return r.JSON.raw }
func (r *LoginSessionCurrentStepDisplayAndWaitDisplayEmpty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginSessionCurrentStepComplete struct {
	Type constant.Complete `json:"type" default:"complete"`
	// A chat account added to Beeper.
	Account Account `json:"account"`
	// Completion instructions, when provided.
	Instructions string `json:"instructions"`
	// Signed-in identity for a bridge. One bridge login can contain multiple chat
	// accounts.
	Login  LoginSessionCurrentStepCompleteLogin `json:"login"`
	StepID string                               `json:"stepID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type         respjson.Field
		Account      respjson.Field
		Instructions respjson.Field
		Login        respjson.Field
		StepID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSessionCurrentStepComplete) RawJSON() string { return r.JSON.raw }
func (r *LoginSessionCurrentStepComplete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in identity for a bridge. One bridge login can contain multiple chat
// accounts.
type LoginSessionCurrentStepCompleteLogin struct {
	// Bridge ID.
	BridgeID string `json:"bridgeID" api:"required"`
	// Bridge login ID.
	LoginID string `json:"loginID" api:"required"`
	// Any of "current-device", "all-devices".
	RemoveScopes []string `json:"removeScopes" api:"required"`
	// Any of "connected", "connecting", "needs_login", "logged_out", "unknown".
	Status string `json:"status" api:"required"`
	// Chat accounts that belong to this bridge login, when known.
	AccountIDs []string `json:"accountIDs"`
	// Human-friendly bridge login status text.
	StatusText string `json:"statusText"`
	// User the account belongs to.
	User shared.User `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BridgeID     respjson.Field
		LoginID      respjson.Field
		RemoveScopes respjson.Field
		Status       respjson.Field
		AccountIDs   respjson.Field
		StatusText   respjson.Field
		User         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSessionCurrentStepCompleteLogin) RawJSON() string { return r.JSON.raw }
func (r *LoginSessionCurrentStepCompleteLogin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in identity for a bridge. One bridge login can contain multiple chat
// accounts.
type LoginSessionLogin struct {
	// Bridge ID.
	BridgeID string `json:"bridgeID" api:"required"`
	// Bridge login ID.
	LoginID string `json:"loginID" api:"required"`
	// Any of "current-device", "all-devices".
	RemoveScopes []string `json:"removeScopes" api:"required"`
	// Any of "connected", "connecting", "needs_login", "logged_out", "unknown".
	Status string `json:"status" api:"required"`
	// Chat accounts that belong to this bridge login, when known.
	AccountIDs []string `json:"accountIDs"`
	// Human-friendly bridge login status text.
	StatusText string `json:"statusText"`
	// User the account belongs to.
	User shared.User `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BridgeID     respjson.Field
		LoginID      respjson.Field
		RemoveScopes respjson.Field
		Status       respjson.Field
		AccountIDs   respjson.Field
		StatusText   respjson.Field
		User         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoginSessionLogin) RawJSON() string { return r.JSON.raw }
func (r *LoginSessionLogin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced network capabilities for account lookup and group creation.
type ProvisioningCapabilities struct {
	GroupCreation map[string]GroupTypeCapabilities `json:"group_creation" api:"required"`
	// Identifier lookup capabilities for this bridge.
	ResolveIdentifier ResolveIdentifierCapabilities `json:"resolve_identifier" api:"required"`
	ImagePackImport   bool                          `json:"image_pack_import"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupCreation     respjson.Field
		ResolveIdentifier respjson.Field
		ImagePackImport   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProvisioningCapabilities) RawJSON() string { return r.JSON.raw }
func (r *ProvisioningCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier lookup capabilities for this bridge.
type ResolveIdentifierCapabilities struct {
	AnyPhone       bool `json:"any_phone" api:"required"`
	ContactList    bool `json:"contact_list" api:"required"`
	CreateDM       bool `json:"create_dm" api:"required"`
	LookupEmail    bool `json:"lookup_email" api:"required"`
	LookupPhone    bool `json:"lookup_phone" api:"required"`
	LookupUsername bool `json:"lookup_username" api:"required"`
	Search         bool `json:"search" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnyPhone       respjson.Field
		ContactList    respjson.Field
		CreateDM       respjson.Field
		LookupEmail    respjson.Field
		LookupPhone    respjson.Field
		LookupUsername respjson.Field
		Search         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResolveIdentifierCapabilities) RawJSON() string { return r.JSON.raw }
func (r *ResolveIdentifierCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Available bridge that can connect or reconnect chat accounts.
type BridgeGetResponse struct {
	// Bridge ID. Use with bridge endpoints.
	ID string `json:"id" api:"required"`
	// Connected accounts for this bridge. Uses the same Account schema as GET
	// /v1/accounts.
	Accounts []Account `json:"accounts" api:"required"`
	// Number of active accounts for this network on this device.
	ActiveAccountCount int64 `json:"activeAccountCount" api:"required"`
	// Human-friendly bridge name shown in Beeper.
	DisplayName string `json:"displayName" api:"required"`
	// Where accounts for this bridge run: on this device or in Beeper Cloud.
	//
	// Any of "cloud", "self-hosted", "local", "platform-sdk".
	Provider BridgeGetResponseProvider `json:"provider" api:"required"`
	// Whether this bridge can currently be used to connect new accounts.
	//
	// Any of "available", "connected", "limit_reached", "temporarily_unavailable",
	// "disabled".
	Status BridgeGetResponseStatus `json:"status" api:"required"`
	// Whether this bridge can have multiple active accounts for the same network.
	SupportsMultipleAccounts bool `json:"supportsMultipleAccounts" api:"required"`
	// Underlying bridge type, such as matrix, discordgo, slackgo, whatsapp, telegram,
	// or twitter.
	Type string `json:"type" api:"required"`
	// Network grouping used for account counts and limits.
	Network string `json:"network"`
	// Human-friendly status text matching Beeper account management language.
	StatusText string `json:"statusText"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Accounts                 respjson.Field
		ActiveAccountCount       respjson.Field
		DisplayName              respjson.Field
		Provider                 respjson.Field
		Status                   respjson.Field
		SupportsMultipleAccounts respjson.Field
		Type                     respjson.Field
		Network                  respjson.Field
		StatusText               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where accounts for this bridge run: on this device or in Beeper Cloud.
type BridgeGetResponseProvider string

const (
	BridgeGetResponseProviderCloud       BridgeGetResponseProvider = "cloud"
	BridgeGetResponseProviderSelfHosted  BridgeGetResponseProvider = "self-hosted"
	BridgeGetResponseProviderLocal       BridgeGetResponseProvider = "local"
	BridgeGetResponseProviderPlatformSDK BridgeGetResponseProvider = "platform-sdk"
)

// Whether this bridge can currently be used to connect new accounts.
type BridgeGetResponseStatus string

const (
	BridgeGetResponseStatusAvailable              BridgeGetResponseStatus = "available"
	BridgeGetResponseStatusConnected              BridgeGetResponseStatus = "connected"
	BridgeGetResponseStatusLimitReached           BridgeGetResponseStatus = "limit_reached"
	BridgeGetResponseStatusTemporarilyUnavailable BridgeGetResponseStatus = "temporarily_unavailable"
	BridgeGetResponseStatusDisabled               BridgeGetResponseStatus = "disabled"
)

// Available bridges and their connected accounts.
type BridgeListResponse struct {
	Items []Bridge `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeListResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

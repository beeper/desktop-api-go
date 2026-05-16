// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/param"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// Complete first-party Beeper app login
//
// AppLoginService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppLoginService] method instead.
type AppLoginService struct {
	Options []option.RequestOption
}

// NewAppLoginService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppLoginService(opts ...option.RequestOption) (r AppLoginService) {
	r = AppLoginService{}
	r.Options = opts
	return
}

// Send a sign-in code to the user email address.
func (r *AppLoginService) Email(ctx context.Context, body AppLoginEmailParams, opts ...option.RequestOption) (res *AppLoginEmailResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/login/email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a Beeper account after the user chooses a username and accepts the Terms
// of Use.
func (r *AppLoginService) Register(ctx context.Context, body AppLoginRegisterParams, opts ...option.RequestOption) (res *AppLoginRegisterResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/login/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Finish sign-in with the code sent to the user email address. If the user needs a
// new account, the response includes account creation copy and username
// suggestions.
func (r *AppLoginService) Response(ctx context.Context, body AppLoginResponseParams, opts ...option.RequestOption) (res *AppLoginResponseResponseUnion, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/login/response"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Start a first-party Beeper Desktop sign-in session.
func (r *AppLoginService) Start(ctx context.Context, opts ...option.RequestOption) (res *AppLoginStartResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/login/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type AppLoginEmailResponse = any

type AppLoginRegisterResponse struct {
	// Current onboarding state after sign-in.
	AppState AppLoginRegisterResponseAppState `json:"appState" api:"required"`
	// Desktop API credentials for the signed-in app session.
	DesktopAPI AppLoginRegisterResponseDesktopAPI `json:"desktopAPI" api:"required"`
	// Account credentials for first-party app setup.
	Matrix AppLoginRegisterResponseMatrix `json:"matrix" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState    respjson.Field
		DesktopAPI  respjson.Field
		Matrix      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponse) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after sign-in.
type AppLoginRegisterResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppLoginRegisterResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppLoginRegisterResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppLoginRegisterResponseAppStateVerification `json:"verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		E2ee         respjson.Field
		State        respjson.Field
		Matrix       respjson.Field
		Verification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppLoginRegisterResponseAppStateE2ee struct {
	// Whether this account can verify trusted devices.
	CrossSigning bool `json:"crossSigning" api:"required"`
	// Whether the first encrypted message sync is complete.
	FirstSyncDone bool `json:"firstSyncDone" api:"required"`
	// Whether the user confirmed that they saved their recovery key.
	HasBackedUpCode bool `json:"hasBackedUpCode" api:"required"`
	// Whether encrypted messaging setup has started.
	Initialized bool `json:"initialized" api:"required"`
	// Whether encrypted message backup is available.
	KeyBackup bool `json:"keyBackup" api:"required"`
	// Encrypted messaging keys available on this device.
	Secrets AppLoginRegisterResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
	// Whether secure key storage is available.
	SecretStorage bool `json:"secretStorage" api:"required"`
	// Whether this device is trusted for encrypted messages.
	Verified bool `json:"verified" api:"required"`
	// Unix timestamp for when the recovery key was created.
	RecoveryCodeGeneratedAt float64 `json:"recoveryCodeGeneratedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrossSigning            respjson.Field
		FirstSyncDone           respjson.Field
		HasBackedUpCode         respjson.Field
		Initialized             respjson.Field
		KeyBackup               respjson.Field
		Secrets                 respjson.Field
		SecretStorage           respjson.Field
		Verified                respjson.Field
		RecoveryCodeGeneratedAt respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppLoginRegisterResponseAppStateE2eeSecrets struct {
	// Whether the account identity key is available.
	MasterKey bool `json:"masterKey" api:"required"`
	// Whether the encrypted message backup key is available.
	MegolmBackupKey bool `json:"megolmBackupKey" api:"required"`
	// Whether a recovery key is available.
	RecoveryCode bool `json:"recoveryCode" api:"required"`
	// Whether the device trust key is available.
	SelfSigningKey bool `json:"selfSigningKey" api:"required"`
	// Whether the user trust key is available.
	UserSigningKey bool `json:"userSigningKey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MasterKey       respjson.Field
		MegolmBackupKey respjson.Field
		RecoveryCode    respjson.Field
		SelfSigningKey  respjson.Field
		UserSigningKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseAppStateE2eeSecrets) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppLoginRegisterResponseAppStateMatrix struct {
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper server URL for this account.
	Homeserver string `json:"homeserver" api:"required"`
	// Signed-in Beeper user ID.
	UserID string `json:"userID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceID    respjson.Field
		Homeserver  respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppLoginRegisterResponseAppStateVerification struct {
	// Verification actions that are valid for the current state.
	//
	// Any of "create", "qr.scan", "accept", "cancel", "qr.confirmScanned",
	// "sas.start", "sas.confirm".
	AvailableActions []string `json:"availableActions" api:"required"`
	// Current trusted-device verification state.
	//
	// Any of "idle", "requested", "ready", "sas_ready", "qr_scanned", "done",
	// "cancelled", "error".
	State string `json:"state" api:"required"`
	// Verification error details, if verification stopped.
	Error AppLoginRegisterResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppLoginRegisterResponseAppStateVerificationSas `json:"sas"`
	// Whether emoji comparison is available.
	SupportsSas bool `json:"supportsSAS"`
	// Whether QR code verification is available.
	SupportsScanQrCode bool `json:"supportsScanQRCode"`
	// Verification ID to pass in verification action paths.
	VerificationID string `json:"verificationID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableActions   respjson.Field
		State              respjson.Field
		Error              respjson.Field
		From               respjson.Field
		FromDevice         respjson.Field
		OtherDevice        respjson.Field
		QrData             respjson.Field
		Sas                respjson.Field
		SupportsSas        respjson.Field
		SupportsScanQrCode respjson.Field
		VerificationID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseAppStateVerification) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppLoginRegisterResponseAppStateVerificationError struct {
	// Verification error code.
	Code string `json:"code" api:"required"`
	// User-facing verification error message.
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseAppStateVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppLoginRegisterResponseAppStateVerificationSas struct {
	// Number sequence to compare on both devices.
	Decimals string `json:"decimals" api:"required"`
	// Emoji sequence to compare on both devices.
	Emojis string `json:"emojis" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Decimals    respjson.Field
		Emojis      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseAppStateVerificationSas) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Desktop API credentials for the signed-in app session.
type AppLoginRegisterResponseDesktopAPI struct {
	// Desktop API access token for this app session.
	AccessToken string `json:"accessToken" api:"required"`
	// Granted Desktop API scopes.
	//
	// Any of "read write".
	Scope string `json:"scope" api:"required"`
	// Access token type.
	//
	// Any of "Bearer".
	TokenType string `json:"tokenType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		Scope       respjson.Field
		TokenType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseDesktopAPI) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseDesktopAPI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account credentials for first-party app setup.
type AppLoginRegisterResponseMatrix struct {
	// Account access token. Returned once for first-party app setup.
	AccessToken string `json:"accessToken" api:"required"`
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper server URL for this account.
	Homeserver string `json:"homeserver" api:"required"`
	// Signed-in Beeper user ID.
	UserID string `json:"userID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		DeviceID    respjson.Field
		Homeserver  respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AppLoginResponseResponseUnion contains all possible properties and values from
// [AppLoginResponseResponseObject], [AppLoginResponseResponseObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AppLoginResponseResponseUnion struct {
	// This field is from variant [AppLoginResponseResponseObject].
	AppState AppLoginResponseResponseObjectAppState `json:"appState"`
	// This field is from variant [AppLoginResponseResponseObject].
	DesktopAPI AppLoginResponseResponseObjectDesktopAPI `json:"desktopAPI"`
	// This field is from variant [AppLoginResponseResponseObject].
	Matrix AppLoginResponseResponseObjectMatrix `json:"matrix"`
	// This field is from variant [AppLoginResponseResponseObject2].
	Copy AppLoginResponseResponseObject2Copy `json:"copy"`
	// This field is from variant [AppLoginResponseResponseObject2].
	LeadToken string `json:"leadToken"`
	// This field is from variant [AppLoginResponseResponseObject2].
	RegistrationRequired bool `json:"registrationRequired"`
	// This field is from variant [AppLoginResponseResponseObject2].
	Request string `json:"request"`
	// This field is from variant [AppLoginResponseResponseObject2].
	UsernameSuggestions []string `json:"usernameSuggestions"`
	JSON                struct {
		AppState             respjson.Field
		DesktopAPI           respjson.Field
		Matrix               respjson.Field
		Copy                 respjson.Field
		LeadToken            respjson.Field
		RegistrationRequired respjson.Field
		Request              respjson.Field
		UsernameSuggestions  respjson.Field
		raw                  string
	} `json:"-"`
}

func (u AppLoginResponseResponseUnion) AsAppLoginResponseResponseObject() (v AppLoginResponseResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppLoginResponseResponseUnion) AsAppLoginResponseResponseObject2() (v AppLoginResponseResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AppLoginResponseResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AppLoginResponseResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginResponseResponseObject struct {
	// Current onboarding state after sign-in.
	AppState AppLoginResponseResponseObjectAppState `json:"appState" api:"required"`
	// Desktop API credentials for the signed-in app session.
	DesktopAPI AppLoginResponseResponseObjectDesktopAPI `json:"desktopAPI" api:"required"`
	// Account credentials for first-party app setup.
	Matrix AppLoginResponseResponseObjectMatrix `json:"matrix" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState    respjson.Field
		DesktopAPI  respjson.Field
		Matrix      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObject) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after sign-in.
type AppLoginResponseResponseObjectAppState struct {
	// Encrypted messaging setup status.
	E2ee AppLoginResponseResponseObjectAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppLoginResponseResponseObjectAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppLoginResponseResponseObjectAppStateVerification `json:"verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		E2ee         respjson.Field
		State        respjson.Field
		Matrix       respjson.Field
		Verification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObjectAppState) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObjectAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppLoginResponseResponseObjectAppStateE2ee struct {
	// Whether this account can verify trusted devices.
	CrossSigning bool `json:"crossSigning" api:"required"`
	// Whether the first encrypted message sync is complete.
	FirstSyncDone bool `json:"firstSyncDone" api:"required"`
	// Whether the user confirmed that they saved their recovery key.
	HasBackedUpCode bool `json:"hasBackedUpCode" api:"required"`
	// Whether encrypted messaging setup has started.
	Initialized bool `json:"initialized" api:"required"`
	// Whether encrypted message backup is available.
	KeyBackup bool `json:"keyBackup" api:"required"`
	// Encrypted messaging keys available on this device.
	Secrets AppLoginResponseResponseObjectAppStateE2eeSecrets `json:"secrets" api:"required"`
	// Whether secure key storage is available.
	SecretStorage bool `json:"secretStorage" api:"required"`
	// Whether this device is trusted for encrypted messages.
	Verified bool `json:"verified" api:"required"`
	// Unix timestamp for when the recovery key was created.
	RecoveryCodeGeneratedAt float64 `json:"recoveryCodeGeneratedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrossSigning            respjson.Field
		FirstSyncDone           respjson.Field
		HasBackedUpCode         respjson.Field
		Initialized             respjson.Field
		KeyBackup               respjson.Field
		Secrets                 respjson.Field
		SecretStorage           respjson.Field
		Verified                respjson.Field
		RecoveryCodeGeneratedAt respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObjectAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObjectAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppLoginResponseResponseObjectAppStateE2eeSecrets struct {
	// Whether the account identity key is available.
	MasterKey bool `json:"masterKey" api:"required"`
	// Whether the encrypted message backup key is available.
	MegolmBackupKey bool `json:"megolmBackupKey" api:"required"`
	// Whether a recovery key is available.
	RecoveryCode bool `json:"recoveryCode" api:"required"`
	// Whether the device trust key is available.
	SelfSigningKey bool `json:"selfSigningKey" api:"required"`
	// Whether the user trust key is available.
	UserSigningKey bool `json:"userSigningKey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MasterKey       respjson.Field
		MegolmBackupKey respjson.Field
		RecoveryCode    respjson.Field
		SelfSigningKey  respjson.Field
		UserSigningKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObjectAppStateE2eeSecrets) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObjectAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppLoginResponseResponseObjectAppStateMatrix struct {
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper server URL for this account.
	Homeserver string `json:"homeserver" api:"required"`
	// Signed-in Beeper user ID.
	UserID string `json:"userID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceID    respjson.Field
		Homeserver  respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObjectAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObjectAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppLoginResponseResponseObjectAppStateVerification struct {
	// Verification actions that are valid for the current state.
	//
	// Any of "create", "qr.scan", "accept", "cancel", "qr.confirmScanned",
	// "sas.start", "sas.confirm".
	AvailableActions []string `json:"availableActions" api:"required"`
	// Current trusted-device verification state.
	//
	// Any of "idle", "requested", "ready", "sas_ready", "qr_scanned", "done",
	// "cancelled", "error".
	State string `json:"state" api:"required"`
	// Verification error details, if verification stopped.
	Error AppLoginResponseResponseObjectAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppLoginResponseResponseObjectAppStateVerificationSas `json:"sas"`
	// Whether emoji comparison is available.
	SupportsSas bool `json:"supportsSAS"`
	// Whether QR code verification is available.
	SupportsScanQrCode bool `json:"supportsScanQRCode"`
	// Verification ID to pass in verification action paths.
	VerificationID string `json:"verificationID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableActions   respjson.Field
		State              respjson.Field
		Error              respjson.Field
		From               respjson.Field
		FromDevice         respjson.Field
		OtherDevice        respjson.Field
		QrData             respjson.Field
		Sas                respjson.Field
		SupportsSas        respjson.Field
		SupportsScanQrCode respjson.Field
		VerificationID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObjectAppStateVerification) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObjectAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppLoginResponseResponseObjectAppStateVerificationError struct {
	// Verification error code.
	Code string `json:"code" api:"required"`
	// User-facing verification error message.
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObjectAppStateVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObjectAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppLoginResponseResponseObjectAppStateVerificationSas struct {
	// Number sequence to compare on both devices.
	Decimals string `json:"decimals" api:"required"`
	// Emoji sequence to compare on both devices.
	Emojis string `json:"emojis" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Decimals    respjson.Field
		Emojis      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObjectAppStateVerificationSas) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObjectAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Desktop API credentials for the signed-in app session.
type AppLoginResponseResponseObjectDesktopAPI struct {
	// Desktop API access token for this app session.
	AccessToken string `json:"accessToken" api:"required"`
	// Granted Desktop API scopes.
	//
	// Any of "read write".
	Scope string `json:"scope" api:"required"`
	// Access token type.
	//
	// Any of "Bearer".
	TokenType string `json:"tokenType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		Scope       respjson.Field
		TokenType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObjectDesktopAPI) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObjectDesktopAPI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account credentials for first-party app setup.
type AppLoginResponseResponseObjectMatrix struct {
	// Account access token. Returned once for first-party app setup.
	AccessToken string `json:"accessToken" api:"required"`
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper server URL for this account.
	Homeserver string `json:"homeserver" api:"required"`
	// Signed-in Beeper user ID.
	UserID string `json:"userID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		DeviceID    respjson.Field
		Homeserver  respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObjectMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObjectMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginResponseResponseObject2 struct {
	// Copy to display during account creation.
	Copy AppLoginResponseResponseObject2Copy `json:"copy" api:"required"`
	// Registration token returned by Beeper.
	LeadToken string `json:"leadToken" api:"required"`
	// Indicates that the user needs to create a Beeper account.
	//
	// Any of true.
	RegistrationRequired bool `json:"registrationRequired" api:"required"`
	// Login request ID to use when creating the account.
	Request string `json:"request" api:"required"`
	// Suggested usernames for the new account.
	UsernameSuggestions []string `json:"usernameSuggestions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Copy                 respjson.Field
		LeadToken            respjson.Field
		RegistrationRequired respjson.Field
		Request              respjson.Field
		UsernameSuggestions  respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Copy to display during account creation.
type AppLoginResponseResponseObject2Copy struct {
	// Submit button label.
	//
	// Any of "Continue".
	Submit string `json:"submit" api:"required"`
	// Terms and privacy notice to show before account creation.
	//
	// Any of "By continuing, you agree to the Terms of Use and acknowledge the Privacy
	// Policy.".
	Terms string `json:"terms" api:"required"`
	// Title for the username step.
	//
	// Any of "Choose your username".
	Title string `json:"title" api:"required"`
	// Placeholder for the username field.
	//
	// Any of "Username".
	UsernamePlaceholder string `json:"usernamePlaceholder" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Submit              respjson.Field
		Terms               respjson.Field
		Title               respjson.Field
		UsernamePlaceholder respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseObject2Copy) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseObject2Copy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginStartResponse struct {
	// Login request ID to use in the next sign-in step.
	Request string `json:"request" api:"required"`
	// Available sign-in methods for this request.
	Type []string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginStartResponse) RawJSON() string { return r.JSON.raw }
func (r *AppLoginStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginEmailParams struct {
	// Email address to send the sign-in code to.
	Email string `json:"email" api:"required" format:"email"`
	// Login request ID returned by the start step.
	Request string `json:"request" api:"required"`
	paramObj
}

func (r AppLoginEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow AppLoginEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppLoginEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginRegisterParams struct {
	// Confirms that the user accepted the Terms of Use and acknowledged the Privacy
	// Policy.
	//
	// Any of true.
	AcceptTerms bool `json:"acceptTerms,omitzero" api:"required"`
	// Registration token returned by Beeper.
	LeadToken string `json:"leadToken" api:"required"`
	// Login request ID returned by the start step.
	Request string `json:"request" api:"required"`
	// Username selected by the user.
	Username string `json:"username" api:"required"`
	paramObj
}

func (r AppLoginRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow AppLoginRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppLoginRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginResponseParams struct {
	// Login request ID returned by the start step.
	Request string `json:"request" api:"required"`
	// Sign-in code from the user email.
	Response string `json:"response" api:"required"`
	paramObj
}

func (r AppLoginResponseParams) MarshalJSON() (data []byte, err error) {
	type shadow AppLoginResponseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppLoginResponseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

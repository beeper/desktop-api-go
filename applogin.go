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
	"github.com/beeper/desktop-api-go/v5/shared/constant"
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
	Options      []option.RequestOption
	Verification AppLoginVerificationService
}

// NewAppLoginService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppLoginService(opts ...option.RequestOption) (r AppLoginService) {
	r = AppLoginService{}
	r.Options = opts
	r.Verification = NewAppLoginVerificationService(opts...)
	return
}

// Send a sign-in code to the user email address for app setup.
func (r *AppLoginService) Email(ctx context.Context, body AppLoginEmailParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/app/setup/email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Create a Beeper account after the user chooses a username and accepts the Terms
// of Use.
func (r *AppLoginService) Register(ctx context.Context, body AppLoginRegisterParams, opts ...option.RequestOption) (res *AppLoginRegisterResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/setup/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Finish setup sign-in with the code sent to the user email address. If the user
// needs a new account, the response includes account creation copy and username
// suggestions.
func (r *AppLoginService) Response(ctx context.Context, body AppLoginResponseParams, opts ...option.RequestOption) (res *AppLoginResponseResponseUnion, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/setup/response"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Start setting up Beeper Desktop or Beeper Server. The flow supports existing
// Beeper accounts and new account creation.
func (r *AppLoginService) Start(ctx context.Context, opts ...option.RequestOption) (res *AppLoginStartResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/setup/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type AppLoginRegisterResponse struct {
	// Account credentials for first-party app setup.
	Matrix AppLoginRegisterResponseMatrix `json:"matrix" api:"required"`
	// Current app sign-in and encrypted messaging setup state after sign-in.
	Session AppLoginRegisterResponseSession `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Matrix      respjson.Field
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponse) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account credentials for first-party app setup.
type AppLoginRegisterResponseMatrix struct {
	// Beeper account access token. Returned once for first-party app setup.
	AccessToken string `json:"accessToken" api:"required"`
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper homeserver URL for this account.
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

// Current app sign-in and encrypted messaging setup state after sign-in.
type AppLoginRegisterResponseSession struct {
	// Encrypted messaging setup status.
	E2EE AppLoginRegisterResponseSessionE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppLoginRegisterResponseSessionMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppLoginRegisterResponseSessionVerification `json:"verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		E2EE         respjson.Field
		State        respjson.Field
		Matrix       respjson.Field
		Verification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseSession) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppLoginRegisterResponseSessionE2EE struct {
	// Whether this account can verify trusted devices.
	CrossSigning bool `json:"crossSigning" api:"required"`
	// Whether the first encrypted message sync is complete.
	FirstSyncDone bool `json:"firstSyncDone" api:"required"`
	// Whether the user confirmed that they saved their recovery key.
	HasBackedUpRecoveryKey bool `json:"hasBackedUpRecoveryKey" api:"required"`
	// Whether encrypted messaging setup has started.
	Initialized bool `json:"initialized" api:"required"`
	// Whether encrypted message backup is available.
	KeyBackup bool `json:"keyBackup" api:"required"`
	// Encrypted messaging keys available on this device.
	Secrets AppLoginRegisterResponseSessionE2EESecrets `json:"secrets" api:"required"`
	// Whether secure key storage is available.
	SecretStorage bool `json:"secretStorage" api:"required"`
	// Whether this device is trusted for encrypted messages.
	Verified bool `json:"verified" api:"required"`
	// Unix timestamp for when the recovery key was created.
	RecoveryKeyGeneratedAt float64 `json:"recoveryKeyGeneratedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrossSigning           respjson.Field
		FirstSyncDone          respjson.Field
		HasBackedUpRecoveryKey respjson.Field
		Initialized            respjson.Field
		KeyBackup              respjson.Field
		Secrets                respjson.Field
		SecretStorage          respjson.Field
		Verified               respjson.Field
		RecoveryKeyGeneratedAt respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseSessionE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseSessionE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppLoginRegisterResponseSessionE2EESecrets struct {
	// Whether the account identity key is available.
	MasterKey bool `json:"masterKey" api:"required"`
	// Whether the encrypted message backup key is available.
	MegolmBackupKey bool `json:"megolmBackupKey" api:"required"`
	// Whether a recovery key is available.
	RecoveryKey bool `json:"recoveryKey" api:"required"`
	// Whether the device trust key is available.
	SelfSigningKey bool `json:"selfSigningKey" api:"required"`
	// Whether the user trust key is available.
	UserSigningKey bool `json:"userSigningKey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MasterKey       respjson.Field
		MegolmBackupKey respjson.Field
		RecoveryKey     respjson.Field
		SelfSigningKey  respjson.Field
		UserSigningKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseSessionE2EESecrets) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseSessionE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppLoginRegisterResponseSessionMatrix struct {
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper homeserver URL for this account.
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
func (r AppLoginRegisterResponseSessionMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseSessionMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppLoginRegisterResponseSessionVerification struct {
	// Verification ID to pass in verification action paths.
	ID string `json:"id" api:"required"`
	// Verification actions that are valid for the current state.
	//
	// Any of "accept", "cancel", "qr.confirmScanned", "sas.start", "sas.confirm".
	AvailableActions []string `json:"availableActions" api:"required"`
	// Whether this device started or received the verification.
	//
	// Any of "incoming", "outgoing".
	Direction string `json:"direction" api:"required"`
	// Verification methods supported for this transaction.
	//
	// Any of "qr", "sas".
	Methods []string `json:"methods" api:"required"`
	// Why this verification exists.
	//
	// Any of "login", "device".
	Purpose string `json:"purpose" api:"required"`
	// Current trusted-device verification state.
	//
	// Any of "requested", "ready", "sas_ready", "qr_scanned", "done", "cancelled",
	// "error".
	State string `json:"state" api:"required"`
	// Verification error details, if verification stopped.
	Error AppLoginRegisterResponseSessionVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppLoginRegisterResponseSessionVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	Qr AppLoginRegisterResponseSessionVerificationQr `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppLoginRegisterResponseSessionVerificationSAS `json:"sas"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AvailableActions respjson.Field
		Direction        respjson.Field
		Methods          respjson.Field
		Purpose          respjson.Field
		State            respjson.Field
		Error            respjson.Field
		OtherDevice      respjson.Field
		OtherUserID      respjson.Field
		Qr               respjson.Field
		SAS              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseSessionVerification) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseSessionVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppLoginRegisterResponseSessionVerificationError struct {
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
func (r AppLoginRegisterResponseSessionVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseSessionVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppLoginRegisterResponseSessionVerificationOtherDevice struct {
	// Other device ID.
	ID string `json:"id" api:"required"`
	// Other device display name, if known.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseSessionVerificationOtherDevice) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseSessionVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppLoginRegisterResponseSessionVerificationQr struct {
	// QR code payload to display for verification.
	Data string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseSessionVerificationQr) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseSessionVerificationQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppLoginRegisterResponseSessionVerificationSAS struct {
	// Emoji sequence to compare on both devices.
	Emojis string `json:"emojis" api:"required"`
	// Number sequence to compare on both devices.
	Decimals string `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Emojis      respjson.Field
		Decimals    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginRegisterResponseSessionVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppLoginRegisterResponseSessionVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AppLoginResponseResponseUnion contains all possible properties and values from
// [AppLoginResponseResponseSuccess],
// [AppLoginResponseResponseRegistrationRequired].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AppLoginResponseResponseUnion struct {
	// This field is from variant [AppLoginResponseResponseSuccess].
	Matrix AppLoginResponseResponseSuccessMatrix `json:"matrix"`
	// This field is from variant [AppLoginResponseResponseSuccess].
	Session AppLoginResponseResponseSuccessSession `json:"session"`
	// This field is from variant [AppLoginResponseResponseRegistrationRequired].
	Copy AppLoginResponseResponseRegistrationRequiredCopy `json:"copy"`
	// This field is from variant [AppLoginResponseResponseRegistrationRequired].
	LeadToken string `json:"leadToken"`
	// This field is from variant [AppLoginResponseResponseRegistrationRequired].
	RegistrationRequired bool `json:"registrationRequired"`
	// This field is from variant [AppLoginResponseResponseRegistrationRequired].
	SetupRequestID string `json:"setupRequestID"`
	// This field is from variant [AppLoginResponseResponseRegistrationRequired].
	UsernameSuggestions []string `json:"usernameSuggestions"`
	JSON                struct {
		Matrix               respjson.Field
		Session              respjson.Field
		Copy                 respjson.Field
		LeadToken            respjson.Field
		RegistrationRequired respjson.Field
		SetupRequestID       respjson.Field
		UsernameSuggestions  respjson.Field
		raw                  string
	} `json:"-"`
}

func (u AppLoginResponseResponseUnion) AsSuccess() (v AppLoginResponseResponseSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppLoginResponseResponseUnion) AsRegistrationRequired() (v AppLoginResponseResponseRegistrationRequired) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AppLoginResponseResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AppLoginResponseResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginResponseResponseSuccess struct {
	// Account credentials for first-party app setup.
	Matrix AppLoginResponseResponseSuccessMatrix `json:"matrix" api:"required"`
	// Current app sign-in and encrypted messaging setup state after sign-in.
	Session AppLoginResponseResponseSuccessSession `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Matrix      respjson.Field
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseSuccess) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account credentials for first-party app setup.
type AppLoginResponseResponseSuccessMatrix struct {
	// Beeper account access token. Returned once for first-party app setup.
	AccessToken string `json:"accessToken" api:"required"`
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper homeserver URL for this account.
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
func (r AppLoginResponseResponseSuccessMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccessMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current app sign-in and encrypted messaging setup state after sign-in.
type AppLoginResponseResponseSuccessSession struct {
	// Encrypted messaging setup status.
	E2EE AppLoginResponseResponseSuccessSessionE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppLoginResponseResponseSuccessSessionMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppLoginResponseResponseSuccessSessionVerification `json:"verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		E2EE         respjson.Field
		State        respjson.Field
		Matrix       respjson.Field
		Verification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseSuccessSession) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccessSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppLoginResponseResponseSuccessSessionE2EE struct {
	// Whether this account can verify trusted devices.
	CrossSigning bool `json:"crossSigning" api:"required"`
	// Whether the first encrypted message sync is complete.
	FirstSyncDone bool `json:"firstSyncDone" api:"required"`
	// Whether the user confirmed that they saved their recovery key.
	HasBackedUpRecoveryKey bool `json:"hasBackedUpRecoveryKey" api:"required"`
	// Whether encrypted messaging setup has started.
	Initialized bool `json:"initialized" api:"required"`
	// Whether encrypted message backup is available.
	KeyBackup bool `json:"keyBackup" api:"required"`
	// Encrypted messaging keys available on this device.
	Secrets AppLoginResponseResponseSuccessSessionE2EESecrets `json:"secrets" api:"required"`
	// Whether secure key storage is available.
	SecretStorage bool `json:"secretStorage" api:"required"`
	// Whether this device is trusted for encrypted messages.
	Verified bool `json:"verified" api:"required"`
	// Unix timestamp for when the recovery key was created.
	RecoveryKeyGeneratedAt float64 `json:"recoveryKeyGeneratedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrossSigning           respjson.Field
		FirstSyncDone          respjson.Field
		HasBackedUpRecoveryKey respjson.Field
		Initialized            respjson.Field
		KeyBackup              respjson.Field
		Secrets                respjson.Field
		SecretStorage          respjson.Field
		Verified               respjson.Field
		RecoveryKeyGeneratedAt respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseSuccessSessionE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccessSessionE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppLoginResponseResponseSuccessSessionE2EESecrets struct {
	// Whether the account identity key is available.
	MasterKey bool `json:"masterKey" api:"required"`
	// Whether the encrypted message backup key is available.
	MegolmBackupKey bool `json:"megolmBackupKey" api:"required"`
	// Whether a recovery key is available.
	RecoveryKey bool `json:"recoveryKey" api:"required"`
	// Whether the device trust key is available.
	SelfSigningKey bool `json:"selfSigningKey" api:"required"`
	// Whether the user trust key is available.
	UserSigningKey bool `json:"userSigningKey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MasterKey       respjson.Field
		MegolmBackupKey respjson.Field
		RecoveryKey     respjson.Field
		SelfSigningKey  respjson.Field
		UserSigningKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseSuccessSessionE2EESecrets) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccessSessionE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppLoginResponseResponseSuccessSessionMatrix struct {
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper homeserver URL for this account.
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
func (r AppLoginResponseResponseSuccessSessionMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccessSessionMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppLoginResponseResponseSuccessSessionVerification struct {
	// Verification ID to pass in verification action paths.
	ID string `json:"id" api:"required"`
	// Verification actions that are valid for the current state.
	//
	// Any of "accept", "cancel", "qr.confirmScanned", "sas.start", "sas.confirm".
	AvailableActions []string `json:"availableActions" api:"required"`
	// Whether this device started or received the verification.
	//
	// Any of "incoming", "outgoing".
	Direction string `json:"direction" api:"required"`
	// Verification methods supported for this transaction.
	//
	// Any of "qr", "sas".
	Methods []string `json:"methods" api:"required"`
	// Why this verification exists.
	//
	// Any of "login", "device".
	Purpose string `json:"purpose" api:"required"`
	// Current trusted-device verification state.
	//
	// Any of "requested", "ready", "sas_ready", "qr_scanned", "done", "cancelled",
	// "error".
	State string `json:"state" api:"required"`
	// Verification error details, if verification stopped.
	Error AppLoginResponseResponseSuccessSessionVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppLoginResponseResponseSuccessSessionVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	Qr AppLoginResponseResponseSuccessSessionVerificationQr `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppLoginResponseResponseSuccessSessionVerificationSAS `json:"sas"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AvailableActions respjson.Field
		Direction        respjson.Field
		Methods          respjson.Field
		Purpose          respjson.Field
		State            respjson.Field
		Error            respjson.Field
		OtherDevice      respjson.Field
		OtherUserID      respjson.Field
		Qr               respjson.Field
		SAS              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseSuccessSessionVerification) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccessSessionVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppLoginResponseResponseSuccessSessionVerificationError struct {
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
func (r AppLoginResponseResponseSuccessSessionVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccessSessionVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppLoginResponseResponseSuccessSessionVerificationOtherDevice struct {
	// Other device ID.
	ID string `json:"id" api:"required"`
	// Other device display name, if known.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseSuccessSessionVerificationOtherDevice) RawJSON() string {
	return r.JSON.raw
}
func (r *AppLoginResponseResponseSuccessSessionVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppLoginResponseResponseSuccessSessionVerificationQr struct {
	// QR code payload to display for verification.
	Data string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseSuccessSessionVerificationQr) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccessSessionVerificationQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppLoginResponseResponseSuccessSessionVerificationSAS struct {
	// Emoji sequence to compare on both devices.
	Emojis string `json:"emojis" api:"required"`
	// Number sequence to compare on both devices.
	Decimals string `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Emojis      respjson.Field
		Decimals    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseSuccessSessionVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseSuccessSessionVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginResponseResponseRegistrationRequired struct {
	// Copy to display during account creation.
	Copy AppLoginResponseResponseRegistrationRequiredCopy `json:"copy" api:"required"`
	// Registration token returned by Beeper.
	LeadToken string `json:"leadToken" api:"required"`
	// Indicates that the user needs to create a Beeper account.
	RegistrationRequired bool `json:"registrationRequired" api:"required"`
	// Setup request ID to use when creating the account.
	SetupRequestID string `json:"setupRequestID" api:"required"`
	// Suggested usernames for the new account.
	UsernameSuggestions []string `json:"usernameSuggestions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Copy                 respjson.Field
		LeadToken            respjson.Field
		RegistrationRequired respjson.Field
		SetupRequestID       respjson.Field
		UsernameSuggestions  respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginResponseResponseRegistrationRequired) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseRegistrationRequired) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Copy to display during account creation.
type AppLoginResponseResponseRegistrationRequiredCopy struct {
	// Submit button label.
	Submit constant.Continue `json:"submit" default:"Continue"`
	// Terms and privacy notice to show before account creation.
	Terms constant.ByContinuingYouAgreeToTheTermsOfUseAndAcknowledgeThePrivacyPolicy `json:"terms" default:"By continuing, you agree to the Terms of Use and acknowledge the Privacy Policy."`
	// Title for the username step.
	Title constant.ChooseYourUsername `json:"title" default:"Choose your username"`
	// Placeholder for the username field.
	UsernamePlaceholder constant.Username `json:"usernamePlaceholder" default:"Username"`
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
func (r AppLoginResponseResponseRegistrationRequiredCopy) RawJSON() string { return r.JSON.raw }
func (r *AppLoginResponseResponseRegistrationRequiredCopy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginStartResponse struct {
	// Setup request ID to use in the next sign-in step.
	SetupRequestID string `json:"setupRequestID" api:"required"`
	// Available sign-in methods for this setup request.
	SignInMethods []string `json:"signInMethods" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SetupRequestID respjson.Field
		SignInMethods  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
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
	// Setup request ID returned by the start step.
	SetupRequestID string `json:"setupRequestID" api:"required"`
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
	// Registration token returned by Beeper.
	LeadToken string `json:"leadToken" api:"required"`
	// Setup request ID returned by the start step.
	SetupRequestID string `json:"setupRequestID" api:"required"`
	// Username selected by the user.
	Username string `json:"username" api:"required"`
	// Confirms that the user agreed to our
	// [terms of use](https://www.beeper.com/terms-onboarding) and has read our
	// [privacy policy](https://www.beeper.com/privacy).
	//
	// This field can be elided, and will marshal its zero value as true.
	AcceptTerms bool `json:"acceptTerms,omitzero" api:"required"`
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
	// Sign-in code from the user email.
	Response string `json:"response" api:"required"`
	// Setup request ID returned by the start step.
	SetupRequestID string `json:"setupRequestID" api:"required"`
	paramObj
}

func (r AppLoginResponseParams) MarshalJSON() (data []byte, err error) {
	type shadow AppLoginResponseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppLoginResponseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

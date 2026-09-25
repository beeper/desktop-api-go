// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v6/internal/apijson"
	"github.com/beeper/desktop-api-go/v6/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v6/option"
	"github.com/beeper/desktop-api-go/v6/packages/param"
	"github.com/beeper/desktop-api-go/v6/packages/respjson"
	"github.com/beeper/desktop-api-go/v6/shared/constant"
)

// Complete first-party Beeper app setup
//
// AppSetupService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppSetupService] method instead.
type AppSetupService struct {
	Options []option.RequestOption
	// Manage recovery-key setup for encrypted messages
	RecoveryKey AppSetupRecoveryKeyService
	// Manage device verification transactions
	Verifications AppSetupVerificationService
}

// NewAppSetupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppSetupService(opts ...option.RequestOption) (r AppSetupService) {
	r = AppSetupService{}
	r.Options = opts
	r.RecoveryKey = NewAppSetupRecoveryKeyService(opts...)
	r.Verifications = NewAppSetupVerificationService(opts...)
	return
}

// Return the current Beeper Desktop or Beeper Server sign-in and encrypted
// messaging setup state. This endpoint is public before sign-in so apps can
// discover that sign-in is needed; after sign-in, pass a read token.
func (r *AppSetupService) Get(ctx context.Context, opts ...option.RequestOption) (res *AppSetupGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/setup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Send a sign-in code to the user email address for app setup.
func (r *AppSetupService) Email(ctx context.Context, body AppSetupEmailParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/app/setup/email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Create a Beeper account after the user chooses a username and accepts the Terms
// of Use.
func (r *AppSetupService) Register(ctx context.Context, body AppSetupRegisterParams, opts ...option.RequestOption) (res *AppSetupRegisterResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/setup/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Finish setup sign-in with the code sent to the user email address. If the user
// needs a new account, the response includes account creation copy and username
// suggestions.
func (r *AppSetupService) Response(ctx context.Context, body AppSetupResponseParams, opts ...option.RequestOption) (res *AppSetupResponseResponseUnion, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/setup/response"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Start setting up Beeper Desktop or Beeper Server. The flow supports existing
// Beeper accounts and new account creation.
func (r *AppSetupService) Start(ctx context.Context, opts ...option.RequestOption) (res *AppSetupStartResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "v1/app/setup/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type AppSetupGetResponse struct {
	// Encrypted messaging setup status.
	E2EE AppSetupGetResponseE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State AppSetupGetResponseState `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppSetupGetResponseMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppSetupGetResponseVerification `json:"verification"`
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
func (r AppSetupGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AppSetupGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppSetupGetResponseE2EE struct {
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
	Secrets AppSetupGetResponseE2EESecrets `json:"secrets" api:"required"`
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
func (r AppSetupGetResponseE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppSetupGetResponseE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppSetupGetResponseE2EESecrets struct {
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
func (r AppSetupGetResponseE2EESecrets) RawJSON() string { return r.JSON.raw }
func (r *AppSetupGetResponseE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
// Server.
type AppSetupGetResponseState string

const (
	AppSetupGetResponseStateNeedsLogin             AppSetupGetResponseState = "needs-login"
	AppSetupGetResponseStateInitializing           AppSetupGetResponseState = "initializing"
	AppSetupGetResponseStateNeedsCrossSigningSetup AppSetupGetResponseState = "needs-cross-signing-setup"
	AppSetupGetResponseStateNeedsVerification      AppSetupGetResponseState = "needs-verification"
	AppSetupGetResponseStateNeedsSecrets           AppSetupGetResponseState = "needs-secrets"
	AppSetupGetResponseStateNeedsFirstSync         AppSetupGetResponseState = "needs-first-sync"
	AppSetupGetResponseStateReady                  AppSetupGetResponseState = "ready"
)

// Signed-in account details. Omitted until sign-in is complete.
type AppSetupGetResponseMatrix struct {
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
func (r AppSetupGetResponseMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppSetupGetResponseMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppSetupGetResponseVerification struct {
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
	Error AppSetupGetResponseVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppSetupGetResponseVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	QR AppSetupGetResponseVerificationQR `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppSetupGetResponseVerificationSAS `json:"sas"`
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
		QR               respjson.Field
		SAS              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSetupGetResponseVerification) RawJSON() string { return r.JSON.raw }
func (r *AppSetupGetResponseVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppSetupGetResponseVerificationError struct {
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
func (r AppSetupGetResponseVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppSetupGetResponseVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppSetupGetResponseVerificationOtherDevice struct {
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
func (r AppSetupGetResponseVerificationOtherDevice) RawJSON() string { return r.JSON.raw }
func (r *AppSetupGetResponseVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppSetupGetResponseVerificationQR struct {
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
func (r AppSetupGetResponseVerificationQR) RawJSON() string { return r.JSON.raw }
func (r *AppSetupGetResponseVerificationQR) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppSetupGetResponseVerificationSAS struct {
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
func (r AppSetupGetResponseVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppSetupGetResponseVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppSetupRegisterResponse struct {
	// Account credentials for first-party app setup.
	Matrix AppSetupRegisterResponseMatrix `json:"matrix" api:"required"`
	// Current app sign-in and encrypted messaging setup state after sign-in.
	Session AppSetupRegisterResponseSession `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Matrix      respjson.Field
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSetupRegisterResponse) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account credentials for first-party app setup.
type AppSetupRegisterResponseMatrix struct {
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
func (r AppSetupRegisterResponseMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current app sign-in and encrypted messaging setup state after sign-in.
type AppSetupRegisterResponseSession struct {
	// Encrypted messaging setup status.
	E2EE AppSetupRegisterResponseSessionE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppSetupRegisterResponseSessionMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppSetupRegisterResponseSessionVerification `json:"verification"`
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
func (r AppSetupRegisterResponseSession) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppSetupRegisterResponseSessionE2EE struct {
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
	Secrets AppSetupRegisterResponseSessionE2EESecrets `json:"secrets" api:"required"`
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
func (r AppSetupRegisterResponseSessionE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseSessionE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppSetupRegisterResponseSessionE2EESecrets struct {
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
func (r AppSetupRegisterResponseSessionE2EESecrets) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseSessionE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppSetupRegisterResponseSessionMatrix struct {
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
func (r AppSetupRegisterResponseSessionMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseSessionMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppSetupRegisterResponseSessionVerification struct {
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
	Error AppSetupRegisterResponseSessionVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppSetupRegisterResponseSessionVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	QR AppSetupRegisterResponseSessionVerificationQR `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppSetupRegisterResponseSessionVerificationSAS `json:"sas"`
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
		QR               respjson.Field
		SAS              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSetupRegisterResponseSessionVerification) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseSessionVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppSetupRegisterResponseSessionVerificationError struct {
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
func (r AppSetupRegisterResponseSessionVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseSessionVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppSetupRegisterResponseSessionVerificationOtherDevice struct {
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
func (r AppSetupRegisterResponseSessionVerificationOtherDevice) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseSessionVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppSetupRegisterResponseSessionVerificationQR struct {
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
func (r AppSetupRegisterResponseSessionVerificationQR) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseSessionVerificationQR) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppSetupRegisterResponseSessionVerificationSAS struct {
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
func (r AppSetupRegisterResponseSessionVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRegisterResponseSessionVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AppSetupResponseResponseUnion contains all possible properties and values from
// [AppSetupResponseResponseSuccess],
// [AppSetupResponseResponseRegistrationRequired].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AppSetupResponseResponseUnion struct {
	// This field is from variant [AppSetupResponseResponseSuccess].
	Matrix AppSetupResponseResponseSuccessMatrix `json:"matrix"`
	// This field is from variant [AppSetupResponseResponseSuccess].
	Session AppSetupResponseResponseSuccessSession `json:"session"`
	// This field is from variant [AppSetupResponseResponseRegistrationRequired].
	Copy AppSetupResponseResponseRegistrationRequiredCopy `json:"copy"`
	// This field is from variant [AppSetupResponseResponseRegistrationRequired].
	LeadToken string `json:"leadToken"`
	// This field is from variant [AppSetupResponseResponseRegistrationRequired].
	RegistrationRequired bool `json:"registrationRequired"`
	// This field is from variant [AppSetupResponseResponseRegistrationRequired].
	SetupRequestID string `json:"setupRequestID"`
	// This field is from variant [AppSetupResponseResponseRegistrationRequired].
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

func (u AppSetupResponseResponseUnion) AsSuccess() (v AppSetupResponseResponseSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppSetupResponseResponseUnion) AsRegistrationRequired() (v AppSetupResponseResponseRegistrationRequired) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AppSetupResponseResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AppSetupResponseResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppSetupResponseResponseSuccess struct {
	// Account credentials for first-party app setup.
	Matrix AppSetupResponseResponseSuccessMatrix `json:"matrix" api:"required"`
	// Current app sign-in and encrypted messaging setup state after sign-in.
	Session AppSetupResponseResponseSuccessSession `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Matrix      respjson.Field
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSetupResponseResponseSuccess) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account credentials for first-party app setup.
type AppSetupResponseResponseSuccessMatrix struct {
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
func (r AppSetupResponseResponseSuccessMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccessMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current app sign-in and encrypted messaging setup state after sign-in.
type AppSetupResponseResponseSuccessSession struct {
	// Encrypted messaging setup status.
	E2EE AppSetupResponseResponseSuccessSessionE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppSetupResponseResponseSuccessSessionMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppSetupResponseResponseSuccessSessionVerification `json:"verification"`
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
func (r AppSetupResponseResponseSuccessSession) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccessSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppSetupResponseResponseSuccessSessionE2EE struct {
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
	Secrets AppSetupResponseResponseSuccessSessionE2EESecrets `json:"secrets" api:"required"`
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
func (r AppSetupResponseResponseSuccessSessionE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccessSessionE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppSetupResponseResponseSuccessSessionE2EESecrets struct {
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
func (r AppSetupResponseResponseSuccessSessionE2EESecrets) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccessSessionE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppSetupResponseResponseSuccessSessionMatrix struct {
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
func (r AppSetupResponseResponseSuccessSessionMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccessSessionMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppSetupResponseResponseSuccessSessionVerification struct {
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
	Error AppSetupResponseResponseSuccessSessionVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppSetupResponseResponseSuccessSessionVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	QR AppSetupResponseResponseSuccessSessionVerificationQR `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppSetupResponseResponseSuccessSessionVerificationSAS `json:"sas"`
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
		QR               respjson.Field
		SAS              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSetupResponseResponseSuccessSessionVerification) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccessSessionVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppSetupResponseResponseSuccessSessionVerificationError struct {
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
func (r AppSetupResponseResponseSuccessSessionVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccessSessionVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppSetupResponseResponseSuccessSessionVerificationOtherDevice struct {
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
func (r AppSetupResponseResponseSuccessSessionVerificationOtherDevice) RawJSON() string {
	return r.JSON.raw
}
func (r *AppSetupResponseResponseSuccessSessionVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppSetupResponseResponseSuccessSessionVerificationQR struct {
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
func (r AppSetupResponseResponseSuccessSessionVerificationQR) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccessSessionVerificationQR) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppSetupResponseResponseSuccessSessionVerificationSAS struct {
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
func (r AppSetupResponseResponseSuccessSessionVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseSuccessSessionVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppSetupResponseResponseRegistrationRequired struct {
	// Copy to display during account creation.
	Copy AppSetupResponseResponseRegistrationRequiredCopy `json:"copy" api:"required"`
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
func (r AppSetupResponseResponseRegistrationRequired) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseRegistrationRequired) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Copy to display during account creation.
type AppSetupResponseResponseRegistrationRequiredCopy struct {
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
func (r AppSetupResponseResponseRegistrationRequiredCopy) RawJSON() string { return r.JSON.raw }
func (r *AppSetupResponseResponseRegistrationRequiredCopy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppSetupStartResponse struct {
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
func (r AppSetupStartResponse) RawJSON() string { return r.JSON.raw }
func (r *AppSetupStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppSetupEmailParams struct {
	// Email address to send the sign-in code to.
	Email string `json:"email" api:"required" format:"email"`
	// Setup request ID returned by the start step.
	SetupRequestID string `json:"setupRequestID" api:"required"`
	paramObj
}

func (r AppSetupEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow AppSetupEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppSetupEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppSetupRegisterParams struct {
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

func (r AppSetupRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow AppSetupRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppSetupRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppSetupResponseParams struct {
	// Sign-in code from the user email.
	Response string `json:"response" api:"required"`
	// Setup request ID returned by the start step.
	SetupRequestID string `json:"setupRequestID" api:"required"`
	paramObj
}

func (r AppSetupResponseParams) MarshalJSON() (data []byte, err error) {
	type shadow AppSetupResponseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppSetupResponseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

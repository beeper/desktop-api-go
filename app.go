// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// Manage Beeper app login and encrypted messaging setup
//
// AppService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
	// Complete first-party Beeper app login
	Login AppLoginService
	// Manage encrypted messaging setup
	E2ee AppE2eeService
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r AppService) {
	r = AppService{}
	r.Options = opts
	r.Login = NewAppLoginService(opts...)
	r.E2ee = NewAppE2eeService(opts...)
	return
}

// Return the current Beeper Desktop sign-in and encrypted messaging setup state.
// This endpoint is public before sign-in so apps can discover that login is
// needed; after sign-in, pass a read token.
func (r *AppService) Status(ctx context.Context, opts ...option.RequestOption) (res *AppStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AppStatusResponse struct {
	// Encrypted messaging setup status.
	E2ee AppStatusResponseE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State AppStatusResponseState `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppStatusResponseMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppStatusResponseVerification `json:"verification"`
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
func (r AppStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *AppStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppStatusResponseE2ee struct {
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
	Secrets AppStatusResponseE2eeSecrets `json:"secrets" api:"required"`
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
func (r AppStatusResponseE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppStatusResponseE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppStatusResponseE2eeSecrets struct {
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
func (r AppStatusResponseE2eeSecrets) RawJSON() string { return r.JSON.raw }
func (r *AppStatusResponseE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state for Beeper Desktop.
type AppStatusResponseState string

const (
	AppStatusResponseStateNeedsLogin             AppStatusResponseState = "needs-login"
	AppStatusResponseStateInitializing           AppStatusResponseState = "initializing"
	AppStatusResponseStateNeedsCrossSigningSetup AppStatusResponseState = "needs-cross-signing-setup"
	AppStatusResponseStateNeedsVerification      AppStatusResponseState = "needs-verification"
	AppStatusResponseStateNeedsSecrets           AppStatusResponseState = "needs-secrets"
	AppStatusResponseStateNeedsFirstSync         AppStatusResponseState = "needs-first-sync"
	AppStatusResponseStateReady                  AppStatusResponseState = "ready"
)

// Signed-in account details. Omitted until sign-in is complete.
type AppStatusResponseMatrix struct {
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
func (r AppStatusResponseMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppStatusResponseMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppStatusResponseVerification struct {
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
	Error AppStatusResponseVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppStatusResponseVerificationSas `json:"sas"`
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
func (r AppStatusResponseVerification) RawJSON() string { return r.JSON.raw }
func (r *AppStatusResponseVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppStatusResponseVerificationError struct {
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
func (r AppStatusResponseVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppStatusResponseVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppStatusResponseVerificationSas struct {
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
func (r AppStatusResponseVerificationSas) RawJSON() string { return r.JSON.raw }
func (r *AppStatusResponseVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// First-party sign-in and encrypted messaging setup for Beeper Desktop and Beeper
// Server.
//
// AppVerificationQrService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppVerificationQrService] method instead.
type AppVerificationQrService struct {
	Options []option.RequestOption
}

// NewAppVerificationQrService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAppVerificationQrService(opts ...option.RequestOption) (r AppVerificationQrService) {
	r = AppVerificationQrService{}
	r.Options = opts
	return
}

// Confirm that another device scanned this device QR code.
func (r *AppVerificationQrService) ConfirmScanned(ctx context.Context, verificationID string, opts ...option.RequestOption) (res *AppVerificationQrConfirmScannedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if verificationID == "" {
		err = errors.New("missing required verificationID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/app/setup/verifications/%s/qr/confirm-scanned", verificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Submit the QR code scanned from another signed-in device.
func (r *AppVerificationQrService) Scan(ctx context.Context, body AppVerificationQrScanParams, opts ...option.RequestOption) (res *AppVerificationQrScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/setup/verifications/qr/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AppVerificationQrConfirmScannedResponse struct {
	// Current app sign-in and encrypted messaging setup state.
	Session AppVerificationQrConfirmScannedResponseSession `json:"session" api:"required"`
	// Trusted device verification progress.
	Verification AppVerificationQrConfirmScannedResponseVerification `json:"verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Session      respjson.Field
		Verification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppVerificationQrConfirmScannedResponse) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrConfirmScannedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current app sign-in and encrypted messaging setup state.
type AppVerificationQrConfirmScannedResponseSession struct {
	// Encrypted messaging setup status.
	E2EE AppVerificationQrConfirmScannedResponseSessionE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppVerificationQrConfirmScannedResponseSessionMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppVerificationQrConfirmScannedResponseSessionVerification `json:"verification"`
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
func (r AppVerificationQrConfirmScannedResponseSession) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrConfirmScannedResponseSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppVerificationQrConfirmScannedResponseSessionE2EE struct {
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
	Secrets AppVerificationQrConfirmScannedResponseSessionE2EESecrets `json:"secrets" api:"required"`
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
func (r AppVerificationQrConfirmScannedResponseSessionE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrConfirmScannedResponseSessionE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppVerificationQrConfirmScannedResponseSessionE2EESecrets struct {
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
func (r AppVerificationQrConfirmScannedResponseSessionE2EESecrets) RawJSON() string {
	return r.JSON.raw
}
func (r *AppVerificationQrConfirmScannedResponseSessionE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppVerificationQrConfirmScannedResponseSessionMatrix struct {
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
func (r AppVerificationQrConfirmScannedResponseSessionMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrConfirmScannedResponseSessionMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppVerificationQrConfirmScannedResponseSessionVerification struct {
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
	Error AppVerificationQrConfirmScannedResponseSessionVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppVerificationQrConfirmScannedResponseSessionVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	Qr AppVerificationQrConfirmScannedResponseSessionVerificationQr `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppVerificationQrConfirmScannedResponseSessionVerificationSAS `json:"sas"`
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
func (r AppVerificationQrConfirmScannedResponseSessionVerification) RawJSON() string {
	return r.JSON.raw
}
func (r *AppVerificationQrConfirmScannedResponseSessionVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppVerificationQrConfirmScannedResponseSessionVerificationError struct {
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
func (r AppVerificationQrConfirmScannedResponseSessionVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppVerificationQrConfirmScannedResponseSessionVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppVerificationQrConfirmScannedResponseSessionVerificationOtherDevice struct {
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
func (r AppVerificationQrConfirmScannedResponseSessionVerificationOtherDevice) RawJSON() string {
	return r.JSON.raw
}
func (r *AppVerificationQrConfirmScannedResponseSessionVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppVerificationQrConfirmScannedResponseSessionVerificationQr struct {
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
func (r AppVerificationQrConfirmScannedResponseSessionVerificationQr) RawJSON() string {
	return r.JSON.raw
}
func (r *AppVerificationQrConfirmScannedResponseSessionVerificationQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppVerificationQrConfirmScannedResponseSessionVerificationSAS struct {
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
func (r AppVerificationQrConfirmScannedResponseSessionVerificationSAS) RawJSON() string {
	return r.JSON.raw
}
func (r *AppVerificationQrConfirmScannedResponseSessionVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppVerificationQrConfirmScannedResponseVerification struct {
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
	Error AppVerificationQrConfirmScannedResponseVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppVerificationQrConfirmScannedResponseVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	Qr AppVerificationQrConfirmScannedResponseVerificationQr `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppVerificationQrConfirmScannedResponseVerificationSAS `json:"sas"`
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
func (r AppVerificationQrConfirmScannedResponseVerification) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrConfirmScannedResponseVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppVerificationQrConfirmScannedResponseVerificationError struct {
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
func (r AppVerificationQrConfirmScannedResponseVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrConfirmScannedResponseVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppVerificationQrConfirmScannedResponseVerificationOtherDevice struct {
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
func (r AppVerificationQrConfirmScannedResponseVerificationOtherDevice) RawJSON() string {
	return r.JSON.raw
}
func (r *AppVerificationQrConfirmScannedResponseVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppVerificationQrConfirmScannedResponseVerificationQr struct {
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
func (r AppVerificationQrConfirmScannedResponseVerificationQr) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrConfirmScannedResponseVerificationQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppVerificationQrConfirmScannedResponseVerificationSAS struct {
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
func (r AppVerificationQrConfirmScannedResponseVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrConfirmScannedResponseVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppVerificationQrScanResponse struct {
	// Current app sign-in and encrypted messaging setup state.
	Session AppVerificationQrScanResponseSession `json:"session" api:"required"`
	// Trusted device verification progress.
	Verification AppVerificationQrScanResponseVerification `json:"verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Session      respjson.Field
		Verification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppVerificationQrScanResponse) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current app sign-in and encrypted messaging setup state.
type AppVerificationQrScanResponseSession struct {
	// Encrypted messaging setup status.
	E2EE AppVerificationQrScanResponseSessionE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppVerificationQrScanResponseSessionMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppVerificationQrScanResponseSessionVerification `json:"verification"`
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
func (r AppVerificationQrScanResponseSession) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppVerificationQrScanResponseSessionE2EE struct {
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
	Secrets AppVerificationQrScanResponseSessionE2EESecrets `json:"secrets" api:"required"`
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
func (r AppVerificationQrScanResponseSessionE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseSessionE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppVerificationQrScanResponseSessionE2EESecrets struct {
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
func (r AppVerificationQrScanResponseSessionE2EESecrets) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseSessionE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppVerificationQrScanResponseSessionMatrix struct {
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
func (r AppVerificationQrScanResponseSessionMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseSessionMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppVerificationQrScanResponseSessionVerification struct {
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
	Error AppVerificationQrScanResponseSessionVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppVerificationQrScanResponseSessionVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	Qr AppVerificationQrScanResponseSessionVerificationQr `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppVerificationQrScanResponseSessionVerificationSAS `json:"sas"`
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
func (r AppVerificationQrScanResponseSessionVerification) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseSessionVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppVerificationQrScanResponseSessionVerificationError struct {
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
func (r AppVerificationQrScanResponseSessionVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseSessionVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppVerificationQrScanResponseSessionVerificationOtherDevice struct {
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
func (r AppVerificationQrScanResponseSessionVerificationOtherDevice) RawJSON() string {
	return r.JSON.raw
}
func (r *AppVerificationQrScanResponseSessionVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppVerificationQrScanResponseSessionVerificationQr struct {
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
func (r AppVerificationQrScanResponseSessionVerificationQr) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseSessionVerificationQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppVerificationQrScanResponseSessionVerificationSAS struct {
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
func (r AppVerificationQrScanResponseSessionVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseSessionVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppVerificationQrScanResponseVerification struct {
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
	Error AppVerificationQrScanResponseVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppVerificationQrScanResponseVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	Qr AppVerificationQrScanResponseVerificationQr `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppVerificationQrScanResponseVerificationSAS `json:"sas"`
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
func (r AppVerificationQrScanResponseVerification) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppVerificationQrScanResponseVerificationError struct {
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
func (r AppVerificationQrScanResponseVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppVerificationQrScanResponseVerificationOtherDevice struct {
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
func (r AppVerificationQrScanResponseVerificationOtherDevice) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppVerificationQrScanResponseVerificationQr struct {
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
func (r AppVerificationQrScanResponseVerificationQr) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseVerificationQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppVerificationQrScanResponseVerificationSAS struct {
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
func (r AppVerificationQrScanResponseVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppVerificationQrScanResponseVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppVerificationQrScanParams struct {
	// QR code payload scanned from the other device.
	Data string `json:"data" api:"required"`
	paramObj
}

func (r AppVerificationQrScanParams) MarshalJSON() (data []byte, err error) {
	type shadow AppVerificationQrScanParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppVerificationQrScanParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"github.com/beeper/desktop-api-go/v5/option"
)

// Manage encrypted messaging setup
//
// AppE2eeService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppE2eeService] method instead.
type AppE2eeService struct {
	Options []option.RequestOption
	// First-party sign-in and encrypted messaging setup for Beeper Desktop.
	RecoveryCode AppE2eeRecoveryCodeService
	// First-party sign-in and encrypted messaging setup for Beeper Desktop.
	Verification AppE2eeVerificationService
}

// NewAppE2eeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppE2eeService(opts ...option.RequestOption) (r AppE2eeService) {
	r = AppE2eeService{}
	r.Options = opts
	r.RecoveryCode = NewAppE2eeRecoveryCodeService(opts...)
	r.Verification = NewAppE2eeVerificationService(opts...)
	return
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"github.com/beeper/desktop-api-go/v5/option"
)

// AppVerificationSASService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppVerificationSASService] method instead.
type AppVerificationSASService struct {
	Options []option.RequestOption
}

// NewAppVerificationSASService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAppVerificationSASService(opts ...option.RequestOption) (r AppVerificationSASService) {
	r = AppVerificationSASService{}
	r.Options = opts
	return
}

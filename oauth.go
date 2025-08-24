// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo

import (
	"github.com/beeper/desktop-api-go/option"
)

// OAuth2 authentication and token management
//
// OAuthService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthService] method instead.
type OAuthService struct {
	Options []option.RequestOption
}

// NewOAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOAuthService(opts ...option.RequestOption) (r OAuthService) {
	r = OAuthService{}
	r.Options = opts
	return
}

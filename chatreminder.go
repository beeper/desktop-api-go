// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"github.com/beeper/desktop-api-go/option"
)

// Reminders operations
//
// ChatReminderService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatReminderService] method instead.
type ChatReminderService struct {
	Options []option.RequestOption
}

// NewChatReminderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatReminderService(opts ...option.RequestOption) (r ChatReminderService) {
	r = ChatReminderService{}
	r.Options = opts
	return
}

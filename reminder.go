// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo

import (
	"context"
	"net/http"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/shared"
)

// Set and clear reminders for chats
//
// ReminderService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReminderService] method instead.
type ReminderService struct {
	Options []option.RequestOption
}

// NewReminderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReminderService(opts ...option.RequestOption) (r ReminderService) {
	r = ReminderService{}
	r.Options = opts
	return
}

// Clear an existing reminder from a chat
func (r *ReminderService) ClearChatReminder(ctx context.Context, body ReminderClearChatReminderParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/clear-chat-reminder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Set a reminder for a chat at a specific time
func (r *ReminderService) SetChatReminder(ctx context.Context, body ReminderSetChatReminderParams, opts ...option.RequestOption) (res *shared.BaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/set-chat-reminder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ReminderClearChatReminderParams struct {
	// The identifier of the chat to clear reminder from
	ChatID string `json:"chatID,required"`
	paramObj
}

func (r ReminderClearChatReminderParams) MarshalJSON() (data []byte, err error) {
	type shadow ReminderClearChatReminderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReminderClearChatReminderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReminderSetChatReminderParams struct {
	// The identifier of the chat to set reminder for
	ChatID string `json:"chatID,required"`
	// Reminder configuration
	Reminder ReminderSetChatReminderParamsReminder `json:"reminder,omitzero,required"`
	paramObj
}

func (r ReminderSetChatReminderParams) MarshalJSON() (data []byte, err error) {
	type shadow ReminderSetChatReminderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReminderSetChatReminderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reminder configuration
//
// The property RemindAtMs is required.
type ReminderSetChatReminderParamsReminder struct {
	// Unix timestamp in milliseconds when reminder should trigger
	RemindAtMs float64 `json:"remindAtMs,required"`
	// Cancel reminder if someone messages in the chat
	DismissOnIncomingMessage param.Opt[bool] `json:"dismissOnIncomingMessage,omitzero"`
	paramObj
}

func (r ReminderSetChatReminderParamsReminder) MarshalJSON() (data []byte, err error) {
	type shadow ReminderSetChatReminderParamsReminder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReminderSetChatReminderParamsReminder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

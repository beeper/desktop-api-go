// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/param"
)

// Manage reminders for chats
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

// Set a reminder for a chat at a specific time
func (r *ChatReminderService) New(ctx context.Context, chatID string, body ChatReminderNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return
	}
	path := fmt.Sprintf("v1/chats/%s/reminders", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Clear an existing reminder from a chat
func (r *ChatReminderService) Delete(ctx context.Context, chatID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if chatID == "" {
		err = errors.New("missing required chatID parameter")
		return
	}
	path := fmt.Sprintf("v1/chats/%s/reminders", chatID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ChatReminderNewParams struct {
	// Reminder configuration
	Reminder ChatReminderNewParamsReminder `json:"reminder,omitzero,required"`
	paramObj
}

func (r ChatReminderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatReminderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatReminderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reminder configuration
//
// The property RemindAtMs is required.
type ChatReminderNewParamsReminder struct {
	// Unix timestamp in milliseconds when reminder should trigger
	RemindAtMs float64 `json:"remindAtMs,required"`
	// Cancel reminder if someone messages in the chat
	DismissOnIncomingMessage param.Opt[bool] `json:"dismissOnIncomingMessage,omitzero"`
	paramObj
}

func (r ChatReminderNewParamsReminder) MarshalJSON() (data []byte, err error) {
	type shadow ChatReminderNewParamsReminder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatReminderNewParamsReminder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

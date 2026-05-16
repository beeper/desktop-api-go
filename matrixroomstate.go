// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/apiquery"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// MatrixRoomStateService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixRoomStateService] method instead.
type MatrixRoomStateService struct {
	Options []option.RequestOption
}

// NewMatrixRoomStateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatrixRoomStateService(opts ...option.RequestOption) (r MatrixRoomStateService) {
	r = MatrixRoomStateService{}
	r.Options = opts
	return
}

// Looks up the contents of a state event in a room. If the user is joined to the
// room then the state is taken from the current state of the room. If the user has
// left the room then the state is taken from the state of the room when they left.
func (r *MatrixRoomStateService) Get(ctx context.Context, stateKey string, params MatrixRoomStateGetParams, opts ...option.RequestOption) (res *MatrixRoomStateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.RoomID == "" {
		err = errors.New("missing required roomId parameter")
		return nil, err
	}
	if params.EventType == "" {
		err = errors.New("missing required eventType parameter")
		return nil, err
	}
	if stateKey == "" {
		err = errors.New("missing required stateKey parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/v3/rooms/%s/state/%s/%s", params.RoomID, params.EventType, stateKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get the state events for the current state of a room.
func (r *MatrixRoomStateService) List(ctx context.Context, roomID string, opts ...option.RequestOption) (res *[]MatrixRoomStateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomID == "" {
		err = errors.New("missing required roomId parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/v3/rooms/%s/state", roomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MatrixRoomStateGetResponse map[string]any

// The format used for events when they are returned from a homeserver to a client
// via the Client-Server API, or sent to an Application Service via the Application
// Services API.
type MatrixRoomStateListResponse struct {
	// The body of this event, as created by the client which sent it.
	Content any `json:"content" api:"required"`
	// The globally unique identifier for this event.
	EventID string `json:"event_id" api:"required"`
	// Timestamp (in milliseconds since the unix epoch) on originating homeserver when
	// this event was sent.
	OriginServerTs int64 `json:"origin_server_ts" api:"required"`
	// The ID of the room associated with this event.
	RoomID string `json:"room_id" api:"required"`
	// Contains the fully-qualified ID of the user who sent this event.
	Sender string `json:"sender" api:"required"`
	// The type of the event.
	Type string `json:"type" api:"required"`
	// Present if, and only if, this event is a _state_ event. The key making this
	// piece of state unique in the room. Note that it is often an empty string.
	//
	// State keys starting with an `@` are reserved for referencing user IDs, such as
	// room members. With the exception of a few events, state events set with a given
	// user's ID as the state key MUST only be set by that user.
	StateKey string                              `json:"state_key"`
	Unsigned MatrixRoomStateListResponseUnsigned `json:"unsigned"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content        respjson.Field
		EventID        respjson.Field
		OriginServerTs respjson.Field
		RoomID         respjson.Field
		Sender         respjson.Field
		Type           respjson.Field
		StateKey       respjson.Field
		Unsigned       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixRoomStateListResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixRoomStateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixRoomStateListResponseUnsigned struct {
	// The time in milliseconds that has elapsed since the event was sent. This field
	// is generated by the local homeserver, and may be incorrect if the local time on
	// at least one of the two servers is out of sync, which can cause the age to
	// either be negative or greater than it actually is.
	Age int64 `json:"age"`
	// The room membership of the user making the request, at the time of the event.
	//
	// This property is the value of the `membership` property of the requesting user's
	// [`m.room.member`](https://spec.matrix.org/v1.18/client-server-api#mroommember)
	// state at the point of the event, including any changes caused by the event. If
	// the user had yet to join the room at the time of the event (i.e, they have no
	// `m.room.member` state), this property is set to `leave`.
	//
	// Homeservers SHOULD populate this property wherever practical, but they MAY omit
	// it if necessary (for example, if calculating the value is expensive, servers
	// might choose to only implement it in encrypted rooms). The property is _not_
	// normally populated in events pushed to application services via the application
	// service transaction API (where there is no clear definition of "requesting
	// user").
	Membership string `json:"membership"`
	// The previous `content` for this event. This field is generated by the local
	// homeserver, and is only returned if the event is a state event, and the client
	// has permission to see the previous content.
	PrevContent     any `json:"prev_content"`
	RedactedBecause any `json:"redacted_because"`
	// The client-supplied
	// [transaction ID](https://spec.matrix.org/v1.18/client-server-api/#transaction-identifiers),
	// for example, provided via
	// `PUT /_matrix/client/v3/rooms/{roomId}/send/{eventType}/{txnId}`, if the client
	// being given the event is the same one which sent it.
	TransactionID string `json:"transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Age             respjson.Field
		Membership      respjson.Field
		PrevContent     respjson.Field
		RedactedBecause respjson.Field
		TransactionID   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixRoomStateListResponseUnsigned) RawJSON() string { return r.JSON.raw }
func (r *MatrixRoomStateListResponseUnsigned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixRoomStateGetParams struct {
	RoomID    string `path:"roomId" api:"required" json:"-"`
	EventType string `path:"eventType" api:"required" json:"-"`
	// The format to use for the returned data. `content` (the default) will return
	// only the content of the state event. `event` will return the entire event in the
	// usual format suitable for clients, including fields like event ID, sender and
	// timestamp.
	//
	// Any of "content", "event".
	Format MatrixRoomStateGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MatrixRoomStateGetParams]'s query parameters as
// `url.Values`.
func (r MatrixRoomStateGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The format to use for the returned data. `content` (the default) will return
// only the content of the state event. `event` will return the entire event in the
// usual format suitable for clients, including fields like event ID, sender and
// timestamp.
type MatrixRoomStateGetParamsFormat string

const (
	MatrixRoomStateGetParamsFormatContent MatrixRoomStateGetParamsFormat = "content"
	MatrixRoomStateGetParamsFormatEvent   MatrixRoomStateGetParamsFormat = "event"
)

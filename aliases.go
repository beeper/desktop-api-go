// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"github.com/beeper/desktop-api-go/internal/apierror"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type Attachment = shared.Attachment

// Attachment type.
//
// This is an alias to an internal type.
type AttachmentType = shared.AttachmentType

// Equals "unknown"
const AttachmentTypeUnknown = shared.AttachmentTypeUnknown

// Equals "img"
const AttachmentTypeImg = shared.AttachmentTypeImg

// Equals "video"
const AttachmentTypeVideo = shared.AttachmentTypeVideo

// Equals "audio"
const AttachmentTypeAudio = shared.AttachmentTypeAudio

// Pixel dimensions of the attachment: width/height in px.
//
// This is an alias to an internal type.
type AttachmentSize = shared.AttachmentSize

// This is an alias to an internal type.
type Message = shared.Message

// This is an alias to an internal type.
type Reaction = shared.Reaction

// User the account belongs to.
//
// This is an alias to an internal type.
type User = shared.User

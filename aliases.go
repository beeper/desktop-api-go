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

// Message content type. Useful for distinguishing reactions, media messages, and
// state events from regular text messages.
//
// This is an alias to an internal type.
type MessageType = shared.MessageType

// Equals "TEXT"
const MessageTypeText = shared.MessageTypeText

// Equals "NOTICE"
const MessageTypeNotice = shared.MessageTypeNotice

// Equals "IMAGE"
const MessageTypeImage = shared.MessageTypeImage

// Equals "VIDEO"
const MessageTypeVideo = shared.MessageTypeVideo

// Equals "VOICE"
const MessageTypeVoice = shared.MessageTypeVoice

// Equals "AUDIO"
const MessageTypeAudio = shared.MessageTypeAudio

// Equals "FILE"
const MessageTypeFile = shared.MessageTypeFile

// Equals "STICKER"
const MessageTypeSticker = shared.MessageTypeSticker

// Equals "LOCATION"
const MessageTypeLocation = shared.MessageTypeLocation

// Equals "REACTION"
const MessageTypeReaction = shared.MessageTypeReaction

// This is an alias to an internal type.
type Reaction = shared.Reaction

// User the account belongs to.
//
// This is an alias to an internal type.
type User = shared.User

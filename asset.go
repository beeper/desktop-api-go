// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
)

// Manage assets in Beeper Desktop, like message attachments
//
// AssetService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r AssetService) {
	r = AssetService{}
	r.Options = opts
	return
}

// Download a Matrix asset using its mxc:// or localmxc:// URL to the device
// running Beeper Desktop and return the local file URL.
func (r *AssetService) Download(ctx context.Context, body AssetDownloadParams, opts ...option.RequestOption) (res *AssetDownloadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/assets/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Upload a file to a temporary location. Supports JSON body with base64 `content`
// field, or multipart/form-data with `file` field. Returns a local file URL that
// can be used when sending messages with attachments.
func (r *AssetService) Upload(ctx context.Context, body AssetUploadParams, opts ...option.RequestOption) (res *AssetUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/assets/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AssetDownloadResponse struct {
	// Error message if the download failed.
	Error string `json:"error"`
	// Local file URL to the downloaded asset.
	SrcURL string `json:"srcURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		SrcURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetDownloadResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetDownloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetUploadResponse struct {
	// Duration in seconds (audio/videos)
	Duration float64 `json:"duration"`
	// Error message if upload failed
	Error string `json:"error"`
	// Resolved filename
	FileName string `json:"fileName"`
	// File size in bytes
	FileSize float64 `json:"fileSize"`
	// Height in pixels (images/videos)
	Height float64 `json:"height"`
	// Detected or provided MIME type
	MimeType string `json:"mimeType"`
	// Local file URL (file://) for the uploaded asset
	SrcURL string `json:"srcURL"`
	// Unique upload ID for this asset
	UploadID string `json:"uploadID"`
	// Width in pixels (images/videos)
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Error       respjson.Field
		FileName    respjson.Field
		FileSize    respjson.Field
		Height      respjson.Field
		MimeType    respjson.Field
		SrcURL      respjson.Field
		UploadID    respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetDownloadParams struct {
	// Matrix content URL (mxc:// or localmxc://) for the asset to download.
	URL string `json:"url,required"`
	paramObj
}

func (r AssetDownloadParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetDownloadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetDownloadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetUploadParams struct {
	// Base64-encoded file content (max ~500MB decoded)
	Content string `json:"content,required"`
	// Original filename. Generated if omitted
	FileName param.Opt[string] `json:"fileName,omitzero"`
	// MIME type. Auto-detected from magic bytes if omitted
	MimeType param.Opt[string] `json:"mimeType,omitzero"`
	paramObj
}

func (r AssetUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

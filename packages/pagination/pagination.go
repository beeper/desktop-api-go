// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/stainless-sdks/beeper-desktop-api-go/internal/apijson"
	"github.com/stainless-sdks/beeper-desktop-api-go/internal/requestconfig"
	"github.com/stainless-sdks/beeper-desktop-api-go/option"
	"github.com/stainless-sdks/beeper-desktop-api-go/packages/param"
	"github.com/stainless-sdks/beeper-desktop-api-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Cursor[T any] struct {
	Items        []T    `json:"items"`
	HasMore      bool   `json:"hasMore"`
	OldestCursor string `json:"oldestCursor,nullable"`
	NewestCursor string `json:"newestCursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items        respjson.Field
		HasMore      respjson.Field
		OldestCursor respjson.Field
		NewestCursor respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r Cursor[T]) RawJSON() string { return r.JSON.raw }
func (r *Cursor[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *Cursor[T]) GetNextPage() (res *Cursor[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	next := r.OldestCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Cursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &Cursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorAutoPager[T any] struct {
	page *Cursor[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorAutoPager[T any](page *Cursor[T], err error) *CursorAutoPager[T] {
	return &CursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorAutoPager[T]) Index() int {
	return r.run
}

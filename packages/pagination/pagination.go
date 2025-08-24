// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"reflect"

	"github.com/beeper/beeper-desktop-api-go/internal/apijson"
	"github.com/beeper/beeper-desktop-api-go/internal/requestconfig"
	"github.com/beeper/beeper-desktop-api-go/option"
	"github.com/beeper/beeper-desktop-api-go/packages/param"
	"github.com/beeper/beeper-desktop-api-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CursorID[T any] struct {
	Data    []T  `json:"data"`
	HasMore bool `json:"has_more"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorID[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorID[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorID[T]) GetNextPage() (res *CursorID[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	items := r.Data
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("starting_after", field.Interface().(string)))
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

func (r *CursorID[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorID[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDAutoPager[T any] struct {
	page *CursorID[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDAutoPager[T any](page *CursorID[T], err error) *CursorIDAutoPager[T] {
	return &CursorIDAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDAutoPager[T]) Index() int {
	return r.run
}

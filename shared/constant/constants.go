// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/beeper/desktop-api-go/v5/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Complete string       // Always "complete"
type Cookies string        // Always "cookies"
type DisplayAndWait string // Always "display_and_wait"
type UserInput string      // Always "user_input"

func (c Complete) Default() Complete             { return "complete" }
func (c Cookies) Default() Cookies               { return "cookies" }
func (c DisplayAndWait) Default() DisplayAndWait { return "display_and_wait" }
func (c UserInput) Default() UserInput           { return "user_input" }

func (c Complete) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Cookies) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c DisplayAndWait) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c UserInput) MarshalJSON() ([]byte, error)      { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}

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

type ByContinuingYouAgreeToTheTermsOfUseAndAcknowledgeThePrivacyPolicy string // Always "By continuing, you agree to the Terms of Use and acknowledge the Privacy Policy."
type Cancelled string                                                         // Always "cancelled"
type ChooseYourUsername string                                                // Always "Choose your username"
type Complete string                                                          // Always "complete"
type Continue string                                                          // Always "Continue"
type Cookies string                                                           // Always "cookies"
type DisplayAndWait string                                                    // Always "display_and_wait"
type Emoji string                                                             // Always "emoji"
type Nothing string                                                           // Always "nothing"
type Qr string                                                                // Always "qr"
type UserInput string                                                         // Always "user_input"
type Username string                                                          // Always "Username"

func (c ByContinuingYouAgreeToTheTermsOfUseAndAcknowledgeThePrivacyPolicy) Default() ByContinuingYouAgreeToTheTermsOfUseAndAcknowledgeThePrivacyPolicy {
	return "By continuing, you agree to the Terms of Use and acknowledge the Privacy Policy."
}
func (c Cancelled) Default() Cancelled                   { return "cancelled" }
func (c ChooseYourUsername) Default() ChooseYourUsername { return "Choose your username" }
func (c Complete) Default() Complete                     { return "complete" }
func (c Continue) Default() Continue                     { return "Continue" }
func (c Cookies) Default() Cookies                       { return "cookies" }
func (c DisplayAndWait) Default() DisplayAndWait         { return "display_and_wait" }
func (c Emoji) Default() Emoji                           { return "emoji" }
func (c Nothing) Default() Nothing                       { return "nothing" }
func (c Qr) Default() Qr                                 { return "qr" }
func (c UserInput) Default() UserInput                   { return "user_input" }
func (c Username) Default() Username                     { return "Username" }

func (c ByContinuingYouAgreeToTheTermsOfUseAndAcknowledgeThePrivacyPolicy) MarshalJSON() ([]byte, error) {
	return marshalString(c)
}
func (c Cancelled) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c ChooseYourUsername) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Complete) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Continue) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Cookies) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c DisplayAndWait) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Emoji) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Nothing) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Qr) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c UserInput) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Username) MarshalJSON() ([]byte, error)           { return marshalString(c) }

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

package forms

import (
	"fmt"
	"regexp"
)

type Errors map[string][]string

var EmailRx = regexp.MustCompile(`^\S+@\S+$`)

type Lengthable[Q any, U comparable] interface {
	[]Q | map[U]Q
}

type NumericComparable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Validate records the provided error, if not nil, inside the errors list
// marked against the provided field.
func AddError(field string, errors Errors, msg string) {
	errors[field] = append(errors[field], msg)
}

// StringLength Checks that a string has either an exact count of characters,
// or fits within the specified range of m to n (inclusive).
func IsStringLength(
	field string,
	errors Errors,
	v string,
	m int,
	n int,
) {
	var msg string
	if m == n {
		msg = fmt.Sprintf("Must be exactly %d characters long", m)
	} else {
		msg = fmt.Sprintf("Must be between %d and %d characters long", m, n)
	}

	if len(v) < m || len(v) > n {
		AddError(field, errors, msg)
	}
}

// NumberBetween Checks that the integer typed variable is exactly m == n in
// size, or between m and n inclusive.
func IsNumberBetween[T NumericComparable](
	field string,
	errors Errors,
	v T,
	m T,
	n T,
) {
	var msg string

	if m == n {
		msg = fmt.Sprintf("Must be exactly %d, but was %d", m, v)
	} else {
		msg = fmt.Sprintf("Must be between %d and %d, but was %d", m, n, v)
	}

	if v < m || v > n {
		AddError(field, errors, msg)
	}
}

// Size checks that an array or map has either exactly m == n entries, or
// between m and n entries (inclusive)
func IsSize[T Lengthable[Q, U], Q any, U comparable](
	field string,
	errors Errors,
	v T,
	m int,
	n int,
) {
	var msg string
	if m == n {
		msg = fmt.Sprintf("Must have exactly %d entries, but had %d", m, len(v))
	} else {
		msg = fmt.Sprintf("Must have between %d and %d entries, but had %d", m, n, len(v))
	}

	if len(v) < m || len(v) > n {
		AddError(field, errors, msg)
	}
}

// MinSize checks that an array or map has at least n entries
func IsMinSize[T Lengthable[Q, U], Q any, U comparable](
	field string,
	errors Errors,
	v T,
	n int,
) {
	entry := "entry"
	if n > 1 {
		entry = "entries"
	}

	var msg string
	msg = fmt.Sprintf("Must have a minimum of %d %s, but had %d", n, entry, len(v))

	if len(v) < n {
		AddError(field, errors, msg)
	}
}

// Regex Confirms that value matches the provided regex
func IsRegex(
	field string,
	errors Errors,
	v string,
	rx *regexp.Regexp,
	message string,
) {
	if !rx.MatchString(v) {
		AddError(field, errors, message)
	}
}

// Email Confirms that value matches our provided email regex.  For a custom
// email regex, use Regex.
func IsEmail(
	field string,
	errors Errors,
	v string,
) {
	IsRegex(field, errors, v, EmailRx, "Email address is invalid")
}

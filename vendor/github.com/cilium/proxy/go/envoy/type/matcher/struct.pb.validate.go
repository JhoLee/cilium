// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/matcher/struct.proto

package matcher

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on StructMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StructMatcher) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StructMatcher with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StructMatcherMultiError, or
// nil if none found.
func (m *StructMatcher) ValidateAll() error {
	return m.validate(true)
}

func (m *StructMatcher) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetPath()) < 1 {
		err := StructMatcherValidationError{
			field:  "Path",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetPath() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StructMatcherValidationError{
						field:  fmt.Sprintf("Path[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StructMatcherValidationError{
						field:  fmt.Sprintf("Path[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StructMatcherValidationError{
					field:  fmt.Sprintf("Path[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetValue() == nil {
		err := StructMatcherValidationError{
			field:  "Value",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StructMatcherValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StructMatcherValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StructMatcherValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StructMatcherMultiError(errors)
	}
	return nil
}

// StructMatcherMultiError is an error wrapping multiple validation errors
// returned by StructMatcher.ValidateAll() if the designated constraints
// aren't met.
type StructMatcherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StructMatcherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StructMatcherMultiError) AllErrors() []error { return m }

// StructMatcherValidationError is the validation error returned by
// StructMatcher.Validate if the designated constraints aren't met.
type StructMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StructMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StructMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StructMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StructMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StructMatcherValidationError) ErrorName() string { return "StructMatcherValidationError" }

// Error satisfies the builtin error interface
func (e StructMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStructMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StructMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StructMatcherValidationError{}

// Validate checks the field values on StructMatcher_PathSegment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StructMatcher_PathSegment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StructMatcher_PathSegment with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StructMatcher_PathSegmentMultiError, or nil if none found.
func (m *StructMatcher_PathSegment) ValidateAll() error {
	return m.validate(true)
}

func (m *StructMatcher_PathSegment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Segment.(type) {

	case *StructMatcher_PathSegment_Key:

		if utf8.RuneCountInString(m.GetKey()) < 1 {
			err := StructMatcher_PathSegmentValidationError{
				field:  "Key",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		err := StructMatcher_PathSegmentValidationError{
			field:  "Segment",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return StructMatcher_PathSegmentMultiError(errors)
	}
	return nil
}

// StructMatcher_PathSegmentMultiError is an error wrapping multiple validation
// errors returned by StructMatcher_PathSegment.ValidateAll() if the
// designated constraints aren't met.
type StructMatcher_PathSegmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StructMatcher_PathSegmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StructMatcher_PathSegmentMultiError) AllErrors() []error { return m }

// StructMatcher_PathSegmentValidationError is the validation error returned by
// StructMatcher_PathSegment.Validate if the designated constraints aren't met.
type StructMatcher_PathSegmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StructMatcher_PathSegmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StructMatcher_PathSegmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StructMatcher_PathSegmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StructMatcher_PathSegmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StructMatcher_PathSegmentValidationError) ErrorName() string {
	return "StructMatcher_PathSegmentValidationError"
}

// Error satisfies the builtin error interface
func (e StructMatcher_PathSegmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStructMatcher_PathSegment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StructMatcher_PathSegmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StructMatcher_PathSegmentValidationError{}

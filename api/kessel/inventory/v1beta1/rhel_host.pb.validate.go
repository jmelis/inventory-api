// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/rhel_host.proto

package v1beta1

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

// Validate checks the field values on RhelHost with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RhelHost) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RhelHost with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RhelHostMultiError, or nil
// if none found.
func (m *RhelHost) ValidateAll() error {
	return m.validate(true)
}

func (m *RhelHost) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RhelHostValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RhelHostValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RhelHostValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetReporterData() == nil {
		err := RhelHostValidationError{
			field:  "ReporterData",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetReporterData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RhelHostValidationError{
					field:  "ReporterData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RhelHostValidationError{
					field:  "ReporterData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReporterData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RhelHostValidationError{
				field:  "ReporterData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RhelHostMultiError(errors)
	}

	return nil
}

// RhelHostMultiError is an error wrapping multiple validation errors returned
// by RhelHost.ValidateAll() if the designated constraints aren't met.
type RhelHostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RhelHostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RhelHostMultiError) AllErrors() []error { return m }

// RhelHostValidationError is the validation error returned by
// RhelHost.Validate if the designated constraints aren't met.
type RhelHostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RhelHostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RhelHostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RhelHostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RhelHostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RhelHostValidationError) ErrorName() string { return "RhelHostValidationError" }

// Error satisfies the builtin error interface
func (e RhelHostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRhelHost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RhelHostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RhelHostValidationError{}

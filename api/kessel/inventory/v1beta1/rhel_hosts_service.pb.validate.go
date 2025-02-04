// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/rhel_hosts_service.proto

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

// Validate checks the field values on CreateRhelHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRhelHostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRhelHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRhelHostRequestMultiError, or nil if none found.
func (m *CreateRhelHostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRhelHostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetRhelHost() == nil {
		err := CreateRhelHostRequestValidationError{
			field:  "RhelHost",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRhelHost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRhelHostRequestValidationError{
					field:  "RhelHost",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRhelHostRequestValidationError{
					field:  "RhelHost",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRhelHost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRhelHostRequestValidationError{
				field:  "RhelHost",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRhelHostRequestMultiError(errors)
	}

	return nil
}

// CreateRhelHostRequestMultiError is an error wrapping multiple validation
// errors returned by CreateRhelHostRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateRhelHostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRhelHostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRhelHostRequestMultiError) AllErrors() []error { return m }

// CreateRhelHostRequestValidationError is the validation error returned by
// CreateRhelHostRequest.Validate if the designated constraints aren't met.
type CreateRhelHostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRhelHostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRhelHostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRhelHostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRhelHostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRhelHostRequestValidationError) ErrorName() string {
	return "CreateRhelHostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRhelHostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRhelHostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRhelHostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRhelHostRequestValidationError{}

// Validate checks the field values on CreateRhelHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateRhelHostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRhelHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRhelHostResponseMultiError, or nil if none found.
func (m *CreateRhelHostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRhelHostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateRhelHostResponseMultiError(errors)
	}

	return nil
}

// CreateRhelHostResponseMultiError is an error wrapping multiple validation
// errors returned by CreateRhelHostResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateRhelHostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRhelHostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRhelHostResponseMultiError) AllErrors() []error { return m }

// CreateRhelHostResponseValidationError is the validation error returned by
// CreateRhelHostResponse.Validate if the designated constraints aren't met.
type CreateRhelHostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRhelHostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRhelHostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRhelHostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRhelHostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRhelHostResponseValidationError) ErrorName() string {
	return "CreateRhelHostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRhelHostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRhelHostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRhelHostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRhelHostResponseValidationError{}

// Validate checks the field values on UpdateRhelHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateRhelHostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRhelHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRhelHostRequestMultiError, or nil if none found.
func (m *UpdateRhelHostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRhelHostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetResource()) < 1 {
		err := UpdateRhelHostRequestValidationError{
			field:  "Resource",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetRhelHost() == nil {
		err := UpdateRhelHostRequestValidationError{
			field:  "RhelHost",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRhelHost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateRhelHostRequestValidationError{
					field:  "RhelHost",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateRhelHostRequestValidationError{
					field:  "RhelHost",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRhelHost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRhelHostRequestValidationError{
				field:  "RhelHost",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateRhelHostRequestMultiError(errors)
	}

	return nil
}

// UpdateRhelHostRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateRhelHostRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateRhelHostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRhelHostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRhelHostRequestMultiError) AllErrors() []error { return m }

// UpdateRhelHostRequestValidationError is the validation error returned by
// UpdateRhelHostRequest.Validate if the designated constraints aren't met.
type UpdateRhelHostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRhelHostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRhelHostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRhelHostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRhelHostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRhelHostRequestValidationError) ErrorName() string {
	return "UpdateRhelHostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRhelHostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRhelHostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRhelHostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRhelHostRequestValidationError{}

// Validate checks the field values on UpdateRhelHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateRhelHostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRhelHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRhelHostResponseMultiError, or nil if none found.
func (m *UpdateRhelHostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRhelHostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateRhelHostResponseMultiError(errors)
	}

	return nil
}

// UpdateRhelHostResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateRhelHostResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateRhelHostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRhelHostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRhelHostResponseMultiError) AllErrors() []error { return m }

// UpdateRhelHostResponseValidationError is the validation error returned by
// UpdateRhelHostResponse.Validate if the designated constraints aren't met.
type UpdateRhelHostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRhelHostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRhelHostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRhelHostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRhelHostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRhelHostResponseValidationError) ErrorName() string {
	return "UpdateRhelHostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRhelHostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRhelHostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRhelHostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRhelHostResponseValidationError{}

// Validate checks the field values on DeleteRhelHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRhelHostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRhelHostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRhelHostRequestMultiError, or nil if none found.
func (m *DeleteRhelHostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRhelHostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetResource()) < 1 {
		err := DeleteRhelHostRequestValidationError{
			field:  "Resource",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteRhelHostRequestMultiError(errors)
	}

	return nil
}

// DeleteRhelHostRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteRhelHostRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteRhelHostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRhelHostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRhelHostRequestMultiError) AllErrors() []error { return m }

// DeleteRhelHostRequestValidationError is the validation error returned by
// DeleteRhelHostRequest.Validate if the designated constraints aren't met.
type DeleteRhelHostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRhelHostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRhelHostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRhelHostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRhelHostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRhelHostRequestValidationError) ErrorName() string {
	return "DeleteRhelHostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRhelHostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRhelHostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRhelHostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRhelHostRequestValidationError{}

// Validate checks the field values on DeleteRhelHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRhelHostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRhelHostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRhelHostResponseMultiError, or nil if none found.
func (m *DeleteRhelHostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRhelHostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteRhelHostResponseMultiError(errors)
	}

	return nil
}

// DeleteRhelHostResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteRhelHostResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteRhelHostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRhelHostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRhelHostResponseMultiError) AllErrors() []error { return m }

// DeleteRhelHostResponseValidationError is the validation error returned by
// DeleteRhelHostResponse.Validate if the designated constraints aren't met.
type DeleteRhelHostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRhelHostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRhelHostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRhelHostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRhelHostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRhelHostResponseValidationError) ErrorName() string {
	return "DeleteRhelHostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRhelHostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRhelHostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRhelHostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRhelHostResponseValidationError{}

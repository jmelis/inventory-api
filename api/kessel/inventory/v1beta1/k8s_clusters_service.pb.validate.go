// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/k8s_clusters_service.proto

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

// Validate checks the field values on CreateK8SClusterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateK8SClusterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateK8SClusterRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateK8SClusterRequestMultiError, or nil if none found.
func (m *CreateK8SClusterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateK8SClusterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetK8SCluster()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateK8SClusterRequestValidationError{
					field:  "K8SCluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateK8SClusterRequestValidationError{
					field:  "K8SCluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetK8SCluster()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateK8SClusterRequestValidationError{
				field:  "K8SCluster",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateK8SClusterRequestMultiError(errors)
	}

	return nil
}

// CreateK8SClusterRequestMultiError is an error wrapping multiple validation
// errors returned by CreateK8SClusterRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateK8SClusterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateK8SClusterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateK8SClusterRequestMultiError) AllErrors() []error { return m }

// CreateK8SClusterRequestValidationError is the validation error returned by
// CreateK8SClusterRequest.Validate if the designated constraints aren't met.
type CreateK8SClusterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateK8SClusterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateK8SClusterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateK8SClusterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateK8SClusterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateK8SClusterRequestValidationError) ErrorName() string {
	return "CreateK8SClusterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateK8SClusterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateK8SClusterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateK8SClusterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateK8SClusterRequestValidationError{}

// Validate checks the field values on CreateK8SClusterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateK8SClusterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateK8SClusterResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateK8SClusterResponseMultiError, or nil if none found.
func (m *CreateK8SClusterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateK8SClusterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateK8SClusterResponseMultiError(errors)
	}

	return nil
}

// CreateK8SClusterResponseMultiError is an error wrapping multiple validation
// errors returned by CreateK8SClusterResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateK8SClusterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateK8SClusterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateK8SClusterResponseMultiError) AllErrors() []error { return m }

// CreateK8SClusterResponseValidationError is the validation error returned by
// CreateK8SClusterResponse.Validate if the designated constraints aren't met.
type CreateK8SClusterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateK8SClusterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateK8SClusterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateK8SClusterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateK8SClusterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateK8SClusterResponseValidationError) ErrorName() string {
	return "CreateK8SClusterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateK8SClusterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateK8SClusterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateK8SClusterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateK8SClusterResponseValidationError{}

// Validate checks the field values on UpdateK8SClusterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateK8SClusterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateK8SClusterRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateK8SClusterRequestMultiError, or nil if none found.
func (m *UpdateK8SClusterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateK8SClusterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Resource

	if all {
		switch v := interface{}(m.GetK8SCluster()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateK8SClusterRequestValidationError{
					field:  "K8SCluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateK8SClusterRequestValidationError{
					field:  "K8SCluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetK8SCluster()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateK8SClusterRequestValidationError{
				field:  "K8SCluster",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateK8SClusterRequestMultiError(errors)
	}

	return nil
}

// UpdateK8SClusterRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateK8SClusterRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateK8SClusterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateK8SClusterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateK8SClusterRequestMultiError) AllErrors() []error { return m }

// UpdateK8SClusterRequestValidationError is the validation error returned by
// UpdateK8SClusterRequest.Validate if the designated constraints aren't met.
type UpdateK8SClusterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateK8SClusterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateK8SClusterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateK8SClusterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateK8SClusterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateK8SClusterRequestValidationError) ErrorName() string {
	return "UpdateK8SClusterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateK8SClusterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateK8SClusterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateK8SClusterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateK8SClusterRequestValidationError{}

// Validate checks the field values on UpdateK8SClusterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateK8SClusterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateK8SClusterResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateK8SClusterResponseMultiError, or nil if none found.
func (m *UpdateK8SClusterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateK8SClusterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateK8SClusterResponseMultiError(errors)
	}

	return nil
}

// UpdateK8SClusterResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateK8SClusterResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateK8SClusterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateK8SClusterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateK8SClusterResponseMultiError) AllErrors() []error { return m }

// UpdateK8SClusterResponseValidationError is the validation error returned by
// UpdateK8SClusterResponse.Validate if the designated constraints aren't met.
type UpdateK8SClusterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateK8SClusterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateK8SClusterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateK8SClusterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateK8SClusterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateK8SClusterResponseValidationError) ErrorName() string {
	return "UpdateK8SClusterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateK8SClusterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateK8SClusterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateK8SClusterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateK8SClusterResponseValidationError{}

// Validate checks the field values on DeleteK8SClusterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteK8SClusterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteK8SClusterRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteK8SClusterRequestMultiError, or nil if none found.
func (m *DeleteK8SClusterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteK8SClusterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Resource

	if len(errors) > 0 {
		return DeleteK8SClusterRequestMultiError(errors)
	}

	return nil
}

// DeleteK8SClusterRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteK8SClusterRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteK8SClusterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteK8SClusterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteK8SClusterRequestMultiError) AllErrors() []error { return m }

// DeleteK8SClusterRequestValidationError is the validation error returned by
// DeleteK8SClusterRequest.Validate if the designated constraints aren't met.
type DeleteK8SClusterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteK8SClusterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteK8SClusterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteK8SClusterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteK8SClusterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteK8SClusterRequestValidationError) ErrorName() string {
	return "DeleteK8SClusterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteK8SClusterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteK8SClusterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteK8SClusterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteK8SClusterRequestValidationError{}

// Validate checks the field values on DeleteK8SClusterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteK8SClusterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteK8SClusterResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteK8SClusterResponseMultiError, or nil if none found.
func (m *DeleteK8SClusterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteK8SClusterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteK8SClusterResponseMultiError(errors)
	}

	return nil
}

// DeleteK8SClusterResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteK8SClusterResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteK8SClusterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteK8SClusterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteK8SClusterResponseMultiError) AllErrors() []error { return m }

// DeleteK8SClusterResponseValidationError is the validation error returned by
// DeleteK8SClusterResponse.Validate if the designated constraints aren't met.
type DeleteK8SClusterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteK8SClusterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteK8SClusterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteK8SClusterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteK8SClusterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteK8SClusterResponseValidationError) ErrorName() string {
	return "DeleteK8SClusterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteK8SClusterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteK8SClusterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteK8SClusterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteK8SClusterResponseValidationError{}

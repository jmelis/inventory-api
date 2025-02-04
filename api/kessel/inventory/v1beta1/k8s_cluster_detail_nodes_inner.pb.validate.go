// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/k8s_cluster_detail_nodes_inner.proto

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

// Validate checks the field values on K8SClusterDetailNodesInner with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *K8SClusterDetailNodesInner) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on K8SClusterDetailNodesInner with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// K8SClusterDetailNodesInnerMultiError, or nil if none found.
func (m *K8SClusterDetailNodesInner) ValidateAll() error {
	return m.validate(true)
}

func (m *K8SClusterDetailNodesInner) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := K8SClusterDetailNodesInnerValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCpu()) < 1 {
		err := K8SClusterDetailNodesInnerValidationError{
			field:  "Cpu",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetMemory()) < 1 {
		err := K8SClusterDetailNodesInnerValidationError{
			field:  "Memory",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetLabels() {
		_, _ = idx, item

		if item == nil {
			err := K8SClusterDetailNodesInnerValidationError{
				field:  fmt.Sprintf("Labels[%v]", idx),
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, K8SClusterDetailNodesInnerValidationError{
						field:  fmt.Sprintf("Labels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, K8SClusterDetailNodesInnerValidationError{
						field:  fmt.Sprintf("Labels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return K8SClusterDetailNodesInnerValidationError{
					field:  fmt.Sprintf("Labels[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return K8SClusterDetailNodesInnerMultiError(errors)
	}

	return nil
}

// K8SClusterDetailNodesInnerMultiError is an error wrapping multiple
// validation errors returned by K8SClusterDetailNodesInner.ValidateAll() if
// the designated constraints aren't met.
type K8SClusterDetailNodesInnerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m K8SClusterDetailNodesInnerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m K8SClusterDetailNodesInnerMultiError) AllErrors() []error { return m }

// K8SClusterDetailNodesInnerValidationError is the validation error returned
// by K8SClusterDetailNodesInner.Validate if the designated constraints aren't met.
type K8SClusterDetailNodesInnerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e K8SClusterDetailNodesInnerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e K8SClusterDetailNodesInnerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e K8SClusterDetailNodesInnerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e K8SClusterDetailNodesInnerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e K8SClusterDetailNodesInnerValidationError) ErrorName() string {
	return "K8SClusterDetailNodesInnerValidationError"
}

// Error satisfies the builtin error interface
func (e K8SClusterDetailNodesInnerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sK8SClusterDetailNodesInner.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = K8SClusterDetailNodesInnerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = K8SClusterDetailNodesInnerValidationError{}

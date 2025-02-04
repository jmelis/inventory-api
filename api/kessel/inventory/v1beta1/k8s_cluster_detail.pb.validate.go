// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/k8s_cluster_detail.proto

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

// Validate checks the field values on K8SClusterDetail with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *K8SClusterDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on K8SClusterDetail with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// K8SClusterDetailMultiError, or nil if none found.
func (m *K8SClusterDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *K8SClusterDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetExternalClusterId()) < 1 {
		err := K8SClusterDetailValidationError{
			field:  "ExternalClusterId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _K8SClusterDetail_ClusterStatus_NotInLookup[m.GetClusterStatus()]; ok {
		err := K8SClusterDetailValidationError{
			field:  "ClusterStatus",
			reason: "value must not be in list [CLUSTER_STATUS_UNSPECIFIED]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := K8SClusterDetail_ClusterStatus_name[int32(m.GetClusterStatus())]; !ok {
		err := K8SClusterDetailValidationError{
			field:  "ClusterStatus",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for KubeVersion

	if _, ok := _K8SClusterDetail_KubeVendor_NotInLookup[m.GetKubeVendor()]; ok {
		err := K8SClusterDetailValidationError{
			field:  "KubeVendor",
			reason: "value must not be in list [KUBE_VENDOR_UNSPECIFIED]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := K8SClusterDetail_KubeVendor_name[int32(m.GetKubeVendor())]; !ok {
		err := K8SClusterDetailValidationError{
			field:  "KubeVendor",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetVendorVersion()) < 1 {
		err := K8SClusterDetailValidationError{
			field:  "VendorVersion",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _K8SClusterDetail_CloudPlatform_NotInLookup[m.GetCloudPlatform()]; ok {
		err := K8SClusterDetailValidationError{
			field:  "CloudPlatform",
			reason: "value must not be in list [CLOUD_PLATFORM_UNSPECIFIED]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := K8SClusterDetail_CloudPlatform_name[int32(m.GetCloudPlatform())]; !ok {
		err := K8SClusterDetailValidationError{
			field:  "CloudPlatform",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetNodes()) < 1 {
		err := K8SClusterDetailValidationError{
			field:  "Nodes",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetNodes() {
		_, _ = idx, item

		if item == nil {
			err := K8SClusterDetailValidationError{
				field:  fmt.Sprintf("Nodes[%v]", idx),
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
					errors = append(errors, K8SClusterDetailValidationError{
						field:  fmt.Sprintf("Nodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, K8SClusterDetailValidationError{
						field:  fmt.Sprintf("Nodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return K8SClusterDetailValidationError{
					field:  fmt.Sprintf("Nodes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return K8SClusterDetailMultiError(errors)
	}

	return nil
}

// K8SClusterDetailMultiError is an error wrapping multiple validation errors
// returned by K8SClusterDetail.ValidateAll() if the designated constraints
// aren't met.
type K8SClusterDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m K8SClusterDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m K8SClusterDetailMultiError) AllErrors() []error { return m }

// K8SClusterDetailValidationError is the validation error returned by
// K8SClusterDetail.Validate if the designated constraints aren't met.
type K8SClusterDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e K8SClusterDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e K8SClusterDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e K8SClusterDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e K8SClusterDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e K8SClusterDetailValidationError) ErrorName() string { return "K8SClusterDetailValidationError" }

// Error satisfies the builtin error interface
func (e K8SClusterDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sK8SClusterDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = K8SClusterDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = K8SClusterDetailValidationError{}

var _K8SClusterDetail_ClusterStatus_NotInLookup = map[K8SClusterDetail_ClusterStatus]struct{}{
	0: {},
}

var _K8SClusterDetail_KubeVendor_NotInLookup = map[K8SClusterDetail_KubeVendor]struct{}{
	0: {},
}

var _K8SClusterDetail_CloudPlatform_NotInLookup = map[K8SClusterDetail_CloudPlatform]struct{}{
	0: {},
}

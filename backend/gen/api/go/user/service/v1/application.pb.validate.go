// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/service/v1/application.proto

package servicev1

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

// Validate checks the field values on Application with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Application) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Application with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ApplicationMultiError, or
// nil if none found.
func (m *Application) ValidateAll() error {
	return m.validate(true)
}

func (m *Application) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.AppId != nil {
		// no validation rules for AppId
	}

	if m.AppKey != nil {
		// no validation rules for AppKey
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.CreatorId != nil {
		// no validation rules for CreatorId
	}

	if m.OwnerId != nil {
		// no validation rules for OwnerId
	}

	if m.Remark != nil {
		// no validation rules for Remark
	}

	if m.KeepMonth != nil {
		// no validation rules for KeepMonth
	}

	if m.CreateTime != nil {
		// no validation rules for CreateTime
	}

	if m.UpdateTime != nil {
		// no validation rules for UpdateTime
	}

	if m.DeleteTime != nil {
		// no validation rules for DeleteTime
	}

	if len(errors) > 0 {
		return ApplicationMultiError(errors)
	}

	return nil
}

// ApplicationMultiError is an error wrapping multiple validation errors
// returned by Application.ValidateAll() if the designated constraints aren't met.
type ApplicationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplicationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplicationMultiError) AllErrors() []error { return m }

// ApplicationValidationError is the validation error returned by
// Application.Validate if the designated constraints aren't met.
type ApplicationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationValidationError) ErrorName() string { return "ApplicationValidationError" }

// Error satisfies the builtin error interface
func (e ApplicationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplication.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationValidationError{}

// Validate checks the field values on ListApplicationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListApplicationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListApplicationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListApplicationResponseMultiError, or nil if none found.
func (m *ListApplicationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListApplicationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListApplicationResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListApplicationResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListApplicationResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListApplicationResponseMultiError(errors)
	}

	return nil
}

// ListApplicationResponseMultiError is an error wrapping multiple validation
// errors returned by ListApplicationResponse.ValidateAll() if the designated
// constraints aren't met.
type ListApplicationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListApplicationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListApplicationResponseMultiError) AllErrors() []error { return m }

// ListApplicationResponseValidationError is the validation error returned by
// ListApplicationResponse.Validate if the designated constraints aren't met.
type ListApplicationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListApplicationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListApplicationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListApplicationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListApplicationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListApplicationResponseValidationError) ErrorName() string {
	return "ListApplicationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListApplicationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListApplicationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListApplicationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListApplicationResponseValidationError{}

// Validate checks the field values on GetApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetApplicationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetApplicationRequestMultiError, or nil if none found.
func (m *GetApplicationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetApplicationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetApplicationRequestMultiError(errors)
	}

	return nil
}

// GetApplicationRequestMultiError is an error wrapping multiple validation
// errors returned by GetApplicationRequest.ValidateAll() if the designated
// constraints aren't met.
type GetApplicationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetApplicationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetApplicationRequestMultiError) AllErrors() []error { return m }

// GetApplicationRequestValidationError is the validation error returned by
// GetApplicationRequest.Validate if the designated constraints aren't met.
type GetApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApplicationRequestValidationError) ErrorName() string {
	return "GetApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApplicationRequestValidationError{}

// Validate checks the field values on CreateApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateApplicationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateApplicationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateApplicationRequestMultiError, or nil if none found.
func (m *CreateApplicationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateApplicationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetApp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateApplicationRequestValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateApplicationRequestValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateApplicationRequestValidationError{
				field:  "App",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return CreateApplicationRequestMultiError(errors)
	}

	return nil
}

// CreateApplicationRequestMultiError is an error wrapping multiple validation
// errors returned by CreateApplicationRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateApplicationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateApplicationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateApplicationRequestMultiError) AllErrors() []error { return m }

// CreateApplicationRequestValidationError is the validation error returned by
// CreateApplicationRequest.Validate if the designated constraints aren't met.
type CreateApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateApplicationRequestValidationError) ErrorName() string {
	return "CreateApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateApplicationRequestValidationError{}

// Validate checks the field values on UpdateApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateApplicationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateApplicationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateApplicationRequestMultiError, or nil if none found.
func (m *UpdateApplicationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateApplicationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetApp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateApplicationRequestValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateApplicationRequestValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateApplicationRequestValidationError{
				field:  "App",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return UpdateApplicationRequestMultiError(errors)
	}

	return nil
}

// UpdateApplicationRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateApplicationRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateApplicationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateApplicationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateApplicationRequestMultiError) AllErrors() []error { return m }

// UpdateApplicationRequestValidationError is the validation error returned by
// UpdateApplicationRequest.Validate if the designated constraints aren't met.
type UpdateApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateApplicationRequestValidationError) ErrorName() string {
	return "UpdateApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateApplicationRequestValidationError{}

// Validate checks the field values on DeleteApplicationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteApplicationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteApplicationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteApplicationRequestMultiError, or nil if none found.
func (m *DeleteApplicationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteApplicationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return DeleteApplicationRequestMultiError(errors)
	}

	return nil
}

// DeleteApplicationRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteApplicationRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteApplicationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteApplicationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteApplicationRequestMultiError) AllErrors() []error { return m }

// DeleteApplicationRequestValidationError is the validation error returned by
// DeleteApplicationRequest.Validate if the designated constraints aren't met.
type DeleteApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteApplicationRequestValidationError) ErrorName() string {
	return "DeleteApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteApplicationRequestValidationError{}

// Validate checks the field values on GetApplicationByAppIdRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetApplicationByAppIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetApplicationByAppIdRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetApplicationByAppIdRequestMultiError, or nil if none found.
func (m *GetApplicationByAppIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetApplicationByAppIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	if len(errors) > 0 {
		return GetApplicationByAppIdRequestMultiError(errors)
	}

	return nil
}

// GetApplicationByAppIdRequestMultiError is an error wrapping multiple
// validation errors returned by GetApplicationByAppIdRequest.ValidateAll() if
// the designated constraints aren't met.
type GetApplicationByAppIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetApplicationByAppIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetApplicationByAppIdRequestMultiError) AllErrors() []error { return m }

// GetApplicationByAppIdRequestValidationError is the validation error returned
// by GetApplicationByAppIdRequest.Validate if the designated constraints
// aren't met.
type GetApplicationByAppIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApplicationByAppIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApplicationByAppIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetApplicationByAppIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApplicationByAppIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApplicationByAppIdRequestValidationError) ErrorName() string {
	return "GetApplicationByAppIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetApplicationByAppIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApplicationByAppIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApplicationByAppIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApplicationByAppIdRequestValidationError{}

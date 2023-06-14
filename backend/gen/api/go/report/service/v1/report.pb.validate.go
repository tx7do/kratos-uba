// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: report/service/v1/report.proto

package v1

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

// Validate checks the field values on UserReport with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserReport) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserReport with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserReportMultiError, or
// nil if none found.
func (m *UserReport) ValidateAll() error {
	return m.validate(true)
}

func (m *UserReport) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for DistinctId

	// no validation rules for TableId

	// no validation rules for Ip

	// no validation rules for ReportType

	// no validation rules for Debug

	// no validation rules for RequestData

	// no validation rules for ReportTime

	// no validation rules for ConsumptionTime

	// no validation rules for EventName

	// no validation rules for Offset

	if len(errors) > 0 {
		return UserReportMultiError(errors)
	}

	return nil
}

// UserReportMultiError is an error wrapping multiple validation errors
// returned by UserReport.ValidateAll() if the designated constraints aren't met.
type UserReportMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserReportMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserReportMultiError) AllErrors() []error { return m }

// UserReportValidationError is the validation error returned by
// UserReport.Validate if the designated constraints aren't met.
type UserReportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserReportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserReportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserReportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserReportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserReportValidationError) ErrorName() string { return "UserReportValidationError" }

// Error satisfies the builtin error interface
func (e UserReportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserReport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserReportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserReportValidationError{}

// Validate checks the field values on EventReport with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EventReport) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventReport with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EventReportMultiError, or
// nil if none found.
func (m *EventReport) ValidateAll() error {
	return m.validate(true)
}

func (m *EventReport) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccountId

	// no validation rules for DistinctId

	// no validation rules for UserId

	// no validation rules for PartEvent

	// no validation rules for PartDate

	// no validation rules for MpPlatform

	// no validation rules for LibVersion

	// no validation rules for Os

	// no validation rules for OsVersion

	// no validation rules for ScreenWidth

	// no validation rules for ScreenHeight

	// no validation rules for DeviceId

	// no validation rules for NetworkType

	// no validation rules for DeviceModel

	// no validation rules for CountryCode

	// no validation rules for Province

	// no validation rules for City

	// no validation rules for Lib

	// no validation rules for Scene

	// no validation rules for Manufacturer

	// no validation rules for Ip

	// no validation rules for TableId

	// no validation rules for Properties

	if len(errors) > 0 {
		return EventReportMultiError(errors)
	}

	return nil
}

// EventReportMultiError is an error wrapping multiple validation errors
// returned by EventReport.ValidateAll() if the designated constraints aren't met.
type EventReportMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventReportMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventReportMultiError) AllErrors() []error { return m }

// EventReportValidationError is the validation error returned by
// EventReport.Validate if the designated constraints aren't met.
type EventReportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventReportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventReportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventReportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventReportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventReportValidationError) ErrorName() string { return "EventReportValidationError" }

// Error satisfies the builtin error interface
func (e EventReportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventReport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventReportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventReportValidationError{}

// Validate checks the field values on AcceptStatusReportData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AcceptStatusReportData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AcceptStatusReportData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AcceptStatusReportDataMultiError, or nil if none found.
func (m *AcceptStatusReportData) ValidateAll() error {
	return m.validate(true)
}

func (m *AcceptStatusReportData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.DataName != nil {
		// no validation rules for DataName
	}

	if m.ReportType != nil {
		// no validation rules for ReportType
	}

	if m.ReportData != nil {
		// no validation rules for ReportData
	}

	if m.ErrorReason != nil {
		// no validation rules for ErrorReason
	}

	if m.ErrorHandling != nil {
		// no validation rules for ErrorHandling
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.PartDate != nil {
		// no validation rules for PartDate
	}

	if len(errors) > 0 {
		return AcceptStatusReportDataMultiError(errors)
	}

	return nil
}

// AcceptStatusReportDataMultiError is an error wrapping multiple validation
// errors returned by AcceptStatusReportData.ValidateAll() if the designated
// constraints aren't met.
type AcceptStatusReportDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AcceptStatusReportDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AcceptStatusReportDataMultiError) AllErrors() []error { return m }

// AcceptStatusReportDataValidationError is the validation error returned by
// AcceptStatusReportData.Validate if the designated constraints aren't met.
type AcceptStatusReportDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AcceptStatusReportDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AcceptStatusReportDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AcceptStatusReportDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AcceptStatusReportDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AcceptStatusReportDataValidationError) ErrorName() string {
	return "AcceptStatusReportDataValidationError"
}

// Error satisfies the builtin error interface
func (e AcceptStatusReportDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAcceptStatusReportData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AcceptStatusReportDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AcceptStatusReportDataValidationError{}

// Validate checks the field values on RealTimeWarehousingData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RealTimeWarehousingData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RealTimeWarehousingData with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RealTimeWarehousingDataMultiError, or nil if none found.
func (m *RealTimeWarehousingData) ValidateAll() error {
	return m.validate(true)
}

func (m *RealTimeWarehousingData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.EventName != nil {
		// no validation rules for EventName
	}

	if m.ReportData != nil {
		// no validation rules for ReportData
	}

	if m.CreateTime != nil {
		// no validation rules for CreateTime
	}

	if len(errors) > 0 {
		return RealTimeWarehousingDataMultiError(errors)
	}

	return nil
}

// RealTimeWarehousingDataMultiError is an error wrapping multiple validation
// errors returned by RealTimeWarehousingData.ValidateAll() if the designated
// constraints aren't met.
type RealTimeWarehousingDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RealTimeWarehousingDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RealTimeWarehousingDataMultiError) AllErrors() []error { return m }

// RealTimeWarehousingDataValidationError is the validation error returned by
// RealTimeWarehousingData.Validate if the designated constraints aren't met.
type RealTimeWarehousingDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RealTimeWarehousingDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RealTimeWarehousingDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RealTimeWarehousingDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RealTimeWarehousingDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RealTimeWarehousingDataValidationError) ErrorName() string {
	return "RealTimeWarehousingDataValidationError"
}

// Error satisfies the builtin error interface
func (e RealTimeWarehousingDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRealTimeWarehousingData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RealTimeWarehousingDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RealTimeWarehousingDataValidationError{}

// Validate checks the field values on KafkaData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *KafkaData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KafkaData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in KafkaDataMultiError, or nil
// if none found.
func (m *KafkaData) ValidateAll() error {
	return m.validate(true)
}

func (m *KafkaData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.AppId != nil {
		// no validation rules for AppId
	}

	if m.DistinctId != nil {
		// no validation rules for DistinctId
	}

	if m.TableId != nil {
		// no validation rules for TableId
	}

	if m.Ip != nil {
		// no validation rules for Ip
	}

	if m.ReportType != nil {
		// no validation rules for ReportType
	}

	if m.Debug != nil {
		// no validation rules for Debug
	}

	if m.ReqData != nil {
		// no validation rules for ReqData
	}

	if m.ReportTime != nil {
		// no validation rules for ReportTime
	}

	if m.ConsumptionTime != nil {
		// no validation rules for ConsumptionTime
	}

	if m.EventName != nil {
		// no validation rules for EventName
	}

	if len(errors) > 0 {
		return KafkaDataMultiError(errors)
	}

	return nil
}

// KafkaDataMultiError is an error wrapping multiple validation errors returned
// by KafkaData.ValidateAll() if the designated constraints aren't met.
type KafkaDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KafkaDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KafkaDataMultiError) AllErrors() []error { return m }

// KafkaDataValidationError is the validation error returned by
// KafkaData.Validate if the designated constraints aren't met.
type KafkaDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KafkaDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KafkaDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KafkaDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KafkaDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KafkaDataValidationError) ErrorName() string { return "KafkaDataValidationError" }

// Error satisfies the builtin error interface
func (e KafkaDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKafkaData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KafkaDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KafkaDataValidationError{}

// Validate checks the field values on Report with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Report) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Report with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ReportMultiError, or nil if none found.
func (m *Report) ValidateAll() error {
	return m.validate(true)
}

func (m *Report) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return ReportMultiError(errors)
	}

	return nil
}

// ReportMultiError is an error wrapping multiple validation errors returned by
// Report.ValidateAll() if the designated constraints aren't met.
type ReportMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReportMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReportMultiError) AllErrors() []error { return m }

// ReportValidationError is the validation error returned by Report.Validate if
// the designated constraints aren't met.
type ReportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReportValidationError) ErrorName() string { return "ReportValidationError" }

// Error satisfies the builtin error interface
func (e ReportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReportValidationError{}

// Validate checks the field values on ListReportResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListReportResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReportResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReportResponseMultiError, or nil if none found.
func (m *ListReportResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReportResponse) validate(all bool) error {
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
					errors = append(errors, ListReportResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListReportResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListReportResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListReportResponseMultiError(errors)
	}

	return nil
}

// ListReportResponseMultiError is an error wrapping multiple validation errors
// returned by ListReportResponse.ValidateAll() if the designated constraints
// aren't met.
type ListReportResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReportResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReportResponseMultiError) AllErrors() []error { return m }

// ListReportResponseValidationError is the validation error returned by
// ListReportResponse.Validate if the designated constraints aren't met.
type ListReportResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReportResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReportResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReportResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReportResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReportResponseValidationError) ErrorName() string {
	return "ListReportResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListReportResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReportResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReportResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReportResponseValidationError{}

// Validate checks the field values on GetReportRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetReportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetReportRequestMultiError, or nil if none found.
func (m *GetReportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetReportRequestMultiError(errors)
	}

	return nil
}

// GetReportRequestMultiError is an error wrapping multiple validation errors
// returned by GetReportRequest.ValidateAll() if the designated constraints
// aren't met.
type GetReportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReportRequestMultiError) AllErrors() []error { return m }

// GetReportRequestValidationError is the validation error returned by
// GetReportRequest.Validate if the designated constraints aren't met.
type GetReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReportRequestValidationError) ErrorName() string { return "GetReportRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReportRequestValidationError{}

// Validate checks the field values on CreateReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateReportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateReportRequestMultiError, or nil if none found.
func (m *CreateReportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateReportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateReportRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateReportRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateReportRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return CreateReportRequestMultiError(errors)
	}

	return nil
}

// CreateReportRequestMultiError is an error wrapping multiple validation
// errors returned by CreateReportRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateReportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateReportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateReportRequestMultiError) AllErrors() []error { return m }

// CreateReportRequestValidationError is the validation error returned by
// CreateReportRequest.Validate if the designated constraints aren't met.
type CreateReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateReportRequestValidationError) ErrorName() string {
	return "CreateReportRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateReportRequestValidationError{}

// Validate checks the field values on UpdateReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateReportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateReportRequestMultiError, or nil if none found.
func (m *UpdateReportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateReportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateReportRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateReportRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateReportRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return UpdateReportRequestMultiError(errors)
	}

	return nil
}

// UpdateReportRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateReportRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateReportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateReportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateReportRequestMultiError) AllErrors() []error { return m }

// UpdateReportRequestValidationError is the validation error returned by
// UpdateReportRequest.Validate if the designated constraints aren't met.
type UpdateReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateReportRequestValidationError) ErrorName() string {
	return "UpdateReportRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateReportRequestValidationError{}

// Validate checks the field values on DeleteReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteReportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteReportRequestMultiError, or nil if none found.
func (m *DeleteReportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteReportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for OperatorId

	if len(errors) > 0 {
		return DeleteReportRequestMultiError(errors)
	}

	return nil
}

// DeleteReportRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteReportRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteReportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteReportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteReportRequestMultiError) AllErrors() []error { return m }

// DeleteReportRequestValidationError is the validation error returned by
// DeleteReportRequest.Validate if the designated constraints aren't met.
type DeleteReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteReportRequestValidationError) ErrorName() string {
	return "DeleteReportRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteReportRequestValidationError{}

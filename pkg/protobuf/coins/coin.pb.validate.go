// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/coin.proto

package coins

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

// Validate checks the field values on Coin with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Coin) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Coin with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CoinMultiError, or nil if none found.
func (m *Coin) ValidateAll() error {
	return m.validate(true)
}

func (m *Coin) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Code

	// no validation rules for Icon

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CoinValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CoinValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CoinValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CoinMultiError(errors)
	}

	return nil
}

// CoinMultiError is an error wrapping multiple validation errors returned by
// Coin.ValidateAll() if the designated constraints aren't met.
type CoinMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CoinMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CoinMultiError) AllErrors() []error { return m }

// CoinValidationError is the validation error returned by Coin.Validate if the
// designated constraints aren't met.
type CoinValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CoinValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CoinValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CoinValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CoinValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CoinValidationError) ErrorName() string { return "CoinValidationError" }

// Error satisfies the builtin error interface
func (e CoinValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCoin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CoinValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CoinValidationError{}

// Validate checks the field values on Info with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Info) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Info with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in InfoMultiError, or nil if none found.
func (m *Info) ValidateAll() error {
	return m.validate(true)
}

func (m *Info) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for IsActive

	// no validation rules for HasSmartContracts

	// no validation rules for Platform

	// no validation rules for DateStart

	// no validation rules for MaxSupply

	// no validation rules for KeyFeatures

	// no validation rules for Usage

	// no validation rules for Site

	// no validation rules for Chat

	if len(errors) > 0 {
		return InfoMultiError(errors)
	}

	return nil
}

// InfoMultiError is an error wrapping multiple validation errors returned by
// Info.ValidateAll() if the designated constraints aren't met.
type InfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InfoMultiError) AllErrors() []error { return m }

// InfoValidationError is the validation error returned by Info.Validate if the
// designated constraints aren't met.
type InfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoValidationError) ErrorName() string { return "InfoValidationError" }

// Error satisfies the builtin error interface
func (e InfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoValidationError{}

// Validate checks the field values on GetCoinsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCoinsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCoinsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCoinsRequestMultiError, or nil if none found.
func (m *GetCoinsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCoinsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	if len(errors) > 0 {
		return GetCoinsRequestMultiError(errors)
	}

	return nil
}

// GetCoinsRequestMultiError is an error wrapping multiple validation errors
// returned by GetCoinsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCoinsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCoinsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCoinsRequestMultiError) AllErrors() []error { return m }

// GetCoinsRequestValidationError is the validation error returned by
// GetCoinsRequest.Validate if the designated constraints aren't met.
type GetCoinsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCoinsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCoinsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCoinsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCoinsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCoinsRequestValidationError) ErrorName() string { return "GetCoinsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetCoinsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCoinsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCoinsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCoinsRequestValidationError{}

// Validate checks the field values on GetCoinsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCoinsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCoinsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCoinsResponseMultiError, or nil if none found.
func (m *GetCoinsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCoinsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCoins() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetCoinsResponseValidationError{
						field:  fmt.Sprintf("Coins[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetCoinsResponseValidationError{
						field:  fmt.Sprintf("Coins[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetCoinsResponseValidationError{
					field:  fmt.Sprintf("Coins[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetCoinsResponseMultiError(errors)
	}

	return nil
}

// GetCoinsResponseMultiError is an error wrapping multiple validation errors
// returned by GetCoinsResponse.ValidateAll() if the designated constraints
// aren't met.
type GetCoinsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCoinsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCoinsResponseMultiError) AllErrors() []error { return m }

// GetCoinsResponseValidationError is the validation error returned by
// GetCoinsResponse.Validate if the designated constraints aren't met.
type GetCoinsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCoinsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCoinsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCoinsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCoinsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCoinsResponseValidationError) ErrorName() string { return "GetCoinsResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetCoinsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCoinsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCoinsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCoinsResponseValidationError{}

// Validate checks the field values on GetCoinRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCoinRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCoinRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCoinRequestMultiError,
// or nil if none found.
func (m *GetCoinRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCoinRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetCoinRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetCoinRequestMultiError(errors)
	}

	return nil
}

// GetCoinRequestMultiError is an error wrapping multiple validation errors
// returned by GetCoinRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCoinRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCoinRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCoinRequestMultiError) AllErrors() []error { return m }

// GetCoinRequestValidationError is the validation error returned by
// GetCoinRequest.Validate if the designated constraints aren't met.
type GetCoinRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCoinRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCoinRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCoinRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCoinRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCoinRequestValidationError) ErrorName() string { return "GetCoinRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetCoinRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCoinRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCoinRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCoinRequestValidationError{}

// Validate checks the field values on GetCoinResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCoinResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCoinResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCoinResponseMultiError, or nil if none found.
func (m *GetCoinResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCoinResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCoin()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetCoinResponseValidationError{
					field:  "Coin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetCoinResponseValidationError{
					field:  "Coin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCoin()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCoinResponseValidationError{
				field:  "Coin",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetCoinResponseMultiError(errors)
	}

	return nil
}

// GetCoinResponseMultiError is an error wrapping multiple validation errors
// returned by GetCoinResponse.ValidateAll() if the designated constraints
// aren't met.
type GetCoinResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCoinResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCoinResponseMultiError) AllErrors() []error { return m }

// GetCoinResponseValidationError is the validation error returned by
// GetCoinResponse.Validate if the designated constraints aren't met.
type GetCoinResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCoinResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCoinResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCoinResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCoinResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCoinResponseValidationError) ErrorName() string { return "GetCoinResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetCoinResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCoinResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCoinResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCoinResponseValidationError{}

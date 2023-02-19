// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/listener/v3/udp_listener_config.proto

package listenerv3

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

// Validate checks the field values on UdpListenerConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UdpListenerConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UdpListenerConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UdpListenerConfigMultiError, or nil if none found.
func (m *UdpListenerConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *UdpListenerConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDownstreamSocketConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UdpListenerConfigValidationError{
					field:  "DownstreamSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UdpListenerConfigValidationError{
					field:  "DownstreamSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDownstreamSocketConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UdpListenerConfigValidationError{
				field:  "DownstreamSocketConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetQuicOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UdpListenerConfigValidationError{
					field:  "QuicOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UdpListenerConfigValidationError{
					field:  "QuicOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuicOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UdpListenerConfigValidationError{
				field:  "QuicOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUdpPacketPacketWriterConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UdpListenerConfigValidationError{
					field:  "UdpPacketPacketWriterConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UdpListenerConfigValidationError{
					field:  "UdpPacketPacketWriterConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUdpPacketPacketWriterConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UdpListenerConfigValidationError{
				field:  "UdpPacketPacketWriterConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UdpListenerConfigMultiError(errors)
	}
	return nil
}

// UdpListenerConfigMultiError is an error wrapping multiple validation errors
// returned by UdpListenerConfig.ValidateAll() if the designated constraints
// aren't met.
type UdpListenerConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UdpListenerConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UdpListenerConfigMultiError) AllErrors() []error { return m }

// UdpListenerConfigValidationError is the validation error returned by
// UdpListenerConfig.Validate if the designated constraints aren't met.
type UdpListenerConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpListenerConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpListenerConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpListenerConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpListenerConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpListenerConfigValidationError) ErrorName() string {
	return "UdpListenerConfigValidationError"
}

// Error satisfies the builtin error interface
func (e UdpListenerConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpListenerConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpListenerConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpListenerConfigValidationError{}

// Validate checks the field values on ActiveRawUdpListenerConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ActiveRawUdpListenerConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActiveRawUdpListenerConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ActiveRawUdpListenerConfigMultiError, or nil if none found.
func (m *ActiveRawUdpListenerConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ActiveRawUdpListenerConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ActiveRawUdpListenerConfigMultiError(errors)
	}
	return nil
}

// ActiveRawUdpListenerConfigMultiError is an error wrapping multiple
// validation errors returned by ActiveRawUdpListenerConfig.ValidateAll() if
// the designated constraints aren't met.
type ActiveRawUdpListenerConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActiveRawUdpListenerConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActiveRawUdpListenerConfigMultiError) AllErrors() []error { return m }

// ActiveRawUdpListenerConfigValidationError is the validation error returned
// by ActiveRawUdpListenerConfig.Validate if the designated constraints aren't met.
type ActiveRawUdpListenerConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveRawUdpListenerConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveRawUdpListenerConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveRawUdpListenerConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveRawUdpListenerConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveRawUdpListenerConfigValidationError) ErrorName() string {
	return "ActiveRawUdpListenerConfigValidationError"
}

// Error satisfies the builtin error interface
func (e ActiveRawUdpListenerConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveRawUdpListenerConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveRawUdpListenerConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveRawUdpListenerConfigValidationError{}

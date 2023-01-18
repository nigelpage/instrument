/**
 * Author: Nigel Page
 * File: error.go
 */

package instrument

import (
	"fmt"
	"strings"
	"time"
)

type ErrorLevel int16

const (
	Information ErrorLevel = iota
	Warning
	Error
	Fatal
)

type StructuredError struct {
	level       ErrorLevel    // error level
	Code        string        // error code
	description string        // error description
	When        time.Time     // time error occured
	Values      []interface{} // any additional information
}

func NewStructuredError(level ErrorLevel, code, descr string, values []interface{}) *StructuredError {
	return &StructuredError{
		level:       level,
		Code:        strings.ToUpper(code),
		description: descr,
		When:        time.Now(),
		Values:      values,
	}
}

// Error() supports the standard error interface
func (se *StructuredError) Error() string {
	var l string // error level
	switch se.level {
	case Information:
		l = "INFORMATION"
	case Warning:
		l = "WARNING"
	case Error:
		l = "ERROR"
	case Fatal:
		l = "FATAL"
	}

	var v string = "" // error values
	if se.Values != nil {
		var sb strings.Builder
		sb.WriteString(" : ")
		f := true
		t := "%v"
		for _, s := range se.Values {
			if f {
				f = false
				t = ", %v"
			}
			sb.WriteString(fmt.Sprintf(t, s))
		}
		v = sb.String()
	}

	fs := "%s: %s at %s, " + se.description + "%v"

	return fmt.Sprintf(fs, l, se.Code, se.When.Format(time.RFC1123), v)
}

/*
Needs to be changed when Go 1.20 is released as signature changes to Unwrap() []error
*/
func (se *StructuredError) Unwrap() error {
	var sea error

	return sea
}

// IsErrorCode() checks the error code against a provided one
func (e *StructuredError) IsErrorCode(code string) bool {
	return e.Code == strings.ToUpper(code)
}

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

var Level = ErrorLevel(0).NONE()

func (ErrorLevel) NONE() ErrorLevel        { return ErrorLevel(0) }
func (ErrorLevel) INFORMATION() ErrorLevel { return ErrorLevel(1) }
func (ErrorLevel) WARNING() ErrorLevel     { return ErrorLevel(2) }
func (ErrorLevel) ERROR() ErrorLevel       { return ErrorLevel(3) }
func (ErrorLevel) FATAL() ErrorLevel       { return ErrorLevel(4) }

type StructuredError struct {
	level       ErrorLevel    // error level
	Code        string        // error code
	description string        // error description
	When        time.Time     // time error occured
	Values      []interface{} // any additional information
}

func NewStructuredError(level ErrorLevel, code, descr string) *StructuredError {
	return &StructuredError{
		level:       level,
		Code:        code,
		description: descr,
		When:        time.Now(),
	}
}

// Error() supports the standard error interface
func (se *StructuredError) Error() string {
	var l string // error level
	switch se.level {
	case Level.INFORMATION():
		l = "INFORMATION"
	case Level.WARNING():
		l = "WARNING"
	case Level.ERROR():
		l = "ERROR"
	case Level.FATAL():
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

	return fmt.Sprintf("%s: %s at %s, %s%v",
		l, se.Code, se.When.Format(time.RFC1123), se.description, v)
}

func (se *StructuredError) Unwrap() []StructuredError {
	var sea []StructuredError

	return sea
}

// IsErrorCode() checks the error code against a provided one
func (e *StructuredError) IsErrorCode(code string) bool {
	return e.Code == code
}

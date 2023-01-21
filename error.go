/**
 * Author: Nigel Page
 * File: error.go
 */

package instrument

import (
	"encoding/json"
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
	Level       ErrorLevel    // error level
	Code        string        // error code
	Description string        // error description
	When        time.Time     // time error occured
	Values      []interface{} // any additional information
}

func NewStructuredError(level ErrorLevel, code, descr string, values []interface{}) *StructuredError {
	return &StructuredError{
		Level:       level,
		Code:        strings.ToUpper(code),
		Description: descr,
		When:        time.Now(),
		Values:      values,
	}
}

// Error() supports the standard error interface
func (se *StructuredError) Error() string {
	var l string // error level
	switch se.Level {
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
			sb.WriteString(fmt.Sprintf(t, s))
			if f {
				f = false
				t = ", %v"
			}
		}
		v = sb.String()
	}

	fs := "%s: %s at %s, " + se.Description + "%s"

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

func (e *StructuredError) ToJson() (string, error) {
	b, err := json.Marshal(e)
	return string(b[:]), err // [:] converts from array to slice without copying!
}

func FromJson(s string) (*StructuredError, error) {
	var se StructuredError
	err := json.Unmarshal([]byte(s), &se)
	return &se, err
}

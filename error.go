/**
 * Author: Nigel Page
 * File: error.go
 */

package instrument

import (
	"errors"
	//"encoding/json"
	"fmt"
	"strings"
	"time"
)

/*
* The format of StructuredError is based on the OpenTelemetry log data-model
* https://opentelemetry.io/docs/reference/specification/logs/data-model/
 */

type Severity uint16

const (
	Trace = 1
	Trace2 = 2
	Trace3 = 3
	Trace4 = 4
	Debug = 5
	Debug2 = 6
	Debug3 = 7
	Debug4 = 8
	Info = 9
	Info2 = 10
	Info3 = 11
	Info4 = 12
	Warn = 13
	Warn2 = 14
	Warn3 = 15
	Warn4 = 16
	Error = 17
	Error2 = 18
	Error3 = 19
	Error4 = 20
	Fatal = 21
	Fatal2 = 22
	Fatal3 = 23
	Fatal4 = 24
)

var severityToStr = map[Severity]string{
	Trace: "Trace",
	Trace2: "Trace2",
	Trace3: "Trace3",
	Trace4: "Trace4",
	Debug: "Debug",
	Debug2: "Debug2",
	Debug3: "Debug3",
	Debug4: "Debug4",
	Info: "Info",
	Info2: "Info2",
	Info3: "Info3",
	Info4: "Info4",
	Warn: "Warn",
	Warn2: "Warn2",
	Warn3: "Warn3",
	Warn4: "Warn4",
	Error: "Error",
	Error2: "Error2",
	Error3: "Error3",
	Error4: "Error4",
	Fatal: "Fatal",
	Fatal2: "Fatal2",
	Fatal3: "Fatal3",
	Fatal4: "Fatal4",
}

var strToSeverity = map[string]Severity{
	`"Trace"`: Trace,
	`"Trace2"`: Trace2,
	`"Trace3"`: Trace3,
	`"Trace4"`: Trace4,
	`"Debug"`: Debug,
	`"Debug2"`: Debug2,
	`"Debug3"`: Debug3,
	`"Debug4"`: Debug4,
	`"Info"`: Info,
	`"Info2"`: Info2,
	`"Info3"`: Info3,
	`"Info4"`: Info4,
	`"Warn"`: Warn,
	`"Warn2"`: Warn2,
	`"Warn3"`: Warn3,
	`"Warn4"`: Warn4,
	`"Error"`: Error,
	`"Error2"`: Error2,
	`"Error3"`: Error3,
	`"Error4"`: Error4,
	`"Fatal"`: Fatal,
	`"Fatal2"`: Fatal2,
	`"Fatal3"`: Fatal3,
	`"Fatal4"`: Fatal4,
}

func (s Severity) String() string {
	return severityToStr[s]
}

type StructuredError struct {
	Severity   Severity               // error severity
	Code       string                 // error code
	Message    string                 // error message
	When       int64                  // time error occured (nanoseconds since Unix epoch - OpenTelemetry)
	Attributes map[string]interface{} // any additional information
	wrapped	   error				  // wrapped error
}

func NewStructuredError(severity Severity, code, msg string, attributes map[string]interface{}) (*StructuredError, error) {
	if severity < Trace || severity > Fatal4 {
		return nil, errors.New("invalid severity")
	}
	return &StructuredError{
		Severity:   severity,
		Code:       strings.ToUpper(code),
		Message:    msg,
		When:       time.Now().Unix(),
		Attributes: attributes,
	}, nil
	
}

// Error() supports the standard error interface
func (se *StructuredError) Error() string {
	var v string = "" // error values
	if se.Attributes != nil {
		var sb strings.Builder
		sb.WriteString(" : ")
		f := true
		t := "%s=%v"
		for k, v := range se.Attributes {
			sb.WriteString(fmt.Sprintf(t, k, v))
			if f {
				f = false
				t = ", " + t
			}
		}
		v = sb.String()
	}

	fs := "%s: %s at %s, " + se.Message + "%s"
	if se.wrapped != nil {
		fs = fs + "%n  "
		// TODO: deal with repeat unwrapping and formatting
		
	}

	return fmt.Sprintf(fs, se.Severity.String(), se.Code, time.Unix(se.When, 0), v)
}

// IsErrorCode() checks the error code against a provided one
func (e *StructuredError) IsErrorCode(ec string) bool {
	return e.Code == strings.ToUpper(ec)
}

/*
func (se *StructuredError) Unwrap() error {
	if se.wrapped == nil {
		return nil
	} else
	{
		use, we := fromJson(se.wrapped)
		if use != nil {
			return use
		} else {
			return we.Unwrap()
		}
	}
}

func (e *StructuredError) ToJson() (string, error) {
	b, err := json.Marshal(e)
	return string(b[:]), err // [:] converts from array to slice without copying!
}

func fromJson(e error) (*StructuredError, error) {
	var se StructuredError
	err := json.Unmarshal([]byte(e), &se)
	if err != nil {
		return nil, s.wrapped
	}
	return &se, nil
}
*/
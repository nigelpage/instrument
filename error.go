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

/*
* The format of StructuredError is based on the OpenTelemetry log data-model
* https://opentelemetry.io/docs/reference/specification/logs/data-model/
*/

type Severity int16

const (
	TRACE = 1
	DEBUG = 5
	INFO = 9
	WARN = 13
	ERROR = 17
	FATAL = 21
)

type StructuredError struct {
	Severity       Severity      // error severity
	Code        string        // error code
	Description string        // error description
	When        time.Time     // time error occured
	Values      []interface{} // any additional information
}

func NewStructuredError(severity Severity, code, descr string, values []interface{}) *StructuredError {
	return &StructuredError{
		Severity:       severity,
		Code:        strings.ToUpper(code),
		Description: descr,
		When:        time.Now(),
		Values:      values,
	}
}

// Error() supports the standard error interface
func (se *StructuredError) Error() string {
	var sev string // error severity
	switch se.Severity {
	case TRACE:
		sev = "TRACE"
	case DEBUG:
		sev = "DEBUG"
	case INFO:
		sev = "INFO"
	case WARN:
		sev = "WARN"
	case ERROR:
		sev = "ERROR"
	case FATAL:
		sev = "FATAL"
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

	return fmt.Sprintf(fs, sev, se.Code, se.When.Format(time.RFC1123), v)
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

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
	NOTSPECIFIED Severity = iota
	TRACE
	TRACE2
	TRACE3
	TRACE4
	DEBUG
	DEBUG2
	DEBUG3
	DEBUG4
	INFO
	INFO2
	INFO3
	INFO4
	WARN
	WARN2
	WARN3
	WARN4
	ERROR
	ERROR2
	ERROR3
	ERROR4
	FATAL
	FATAL2
	FATAL3
	FATAL4
)

type StructuredError struct {
	Severity   Severity               // error severity
	Code       string                 // error code
	Message    string                 // error message
	When       int64                  // time error occured (nanoseconds since Unix epoch - OpenTelemetry)
	Attributes map[string]interface{} // any additional information
}

func NewStructuredError(severity Severity, code, msg string, attributes map[string]interface{}) *StructuredError {
	return &StructuredError{
		Severity:   severity,
		Code:       strings.ToUpper(code),
		Message:    msg,
		When:       time.Now().Unix(),
		Attributes: attributes,
	}
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

	return fmt.Sprintf(fs, se.severityText(), se.Code, time.Unix(se.When, 0), v)
}

func (se *StructuredError) severityText() string {
	var sev string // error severity
	switch se.Severity {
	case NOTSPECIFIED:
		sev = "NOTSPECIFIED"
	case TRACE:
		sev = "TRACE"
	case TRACE2:
		sev = "TRACE2"
	case TRACE3:
		sev = "TRACE3"
	case TRACE4:
		sev = "TRACE4"
	case DEBUG:
		sev = "DEBUG"
	case DEBUG2:
		sev = "DEBUG2"
	case DEBUG3:
		sev = "DEBUG3"
	case DEBUG4:
		sev = "DEBUG4"
	case INFO:
		sev = "INFO"
	case INFO2:
		sev = "INFO2"
	case INFO3:
		sev = "INFO3"
	case INFO4:
		sev = "INFO4"
	case WARN:
		sev = "WARN"
	case WARN2:
		sev = "WARN2"
	case WARN3:
		sev = "WARN3"
	case WARN4:
		sev = "WARN4"
	case ERROR:
		sev = "ERROR"
	case ERROR2:
		sev = "ERROR2"
	case ERROR3:
		sev = "ERROR3"
	case ERROR4:
		sev = "ERROR4"
	case FATAL:
		sev = "FATAL"
	case FATAL2:
		sev = "FATAL2"
	case FATAL3:
		sev = "FATAL3"
	case FATAL4:
		sev = "FATAL4"
	}

	return sev
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

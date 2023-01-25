/**
 * Author: Nigel Page
 * File: error_test.go
 */

package instrument

import (
	"strings"
	"testing"
)

var s Severity = INFO
var c = "filecreated"
var d = "new file created"
var a = map[string]interface{}{
	"filename": "Readme.md",
	"lines":    42,
}

func TestNewStructuredError(t *testing.T) {
	var e = NewStructuredError(s, c, d, a)

	if e.Severity != s {
		t.Errorf("NewStructredError did not return expected error severity")
	}

	if !e.IsErrorCode(c) {
		t.Errorf("NewStructuredError did not return expected error code")
	}

}

func TestStructuredErrorIsErrorCode(t *testing.T) {
	var e = NewStructuredError(s, c, d, a)

	if e.Code != strings.ToUpper(c) {
		t.Errorf("StructuredError.IsErrorCode did not enforce uppercase error code")
	}
}

func TestError(t *testing.T) {
	var e = NewStructuredError(s, c, d, a)
	var m = e.Error()
	if m == "" {
		t.Errorf("StructuredError.Error() did not return a formatted string")
	}
}

func TestJson(t *testing.T) {
	se := NewStructuredError(s, c, d, a)
	j, e := se.ToJson()
	if e != nil {
		t.Errorf("StructuredErrorToJson failed to serialize StructuredError to Json")
	}

	sed, e := FromJson(j)
	if e != nil {
		t.Errorf("StructuredErrorFromJson failed to deserialize StructuredError from Json")
	}

	if sed.Severity != se.Severity || sed.Code != se.Code || sed.Description != se.Description || sed.When != se.When {
		t.Errorf("StructuredErrorFromJson failed to deserialize StructuredError correctly")
	}
}

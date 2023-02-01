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
var m = "new file created"
var a = map[string]interface{}{
	"filename": "Readme.md",
	"lines":    42,
}

func TestNewStructuredError(t *testing.T) {
	e, err := NewStructuredError(s, c, m, a, nil)
	if err != nil {
		t.Errorf("NewStructuredError did not validate severity correctly")
	}

	if e.Severity != s {
		t.Errorf("NewStructredError did not return expected error severity")
	}

	if !e.IsErrorCode(c) {
		t.Errorf("NewStructuredError did not return expected error code")
	}

}

func TestStructuredErrorIsErrorCode(t *testing.T) {
	e, _ := NewStructuredError(s, c, m, a, nil)

	if e.Code != strings.ToUpper(c) {
		t.Errorf("StructuredError.IsErrorCode did not enforce uppercase error code")
	}
}

func TestError(t *testing.T) {
	e, _ := NewStructuredError(s, c, m, a, nil)
	var m = e.Error()
	if m == "" {
		t.Errorf("StructuredError.Error() did not return a formatted string")
	}
}

func TestJson(t *testing.T) {
	se, _ := NewStructuredError(s, c, m, a, nil)
	j, e := se.ToJson()
	if e != nil {
		t.Errorf("StructuredErrorToJson failed to serialize StructuredError to Json")
	}

	sed, e := FromJson(j)
	if e != nil {
		t.Errorf("StructuredErrorFromJson failed to deserialize StructuredError from Json")
	}

	if sed.Severity != se.Severity || sed.Code != se.Code || sed.Message != se.Message || sed.When != se.When {
		t.Errorf("StructuredErrorFromJson failed to deserialize StructuredError correctly")
	}
}

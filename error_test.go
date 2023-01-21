/**
 * Author: Nigel Page
 * File: error_test.go
 */

package instrument

import (
	"strings"
	"testing"
)

var l ErrorLevel = Warning
var c = "filecreated"
var d = "new file created"
var v = []interface{}{
	"Readme.md",
	42,
}

func TestNewStructuredError(t *testing.T) {
	var e = NewStructuredError(l, c, d, v)

	if e.Level != l {
		t.Errorf("NewStructredError did not return expected error level")
	}

	if !e.IsErrorCode(c) {
		t.Errorf("NewStructuredError did not return expected error code")
	}

}

func TestStructuredErrorIsErrorCode(t *testing.T) {
	var e = NewStructuredError(l, c, d, v)

	if e.Code != strings.ToUpper(c) {
		t.Errorf("StructuredError.IsErrorCode did not enforce uppercase error code")
	}
}

func TestError(t *testing.T) {
	var e = NewStructuredError(l, c, d, v)
	var m = e.Error()
	if m == "" {
		t.Errorf("StructuredError.Error() did not return a formatted string")
	}
}

func TestJson(t *testing.T) {
	se := NewStructuredError(l, c, d, v)
	j, e := se.Serialize()
	if e != nil {
		t.Errorf("StructuredErrorSerialize failed to dehydrate StructuredError to Json")
	}

	sed, e := Deserialize(j)
	if e != nil {
		t.Errorf("StructuredErrorDesrialize failed to rehydrate StructuredError from Json")
	}

	if sed != se {
		t.Errorf("StructuredErrorDesrialize failed to rehydrate StructuredError correctly")
	}
}

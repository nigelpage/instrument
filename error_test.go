/**
 * Author: Nigel Page
 * File: error_test.go
 */

package instrument

import (
	"strings"
	"testing"
)

func TestNewStructuredError(t *testing.T) {
	var l ErrorLevel = Information
	var c = "filecreated"
	var e = NewStructuredError(Information, c, "new file created")

	if e.level != l {
		t.Errorf("NewStructredError did not return expected error level")
	}

	if !e.IsErrorCode(c) {
		t.Errorf("NewStructuredError did not return expected error code")
	}

	//fmt.Println(e)
}

func TestStructuredErrorIsErrorCode(t *testing.T) {
	var c = "filecreated"
	var e = NewStructuredError(Information, c, "new file created")

	if e.Code != strings.ToUpper(c) {
		t.Errorf("StructuredError.IsErrorCode did not enforce uppercase error code")
	}
}

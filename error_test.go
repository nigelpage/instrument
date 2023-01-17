/**
 * Author: Nigel Page
 * File: error_test.go
 */

package instrument

import (
	"fmt"
	"testing"
)

func TestNewStructuredError(t *testing.T) {
	var l = Information
	var c = "FILECREATED"
	var e = NewStructuredError(Information, c, "new file created")

	if e.level != l {
		t.Errorf("NewStructredError did not return expected error level")
	}

	if e.Code != c {
		t.Errorf("NewStructuredError did not return expected erro code")
	}

	fmt.Println(e)
}

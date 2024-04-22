package networkfirewall

import (
	"testing"
)

// TestList tests the ViewRule function
func TestList(t *testing.T) {
	s, err := ViewRule("test")
	if err != nil {
		t.Error(err)
	}
	t.Log(*s)
}

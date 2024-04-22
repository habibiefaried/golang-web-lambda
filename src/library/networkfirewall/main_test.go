package networkfirewall

import (
	"testing"
)

// TestList tests the ViewRule function
func TestList(t *testing.T) {
	s, token, err := ViewRule("test")
	if err != nil {
		t.Error(err)
	}
	t.Log(*s, *token)
}

func TestAddRule(t *testing.T) {
	token, err := AddRule("test", "facebook.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(*token)
}

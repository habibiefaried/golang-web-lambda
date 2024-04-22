package networkfirewall

import (
	"github.com/stretchr/testify/assert"
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
	domainsToBeAdded := []string{"facebook.com", "google.com", "yahoo.com"}

	for _, v := range domainsToBeAdded {
		token, err := AddRule("test", v)
		if err != nil {
			t.Error(err)
		}
		t.Log(*token)
	}

}

func TestDeleteRule(t *testing.T) {
	token, err := DeleteRule("test", "facebook.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(*token)
}

func TestWhitelisted(t *testing.T) {
	verdicts := map[string]bool{
		"facebook.com": false,
		"google.com":   true,
	}

	for k, v := range verdicts {
		ans, err := IsDomainWhitelisted("test", k)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, v, ans, "Expected result doesn't match")
	}
}

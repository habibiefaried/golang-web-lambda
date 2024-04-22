package networkfirewall

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDomainValid(t *testing.T) {
	verdicts := map[string]bool{
		"facebook.com":  true,
		"google.com":    true,
		"sid":           false,
		"aasabvavasdsa": false,
		"invalid-.com":  false,
		"in..valid.com": false,
		"toolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolong.com": true,
		"valid123.com":               true,
		"1.com":                      true,
		"example.subdomain123.org":   true,
		"sub.example.domain2013.org": true,
		"sub.example.2021.ok.com":    true,
	}

	for k, v := range verdicts {
		assert.Equal(t, v, isDomainValid(k), fmt.Sprintf("Expected result of %v doesn't match", k))
	}
}

func TestAddRule(t *testing.T) {
	domainsToBeAdded := []string{"facebook.com", "google.com", "yahoo.com"}

	for _, v := range domainsToBeAdded {
		token, err := AddRule(os.Getenv("RULEGROUPNAME"), v)
		if err != nil {
			t.Error(err)
		}
		t.Log(*token)
	}

}

func TestList(t *testing.T) {
	s, token, err := ViewRule(os.Getenv("RULEGROUPNAME"))
	if err != nil {
		t.Error(err)
	}
	t.Log(*s, *token)
}

func TestDeleteRule(t *testing.T) {
	token, err := DeleteRule(os.Getenv("RULEGROUPNAME"), "facebook.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(*token)
}

func TestWhitelisted(t *testing.T) {
	verdicts := map[string]bool{
		"facebook.com": false,
		"google.com":   true,
		"yahoo.com":    true,
	}

	for k, v := range verdicts {
		ans, err := IsDomainWhitelisted(os.Getenv("RULEGROUPNAME"), k)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, v, ans, "Expected result doesn't match")
	}

	_, _ = DeleteRule(os.Getenv("RULEGROUPNAME"), "google.com")
	_, _ = DeleteRule(os.Getenv("RULEGROUPNAME"), "yahoo.com")
}

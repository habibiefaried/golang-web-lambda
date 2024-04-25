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
	domainsToBeAdded := []RequestBody{
		&RequestBody{
			Domain: "facebook.com",
			Port:   "80",
		},
		&RequestBody{
			Domain: "google.com",
			Port:   "443",
		},
		&RequestBody{
			Domain: "yahoo.com",
			Port:   "80",
		},
	}

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
	token, err := DeleteRule(os.Getenv("RULEGROUPNAME"), &RequestBody{
		Domain: "facebook.com",
		Port:   "80",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(*token)
}

func TestWhitelisted(t *testing.T) {
	verdicts := make(map[RequestBody]bool)
	verdicts[RequestBody{Domain: "facebook.com", Port: "80"}] = false
	verdicts[RequestBody{Domain: "google.com", Port: "80"}] = false
	verdicts[RequestBody{Domain: "google.com", Port: "443"}] = true
	verdicts[RequestBody{Domain: "google.com", Port: "80"}] = true
	verdicts[RequestBody{Domain: "google.com", Port: "443"}] = false

	for k, v := range verdicts {
		ans, err := IsDomainWhitelisted(os.Getenv("RULEGROUPNAME"), k)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, v, ans, "Expected result doesn't match")
	}

	_, _ = DeleteRule(os.Getenv("RULEGROUPNAME"), &RequestBody{
		Domain: "google.com",
		Port:   "443",
	})
	_, _ = DeleteRule(os.Getenv("RULEGROUPNAME"), &RequestBody{
		Domain: "yahoo.com",
		Port:   "80",
	})
}

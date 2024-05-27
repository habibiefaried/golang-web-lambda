package networkfirewallv2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidDomain(t *testing.T) {
	testDomains := map[string]bool{
		"example.com":     true,
		"sub.example.com": true,
		"invalid_domain":  false,
		"example":         false,
	}
	for k, ans := range testDomains {
		assert.Equal(t, ans, IsValidDomain(k))
	}
}

func TestIP(t *testing.T) {
	testIPs := map[string]bool{
		"192.168.1.1":                    true,
		"255.255.255.255":                true,
		"invalidIP":                      false,
		"1234::5678:90ab:cdef:1234:5678": true,
	}
	for k, ans := range testIPs {
		assert.Equal(t, ans, IsValidIP(k))
	}
}

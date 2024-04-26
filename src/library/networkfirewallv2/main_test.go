package networkfirewallv2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessData(t *testing.T) {
	r := RequestBody{
		ID:  "id-1234",
		URL: "https://example.org/path/subpath",
	}

	err := r.Process()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, r.Port, "443")
	assert.Equal(t, r.Domain, "example.org")
	assert.Equal(t, r.IsTLS, true)
}

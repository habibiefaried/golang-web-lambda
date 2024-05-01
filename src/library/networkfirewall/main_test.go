package networkfirewallv2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestProcessData(t *testing.T) {
	verdicts := make(map[RequestBody]RequestBody)
	verdicts[RequestBody{ID: "id-1234", URL: "https://example.org/path/subpath"}] = RequestBody{Port: "443", Domain: "example.org", IsTLS: true}
	verdicts[RequestBody{ID: "id-1222", URL: "http://example.com/"}] = RequestBody{Port: "80", Domain: "example.com", IsTLS: false}
	verdicts[RequestBody{ID: "id-4325", URL: "http://google.com:1234"}] = RequestBody{Port: "1234", Domain: "google.com", IsTLS: false}
	verdicts[RequestBody{ID: "id-1231", URL: "https://es.aws.internal:9090"}] = RequestBody{Port: "9090", Domain: "es.aws.internal", IsTLS: true}

	for tc, ans := range verdicts {
		err := tc.Process()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, tc.Port, ans.Port)
		assert.Equal(t, tc.Domain, ans.Domain)
		assert.Equal(t, tc.IsTLS, ans.IsTLS)
	}

}

func TestManageRule1(t *testing.T) {
	oldRb := RequestBody{
		ID:  "1",
		URL: "https://example.org/subpath",
	}
	newRb := RequestBody{
		ID:  "1",
		URL: "https://example.com/subpath",
	}
	err := ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	assert.EqualError(t, err, "Old rule {ID:1 URL:https://example.org/subpath IsTLS:true Domain:example.org Port:443} is not found, cannot proceed\n")
}

// TestManageRule2 to test by adding https://google.com , update to http://google.org/a and delete it
func TestManageRule2(t *testing.T) {
	// Add rule
	oldRb := RequestBody{}
	newRb := RequestBody{
		ID:  "1",
		URL: "https://google.com/subpath",
	}

	err := ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	if err != nil {
		t.Error(err)
	}

	// Update
	oldRb = RequestBody{
		ID:  "1",
		URL: "https://google.com/subpath",
	}
	newRb = RequestBody{
		ID:  "1",
		URL: "http://google.org/a",
	}
	err = ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	if err != nil {
		t.Error(err)
	}

	// Delete old Rule
	oldRb = RequestBody{
		ID:  "1",
		URL: "http://google.org/a",
	}
	newRb = RequestBody{}
	err = ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	if err != nil {
		t.Error(err)
	}
}

// TestManageRule3 to test by adding https://google.com:9090/exactpath to merchant-1 and https://google.com:9090/exactpath to merchant-2
// Then clear those
func TestManageRule3(t *testing.T) {
	// Add rule
	oldRb := RequestBody{}
	newRb := RequestBody{
		ID:  "1",
		URL: "https://google.com:9090/exactpath",
	}

	err := ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	if err != nil {
		t.Error(err)
	}

	// Add rule
	oldRb = RequestBody{}
	newRb = RequestBody{
		ID:  "2",
		URL: "https://google.com:9090/exactpath",
	}

	err = ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	if err != nil {
		t.Error(err)
	}

	// Delete rule with wrong merchant ID
	oldRb = RequestBody{
		ID:  "5",
		URL: "https://google.com:9090/exactpath",
	}
	newRb = RequestBody{}
	err = ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	assert.EqualError(t, err, "Old rule {ID:5 URL:https://google.com:9090/exactpath IsTLS:true Domain:google.com Port:9090} is not found, cannot proceed\n")

	// Update rule with wrong merchant ID
	oldRb = RequestBody{
		ID:  "xx-566",
		URL: "https://google.com:9090/exactpath",
	}
	newRb = RequestBody{
		ID:  "xx-566",
		URL: "http://google.com:1234/test",
	}
	err = ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	assert.EqualError(t, err, "Old rule {ID:xx-566 URL:https://google.com:9090/exactpath IsTLS:true Domain:google.com Port:9090} is not found, cannot proceed\n")

	// Delete old Rule
	oldRb = RequestBody{
		ID:  "1",
		URL: "https://google.com:9090/exactpath",
	}
	newRb = RequestBody{}
	err = ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	if err != nil {
		t.Error(err)
	}

	// Delete old Rule
	oldRb = RequestBody{
		ID:  "2",
		URL: "https://google.com:9090/exactpath",
	}
	newRb = RequestBody{}
	err = ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, newRb)
	if err != nil {
		t.Error(err)
	}
}

func TestManageRule4(t *testing.T) {
	domain := "https:"
	oldRb := RequestBody{
		ID:  "1",
		URL: domain,
	}
	err := ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, RequestBody{})
	assert.EqualError(t, err, fmt.Sprintf("Not valid URL: %v\n", domain))
}

// TestManageRule5 will return "parameter missing" because of improper input of old rule
// and empty input for new rule param
func TestManageRule5(t *testing.T) {
	oldRb := RequestBody{
		ID:  "",
		URL: "https://example.org/a",
	}
	err := ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, RequestBody{})
	assert.EqualError(t, err, "Parameter needed is missing\n")
}

// TestManageRule6 will return error try to delete non existent rule only
func TestManageRule6(t *testing.T) {
	oldRb := RequestBody{
		ID:  "123",
		URL: "https://example.org",
	}
	err := ManageRule(os.Getenv("RULEGROUPNAME"), oldRb, RequestBody{})
	assert.EqualError(t, err, "Old rule {ID:123 URL:https://example.org IsTLS:true Domain:example.org Port:443} is not found, cannot proceed\n")
}

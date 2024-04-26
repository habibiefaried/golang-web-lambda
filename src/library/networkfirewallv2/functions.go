package networkfirewallv2

import (
	"fmt"
	"net/url"
)

// Process checks if all required fields are filled, then parse the URL into IsTLS, domain and port
func (r *RequestBody) Process() error {
	if r.ID == "" {
		return fmt.Errorf("ID must not be empty")
	}
	if r.URL == "" {
		return fmt.Errorf("URL must not be empty")
	}

	urlParsed, err := url.Parse(r.URL)
	if err != nil {
		return fmt.Errorf("Error parsing URL: %v", err)
	}

	// Check the scheme to determine if TLS is used
	r.IsTLS = urlParsed.Scheme == "https"

	// Extract the domain
	r.Domain = urlParsed.Hostname()

	// Determine the port; use default ports if not specified
	if urlParsed.Port() == "" {
		if r.IsTLS {
			r.Port = "443" // Default HTTPS port
		} else {
			r.Port = "80" // Default HTTP port
		}
	} else {
		r.Port = urlParsed.Port()
	}

	return nil
}

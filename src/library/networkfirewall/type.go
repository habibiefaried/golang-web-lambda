package networkfirewall

import (
	"fmt"
)

type RequestBody struct {
	Domain string `json:"domain"`
	Port   string `json:"port"`
}

// Validate checks if all required fields are filled.
func (r *RequestBody) Validate() error {
	if r.Domain == "" {
		return fmt.Errorf("Domain must not be empty")
	}
	if r.Port == "" {
		return fmt.Errorf("Port must not be empty")
	}
	return nil
}

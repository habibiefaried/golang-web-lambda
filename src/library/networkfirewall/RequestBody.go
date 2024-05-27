package networkfirewallv2

import (
	"fmt"
	"net/url"
)

// RequestBody will receive the ID and URL that needs to be whitelisted, so then, this must be processed
type RequestBody struct {
	// Data coming from http/kafka request
	ID  string `json:"id"`
	URL string `json:"url"`

	// After process
	IsTLS    bool
	Domain   string
	Port     string
	IsIPAddr bool
}

// Process parse the URL into IsTLS, domain and port
func (r *RequestBody) Process() error {
	if r.IsEmpty() {
		return nil
	}

	urlParsed, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return fmt.Errorf("Error parsing URL: %v", err)
	}

	if IsValidDomain(urlParsed.Hostname()) {
		r.IsIPAddr = false
	} else if IsValidIP(urlParsed.Hostname()) {
		r.IsIPAddr = true
	} else {
		return fmt.Errorf("Not valid URL: %v\n", r.URL)
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

// IsEmpty validates both ID and URL, if any of them empty, then this may not a valid request
// or valid request to only whitelist/add a new domain
// anyway, this will be considered as empty if it is not complete
func (r *RequestBody) IsEmpty() bool {
	return r.ID == "" || r.URL == ""
}

// generatePartSuricataRule is a function to generate Suricata Rule string after $HOME_NET and before sid
// this is used for verification and searching purpose
func (r *RequestBody) generatePartSuricataRule() string {
	if r.IsIPAddr {
		return fmt.Sprintf(`$HOME_NET any -> %v %v (msg:"ID %v";`, r.Domain, r.Port, r.ID)
	} else {
		if r.IsTLS {
			return fmt.Sprintf(`$HOME_NET any -> any %v (tls.sni; content:"%v"; endswith; msg:"ID %v";`, r.Port, r.Domain, r.ID)
		} else {
			return fmt.Sprintf(`$HOME_NET any -> any %v (http.host; content:"%v"; endswith; msg:"ID %v";`, r.Port, r.Domain, r.ID)
		}
	}
}

// generateWholeSuricataRule is a function to generate whole rule to whitelist new domain
func (r *RequestBody) generateWholeSuricataRule(RuleNumber int) string {
	ret := ""
	if r.IsIPAddr {
		ret = "\n" + fmt.Sprintf(`alert tcp $HOME_NET any -> %v %v (msg:"ID %v"; sid:%v;) `, r.Domain, r.Port, r.ID, 300000+RuleNumber) + "\n"
		ret = ret + fmt.Sprintf(`pass tcp $HOME_NET any -> %v %v (msg:"ID %v"; sid:%v;)`, r.Domain, r.Port, r.ID, 600000+RuleNumber)
	} else { // Domain
		if r.IsTLS {
			ret = "\n" + fmt.Sprintf(`alert tls $HOME_NET any -> any %v (tls.sni; content:"%v"; endswith; msg:"ID %v"; sid:%v;) `, r.Port, r.Domain, r.ID, 300000+RuleNumber) + "\n"
			ret = ret + fmt.Sprintf(`pass tls $HOME_NET any -> any %v (tls.sni; content:"%v"; endswith; msg:"ID %v"; sid:%v;)`, r.Port, r.Domain, r.ID, 600000+RuleNumber)
		} else {
			ret = "\n" + fmt.Sprintf(`alert http $HOME_NET any -> any %v (http.host; content:"%v"; endswith; msg:"ID %v"; sid:%v;) `, r.Port, r.Domain, r.ID, 300000+RuleNumber) + "\n"
			ret = ret + fmt.Sprintf(`pass http $HOME_NET any -> any %v (http.host; content:"%v"; endswith; msg:"ID %v"; sid:%v;)`, r.Port, r.Domain, r.ID, 600000+RuleNumber)
		}
	}

	return ret
}

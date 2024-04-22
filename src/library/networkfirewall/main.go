package networkfirewall

import (
	"context"
	"fmt"
	nf "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	ssm "github.com/habibiefaried/golang-web-lambda/library/ssmparam"
	"os"
	"strings"
)

func AddRule(rulegroupname string, domain string) (*string, error) {
	counterSSMParam := os.Getenv("COUNTERSSMPATH")

	if !isDomainValid(domain) {
		return nil, fmt.Errorf("Domain '%v' is invalid", domain)
	}

	c, err := awsAuth()
	if err != nil {
		return nil, err
	}

	rules, token, err := ViewRule(rulegroupname)
	if err != nil {
		return nil, err
	}

	RuleNumber, err := ssm.GetCounter(counterSSMParam)
	if err != nil {
		return nil, err
	}

	inputrule := *rules + "\n" + fmt.Sprintf(`alert tls $HOME_NET any -> any 443 (tls.sni; content:"%v"; endswith; msg:"Matching TLS allowlisted FQDNs"; sid:%v;) `, domain, 30000+RuleNumber) + "\n"
	inputrule = inputrule + fmt.Sprintf(`pass tls $HOME_NET any -> any 443 (tls.sni; content:"%v"; endswith; sid:%v;)`, domain, 40000+RuleNumber) + "\n"
	updateRGoutput, err := updaterulegroupint(c, rulegroupname, inputrule, token)
	if err != nil {
		return nil, err
	}

	err = ssm.IncreaseCounter(counterSSMParam)
	if err != nil {
		return nil, err
	}

	return updateRGoutput.UpdateToken, nil
}

func ViewRule(rulegroupname string) (*string, *string, error) {
	c, err := awsAuth()
	if err != nil {
		return nil, nil, err
	}

	describeRuleOutput, err := c.DescribeRuleGroup(context.Background(), &nf.DescribeRuleGroupInput{
		AnalyzeRuleGroup: false,
		RuleGroupName:    &rulegroupname,
		Type:             types.RuleGroupTypeStateful,
	})

	if err != nil {
		return nil, nil, err
	}

	return describeRuleOutput.RuleGroup.RulesSource.RulesString, describeRuleOutput.UpdateToken, nil
}

func DeleteRule(rulegroupname string, domain string) (*string, error) {
	if !isDomainValid(domain) {
		return nil, fmt.Errorf("Domain '%v' is invalid", domain)
	}

	c, err := awsAuth()
	if err != nil {
		return nil, err
	}

	rules, token, err := ViewRule(rulegroupname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(*rules, "\n")
	var filteredLines []string

	for _, line := range lines {
		if !strings.Contains(line, domain) {
			filteredLines = append(filteredLines, line)
		}
	}

	updateRGoutput, err := updaterulegroupint(c, rulegroupname, strings.Join(filteredLines, "\n"), token)
	if err != nil {
		return nil, err
	}
	return updateRGoutput.UpdateToken, nil
}

func IsDomainWhitelisted(rulegroupname string, domain string) (bool, error) {
	if !isDomainValid(domain) {
		return false, fmt.Errorf("Domain '%v' is invalid", domain)
	}

	rules, _, err := ViewRule(rulegroupname)
	if err != nil {
		return false, err
	}

	return strings.Contains(*rules, domain), nil
}

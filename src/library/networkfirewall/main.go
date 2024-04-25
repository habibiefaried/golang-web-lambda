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

func AddRule(rulegroupname string, rb RequestBody) (*string, error) {
	counterSSMParam := os.Getenv("COUNTERSSMPATH")

	if !isDomainValid(rb.Domain) {
		return nil, fmt.Errorf("Domain '%v' is invalid", rb.Domain)
	}

	c, err := awsAuth()
	if err != nil {
		return nil, err
	}

	rules, token, err := ViewRule(rulegroupname)
	if err != nil {
		return nil, err
	}

	if isRuleExist(*rules, rb) {
		return nil, fmt.Errorf("Duplicated entry of domain '%v' and port '%v'", rb.Domain, rb.Port)
	}

	RuleNumber, err := ssm.GetCounter(counterSSMParam)
	if err != nil {
		return nil, err
	}

	inputrule := *rules + "\n" + fmt.Sprintf(`alert tls $HOME_NET any -> any %v (tls.sni; content:"%v"; endswith; msg:"Matching TLS allowlisted FQDNs"; sid:%v;) `, rb.Port, rb.Domain, 300000+RuleNumber) + "\n"
	inputrule = inputrule + fmt.Sprintf(`pass tls $HOME_NET any -> any %v (tls.sni; content:"%v"; endswith; sid:%v;)`, rb.Port, rb.Domain, 600000+RuleNumber)
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

func DeleteRule(rulegroupname string, rb RequestBody) (*string, error) {
	if !isDomainValid(rb.Domain) {
		return nil, fmt.Errorf("Domain '%v' is invalid", rb.Domain)
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
		if !isRuleExist(line, rb) {
			filteredLines = append(filteredLines, line)
		}
	}

	updateRGoutput, err := updaterulegroupint(c, rulegroupname, strings.Join(filteredLines, "\n"), token)
	if err != nil {
		return nil, err
	}
	return updateRGoutput.UpdateToken, nil
}

func IsDomainWhitelisted(rulegroupname string, rb RequestBody) (bool, error) {
	if !isDomainValid(rb.Domain) {
		return false, fmt.Errorf("Domain '%v' is invalid", rb.Domain)
	}

	rules, _, err := ViewRule(rulegroupname)
	if err != nil {
		return false, err
	}

	return isRuleExist(*rules, rb), nil
}

func isRuleExist(rules string, rb RequestBody) bool {
	return strings.Contains(rules, fmt.Sprintf(`any %v (tls.sni; content:"%v";`, rb.Port, rb.Domain))
}

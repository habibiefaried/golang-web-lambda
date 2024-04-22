package networkfirewall

import (
	"context"
	"fmt"
	nf "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"strings"
)

func AddRule(rulegroupname string, domain string) (*string, error) {
	baseSIDalert := 30000
	baseSIDpass := 40000

	c, err := awsAuth()
	if err != nil {
		return nil, err
	}

	rules, token, err := ViewRule(rulegroupname)
	if err != nil {
		return nil, err
	}

	LineOfRules := strings.Count(*rules, "\n") + 1

	baseSIDalert = baseSIDalert + LineOfRules
	baseSIDpass = baseSIDpass + LineOfRules

	inputrule := *rules + "\n" + fmt.Sprintf(`alert tls $HOME_NET any -> any 443 (tls.sni; content:"%v"; endswith; msg:"Matching TLS allowlisted FQDNs"; sid:%v;)`, domain, baseSIDalert) + "\n"
	inputrule = inputrule + fmt.Sprintf(`pass tls $HOME_NET any -> any 443 (tls.sni; content:"%v"; endswith; sid:%v;)`, domain, baseSIDpass)

	updateRGoutput, err := updaterulegroupint(c, rulegroupname, inputrule, token)

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
	rules, _, err := ViewRule(rulegroupname)
	if err != nil {
		return false, err
	}

	return strings.Contains(*rules, domain), nil
}

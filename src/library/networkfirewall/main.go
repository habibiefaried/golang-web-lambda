package networkfirewall

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	nf "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"strings"
)

func awsAuth() (*nf.Client, error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	// Create an Amazon Network Firewall service client
	client := nf.NewFromConfig(cfg)
	return client, nil
}

func AddRule(rulegroupname string, domain string) (*string, error) {
	baseSID := 10000

	c, err := awsAuth()
	if err != nil {
		return nil, err
	}

	str, token, err := ViewRule(rulegroupname)
	if err != nil {
		return nil, err
	}

	baseSID = baseSID + strings.Count((*str), "\n") + 1
	inputrule := (*str) + "\n" + fmt.Sprintf("pass tcp $HOME_NET any <> %v 443 (flow: not_established; sid:%v;)", domain, baseSID)

	updateRGoutput, err := c.UpdateRuleGroup(context.TODO(), &nf.UpdateRuleGroupInput{
		AnalyzeRuleGroup: false,
		RuleGroupName:    &rulegroupname,
		Rules:            &inputrule,
		UpdateToken:      token,
		Type:             types.RuleGroupTypeStateful,
	})

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

	describeRuleOutput, err := c.DescribeRuleGroup(context.TODO(), &nf.DescribeRuleGroupInput{
		AnalyzeRuleGroup: false,
		RuleGroupName:    &rulegroupname,
		Type:             types.RuleGroupTypeStateful,
	})

	if err != nil {
		return nil, nil, err
	}

	return describeRuleOutput.RuleGroup.RulesSource.RulesString, describeRuleOutput.UpdateToken, nil
}

func DeleteRule(domain string) error {
	return nil
}

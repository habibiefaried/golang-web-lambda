package networkfirewall

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	nf "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
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

func AddRule(domain string) error {
	return nil
}

func ViewRule(rulegroupname string) (*string, error) {
	c, err := awsAuth()
	if err != nil {
		return nil, err
	}

	describeRuleOutput, err := c.DescribeRuleGroup(context.TODO(), &nf.DescribeRuleGroupInput{
		AnalyzeRuleGroup: false,
		RuleGroupName:    &rulegroupname,
		Type:             types.RuleGroupTypeStateful,
	})

	if err != nil {
		return nil, err
	}

	return describeRuleOutput.RuleGroup.RulesSource.RulesString, nil
}

func DeleteRule(domain string) error {
	return nil
}

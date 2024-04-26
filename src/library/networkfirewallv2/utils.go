package networkfirewallv2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	nf "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"os"
	"regexp"
)

func awsAuth() (*nf.Client, error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	// Create an Amazon Network Firewall service client
	client := nf.NewFromConfig(cfg)
	return client, nil
}

func updaterulegroupint(c *nf.Client, rulegroupname string, inputrule string, token *string) (*nf.UpdateRuleGroupOutput, error) {
	IPSets := map[string]types.IPSet{}
	IPSets["HOME_NET"] = types.IPSet{
		Definition: []string{os.Getenv("HOME_NET")},
	}

	return c.UpdateRuleGroup(context.Background(), &nf.UpdateRuleGroupInput{
		RuleGroup: &types.RuleGroup{
			RulesSource: &types.RulesSource{
				RulesString: aws.String(inputrule),
			},
			RuleVariables: &types.RuleVariables{
				IPSets: IPSets,
			},
			StatefulRuleOptions: &types.StatefulRuleOptions{
				RuleOrder: types.RuleOrderStrictOrder,
			},
		},
		RuleGroupName: aws.String(rulegroupname),
		UpdateToken:   token,
		Type:          types.RuleGroupTypeStateful,
	})
}

func isRuleExist(rules string, rb RequestBody) bool {
	if rb.IsTLS {
		return strings.Contains(rules, fmt.Sprintf(`any %v (tls.sni; content:"%v"; msg:"%v";`, rb.Port, rb.Domain, rb.ID))
	} else {
		return strings.Contains(rules, fmt.Sprintf(`any %v (tls.sni; content:"%v"; msg:"%v";`, rb.Port, rb.Domain, rb.ID))
	}
	
}
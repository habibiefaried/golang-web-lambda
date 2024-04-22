package networkfirewall

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	nf "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"regexp"
	"os"
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

// isDomainValid is a function where it validates the string
func isDomainValid(domain string) bool {
	// Regex for basic domain validation
	pattern := `^(?i)(([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])\.){1,126}([a-z]|[a-z][a-z\-]*[a-z])$`

	// Compile the regex
	re := regexp.MustCompile(pattern)
	return re.MatchString(domain)
}

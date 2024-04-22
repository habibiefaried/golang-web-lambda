package networkfirewall

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	nf "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"strconv"
	"strings"
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
		Definition: []string{"10.0.0.0/24"},
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

func getLatestSID(inputrule string) int {
	// Splitting the string into lines
	lines := strings.Split(inputrule, "\n")

	// Getting the last line
	if len(lines) > 0 {
		lastLine := lines[len(lines)-1]
		if lastLine[:2] == "##" {
			num, err := strconv.Atoi(lastLine[3:])
			if err != nil {
				return 0
			}
			return num
		} else {
			return 0
		}
	} else {
		return 0
	}
}

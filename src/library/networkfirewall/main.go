package networkfirewallv2

import (
	"fmt"
	ssm "github.com/habibiefaried/golang-web-lambda/library/ssmparam"
	"os"
	"strings"
)

func ManageRule(rulegroupname string, oldRule RequestBody, newRule RequestBody) error {
	counterSSMParam := os.Getenv("COUNTERSSMPATH")

	c, err := awsAuth()
	if err != nil {
		return err
	}

	err = oldRule.Process()
	if err != nil {
		return err
	}

	err = newRule.Process()
	if err != nil {
		return err
	}

	rules, token, err := ViewRule(rulegroupname)
	if err != nil {
		return err
	}

	inputrule := ""

	RuleNumber, err := ssm.GetCounter(counterSSMParam)
	if err != nil {
		return err
	}

	if oldRule.IsEmpty() && newRule.IsEmpty() {
		return fmt.Errorf("Parameter needed is missing")
	} else {
		if !oldRule.IsEmpty() && !newRule.IsEmpty() { // if both filled
			if strings.Contains(*rules, oldRule.generatePartSuricataRule()) {
				inputrule = deleteRule(*rules, oldRule)
				inputrule = inputrule + newRule.generateWholeSuricataRule(RuleNumber)
			} else {
				return fmt.Errorf("Old rule %+v is not found, cannot proceed", oldRule)
			}
		} else if oldRule.IsEmpty() { // then this is whitelist only
			if !strings.Contains(*rules, newRule.generatePartSuricataRule()) {
				inputrule = *rules + newRule.generateWholeSuricataRule(RuleNumber)
			} else {
				return fmt.Errorf("New rule %+v already exists", newRule)
			}
		} else if newRule.IsEmpty() { // then this will delete the old rule
			if strings.Contains(*rules, oldRule.generatePartSuricataRule()) {
				inputrule = deleteRule(*rules, oldRule)
			} else {
				return fmt.Errorf("Old rule %+v is not found, cannot proceed", oldRule)
			}
		} else {
			return fmt.Errorf("To the case which should be impossible")
		}
	}

	_, err = updateRuleGroupInt(c, rulegroupname, inputrule, token)
	if err != nil {
		return err
	}

	err = ssm.IncreaseCounter(counterSSMParam)
	if err != nil {
		return err
	}

	return nil
}

func deleteRule(rules string, rb RequestBody) string {
	lines := strings.Split(rules, "\n")
	var filteredLines []string

	for _, line := range lines {
		if !strings.Contains(line, rb.generatePartSuricataRule()) { // if the rule exists
			filteredLines = append(filteredLines, line)
		}
	}

	return strings.Join(filteredLines, "\n")
}

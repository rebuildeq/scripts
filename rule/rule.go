package rule

import (
	"fmt"
	"os"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	start := time.Now()
	data, err := os.ReadFile("data/rule.yaml")
	if err != nil {
		return err
	}
	rule := RuleYaml{}
	err = yaml.Unmarshal(data, &rule)
	if err != nil {
		return fmt.Errorf("rule unmarshal: %w", err)
	}

	err = rule.sanitize()
	if err != nil {
		return fmt.Errorf("rule sanitize: %w", err)
	}

	err = generateRuleSQL(&rule)
	if err != nil {
		return fmt.Errorf("generateRuleSQL: %w", err)
	}

	fmt.Println("finished in", time.Since(start).String())
	return nil
}

func generateRuleSQL(sp *RuleYaml) error {
	w, err := os.Create("bin/rule_out.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	for _, rule := range sp.Rules {
		w.WriteString(fmt.Sprintf("UPDATE `rule_values` SET rule_value = '%s' WHERE rule_name = '%s';\n", rule.Value, rule.Name))
	}

	return nil
}

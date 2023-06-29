package rule

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-yaml/yaml"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Build(cmd *cobra.Command, args []string) error {
	start := time.Now()
	fmt.Printf("Rule...")
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
	}()
	data, err := os.ReadFile("rule.yaml")
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

	return nil
}

func generateRuleSQL(sp *RuleYaml) error {
	w, err := os.Create("rule.sql")
	if err != nil {
		return err
	}
	defer w.Close()

	for _, rule := range sp.Rules {
		w.WriteString(fmt.Sprintf("UPDATE `rule_values` SET rule_value = '%s' WHERE rule_name = '%s';\n", rule.Value, rule.Name))
	}

	return nil
}

// Import takes database info and dumps to yaml
func Import(cmd *cobra.Command, args []string) error {
	if !viper.IsSet("db_host") {
		return fmt.Errorf("db_host is not set, pass it as an argument --db_host=... or set it in .rebuildeq.yaml")
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&interpolateParams=true&collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8", viper.GetString("db_user"), viper.GetString("db_pass"), viper.GetString("db_host"), viper.GetString("db_name")))
	if err != nil {
		return fmt.Errorf("db connect: %w", err)
	}
	defer db.Close()

	rule := &RuleYaml{}
	rule.sanitize()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryxContext(ctx, "SELECT ruleset_id, rule_name, rule_value, notes FROM rule_values")
	if err != nil {
		return fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Rule{}
		err = rows.StructScan(&r)
		if err != nil {
			return fmt.Errorf("db struct scan: %w", err)
		}
		rule.Rules = append(rule.Rules, &r)
	}

	w, err := os.Create("rule_dump.yaml")
	if err != nil {
		return err
	}
	defer w.Close()

	enc := yaml.NewEncoder(w)
	err = enc.Encode(rule)
	if err != nil {
		return err
	}
	fmt.Println("Created rule_dump.yaml")
	return nil
}

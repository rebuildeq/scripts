package rule

import "fmt"

type RuleYaml struct {
	Rules []*struct {
		ID    int    `yaml:"id"`
		Name  string `yaml:"name"`
		Value string `yaml:"value"`
	} `yaml:"rules"`
}

func (e *RuleYaml) sanitize() error {
	for _, rule := range e.Rules {
		if rule.ID == 0 {
			rule.ID = 1
		}
		if rule.Name == "" {
			return fmt.Errorf("rule name must not be empty")
		}
		if rule.Value == "" {
			return fmt.Errorf("rule value must not be empty")
		}
	}
	return nil
}

package aa

import (
	"fmt"
	"os"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
)

func Build(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	start := time.Now()
	db := &dbReader{}
	fmt.Printf("AA...")
	var err error
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("AA changed", db.changedDBStrCount, "DBStr entries")
	}()
	err = generate(db)

	return nil
}

func generate(db *dbReader) error {
	data, err := os.ReadFile("aa.yaml")
	if err != nil {
		return err
	}
	aa := AAYaml{}
	err = yaml.Unmarshal(data, &aa)
	if err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	err = aa.sanitize()
	if err != nil {
		return fmt.Errorf("sanitize: %w", err)
	}

	err = generateAASQL(&aa)
	if err != nil {
		return fmt.Errorf("generateAASQL: %w", err)
	}

	err = modifyDBStr(db, &aa)
	if err != nil {
		return fmt.Errorf("modifyDBStr: %w", err)
	}

	err = generateWeb(&aa)
	if err != nil {
		return fmt.Errorf("generateWeb: %w", err)
	}
	return nil
}

func Import(cmd *cobra.Command, args []string) error {
	return nil
}

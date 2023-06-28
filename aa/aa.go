package aa

import (
	"fmt"
	"os"
	"time"

	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	start := time.Now()
	db := &dbReader{}
	fmt.Printf("AA...")
	defer func() {
		fmt.Println(" finished in", time.Since(start).String())
		fmt.Println("AA changed", db.changedDBStrCount, "DBStr entries")
	}()
	data, err := os.ReadFile("aa.yaml")
	if err != nil {
		return err
	}
	aa := AAYaml{}
	err = yaml.Unmarshal(data, &aa)
	if err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	err = generateAASQL(&aa)
	if err != nil {
		return fmt.Errorf("generateAASQL: %w", err)
	}

	err = modifyDBStr(db, &aa)
	if err != nil {
		return fmt.Errorf("modifyDBStr: %w", err)
	}
	return nil
}

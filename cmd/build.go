/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/xackery/rebuildeq/aa"
	"github.com/xackery/rebuildeq/rule"
	"github.com/xackery/rebuildeq/spell"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: buildRun,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func buildRun(cmd *cobra.Command, args []string) error {
	start := time.Now()
	if len(args) == 0 {
		fmt.Println("usage: rebuildeq build [rule|spell|aa|all]")
		os.Exit(1)
	}
	fmt.Println("Starting build with args", strings.Join(args, " "))
	defer func() {
		fmt.Println("Finished build in", time.Since(start).String())
	}()

	_, err := os.Stat("dbstr_us_original.txt")
	if err != nil {
		fmt.Println("Please copy dbstr_us.txt into this path, and rename it to dbstr_us_original.txt")
		os.Exit(1)
	}

	_, err = os.Stat("spells_us_original.txt")
	if err != nil {
		fmt.Println("Please copy spells_us.txt into this path, and rename it to spells_us_original.txt")
		os.Exit(1)
	}

	for _, arg := range args {
		switch arg {
		case "rule":
			err := rule.Build(cmd, args)
			if err != nil {
				return fmt.Errorf("rule: %w", err)
			}
		case "spell":
			err := spell.Build(cmd, args)
			if err != nil {
				return fmt.Errorf("spell: %w", err)
			}
		case "aa":
			err := aa.Build(cmd, args)
			if err != nil {
				return fmt.Errorf("aa: %w", err)
			}
		case "all":
			err := rule.Build(cmd, args)
			if err != nil {
				return fmt.Errorf("rule: %w", err)
			}
			err = spell.Build(cmd, args)
			if err != nil {
				return fmt.Errorf("spell: %w", err)
			}
			err = aa.Build(cmd, args)
			if err != nil {
				return fmt.Errorf("aa: %w", err)
			}
			return nil
		default:
			return fmt.Errorf("unknown build target: %s", arg)
		}
	}

	return nil
}

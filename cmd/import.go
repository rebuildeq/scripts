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
	"github.com/xackery/rebuildeq/charcreate"
	"github.com/xackery/rebuildeq/item"
	"github.com/xackery/rebuildeq/npc"
	"github.com/xackery/rebuildeq/rule"
	"github.com/xackery/rebuildeq/spell"
	"github.com/xackery/rebuildeq/task"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: importRun,
}

func init() {
	rootCmd.AddCommand(importCmd)
}

func importRun(cmd *cobra.Command, args []string) error {
	start := time.Now()
	if len(args) == 0 {
		fmt.Println("usage: rebuildeq import [aa|rule|spell|aa|task|charcreate|all]")
		os.Exit(1)
	}
	fmt.Println("Starting import with args", strings.Join(args, " "))
	defer func() {
		fmt.Println("Finished import in", time.Since(start).String())
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

	err = importProcess(cmd, args)
	if err != nil {
		fmt.Println("Failed import:", err)
		os.Exit(1)
	}

	return nil
}

func importProcess(cmd *cobra.Command, args []string) error {
	isRan := false
	for _, arg := range args {
		switch arg {
		case "rule":
			err := rule.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("rule: %w", err)
			}
		case "spell":
			err := spell.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("spell: %w", err)
			}
		case "aa":
			err := aa.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("aa: %w", err)
			}
		case "task":
			err := task.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("task: %w", err)
			}
		case "charcreate":
			err := charcreate.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("charcreate: %w", err)
			}
		case "npc":
			err := npc.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("npc: %w", err)
			}
		case "item":
			err := item.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("item: %w", err)
			}
		case "all":
			err := rule.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("rule: %w", err)
			}
			err = spell.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("spell: %w", err)
			}
			err = aa.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("aa: %w", err)
			}
			err = task.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("task: %w", err)
			}
			err = charcreate.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("charcreate: %w", err)
			}
			err = npc.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("npc: %w", err)
			}
			err = item.Import(cmd, args)
			if err != nil {
				return fmt.Errorf("item: %w", err)
			}
			return nil
		default:
			if !isRan {
				return fmt.Errorf("unknown import target: %s", arg)
			}
		}
		isRan = true
	}
	return nil
}

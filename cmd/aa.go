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

	"github.com/spf13/cobra"
)

// aaCmd represents the aa command
var aaCmd = &cobra.Command{
	Use:   "aa",
	Short: "Generate SQL/dbstr changes for AAs",
	Long:  `When you need an AA generated`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aa called")
	},
}

func init() {
	rootCmd.AddCommand(aaCmd)
}

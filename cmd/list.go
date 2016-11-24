// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"todo/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long: `List todo items`,
	Run: runCommand,
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func runCommand(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)

	if err != nil {
		log.Printf("%v", err)
	}
	fmt.Println(items)

}

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

	"github.com/spf13/cobra"

	"todo/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Long: `Add will save a new todo item and add it to the list`,
	Run: addRun,
}

var priority int

func init() {
	// called before main, use this for configuration
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 3, "Priority: 1-5")
}

func addRun(cmd *cobra.Command, args []string) {
	added := []todo.Item{}
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}
	
	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
		added = append(added, item)
	}
	err = todo.SaveItems(dataFile, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	fmt.Println("Added:", added)
}

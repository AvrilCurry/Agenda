// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/spf13/cobra"
)

// quitCmd represents the quit command
var quitCmd = &cobra.Command{
	Use:   "quit",
	Short: "Meeting quit",
	Long:  "Quit an existed meeting by title",
	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")

		var isTitleMissing = false

		if title == "Anonymous" {
			isTitleMissing = true
		}

		if isTitleMissing {
			log.Printf("Error: [Missing option \"--title\"] occur.\n")
			os.Exit(2)
		}

		if err == nil {
			// Todo Somethings
			fmt.Println("Agenda Command is \"quit\".\ncalled with:")
			fmt.Printf("\ttitle: %s\n", title)
		}
	},
}

func init() {
	rootCmd.AddCommand(quitCmd)

	quitCmd.Flags().StringP("title", "", "Anonymous", "title for quit")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

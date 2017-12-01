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

// deleteparCmd represents the deletepar command
var deleteparCmd = &cobra.Command{
	Use:   "deletepar",
	Short: "Delete participators from an existed meeting",
	Long:  "Delete participators from an existed meeting by title and participator",
	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")
		participator, err := cmd.Flags().GetStringSlice("participator")

		var isTitleMissing, isTitleMissingValue = false, false
		var isParticipatorMissing = false

		if title == "Anonymous" {
			isTitleMissing = true
		}
		if title[0] == '-' && title[1] == '-' {
			isTitleMissingValue = true
		}

		if len(participator) == 0 {
			isParticipatorMissing = true
		}

		if isTitleMissing {
			log.Printf("Error: [Missing option \"--title\"] occur.\n")
			os.Exit(2)
		} else if isTitleMissingValue {
			log.Printf("Error: [\"--title\" doesn't own an argument value] occur.\n")
			os.Exit(3)
		} else if isParticipatorMissing {
			/*
				使用warning，是因为这对结果不会产生影响，这在增删参与者是允许的，但如果是创建会议时是不可以的，
				因为至少得有一个人。
			*/
			log.Printf("Warning: [\"--participator\" own an empty argument value \"[]\"].\n")
			os.Exit(0)
		}

		if err == nil {
			// Todo Somethings
			fmt.Println("Agenda Command is \"deletepar\".\ncalled with:")
			fmt.Printf("\ttitle: %s\n\tparticipator: %s\n", title, participator)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteparCmd)

	deleteparCmd.Flags().StringP("title", "", "Anonymous", "Title for deletepar")
	deleteparCmd.Flags().StringSliceP("participator", "", nil, "Participator for deletepar")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteparCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteparCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

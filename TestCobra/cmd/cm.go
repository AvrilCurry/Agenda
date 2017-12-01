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
	"regexp"

	"github.com/spf13/cobra"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "Meeting Create",
	Long:  "Create a meeting by title, partcipators, startTime and endTime",
	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")
		participator, err := cmd.Flags().GetStringSlice("participator")
		startTime, err := cmd.Flags().GetString("startTime")
		endTime, err := cmd.Flags().GetString("endTime")

		var isTitleMissing, isTitleMissingValue = false, false
		var isParticipatorMissing, isParticipatorMissingValue = false, false
		var isStartTimeMissing, isStartTimeMissingValue = false, false
		var isEndTimeMissing = false

		fmt.Println(participator)

		if title == "Anonymous" {
			isTitleMissing = true
		}
		if title[0] == '-' && title[1] == '-' {
			isTitleMissingValue = true
		}

		if len(participator) == 0 {
			isParticipatorMissing = true
		}
		if len(participator) != 0 && participator[0][0] == '-' && participator[0][1] == '-' {
			isParticipatorMissingValue = true
		}

		if startTime == "Anonymous" {
			isStartTimeMissing = true
		}
		if startTime[0] == '-' && startTime[1] == '-' {
			isStartTimeMissingValue = true
		}

		if endTime == "Anonymous" {
			isEndTimeMissing = true
		}

		if isTitleMissing {
			log.Printf("Error: [Missing option\"--title\"] occur.\n")
			os.Exit(2)
		} else if isTitleMissingValue {
			log.Printf("Error: [\"--title\" doesn't own an argument value] occur.\n")
			os.Exit(3)
		} else if isParticipatorMissing {
			log.Printf("Error: [Missing option\"--participator\"].\n")
			os.Exit(4)
		} else if isParticipatorMissingValue {
			log.Printf("Error: [\"--participator\" doesn't own an argument value] occur.\n")
			os.Exit(5)
		} else if isStartTimeMissing {
			log.Printf("Error: [Missing option\"--startTime\"] occur.\n")
			os.Exit(6)
		} else if isStartTimeMissingValue {
			log.Printf("Error: [\"--startTime\" doesn't own an argument value] occur.\n")
			os.Exit(7)
		} else if isEndTimeMissing {
			log.Printf("Error: [Missing option\"--endTime\"] occur.\n")
			os.Exit(8)
		}

		// Regular Expression
		isStartTimeMatch, err := regexp.Match("([0-9|-])*", []byte(startTime))
		isEndTimeMatch, err := regexp.Match("([0-9|-])*", []byte(endTime))

		if !isStartTimeMatch {
			log.Printf("Error: [%s doesn't match the rules] occur.\n", startTime)
			configrules()
			os.Exit(9)
		}
		if !isEndTimeMatch {
			log.Printf("Error: [%s doesn't match the rules] occur.\n", endTime)
			configrules()
			os.Exit(10)
		}
		if startTime >= endTime {
			log.Printf("Error: [%s shouldn't larger than %s] occur.\n", startTime, endTime)
			os.Exit(11)
		}

		if err == nil {
			// Todo Somethings
			fmt.Println("Agenda Command is \"cm\".\ncalled with:")
			fmt.Printf("\ttitle: %s\n\tparticipator: %s\n\tstartTime: %s\n\tendTime: %s\n", title, participator, startTime, endTime)
		}
	},
}

func configrules() {
	fmt.Printf("--startTime should be made up of the number between 0-9 and character \"-\".\n")
	fmt.Printf("--endTime should be made up of the number between 0-9 and character \"-\".\n")
	fmt.Printf("--email should be made up of the pattern like \"xxxx@xxx.com\".\n")
	fmt.Printf("--telephone should begin as '1' and be made up of the number between 0-9 only have 11 numbers.\n")
}

func init() {
	rootCmd.AddCommand(cmCmd)

	cmCmd.Flags().StringP("title", "", "Anonymous", "Title for cm")
	cmCmd.Flags().StringSliceP("participator", "", nil, "Participator for cm")
	cmCmd.Flags().StringP("startTime", "", "Anonymous", "StartTime for cm")
	cmCmd.Flags().StringP("endTime", "", "Anonymous", "EndTime for cm")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

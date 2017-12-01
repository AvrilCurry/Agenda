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

// searchMeetingCmd represents the searchMeeting command
var searchMeetingCmd = &cobra.Command{
	Use:   "searchMeeting",
	Short: "Meeting search",
	Long:  "Search meetings during a period by startTime and endTime",
	Run: func(cmd *cobra.Command, args []string) {
		startTime, err := cmd.Flags().GetString("startTime")
		endTime, err := cmd.Flags().GetString("endTime")

		var isStartTimeMissing, isStartTimeMissingValue = false, false
		var isEndTimeMissing = false

		if startTime == "Anonymous" {
			isStartTimeMissing = true
		}
		if startTime[0] == '-' && startTime[1] == '-' {
			isStartTimeMissingValue = true
		}

		if endTime == "Anonymous" {
			isEndTimeMissing = true
		}

		if isStartTimeMissing {
			log.Printf("Error: [Missing option \"--startTime\"] occur.\n")
			os.Exit(2)
		} else if isStartTimeMissingValue {
			log.Printf("Error: [\"--startTime\" doesn't own an argument value] occur.\n")
			os.Exit(3)
		} else if isEndTimeMissing {
			log.Printf("Error: [Missing option \"--endTime\"] occur.\n")
			os.Exit(4)
		}

		// Regular Expression
		isStartTimeMatch, err := regexp.Match("([0-9|-])*", []byte(startTime))
		isEndTimeMatch, err := regexp.Match("([0-9|-])*", []byte(endTime))

		if !isStartTimeMatch {
			log.Printf("Error: [%s doesn't match the rules] occur.\n", startTime)
			configrules()
			os.Exit(5)
		}
		if !isEndTimeMatch {
			log.Printf("Error: [%s doesn't match the rules] occur.\n", endTime)
			configrules()
			os.Exit(6)
		}
		if startTime >= endTime {
			log.Printf("Error: [%s shouldn't larger than %s] occur.\n", startTime, endTime)
			os.Exit(7)
		}

		if err == nil {
			// Todo Somethings
			fmt.Println("Agenda Command is \"searchMeeting\".\ncalled with:")
			fmt.Printf("\tstartTime: %s\n\tendTime: %s\n", startTime, endTime)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchMeetingCmd)

	searchMeetingCmd.Flags().StringP("startTime", "", "Anonymous", "StartTime for searchMeeting")
	searchMeetingCmd.Flags().StringP("endTime", "", "Anonymous", "EndTime for searchMeeting")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

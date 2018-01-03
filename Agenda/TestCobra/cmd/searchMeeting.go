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
	"Agenda/TestCobra/entity"
	"Agenda/TestCobra/otherFunc"
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
		username, err := cmd.Flags().GetString("username")
		startTime, err := cmd.Flags().GetString("startTime")
		endTime, err := cmd.Flags().GetString("endTime")

		var isUsernameMissing, isUsernameMissingValue = false, false
		var isStartTimeMissing, isStartTimeMissingValue = false, false
		var isEndTimeMissing = false

		if username == "Anonymous" {
			isUsernameMissing = true
		}
		if username[0] == '-' && username[1] == '-' {
			isUsernameMissingValue = true
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

		fout, err := os.OpenFile("./log/error.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
		MyErrorLogger := log.New(fout, "[Error]: ", log.Ldate|log.Ltime)
		CommandInfo := "Running at agenda.go searchMeeting."

		if isUsernameMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-u/--username\"] occur.\n\n", CommandInfo)
			os.Exit(2)
		} else if isUsernameMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"-u/--username\" doesn't own an argument value] occur.\n", CommandInfo)
			os.Exit(3)
		} else if isStartTimeMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"--startTime\"] occur.\n\n", CommandInfo)
			os.Exit(4)
		} else if isStartTimeMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"--startTime\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(5)
		} else if isEndTimeMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"--endTime\"] occur.\n\n", CommandInfo)
			os.Exit(6)
		}

		// Regular Expression
		isStartTimeMatch, err := regexp.Match("[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}", []byte(startTime))
		isEndTimeMatch, err := regexp.Match("[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}", []byte(endTime))
		if !isStartTimeMatch {
			MyErrorLogger.Printf("%s\n\tError: [startTime %s doesn't match the rules] occur.\n\n", CommandInfo, startTime)
			otherFunc.Configrules()
			os.Exit(7)
		}
		if !isEndTimeMatch {
			MyErrorLogger.Printf("%s\n\tError: [endTime %s doesn't match the rules] occur.\n\n", CommandInfo, endTime)
			otherFunc.Configrules()
			os.Exit(8)
		}
		if startTime >= endTime {
			MyErrorLogger.Printf("%s\n\tError: [startTime %s shouldn't larger than endTime %s] occur.\n\n", CommandInfo, startTime, endTime)
			os.Exit(9)
		}

		if err == nil {
			MyCorrectLogger := log.New(fout, "[Correct]: ", log.Ldate|log.Ltime)
			MyWrongLogger := log.New(fout, "[Wrong]: ", log.Ldate|log.Ltime)
			outputInfo := "Agenda Command is \"searchMeeting\".\n\tcalled with:\n\t\tusername: %s\n\t\tstartTime: %s\n\t\tendTime: %s\n"

			value, _ := entity.MeetingSearch(username, startTime, endTime)
			switch value {
			case 1:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Search the Meeting! \"%s\" hasn't register yet!\n\n", username, startTime, endTime, username)
			case 2:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Search the Meeting! \"%s\" hasn't log in yet!\n\n", username, startTime, endTime, username)
			default:
				MyCorrectLogger.Printf(outputInfo+"\tOutput:\n\t\tSucceed to Search the Meeting!\n\t\t%v\n\n", username, startTime, endTime, value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchMeetingCmd)

	searchMeetingCmd.Flags().StringP("username", "u", "Anonymous", "Username for searchMeeting")
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

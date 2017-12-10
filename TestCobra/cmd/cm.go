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

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "Meeting Create",
	Long:  "Create a meeting by title, partcipators, startTime and endTime",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		title, err := cmd.Flags().GetString("title")
		participator, err := cmd.Flags().GetStringSlice("participator")
		startTime, err := cmd.Flags().GetString("startTime")
		endTime, err := cmd.Flags().GetString("endTime")

		var isUsernameMissing, isUsernameMissingValue = false, false
		var isTitleMissing, isTitleMissingValue = false, false
		var isParticipatorMissing, isParticipatorMissingValue = false, false
		var isStartTimeMissing, isStartTimeMissingValue = false, false
		var isEndTimeMissing = false

		//fmt.Println(participator)

		if username == "Anonymous" {
			isUsernameMissing = true
		}
		if username[0] == '-' && username[1] == '-' {
			isUsernameMissingValue = true
		}

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

		fout, err := os.OpenFile("./log/error.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
		MyErrorLogger := log.New(fout, "[Error]: ", log.Ldate|log.Ltime)
		CommandInfo := "Running at agenda.go cm."

		if isUsernameMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option\"-u/--username\"] occur.\n\n", CommandInfo)
			os.Exit(2)
		} else if isUsernameMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"-u/--username\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(3)
		} else if isTitleMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option\"--title\"] occur.\n\n", CommandInfo)
			os.Exit(4)
		} else if isTitleMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"--title\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(5)
		} else if isParticipatorMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option\"--participator\"].\n\n", CommandInfo)
			os.Exit(6)
		} else if isParticipatorMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"--participator\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(7)
		} else if isStartTimeMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option\"--startTime\"] occur.\n\n", CommandInfo)
			os.Exit(8)
		} else if isStartTimeMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"--startTime\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(9)
		} else if isEndTimeMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option\"--endTime\"] occur.\n\n", CommandInfo)
			os.Exit(10)
		}

		// Regular Expression
		isStartTimeMatch, err := regexp.Match("[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}", []byte(startTime))
		isEndTimeMatch, err := regexp.Match("[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}", []byte(endTime))

		if !isStartTimeMatch {
			MyErrorLogger.Printf("%s\n\tError: [startTime %s doesn't match the rules] occur.\n\n", CommandInfo, startTime)
			otherFunc.Configrules()
			os.Exit(11)
		}
		if !isEndTimeMatch {
			MyErrorLogger.Printf("%s\n\tError: [endTime %s doesn't match the rules] occur.\n\n", CommandInfo, endTime)
			otherFunc.Configrules()
			os.Exit(12)
		}
		if startTime >= endTime {
			MyErrorLogger.Printf("%s\n\tError: [startTime %s shouldn't larger than endTime %s] occur.\n\n", CommandInfo, startTime, endTime)
			os.Exit(13)
		}

		if err == nil {
			// Todo Somethings
			MyCorrectLogger := log.New(fout, "[Correct]: ", log.Ldate|log.Ltime)
			MyWrongLogger := log.New(fout, "[Wrong]: ", log.Ldate|log.Ltime)
			outputInfo := "Agenda Command is \"cm\".\n\tcalled with:\n\t\tusername: %s\n\t\ttitle: %s\n\t\tparticipator: %s\n\t\tstartTime: %s\n\t\tendTime: %s\n"

			res, value, _ := entity.CreateMeeting(username, title, participator, startTime, endTime)
			switch res {
			case 0:
				MyCorrectLogger.Printf(outputInfo+"\tOutput:\n\t\tUser \"%s\" Succeed to Create Meeting \"%s\"!\n\n", username, title, participator, startTime, endTime, username, title)
			case 1:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tUser \"%s\" Fail to Create Meeting! User hasn't register yet!\n\n", username, title, participator, startTime, endTime, username)
			case 2:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tUser \"%s\" Fail to Create Meeting! User hasn't log in yet!\n\n", username, title, participator, startTime, endTime, username)
			case 3:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tUser \"%s\" Fail to Create Meeting! Meeting \"%s\" has been Created!\n\n", username, title, participator, startTime, endTime, username, title)
			case 4:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tUser \"%s\" Fail to Create Meeting! %v hasn't register yet!\n\n", username, title, participator, startTime, endTime, username, value)
			case 5:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tUser \"%s\" Fail to Create Meeting! %v can't not participate the meeting during this period!\n\n", username, title, participator, startTime, endTime, username, value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)

	cmCmd.Flags().StringP("username", "u", "Anonymous", "Username for cm")
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

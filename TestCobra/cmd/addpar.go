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
	"log"
	"os"

	"github.com/spf13/cobra"
)

// addparCmd represents the addpar command
var addparCmd = &cobra.Command{
	Use:   "addpar",
	Short: "Add participator into an existed meeting",
	Long:  "Add participator into an existed meeting by title and participator",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		title, err := cmd.Flags().GetString("title")
		participator, err := cmd.Flags().GetStringSlice("participator")

		var isUsernameMissing, isUsernameMissingValue = false, false
		var isTitleMissing, isTitleMissingValue = false, false
		var isParticipatorMissing = false

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

		fout, err := os.OpenFile("./log/error.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
		MyErrorLogger := log.New(fout, "[Error]: ", log.Ldate|log.Ltime)
		CommandInfo := "Running at agenda.go addpar."

		if isUsernameMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-u/--username\"] occur.\n\n", CommandInfo)
			os.Exit(2)
		} else if isUsernameMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"-u/--username\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(3)
		} else if isTitleMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"--title\"] occur.\n\n", CommandInfo)
			os.Exit(4)
		} else if isTitleMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"--title\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(5)
		} else if isParticipatorMissing {
			/*
				使用warning，是因为这对结果不会产生影响，这在增删参与者是允许的，但如果是创建会议时是不可以的，
				因为至少得有一个人。
			*/
			MyWarningLogger := log.New(fout, "[Warning]: ", log.Ldate|log.Ltime)
			MyWarningLogger.Printf("%s\n\tWarning: [\"--participator\" own an empty argument value \"[]\"].\n\n", CommandInfo)
			os.Exit(0)
		}

		if err == nil {
			// Todo Somethings
			MyCorrectLogger := log.New(fout, "[Correct]: ", log.Ldate|log.Ltime)
			MyWrongLogger := log.New(fout, "[Wrong]: ", log.Ldate|log.Ltime)
			outputInfo := "Agenda Command is \"addpar\".\n\tcalled with:\n\t\tusername: %s\n\t\ttitle: %s\n\t\tparticipator: %s\n"

			res, value, _ := entity.AddParticipator(username, title, participator)
			switch res {
			case 0:
				MyCorrectLogger.Printf(outputInfo+"\tOutput:\n\t\tSucceed to Add Participator %v to the Meeting \"%s\"!\n\n", username, title, participator, participator, title)
			case 1:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Add Participator %v to the Meeting \"%s\"! \"%s\" hasn't register yet!\n\n", username, title, participator, participator, title, username)
			case 2:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Add Participator %v to the Meeting \"%s\"! \"%s\" hasn't log in yet!\n\n", username, title, participator, participator, title, username)
			case 3:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Add Participator %v to the Meeting! \"%s\" hasn't been Created  yet!\n\n", username, title, participator, participator, title)
			case 4:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Add Participator %v to the Meeting! \"%s\" has no right to add participator!\n\n", username, title, participator, participator, username)
			case 5:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Add Participator %v to the Meeting! Participator %v hasn't register yet!\n\n", username, title, participator, participator, value)
			case 6:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Add Participator %v to the Meeting! Participator %v can't not participate the meeting during this period!\n\n", username, title, participator, participator, value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addparCmd)

	addparCmd.Flags().StringP("username", "u", "Anonymous", "Username for addpar")
	addparCmd.Flags().StringP("title", "", "Anonymous", "Title for addpar")
	addparCmd.Flags().StringSliceP("participator", "", nil, "Participator for addpar")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addparCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addparCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Meeting cancel",
	Long:  "Cancel an existed meeting by title",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		title, err := cmd.Flags().GetString("title")

		var isUsernameMissing, isUsernameMissingValue = false, false
		var isTitleMissing = false

		if username == "Anonymous" {
			isUsernameMissing = true
		}
		if username[0] == '-' && username[1] == '-' {
			isUsernameMissingValue = true
		}

		if title == "Anonymous" {
			isTitleMissing = true
		}

		fout, err := os.OpenFile("./log/error.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
		MyErrorLogger := log.New(fout, "[Error]: ", log.Ldate|log.Ltime)
		CommandInfo := "Running at agenda.go cancel."

		if isUsernameMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-u/--username\"] occur.\n\n", CommandInfo)
			os.Exit(2)
		} else if isUsernameMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"-u/--username\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(3)
		} else if isTitleMissing {
			MyErrorLogger.Printf("%s\n\tError: Missing option \"--title\"] occur.\n\n", CommandInfo)
			os.Exit(4)
		}

		if err == nil {
			// Todo Somethings
			MyCorrectLogger := log.New(fout, "[Correct]: ", log.Ldate|log.Ltime)
			MyWrongLogger := log.New(fout, "[Wrong]: ", log.Ldate|log.Ltime)
			outputInfo := "Agenda Command is \"cancel\".\n\tcalled with:\n\t\tusername: %s\n\t\ttitle: %s\n"

			res, _ := entity.CancelMeeting(username, title)
			switch res {
			case 0:
				MyCorrectLogger.Printf(outputInfo+"\tOutput:\n\t\tSucceed to Cancel the Meeting \"%s\"!\n\n", username, title, title)
			case 1:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Cancel the Meeting! \"%s\" hasn't register yet!\n\n", username, title, username)
			case 2:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Cancel the Meeting! \"%s\" hasn't log in yet!\n\n", username, title, username)
			case 3:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Cancel the Meeting! \"%s\" hasn't created yet!\n\n", username, title, title)
			case 4:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Cancel the Meeting! \"%s\" has no right to cancel this meeting!\n\n", username, title, username)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cancelCmd)

	cancelCmd.Flags().StringP("username", "u", "Anonymous", "Username for cancel")
	cancelCmd.Flags().StringP("title", "", "Anonymous", "Title for cancel")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

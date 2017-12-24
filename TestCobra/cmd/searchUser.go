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

// searchCmd represents the search command
var searchUserCmd = &cobra.Command{
	Use:   "searchUser",
	Short: "User search",
	Long:  "Search user by username, e-mail or telephone",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		information, err := cmd.Flags().GetStringSlice("information")

		//var searchInfo []string
		var isUsernameMissing, isUsernameMissingValue = false, false
		/*var isInformationMissing, isUsername, isEmail, isTelephone = false, false, false, false
		var searchName, searchEmail, searchTelephone string*/

		if username == "Anonymous" {
			isUsernameMissing = true
		}
		if username != "Anonymous" && username[0] == '-' {
			isUsernameMissingValue = true
		}

		/*if len(information) == 0 {
			isInformationMissing = true
		}*/

		fout, err := os.OpenFile("./log/error.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
		MyErrorLogger := log.New(fout, "[Error]: ", log.Ldate|log.Ltime)
		CommandInfo := "Running at agenda.go searchUser."

		if isUsernameMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-u/--username\"] occur.\n\n", CommandInfo)
			os.Exit(2)
		} else if isUsernameMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"-u/--username\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(3)
		}

		/*Regular Expression*/
		/*isEmailMatch, err := regexp.Match("^([A-Za-z0-9_\\.-_]+)@([\\da-z]+)\\.com$", []byte(email))
		isTelephoneMatch, err := regexp.Match("^1[0-9]{10}$", []byte(telephone))

		if !isEmailMatch && !isEmailMissing {
			MyErrorLogger.Printf("%s\n\tError: [Email %s doesn't match the rules] occur.\n\n", CommandInfo, email)
			otherFunc.Configrules()
			os.Exit(4)
		}
		if !isTelephoneMatch && !isTelephoneMissing {
			MyErrorLogger.Printf("%s\n\tError: [Telephone %s doesn't match the rules] occur.\n\n", CommandInfo, telephone)
			otherFunc.Configrules()
			os.Exit(5)
		}*/

		if err == nil {
			// Todo Somethings
			MyCorrectLogger := log.New(fout, "[Correct]: ", log.Ldate|log.Ltime)
			MyWrongLogger := log.New(fout, "[Wrong]: ", log.Ldate|log.Ltime)
			MyValueLogger := log.New(fout, "", 0)
			outputInfo := "Agenda Command is \"search\".\n\tcalled with:\n\t\tusername: %s\n\t\tinformation: %s\n"

			res, value, _ := entity.UserSearch(username, information)
			switch res {
			case 0:
				MyCorrectLogger.Printf(outputInfo+"\tOutput:\n\t\tSucceed to Search!\n", username, information)
				for index, item := range value.([]map[string]string) {
					MyValueLogger.Printf("\t\t\tUser%d: %v\n", index, item)
				}
				MyValueLogger.Printf("\n")
			case 1:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to search! User \"%s\" hasn't register yet!\n\n", username, information, username)
			case 2:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to search! User \"%s\" hasn't log in yet!\n\n", username, information, username)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchUserCmd)

	searchUserCmd.Flags().StringP("username", "u", "Anonymous", "Username for searchUser")
	searchUserCmd.Flags().StringSliceP("information", "i", nil, "Information for searchUser")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

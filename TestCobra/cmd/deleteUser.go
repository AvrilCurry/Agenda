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

// deleteCmd represents the delete command
var deleteUserCmd = &cobra.Command{
	Use:   "deleteUser",
	Short: "User delete",
	Long:  "Delete an user by username",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")

		var isUsernameMissing = false

		if username == "Anonymous" {
			isUsernameMissing = true
		}

		fout, err := os.OpenFile("./log/error.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
		MyErrorLogger := log.New(fout, "[Error]: ", log.Ldate|log.Ltime)
		CommandInfo := "Running at agenda.go deleteUser."

		if isUsernameMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"--username\"] occur.\n\n", CommandInfo)
			os.Exit(2)
		}

		if err == nil {
			MyCorrectLogger := log.New(fout, "[Correct]: ", log.Ldate|log.Ltime)
			MyWrongLogger := log.New(fout, "[Wrong]: ", log.Ldate|log.Ltime)
			outputInfo := "Agenda Command is \"deleteUser\".\n\tcalled with:\n\t\tusername: %s\n"

			res, _ := entity.UserDelete(username)
			switch res {
			case 0:
				MyCorrectLogger.Printf(outputInfo+"\tOutput:\n\t\tSucceed to Delete User \"%s\"!\n\n", username, username)
			case 1:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Delete User \"%s\"! User hasn't register yet!\n\n", username, username)
			case 2:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Delete User \"%s\"! User hasn't log in yet!\n\n", username, username)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)

	deleteUserCmd.Flags().StringP("username", "u", "Anonymous", "Username for deleteUser")
	//deleteUserCmd.Flags().StringP("username", "u", "Anonymous", "Username for delete")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

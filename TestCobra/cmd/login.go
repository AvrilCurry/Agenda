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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "User log in",
	Long:  "Log in by username and password",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		password, err := cmd.Flags().GetString("password")

		var isUsernameMissing, isUsernameMissingValue = false, false
		var isPasswordMissing = false

		if username == "Anonymous" {
			isUsernameMissing = true
		}
		if username[0] == '-' {
			isUsernameMissingValue = true
		}

		if password == "Anonymous" {
			isPasswordMissing = true
		}

		fout, err := os.OpenFile("./log/error.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
		MyErrorLogger := log.New(fout, "[Error]: ", log.Ldate|log.Ltime)
		CommandInfo := "Running at agenda.go login."

		if isUsernameMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-u/--username\"] occur.\n\n", CommandInfo)
			os.Exit(2)
		} else if isUsernameMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"-u/--username\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(3)
		} else if isPasswordMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-p/--password\"] occur.\n\n", CommandInfo)
			os.Exit(4)
		}

		if err == nil {
			// Todo Somethings
			MyCorrectLogger := log.New(fout, "[Correct]: ", log.Ldate|log.Ltime)
			MyWrongLogger := log.New(fout, "[Wrong]: ", log.Ldate|log.Ltime)
			outputInfo := "Agenda Command is \"login\".\n\tcalled with:\n\t\tusername: %s\n\t\tpassword: %s\n"

			res, _ := entity.UserLogin(username, password)
			switch res {
			case 0:
				MyCorrectLogger.Printf(outputInfo+"\tOutput:\n\t\tSucceed to Log in \"%s\"!\n\n", username, password, username)
			case 1:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Login in \"%s\"! User doesn't exist!\n\n", username, password, username)
			case 2:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Login in \"%s\"! The password doesn't match the username!\n\n", username, password, username)
			case 3:
				MyWrongLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Login in \"%s\"! User has already log in!\n\n", username, password, username)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "Anonymous", "Username for login")
	loginCmd.Flags().StringP("password", "p", "Anonymous", "Password for register")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

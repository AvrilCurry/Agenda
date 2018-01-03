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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "User register",
	Long:  `Register an user by username, password, e-mail and phone information`,
	Run: func(cmd *cobra.Command, args []string) {

		username, err := cmd.Flags().GetString("username")
		password, err := cmd.Flags().GetString("password")
		email, err := cmd.Flags().GetString("email")
		telephone, err := cmd.Flags().GetString("telephone")

		var isUsernameMissing, isUsernameMissingValue = false, false
		var isPasswordMissing, isPasswordMissingValue = false, false
		var isEmailMissing, isEmailMissingValue = false, false
		var isTelephoneMissing = false

		if username == "Anonymous" {
			isUsernameMissing = true
		}
		if username[0] == '-' && username[1] == '-' {
			isUsernameMissingValue = true
		}

		if password == "Anonymous" {
			isPasswordMissing = true
		}
		if password[0] == '-' {
			isPasswordMissingValue = true
		}

		if email == "Anonymous" {
			isEmailMissing = true
		}
		if email[0] == '-' {
			isEmailMissingValue = true
		}

		if telephone == "Anonymous" {
			isTelephoneMissing = true
		}

		fout, err := os.OpenFile("./log/error.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
		MyErrorLogger := log.New(fout, "[Error]: ", log.Ldate|log.Ltime)
		CommandInfo := "Running at agenda.go register."

		if isUsernameMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-u/--username\"] occur.\n\n", CommandInfo)
			os.Exit(2)
		} else if isUsernameMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"-u/--username\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(3)
		} else if isPasswordMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-p/--password\"] occur.\n\n", CommandInfo)
			os.Exit(4)
		} else if isPasswordMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"-p/--password\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(5)
		} else if isEmailMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-e/--email\"] occur.\n\n", CommandInfo)
			os.Exit(6)
		} else if isEmailMissingValue {
			MyErrorLogger.Printf("%s\n\tError: [\"-e/--email\" doesn't own an argument value] occur.\n\n", CommandInfo)
			os.Exit(7)
		} else if isTelephoneMissing {
			MyErrorLogger.Printf("%s\n\tError: [Missing option \"-t/--telephone\"] occur.\n\n", CommandInfo)
			os.Exit(8)
		}

		/*Regular Expression*/
		isEmailMatch, err := regexp.Match("^([A-Za-z0-9_\\.-_]+)@([\\da-z]+)\\.com$", []byte(email))
		isTelephoneMatch, err := regexp.Match("^1[0-9]{10}$", []byte(telephone))

		if !isEmailMatch {
			MyErrorLogger.Printf("%s\n\tError: [Email %s doesn't match the rules] occur.\n\n", CommandInfo, email)
			otherFunc.Configrules()
			os.Exit(9)
		}
		if !isTelephoneMatch {
			MyErrorLogger.Printf("%s\n\tError: [Telephone %s doesn't match the rules] occur.\n\n", CommandInfo, telephone)
			otherFunc.Configrules()
			os.Exit(10)
		}

		if err == nil {
			// Todo Somethings

			MyCorrectLogger := log.New(fout, "[Correct]: ", log.Ldate|log.Ltime)
			outputInfo := "Agenda Command is \"register\".\n\tcalled with:\n\t\tusername: %s\n\t\tpassword: %s\n\t\te-mail: %s\n\t\ttelephone: %s\n"
			res, _ := entity.UserRegister(username, password, email, telephone)
			switch res {
			case 0:
				MyCorrectLogger.Printf(outputInfo+"\tOutput:\n\t\tSucceed to Create User \"%s\"!\n\n", username, password, email, telephone, username)
			case 1:
				MyCorrectLogger.Printf(outputInfo+"\tOutput:\n\t\tFail to Create User! User \"%s\" already exists!\n\n", username, password, email, telephone, username)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("username", "u", "Anonymous", "Name for register")
	registerCmd.Flags().StringP("password", "p", "Anonymous", "Password for register")
	registerCmd.Flags().StringP("email", "e", "Anonymous", "E-mail for register")
	registerCmd.Flags().StringP("telephone", "t", "Anonymous", "Telephone for register")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

		if isUsernameMissing {
			log.Printf("Error: [Missing option \"-u/--username\"] occur.\n")
			os.Exit(2)
		} else if isUsernameMissingValue {
			log.Printf("Error: [\"-u/--username\" doesn't own an argument value] occur.\n")
			os.Exit(3)
		} else if isPasswordMissing {
			log.Printf("Error: [Missing option \"-p/--password\"] occur.\n")
			os.Exit(4)
		} else if isPasswordMissingValue {
			log.Printf("Error: [\"-p/--password\" doesn't own an argument value] occur.\n")
			os.Exit(5)
		} else if isEmailMissing {
			log.Printf("Error: [Missing option \"-e/--email\"] occur.\n")
			os.Exit(6)
		} else if isEmailMissingValue {
			log.Printf("Error: [\"-e/--email\" doesn't own an argument value] occur.\n")
			os.Exit(7)
		} else if isTelephoneMissing {
			log.Printf("Error: [Missing option \"-t/--telephone\"] occur.\n")
			os.Exit(8)
		}

		/*Regular Expression*/
		isEmailMatch, err := regexp.Match("^([A-Za-z0-9_\\.-_]+)@([\\da-z]+)\\.com$", []byte(email))
		isTelephoneMatch, err := regexp.Match("^1[0-9]{10}$", []byte(telephone))

		if !isEmailMatch {
			log.Printf("Error: [%s doesn't match the rules] occur.\n", email)
			configrules()
			os.Exit(9)
		}
		if !isTelephoneMatch {
			log.Printf("Error: [%s doesn't match the rules] occur.\n", telephone)
			configrules()
			os.Exit(10)
		}

		if err == nil {
			// Todo Somethings
			fmt.Println("Agenda Command is \"register\".\ncalled with:")
			fmt.Printf("\tusername: %s\n\tpassword: %s\n\te-mail: %s\n\ttelephone: %s\n", username, password, email, telephone)
			/*data, err := entity.EncodeJSON(username, password, email, telephone)
			if err == nil {
				result := entity.DecodeJSON(data)
				fmt.Println(data)
				fmt.Println(result, len(result))
			}*/
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

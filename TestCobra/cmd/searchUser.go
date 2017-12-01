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

// searchCmd represents the search command
var searchUserCmd = &cobra.Command{
	Use:   "searchUser",
	Short: "User search",
	Long:  "Search user by username, e-mail or telephone",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		email, err := cmd.Flags().GetString("email")
		telephone, err := cmd.Flags().GetString("telephone")

		var isUsernameMissing, isUsernameMissingValue = false, false
		var isEmailMissing, isEmailMissingValue = false, false
		var isTelephoneMissing = false

		if username == "Anonymous" {
			isUsernameMissing = true
		}
		if username[0] == '-' {
			isUsernameMissingValue = true
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

		if isUsernameMissingValue {
			log.Printf("Error: [\"-u/--username\" doesn't own an argument value] occur.\n")
			os.Exit(2)
		} else if isEmailMissingValue {
			log.Printf("Error: [\"-e/--email\" doesn't own an argument value] occur.\n")
			os.Exit(3)
		}

		fmt.Println(isUsernameMissing, isEmailMissing, isTelephoneMissing)

		/*Regular Expression*/
		isEmailMatch, err := regexp.Match("^([A-Za-z0-9_\\.-_]+)@([\\da-z]+)\\.com$", []byte(email))
		isTelephoneMatch, err := regexp.Match("^1[0-9]{10}$", []byte(telephone))

		if !isEmailMatch {
			log.Printf("Error: [%s doesn't match the rules] occur.\n", email)
			configrules()
			os.Exit(4)
		}
		if !isTelephoneMatch {
			log.Printf("Error: [%s doesn't match the rules] occur.\n", telephone)
			configrules()
			os.Exit(5)
		}

		if err == nil {
			// Todo Somethings
			fmt.Println("Agenda Command is \"search\".\ncalled with:")
			fmt.Printf("\tusername: %s\n\temail: %s\n\ttelephone: %s\n", username, email, telephone)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchUserCmd)

	searchUserCmd.Flags().StringP("username", "u", "Anonymous", "Username for search")
	searchUserCmd.Flags().StringP("email", "e", "Anonymous", "E-mail for search")
	searchUserCmd.Flags().StringP("telephone", "t", "Anonymous", "Telephone for search")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

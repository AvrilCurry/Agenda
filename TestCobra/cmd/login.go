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
		if username[0] == '-' && username[1] == '-' {
			isUsernameMissingValue = true
		}

		if password == "Anonymous" {
			isPasswordMissing = true
		}

		if isUsernameMissing {
			log.Printf("Error: [Missing option \"--username\"] occur.\n")
			os.Exit(2)
		} else if isUsernameMissingValue {
			log.Printf("Error: [\"--username\" doesn't own an argument value] occur.\n")
			os.Exit(3)
		} else if isPasswordMissing {
			log.Printf("Error: [Missing option \"--password\"] occur.\n")
			os.Exit(4)
		}

		if err == nil {
			// Todo Somethings
			fmt.Println("Agenda Command is \"login\".\ncalled with:")
			fmt.Printf("\tusername: %s\n\tpassword: %s\n", username, password)
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

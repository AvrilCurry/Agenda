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
	"log"
	"os"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Help messages for Agenda",
	Long:  "list all the commands and its descriptions",
	Run: func(cmd *cobra.Command, args []string) {
		res1, err := cmd.Flags().GetBool("addpar")
		res2, err := cmd.Flags().GetBool("cancel")
		res3, err := cmd.Flags().GetBool("clear")
		res4, err := cmd.Flags().GetBool("cm")
		res5, err := cmd.Flags().GetBool("deletepar")
		res6, err := cmd.Flags().GetBool("deleteUser")
		res7, err := cmd.Flags().GetBool("login")
		res8, err := cmd.Flags().GetBool("logout")
		res9, err := cmd.Flags().GetBool("quit")
		res10, err := cmd.Flags().GetBool("register")
		res11, err := cmd.Flags().GetBool("searchMeeting")
		res12, err := cmd.Flags().GetBool("searchUser")

		if err == nil {
			fout, _ := os.OpenFile("./log/error.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
			MyCorrectLogger := log.New(fout, "[Correct]: ", log.Ldate|log.Ltime)
			outputInfo := "Agenda Command is \"help\".\n\tcalled with:"
			helpInfo := "\t\n\t\t\t--addpar\n\t\t\t--cancel\n\t\t\t--clear\n\t\t\t--cm\n\t\t\t--deletepar\n\t\t\t--deleteUser\n\t\t\t--login\n\t\t\t--logout\n\t\t\t--quit\n\t\t\t--register\n\t\t\t--searchMeeting\n\t\t\t--searchUser\n\n"

			if res1 {
				MyCorrectLogger.Printf(outputInfo + " \"--adpar\"\n\t\tOutput:\n\t\t\tagenda addpar --username username --title Title --participator [xxx, ...]\n\n")
			} else if res2 {
				MyCorrectLogger.Printf(outputInfo + " \"--cancel\"\n\t\tOutput:\n\t\t\tagenda cancel --username --title Title\n\n")
			} else if res3 {
				MyCorrectLogger.Printf(outputInfo + " \"--clear\"\n\t\tOutput:\n\t\t\tagenda clear --username username\n\n")
			} else if res4 {
				MyCorrectLogger.Printf(outputInfo + " \"--cm\"\n\t\tOutput:\n\t\t\tagenda cm --username username --title Title --participator [xxx,xxx, ...] --startTime StartTime --endTime EndTime\n\n")
			} else if res5 {
				MyCorrectLogger.Printf(outputInfo + " \"--deletepar\"\n\t\tOutput:\n\t\t\tagenda deletepar --username username --title Title --participator [xxx, ...]\n\n")
			} else if res6 {
				MyCorrectLogger.Printf(outputInfo + " \"--deleteUser\"\n\t\tOutput:\n\t\t\tagenda deleteUser --username username\n\n")
			} else if res7 {
				MyCorrectLogger.Printf(outputInfo + " \"--login\"\n\t\tOutput:\n\t\t\tagenda login -u/--username Username -p/--password Passwor\n\n")
			} else if res8 {
				MyCorrectLogger.Printf(outputInfo + " \"--logout\"\n\t\tOutput:\n\t\t\tagenda logout --username username\n\n")
			} else if res9 {
				MyCorrectLogger.Printf(outputInfo + " \"--quit\"\n\t\tOutput:\n\t\t\tagenda quit --username username --title Title\n\n")
			} else if res10 {
				MyCorrectLogger.Printf(outputInfo + " \"--register\"\n\t\tOutput:\n\t\t\tagenda register -u/-username Username -p/--password Password -e/--email Email -t/--telephone Telephone\n\n")
			} else if res11 {
				MyCorrectLogger.Printf(outputInfo + " \"--searchMeeting\"\n\t\tOutput:\n\t\t\tagenda searchMeeting --username username --startTime StartTime --endTime EndTime\n\n")
			} else if res12 {
				MyCorrectLogger.Printf(outputInfo + " \"--searchUser\"\n\t\tOutput:\n\t\t\tagenda searchUser -u/--username username -i/--information [Username/Email/Telephone]\n\n")
			} else {
				MyCorrectLogger.Printf(outputInfo + " \"no options\"\n\t\tOutput:" + helpInfo + "\n")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

	helpCmd.Flags().BoolP("addpar", "", false, "addpar For help")
	helpCmd.Flags().BoolP("cancel", "", false, "cancel For help")
	helpCmd.Flags().BoolP("clear", "", false, "clear For help")
	helpCmd.Flags().BoolP("cm", "", false, "cm For help")
	helpCmd.Flags().BoolP("deletepar", "", false, "deletepar For help")
	helpCmd.Flags().BoolP("deleteUser", "", false, "deleteUser For help")
	helpCmd.Flags().BoolP("login", "", false, "login For help")
	helpCmd.Flags().BoolP("logout", "", false, "logout For help")
	helpCmd.Flags().BoolP("quit", "", false, "quit For help")
	helpCmd.Flags().BoolP("register", "", false, "register For help")
	helpCmd.Flags().BoolP("searchMeeting", "", false, "searchMeeting For help")
	helpCmd.Flags().BoolP("searchUser", "", false, "searchUser For help")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

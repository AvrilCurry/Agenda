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
	"os"

	"github.com/spf13/cobra"
	"github.com/voxelbrain/goptions"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Help messages for Agenda",
	Long:  "list all the commands and its descriptions",
	Run: func(cmd *cobra.Command, args []string) {
		/*helpInformation, _ := cmd.Flags().GetBool("name")
		if helpInformation {
			fmt.Println("help called by register")
		} else {
			fmt.Println("help called")
		}*/
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

	var options struct {
		Register bool `goptions:"--register, description='help called by register'"`
		Login    bool `goptions:"--login, description='help called by login'"`
	}

	arg := []string{"--register", "--login"}
	fs := goptions.NewFlagSet("goptions", &options)
	err := fs.Parse(arg)
	if err == nil {
		fmt.Println("Yes")
	}
	if err == goptions.ErrHelpRequest {
		fs.PrintHelp(os.Stdout)
		return
	} else if err != nil {
		fmt.Printf("Failure: %s\n", err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

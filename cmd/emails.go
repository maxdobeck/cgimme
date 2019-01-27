// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

// emailsCmd represents the emails command
var emailsCmd = &cobra.Command{
	Use:   "emails",
	Short: "search for emails",
	Long: `Find all emails in your clipboard or an intput file.  By default will check the clipboard but can instead check a user specified file.  Any emails will be printed to the terminal and then copied into the clipboard.

Emails consist of a string separated by spaces that contain the @ (at) symbol.`,
	Run: func(cmd *cobra.Command, args []string) {
		src, err := cmd.Flags().GetString("file")
		if err != nil {
			panic(err)
		}
		if src != "" {
			found, err := parseFile(src)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(found["email"])
		} else {
			found, err := parseClipboard()
			if err != nil {
				fmt.Println(err)
			}
			cpEmails(found["email"])
		}
	},
}

func init() {
	rootCmd.AddCommand(emailsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	emailsCmd.Flags().StringP("file", "f", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// emailsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func parseFile(file string) (map[string][]string, error) {
	found := make(map[string][]string)
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	content, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	words := strings.Fields(string(content))
	for _, word := range words {
		if strings.ContainsRune(word, 64) {
			found["email"] = append(found["email"], word)
		}
	}
	return found, nil
}

func parseClipboard() (map[string][]string, error) {
	found := make(map[string][]string)
	src := getClipboard()
	words := strings.Fields(src)
	for _, word := range words {
		if strings.ContainsRune(word, 64) {
			found["email"] = append(found["email"], word)
		}
	}
	return found, nil
}

func getClipboard() string {
	content, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Problem reading from your clipboard\n: ", err)
	}
	return content
}

func cpEmails(data []string) {
	var cp string
	for _, email := range data {
		cp = cp + email + "\n"
	}
	err := clipboard.WriteAll(cp)
	if err != nil {
		panic(err)
	}
	fmt.Println(cp)
}

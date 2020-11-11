/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/treethought/markov/pkg"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen [corpus]",
	Short: "Generate markov text from corpus file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := markov.New()

		corpus := args[0]
		data, err := ioutil.ReadFile(corpus)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.FromString(string(data))
		c.Generate()

	},
}

func init() {
	log.SetLevel(log.InfoLevel)
	rootCmd.AddCommand(genCmd)
}

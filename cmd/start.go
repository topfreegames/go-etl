// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/topfreegames/go-etl/app"
	"github.com/topfreegames/go-etl/reader"
)

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starting worker",
	Long:  `starting worker`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("reading config")
		readConfig()

		log.Print("getting job reader")
		reader := reader.NewConfigReader()

		log.Print("configuring server")
		app, err := app.NewApp(reader)
		if err != nil {
			log.Panic(err)
		}

		log.Print("starting app")
		app.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

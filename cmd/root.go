/*
Copyright © 2023 Josh Hill josh.hill@vetspire.com

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
	"context"
	"os"

	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
	token  string
	ctx    context.Context
	client *github.Client
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ghman",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Github Personal Access Token with correct permissions.")
}

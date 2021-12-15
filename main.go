/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cfwctl",
	Short: "Cloud Firewall Control tools",
	Long:  `Cloud Firewall Control tools`,
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get rules list",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		Get()
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a rule",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		Add()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cfwctl.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(addCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cfwctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cfwctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func Get() {
	client := NewClient()
	resopnse := client.GetRules()
	fmt.Print(resopnse)
}

func Add() {
	myPublicIP := GetPublicIP()
	if myPublicIP == "" {
		fmt.Println("get my public ip failed")
		os.Exit(1)
	}
	firewallRules := []*lighthouse.FirewallRule{
		&lighthouse.FirewallRule{
			Protocol:                common.StringPtr("ALL"),
			Port:                    common.StringPtr("ALL"),
			CidrBlock:               common.StringPtr(myPublicIP),
			Action:                  common.StringPtr("ACCEPT"),
			FirewallRuleDescription: common.StringPtr("add-by-cfwctl"),
		},
	}
	client := NewClient()
	client.AddRules(firewallRules)
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}

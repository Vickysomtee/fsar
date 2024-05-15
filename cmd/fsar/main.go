/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vickysomtee/fsar"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fsar",
	Short: "A brief description of your application",
}

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Creates directory for the storage operations",
	Run: func(cmd *cobra.Command, args []string) {
		dirName, _ := cmd.Flags().GetString("name")
		fsar.MountFsar(dirName)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(mountCmd)
	mountCmd.Flags().String("name", "", "Name of directory where buckets will be created")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

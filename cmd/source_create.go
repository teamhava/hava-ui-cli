/*
Copyright © 2023 Hava Pty Ltd support@hava.io
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create sources to Hava",
	Long:  `Create sources to Hava for different providers (AWS/Azure/Google Cloud)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()

	},
}

func init() {
	sourceCmd.AddCommand(sourceCreateCmd)

	sourceCreateCmd.PersistentFlags().String("name", "", "The name for this source")
	sourceCreateCmd.MarkFlagRequired("name")
}

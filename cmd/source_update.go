/*
Copyright © 2023 Hava Pty Ltd support@hava.io
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update sources to Hava",
	Long:  `Update sources to Hava for different providers (AWS/Azure/Google Cloud)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	sourceCmd.AddCommand(sourceUpdateCmd)

	sourceUpdateCmd.PersistentFlags().String("name", "", "The name for this source")
	sourceUpdateCmd.PersistentFlags().String("source-id", "", "sourceId of AWS|Azure|GCP source")
	sourceUpdateCmd.MarkFlagRequired("source-id")
}

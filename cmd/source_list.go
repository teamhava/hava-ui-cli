/*
Copyright © 2023 Hava Pty Ltd support@hava.io
*/
package cmd

import (
	"github.com/spf13/cobra"
	apiclient "github.com/teamhava/hava-ui-cli/havaclient"
)

// sourceCmd represents the source command
var sourceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sources of Hava",
	Long: `List sources of Hava
	
	# List all sources
	hava source list

	# List specific source
	hava source list --source-id a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38`,
	Run: func(cmd *cobra.Command, args []string) {

		myclient, err := apiclient.GetNewClient()

		if err != nil {
			o.AddErrorMessage("Error creating havaclient:", err.Error(), true)
		}

		// Get Source flags
		sourceIdFlag, _ := cmd.Flags().GetString("source-id")

		if sourceIdFlag != "" {
			showSourceId(myclient, sourceIdFlag)
		} else {
			showSources(myclient)
		}
	},
}

func init() {
	sourceCmd.AddCommand(sourceListCmd)
	sourceListCmd.Flags().String("source-id", "", "hava source list --sourceId a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38")
}

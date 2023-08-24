/*
Copyright © 2023 Hava Pty Ltd support@hava.io
*/
package cmd

import (
	"context"
	"os"
	"strings"

	"github.com/spf13/cobra"
	havaclient "github.com/teamhava/hava-sdk-go"
	apiclient "github.com/teamhava/hava-ui-cli/havaclient"
)

// sourceCmd represents the source command
var sourceDeleteCmd = &cobra.Command{
	Use:   "delete <source-id>",
	Short: "Delete sources to Hava",
	Long: `Delete sources to Hava for different providers (AWS/Azure/Google Cloud)
	
	Example: 
	hava source delete a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38`,
	Run: func(cmd *cobra.Command, args []string) {

		myclient, err := apiclient.GetNewClient()

		if err != nil {
			o.AddErrorMessage("Error creating havaclient:", err.Error(), true)
		}

		// Retrieve Source to Delete source

		sourceId := strings.ToLower(args[0])

		if sourceId != "" {
			sourceDelete(myclient, sourceId)
		} else {
			cmd.Help()
		}

	},
}

func init() {
	sourceCmd.AddCommand(sourceDeleteCmd)
}

func sourceDelete(api *havaclient.APIClient, sourceId string) {

	if !confirm() {
		o.AddInfoMessage("\n\t**** Canceling `hava source delete`  **** ")
		os.Exit(1)
	}
	ctx := context.Background()

	req := api.SourcesApi.SourcesDestroy(ctx, sourceId)

	_, res, err := req.Execute()

	if err != nil {
		o.AddErrorMessage("Error from Source Delete Request:", err.Error(), false)
	}

	if res.StatusCode != 200 {
		o.AddErrorMessage("Error from `SourcesApi.SourcesDestroy` API, Request Status:", res.Status, true)
	}

	o.AddInfoMessage("\n\tDelete scheduled for the following source:" + sourceId)

}

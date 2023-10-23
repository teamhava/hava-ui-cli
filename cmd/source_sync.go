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
var sourceSyncCmd = &cobra.Command{
	Use:   "sync  <source-id>",
	Short: "Sync sources to Hava",
	Long: `Sync sources to Hava for different providers (AWS/Azure/Google Cloud)
	
	Example:
	hava source sync a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38`,
	Run: func(cmd *cobra.Command, args []string) {

		myclient, err := apiclient.GetNewClient()

		if err != nil {
			o.AddErrorMessage("Error creating havaclient:", err.Error(), true)
		}

		// Get Source flags
		sourceIdFlag := strings.ToLower(args[0])

		if sourceIdFlag != "" {
			syncSourceID(myclient, sourceIdFlag)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	sourceCmd.AddCommand(sourceSyncCmd)
}

func syncSourceID(api *havaclient.APIClient, sourceId string) {

	ctx := context.Background()

	sync := api.SourcesApi.SourcesSync(ctx, sourceId)

	res, err := sync.Execute()

	if err != nil {
		o.AddErrorMessage("Error from Source Sync Request:", err.Error(), false)
	}

	if res.StatusCode != 202 {
		o.AddErrorMessage("Error from `SourcesApi.SourcesSync`, Request Status:", res.Status, true)
	}

	req := api.SourcesApi.SourcesShow(ctx, sourceId)

	source, res, err := req.Execute()

	if err != nil {
		o.AddErrorMessage("Error from Sync Souce ID Request:", err.Error(), false)
	}

	if res.StatusCode == 422 {
		o.AddInfoMessage("Validation Error, Request Status:" + res.Status + " " + err.Error())
		os.Exit(0)
	} else if res.StatusCode == 404 {
		o.AddErrorMessage("Error from Source Sync Request, please check source-id:"+sourceId+" Request Status:", res.Status, true)
	} else if res.StatusCode != 200 {
		o.AddErrorMessage("Error from Source Sync Request, Request Status:", res.Status, true)
	}

	o.AddInfoMessage("Sync scheduled for the following source:")
	o.AddTableHeaders("DisplayName", "Id", "Info", "Name", "State", "Type")
	o.AddTableRows(
		SafeDeref(source.DisplayName), 
		SafeDeref(source.Id), 
		SafeDeref(source.Info), 
		SafeDeref(source.Name), 
		SafeDeref(source.State), 
		SafeDeref(source.Type))

}

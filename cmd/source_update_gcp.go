/*
Copyright © 2023 Hava Pty Ltd support@hava.io
*/
package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	havaclient "github.com/teamhava/hava-sdk-go"
	apiclient "github.com/teamhava/hava-ui-cli/havaclient"
)

// sourceCmd represents the source command
var sourceUpdateGcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Update GCP sources to Hava",
	Long: `Update GCP sources to Hava

	hava update source gcp --name GCPDev --source-id a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38 --config-file $GCP_ENCODED_FILE`,
	Run: func(cmd *cobra.Command, args []string) {

		myclient, err := apiclient.GetNewClient()

		if err != nil {
			o.AddErrorMessage("Error creating havaclient:", err.Error(), true)
		}

		// Retrieve flags to configure source
		sourceName, _ := cmd.Flags().GetString("name")
		configFile, _ := cmd.Flags().GetString("config-file")
		sourceId, _ := cmd.Flags().GetString("source-id")

		//Deteremine which Source to configure AWS|AZURE|GCP|K8s

		if configFile != "" || sourceName != "" || sourceId != "" {
			sourceUpdateGcp(myclient, sourceName, "GCP::ServiceAccountCredentials", configFile, sourceId)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	sourceUpdateCmd.AddCommand(sourceUpdateGcpCmd)
	sourceUpdateGcpCmd.Flags().String("config-file", "", "Base64 encoded credentials file")
}

func sourceUpdateGcp(api *havaclient.APIClient, sourceName string, sourceType string, configFile string, sourceId string) {

	ctx := context.Background()

	if sourceId == "" {
		o.AddErrorMessage("Missing flag :", "--source-id <source-id> missing.", true)
	}

	gcpCredentialsSource := &havaclient.SourcesGCPServiceAccountCredentials{
		Type: &sourceType,
	}

	//Check each value is valid. Empty values should be left out of the struct
	if sourceName != "" {
		gcpCredentialsSource.Name = &sourceName
	}
	if configFile != "" {
		gcpCredentialsSource.EncodedFile = &configFile
	}

	body := havaclient.SourcesGCPServiceAccountCredentialsAsSourcesUpdateRequest(gcpCredentialsSource)

	req := api.SourcesApi.SourcesUpdate(ctx, sourceId).SourcesUpdateRequest(body)

	_, res, err := req.Execute()

	if err != nil {
		o.AddErrorMessage("Error from Source Update Request:", err.Error(), false)
	}

	if res.StatusCode == 422 {
		o.AddInfoMessage("Validation Error, Request Status:" + res.Status + " " + err.Error())
		os.Exit(0)
	} else if res.StatusCode == 404 {
		o.AddErrorMessage("Error from Source Update Request, please check source-id:"+sourceId+" Request Status:", res.Status, true)
	} else if res.StatusCode != 200 {
		o.AddErrorMessage("Error from Source Update Request, Request Status:", res.Status, true)
	}

	o.AddInfoMessage("Update GCP Source for the following source:\n")
	showSourceId(api, sourceId)

}

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
var sourceCreateGcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Create GCP sources to Hava",
	Long: `Create GCP sources to Hava

	hava create source gcp --name GCPDev --config-file $GCP_ENCODED_FILE`,
	Run: func(cmd *cobra.Command, args []string) {

		myclient, err := apiclient.GetNewClient()

		if err != nil {
			o.AddErrorMessage("Error creating havaclient:", err.Error(), true)
		}

		// Retrieve flags to configure source
		sourceName, _ := cmd.Flags().GetString("name")
		configFile, _ := cmd.Flags().GetString("config-file")

		//Deteremine which Source to configure AWS|AZURE|GCP|K8s

		if configFile != "" {
			sourceCreateGcp(myclient, sourceName, "GCP::ServiceAccountCredentials", configFile)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	sourceCreateCmd.AddCommand(sourceCreateGcpCmd)
	sourceCreateGcpCmd.Flags().String("config-file", "", "Base64 encoded credentials file")
	sourceCreateAzureCmd.MarkFlagRequired("config-file")
}

func sourceCreateGcp(api *havaclient.APIClient, sourceName string, sourceType string, configFile string) {

	ctx := context.Background()

	gcpCredentialsSource := &havaclient.SourcesGCPServiceAccountCredentials{
		Name:        &sourceName,
		Type:        &sourceType,
		EncodedFile: &configFile,
	}

	body := havaclient.SourcesGCPServiceAccountCredentialsAsSourcesCreateRequest(gcpCredentialsSource)

	req := api.SourcesApi.SourcesCreate(ctx).SourcesCreateRequest(body)

	source, res, err := req.Execute()

	if err != nil {
		o.AddErrorMessage("Error from Source Create Request:", err.Error(), false)
	}

	if res.StatusCode == 422 {
		o.AddInfoMessage("Source already configured, Request Status:" + res.Status + " " + err.Error())
		os.Exit(0)
	} else if res.StatusCode != 200 {
		o.AddErrorMessage("Error from Source Create Request, Request Status:", res.Status, true)
	}

	o.AddInfoMessage("Created GCP Source for the following source:\n")
	o.AddTableHeaders("DisplayName", "Id", "Info", "Name", "State", "Type")
	o.AddTableRows(
		SafeDeref(source.DisplayName), 
		SafeDeref(source.Id), 
		SafeDeref(source.Info), 
		SafeDeref(source.Name), 
		SafeDeref(source.State), 
		SafeDeref(source.Type))

}

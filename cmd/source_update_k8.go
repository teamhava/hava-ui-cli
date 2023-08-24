/*
Copyright © 2023 Hava Pty Ltd support@hava.io
*/
package cmd

import (
	"github.com/spf13/cobra"
	havaclient "github.com/teamhava/hava-sdk-go"
	apiclient "github.com/teamhava/hava-ui-cli/havaclient"
)

// sourceCmd represents the source command
var sourceUpdatek8Cmd = &cobra.Command{
	Use:   "k8",
	Short: "Update K8 sources to Hava",
	Long: `Update K8 sources to Hava

	hava update source k8 --name K8Dev --source-id a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38 --config-file $KUBE_CONFIG`,
	Hidden: true,
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
			sourceUpdateK8(myclient, sourceName, "K8::ServiceAccountCredentials", configFile, sourceId)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	sourceUpdateCmd.AddCommand(sourceUpdatek8Cmd)
	sourceUpdatek8Cmd.Flags().String("config-file", "", "kubeconfig credentials file")
}

func sourceUpdateK8(api *havaclient.APIClient, sourceName string, sourceType string, configFile string, sourceId string) {

	// ctx := context.Background()

	// k8CredentialsSource := &havaclient.SourcesK8ServiceAccountCredentials{
	// 	Name: &sourceName,
	// 	Type: &sourceType,
	// }

	// //Check each value is valid. Empty values should be left out of the struct
	// if configFile != "" {
	// 	k8CredentialsSource.EncodedFile = &configFile
	// }

	// body := havaclient.SourcesK8ServiceAccountCredentialsAsSourcesUpdateRequest(k8CredentialsSource)

	// req := api.SourcesApi.SourcesUpdate(ctx, sourceId).SourcesUpdateRequest(body)

	// _, res, err := req.Execute()

	// if err != nil {
	// 	o.AddErrorMessage("Error from Source Update Request:", err.Error(), false)
	// }

	// if res.StatusCode == 422 {
	// 	o.AddInfoMessage("Validation Error, Request Status:" + res.Status + " " + err.Error())
	// 	os.Exit(0)
	// } else if res.StatusCode != 200 {
	// 	o.AddErrorMessage("Error from Source Create Request, Request Status:", res.Status, true)
	// }

	// o.AddInfoMessage("Update K8 Source for the following source:\n")
	// showSourceId(api, sourceId)
	o.AddInfoMessage("K8 Source feature coming soon...")

}

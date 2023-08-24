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
var sourceCreateK8Cmd = &cobra.Command{
	Use:   "k8",
	Short: "Create K8 sources to Hava",
	Long: `Create K8 sources to Hava,

	hava create source k8 --name K8Dev --config-file $KUBE_CONFIG`,
	Hidden: true,
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
			sourceCreateK8(myclient, sourceName, "K8::ServiceAccountCredentials", configFile)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	sourceCreateCmd.AddCommand(sourceCreateK8Cmd)
	sourceCreateK8Cmd.Flags().String("config-file", "", "kubeconfig credentials file")
	sourceCreateK8Cmd.MarkFlagRequired("config-file")
}

func sourceCreateK8(api *havaclient.APIClient, sourceName string, sourceType string, configFile string) {

	// ctx := context.Background()

	// k8CredentialsSource := &havaclient.SourcesK8ServiceAccountCredentials{
	// 	Name:        &sourceName,
	// 	Type:        &sourceType,
	// 	EncodedFile: &configFile,
	// }

	// body := havaclient.SourcesK8ServiceAccountCredentialsAsSourcesCreateRequest(k8CredentialsSource)

	// req := api.SourcesApi.SourcesCreate(ctx).SourcesCreateRequest(body)

	// source, res, err := req.Execute()

	// if err != nil {
	// 	o.AddErrorMessage("Error from Source Create Request:", err.Error(), false)
	// }

	// if res.StatusCode == 422 {
	// 	o.AddInfoMessage("Source already configured, Request Status:" + res.Status + " " + err.Error())
	// 	os.Exit(0)
	// } else if res.StatusCode != 200 {
	// 	o.AddErrorMessage("Error from Source Create Request, Request Status:", res.Status, true)
	// }

	// o.AddInfoMessage("Created K8 Source for the following source:\n")
	// o.AddTableHeaders("DisplayName", "Id", "Info", "Name", "State", "Type")
	// o.AddTableRows(*source.DisplayName, *source.Id, *source.Info, *source.Name, *source.State, *source.Type)
	o.AddInfoMessage("K8 Source feature coming soon...")

}

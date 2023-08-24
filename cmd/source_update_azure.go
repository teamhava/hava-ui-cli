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
var sourceUpdateAzureCmd = &cobra.Command{
	Use:   "azure",
	Short: "Update Azure sources to Hava",
	Long: `Update Azure sources to Hava for different providers (AWS/Azure/Google Cloud)

	Azure Example:
	hava source update azure --name AzureDev --source-id a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38 --client-id $ARM_CLIENT_ID --tenant-id $ARM_TENANT_ID --subscription-id $ARM_SUBSCRIPTION_ID`,
	Run: func(cmd *cobra.Command, args []string) {

		myclient, err := apiclient.GetNewClient()

		if err != nil {
			o.AddErrorMessage("Error creating havaclient:", err.Error(), true)
		}

		// Retrieve flags to configure source
		sourceName, _ := cmd.Flags().GetString("name")
		clientId, _ := cmd.Flags().GetString("client-id")
		tenantId, _ := cmd.Flags().GetString("tenant-id")
		subscriptionId, _ := cmd.Flags().GetString("subscription-id")
		secretKey, _ := cmd.Flags().GetString("secret-key")
		sourceId, _ := cmd.Flags().GetString("source-id")

		if clientId != "" || sourceName != "" || sourceId != "" || tenantId != "" || subscriptionId != "" || secretKey != "" {
			sourceUpdateAzure(myclient, sourceName, "Azure::Credentials", clientId, tenantId, subscriptionId, secretKey, sourceId)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	sourceUpdateCmd.AddCommand(sourceUpdateAzureCmd)
	sourceUpdateAzureCmd.Flags().String("subscription-id", "", "The ID of the Azure subscription to import from")
	sourceUpdateAzureCmd.Flags().String("tenant-id", "", "The GUID representing the Active Directory Tenant")
	sourceUpdateAzureCmd.Flags().String("client-id", "", "The Client ID for your Service Principle")
	sourceUpdateAzureCmd.Flags().String("secret-key", "", "The Client Secret for your Service Principle")
}

func sourceUpdateAzure(api *havaclient.APIClient, sourceName string, sourceType string, clientId string, tenantId string, subscriptionId string, secretKey string, sourceId string) {

	ctx := context.Background()

	if sourceId == "" {
		o.AddErrorMessage("Missing flag :", "--source-id <source-id> missing.", true)
	}

	azureCredentialsSource := &havaclient.SourcesAzureCredentials{
		Type: &sourceType,
	}

	//Check each value is valid. Empty values should be left out of the struct
	if sourceName != "" {
		azureCredentialsSource.Name = &sourceName
	}
	if subscriptionId != "" {
		azureCredentialsSource.SubscriptionId = &subscriptionId
	}
	if tenantId != "" {
		azureCredentialsSource.TenantId = &tenantId
	}
	if clientId != "" {
		azureCredentialsSource.ClientId = &clientId
	}
	if secretKey != "" {
		azureCredentialsSource.SecretKey = &secretKey
	}

	sourceUpdateRequest := havaclient.SourcesAzureCredentialsAsSourcesUpdateRequest(azureCredentialsSource)

	req := api.SourcesApi.SourcesUpdate(ctx, sourceId).SourcesUpdateRequest(sourceUpdateRequest)

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

	o.AddInfoMessage("Updated Azure Source for the following source:\n")
	showSourceId(api, sourceId)
}

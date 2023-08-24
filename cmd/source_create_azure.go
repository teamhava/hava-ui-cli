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
var sourceCreateAzureCmd = &cobra.Command{
	Use:   "azure",
	Short: "Create Azure sources to Hava",
	Long: `Create Azure sources to Hava for different providers (AWS/Azure/Google Cloud)

	Azure Example:
	hava source create azure --name AzureDev --client-id $ARM_CLIENT_ID --tenant-id $ARM_TENANT_ID --subscription-id $ARM_SUBSCRIPTION_ID`,
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

		if clientId != "" {
			sourceCreateAzure(myclient, sourceName, "Azure::Credentials", clientId, tenantId, subscriptionId, secretKey)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	sourceCreateCmd.AddCommand(sourceCreateAzureCmd)
	sourceCreateAzureCmd.Flags().String("subscription-id", "", "The ID of the Azure subscription to import from")
	sourceCreateAzureCmd.Flags().String("tenant-id", "", "The GUID representing the Active Directory Tenant")
	sourceCreateAzureCmd.Flags().String("client-id", "", "The Client ID for your Service Principle")
	sourceCreateAzureCmd.Flags().String("secret-key", "", "The Client Secret for your Service Principle")
	sourceCreateAzureCmd.MarkFlagRequired("subscription-id")
	sourceCreateAzureCmd.MarkFlagRequired("tenant-id")
	sourceCreateAzureCmd.MarkFlagRequired("client-id")
	sourceCreateAzureCmd.MarkFlagRequired("secret-key")

}

func sourceCreateAzure(api *havaclient.APIClient, sourceName string, sourceType string, clientId string, tenantId string, subscriptionId string, secretKey string) {

	ctx := context.Background()

	azureCredentialsSource := &havaclient.SourcesAzureCredentials{
		Name:           &sourceName,
		Type:           &sourceType,
		SubscriptionId: &subscriptionId,
		TenantId:       &tenantId,
		ClientId:       &clientId,
		SecretKey:      &secretKey,
	}

	body := havaclient.SourcesAzureCredentialsAsSourcesCreateRequest(azureCredentialsSource)

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

	o.AddInfoMessage("Created Azure Source for the following source:\n")
	o.AddTableHeaders("DisplayName", "Id", "Info", "Name", "State", "Type")
	o.AddTableRows(*source.DisplayName, *source.Id, *source.Info, *source.Name, *source.State, *source.Type)

}

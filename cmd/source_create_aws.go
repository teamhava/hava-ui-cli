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
var sourceCreateAwsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Create AWS sources to Hava",
	Long: `Create AWS sources to Hava for different providers (AWS/Azure/Google Cloud)
	
	Example (AWS KEY + AWS Secret Key method): 
	
		hava source create aws --name dev --access-key $AWS_ACCESS_KEY_ID --secret-key $AWS_SECRET_ACCESS_KEY
		
	Example (AWS Cross Account Role method):
	
		hava source create aws --name dev --role-arn arn:aws:iam::123456789012:role/HavaReadOnly`,
	Run: func(cmd *cobra.Command, args []string) {

		myclient, err := apiclient.GetNewClient()

		if err != nil {
			o.AddErrorMessage("Error creating havaclient:", err.Error(), true)
		}

		// Retrieve flags to configure source
		roleArn, _ := cmd.Flags().GetString("role-arn")
		sourceName, _ := cmd.Flags().GetString("name")
		accessKey, _ := cmd.Flags().GetString("access-key")
		secretKey, _ := cmd.Flags().GetString("secret-key")

		if roleArn != "" {
			sourceCreateAwsCar(myclient, sourceName, "AWS::CrossAccountRole", roleArn, "549f601f9feef215614720ec50ef9182")
		} else if accessKey != "" && secretKey != "" {
			sourceCreateAwsKey(myclient, sourceName, "AWS::Keys", accessKey, secretKey)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	sourceCreateCmd.AddCommand(sourceCreateAwsCmd)
	sourceCreateAwsCmd.Flags().String("access-key", "", "The access key for your AWS account")
	sourceCreateAwsCmd.Flags().String("secret-key", "", "The secret key for your AWS account	")
	sourceCreateAwsCmd.Flags().String("role-arn", "", "The ARN of the role Hava is to assume in your account to import resources")

}

func sourceCreateAwsKey(api *havaclient.APIClient, sourceName string, awsType string, accessKey string, secretKey string) {

	ctx := context.Background()

	awsKeySource := &havaclient.SourcesAWSKey{
		Name:      &sourceName,
		Type:      &awsType,
		AccessKey: &accessKey,
		SecretKey: &secretKey,
	}

	body := havaclient.SourcesAWSKeyAsSourcesCreateRequest(awsKeySource)

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

	o.AddInfoMessage("Created AWS Source for the following source:\n")
	o.AddTableHeaders("DisplayName", "Id", "Info", "Name", "State", "Type")
	o.AddTableRows(
		SafeDeref(source.DisplayName), 
		SafeDeref(source.Id), 
		SafeDeref(source.Info), 
		SafeDeref(source.Name), 
		SafeDeref(source.State), 
		SafeDeref(source.Type))
}

func sourceCreateAwsCar(api *havaclient.APIClient, sourceName string, awsType string, roleArn string, externalID string) {

	ctx := context.Background()

	sawscar := &havaclient.SourcesAWSCAR{
		Name:       &sourceName,
		Type:       &awsType,
		RoleArn:    &roleArn,
		ExternalId: &externalID,
	}

	body := havaclient.SourcesAWSCARAsSourcesCreateRequest(sawscar)

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

	o.AddInfoMessage("Created AWS Source for the following source:\n")
	o.AddTableHeaders("DisplayName", "Id", "Info", "Name", "State", "Type")
	o.AddTableRows(
		SafeDeref(source.DisplayName), 
		SafeDeref(source.Id), 
		SafeDeref(source.Info), 
		SafeDeref(source.Name), 
		SafeDeref(source.State), 
		SafeDeref(source.Type))

}

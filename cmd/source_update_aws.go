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
var sourceUpdateAwsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Update AWS sources to Hava",
	Long: `Update AWS sources to Hava for different providers (AWS/Azure/Google Cloud)
	
	Example (AWS KEY + AWS Secret Key method): 
	
		hava source update aws --name dev --source-id a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38  --access-key $AWS_ACCESS_KEY_ID --secret-key $AWS_SECRET_ACCESS_KEY
		
	Example (AWS Cross Account Role method):
	
		hava source update aws --name dev --source-id a58b7cb1-f9da-42ad-9fc1-8dc61b0d3e38 --role-arn arn:aws:iam::123456789012:role/HavaReadOnly`,
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
		sourceId, _ := cmd.Flags().GetString("source-id")

		if roleArn != "" {
			sourceUpdateAwsCar(myclient, sourceName, "AWS::CrossAccountRole", roleArn, "549f601f9feef215614720ec50ef9182", sourceId)
		} else if accessKey != "" || secretKey != "" {
			sourceUpdateAwsKey(myclient, sourceName, "AWS::Keys", accessKey, secretKey, sourceId)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	sourceUpdateCmd.AddCommand(sourceUpdateAwsCmd)
	sourceUpdateAwsCmd.Flags().String("access-key", "", "The access key for your AWS account")
	sourceUpdateAwsCmd.Flags().String("secret-key", "", "The secret key for your AWS account	")
	sourceUpdateAwsCmd.Flags().String("role-arn", "", "The ARN of the role Hava is to assume in your account to import resources")
}

func sourceUpdateAwsKey(api *havaclient.APIClient, sourceName string, awsType string, accessKey string, secretKey string, sourceId string) {

	ctx := context.Background()

	if sourceId == "" {
		o.AddErrorMessage("Missing flag :", "--source-id <source-id> missing.", true)
	}

	awsKeySource := &havaclient.SourcesAWSKey{
		Type: &awsType,
	}

	//Check each value is valid. Empty values should be left out of the struct
	if sourceName != "" {
		awsKeySource.Name = &sourceName
	}
	if accessKey != "" {
		awsKeySource.AccessKey = &accessKey
	}
	if secretKey != "" {
		awsKeySource.SecretKey = &secretKey
	}

	sourceUpdateRequest := havaclient.SourcesAWSKeyAsSourcesUpdateRequest(awsKeySource)

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

	o.AddInfoMessage("Updated AWS Source for the following source:\n")
	showSourceId(api, sourceId)

}

func sourceUpdateAwsCar(api *havaclient.APIClient, sourceName string, awsType string, roleArn string, externalID string, sourceId string) {

	ctx := context.Background()

	if sourceId == "" {
		o.AddErrorMessage("Missing flag :", "--source-id <source-id> missing.", true)
	}

	sawscar := &havaclient.SourcesAWSCAR{
		Type:       &awsType,
		ExternalId: &externalID,
	}

	//Check each value is valid. Empty values should be left out of the struct
	if sourceName != "" {
		sawscar.Name = &sourceName
	}
	if roleArn != "" {
		sawscar.RoleArn = &roleArn
	}

	sourceUpdateRequest := havaclient.SourcesAWSCARAsSourcesUpdateRequest(sawscar)

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

	o.AddInfoMessage("Updated AWS Source for the following source:\n")
	showSourceId(api, sourceId)

}

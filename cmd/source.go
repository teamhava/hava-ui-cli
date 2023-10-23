/*
Copyright © 2023 Hava Pty Ltd support@hava.io
*/
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	havaclient "github.com/teamhava/hava-sdk-go"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Create/List/Delete/Update sources to Hava",
	Long:  `Create/List/Delete/Update sources to Hava for different providers (AWS/Azure/Google Cloud)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(sourceCmd)
}

func showSourceId(api *havaclient.APIClient, sourceId string) {

	ctx := context.Background()

	req := api.SourcesApi.SourcesShow(ctx, sourceId)

	source, res, err := req.Execute()

	if err != nil {
		o.AddErrorMessage("Error from Show Sources API:", err.Error(), false)
	}

	if res.StatusCode != 200 {
		o.AddErrorMessage("Error when calling `SourcesApi.SourcesShow`, Request Status:", res.Status, true)
	}

	o.AddTableHeaders("DisplayName", "Id", "Info", "Name", "State", "Type")
	o.AddTableRows(
		SafeDeref(source.DisplayName), 
		SafeDeref(source.Id),
		SafeDeref(source.Info), 
		SafeDeref(source.Name), 
		SafeDeref(source.State), 
		SafeDeref(source.Type))

}

func showSources(api *havaclient.APIClient) {

	ctx := context.Background()

	// List all sources
	req := api.SourcesApi.SourcesIndex(ctx)

	source, res, err := req.Execute()

	if err != nil {
		o.AddErrorMessage("Error from Show Sources API:", err.Error(), false)
	}

	if res.StatusCode != 200 {
		o.AddErrorMessage("Error when calling `SourcesApi.SourcesIndex`, Request Status:", res.Status, true)
	}

	//Length of Results

	numResults := len(source.Results)

	if numResults > 0 {

		o.AddTableHeaders("DisplayName", "Id", "Info", "Name", "State", "Type")
		// Loop through results
		for i := 0; i < numResults; i++ {
			o.AddTableRows(
				SafeDeref(source.Results[i].DisplayName), 
				SafeDeref(source.Results[i].Id), 
				SafeDeref(source.Results[i].Info), 
				SafeDeref(source.Results[i].Name), 
				SafeDeref(source.Results[i].State), 
				SafeDeref(source.Results[i].Type))
		}

	} else {
		o.AddInfoMessage("No sources found in Hava.\nPlease check https://app.hava.io/<ORGNAME>/sources and add sources\n or use `hava source create [aws|azure|gcp]`")
	}

}

func confirm() bool {

	var input string

	o.AddUserInputMessage("Do you want to continue with this operation? [y|n]: ")

	auto, err := sourceCmd.Flags().GetBool("autoapprove")

	if err != nil {
		o.AddErrorMessage("Error Retrieving autoapprove flag value: ", err, false)
	}

	// Check if --autoapprove=false
	if !auto {
		_, err := fmt.Scanln(&input)
		if err != nil {
			o.AddErrorMessage("Error autoapprove input value: ", err, true)
		}
	} else {
		input = "y"
		o.AddUserInputMessage("y(autoapprove=true)")
	}

	input = strings.ToLower(input)

	if input == "y" || input == "yes" {
		return true
	}
	return false
}

package havaclient

import (
	"errors"

	"github.com/spf13/viper"
	havaclient "github.com/teamhava/hava-sdk-go"
)

func GetNewClient() (*havaclient.APIClient, error) {
	havaToken := viper.GetString("HAVA_TOKEN")
	havaEndpoint := viper.GetString("HAVA_ENDPOINT")

	if len(havaToken) == 0 {
		return nil, errors.New("Hava API token not found. Have you set the 'HAVA_TOKEN' as an environment variable or `hava_token` in your config file?")
	}

	if havaEndpoint == "" {
		havaEndpoint = "https://api.hava.io"
	}

	cfg := havaclient.NewConfiguration()
	cfg.Servers = havaclient.ServerConfigurations{
		{
			URL:         havaEndpoint,
			Description: "Hava Platform",
		},
	}

	cfg.UserAgent = "havacli 0.1"

	cfg.DefaultHeader["Authorization"] = "Bearer " + havaToken

	return havaclient.NewAPIClient(cfg), nil

}

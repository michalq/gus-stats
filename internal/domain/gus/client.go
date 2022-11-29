package gus

import (
	"github.com/michalq/gus-stats/internal/config"
	gus "github.com/michalq/gus-stats/pkg/client_gus"
)

func NewClient(cfg config.Config) *gus.APIClient {
	gusConfig := gus.NewConfiguration()
	gusConfig.Debug = false
	gusConfig.Servers = gus.ServerConfigurations{
		{
			URL:         "https://bdl.stat.gov.pl/api/v1",
			Description: "No description provided",
		},
	}
	gusConfig.DefaultHeader["X-ClientId"] = cfg.Gus.Client
	return gus.NewAPIClient(gusConfig)
}

package main

import (
	"github.com/ellisdon/terraform-provider-azuredevops/azuredevops"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: azuredevops.Provider,
	})
}

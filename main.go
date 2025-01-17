package main

import (
	"github.com/bishtawi/terraform-provider-auth0/auth0"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return auth0.Provider()
		},
	})
}

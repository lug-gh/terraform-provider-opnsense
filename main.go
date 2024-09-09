package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-provider-opnsense/opnsense"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: opnsense.Provider,
	})
}

package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sarguru/terraform-provider-nsot/nsot"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nsot.Provider,
	})
}

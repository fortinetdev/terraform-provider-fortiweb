package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	//"github.com/fortinetdev/terraform-provider-fortios/fortios"
	"github.com/fortinetdev/terraform-provider-fortiweb/fweb"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		//ProviderFunc: fortios.Provider})
		ProviderFunc: fortiweb.Provider})
}

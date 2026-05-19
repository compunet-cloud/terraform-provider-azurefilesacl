package main

import (
	"context"
	"log"

	"github.com/compunet-cloud/terraform-provider-avd-azurefilesacl/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	err := providerserver.Serve(context.Background(), provider.New(version), providerserver.ServeOpts{
		Address: "registry.terraform.io/compunet/avd/azurefilesacl",
	})
	if err != nil {
		log.Fatal(err)
	}
}

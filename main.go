package main

import (
	"context"
	"log"

	"github.com/compunet-cloud/terraform-provider-azurefilesacl/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	err := providerserver.Serve(context.Background(), provider.New(version), providerserver.ServeOpts{
		Address: "registry.terraform.io/compunet-cloud/azurefilesacl",
	})
	if err != nil {
		log.Fatal(err)
	}
}

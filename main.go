package main

import (
	"context"
	"flag"
	"log"

	"balhau.net/receita-provider/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version = "dev"
)

func main() {

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true if debug mode")
	flag.Parse()

	// Initialize the structure to setup the provider service
	// create the folder in .terraform.d/plugins/terraform.local/local
	// otherwise you'll need to publish this provider to be able to use it
	opts := providerserver.ServeOpts{
		Address: "terraform.local/balhau/receita",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	// If the creation of service give us an error just print it.
	if err != nil {
		log.Fatal(err.Error())
	}

}

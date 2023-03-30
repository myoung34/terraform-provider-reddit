package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"terraform-provider-reddit/reddit"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	err := providerserver.Serve(
		context.Background(),
		reddit.New,
		providerserver.ServeOpts{
			Address: "github.com/myoung34/reddit",
		},
	)

	if err != nil {
		panic(fmt.Sprintf("Error serving provider: %v", err))
	}
}

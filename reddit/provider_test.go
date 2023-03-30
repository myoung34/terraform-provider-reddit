package reddit

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func makeProviderFactoryMap(name string, prov *RedditProvider) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		name: providerserver.NewProtocol6WithError(prov),
	}
}

func TestReddit(t *testing.T) {
	const testConfig = // language=hcl
	`
provider "reddit" {
  readonly = true
}


data "reddit_subreddit" "golang" {
  name = "golang"
}

output "golang" {
  value = data.reddit_subreddit.golang
}

`

	var (
		prov = new(RedditProvider)
	)

	resource.Test(
		t,
		resource.TestCase{
			ProtoV6ProviderFactories: makeProviderFactoryMap("reddit", prov),
			Steps: []resource.TestStep{
				{
					Config: testConfig,
				},
			},
		},
	)
}

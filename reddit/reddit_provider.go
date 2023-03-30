package reddit

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"os"
)

func New() provider.Provider {
	return &RedditProvider{}
}

type RedditProvider struct{} //revive:disable-line:exported

type redditProviderData struct {
	ID       types.String `tfsdk:"id"`
	Secret   types.String `tfsdk:"secret"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	ReadOnly types.Bool   `tfsdk:"readonly"`
}

func (p *RedditProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "reddit"
}

func (p *RedditProvider) Schema(_ context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":       schema.StringAttribute{Optional: true},
			"secret":   schema.StringAttribute{Optional: true, Sensitive: true},
			"username": schema.StringAttribute{Optional: true},
			"password": schema.StringAttribute{Optional: true, Sensitive: true},
			"readonly": schema.BoolAttribute{Optional: true},
		},
	}
}

func (p *RedditProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Reddit client")

	// Retrieve provider data from configuration
	var config redditProviderData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readOnly = config.ReadOnly.ValueBool()

	client, _ := reddit.NewReadonlyClient()
	if !readOnly {

		var secret string
		if config.Secret.IsUnknown() {
			resp.Diagnostics.AddWarning(
				"Unable to create client",
				"Cannot use unknown value as Secret",
			)
			return
		}

		if config.Secret.IsNull() {
			secret = os.Getenv("REDDIT_SECRET")
		} else {
			secret = config.Secret.ValueString()
		}

		if secret == "" {
			// Error vs warning - empty value must stop execution
			resp.Diagnostics.AddError(
				"Unable to find customer id",
				"Secret cannot be an empty string",
			)
			return
		}

		var id string
		if config.ID.IsUnknown() {
			// Cannot connect to client with an unknown value
			resp.Diagnostics.AddWarning(
				"Unable to create client",
				"Cannot use unknown value as apiKey",
			)
			return
		}

		if config.ID.IsNull() {
			id = os.Getenv("REDDIT_ID")
		} else {
			id = config.ID.ValueString()
		}

		if id == "" {
			// Error vs warning - empty value must stop execution
			resp.Diagnostics.AddError(
				"Unable to find id",
				"ID cannot be an empty string",
			)
			return
		}

		var password string
		if config.Password.IsUnknown() {
			// Cannot connect to client with an unknown value
			resp.Diagnostics.AddError(
				"Unable to create client",
				"Cannot use unknown value as password",
			)
			return
		}

		if config.Password.IsNull() {
			password = os.Getenv("REDDIT_PASSWORD")
		} else {
			password = config.Password.ValueString()
		}

		if password == "" {
			// Error vs warning - empty value must stop execution
			resp.Diagnostics.AddError(
				"Unable to find password",
				"Password cannot be an empty string",
			)
			return
		}

		// User must specify a host
		var username string
		if config.Username.IsUnknown() {
			// Cannot connect to client with an unknown value
			resp.Diagnostics.AddError(
				"Unable to create client",
				"Cannot use unknown value as username",
			)
			return
		}

		if config.Username.IsNull() {
			username = os.Getenv("REDDIT_USERNAME")
		} else {
			username = config.Username.ValueString()
		}

		if username == "" {
			// Error vs warning - empty value must stop execution
			resp.Diagnostics.AddError(
				"Unable to find username",
				"Username cannot be an empty string",
			)
			return
		}

		// Create a new reddit client and set it to the provider.client
		client, err := reddit.NewClient(reddit.Credentials{
			ID:       id,
			Secret:   secret,
			Username: username,
			Password: password,
		})
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to create client",
				"Unable to create reddit client:\n\n"+err.Error(),
			)
			return
		}
		resp.DataSourceData = client
		resp.ResourceData = client
	} else {
		resp.DataSourceData = client
		resp.ResourceData = client
	}

	tflog.Info(ctx, "Configured Reddit client", map[string]any{"success": true})

}

func (p *RedditProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *RedditProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		SubredditDataSource,
	}
}

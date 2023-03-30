package reddit

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/vartanbeno/go-reddit/v2/reddit"

	"github.com/hashicorp/terraform-plugin-framework/path"
)

func SubredditDataSource() datasource.DataSource {
	return &subredditDataSource{}
}

type subredditDataSource struct {
	client *reddit.Client
}

func (d *subredditDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subreddit"
}

func (d *subredditDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*reddit.Client)
}

func (d *subredditDataSource) Schema(_ context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name":                   schema.StringAttribute{Optional: true},
			"id":                     schema.StringAttribute{Optional: true},
			"full_id":                schema.StringAttribute{Optional: true},
			"created":                schema.StringAttribute{Optional: true},
			"url":                    schema.StringAttribute{Optional: true},
			"name_prefixed":          schema.StringAttribute{Optional: true},
			"title":                  schema.StringAttribute{Optional: true},
			"description":            schema.StringAttribute{Optional: true},
			"type":                   schema.StringAttribute{Optional: true},
			"suggested_comment_sort": schema.StringAttribute{Optional: true},
			"subscribers":            schema.Int64Attribute{Optional: true},
			"active_user_count":      schema.Int64Attribute{Optional: true},
			"nsfw":                   schema.BoolAttribute{Optional: true},
			"user_is_mod":            schema.BoolAttribute{Optional: true},
			"subscribed":             schema.BoolAttribute{Optional: true},
			"favorite":               schema.BoolAttribute{Optional: true},
		},
	}
}

func (d *subredditDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var subredditName string

	nameAttr := req.Config.GetAttribute(ctx, path.Root("name"), &subredditName)

	resp.Diagnostics.Append(nameAttr...)

	subredditResp, _, err := reddit.DefaultClient().Subreddit.Get(ctx, "golang")

	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to read.",
			"Could not get subreddit with name  "+subredditName+": "+err.Error(),
		)
		return
	}

	var result = Subreddit{
		ID:                   types.StringValue(subredditResp.ID),
		Name:                 types.StringValue(subredditResp.Name),
		FullID:               types.StringValue(subredditResp.FullID),
		Created:              types.StringValue(subredditResp.Created.String()),
		URL:                  types.StringValue(subredditResp.URL),
		NamePrefixed:         types.StringValue(subredditResp.NamePrefixed),
		Title:                types.StringValue(subredditResp.Title),
		Description:          types.StringValue(subredditResp.Description),
		Type:                 types.StringValue(subredditResp.Type),
		SuggestedCommentSort: types.StringValue(subredditResp.SuggestedCommentSort),
		Subscribers:          types.Int64Value(int64(subredditResp.Subscribers)),
		ActiveUserCount:      types.Int64Value(int64(*subredditResp.ActiveUserCount)),
		NSFW:                 types.BoolValue(subredditResp.NSFW),
		UserIsMod:            types.BoolValue(subredditResp.UserIsMod),
		Subscribed:           types.BoolValue(subredditResp.Subscribed),
		Favorite:             types.BoolValue(subredditResp.Favorite),
	}
	diags := resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

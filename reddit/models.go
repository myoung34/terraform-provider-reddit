package reddit

import "github.com/hashicorp/terraform-plugin-framework/types"

type Subreddit struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	FullID               types.String `tfsdk:"full_id"`
	Created              types.String `tfsdk:"created"`
	URL                  types.String `tfsdk:"url"`
	NamePrefixed         types.String `tfsdk:"name_prefixed"`
	Title                types.String `tfsdk:"title"`
	Description          types.String `tfsdk:"description"`
	Type                 types.String `tfsdk:"type"`
	SuggestedCommentSort types.String `tfsdk:"suggested_comment_sort"`
	Subscribers          types.Int64  `tfsdk:"subscribers"`
	ActiveUserCount      types.Int64  `tfsdk:"active_user_count"`
	NSFW                 types.Bool   `tfsdk:"nsfw"`
	UserIsMod            types.Bool   `tfsdk:"user_is_mod"`
	Subscribed           types.Bool   `tfsdk:"subscribed"`
	Favorite             types.Bool   `tfsdk:"favorite"`
}

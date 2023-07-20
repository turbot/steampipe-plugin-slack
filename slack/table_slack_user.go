package slack

import (
	"context"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableSlackUser() *plugin.Table {
	return &plugin.Table{
		Name:        "slack_user",
		Description: "Slack workspace users.",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "email"}),
			Hydrate:    getUser,
		},
		Columns: slackCommonColumns([]*plugin.Column{
			// NOTE:
			// * Profile fields collapsed to top level, there is no meaningful
			//   difference between a user and a profile, so profile_* feels redundant.
			// * Email requires the users:email.read oauth scope, so is often empty
			// * The name field is deprecated by slack, so not as important as it appears.

			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the user."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.DisplayName").NullIfZero(), Description: "Indicates the display name that the user has chosen to identify themselves by in their workspace profile."},

			// Other columns
			{Name: "api_app_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.ApiAppID").NullIfZero(), Description: "If an app user, then this is the unique identifier of the installed Slack application."},
			{Name: "bot_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.BotID").NullIfZero(), Description: "If a bot user, this is the unique identifier of the bot."},
			{Name: "color", Type: proto.ColumnType_STRING, Description: "Used in some clients to display a special username color."},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "True if the user has been deleted."},
			{Name: "display_name_normalized", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.DisplayNameNormalized").NullIfZero(), Description: "The display name, but with any non-Latin characters filtered out."},
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Email").NullIfZero(), Description: "Email address of the user."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.FirstName").NullIfZero(), Description: "First name of the user."},
			{Name: "has_2fa", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Has2FA"), Description: "True if two-factor authentication is enabled for the user."},
			{Name: "image_24", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Image24").NullIfZero(), Description: "URL of the user profile image, size 24x24 pixels."},
			{Name: "image_32", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Image32").NullIfZero(), Description: "URL of the user profile image, size 32x32 pixels."},
			{Name: "image_48", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Image48").NullIfZero(), Description: "URL of the user profile image, size 48x48 pixels."},
			{Name: "image_72", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Image72").NullIfZero(), Description: "URL of the user profile image, size 72x72 pixels."},
			{Name: "image_192", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Image192").NullIfZero(), Description: "URL of the user profile image, size 192x192 pixels."},
			{Name: "image_512", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Image512").NullIfZero(), Description: "URL of the user profile image, size 512x512 pixels."},
			{Name: "image_original", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.ImageOriginal").NullIfZero(), Description: "URL of the user profile image, original size."},
			{Name: "is_admin", Type: proto.ColumnType_BOOL, Description: "True if the user is an administrator of the current workspace."},
			{Name: "is_app_user", Type: proto.ColumnType_BOOL, Description: "True if the user is an owner of the current workspace."},
			{Name: "is_bot", Type: proto.ColumnType_BOOL, Description: "True if the user is a bot."},
			{Name: "is_invited_user", Type: proto.ColumnType_BOOL, Description: "True if the user joined the workspace via an invite."},
			{Name: "is_owner", Type: proto.ColumnType_BOOL, Description: "True if the user is an owner of the current workspace."},
			{Name: "is_primary_owner", Type: proto.ColumnType_BOOL, Description: "True if the user is the primary owner of the current workspace."},
			{Name: "is_restricted", Type: proto.ColumnType_BOOL, Description: "Indicates whether or not the user is a guest user. Use in combination with the is_ultra_restricted field to check if the user is a single-channel guest user."},
			{Name: "is_stranger", Type: proto.ColumnType_BOOL, Description: "If true, this user belongs to a different workspace than the one associated with your app's token, and isn't in any shared channels visible to your app. If false (or this field is not present), the user is either from the same workspace as associated with your app's token, or they are from a different workspace, but are in a shared channel that your app has access to. Read our shared channels docs for more detail."},
			{Name: "is_ultra_restricted", Type: proto.ColumnType_BOOL, Description: "Indicates whether or not the user is a single-channel guest."},
			{Name: "last_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.LastName").NullIfZero(), Description: "Last name of the user."},
			{Name: "locale", Type: proto.ColumnType_STRING, Description: "IETF language code for the user's chosen display language."},
			{Name: "job_title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Title").NullIfZero(), Description: "Job title of the user."},
			//{Name: "name", Type: proto.ColumnType_STRING, Description: "Don't use this. It once indicated the preferred username for a user, but that behavior has fundamentally changed since."},
			{Name: "phone", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Phone").NullIfZero(), Description: "Phone number of the user."},
			{Name: "profile_fields", Type: proto.ColumnType_JSON, Transform: transform.FromField("Profile.Fields"), Description: "Custom fields for the profile."},
			{Name: "real_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.RealName").NullIfZero(), Description: "The real name that the user specified in their workspace profile."},
			{Name: "real_name_normalized", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.RealNameNormalized").NullIfZero(), Description: "The real_name field, but with any non-Latin characters filtered out."},
			{Name: "skype", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Skype").NullIfZero(), Description: "Skype handle of the user."},
			{Name: "status_emoji", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.StatusEmoji").NullIfZero(), Description: "Status emoji the user has set."},
			{Name: "status_expiration", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Profile.StatusExpiration").Transform(intToTime), Description: "Expiration for the user status."},
			{Name: "status_text", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.StatusText").NullIfZero(), Description: "Status text the user has set."},
			{Name: "team_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Profile.Team").NullIfZero(), Description: "The team workspace that the user is a member of."},
			{Name: "tz", Type: proto.ColumnType_STRING, Transform: transform.FromField("TZ").NullIfZero(), Description: "A human-readable string for the geographic timezone-related region this user has specified in their account."},
			{Name: "tz_label", Type: proto.ColumnType_STRING, Transform: transform.FromField("TZLabel").NullIfZero(), Description: "Describes the commonly used name of the timezone."},
			{Name: "tz_offset", Type: proto.ColumnType_INT, Transform: transform.FromField("TZOffset").NullIfZero(), Description: "Indicates the number of seconds to offset UTC time by for this user's timezone."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Updated").Transform(jsonTimeToTime), Description: "Time when the user was last updated."},

			// These work, but are very heavy on the Slack API rate limit
			//{Name: "presence", Type: proto.ColumnType_STRING, Hydrate: getUserPresence},
			//{Name: "presence_online", Type: proto.ColumnType_BOOL, Hydrate: getUserPresence},
		}),
	}
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_user.listUsers", "connection_error", err)
		return nil, err
	}

	// Use 200 as default API limit, as recommended in the docs
	pageLimit := 200
	limit := int(d.QueryContext.GetLimit())

	if limit < pageLimit {
		pageLimit = limit
	}

	// Paginate ourselves instead of api.GetUsersContext to respect the query's limit
	var users []slack.User
	for (limit == -1 || len(users) < limit) && err == nil {
		p := api.GetUsersPaginated(slack.GetUsersOptionLimit(pageLimit))
		p, err = p.Next(ctx)
		if err == nil {
			users = append(users, p.Users...)
		}
	}

	if err != nil {
		plugin.Logger(ctx).Warn("slack_user.listUsers", "query_error", err)
		return nil, err
	}
	for _, user := range users {
		d.StreamListItem(ctx, user)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	api, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("slack_user.getUser", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	email := quals["email"].GetStringValue()

	var info *slack.User

	if len(id) > 0 {
		info, err = api.GetUserInfoContext(ctx, id)
	} else if len(email) > 0 {
		info, err = api.GetUserByEmailContext(ctx, email)
	} else {
		plugin.Logger(ctx).Warn("slack_user.getUser", "invalid_quals", "id and email both empty", "quals", quals)
		return nil, nil
	}

	if err != nil {
		if err.Error() == "user_not_found" || err.Error() == "users_not_found" {
			plugin.Logger(ctx).Warn("slack_user.getUser", "not_found_error", err, "quals", quals)
			return nil, nil
		}
		plugin.Logger(ctx).Error("slack_user.getUser", "query_error", err, "quals", quals)
		return nil, err
	}
	return info, nil
}

// func getUserPresence(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	api, err := connect(ctx, d)
// 	if err != nil {
// 		return nil, err
// 	}
// 	user := h.Item.(slack.User)
// 	for err == nil {
// 		p, err := api.GetUserPresence(user.ID)
// 		if err == nil {
// 			return p, nil
// 		}
// 		if rateLimitedError, ok := err.(*slack.RateLimitedError); ok {
// 			select {
// 			case <-ctx.Done():
// 				// Pass through
// 			case <-time.After(rateLimitedError.RetryAfter):
// 				err = nil
// 			}
// 		}
// 	}
// 	return nil, err
// }

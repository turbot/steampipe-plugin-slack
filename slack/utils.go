package slack

import (
	"context"
	"errors"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func connect(_ context.Context, d *plugin.QueryData) (*slack.Client, error) {
	token := os.Getenv("SLACK_TOKEN")

	slackConfig := GetConfig(d.Connection)
	if &slackConfig != nil {
		if slackConfig.Token != nil {
			token = *slackConfig.Token
		}
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	api := slack.New(token, slack.OptionDebug(false))
	return api, nil
}

func stringFloatToTime(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	timeFloat, err := strconv.ParseFloat(d.Value.(string), 64)
	if err != nil {
		return nil, err
	}
	if timeFloat == 0 {
		return nil, nil
	}
	sec, dec := math.Modf(timeFloat)
	t := time.Unix(int64(sec), int64(dec*(1e9)))
	return t, nil
}

func intToTime(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	i := int64(d.Value.(int))
	// Assume zero value (1970-01-01 00:00:00) means not set (null).
	if i == 0 {
		return nil, nil
	}
	return time.Unix(i, 0), nil
}

func jsonTimeToTime(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	jt := d.Value.(slack.JSONTime)
	// Assume zero value (1970-01-01 00:00:00) means not set (null).
	if jt == 0 {
		return nil, nil
	}
	return jt.Time(), nil
}

package slack

import (
	"context"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(_ context.Context, d *plugin.QueryData) (*slack.Client, error) {
	token := os.Getenv("SLACK_TOKEN")

	slackConfig := GetConfig(d.Connection)
	if slackConfig.Token != nil {
		token = *slackConfig.Token
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

func blockJsonToString(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	blocks := d.Value.(slack.Blocks)
	if blockString, err := parseBlocks(ctx, blocks); err != nil {
		return nil, err
	} else if blockString == "" {
		return nil, nil
	} else {
		return blockString, nil
	}
}

func attachmentJsonToString(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	textSeparator := "\n"
	var attachmentString string
	attachments := d.Value.([]slack.Attachment)
	for _, attachment := range attachments {
		attachmentString += attachment.Text + textSeparator
		if blockString, err := parseBlocks(ctx, attachment.Blocks); err != nil {
			return nil, err
		} else {
			attachmentString += blockString + textSeparator
		}
		if fieldString, err := parseFields(ctx, attachment.Fields); err != nil {
			return nil, err
		} else {
			attachmentString += fieldString + textSeparator
		}
	}

	if attachmentString == "" {
		return nil, nil
	} else {
		return strings.TrimSpace(attachmentString), nil
	}
}

func parseBlocks(ctx context.Context, blocks slack.Blocks) (string, error) {
	textSeparator := "\n"
	var blockString string
	for _, section := range blocks.BlockSet {
		switch blockType := section.BlockType(); blockType {
		case "section":
			plugin.Logger(ctx).Debug("parsing section block")
			sectionBlock := section.(*slack.SectionBlock)
			blockString += sectionBlock.Text.Text + textSeparator
			for _, textBlockObject := range sectionBlock.Fields {
				blockString += textBlockObject.Text + textSeparator
			}
		case "context":
			plugin.Logger(ctx).Debug("parsing context block")
			contextBlock := section.(*slack.ContextBlock)
			for _, element := range contextBlock.ContextElements.Elements {
				if elementType := element.MixedElementType(); elementType == "mixed_text" {
					text := element.(*slack.TextBlockObject).Text
					blockString += text + textSeparator
				}
			}
		}
	}
	return strings.TrimSpace(blockString), nil
}

func parseFields(ctx context.Context, fields []slack.AttachmentField) (string, error) {
	textSeparator := "\n"
	var fieldString string
	for _, field := range fields {
		fieldString += field.Value + textSeparator
	}
	return strings.TrimSpace(fieldString), nil
}

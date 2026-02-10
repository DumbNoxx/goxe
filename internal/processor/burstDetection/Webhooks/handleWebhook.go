package webhooks

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	pkg "github.com/DumbNoxx/goxe/pkg/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

func HandleWebhook(msg string, stats *pipelines.LogBurst) {
	var (
		data []byte
		err  error
	)

	for _, url := range options.Config.WebHookUrls {
		if strings.HasPrefix(url, "https://discord.com") {
			var DataSentWebhook pkg.WebhookDiscord
			var log = pkg.OptionsEmbedsDiscord{
				Title:       msg,
				Description: "The server's acting up.",
				Color:       16777215,
				Author: pkg.AuthorOptionsEmbedsDiscord{
					Name:    "Goxe",
					Url:     "https://github.com/DumbNoxx/Goxe",
					IconUrl: "https://raw.githubusercontent.com/DumbNoxx/Dotfiles-For-Humans/refs/heads/main/src/assets/img/goxe.png",
				},
				Fields: []pkg.FieldEmbedsDiscord{
					{
						Name:   "Errors",
						Value:  "```Check the server, it's overheating.```",
						Inline: false,
					},
					{
						Name:   "Category",
						Value:  stats.Category,
						Inline: true,
					},
					{
						Name:   "Start Time",
						Value:  stats.WindowStart.Format("02-01-2006, 15:04"),
						Inline: true,
					},
					{
						Name:   "Counts",
						Value:  fmt.Sprintf("%d", stats.Count),
						Inline: true,
					},
				},
				Footer: pkg.FooterEmbedsDiscord{
					Text: "Your Log Collector ❤️",
				},
				Timestamp: time.Now(),
			}
			DataSentWebhook.Embeds = append(DataSentWebhook.Embeds, log)
			data, err = json.Marshal(DataSentWebhook)
			sentData(data, err, url)
		}

		if strings.HasPrefix(url, "https://hooks.slack.com") {
			var headerLog = pkg.OptionsBlockSlack{
				Type: "header",
				Text: &pkg.OptionsTextMrkSlack{
					Type:  "plain_text",
					Text:  msg,
					Emoji: true,
				},
			}

			var mrkLog = pkg.OptionsBlockSlack{
				Type: "section",
				Text: &pkg.OptionsTextMrkSlack{
					Type: "mrkdwn",
					Text: fmt.Sprintf(
						"```Check the server, it's overheating.\nCount: %d - Start Time: %v - Category: %s```",
						stats.Count,
						stats.WindowStart.Format("02-01-2006, 15:04"),
						stats.Category,
					),
				},
			}

			var divider = pkg.OptionsBlockSlack{
				Type: "divider",
			}

			var footerLog = pkg.OptionsBlockSlack{
				Type: "context",
				Elements: []pkg.OptionsElementsBlockSlack{
					{
						Type:  "plain_text",
						Text:  "Author: Goxe",
						Emoji: true,
					},
				},
			}

			payload := pkg.WebhookSlack{
				Blocks: []pkg.OptionsBlockSlack{
					headerLog,
					mrkLog,
					divider,
					footerLog,
				},
			}

			data, err = json.Marshal(payload)
			sentData(data, err, url)
		}
	}
}

package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/latte-soft/discord-webhooks-go"
)

//go:embed buildit.png
var imgData []byte

func main() {
	args := os.Args
	argsLen := len(args)
	if argsLen < 2 || argsLen > 2 {
		fmt.Printf("USAGE: %s <WEBHOOK_URL>\n", args[0])
		os.Exit(1)
	}

	webhookUrl := args[1]

	messageId, err := discord.PostMessage(webhookUrl, &discord.Message{
		Embeds: &[]discord.Embed{
			{
				Color: 0x176d34,

				Author: &discord.EmbedAuthor{
					Name: "Author",
				},

				Title: "Title",
				Url:   "https://example.org",

				Fields: &[]discord.EmbedField{
					{
						Name:   "A field",
						Value:  "A value",
						Inline: true,
					},
					{
						Name:   "Another field",
						Value:  "Another value",
						Inline: true,
					},
				},

				Image: &discord.EmbedImage{
					Url: "attachment://buildit.png",
				},

				Footer: &discord.EmbedFooter{
					Text: "Footer",
				},

				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
		},

		Files: &[]discord.File{
			{
				Name: "buildit.png",
				Data: &imgData,
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(*messageId)
}

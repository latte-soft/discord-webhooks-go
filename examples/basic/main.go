package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"log"
	"os"
	"time"

	"github.com/cenkalti/dominantcolor"
	"github.com/latte-soft/discord-webhooks-go"
)

func main() {
	args := os.Args
	argsLen := len(args)
	if argsLen < 2 || argsLen > 2 {
		fmt.Printf("USAGE: %s <WEBHOOK_URL>", args[0])
		os.Exit(1)
	}

	webhookUrl := args[1]

	data, err := os.ReadFile("examples/basic/buildit.png")
	if err != nil {
		log.Fatalln(err)
	}

	img, _, _ := image.Decode(bytes.NewReader(data))
	c := dominantcolor.Find(img)
	embedColor := binary.BigEndian.Uint32([]byte{0, c.R, c.G, c.B})

	messageId, err := discord.PostMessage(webhookUrl, &discord.Message{
		Embeds: &[]discord.Embed{
			{
				Color: embedColor,

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
				Data: &data,
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(*messageId)
}

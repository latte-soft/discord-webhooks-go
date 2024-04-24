# Discord Webhooks for Go

High-level implementation of methods on Discord webhooks in Go, supporting multipart uploads

## Usage

Import the module into your existing Go project with:

```
go get github.com/latte-soft/discord-webhooks-go
```

## Example

To run from cloning the repository (replacing with your own Discord Webhook URL and token): `go run ./examples/basic "https://discord.com/api/webhooks/123/1234567890"`

[examples/basic/main.go](examples/basic/main.go)

![Screenshot of example below](repo-assets/example-screenshot.png)

```go
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
```

## License

See file: [LICENSE](LICENSE)

```
BSD 3-Clause License

Copyright (c) 2024 Latte Softworks <https://latte.to>

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

```

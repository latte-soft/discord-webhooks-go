package discord

// BSD 3-Clause License | Copyright (c) 2024 Latte Softworks <https://latte.to>

// https://discord.com/developers/docs/resources/webhook#execute-webhook
type Message struct {
	Files       *[]File      `json:"-"`
	QueryParams *QueryParams `json:"-"`

	Content    string `json:"content,omitempty"`
	Username   string `json:"username,omitempty"`
	AvatarUrl  string `json:"avatar_url,omitempty"`
	TTS        bool   `json:"tts,omitempty"`
	ThreadName string `json:"thread_name,omitempty"`

	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`

	Embeds *[]Embed `json:"embeds,omitempty"`
}

// For multipart/form-data uploads
type File struct {
	Name string
	Data *[]byte
}

// Because yeah
type QueryParams struct {
	//Wait     bool
	ThreadId string
}

type Embed struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	Color       uint32 `json:"color,omitempty"`

	Footer    *EmbedFooter    `json:"footer,omitempty"`
	Image     *EmbedImage     `json:"image,omitempty"`
	Thumbnail *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video     *EmbedVideo     `json:"video,omitempty"`
	Provider  *EmbedProvider  `json:"provider,omitempty"`
	Author    *EmbedAuthor    `json:"author,omitempty"`

	Fields *[]EmbedField `json:"fields,omitempty"`
}

type AllowedMentions struct {
	Parse *[]string `json:"parse,omitempty"`
	Users *[]string `json:"users,omitempty"`
	Roles *[]string `json:"roles,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconUrl      string `json:"icon_url,omitempty"`
	ProxyIconUrl string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

// Equivalent to `EmbedImage` struct spec
type EmbedThumbnail = EmbedImage
type EmbedVideo = EmbedImage

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	Url          string `json:"url,omitempty"`
	IconUrl      string `json:"icon_url,omitempty"`
	ProxyIconurl string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// https://discord.com/developers/docs/resources/webhook#edit-webhook-message
// Derivative of `discord.Message` though, yeah
type MessageEdit struct {
	Content string `json:"content,omitempty"`

	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`

	Embeds *[]Embed `json:"embeds,omitempty"`
}

// https://discord.com/developers/docs/resources/webhook#get-webhook-with-token
// (With no `user` obj) https://discord.com/developers/docs/resources/webhook#webhook-object
type WebhookInfo struct {
	Id        string `json:"id"`
	GuildId   string `json:"guild_id"`
	ChannelId string `json:"channel_id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Token     string `json:"token"`
	Url       string `json:"url"`

	// Application-owned webhooks only
	ApplicationId string `json:"application_id"`

	Type WebhookType `json:"type"`
}

// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
type WebhookType int

const (
	WebhookTypeIncoming        WebhookType = 1
	WebhookTypeChannelFollower WebhookType = 2
	WebhookTypeApplication     WebhookType = 3
)

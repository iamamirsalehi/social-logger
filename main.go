package main

import (
	"git.coryptex.com/sdk/discord-go/dicord"
)

func main() {
	discordLogger, _ := dicord.NewDiscordLogger([]string{
		"https://discord.com/api/webhooks/917377910271250462/kSf3ej70YXjcwYcM08dBsUPfbBkqEwJ1nNbmHTyX3DbniLlmrzjY7us8W-QWhjqDEpg7",
	})

	_ = discordLogger.Info("key", discordLogger, "asjdjasd", "", "description", "ajsdhkajdhjkashdjk")
}

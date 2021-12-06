package main

import (
	"fmt"
	"git.coryptex.com/sdk/discord-go/dicord"
)

func main() {
	discordLogger, err := dicord.NewSocialLogger([]string{
		"https://discord.com/api/webhooks/917325600807592017/ROcL7e9Hp98M9nIZ-byGZoSNyKQ6kTmCViF5GC8re4Xej_k7GdGu9EJ3CREE4GnIsyTQ",
		"https://discord.com/api/webhooks/917377910271250462/kSf3ej70YXjcwYcM08dBsUPfbBkqEwJ1nNbmHTyX3DbniLlmrzjY7us8W-QWhjqDEpg7",
	})

	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = discordLogger.Info("key 1", "val 1", "key 2", "val 2", "color", "99999", "description", "This is a description", "content", "this is a content")

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

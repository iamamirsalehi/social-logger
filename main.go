package main

import (
	"fmt"
	"github.com/iamamirsalehi/social-logger/dicord"
)

func main() {
	var webhooks []string

	webhooks = append(webhooks, "https://discord.com/api/webhooks/917325600807592017/ROcL7e9Hp98M9nIZ-byGZoSNyKQ6kTmCViF5GC8re4Xej_k7GdGu9EJ3CREE4GnIsyTQ")
	webhooks = append(webhooks, "https://discord.com/api/webhooks/917377910271250462/kSf3ej70YXjcwYcM08dBsUPfbBkqEwJ1nNbmHTyX3DbniLlmrzjY7us8W-QWhjqDEpg7")

	discordLogger, err := dicord.NewSocialLogger(webhooks)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	msgs := make(map[string]interface{})
	msgs = map[string]interface{}{
		"key 1":       "val 1, val 1, val 1",
		"key 2":       "val 2, val 2, val 2",
		"key 3":       "val 3, val 3, val 3",
		"key 4":       "val 4, val 4, val 4",
		"key 5":       "val 5, val 5, val 5",
		"description": "This is a description, This is a description, This is a description",
		"color":       6468445,
		"content":     "This is a content, This is a content, This is a content",
	}
	err = discordLogger.Info(msgs)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

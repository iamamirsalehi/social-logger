package main

import (
	"fmt"
	"github.com/iamamirsalehi/social-logger/dicord"
)

func main() {
	discordLogger, err := dicord.NewSocialLogger([]string{
		"https://discord.com/api/webhooks/917325600807592017/ROcL7e9Hp98M9nIZ-byGZoSNyKQ6kTmCViF5GC8re4Xej_k7GdGu9EJ3CREE4GnIsyTQ",
		"https://discord.com/api/webhooks/917377910271250462/kSf3ej70YXjcwYcM08dBsUPfbBkqEwJ1nNbmHTyX3DbniLlmrzjY7us8W-QWhjqDEpg7",
	})

	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = discordLogger.Info("key 1", "val 1", "key 2", "val 2", "key 3", "val 3", "description", "Midoni chara?")

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

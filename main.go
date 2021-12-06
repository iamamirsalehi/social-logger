package main

import (
	"fmt"
	"github.com/iamamirsalehi/social-logger/dicord"
)

func main() {
	var webhooks []string

	webhooks = append(webhooks, "https://discord.com/api/webhooks/917325600807592017/ROcL7e9Hp98M9nIZ-byGZoSNyKQ6kTmCViF5GC8re4Xej_k7GdGu9EJ3CREE4GnIsyTQ")

	discordLogger, err := dicord.NewSocialLogger(webhooks)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	msgs := make(map[string]interface{})

	msgs["بگایی"] = "بدویید بگا رفتیم"

	err = discordLogger.Info(msgs)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

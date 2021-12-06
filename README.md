# introduction

This package divided into three different packages witch includes **Telegram**, **Slack** and **Discord**. Let's begin with Discord

## Discord

To use discord you need to initial it first, like:

```go
import (
    "fmt"
    "git.coryptex.com/sdk/discord-go/dicord"
)

discordLogger, err := dicord.NewSocialLogger([]string{
    "https://discord.com/api/webhooks/917325600807592017/ROcL7e9Hp98M9nIZ-byGZoSNyKQ6kTmCViF5GC8re4Xej_k7GdGu9EJ3CREE4GnIsyTQ",
    "https://discord.com/api/webhooks/917377910271250462/kSf3ej70YXjcwYcM08dBsUPfbBkqEwJ1nNbmHTyX3DbniLlmrzjY7us8W-QWhjqDEpg7",
})

if err != nil {
    fmt.Println("Error: ", err)
}
```
The ``discordLogger`` variable contains all the methods you need to call to send your message to the discord channel. The following methods are:

```go
discordLogger.Info("key 1", "val 1", "key 2", "val 2",)

discordLogger.Debug("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Warn("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Error("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Fatal("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Trace("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Panic("key 1", "val 1", "key 2", "val 2",) 
```

**Note**: You have to pass **key** and **value** after another. In some cases if you didn't pass the **value** of a key the value will be an empty string.
#### How to set color?
If you want to set a custom color you can pass a key named ``color`` and the value must be a color number, like: ``99999`` which is blue
```go
err = discordLogger.Info("key 1", "val 1", "key 2", "val 2", "color", "99999")

if err != nil {
	fmt.Println("Error: ", err)
}
```
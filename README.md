#introduction

This package divided into three different packages witch includes **Telegram**, **Slack** and **Discord**. Let's begin with Discord

## Discord

To use discord you need to initial it first, like:

```go
import (
    "fmt"
    "github.com/iamamirsalehi/social-logger/dicord"
)

discordLogger, err := dicord.NewSocialLogger([]string{
    "https://discord.com/api/webhooks/917325600807592017/ROcL7e9Hp98M9nIZ-byGZoSNyKQ6kTmCViF5GC8re4Xej_k7GdGu9EJ3CREE4GnIsyTQ",
    "https://discord.com/api/webhooks/917377910271250462/kSf3ej70YXjcwYcM08dBsUPfbBkqEwJ1nNbmHTyX3DbniLlmrzjY7us8W-QWhjqDEpg7",
})

if err != nil {
    fmt.Println("Error: ", err)
}
```
The ``discordLogger`` variable contains all the methods you need to call and send your message to discord. The following methods are its

```go
discordLogger.Info("key 1", "val 1", "key 2", "val 2",)

discordLogger.Debug("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Warn("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Error("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Fatal("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Trace("key 1", "val 1", "key 2", "val 2",) 

discordLogger.Panic("key 1", "val 1", "key 2", "val 2",) 
```

**Note**: you have to pass **key** and **value** after another. In some cases if you didn't pass the **value** of a key the value will be an empty string.
#### How to set color?

```go
err = discordLogger.Info("key 1", "val 1", "key 2", "val 2", )

if err != nil {
	fmt.Println("Error: ", err)
}
```
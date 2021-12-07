# introduction

This package divided into three different packages which includes **Telegram**, **Slack** and **Discord**. Let's begin with Discord

## Discord

To use discord you need to initial it first, like:

```go
import (
    "fmt"
    "git.coryptex.com/sdk/discord-go/dicord"
)

discordLogger, err := dicord.NewSocialLogger([]string{
    "https://discord.com/api/webhooks/your_webhook",
    "https://discord.com/api/webhooks/your_webhook",
})

if err != nil {
    fmt.Println("Error: ", err)
}
```
### How to send message to discord?
```go
err := discordLogger.Info("key 1", "val 1")

if err != nil {
	fmt.Println("Error: ", err)
}
```
Your message sent! :)

The ``discordLogger`` variable contains all the methods you need to call to send your message to the discord channel. The following methods are:

```go
discordLogger.Info("key 1", "val 1", "key 2", "val 2")

discordLogger.Debug("key 1", "val 1", "key 2", "val 2") 

discordLogger.Warn("key 1", "val 1", "key 2", "val 2") 

discordLogger.Error("key 1", "val 1", "key 2", "val 2") 

discordLogger.Fatal("key 1", "val 1", "key 2", "val 2") 

discordLogger.Trace("key 1", "val 1", "key 2", "val 2") 

discordLogger.Panic("key 1", "val 1", "key 2", "val 2") 
```

**Note**: You have to pass **key** and **value** after another. In some cases if you didn't pass the **value** of a key the value will be an empty string.
#### How to set a custom color?
If you want to set a custom color you can pass a key named ``color`` and the value must be a color number, like: ``99999`` which is blue
```go
err := discordLogger.Info("key 1", "val 1", "key 2", "val 2", "color", "99999")

if err != nil {
	fmt.Println("Error: ", err)
}
```
#### How to set custom description?
Simply as above code! you just need to pass ``description`` as key and pass the value:

```go
err := discordLogger.Info("key 1", "val 1", "key 2", "val 2", "color", "99999", "description", "This is a description")

if err != nil {
	fmt.Println("Error: ", err)
}
```

#### How to set custom content?
need to pass ``content`` as key and pass the value:

```go
err := discordLogger.Info("key 1", "val 1", "key 2", "val 2", "color", "99999", "description", "This is a description", "content", "this is a content")

if err != nil {
	fmt.Println("Error: ", err)
}
```

Have fun!
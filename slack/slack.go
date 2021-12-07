package slack

type Slack interface {
	Info(keyVal ...string) error
	Debug(keyVal ...string) error
	Warn(keyVal ...string) error
	Error(keyVal ...string) error
	Fatal(keyVal ...string) error
	Trace(keyVal ...string) error
	Panic(keyVal ...string) error
}

//func New

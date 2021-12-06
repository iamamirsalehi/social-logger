package slack

type Slack interface {
	Info(map[string]interface{}) error
	/*	Debug(map[string]interface{}) error
		Warn(map[string]interface{}) error
		Error(map[string]interface{}) error
		Fatal(map[string]interface{}) error
		Trace(map[string]interface{}) error
		Panic(map[string]interface{}) error*/
}

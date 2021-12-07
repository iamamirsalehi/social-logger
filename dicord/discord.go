package dicord

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Discord interface {
	Info(keyVal ...string) error
	Debug(keyVal ...string) error
	Warn(keyVal ...string) error
	Error(keyVal ...string) error
	Fatal(keyVal ...string) error
	Trace(keyVal ...string) error
	Panic(keyVal ...string) error
}

const (
	InfoColor  = 3447003
	DebugColor = 15105570
	WarnColor  = 16776960
	ErrorColor = 15158332
	FatalColor = 10181046
	TraceColor = 9807270
	PanicColor = 10038562
)

func NewDiscordLogger(webhooks []string) (Discord, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: config,
	}

	netClient := &http.Client{
		Transport: transport,
	}

	return &socialLogger{
		webhooks:  webhooks,
		netClient: netClient,
	}, nil
}

type socialLogger struct {
	webhooks  []string
	netClient *http.Client
}

type Fields struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Embeds struct {
	Description string    `json:"description,omitempty"`
	Color       int       `json:"color,omitempty"`
	Fields      []*Fields `json:"fields,omitempty"`
}

type Params struct {
	Content interface{} `json:"content,omitempty"`
	Embeds  []*Embeds   `json:"embeds"`
}

func (s socialLogger) Info(keyVal ...string) error {
	err := sendMessage(&s, keyVal, InfoColor)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Debug(keyVal ...string) error {
	err := sendMessage(&s, keyVal, DebugColor)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Warn(keyVal ...string) error {
	err := sendMessage(&s, keyVal, WarnColor)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Error(keyVal ...string) error {
	err := sendMessage(&s, keyVal, ErrorColor)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Fatal(keyVal ...string) error {
	err := sendMessage(&s, keyVal, FatalColor)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Trace(keyVal ...string) error {
	err := sendMessage(&s, keyVal, TraceColor)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Panic(keyVal ...string) error {
	err := sendMessage(&s, keyVal, PanicColor)

	if err != nil {
		return err
	}

	return nil
}

func sendMessage(s *socialLogger, data []string, color int) error {
	var wg sync.WaitGroup

	for i := 0; i < len(s.webhooks); i++ {
		wg.Add(1)

		go func(i int, webhooks []string, data []string) {
			defer wg.Done()

			_, _ = s.netClient.Post(webhooks[i], "application/json", prepareData(data, color))

		}(i, s.webhooks, data)
	}

	wg.Wait()

	return nil
}

func prepareData(keyVal []string, color int) *bytes.Buffer {
	var fields []*Fields
	var content string
	description := ""
	keyValLen := len(keyVal)
	key := 0

	for i := 0; i < keyValLen; i += 2 {

		if keyValLen == i+1 {
			break
		}

		if i%2 == 0 {
			key = i
		}

		switch strings.ToUpper(keyVal[key]) {
		case "DESCRIPTION":
			description = keyVal[key+1]
			continue
		case "COLOR":
			color, _ = strconv.Atoi(keyVal[key+1])
			continue
		case "CONTENT":
			content = keyVal[key+1]
			continue
		}

		fields = append(fields, &Fields{
			Name:  keyVal[key],
			Value: keyVal[key+1],
		})
	}

	var embeds []*Embeds

	embeds = append(embeds, &Embeds{
		Description: fmt.Sprintf("%v", description),
		Color:       color,
		Fields:      fields,
	})

	postBody, _ := json.Marshal(Params{
		Content: content,
		Embeds:  embeds,
	})

	responseBody := bytes.NewBuffer(postBody)

	return responseBody
}

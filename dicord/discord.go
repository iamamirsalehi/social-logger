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
	Info  = 3447003
	Debug = 15105570
	Warn  = 16776960
	Error = 15158332
	Fatal = 10181046
	Trace = 9807270
	Panic = 10038562
)

func NewSocialLogger(webhooks []string) (Discord, error) {
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
	err := sendMessage(&s, keyVal, Info)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Debug(keyVal ...string) error {
	err := sendMessage(&s, keyVal, Debug)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Warn(keyVal ...string) error {
	err := sendMessage(&s, keyVal, Warn)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Error(keyVal ...string) error {
	err := sendMessage(&s, keyVal, Error)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Fatal(keyVal ...string) error {
	err := sendMessage(&s, keyVal, Fatal)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Trace(keyVal ...string) error {
	err := sendMessage(&s, keyVal, Trace)

	if err != nil {
		return err
	}

	return nil
}

func (s socialLogger) Panic(keyVal ...string) error {
	err := sendMessage(&s, keyVal, Panic)

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

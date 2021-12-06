package dicord

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Discord interface {
	Info(map[string]interface{}) error
	/*	Debug(map[string]interface{}) error
		Warn(map[string]interface{}) error
		Error(map[string]interface{}) error
		Fatal(map[string]interface{}) error
		Trace(map[string]interface{}) error
		Panic(map[string]interface{}) error*/
}

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

func (s socialLogger) Info(m map[string]interface{}) error {
	var wg sync.WaitGroup

	for i := 0; i < len(s.webhooks); i++ {
		wg.Add(1)

		go func(i int, m map[string]interface{}, webhooks []string) {
			defer wg.Done()

			_, _ = s.netClient.Post(webhooks[i], "application/json", prepareData(m))

		}(i, m, s.webhooks)
	}

	wg.Wait()

	return nil
}

func prepareData(m map[string]interface{}) *bytes.Buffer {
	var fields []*Fields

	var content, color, description interface{}

	for name, value := range m {
		switch strings.ToUpper(name) {
		case "DESCRIPTION":
			description = value
			continue
		case "COLOR":
			color = value
			continue
		case "CONTENT":
			content = value
			continue
		}
		fields = append(fields, &Fields{
			Name:  name,
			Value: value,
		})
	}

	var embeds []*Embeds

	embeds = append(embeds, &Embeds{
		Description: fmt.Sprintf("%v", description),
		Color:       color.(int),
		Fields:      fields,
	})

	postBody, _ := json.Marshal(Params{
		Content: content,
		Embeds:  embeds,
	})

	responseBody := bytes.NewBuffer(postBody)

	return responseBody
}

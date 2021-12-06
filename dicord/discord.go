package dicord

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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

	responseBody := prepareData(m)
	for _, webhook := range s.webhooks {

		resp, err := s.netClient.Post(webhook, "application/json", responseBody)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			defer resp.Body.Close()

			responseBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			return errors.New(string(responseBody))
		}
	}

	return nil
}

func prepareData(m map[string]interface{}) *bytes.Buffer {
	var fields []*Fields

	for name, value := range m {
		fields = append(fields, &Fields{
			Name:  name,
			Value: value,
		})
	}

	var embeds []*Embeds

	embeds = append(embeds, &Embeds{
		Description: "I am a good boy",
		Color:       16007990,
		Fields:      fields,
	})

	postBody, _ := json.Marshal(Params{
		Content: "Mooz",
		Embeds:  embeds,
	})

	responseBody := bytes.NewBuffer(postBody)

	return responseBody
}
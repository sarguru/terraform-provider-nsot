package nsot

import (
	"fmt"
	"log"

	"github.com/sarguru/go-nsot-api"
)

type Config struct {
	Email  string
	Secret string
	Url    string
}

func (c *Config) Client() (*nsot.Client, error) {
	client, err := nsot.NewClient(c.Email, c.Secret, c.Url)

	if err != nil {
		return nil, fmt.Errorf("Error setting up client: %s", err)
	}

	log.Printf("[INFO] NSOT Client configured for user: %s", client.Email)

	return client, nil
}

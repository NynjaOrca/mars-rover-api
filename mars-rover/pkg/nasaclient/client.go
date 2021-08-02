package nasaclient

import (
	"mars-rover-api/mars-rover/pkg/config"
	"net/http"
	"time"
)

type Client struct {
	Config config.APIConfig
	Client http.Client
}

func New(c config.APIConfig) Client {
	return Client{
		Config: c,
		Client: http.Client{
			Timeout: time.Duration(c.Timeout) * time.Millisecond,
		},
	}
}

func (c *Client) GetRoverImages() {

}

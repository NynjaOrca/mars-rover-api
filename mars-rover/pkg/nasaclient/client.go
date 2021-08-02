package nasaclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mars-rover-api/mars-rover/pkg/config"
	"mars-rover-api/mars-rover/pkg/models"
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

func (c *Client) GetRoverImages(earthDate, camera string) error {
	params := fmt.Sprintf("?earth_date=%s&camera=%s&api_key=%s", earthDate, camera, c.Config.Key)
	url := c.Config.URL + c.Config.RelUrl + params
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("unable to build request to NASA: %v", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("call to NASA failed: %v", err)
	}

	if resp == nil {
		return fmt.Errorf("response from NASA is nil")
	}

	if resp.StatusCode != http.StatusOK {
		var badResp map[string]interface{}
		err = json.Unmarshal(readBody(resp), &badResp)
		if err != nil {
			return fmt.Errorf("unable to unmarshal response body after call failed: %v", err)
		}

		return fmt.Errorf("call to NASA failed: %v", badResp)
	}

	xRateLimitRemaining := resp.Header.Get("X-RateLimit-Remaining")
	log.Println("xRateLimitRemaining:", xRateLimitRemaining)

	var nasaResponse models.NasaResponse
	err = json.Unmarshal(readBody(resp), &nasaResponse)
	if err != nil {
		return fmt.Errorf("unable to unmarshal response body after call succeeded: %v", err)
	}

	log.Printf("GetRoverImages(): %+v", nasaResponse)

	return nil
}

func readBody(r *http.Response) []byte {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("unable to read response body: %v", err)
		return nil
	}

	_ = r.Body.Close()
	return b
}

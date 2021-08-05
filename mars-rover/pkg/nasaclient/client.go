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

func (c *Client) GetRoverImages(earthDate, camera string) []string {
	images := make([]string, 0)
	params := fmt.Sprintf("?earth_date=%s&camera=%s&api_key=%s", earthDate, camera, c.Config.Key)
	req, err := http.NewRequest(http.MethodGet, c.Config.URL+c.Config.RelUrl+params, nil)
	if err != nil {
		log.Println("error building request:", err)
		return images
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Println("error calling client:", err)
		return images
	}

	if resp == nil {
		log.Println("error response is nil:", resp.Body)
		return images
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("error response status code:", resp.StatusCode)
		return images
	}

	xRateLimitRemaining := resp.Header.Get("X-RateLimit-Remaining")
	log.Println("xRateLimitRemaining:", xRateLimitRemaining)

	var nasaResponse models.NasaResponse
	err = json.Unmarshal(readBody(resp), &nasaResponse)
	if err != nil {
		log.Println("error unmarshalling:", err)
		return images
	}

	for _, p := range nasaResponse.Photos {
		if p.ImgSrc != "" {
			images = append(images, p.ImgSrc)
			if len(images) == 3 {
				break
			}
		}
	}
	return images
}

func readBody(r *http.Response) []byte {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("unable to read response body:", err)
		return nil
	}

	_ = r.Body.Close()
	return b
}

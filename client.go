package mailchimp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Initial configuration
const (
	username      = "mailchimp.go"
	apiHostFormat = "https://%s.api.mailchimp.com/3.0/"
)

// ClientConfig object used for client creation
type ClientConfig struct {
	APIHost string
	APIKey  string
	Debug   bool
}

// NewClientConfig constructs a ClientConfig object with the environment variables set as default
func NewClientConfig() ClientConfig {
	return ClientConfig{
		APIHost: apiHost,
		APIKey:  apiKey,
	}
}

// Client object
type Client struct {
	config ClientConfig

	Transport http.RoundTripper
}

// GetAPIHost returns the corresponding API Host for Mailchimp's API based on the provided Key
func GetAPIHost(apiKey string) string {
	parts := strings.Split(apiKey, "-")
	if len(parts) > 0 {
		return fmt.Sprintf(apiHostFormat, parts[len(parts)-1])
	}
	return ""
}

// New constructor from API Key
func New(apiKey string) Client {
	return Client{
		config: ClientConfig{
			APIHost: GetAPIHost(apiKey),
			APIKey:  apiKey,
		},
	}
}

// Request executes any Mailchimp API method using the current client configuration
func (c Client) Request(method, endpoint string, params url.Values, data, response interface{}) error {
	// Parse URL
	base, err := url.Parse(c.config.APIHost)
	if err != nil {
		return err
	}
	base.Path = path.Join(base.Path, endpoint)
	// Handle root path redirect
	if endpoint == "" || endpoint == "/" {
		base.Path += "/"
	}

	// Parse params
	if params != nil {
		for k := range params {
			base.Query().Set(k, params.Get(k))
		}
	}

	// Parse data
	var eData []byte
	if data != nil {
		eData, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	// Create request
	req, err := http.NewRequest(method, base.String(), bytes.NewBuffer(eData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Set Auth
	req.SetBasicAuth(username, c.config.APIKey)

	// Debug
	if c.config.Debug {
		log.Printf("NEW REQUEST TO %s", base.String())
	}

	// Perform request
	client := &http.Client{Transport: c.Transport}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Debug
	if c.config.Debug {
		log.Printf("RESPONSE FROM %s: \n%s", base.String(), body)
	}

	// Handle API errors
	if resp.StatusCode >= 400 {
		errResp := ErrorResponse{}
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return err
		}

		err = errResp
	} else {
		// Unmarshal into response
		if len(body) > 0 {
			err = json.Unmarshal(body, response)
		}
	}

	return err
}

// IndexResponse data
type IndexResponse map[string]interface{}

// Index method: Get links to all other resources available in the API.
func (c Client) Index() (IndexResponse, error) {
	var r IndexResponse
	r = make(map[string]interface{})

	err := c.Request("GET", "/", nil, nil, &r)
	return r, err
}

// Lists returns a Lists API client
func (c Client) Lists() Lists {
	return Lists{
		Client: c,
	}
}

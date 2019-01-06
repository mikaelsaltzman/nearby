package places

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NewClient creates a new Client.
func NewClient(url string, client Getter) Client {
	return Client{
		BaseURL:    url,
		HTTPClient: client,
	}
}

// Client is the struct containing the URL and the http client used to make Google Places API requests.
type Client struct {
	BaseURL    string
	HTTPClient Getter
}

// Getter requires the Get method, which is identical to the http.Client method Get.
type Getter interface {
	Get(string) (*http.Response, error)
}

// HTTPClient creates a new Client.
type HTTPClient struct {
	HTTPClient *http.Client
}

// Get makes a Get request using the http client wrapped in the custom http client struct.
func (c HTTPClient) Get(url string) (resp *http.Response, err error) {
	return c.HTTPClient.Get(url)
}

// Response is the struct that the top level Nearby Places API response is decoded into.
type Response struct {
	Results []Result `json:"results"`
}

// Result is the used by the Results slice in the Response struct.
type Result struct {
	Name     string   `json:"name"`
	Geometry Geometry `json:"geometry"`
}

// Geometry contains the coordinates for the result's place.
type Geometry struct {
	Location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
	} `json:"location"`
}

// Params are the key value pairs used by the Google Places API to make a request. These are set by URL parameters, and do also have default values, except for key, which is required.
type Params struct {
	Key       string
	Location  string
	PlaceType string
}

// Nearby is a Client method that makes the Google Places Nearby API request using the provided parameters and the hardcoded radius of 2 km.
func (c *Client) Nearby(ctx context.Context, params Params) (Response, error) {
	r, err := c.HTTPClient.Get(fmt.Sprintf("%s?key=%s&location=%s&type=%s&radius=2000", c.BaseURL, params.Key, params.Location, params.PlaceType))
	if err != nil {
		return Response{}, err
	}

	defer r.Body.Close()
	var res Response
	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return Response{}, err
	}

	return res, nil
}

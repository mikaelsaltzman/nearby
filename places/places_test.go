package places

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockHTTPClient struct {
	response *http.Response
}

func NewMockHTTPClient(r *http.Response) mockHTTPClient {
	return mockHTTPClient{r}
}

func (m mockHTTPClient) Get(url string) (resp *http.Response, err error) {
	return m.response, nil
}

func TestNearby(t *testing.T) {
	m := NewMockHTTPClient(&http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(`
			{
			   "results" : [
			      {
			         "geometry" : {
			            "location" : {
			               "lat" : 11.1111111,
			               "lng" : 22.2222222
			            }
			         },
			         "name" : "The Place",
			         "opening_hours" : {
			            "open_now" : false
			         }
			      },
			      {
			         "geometry" : {
			            "location" : {
			               "lat" : 22.2222222,
			               "lng" : 33.3333333
			            }
			         },
			         "name" : "The Second Place",
			         "opening_hours" : {
			            "open_now" : true 
			         }
			      }
			   ]
			}
		`)),
	})
	c := NewClient("", m)

	nearbyPlaces, err := c.Nearby(Params{})
	if err != nil {
		t.Error()
	}

	if got, want := len(nearbyPlaces.Results), 2; got != want {
		t.Fatalf("got len(nearbyPlaces) = %v, expected %v", got, want)
	}

	if got, want := nearbyPlaces.Results[0].Name, "The Place"; got != want {
		t.Errorf("got len(nearbyPlaces) = %v, expected %v", got, want)
	}

	if got, want := nearbyPlaces.Results[1].Name, "The Second Place"; got != want {
		t.Errorf("got len(nearbyPlaces) = %v, expected %v", got, want)
	}

}

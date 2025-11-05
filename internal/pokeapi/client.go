package pokeapi

import (
	"encoding/json"
	"net/http"
)

type Client interface {
	ListLocationAreas(url string) (ListResp,error)
}


type HTTPClient struct {
	base 	string
	h 		*http.Client	
}

func New(base string) *HTTPClient {
	return &HTTPClient {
		base : base,
		h: http.DefaultClient,
	}
}

func (c *HTTPClient) ListLocationAreas(url string) (ListResp, error) {
	if url == "" {
		url = c.base + "/location-area"
	}

	res, err := c.h.Get(url)
	if err != nil {
		return ListResp{}, err
	}

	defer res.Body.Close()
	var LocationAreas ListResp
	if err := json.NewDecoder(res.Body).Decode(&LocationAreas); err != nil {
		return ListResp{}, err
	}

	return LocationAreas, nil

}
package dkan

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/pkg/errors"
)

type API struct {
	endpoint *url.URL
	mu       sync.Locker
}

func NewAPI(endpoint string) (*API, error) {
	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &API{
		endpoint: url,
		mu:       &sync.Mutex{},
	}, nil
}

func (api *API) ObtainResource(res *Resource) (map[string]interface{}, error) {
	api.mu.Lock()

	values := api.endpoint.Query()
	values.Add("resource_id", res.id)
	values.Add("offset", strconv.Itoa(int(res.offset)))
	values.Add("limit", strconv.Itoa(int(res.limit)))

	api.endpoint.RawQuery = values.Encode()

	r, err := http.Get(api.endpoint.String())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	m := map[string]interface{}{}

	if err = json.Unmarshal(data, &m); err != nil {
		return nil, errors.WithStack(err)
	}

	api.endpoint.Query().Del("resource_id")
	api.endpoint.Query().Del("offset")
	api.endpoint.Query().Del("limit")

	api.mu.Unlock()

	return m, nil
}

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
	}, nil
}

func (api *API) ObtainResource(res *Resource) (map[string]interface{}, error) {
	api.mu.Lock()

	api.endpoint.Query().Add("resource_id", res.id)
	api.endpoint.Query().Add("offset", strconv.Itoa(int(res.offset)))
	api.endpoint.Query().Add("limit", strconv.Itoa(int(res.limit)))

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

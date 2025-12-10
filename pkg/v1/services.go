package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type (
	Service struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Extension string `json:"extension"`
		CreatedAt string `json:"created_at"`
	}

	ServicesFilters struct {
		Name string
	}

	ServicesQueryParams struct {
		Filters ServicesFilters
	}
)

func (q *ServicesQueryParams) queryParamsRaw() string {
	vals := url.Values{}
	if q == nil {
		return vals.Encode()
	}

	var filters []string

	if q.Filters.Name != "" {
		filters = append(filters, "name="+q.Filters.Name)
	}

	if len(filters) != 0 {
		vals.Set("filters", strings.Join(filters, ","))
	}

	return vals.Encode()
}

func (client *ServiceClient) ListServices(ctx context.Context, options *ServicesQueryParams) (*[]Service, *ResponseResult, error) {
	queryParams := ""
	if qRaw := options.queryParamsRaw(); options != nil && qRaw != "" {
		queryParams = "?" + qRaw
	}

	path := fmt.Sprintf("%s/services%s", client.APIUrl, queryParams)
	responseResult, err := client.DoRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}
	var response []Service
	err = responseResult.ExtractResult(&response)
	if err != nil {
		return nil, responseResult, err
	}

	return &response, responseResult, nil
}

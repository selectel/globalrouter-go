package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type (
	Quota struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Scope      string `json:"scope"`
		ScopeValue string `json:"scope_value"`
		Limit      int    `json:"limit"`
	}

	QuotasFilters struct {
		Name       string
		Scope      string
		ScopeValue string
	}

	QuotasQueryParams struct {
		Filters QuotasFilters
	}
)

func (q *QuotasQueryParams) queryParamsRaw() string {
	vals := url.Values{}
	if q == nil {
		return vals.Encode()
	}

	var filters []string

	if q.Filters.Name != "" {
		filters = append(filters, "name="+q.Filters.Name)
	}

	if q.Filters.Scope != "" {
		filters = append(filters, "scope="+q.Filters.Scope)
	}

	if q.Filters.ScopeValue != "" {
		filters = append(filters, "scope_value="+q.Filters.ScopeValue)
	}
	if len(filters) != 0 {
		vals.Set("filters", strings.Join(filters, ","))
	}

	return vals.Encode()
}

func (client *ServiceClient) ListQuotas(ctx context.Context, options *QuotasQueryParams) (*[]Quota, *ResponseResult, error) {
	queryParams := ""
	if qRaw := options.queryParamsRaw(); options != nil && qRaw != "" {
		queryParams = "?" + qRaw
	}

	path := fmt.Sprintf("%s/quotas%s", client.APIUrl, queryParams)
	responseResult, err := client.DoRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}
	var response []Quota
	err = responseResult.ExtractResult(&response)
	if err != nil {
		return nil, responseResult, err
	}

	return &response, responseResult, nil
}

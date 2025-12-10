package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type (
	Router struct {
		ID             string   `json:"id"`
		Name           string   `json:"name"`
		Status         string   `json:"status"`
		CreatedAt      string   `json:"created_at"`
		UpdatedAt      string   `json:"updated_at,omitempty"`
		Enabled        bool     `json:"enabled"`
		Tags           []string `json:"tags"`
		AccountID      string   `json:"account_id"`
		ProjectID      string   `json:"project_id,omitempty"`
		LeakUUID       string   `json:"leak_uuid,omitempty"`
		NetopsRouterID string   `json:"netops_router_id"`
		PrefixPoolID   string   `json:"prefix_pool_id"`
		VpnID          int      `json:"vpn_id"`
	}

	RouterCreateRequest struct {
		Name string   `json:"name,omitempty"`
		Tags []string `json:"tags,omitempty"`
	}

	RouterUpdateRequest struct {
		Name *string   `json:"name,omitempty"`
		Tags *[]string `json:"tags,omitempty"`
	}

	RoutersFilters struct {
		Name string
	}

	RoutersQueryParams struct {
		Filters RoutersFilters
	}
)

func (q *RoutersQueryParams) queryParamsRaw() string {
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

func (client *ServiceClient) Routers(ctx context.Context, q *RoutersQueryParams) (*[]Router, *ResponseResult, error) {
	queryParams := ""
	if qRaw := q.queryParamsRaw(); q != nil && qRaw != "" {
		queryParams = "?" + qRaw
	}

	u := fmt.Sprintf("%s/routers%s", client.APIUrl, queryParams)

	responseResult, err := client.DoRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var response []Router
	err = responseResult.ExtractResult(&response)
	if err != nil {
		return nil, responseResult, err
	}

	return &response, responseResult, nil
}

func (client *ServiceClient) Router(ctx context.Context, routerID string) (*Router, *ResponseResult, error) {
	u := fmt.Sprintf("%s/routers/%s", client.APIUrl, routerID)

	responseResult, err := client.DoRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var result *Router
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) RouterCreate(ctx context.Context, req *RouterCreateRequest) (*Router, *ResponseResult, error) {
	u := fmt.Sprintf("%s/routers", client.APIUrl)

	body, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}

	responseResult, err := client.DoRequest(ctx, http.MethodPost, u, bytes.NewReader(body))
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var result *Router
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) RouterUpdate(ctx context.Context, routerID string, req *RouterUpdateRequest) (*Router, *ResponseResult, error) {
	u := fmt.Sprintf("%s/routers/%s", client.APIUrl, routerID)

	body, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}

	responseResult, err := client.DoRequest(ctx, http.MethodPut, u, bytes.NewReader(body))
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var result *Router
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) RouterDelete(ctx context.Context, routerID string) (*ResponseResult, error) {
	u := fmt.Sprintf("%s/routers/%s", client.APIUrl, routerID)

	responseResult, err := client.DoRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}

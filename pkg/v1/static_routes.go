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
	StaticRoute struct {
		ID                  string   `json:"id"`
		Name                string   `json:"name"`
		RouterID            string   `json:"router_id"`
		NextHop             string   `json:"next_hop"`
		Cidr                string   `json:"cidr"`
		Tags                []string `json:"tags"`
		ProjectID           string   `json:"project_id,omitempty"`
		Status              string   `json:"status"`
		CreatedAt           string   `json:"created_at"`
		UpdatedAt           string   `json:"updated_at"`
		AccountID           string   `json:"account_id"`
		NetopsStaticRouteID string   `json:"netops_static_route_id"`
		SubnetID            string   `json:"subnet_id"`
	}

	StaticRouteCreateRequest struct {
		RouterID  string   `json:"router_id"`
		NextHop   string   `json:"next_hop"`
		Cidr      string   `json:"cidr"`
		Name      string   `json:"name,omitempty"`
		ProjectID string   `json:"project_id,omitempty"`
		Tags      []string `json:"tags,omitempty"`
	}

	StaticRouteUpdateRequest struct {
		Name *string   `json:"name,omitempty"`
		Tags *[]string `json:"tags,omitempty"`
	}

	StaticRoutesFilters struct {
		Name string
	}

	StaticRoutesQueryParams struct {
		Filters StaticRoutesFilters
	}
)

func (q *StaticRoutesQueryParams) queryParamsRaw() string {
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

func (client *ServiceClient) StaticRoutes(ctx context.Context, q *StaticRoutesQueryParams) (*[]StaticRoute, *ResponseResult, error) {
	queryParams := ""
	if qRaw := q.queryParamsRaw(); q != nil && qRaw != "" {
		queryParams = "?" + qRaw
	}

	u := fmt.Sprintf("%s/static_routes%s", client.APIUrl, queryParams)

	responseResult, err := client.DoRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var response []StaticRoute
	err = responseResult.ExtractResult(&response)
	if err != nil {
		return nil, responseResult, err
	}

	return &response, responseResult, nil
}

func (client *ServiceClient) StaticRoute(ctx context.Context, subnetID string) (*StaticRoute, *ResponseResult, error) {
	u := fmt.Sprintf("%s/static_routes/%s", client.APIUrl, subnetID)

	responseResult, err := client.DoRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var result *StaticRoute
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) StaticRouteCreate(ctx context.Context, req *StaticRouteCreateRequest) (*StaticRoute, *ResponseResult, error) {
	u := fmt.Sprintf("%s/static_routes", client.APIUrl)

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

	var result *StaticRoute
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) StaticRouteUpdate(ctx context.Context, subnetID string, req *StaticRouteUpdateRequest) (*StaticRoute, *ResponseResult, error) {
	u := fmt.Sprintf("%s/static_routes/%s", client.APIUrl, subnetID)

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

	var result *StaticRoute
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) StaticRouteDelete(ctx context.Context, subnetID string) (*ResponseResult, error) {
	u := fmt.Sprintf("%s/static_routes/%s", client.APIUrl, subnetID)

	responseResult, err := client.DoRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}

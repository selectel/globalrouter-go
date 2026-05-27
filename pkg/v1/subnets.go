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
	Subnet struct {
		ID               string   `json:"id"`
		Name             string   `json:"name"`
		NetworkID        string   `json:"network_id,omitempty"`
		Gateway          string   `json:"gateway"`
		Cidr             string   `json:"cidr"`
		ServiceAddresses []string `json:"service_addresses"`
		OsSubnetID       string   `json:"os_subnet_id,omitempty"`
		Tags             []string `json:"tags"`
		ProjectID        string   `json:"project_id,omitempty"`
		Status           string   `json:"status"`
		CreatedAt        string   `json:"created_at"`
		UpdatedAt        string   `json:"updated_at"`
		AccountID        string   `json:"account_id"`
		NetopsSubnetID   string   `json:"netops_subnet_id"`
		SvSubnetID       string   `json:"sv_subnet_id,omitempty"`
	}

	VPCSubnetCreateRequest struct {
		NetworkID        string   `json:"network_id"`
		Gateway          string   `json:"gateway,omitempty"`
		Cidr             string   `json:"cidr"`
		OsSubnetID       string   `json:"os_subnet_id"`
		Name             string   `json:"name,omitempty"`
		ProjectID        string   `json:"project_id,omitempty"`
		ServiceAddresses []string `json:"service_addresses,omitempty"`
		Tags             []string `json:"tags,omitempty"`
	}

	DedicatedSubnetCreateRequest struct {
		NetworkID        string   `json:"network_id"`
		Gateway          string   `json:"gateway,omitempty"`
		Cidr             string   `json:"cidr"`
		Name             string   `json:"name,omitempty"`
		ServiceAddresses []string `json:"service_addresses,omitempty"`
		Tags             []string `json:"tags,omitempty"`
	}

	SubnetUpdateRequest struct {
		Name *string   `json:"name,omitempty"`
		Tags *[]string `json:"tags,omitempty"`
	}

	SubnetsFilters struct {
		Name string
	}

	SubnetsQueryParams struct {
		Filters SubnetsFilters
	}
)

func (q *SubnetsQueryParams) queryParamsRaw() string {
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

func (client *ServiceClient) Subnets(ctx context.Context, q *SubnetsQueryParams) (*[]Subnet, *ResponseResult, error) {
	queryParams := ""
	if qRaw := q.queryParamsRaw(); q != nil && qRaw != "" {
		queryParams = "?" + qRaw
	}

	u := fmt.Sprintf("%s/subnets%s", client.APIUrl, queryParams)

	responseResult, err := client.DoRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var response []Subnet
	err = responseResult.ExtractResult(&response)
	if err != nil {
		return nil, responseResult, err
	}

	return &response, responseResult, nil
}

func (client *ServiceClient) Subnet(ctx context.Context, subnetID string) (*Subnet, *ResponseResult, error) {
	u := fmt.Sprintf("%s/subnets/%s", client.APIUrl, subnetID)

	responseResult, err := client.DoRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var result *Subnet
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) VPCSubnetCreate(ctx context.Context, req *VPCSubnetCreateRequest) (*Subnet, *ResponseResult, error) {
	u := fmt.Sprintf("%s/subnets", client.APIUrl)

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

	var result *Subnet
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) DedicatedSubnetCreate(ctx context.Context, req *DedicatedSubnetCreateRequest) (*Subnet, *ResponseResult, error) {
	u := fmt.Sprintf("%s/subnets", client.APIUrl)

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

	var result *Subnet
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) SubnetUpdate(ctx context.Context, subnetID string, req *SubnetUpdateRequest) (*Subnet, *ResponseResult, error) {
	u := fmt.Sprintf("%s/subnets/%s", client.APIUrl, subnetID)

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

	var result *Subnet
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) SubnetDisconnect(ctx context.Context, subnetID string) (*ResponseResult, error) {
	// Always disconnect subnet
	u := fmt.Sprintf("%s/subnets/%s?disconnect=true", client.APIUrl, subnetID)

	responseResult, err := client.DoRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}

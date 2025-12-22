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
	Network struct {
		ID             string   `json:"id"`
		Name           string   `json:"name"`
		RouterID       string   `json:"router_id"`
		ZoneID         string   `json:"zone_id"`
		Status         string   `json:"status"`
		CreatedAt      string   `json:"created_at"`
		UpdatedAt      string   `json:"updated_at"`
		AccountID      string   `json:"account_id"`
		ProjectID      string   `json:"project_id,omitempty"`
		Tags           []string `json:"tags"`
		OsNetworkID    string   `json:"os_network_id,omitempty"`
		Vlan           int      `json:"vlan,omitempty"`
		InnerVlan      int      `json:"inner_vlan,omitempty"`
		NetopsVlanUUID string   `json:"netops_vlan_uuid"`
		SvNetworkID    string   `json:"sv_network_id,omitempty"`
		VdcName        string   `json:"vdc_name,omitempty"`
	}

	VPCNetworkCreateRequest struct {
		RouterID    string   `json:"router_id"`
		ZoneID      string   `json:"zone_id"`
		ProjectID   string   `json:"project_id"`
		OsNetworkID string   `json:"os_network_id"`
		Name        string   `json:"name,omitempty"`
		Tags        []string `json:"tags,omitempty"`
	}

	DedicatedNetworkCreateRequest struct {
		RouterID  string   `json:"router_id"`
		ZoneID    string   `json:"zone_id"`
		Vlan      int      `json:"vlan"`
		InnerVlan int      `json:"inner_vlan,omitempty"`
		Name      string   `json:"name,omitempty"`
		Tags      []string `json:"tags,omitempty"`
	}

	NetworkUpdateRequest struct {
		Name *string   `json:"name,omitempty"`
		Tags *[]string `json:"tags,omitempty"`
	}

	NetworksFilters struct {
		Name string
	}

	NetworksQueryParams struct {
		Filters NetworksFilters
	}
)

func (q *NetworksQueryParams) queryParamsRaw() string {
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

func (client *ServiceClient) Networks(ctx context.Context, q *NetworksQueryParams) (*[]Network, *ResponseResult, error) {
	queryParams := ""
	if qRaw := q.queryParamsRaw(); q != nil && qRaw != "" {
		queryParams = "?" + qRaw
	}

	u := fmt.Sprintf("%s/networks%s", client.APIUrl, queryParams)

	responseResult, err := client.DoRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var response []Network
	err = responseResult.ExtractResult(&response)
	if err != nil {
		return nil, responseResult, err
	}

	return &response, responseResult, nil
}

func (client *ServiceClient) Network(ctx context.Context, networkID string) (*Network, *ResponseResult, error) {
	u := fmt.Sprintf("%s/networks/%s", client.APIUrl, networkID)

	responseResult, err := client.DoRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var result *Network
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) VPCNetworkCreate(ctx context.Context, req *VPCNetworkCreateRequest) (*Network, *ResponseResult, error) {
	u := fmt.Sprintf("%s/networks", client.APIUrl)

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

	var result *Network
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) DedicatedNetworkCreate(ctx context.Context, req *DedicatedNetworkCreateRequest) (*Network, *ResponseResult, error) {
	u := fmt.Sprintf("%s/networks", client.APIUrl)

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

	var result *Network
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) NetworkUpdate(ctx context.Context, networkID string, req *NetworkUpdateRequest) (*Network, *ResponseResult, error) {
	u := fmt.Sprintf("%s/networks/%s", client.APIUrl, networkID)

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

	var result *Network
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result, responseResult, nil
}

func (client *ServiceClient) NetworkDisconnect(ctx context.Context, networkID string) (*ResponseResult, error) {
	// Always disconnect network
	u := fmt.Sprintf("%s/networks/%s?disconnect=true", client.APIUrl, networkID)

	responseResult, err := client.DoRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}

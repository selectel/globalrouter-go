package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type (
	ZoneGroup struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at,omitempty"`
	}

	ZoneGroupsFilters struct {
		Name string
	}

	ZoneGroupsQueryParams struct {
		Filters ZoneGroupsFilters
	}
)

func (q *ZoneGroupsQueryParams) queryParamsRaw() string {
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

func (client *ServiceClient) ListZoneGroups(ctx context.Context, options *ZoneGroupsQueryParams) (*[]ZoneGroup, *ResponseResult, error) {
	queryParams := ""
	if qRaw := options.queryParamsRaw(); options != nil && qRaw != "" {
		queryParams = "?" + qRaw
	}

	path := fmt.Sprintf("%s/zone_groups%s", client.APIUrl, queryParams)
	responseResult, err := client.DoRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}
	var response []ZoneGroup
	err = responseResult.ExtractResult(&response)
	if err != nil {
		return nil, responseResult, err
	}

	return &response, responseResult, nil
}

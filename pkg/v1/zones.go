package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type (
	Zone struct {
		ID          string      `json:"id"`
		Name        string      `json:"name"`
		VisibleName string      `json:"visible_name"`
		Service     string      `json:"service"`
		Enable      bool        `json:"enable"`
		AllowCreate bool        `json:"allow_create"`
		AllowUpdate bool        `json:"allow_update"`
		AllowDelete bool        `json:"allow_delete"`
		CreatedAt   string      `json:"created_at"`
		UpdatedAt   string      `json:"updated_at,omitempty"`
		Options     string      `json:"options,omitempty"`
		ZoneGroups  []ZoneGroup `json:"groups"`
	}

	ZonesFilters struct {
		Name      string
		ServiceID string
		Enable    bool
	}

	ZonesQueryParams struct {
		Filters ZonesFilters
	}
)

func (q *ZonesQueryParams) queryParamsRaw() string {
	vals := url.Values{}
	if q == nil {
		return vals.Encode()
	}

	var filters []string

	if q.Filters.Name != "" {
		filters = append(filters, "name="+q.Filters.Name)
	}

	if q.Filters.ServiceID != "" {
		filters = append(filters, "service_id="+q.Filters.ServiceID)
	}

	if q.Filters.Enable {
		filters = append(filters, "enable=true")
	}
	if len(filters) != 0 {
		vals.Set("filters", strings.Join(filters, ","))
	}

	return vals.Encode()
}

func (client *ServiceClient) ListZones(ctx context.Context, options *ZonesQueryParams) (*[]Zone, *ResponseResult, error) {
	queryParams := ""
	if qRaw := options.queryParamsRaw(); options != nil && qRaw != "" {
		queryParams = "?" + qRaw
	}

	path := fmt.Sprintf("%s/zones%s", client.APIUrl, queryParams)
	responseResult, err := client.DoRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}
	var response []Zone
	err = responseResult.ExtractResult(&response)
	if err != nil {
		return nil, responseResult, err
	}

	return &response, responseResult, nil
}

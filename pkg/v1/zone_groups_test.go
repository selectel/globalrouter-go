package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceClient_ZoneGroups(t *testing.T) {
	t.Run("SuccessWithQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "123e4567-e89b-12d3-a456-111111111111",
			"name": "zone_group_group1",
			"description": "some data",
			"created_at": "2160-01-01 00:00:00",
			"updated_at": "2160-01-01 00:00:00"
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListZoneGroups(context.Background(), &ZoneGroupsQueryParams{Filters: ZoneGroupsFilters{Name: "zone_group_group1"}})

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := []ZoneGroup{
			{
				ID:          "123e4567-e89b-12d3-a456-111111111111",
				Name:        "zone_group_group1",
				Description: "some data",
				CreatedAt:   "2160-01-01 00:00:00",
				UpdatedAt:   "2160-01-01 00:00:00",
			},
		}
		require.Equal(t, &want, res)
	})

	t.Run("SuccessWithoutQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "123e4567-e89b-12d3-a456-111111111111",
			"name": "zone_group_group1",
			"description": "some data",
			"created_at": "2160-01-01 00:00:00",
			"updated_at": "2160-01-01 00:00:00"
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListZoneGroups(context.Background(), nil)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := []ZoneGroup{
			{
				ID:          "123e4567-e89b-12d3-a456-111111111111",
				Name:        "zone_group_group1",
				Description: "some data",
				CreatedAt:   "2160-01-01 00:00:00",
				UpdatedAt:   "2160-01-01 00:00:00",
			},
		}
		require.Equal(t, &want, res)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListZoneGroups(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", NewFakeTransport(fakeResp, nil))

		// Execute
		res, respRes, err := client.ListZoneGroups(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, res)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := NewFakeTransport(nil, errors.New("network failure"))
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListZoneGroups(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.Nil(t, respRes)
	})
}

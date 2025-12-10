package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceClient_Zones(t *testing.T) {
	t.Run("SuccessWithQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "123e4567-e89b-12d3-a456-426655440000",
			"name": "ru-1",
			"visible_name": "ru-1",
			"service": "vpc",
			"enable": true,
			"allow_create": true,
			"allow_update": true,
			"allow_delete": true,
			"created_at": "2170-01-01 00:00:00",
			"updated_at": "2170-01-01 00:00:00",
			"options": "",
			"groups": [
				{
				"id": "123e4567-e89b-12d3-a456-111111111111",
				"name": "zone_group1",
				"description": "some data",
				"created_at": "2160-01-01 00:00:00",
				"updated_at": "2160-01-01 00:00:00"
				}
			]
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListZones(context.Background(), &ZonesQueryParams{Filters: ZonesFilters{Name: "ru-1", Enable: true}})

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := []Zone{
			{
				ID:          "123e4567-e89b-12d3-a456-426655440000",
				Name:        "ru-1",
				VisibleName: "ru-1",
				Service:     "vpc",
				Enable:      true,
				AllowCreate: true,
				AllowUpdate: true,
				AllowDelete: true,
				CreatedAt:   "2170-01-01 00:00:00",
				UpdatedAt:   "2170-01-01 00:00:00",
				Options:     "",
				ZoneGroups: []ZoneGroup{
					{
						ID:          "123e4567-e89b-12d3-a456-111111111111",
						Name:        "zone_group1",
						Description: "some data",
						CreatedAt:   "2160-01-01 00:00:00",
						UpdatedAt:   "2160-01-01 00:00:00",
					},
				},
			},
		}
		require.Equal(t, &want, res)
	})

	t.Run("SuccessWithoutQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "123e4567-e89b-12d3-a456-426655440000",
			"name": "ru-1",
			"visible_name": "ru-1",
			"service": "vpc",
			"enable": true,
			"allow_create": true,
			"allow_update": true,
			"allow_delete": true,
			"created_at": "2170-01-01 00:00:00",
			"updated_at": "2170-01-01 00:00:00",
			"options": "",
			"groups": []
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListZones(context.Background(), nil)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := []Zone{
			{
				ID:          "123e4567-e89b-12d3-a456-426655440000",
				Name:        "ru-1",
				VisibleName: "ru-1",
				Service:     "vpc",
				Enable:      true,
				AllowCreate: true,
				AllowUpdate: true,
				AllowDelete: true,
				CreatedAt:   "2170-01-01 00:00:00",
				UpdatedAt:   "2170-01-01 00:00:00",
				Options:     "",
				ZoneGroups:  []ZoneGroup{},
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
		res, respRes, err := client.ListZones(context.Background(), nil)

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
		res, respRes, err := client.ListZones(context.Background(), nil)

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
		res, respRes, err := client.ListZones(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.Nil(t, respRes)
	})
}

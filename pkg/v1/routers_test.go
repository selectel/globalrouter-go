package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceClient_Routers(t *testing.T) {
	t.Run("SuccessWithQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "11111111-1111-1111-1111-111111111111",
			"name": "test_rtr",
			"updated_at": "",
			"created_at": "2025-11-16T08:27:19.076775",
			"status": "ACTIVE",
			"enabled": true,
			"account_id": "777777",
			"project_id": null,
			"netops_router_id": "11111111-1111-1111-1111-111111111111",
			"tags": ["blue", "red"],
			"leak_uuid": null,
			"prefix_pool_id": "88a5f82b-3151-4d84-b111-2e06ee44d899",
			"vpn_id": 1,
			"networks": [],
			"prefix_pool": {
				"id": null,
				"prefixes": null,
				"default": null
			}
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Routers(context.Background(), &RoutersQueryParams{Filters: RoutersFilters{Name: "test_rtr"}})

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &[]Router{
			{
				ID:             "11111111-1111-1111-1111-111111111111",
				Name:           "test_rtr",
				Status:         "ACTIVE",
				CreatedAt:      "2025-11-16T08:27:19.076775",
				UpdatedAt:      "",
				Enabled:        true,
				Tags:           []string{"blue", "red"},
				AccountID:      "777777",
				ProjectID:      "",
				LeakUUID:       "",
				NetopsRouterID: "11111111-1111-1111-1111-111111111111",
				PrefixPoolID:   "88a5f82b-3151-4d84-b111-2e06ee44d899",
				VpnID:          1,
			},
		}
		require.Equal(t, want, res)
	})

	t.Run("SuccessWithoutQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "11111111-1111-1111-1111-111111111111",
			"name": "test_rtr",
			"updated_at": "",
			"created_at": "2025-11-16T08:27:19.076775",
			"status": "ACTIVE",
			"enabled": true,
			"account_id": "777777",
			"project_id": null,
			"netops_router_id": "11111111-1111-1111-1111-111111111111",
			"tags": ["tag1"],
			"leak_uuid": null,
			"prefix_pool_id": "88a5f82b-3151-4d84-b111-2e06ee44d899",
			"vpn_id": 1,
			"networks": [],
			"prefix_pool": {
				"id": null,
				"prefixes": null,
				"default": null
			}
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Routers(context.Background(), nil)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &[]Router{
			{
				ID:             "11111111-1111-1111-1111-111111111111",
				Name:           "test_rtr",
				Status:         "ACTIVE",
				CreatedAt:      "2025-11-16T08:27:19.076775",
				UpdatedAt:      "",
				Enabled:        true,
				Tags:           []string{"tag1"},
				AccountID:      "777777",
				ProjectID:      "",
				LeakUUID:       "",
				NetopsRouterID: "11111111-1111-1111-1111-111111111111",
				PrefixPoolID:   "88a5f82b-3151-4d84-b111-2e06ee44d899",
				VpnID:          1,
			},
		}
		require.Equal(t, want, res)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Routers(context.Background(), nil)

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
		client := newFakeClient("http://fake", &NewFakeTransport{resp: fakeResp, err: nil})

		// Execute
		res, respRes, err := client.Routers(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, res)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := &NewFakeTransport{resp: nil, err: errors.New("network failure")}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Routers(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_Router(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "11111111-1111-1111-1111-111111111111",
			"name": "test_rtr",
			"updated_at": "2025-11-16T08:27:19.782861",
			"created_at": "2025-11-16T08:27:19.076775",
			"status": "ACTIVE",
			"enabled": true,
			"account_id": "777777",
			"project_id": null,
			"netops_router_id": "11111111-1111-1111-1111-111111111111",
			"tags": ["blue", "red"],
			"leak_uuid": null,
			"prefix_pool_id": "88a5f82b-3151-4d84-b111-2e06ee44d899",
			"vpn_id": 1,
			"networks": [],
			"prefix_pool": {
				"id": null,
				"prefixes": null,
				"default": null
			}
		}`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.Router(context.Background(), "plan-id-1")

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		wantRouter := &Router{
			ID:             "11111111-1111-1111-1111-111111111111",
			Name:           "test_rtr",
			Status:         "ACTIVE",
			CreatedAt:      "2025-11-16T08:27:19.076775",
			UpdatedAt:      "2025-11-16T08:27:19.782861",
			Enabled:        true,
			Tags:           []string{"blue", "red"},
			AccountID:      "777777",
			ProjectID:      "",
			LeakUUID:       "",
			NetopsRouterID: "11111111-1111-1111-1111-111111111111",
			PrefixPoolID:   "88a5f82b-3151-4d84-b111-2e06ee44d899",
			VpnID:          1,
		}
		require.Equal(t, wantRouter, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.Router(context.Background(), "plan-id-1")

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", &NewFakeTransport{resp: fakeResp, err: nil})

		// Execute
		plan, respRes, err := client.Router(context.Background(), "plan-id-1")

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, plan)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := &NewFakeTransport{resp: nil, err: errors.New("network failure")}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.Router(context.Background(), "plan-id-1")

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_RouterCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "11111111-1111-1111-1111-111111111111",
			"name": "test_rtr",
			"updated_at": "2025-11-16T08:27:19.782861",
			"created_at": "2025-11-16T08:27:19.076775",
			"status": "ACTIVE",
			"enabled": true,
			"account_id": "777777",
			"project_id": null,
			"netops_router_id": "11111111-1111-1111-1111-111111111111",
			"tags": ["blue", "red"],
			"leak_uuid": null,
			"prefix_pool_id": "88a5f82b-3151-4d84-b111-2e06ee44d899",
			"vpn_id": 1,
			"networks": [],
			"prefix_pool": {
				"id": null,
				"prefixes": null,
				"default": null
			}
		}`
		fakeResp := NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		createReq := &RouterCreateRequest{
			Name: "test_rtr",
			Tags: []string{"blue", "red"},
		}

		// Execute
		plan, respRes, err := client.RouterCreate(context.Background(), createReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 201, respRes.StatusCode)
		wantRouter := &Router{
			ID:             "11111111-1111-1111-1111-111111111111",
			Name:           "test_rtr",
			Status:         "ACTIVE",
			CreatedAt:      "2025-11-16T08:27:19.076775",
			UpdatedAt:      "2025-11-16T08:27:19.782861",
			Enabled:        true,
			Tags:           []string{"blue", "red"},
			AccountID:      "777777",
			ProjectID:      "",
			LeakUUID:       "",
			NetopsRouterID: "11111111-1111-1111-1111-111111111111",
			PrefixPoolID:   "88a5f82b-3151-4d84-b111-2e06ee44d899",
			VpnID:          1,
		}
		require.Equal(t, wantRouter, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.RouterCreate(context.Background(), &RouterCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.NotNil(t, respRes)
		require.Equal(t, 201, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", &NewFakeTransport{resp: fakeResp, err: nil})

		// Execute
		plan, respRes, err := client.RouterCreate(context.Background(), &RouterCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, plan)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := &NewFakeTransport{resp: nil, err: errors.New("network failure")}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.RouterCreate(context.Background(), &RouterCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_RouterUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "11111111-1111-1111-1111-111111111111",
			"name": "updated_test_rtr",
			"updated_at": "2025-11-16T08:27:19.782861",
			"created_at": "2025-11-16T08:27:19.076775",
			"status": "ACTIVE",
			"enabled": true,
			"account_id": "777777",
			"project_id": null,
			"netops_router_id": "11111111-1111-1111-1111-111111111111",
			"tags": ["tag1"],
			"leak_uuid": null,
			"prefix_pool_id": "88a5f82b-3151-4d84-b111-2e06ee44d899",
			"vpn_id": 1,
			"networks": [],
			"prefix_pool": {
				"id": null,
				"prefixes": null,
				"default": null
			}
		}`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		newName := "updated_test_rtr"
		tags := []string{"tag1"}
		updateReq := &RouterUpdateRequest{
			Name: &newName,
			Tags: &tags,
		}

		// Execute
		plan, respRes, err := client.RouterUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", updateReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		wantRouter := &Router{
			ID:             "11111111-1111-1111-1111-111111111111",
			Name:           "updated_test_rtr",
			Status:         "ACTIVE",
			CreatedAt:      "2025-11-16T08:27:19.076775",
			UpdatedAt:      "2025-11-16T08:27:19.782861",
			Enabled:        true,
			Tags:           []string{"tag1"},
			AccountID:      "777777",
			ProjectID:      "",
			LeakUUID:       "",
			NetopsRouterID: "11111111-1111-1111-1111-111111111111",
			PrefixPoolID:   "88a5f82b-3151-4d84-b111-2e06ee44d899",
			VpnID:          1,
		}
		require.Equal(t, wantRouter, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.RouterUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", &RouterUpdateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", &NewFakeTransport{resp: fakeResp, err: nil})

		// Execute
		plan, respRes, err := client.RouterUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", &RouterUpdateRequest{})

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, plan)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := &NewFakeTransport{resp: nil, err: errors.New("network failure")}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.RouterUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", &RouterUpdateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_RouterDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		fakeResp := NewFakeResponse(204, "") //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		respRes, err := client.RouterDelete(context.Background(), "11111111-1111-1111-1111-111111111111")

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 204, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", &NewFakeTransport{resp: fakeResp, err: nil})

		// Execute
		respRes, err := client.RouterDelete(context.Background(), "11111111-1111-1111-1111-111111111111")

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := &NewFakeTransport{resp: nil, err: errors.New("network failure")}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		respRes, err := client.RouterDelete(context.Background(), "11111111-1111-1111-1111-111111111111")

		// Analyse
		require.Error(t, err)
		require.Nil(t, respRes)
	})
}

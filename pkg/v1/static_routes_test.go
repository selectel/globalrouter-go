package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceClient_StaticRoutes(t *testing.T) {
	t.Run("SuccessWithQuery", func(t *testing.T) {
		// Prepare
		body := `[{
  			"id": "33333333-3333-3333-3333-333333333333",
  			"name": "stat_route",
  			"cidr": "0.0.0.0/0",
  			"next_hop": "10.20.0.252",
  			"router_id": "11111111-1111-1111-1111-111111111111",
  			"created_at": "2025-10-22T09:22:37.294081",
  			"updated_at": "2025-10-22T09:22:48.728094",
  			"status": "ACTIVE",
  			"account_id": "777777",
  			"project_id": null,
  			"subnet_id": "22222222-2222-2222-2222-222222222222",
  			"netops_static_route_id": "33333333-3333-3333-3333-333333333333",
  			"tags": ["net"]
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.StaticRoutes(context.Background(), &StaticRoutesQueryParams{Filters: StaticRoutesFilters{Name: "stat_route"}})

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &[]StaticRoute{
			{
				ID:                  "33333333-3333-3333-3333-333333333333",
				Name:                "stat_route",
				RouterID:            "11111111-1111-1111-1111-111111111111",
				NextHop:             "10.20.0.252",
				Cidr:                "0.0.0.0/0",
				Status:              "ACTIVE",
				CreatedAt:           "2025-10-22T09:22:37.294081",
				UpdatedAt:           "2025-10-22T09:22:48.728094",
				Tags:                []string{"net"},
				AccountID:           "777777",
				ProjectID:           "",
				NetopsStaticRouteID: "33333333-3333-3333-3333-333333333333",
				SubnetID:            "22222222-2222-2222-2222-222222222222",
			},
		}
		require.Equal(t, want, res)
	})

	t.Run("SuccessWithoutQuery", func(t *testing.T) {
		// Prepare
		body := `[{
  			"id": "33333333-3333-3333-3333-333333333333",
  			"name": "stat_route",
  			"cidr": "0.0.0.0/0",
  			"next_hop": "10.20.0.252",
  			"router_id": "11111111-1111-1111-1111-111111111111",
  			"created_at": "2025-10-22T09:22:37.294081",
  			"updated_at": "2025-10-22T09:22:48.728094",
  			"status": "ACTIVE",
  			"account_id": "777777",
  			"project_id": null,
  			"subnet_id": "22222222-2222-2222-2222-222222222222",
  			"netops_static_route_id": "33333333-3333-3333-3333-333333333333",
  			"tags": ["net"]
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.StaticRoutes(context.Background(), nil)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &[]StaticRoute{
			{
				ID:                  "33333333-3333-3333-3333-333333333333",
				Name:                "stat_route",
				RouterID:            "11111111-1111-1111-1111-111111111111",
				NextHop:             "10.20.0.252",
				Cidr:                "0.0.0.0/0",
				Status:              "ACTIVE",
				CreatedAt:           "2025-10-22T09:22:37.294081",
				UpdatedAt:           "2025-10-22T09:22:48.728094",
				Tags:                []string{"net"},
				AccountID:           "777777",
				ProjectID:           "",
				NetopsStaticRouteID: "33333333-3333-3333-3333-333333333333",
				SubnetID:            "22222222-2222-2222-2222-222222222222",
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
		res, respRes, err := client.StaticRoutes(context.Background(), nil)

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
		res, respRes, err := client.StaticRoutes(context.Background(), nil)

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
		res, respRes, err := client.StaticRoutes(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_StaticRoute(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
  			"id": "33333333-3333-3333-3333-333333333333",
  			"name": "stat_route",
  			"cidr": "0.0.0.0/0",
  			"next_hop": "10.20.0.252",
  			"router_id": "11111111-1111-1111-1111-111111111111",
  			"created_at": "2025-10-22T09:22:37.294081",
  			"updated_at": "2025-10-22T09:22:48.728094",
  			"status": "ACTIVE",
  			"account_id": "777777",
  			"project_id": null,
  			"subnet_id": "22222222-2222-2222-2222-222222222222",
  			"netops_static_route_id": "33333333-3333-3333-3333-333333333333",
  			"tags": ["net"]
		}`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.StaticRoute(context.Background(), "plan-id-1")

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		wantStaticRoute := &StaticRoute{
			ID:                  "33333333-3333-3333-3333-333333333333",
			Name:                "stat_route",
			RouterID:            "11111111-1111-1111-1111-111111111111",
			NextHop:             "10.20.0.252",
			Cidr:                "0.0.0.0/0",
			Status:              "ACTIVE",
			CreatedAt:           "2025-10-22T09:22:37.294081",
			UpdatedAt:           "2025-10-22T09:22:48.728094",
			Tags:                []string{"net"},
			AccountID:           "777777",
			ProjectID:           "",
			NetopsStaticRouteID: "33333333-3333-3333-3333-333333333333",
			SubnetID:            "22222222-2222-2222-2222-222222222222",
		}
		require.Equal(t, wantStaticRoute, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.StaticRoute(context.Background(), "plan-id-1")

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
		plan, respRes, err := client.StaticRoute(context.Background(), "plan-id-1")

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
		plan, respRes, err := client.StaticRoute(context.Background(), "plan-id-1")

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_StaticRouteCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
  			"id": "33333333-3333-3333-3333-333333333333",
  			"name": "stat_route",
  			"cidr": "0.0.0.0/0",
  			"next_hop": "10.20.0.252",
  			"router_id": "11111111-1111-1111-1111-111111111111",
  			"created_at": "2025-10-22T09:22:37.294081",
  			"updated_at": "2025-10-22T09:22:48.728094",
  			"status": "ACTIVE",
  			"account_id": "777777",
  			"project_id": null,
  			"subnet_id": "22222222-2222-2222-2222-222222222222",
  			"netops_static_route_id": "33333333-3333-3333-3333-333333333333",
  			"tags": ["net"]
		}`
		fakeResp := NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		createReq := &StaticRouteCreateRequest{
			Name: "stat_route",
			Tags: []string{"blue", "red"},
		}

		// Execute
		plan, respRes, err := client.StaticRouteCreate(context.Background(), createReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 201, respRes.StatusCode)
		wantStaticRoute := &StaticRoute{
			ID:                  "33333333-3333-3333-3333-333333333333",
			Name:                "stat_route",
			RouterID:            "11111111-1111-1111-1111-111111111111",
			NextHop:             "10.20.0.252",
			Cidr:                "0.0.0.0/0",
			Status:              "ACTIVE",
			CreatedAt:           "2025-10-22T09:22:37.294081",
			UpdatedAt:           "2025-10-22T09:22:48.728094",
			Tags:                []string{"net"},
			AccountID:           "777777",
			ProjectID:           "",
			NetopsStaticRouteID: "33333333-3333-3333-3333-333333333333",
			SubnetID:            "22222222-2222-2222-2222-222222222222",
		}
		require.Equal(t, wantStaticRoute, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.StaticRouteCreate(context.Background(), &StaticRouteCreateRequest{})

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
		plan, respRes, err := client.StaticRouteCreate(context.Background(), &StaticRouteCreateRequest{})

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
		plan, respRes, err := client.StaticRouteCreate(context.Background(), &StaticRouteCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_StaticRouteUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
  			"id": "33333333-3333-3333-3333-333333333333",
  			"name": "updated_stat_route",
  			"cidr": "0.0.0.0/0",
  			"next_hop": "10.20.0.252",
  			"router_id": "11111111-1111-1111-1111-111111111111",
  			"created_at": "2025-10-22T09:22:37.294081",
  			"updated_at": "2025-10-22T09:22:48.728094",
  			"status": "ACTIVE",
  			"account_id": "777777",
  			"project_id": null,
  			"subnet_id": "22222222-2222-2222-2222-222222222222",
  			"netops_static_route_id": "33333333-3333-3333-3333-333333333333",
  			"tags": ["tag1"]
		}`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		newName := "updated_stat_route"
		tags := []string{"tag1"}
		updateReq := &StaticRouteUpdateRequest{
			Name: &newName,
			Tags: &tags,
		}

		// Execute
		plan, respRes, err := client.StaticRouteUpdate(context.Background(), "33333333-3333-3333-3333-333333333333", updateReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		wantStaticRoute := &StaticRoute{
			ID:                  "33333333-3333-3333-3333-333333333333",
			Name:                "updated_stat_route",
			RouterID:            "11111111-1111-1111-1111-111111111111",
			NextHop:             "10.20.0.252",
			Cidr:                "0.0.0.0/0",
			Status:              "ACTIVE",
			CreatedAt:           "2025-10-22T09:22:37.294081",
			UpdatedAt:           "2025-10-22T09:22:48.728094",
			Tags:                []string{"tag1"},
			AccountID:           "777777",
			ProjectID:           "",
			NetopsStaticRouteID: "33333333-3333-3333-3333-333333333333",
			SubnetID:            "22222222-2222-2222-2222-222222222222",
		}
		require.Equal(t, wantStaticRoute, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.StaticRouteUpdate(context.Background(), "33333333-3333-3333-3333-333333333333", &StaticRouteUpdateRequest{})

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
		plan, respRes, err := client.StaticRouteUpdate(context.Background(), "33333333-3333-3333-3333-333333333333", &StaticRouteUpdateRequest{})

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
		plan, respRes, err := client.StaticRouteUpdate(context.Background(), "33333333-3333-3333-3333-333333333333", &StaticRouteUpdateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_StaticRouteDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		fakeResp := NewFakeResponse(204, "") //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		respRes, err := client.StaticRouteDelete(context.Background(), "33333333-3333-3333-3333-333333333333")

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
		respRes, err := client.StaticRouteDelete(context.Background(), "33333333-3333-3333-3333-333333333333")

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
		respRes, err := client.StaticRouteDelete(context.Background(), "33333333-3333-3333-3333-333333333333")

		// Analyse
		require.Error(t, err)
		require.Nil(t, respRes)
	})
}

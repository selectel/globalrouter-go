package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceClient_Networks(t *testing.T) {
	t.Run("SuccessWithQuery", func(t *testing.T) {
		// Prepare
		body := `[
			{
				"id": "11111111-1111-1111-1111-111111111111",
				"name": "cloud_net",
				"vlan": 21412,
				"router_id": "22222222-2222-2222-2222-222222222222",
				"zone_id": "33333333-3333-3333-3333-333333333333",
				"project_id": "77777777777777777777777777777777",
				"vdc_name": null,
				"updated_at": null,
				"created_at": "2025-12-08T18:57:57.412757",
				"status": "ACTIVE",
				"account_id": "777777",
				"sv_network_id": "44444444-4444-4444-4444-444444444444",
				"netops_vlan_uuid": "11111111-1111-1111-1111-111111111111",
				"tags": ["tag11"],
				"inner_vlan": null,
				"os_network_id": "55555555-5555-5555-5555-555555555555"
			}
		]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Networks(context.Background(), &NetworksQueryParams{Filters: NetworksFilters{Name: "cloud_net"}})

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &[]Network{
			{
				ID:             "11111111-1111-1111-1111-111111111111",
				Name:           "cloud_net",
				RouterID:       "22222222-2222-2222-2222-222222222222",
				ZoneID:         "33333333-3333-3333-3333-333333333333",
				Status:         "ACTIVE",
				CreatedAt:      "2025-12-08T18:57:57.412757",
				UpdatedAt:      "",
				AccountID:      "777777",
				ProjectID:      "77777777777777777777777777777777",
				Tags:           []string{"tag11"},
				OsNetworkID:    "55555555-5555-5555-5555-555555555555",
				Vlan:           21412,
				InnerVlan:      0,
				NetopsVlanUUID: "11111111-1111-1111-1111-111111111111",
				SvNetworkID:    "44444444-4444-4444-4444-444444444444",
				VdcName:        "",
			},
		}
		require.Equal(t, want, res)
	})

	t.Run("SuccessWithoutQuery", func(t *testing.T) {
		// Prepare
		body := `[
			{
				"id": "11111111-1111-1111-1111-111111111111",
				"name": "cloud_net",
				"vlan": 21412,
				"router_id": "22222222-2222-2222-2222-222222222222",
				"zone_id": "33333333-3333-3333-3333-333333333333",
				"project_id": "77777777777777777777777777777777",
				"vdc_name": null,
				"updated_at": null,
				"created_at": "2025-12-08T18:57:57.412757",
				"status": "ACTIVE",
				"account_id": "777777",
				"sv_network_id": "44444444-4444-4444-4444-444444444444",
				"netops_vlan_uuid": "11111111-1111-1111-1111-111111111111",
				"tags": ["tag11"],
				"inner_vlan": null,
				"os_network_id": "55555555-5555-5555-5555-555555555555"
			},
			{
				"id": "88888888-8888-8888-8888-888888888888",
				"name": "dedicated_net",
				"vlan": 16461,
				"router_id": "22222222-2222-2222-2222-222222222222",
				"zone_id": "99999999-9999-9999-9999-999999999999",
				"project_id": null,
				"vdc_name": null,
				"updated_at": "2025-12-05T17:00:25.690283",
				"created_at": "2025-12-05T17:00:13.637735",
				"status": "ACTIVE",
				"account_id": "777777",
				"sv_network_id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
				"netops_vlan_uuid": "88888888-8888-8888-8888-888888888888",
				"tags": ["tag22"],
				"inner_vlan": null,
				"os_network_id": null
			}
		]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Networks(context.Background(), nil)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &[]Network{
			{
				ID:             "11111111-1111-1111-1111-111111111111",
				Name:           "cloud_net",
				RouterID:       "22222222-2222-2222-2222-222222222222",
				ZoneID:         "33333333-3333-3333-3333-333333333333",
				Status:         "ACTIVE",
				CreatedAt:      "2025-12-08T18:57:57.412757",
				UpdatedAt:      "",
				AccountID:      "777777",
				ProjectID:      "77777777777777777777777777777777",
				Tags:           []string{"tag11"},
				OsNetworkID:    "55555555-5555-5555-5555-555555555555",
				Vlan:           21412,
				InnerVlan:      0,
				NetopsVlanUUID: "11111111-1111-1111-1111-111111111111",
				SvNetworkID:    "44444444-4444-4444-4444-444444444444",
				VdcName:        "",
			},
			{
				ID:             "88888888-8888-8888-8888-888888888888",
				Name:           "dedicated_net",
				RouterID:       "22222222-2222-2222-2222-222222222222",
				ZoneID:         "99999999-9999-9999-9999-999999999999",
				Status:         "ACTIVE",
				CreatedAt:      "2025-12-05T17:00:13.637735",
				UpdatedAt:      "2025-12-05T17:00:25.690283",
				AccountID:      "777777",
				ProjectID:      "",
				Tags:           []string{"tag22"},
				OsNetworkID:    "",
				Vlan:           16461,
				InnerVlan:      0,
				NetopsVlanUUID: "88888888-8888-8888-8888-888888888888",
				SvNetworkID:    "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
				VdcName:        "",
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
		res, respRes, err := client.Networks(context.Background(), nil)

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
		res, respRes, err := client.Networks(context.Background(), nil)

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
		res, respRes, err := client.Networks(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_Network(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "88888888-8888-8888-8888-888888888888",
			"name": "dedicated_net",
			"vlan": 16461,
			"router_id": "22222222-2222-2222-2222-222222222222",
			"zone_id": "99999999-9999-9999-9999-999999999999",
			"project_id": null,
			"vdc_name": null,
			"updated_at": "2025-12-05T17:00:25.690283",
			"created_at": "2025-12-05T17:00:13.637735",
			"status": "ACTIVE",
			"account_id": "777777",
			"sv_network_id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			"netops_vlan_uuid": "88888888-8888-8888-8888-888888888888",
			"tags": ["tag22"],
			"inner_vlan": null,
			"os_network_id": null
		}`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.Network(context.Background(), "plan-id-1")

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		wantNetwork := &Network{
			ID:             "88888888-8888-8888-8888-888888888888",
			Name:           "dedicated_net",
			RouterID:       "22222222-2222-2222-2222-222222222222",
			ZoneID:         "99999999-9999-9999-9999-999999999999",
			Status:         "ACTIVE",
			CreatedAt:      "2025-12-05T17:00:13.637735",
			UpdatedAt:      "2025-12-05T17:00:25.690283",
			AccountID:      "777777",
			ProjectID:      "",
			Tags:           []string{"tag22"},
			OsNetworkID:    "",
			Vlan:           16461,
			InnerVlan:      0,
			NetopsVlanUUID: "88888888-8888-8888-8888-888888888888",
			SvNetworkID:    "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			VdcName:        "",
		}
		require.Equal(t, wantNetwork, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.Network(context.Background(), "plan-id-1")

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
		plan, respRes, err := client.Network(context.Background(), "plan-id-1")

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
		plan, respRes, err := client.Network(context.Background(), "plan-id-1")

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_VPCNetworkCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "11111111-1111-1111-1111-111111111111",
			"name": "cloud_net",
			"vlan": 21412,
			"router_id": "22222222-2222-2222-2222-222222222222",
			"zone_id": "33333333-3333-3333-3333-333333333333",
			"project_id": "77777777777777777777777777777777",
			"vdc_name": null,
			"updated_at": null,
			"created_at": "2025-12-08T18:57:57.412757",
			"status": "ACTIVE",
			"account_id": "777777",
			"sv_network_id": "44444444-4444-4444-4444-444444444444",
			"netops_vlan_uuid": "11111111-1111-1111-1111-111111111111",
			"tags": ["tag11"],
			"inner_vlan": null,
			"os_network_id": "55555555-5555-5555-5555-555555555555"
		}`
		fakeResp := NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		createReq := &VPCNetworkCreateRequest{
			Name:        "cloud_net",
			RouterID:    "22222222-2222-2222-2222-222222222222",
			ZoneID:      "33333333-3333-3333-3333-333333333333",
			ProjectID:   "77777777777777777777777777777777",
			OsNetworkID: "55555555-5555-5555-5555-555555555555",
			Tags:        []string{"tag11"},
		}

		// Execute
		plan, respRes, err := client.VPCNetworkCreate(context.Background(), createReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 201, respRes.StatusCode)
		wantNetwork := &Network{
			ID:             "11111111-1111-1111-1111-111111111111",
			Name:           "cloud_net",
			RouterID:       "22222222-2222-2222-2222-222222222222",
			ZoneID:         "33333333-3333-3333-3333-333333333333",
			Status:         "ACTIVE",
			CreatedAt:      "2025-12-08T18:57:57.412757",
			UpdatedAt:      "",
			AccountID:      "777777",
			ProjectID:      "77777777777777777777777777777777",
			Tags:           []string{"tag11"},
			OsNetworkID:    "55555555-5555-5555-5555-555555555555",
			Vlan:           21412,
			InnerVlan:      0,
			NetopsVlanUUID: "11111111-1111-1111-1111-111111111111",
			SvNetworkID:    "44444444-4444-4444-4444-444444444444",
			VdcName:        "",
		}
		require.Equal(t, wantNetwork, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.VPCNetworkCreate(context.Background(), &VPCNetworkCreateRequest{})

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
		plan, respRes, err := client.VPCNetworkCreate(context.Background(), &VPCNetworkCreateRequest{})

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
		plan, respRes, err := client.VPCNetworkCreate(context.Background(), &VPCNetworkCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_DedicatedNetworkCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "88888888-8888-8888-8888-888888888888",
			"name": "dedicated_net",
			"vlan": 16461,
			"router_id": "22222222-2222-2222-2222-222222222222",
			"zone_id": "99999999-9999-9999-9999-999999999999",
			"project_id": null,
			"vdc_name": null,
			"updated_at": "2025-12-05T17:00:25.690283",
			"created_at": "2025-12-05T17:00:13.637735",
			"status": "ACTIVE",
			"account_id": "777777",
			"sv_network_id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			"netops_vlan_uuid": "88888888-8888-8888-8888-888888888888",
			"tags": ["tag22"],
			"inner_vlan": 321,
			"os_network_id": null
		}`
		fakeResp := NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		createReq := &DedicatedNetworkCreateRequest{
			RouterID:  "22222222-2222-2222-2222-222222222222",
			ZoneID:    "99999999-9999-9999-9999-999999999999",
			Vlan:      16461,
			InnerVlan: 321,
			Name:      "dedicated_net",
			Tags:      []string{"tag22"},
		}

		// Execute
		plan, respRes, err := client.DedicatedNetworkCreate(context.Background(), createReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 201, respRes.StatusCode)
		wantNetwork := &Network{
			ID:             "88888888-8888-8888-8888-888888888888",
			Name:           "dedicated_net",
			RouterID:       "22222222-2222-2222-2222-222222222222",
			ZoneID:         "99999999-9999-9999-9999-999999999999",
			Status:         "ACTIVE",
			CreatedAt:      "2025-12-05T17:00:13.637735",
			UpdatedAt:      "2025-12-05T17:00:25.690283",
			AccountID:      "777777",
			ProjectID:      "",
			Tags:           []string{"tag22"},
			OsNetworkID:    "",
			Vlan:           16461,
			InnerVlan:      321,
			NetopsVlanUUID: "88888888-8888-8888-8888-888888888888",
			SvNetworkID:    "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
			VdcName:        "",
		}
		require.Equal(t, wantNetwork, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.DedicatedNetworkCreate(context.Background(), &DedicatedNetworkCreateRequest{})

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
		plan, respRes, err := client.DedicatedNetworkCreate(context.Background(), &DedicatedNetworkCreateRequest{})

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
		plan, respRes, err := client.DedicatedNetworkCreate(context.Background(), &DedicatedNetworkCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_NetworkUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "11111111-1111-1111-1111-111111111111",
			"name": "updated_cloud_net",
			"vlan": 21412,
			"router_id": "22222222-2222-2222-2222-222222222222",
			"zone_id": "33333333-3333-3333-3333-333333333333",
			"project_id": "77777777777777777777777777777777",
			"vdc_name": null,
			"updated_at": null,
			"created_at": "2025-12-08T18:57:57.412757",
			"status": "ACTIVE",
			"account_id": "777777",
			"sv_network_id": "44444444-4444-4444-4444-444444444444",
			"netops_vlan_uuid": "11111111-1111-1111-1111-111111111111",
			"tags": ["tag1"],
			"inner_vlan": null,
			"os_network_id": "55555555-5555-5555-5555-555555555555"
		}`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		newName := "updated_cloud_net"
		tags := []string{"tag1"}
		updateReq := &NetworkUpdateRequest{
			Name: &newName,
			Tags: &tags,
		}

		// Execute
		plan, respRes, err := client.NetworkUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", updateReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		wantNetwork := &Network{
			ID:             "11111111-1111-1111-1111-111111111111",
			Name:           "updated_cloud_net",
			RouterID:       "22222222-2222-2222-2222-222222222222",
			ZoneID:         "33333333-3333-3333-3333-333333333333",
			Status:         "ACTIVE",
			CreatedAt:      "2025-12-08T18:57:57.412757",
			UpdatedAt:      "",
			AccountID:      "777777",
			ProjectID:      "77777777777777777777777777777777",
			Tags:           []string{"tag1"},
			OsNetworkID:    "55555555-5555-5555-5555-555555555555",
			Vlan:           21412,
			InnerVlan:      0,
			NetopsVlanUUID: "11111111-1111-1111-1111-111111111111",
			SvNetworkID:    "44444444-4444-4444-4444-444444444444",
			VdcName:        "",
		}
		require.Equal(t, wantNetwork, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.NetworkUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", &NetworkUpdateRequest{})

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
		plan, respRes, err := client.NetworkUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", &NetworkUpdateRequest{})

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
		plan, respRes, err := client.NetworkUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", &NetworkUpdateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_NetworkDisconnect(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		fakeResp := NewFakeResponse(204, "") //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		respRes, err := client.NetworkDisconnect(context.Background(), "11111111-1111-1111-1111-111111111111")

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
		respRes, err := client.NetworkDisconnect(context.Background(), "11111111-1111-1111-1111-111111111111")

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
		respRes, err := client.NetworkDisconnect(context.Background(), "11111111-1111-1111-1111-111111111111")

		// Analyse
		require.Error(t, err)
		require.Nil(t, respRes)
	})
}

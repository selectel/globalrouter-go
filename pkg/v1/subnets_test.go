package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/selectel/globalrouter-go/pkg/httptest"
	"github.com/stretchr/testify/require"
)

func TestServiceClient_Subnets(t *testing.T) {
	t.Run("SuccessWithQuery", func(t *testing.T) {
		// Prepare
		body := `[
			{
				"id": "33333333-3333-3333-3333-333333333333",
				"name": "subnet_cloud",
				"gateway": "10.0.14.1",
				"cidr": "10.0.14.0/24",
				"network_id": "44444444-4444-4444-4444-444444444444",
				"updated_at": "2025-12-12T09:57:48.609636",
				"created_at": "2025-12-12T09:57:27.524285",
				"status": "ACTIVE",
				"account_id": "777777",
				"project_id": null,
				"service_addresses": [
					"10.0.14.253",
					"10.0.14.254"
				],
				"os_subnet_id": "55555555-5555-5555-5555-555555555555",
				"sv_subnet_id": "77777777-7777-7777-7777-777777777777",
				"netops_subnet_id": "33333333-3333-3333-3333-333333333333",
				"tags": ["tag2"]
			}
		]`
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Subnets(context.Background(), &SubnetsQueryParams{Filters: SubnetsFilters{Name: "subnet_cloud"}})

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &[]Subnet{
			{
				ID:               "33333333-3333-3333-3333-333333333333",
				Name:             "subnet_cloud",
				NetworkID:        "44444444-4444-4444-4444-444444444444",
				Gateway:          "10.0.14.1",
				Cidr:             "10.0.14.0/24",
				ServiceAddresses: []string{"10.0.14.253", "10.0.14.254"},
				OsSubnetID:       "55555555-5555-5555-5555-555555555555",
				Status:           "ACTIVE",
				CreatedAt:        "2025-12-12T09:57:27.524285",
				UpdatedAt:        "2025-12-12T09:57:48.609636",
				AccountID:        "777777",
				ProjectID:        "",
				Tags:             []string{"tag2"},
				NetopsSubnetID:   "33333333-3333-3333-3333-333333333333",
				SvSubnetID:       "77777777-7777-7777-7777-777777777777",
			},
		}
		require.Equal(t, want, res)
	})

	t.Run("SuccessWithoutQuery", func(t *testing.T) {
		// Prepare
		body := `[
			{
				"id": "33333333-3333-3333-3333-333333333333",
				"name": "subnet_cloud",
				"gateway": "10.0.14.1",
				"cidr": "10.0.14.0/24",
				"network_id": "44444444-4444-4444-4444-444444444444",
				"updated_at": "2025-12-12T09:57:48.609636",
				"created_at": "2025-12-12T09:57:27.524285",
				"status": "ACTIVE",
				"account_id": "777777",
				"project_id": null,
				"service_addresses": [
					"10.0.14.253",
					"10.0.14.254"
				],
				"os_subnet_id": "55555555-5555-5555-5555-555555555555",
				"sv_subnet_id": "77777777-7777-7777-7777-777777777777",
				"netops_subnet_id": "33333333-3333-3333-3333-333333333333",
				"tags": ["tag2"]
			},
			{
				"id": "11111111-1111-1111-1111-111111111111",
				"name": "subnet_dedicated",
				"gateway": "192.168.0.1",
				"cidr": "192.168.0.0/24",
				"network_id": "22222222-2222-2222-2222-222222222222",
				"updated_at": "2025-12-05T17:07:10.074490",
				"created_at": "2025-12-05T17:06:58.204690",
				"status": "ACTIVE",
				"account_id": "777777",
				"project_id": null,
				"service_addresses": [
					"192.168.0.253",
					"192.168.0.254"
				],
				"os_subnet_id": null,
				"sv_subnet_id": null,
				"netops_subnet_id": "11111111-1111-1111-1111-111111111111",
				"tags": ["tag1"]
			}
		]`
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Subnets(context.Background(), nil)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &[]Subnet{
			{
				ID:               "33333333-3333-3333-3333-333333333333",
				Name:             "subnet_cloud",
				NetworkID:        "44444444-4444-4444-4444-444444444444",
				Gateway:          "10.0.14.1",
				Cidr:             "10.0.14.0/24",
				ServiceAddresses: []string{"10.0.14.253", "10.0.14.254"},
				OsSubnetID:       "55555555-5555-5555-5555-555555555555",
				Status:           "ACTIVE",
				CreatedAt:        "2025-12-12T09:57:27.524285",
				UpdatedAt:        "2025-12-12T09:57:48.609636",
				AccountID:        "777777",
				ProjectID:        "",
				Tags:             []string{"tag2"},
				NetopsSubnetID:   "33333333-3333-3333-3333-333333333333",
				SvSubnetID:       "77777777-7777-7777-7777-777777777777",
			},
			{
				ID:               "11111111-1111-1111-1111-111111111111",
				Name:             "subnet_dedicated",
				NetworkID:        "22222222-2222-2222-2222-222222222222",
				Gateway:          "192.168.0.1",
				Cidr:             "192.168.0.0/24",
				ServiceAddresses: []string{"192.168.0.253", "192.168.0.254"},
				OsSubnetID:       "",
				Status:           "ACTIVE",
				CreatedAt:        "2025-12-05T17:06:58.204690",
				UpdatedAt:        "2025-12-05T17:07:10.074490",
				AccountID:        "777777",
				ProjectID:        "",
				Tags:             []string{"tag1"},
				NetopsSubnetID:   "11111111-1111-1111-1111-111111111111",
				SvSubnetID:       "",
			},
		}
		require.Equal(t, want, res)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Subnets(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := httptest.NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", httptest.NewFakeTransport(fakeResp, nil))

		// Execute
		res, respRes, err := client.Subnets(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, res)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := httptest.NewFakeTransport(nil, errors.New("subnet failure"))
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Subnets(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_Subnet(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "33333333-3333-3333-3333-333333333333",
			"name": "subnet_cloud",
			"gateway": "10.0.14.1",
			"cidr": "10.0.14.0/24",
			"network_id": "44444444-4444-4444-4444-444444444444",
			"updated_at": "2025-12-12T09:57:48.609636",
			"created_at": "2025-12-12T09:57:27.524285",
			"status": "ACTIVE",
			"account_id": "777777",
			"project_id": null,
			"service_addresses": [
				"10.0.14.253",
				"10.0.14.254"
			],
			"os_subnet_id": "55555555-5555-5555-5555-555555555555",
			"sv_subnet_id": "77777777-7777-7777-7777-777777777777",
			"netops_subnet_id": "33333333-3333-3333-3333-333333333333",
			"tags": ["tag2"]
		}`
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.Subnet(context.Background(), "plan-id-1")

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		wantSubnet := &Subnet{
			ID:               "33333333-3333-3333-3333-333333333333",
			Name:             "subnet_cloud",
			NetworkID:        "44444444-4444-4444-4444-444444444444",
			Gateway:          "10.0.14.1",
			Cidr:             "10.0.14.0/24",
			ServiceAddresses: []string{"10.0.14.253", "10.0.14.254"},
			OsSubnetID:       "55555555-5555-5555-5555-555555555555",
			Status:           "ACTIVE",
			CreatedAt:        "2025-12-12T09:57:27.524285",
			UpdatedAt:        "2025-12-12T09:57:48.609636",
			AccountID:        "777777",
			ProjectID:        "",
			Tags:             []string{"tag2"},
			NetopsSubnetID:   "33333333-3333-3333-3333-333333333333",
			SvSubnetID:       "77777777-7777-7777-7777-777777777777",
		}
		require.Equal(t, wantSubnet, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.Subnet(context.Background(), "plan-id-1")

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := httptest.NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", httptest.NewFakeTransport(fakeResp, nil))

		// Execute
		plan, respRes, err := client.Subnet(context.Background(), "plan-id-1")

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, plan)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := httptest.NewFakeTransport(nil, errors.New("subnet failure"))
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.Subnet(context.Background(), "plan-id-1")

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_VPCSubnetCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "33333333-3333-3333-3333-333333333333",
			"name": "subnet_cloud",
			"gateway": "10.0.14.1",
			"cidr": "10.0.14.0/24",
			"network_id": "44444444-4444-4444-4444-444444444444",
			"updated_at": "2025-12-12T09:57:48.609636",
			"created_at": "2025-12-12T09:57:27.524285",
			"status": "ACTIVE",
			"account_id": "777777",
			"project_id": null,
			"service_addresses": [
				"10.0.14.253",
				"10.0.14.254"
			],
			"os_subnet_id": "55555555-5555-5555-5555-555555555555",
			"sv_subnet_id": "77777777-7777-7777-7777-777777777777",
			"netops_subnet_id": "33333333-3333-3333-3333-333333333333",
			"tags": ["tag2"]
		}`
		fakeResp := httptest.NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		createReq := &VPCSubnetCreateRequest{
			NetworkID:        "22222222-2222-2222-2222-222222222222",
			Gateway:          "10.0.14.1",
			Cidr:             "10.0.14.0/24",
			ServiceAddresses: []string{"10.0.14.253", "10.0.14.254"},
			Name:             "subnet_cloud",
			Tags:             []string{"tag2"},
		}

		// Execute
		plan, respRes, err := client.VPCSubnetCreate(context.Background(), createReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 201, respRes.StatusCode)
		wantSubnet := &Subnet{
			ID:               "33333333-3333-3333-3333-333333333333",
			Name:             "subnet_cloud",
			NetworkID:        "44444444-4444-4444-4444-444444444444",
			Gateway:          "10.0.14.1",
			Cidr:             "10.0.14.0/24",
			ServiceAddresses: []string{"10.0.14.253", "10.0.14.254"},
			OsSubnetID:       "55555555-5555-5555-5555-555555555555",
			Status:           "ACTIVE",
			CreatedAt:        "2025-12-12T09:57:27.524285",
			UpdatedAt:        "2025-12-12T09:57:48.609636",
			AccountID:        "777777",
			ProjectID:        "",
			Tags:             []string{"tag2"},
			NetopsSubnetID:   "33333333-3333-3333-3333-333333333333",
			SvSubnetID:       "77777777-7777-7777-7777-777777777777",
		}
		require.Equal(t, wantSubnet, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := httptest.NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.VPCSubnetCreate(context.Background(), &VPCSubnetCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.NotNil(t, respRes)
		require.Equal(t, 201, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := httptest.NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", httptest.NewFakeTransport(fakeResp, nil))

		// Execute
		plan, respRes, err := client.VPCSubnetCreate(context.Background(), &VPCSubnetCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, plan)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := httptest.NewFakeTransport(nil, errors.New("subnet failure"))
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.VPCSubnetCreate(context.Background(), &VPCSubnetCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_DedicatedSubnetCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "11111111-1111-1111-1111-111111111111",
			"name": "subnet_dedicated",
			"gateway": "192.168.0.1",
			"cidr": "192.168.0.0/24",
			"network_id": "22222222-2222-2222-2222-222222222222",
			"updated_at": "2025-12-05T17:07:10.074490",
			"created_at": "2025-12-05T17:06:58.204690",
			"status": "ACTIVE",
			"account_id": "777777",
			"project_id": null,
			"service_addresses": [
				"192.168.0.253",
				"192.168.0.254"
			],
			"os_subnet_id": null,
			"sv_subnet_id": null,
			"netops_subnet_id": "11111111-1111-1111-1111-111111111111",
			"tags": ["tag1"]
		}`
		fakeResp := httptest.NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		createReq := &DedicatedSubnetCreateRequest{
			NetworkID:        "22222222-2222-2222-2222-222222222222",
			Gateway:          "192.168.0.1",
			Cidr:             "192.168.0.0/24",
			ServiceAddresses: []string{"192.168.0.253", "192.168.0.254"},
			Name:             "subnet_dedicated",
			Tags:             []string{"tag1"},
		}

		// Execute
		plan, respRes, err := client.DedicatedSubnetCreate(context.Background(), createReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 201, respRes.StatusCode)
		wantSubnet := &Subnet{
			ID:               "11111111-1111-1111-1111-111111111111",
			Name:             "subnet_dedicated",
			NetworkID:        "22222222-2222-2222-2222-222222222222",
			Gateway:          "192.168.0.1",
			Cidr:             "192.168.0.0/24",
			ServiceAddresses: []string{"192.168.0.253", "192.168.0.254"},
			OsSubnetID:       "",
			Status:           "ACTIVE",
			CreatedAt:        "2025-12-05T17:06:58.204690",
			UpdatedAt:        "2025-12-05T17:07:10.074490",
			AccountID:        "777777",
			ProjectID:        "",
			Tags:             []string{"tag1"},
			NetopsSubnetID:   "11111111-1111-1111-1111-111111111111",
			SvSubnetID:       "",
		}
		require.Equal(t, wantSubnet, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := httptest.NewFakeResponse(201, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.DedicatedSubnetCreate(context.Background(), &DedicatedSubnetCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.NotNil(t, respRes)
		require.Equal(t, 201, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := httptest.NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", httptest.NewFakeTransport(fakeResp, nil))

		// Execute
		plan, respRes, err := client.DedicatedSubnetCreate(context.Background(), &DedicatedSubnetCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, plan)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := httptest.NewFakeTransport(nil, errors.New("subnet failure"))
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.DedicatedSubnetCreate(context.Background(), &DedicatedSubnetCreateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_SubnetUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		body := `{
			"id": "11111111-1111-1111-1111-111111111111",
			"name": "subnet_dedicated",
			"gateway": "192.168.0.1",
			"cidr": "192.168.0.0/24",
			"network_id": "22222222-2222-2222-2222-222222222222",
			"updated_at": "2025-12-05T17:07:10.074490",
			"created_at": "2025-12-05T17:06:58.204690",
			"status": "ACTIVE",
			"account_id": "777777",
			"project_id": null,
			"service_addresses": [
				"192.168.0.253",
				"192.168.0.254"
			],
			"os_subnet_id": null,
			"sv_subnet_id": null,
			"netops_subnet_id": "11111111-1111-1111-1111-111111111111",
			"tags": ["tag1"]
		}`
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		newName := "updated_cloud_net"
		tags := []string{"tag1"}
		updateReq := &SubnetUpdateRequest{
			Name: &newName,
			Tags: &tags,
		}

		// Execute
		plan, respRes, err := client.SubnetUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", updateReq)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		wantSubnet := &Subnet{
			ID:               "11111111-1111-1111-1111-111111111111",
			Name:             "subnet_dedicated",
			NetworkID:        "22222222-2222-2222-2222-222222222222",
			Gateway:          "192.168.0.1",
			Cidr:             "192.168.0.0/24",
			ServiceAddresses: []string{"192.168.0.253", "192.168.0.254"},
			OsSubnetID:       "",
			Status:           "ACTIVE",
			CreatedAt:        "2025-12-05T17:06:58.204690",
			UpdatedAt:        "2025-12-05T17:07:10.074490",
			AccountID:        "777777",
			ProjectID:        "",
			Tags:             []string{"tag1"},
			NetopsSubnetID:   "11111111-1111-1111-1111-111111111111",
			SvSubnetID:       "",
		}
		require.Equal(t, wantSubnet, plan)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.SubnetUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", &SubnetUpdateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := httptest.NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", httptest.NewFakeTransport(fakeResp, nil))

		// Execute
		plan, respRes, err := client.SubnetUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", &SubnetUpdateRequest{})

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, plan)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := httptest.NewFakeTransport(nil, errors.New("subnet failure"))
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		plan, respRes, err := client.SubnetUpdate(context.Background(), "11111111-1111-1111-1111-111111111111", &SubnetUpdateRequest{})

		// Analyse
		require.Error(t, err)
		require.Nil(t, plan)
		require.Nil(t, respRes)
	})
}

func TestServiceClient_SubnetDisconnect(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Prepare
		fakeResp := httptest.NewFakeResponse(204, "") //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		respRes, err := client.SubnetDisconnect(context.Background(), "11111111-1111-1111-1111-111111111111")

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 204, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := httptest.NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", httptest.NewFakeTransport(fakeResp, nil))

		// Execute
		respRes, err := client.SubnetDisconnect(context.Background(), "11111111-1111-1111-1111-111111111111")

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := httptest.NewFakeTransport(nil, errors.New("subnet failure"))
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		respRes, err := client.SubnetDisconnect(context.Background(), "11111111-1111-1111-1111-111111111111")

		// Analyse
		require.Error(t, err)
		require.Nil(t, respRes)
	})
}

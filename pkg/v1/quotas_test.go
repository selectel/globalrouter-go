package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceClient_Quotas(t *testing.T) {
	t.Run("SuccessWithQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "123e4567-e89b-12d3-a456-426655440000",
			"name": "routers",
			"scope": "account_id",
			"scope_value": "12345",
			"limit": 10
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListQuotas(context.Background(), &QuotasQueryParams{Filters: QuotasFilters{Name: "routers", Scope: "account_id", ScopeValue: "12345"}})

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := []Quota{
			{
				ID:         "123e4567-e89b-12d3-a456-426655440000",
				Name:       "routers",
				Scope:      "account_id",
				ScopeValue: "12345",
				Limit:      10,
			},
		}
		require.Equal(t, &want, res)
	})

	t.Run("SuccessWithoutQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "123e4567-e89b-12d3-a456-426655440000",
			"name": "routers",
			"scope": "account_id",
			"scope_value": "12345",
			"limit": 10
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListQuotas(context.Background(), nil)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := []Quota{
			{
				ID:         "123e4567-e89b-12d3-a456-426655440000",
				Name:       "routers",
				Scope:      "account_id",
				ScopeValue: "12345",
				Limit:      10,
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
		res, respRes, err := client.ListQuotas(context.Background(), nil)

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
		res, respRes, err := client.ListQuotas(context.Background(), nil)

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
		res, respRes, err := client.ListQuotas(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.Nil(t, respRes)
	})
}

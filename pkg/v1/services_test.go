package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceClient_Services(t *testing.T) {
	t.Run("SuccessWithQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "123e4567-e89b-12d3-a456-426655440000",
			"name": "vpc",
			"extension": "vpc",
			"created_at": "2170-01-01 00:00:00"
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListServices(context.Background(), &ServicesQueryParams{Filters: ServicesFilters{Name: "vpc"}})

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := []Service{
			{
				ID:        "123e4567-e89b-12d3-a456-426655440000",
				Name:      "vpc",
				CreatedAt: "2170-01-01 00:00:00",
				Extension: "vpc",
			},
		}
		require.Equal(t, &want, res)
	})

	t.Run("SuccessWithoutQuery", func(t *testing.T) {
		// Prepare
		body := `[{
			"id": "123e4567-e89b-12d3-a456-426655440000",
			"name": "vpc",
			"extension": "vpc",
			"created_at": "2170-01-01 00:00:00"
		}]`
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListServices(context.Background(), nil)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := []Service{
			{
				ID:        "123e4567-e89b-12d3-a456-426655440000",
				Name:      "vpc",
				CreatedAt: "2170-01-01 00:00:00",
				Extension: "vpc",
			},
		}
		require.Equal(t, &want, res)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := &NewFakeTransport{resp: fakeResp, err: nil}
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.ListServices(context.Background(), nil)

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
		res, respRes, err := client.ListServices(context.Background(), nil)

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
		res, respRes, err := client.ListServices(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.Nil(t, respRes)
	})
}

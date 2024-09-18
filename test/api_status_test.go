/*
API Documentation

Testing StatusAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package nautobot

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/nautobot/go-nautobot/v2"
)

func Test_nautobot_StatusAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StatusAPIService StatusRetrieve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StatusAPI.StatusRetrieve(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

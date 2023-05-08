/*
Puupee API

Testing AbpTenantApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package puupee

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/puupee/puupee-api-go"
)

func Test_puupee_AbpTenantApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AbpTenantApiService ApiAbpMultiTenancyTenantsByIdIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AbpTenantApi.ApiAbpMultiTenancyTenantsByIdIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AbpTenantApiService ApiAbpMultiTenancyTenantsByNameNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.AbpTenantApi.ApiAbpMultiTenancyTenantsByNameNameGet(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

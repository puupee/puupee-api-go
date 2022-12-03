/*
Puupee API

Testing TenantApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package puupee

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_puupee_TenantApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test TenantApiService ApiMultiTenancyTenantsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantApi.ApiMultiTenancyTenantsGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantApiService ApiMultiTenancyTenantsIdDefaultConnectionStringDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringDelete(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantApiService ApiMultiTenancyTenantsIdDefaultConnectionStringGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringGet(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantApiService ApiMultiTenancyTenantsIdDefaultConnectionStringPut", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdDefaultConnectionStringPut(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantApiService ApiMultiTenancyTenantsIdDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdDelete(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantApiService ApiMultiTenancyTenantsIdGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdGet(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantApiService ApiMultiTenancyTenantsIdPut", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.TenantApi.ApiMultiTenancyTenantsIdPut(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantApiService ApiMultiTenancyTenantsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantApi.ApiMultiTenancyTenantsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}

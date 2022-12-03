/*
Puupee API

Testing StorageObjectApiService

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

func Test_puupee_StorageObjectApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test StorageObjectApiService ApiAppStorageObjectFileGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.StorageObjectApi.ApiAppStorageObjectFileGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StorageObjectApiService ApiAppStorageObjectFileOrCredentialsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.StorageObjectApi.ApiAppStorageObjectFileOrCredentialsGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StorageObjectApiService ApiAppStorageObjectPreSignUrlPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.StorageObjectApi.ApiAppStorageObjectPreSignUrlPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StorageObjectApiService ApiAppStorageObjectThumbGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.StorageObjectApi.ApiAppStorageObjectThumbGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}

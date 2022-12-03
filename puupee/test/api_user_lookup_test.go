/*
Puupee API

Testing UserLookupApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package doggy

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_doggy_UserLookupApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test UserLookupApiService ApiIdentityUsersLookupByUsernameUserNameGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var userName string

        resp, httpRes, err := apiClient.UserLookupApi.ApiIdentityUsersLookupByUsernameUserNameGet(context.Background(), userName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test UserLookupApiService ApiIdentityUsersLookupCountGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.UserLookupApi.ApiIdentityUsersLookupCountGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test UserLookupApiService ApiIdentityUsersLookupIdGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.UserLookupApi.ApiIdentityUsersLookupIdGet(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test UserLookupApiService ApiIdentityUsersLookupSearchGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.UserLookupApi.ApiIdentityUsersLookupSearchGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}

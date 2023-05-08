/*
Puupee API

Testing SyncStateApiService

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

func Test_puupee_SyncStateApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SyncStateApiService ApiAppSyncStateGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SyncStateApi.ApiAppSyncStateGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SyncStateApiService ApiAppSyncStatePuupeeChangedEtoPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SyncStateApi.ApiAppSyncStatePuupeeChangedEtoPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

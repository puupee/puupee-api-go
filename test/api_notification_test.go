/*
Puupee API

Testing NotificationApiService

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

func Test_puupee_NotificationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NotificationApiService ApiAppNotificationBarkApiKeyMessageGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiKey string
		var message string

		httpRes, err := apiClient.NotificationApi.ApiAppNotificationBarkApiKeyMessageGet(context.Background(), apiKey, message).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationApiService ApiAppNotificationGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NotificationApi.ApiAppNotificationGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationApiService ApiAppNotificationPushPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.NotificationApi.ApiAppNotificationPushPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

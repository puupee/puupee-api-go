/*
Puupee API

Testing NotificationAPIService

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

func Test_puupee_NotificationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NotificationAPIService Bark", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiKey string
		var message string

		httpRes, err := apiClient.NotificationAPI.Bark(context.Background(), apiKey, message).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationAPIService GetNotificationList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NotificationAPI.GetNotificationList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationAPIService Push", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.NotificationAPI.Push(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

/*
Puupee API

Testing EmailSettingsApiService

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

func Test_puupee_EmailSettingsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EmailSettingsApiService ApiSettingManagementEmailingGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailSettingsApi.ApiSettingManagementEmailingGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailSettingsApiService ApiSettingManagementEmailingPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.EmailSettingsApi.ApiSettingManagementEmailingPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailSettingsApiService ApiSettingManagementEmailingSendTestEmailPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.EmailSettingsApi.ApiSettingManagementEmailingSendTestEmailPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

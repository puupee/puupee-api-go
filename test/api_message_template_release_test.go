/*
Puupee API

Testing MessageTemplateReleaseApiService

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

func Test_puupee_MessageTemplateReleaseApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MessageTemplateReleaseApiService ApiAppMessageTemplateReleaseGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MessageTemplateReleaseApi.ApiAppMessageTemplateReleaseGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MessageTemplateReleaseApiService ApiAppMessageTemplateReleaseIdGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.MessageTemplateReleaseApi.ApiAppMessageTemplateReleaseIdGet(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MessageTemplateReleaseApiService ApiAppMessageTemplateReleasePost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MessageTemplateReleaseApi.ApiAppMessageTemplateReleasePost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
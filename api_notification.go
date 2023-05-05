/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// NotificationApiService NotificationApi service
type NotificationApiService service

type ApiApiAppNotificationBarkApiKeyMessageGetRequest struct {
	ctx context.Context
	ApiService *NotificationApiService
	apiKey string
	message string
	automaticallyCopy *int32
	copy *string
	url *string
	isArchive *string
	group *string
	icon *string
	name *string
	value *string
}

func (r ApiApiAppNotificationBarkApiKeyMessageGetRequest) AutomaticallyCopy(automaticallyCopy int32) ApiApiAppNotificationBarkApiKeyMessageGetRequest {
	r.automaticallyCopy = &automaticallyCopy
	return r
}

func (r ApiApiAppNotificationBarkApiKeyMessageGetRequest) Copy(copy string) ApiApiAppNotificationBarkApiKeyMessageGetRequest {
	r.copy = &copy
	return r
}

func (r ApiApiAppNotificationBarkApiKeyMessageGetRequest) Url(url string) ApiApiAppNotificationBarkApiKeyMessageGetRequest {
	r.url = &url
	return r
}

func (r ApiApiAppNotificationBarkApiKeyMessageGetRequest) IsArchive(isArchive string) ApiApiAppNotificationBarkApiKeyMessageGetRequest {
	r.isArchive = &isArchive
	return r
}

func (r ApiApiAppNotificationBarkApiKeyMessageGetRequest) Group(group string) ApiApiAppNotificationBarkApiKeyMessageGetRequest {
	r.group = &group
	return r
}

func (r ApiApiAppNotificationBarkApiKeyMessageGetRequest) Icon(icon string) ApiApiAppNotificationBarkApiKeyMessageGetRequest {
	r.icon = &icon
	return r
}

func (r ApiApiAppNotificationBarkApiKeyMessageGetRequest) Name(name string) ApiApiAppNotificationBarkApiKeyMessageGetRequest {
	r.name = &name
	return r
}

func (r ApiApiAppNotificationBarkApiKeyMessageGetRequest) Value(value string) ApiApiAppNotificationBarkApiKeyMessageGetRequest {
	r.value = &value
	return r
}

func (r ApiApiAppNotificationBarkApiKeyMessageGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiAppNotificationBarkApiKeyMessageGetExecute(r)
}

/*
ApiAppNotificationBarkApiKeyMessageGet Method for ApiAppNotificationBarkApiKeyMessageGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiKey
 @param message
 @return ApiApiAppNotificationBarkApiKeyMessageGetRequest
*/
func (a *NotificationApiService) ApiAppNotificationBarkApiKeyMessageGet(ctx context.Context, apiKey string, message string) ApiApiAppNotificationBarkApiKeyMessageGetRequest {
	return ApiApiAppNotificationBarkApiKeyMessageGetRequest{
		ApiService: a,
		ctx: ctx,
		apiKey: apiKey,
		message: message,
	}
}

// Execute executes the request
func (a *NotificationApiService) ApiAppNotificationBarkApiKeyMessageGetExecute(r ApiApiAppNotificationBarkApiKeyMessageGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.ApiAppNotificationBarkApiKeyMessageGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/app/notification/bark/{apiKey}/{message}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiKey"+"}", url.PathEscape(parameterValueToString(r.apiKey, "apiKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"message"+"}", url.PathEscape(parameterValueToString(r.message, "message")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.automaticallyCopy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "automaticallyCopy", r.automaticallyCopy, "")
	}
	if r.copy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "copy", r.copy, "")
	}
	if r.url != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "url", r.url, "")
	}
	if r.isArchive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isArchive", r.isArchive, "")
	}
	if r.group != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group", r.group, "")
	}
	if r.icon != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "icon", r.icon, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Name", r.name, "")
	}
	if r.value != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Value", r.value, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiAppNotificationGetRequest struct {
	ctx context.Context
	ApiService *NotificationApiService
	sorting *string
	skipCount *int32
	maxResultCount *int32
}

func (r ApiApiAppNotificationGetRequest) Sorting(sorting string) ApiApiAppNotificationGetRequest {
	r.sorting = &sorting
	return r
}

func (r ApiApiAppNotificationGetRequest) SkipCount(skipCount int32) ApiApiAppNotificationGetRequest {
	r.skipCount = &skipCount
	return r
}

func (r ApiApiAppNotificationGetRequest) MaxResultCount(maxResultCount int32) ApiApiAppNotificationGetRequest {
	r.maxResultCount = &maxResultCount
	return r
}

func (r ApiApiAppNotificationGetRequest) Execute() (*NotificationInfoDtoPagedResultDto, *http.Response, error) {
	return r.ApiService.ApiAppNotificationGetExecute(r)
}

/*
ApiAppNotificationGet Method for ApiAppNotificationGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiAppNotificationGetRequest
*/
func (a *NotificationApiService) ApiAppNotificationGet(ctx context.Context) ApiApiAppNotificationGetRequest {
	return ApiApiAppNotificationGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NotificationInfoDtoPagedResultDto
func (a *NotificationApiService) ApiAppNotificationGetExecute(r ApiApiAppNotificationGetRequest) (*NotificationInfoDtoPagedResultDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotificationInfoDtoPagedResultDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.ApiAppNotificationGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/app/notification"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sorting != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Sorting", r.sorting, "")
	}
	if r.skipCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SkipCount", r.skipCount, "")
	}
	if r.maxResultCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "MaxResultCount", r.maxResultCount, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiAppNotificationPushPostRequest struct {
	ctx context.Context
	ApiService *NotificationApiService
	body *CreatePushNotificationDto
}

func (r ApiApiAppNotificationPushPostRequest) Body(body CreatePushNotificationDto) ApiApiAppNotificationPushPostRequest {
	r.body = &body
	return r
}

func (r ApiApiAppNotificationPushPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiAppNotificationPushPostExecute(r)
}

/*
ApiAppNotificationPushPost Method for ApiAppNotificationPushPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiAppNotificationPushPostRequest
*/
func (a *NotificationApiService) ApiAppNotificationPushPost(ctx context.Context) ApiApiAppNotificationPushPostRequest {
	return ApiApiAppNotificationPushPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NotificationApiService) ApiAppNotificationPushPostExecute(r ApiApiAppNotificationPushPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationApiService.ApiAppNotificationPushPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/app/notification/push"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RemoteServiceErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

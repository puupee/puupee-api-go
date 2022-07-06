# CreateUpdateNotificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**NotificationInfoId** | Pointer to **string** |  | [optional] 
**NotificationMethod** | Pointer to **NullableString** |  | [optional] 
**Success** | Pointer to **NullableBool** |  | [optional] 
**CompletionTime** | Pointer to **NullableTime** |  | [optional] 
**FailureReason** | Pointer to **NullableString** |  | [optional] 
**RetryNotificationId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateUpdateNotificationDto

`func NewCreateUpdateNotificationDto() *CreateUpdateNotificationDto`

NewCreateUpdateNotificationDto instantiates a new CreateUpdateNotificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateNotificationDtoWithDefaults

`func NewCreateUpdateNotificationDtoWithDefaults() *CreateUpdateNotificationDto`

NewCreateUpdateNotificationDtoWithDefaults instantiates a new CreateUpdateNotificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateUpdateNotificationDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateUpdateNotificationDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateUpdateNotificationDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateUpdateNotificationDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetNotificationInfoId

`func (o *CreateUpdateNotificationDto) GetNotificationInfoId() string`

GetNotificationInfoId returns the NotificationInfoId field if non-nil, zero value otherwise.

### GetNotificationInfoIdOk

`func (o *CreateUpdateNotificationDto) GetNotificationInfoIdOk() (*string, bool)`

GetNotificationInfoIdOk returns a tuple with the NotificationInfoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationInfoId

`func (o *CreateUpdateNotificationDto) SetNotificationInfoId(v string)`

SetNotificationInfoId sets NotificationInfoId field to given value.

### HasNotificationInfoId

`func (o *CreateUpdateNotificationDto) HasNotificationInfoId() bool`

HasNotificationInfoId returns a boolean if a field has been set.

### GetNotificationMethod

`func (o *CreateUpdateNotificationDto) GetNotificationMethod() string`

GetNotificationMethod returns the NotificationMethod field if non-nil, zero value otherwise.

### GetNotificationMethodOk

`func (o *CreateUpdateNotificationDto) GetNotificationMethodOk() (*string, bool)`

GetNotificationMethodOk returns a tuple with the NotificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMethod

`func (o *CreateUpdateNotificationDto) SetNotificationMethod(v string)`

SetNotificationMethod sets NotificationMethod field to given value.

### HasNotificationMethod

`func (o *CreateUpdateNotificationDto) HasNotificationMethod() bool`

HasNotificationMethod returns a boolean if a field has been set.

### SetNotificationMethodNil

`func (o *CreateUpdateNotificationDto) SetNotificationMethodNil(b bool)`

 SetNotificationMethodNil sets the value for NotificationMethod to be an explicit nil

### UnsetNotificationMethod
`func (o *CreateUpdateNotificationDto) UnsetNotificationMethod()`

UnsetNotificationMethod ensures that no value is present for NotificationMethod, not even an explicit nil
### GetSuccess

`func (o *CreateUpdateNotificationDto) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateUpdateNotificationDto) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateUpdateNotificationDto) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateUpdateNotificationDto) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### SetSuccessNil

`func (o *CreateUpdateNotificationDto) SetSuccessNil(b bool)`

 SetSuccessNil sets the value for Success to be an explicit nil

### UnsetSuccess
`func (o *CreateUpdateNotificationDto) UnsetSuccess()`

UnsetSuccess ensures that no value is present for Success, not even an explicit nil
### GetCompletionTime

`func (o *CreateUpdateNotificationDto) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *CreateUpdateNotificationDto) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *CreateUpdateNotificationDto) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *CreateUpdateNotificationDto) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### SetCompletionTimeNil

`func (o *CreateUpdateNotificationDto) SetCompletionTimeNil(b bool)`

 SetCompletionTimeNil sets the value for CompletionTime to be an explicit nil

### UnsetCompletionTime
`func (o *CreateUpdateNotificationDto) UnsetCompletionTime()`

UnsetCompletionTime ensures that no value is present for CompletionTime, not even an explicit nil
### GetFailureReason

`func (o *CreateUpdateNotificationDto) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *CreateUpdateNotificationDto) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *CreateUpdateNotificationDto) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *CreateUpdateNotificationDto) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *CreateUpdateNotificationDto) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *CreateUpdateNotificationDto) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetRetryNotificationId

`func (o *CreateUpdateNotificationDto) GetRetryNotificationId() string`

GetRetryNotificationId returns the RetryNotificationId field if non-nil, zero value otherwise.

### GetRetryNotificationIdOk

`func (o *CreateUpdateNotificationDto) GetRetryNotificationIdOk() (*string, bool)`

GetRetryNotificationIdOk returns a tuple with the RetryNotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryNotificationId

`func (o *CreateUpdateNotificationDto) SetRetryNotificationId(v string)`

SetRetryNotificationId sets RetryNotificationId field to given value.

### HasRetryNotificationId

`func (o *CreateUpdateNotificationDto) HasRetryNotificationId() bool`

HasRetryNotificationId returns a boolean if a field has been set.

### SetRetryNotificationIdNil

`func (o *CreateUpdateNotificationDto) SetRetryNotificationIdNil(b bool)`

 SetRetryNotificationIdNil sets the value for RetryNotificationId to be an explicit nil

### UnsetRetryNotificationId
`func (o *CreateUpdateNotificationDto) UnsetRetryNotificationId()`

UnsetRetryNotificationId ensures that no value is present for RetryNotificationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



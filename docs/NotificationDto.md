# NotificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**NotificationInfoId** | Pointer to **string** |  | [optional] 
**NotificationMethod** | Pointer to **NullableString** |  | [optional] 
**Success** | Pointer to **NullableBool** |  | [optional] 
**CompletionTime** | Pointer to **NullableTime** |  | [optional] 
**FailureReason** | Pointer to **NullableString** |  | [optional] 
**RetryNotificationId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNotificationDto

`func NewNotificationDto() *NotificationDto`

NewNotificationDto instantiates a new NotificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDtoWithDefaults

`func NewNotificationDtoWithDefaults() *NotificationDto`

NewNotificationDtoWithDefaults instantiates a new NotificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *NotificationDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *NotificationDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *NotificationDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *NotificationDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *NotificationDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *NotificationDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *NotificationDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *NotificationDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *NotificationDto) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *NotificationDto) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetUserId

`func (o *NotificationDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *NotificationDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *NotificationDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *NotificationDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetNotificationInfoId

`func (o *NotificationDto) GetNotificationInfoId() string`

GetNotificationInfoId returns the NotificationInfoId field if non-nil, zero value otherwise.

### GetNotificationInfoIdOk

`func (o *NotificationDto) GetNotificationInfoIdOk() (*string, bool)`

GetNotificationInfoIdOk returns a tuple with the NotificationInfoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationInfoId

`func (o *NotificationDto) SetNotificationInfoId(v string)`

SetNotificationInfoId sets NotificationInfoId field to given value.

### HasNotificationInfoId

`func (o *NotificationDto) HasNotificationInfoId() bool`

HasNotificationInfoId returns a boolean if a field has been set.

### GetNotificationMethod

`func (o *NotificationDto) GetNotificationMethod() string`

GetNotificationMethod returns the NotificationMethod field if non-nil, zero value otherwise.

### GetNotificationMethodOk

`func (o *NotificationDto) GetNotificationMethodOk() (*string, bool)`

GetNotificationMethodOk returns a tuple with the NotificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMethod

`func (o *NotificationDto) SetNotificationMethod(v string)`

SetNotificationMethod sets NotificationMethod field to given value.

### HasNotificationMethod

`func (o *NotificationDto) HasNotificationMethod() bool`

HasNotificationMethod returns a boolean if a field has been set.

### SetNotificationMethodNil

`func (o *NotificationDto) SetNotificationMethodNil(b bool)`

 SetNotificationMethodNil sets the value for NotificationMethod to be an explicit nil

### UnsetNotificationMethod
`func (o *NotificationDto) UnsetNotificationMethod()`

UnsetNotificationMethod ensures that no value is present for NotificationMethod, not even an explicit nil
### GetSuccess

`func (o *NotificationDto) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *NotificationDto) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *NotificationDto) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *NotificationDto) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### SetSuccessNil

`func (o *NotificationDto) SetSuccessNil(b bool)`

 SetSuccessNil sets the value for Success to be an explicit nil

### UnsetSuccess
`func (o *NotificationDto) UnsetSuccess()`

UnsetSuccess ensures that no value is present for Success, not even an explicit nil
### GetCompletionTime

`func (o *NotificationDto) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *NotificationDto) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *NotificationDto) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *NotificationDto) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### SetCompletionTimeNil

`func (o *NotificationDto) SetCompletionTimeNil(b bool)`

 SetCompletionTimeNil sets the value for CompletionTime to be an explicit nil

### UnsetCompletionTime
`func (o *NotificationDto) UnsetCompletionTime()`

UnsetCompletionTime ensures that no value is present for CompletionTime, not even an explicit nil
### GetFailureReason

`func (o *NotificationDto) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *NotificationDto) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *NotificationDto) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *NotificationDto) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *NotificationDto) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *NotificationDto) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetRetryNotificationId

`func (o *NotificationDto) GetRetryNotificationId() string`

GetRetryNotificationId returns the RetryNotificationId field if non-nil, zero value otherwise.

### GetRetryNotificationIdOk

`func (o *NotificationDto) GetRetryNotificationIdOk() (*string, bool)`

GetRetryNotificationIdOk returns a tuple with the RetryNotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryNotificationId

`func (o *NotificationDto) SetRetryNotificationId(v string)`

SetRetryNotificationId sets RetryNotificationId field to given value.

### HasRetryNotificationId

`func (o *NotificationDto) HasRetryNotificationId() bool`

HasRetryNotificationId returns a boolean if a field has been set.

### SetRetryNotificationIdNil

`func (o *NotificationDto) SetRetryNotificationIdNil(b bool)`

 SetRetryNotificationIdNil sets the value for RetryNotificationId to be an explicit nil

### UnsetRetryNotificationId
`func (o *NotificationDto) UnsetRetryNotificationId()`

UnsetRetryNotificationId ensures that no value is present for RetryNotificationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



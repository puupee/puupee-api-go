# VerifyReceiptResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatorId** | Pointer to **string** |  | [optional] [readonly] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 
**LastModifierId** | Pointer to **string** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeleterId** | Pointer to **string** |  | [optional] 
**DeletionTime** | Pointer to **time.Time** |  | [optional] 
**Deleter** | Pointer to [**IdentityUser**](IdentityUser.md) |  | [optional] 
**Creator** | Pointer to [**IdentityUser**](IdentityUser.md) |  | [optional] 
**LastModifier** | Pointer to [**IdentityUser**](IdentityUser.md) |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**ReceiptData** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to [**AppPlatform**](AppPlatform.md) |  | [optional] 
**DeviceToken** | Pointer to **string** |  | [optional] 
**Ok** | Pointer to **bool** |  | [optional] 
**StatusCode** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ResultData** | Pointer to **string** |  | [optional] 
**RecordId** | Pointer to **string** |  | [optional] 
**AppleVerifyReceiptResult** | Pointer to [**AppleVerifyReceiptResult**](AppleVerifyReceiptResult.md) |  | [optional] 

## Methods

### NewVerifyReceiptResult

`func NewVerifyReceiptResult() *VerifyReceiptResult`

NewVerifyReceiptResult instantiates a new VerifyReceiptResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyReceiptResultWithDefaults

`func NewVerifyReceiptResultWithDefaults() *VerifyReceiptResult`

NewVerifyReceiptResultWithDefaults instantiates a new VerifyReceiptResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerifyReceiptResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifyReceiptResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifyReceiptResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerifyReceiptResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *VerifyReceiptResult) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VerifyReceiptResult) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VerifyReceiptResult) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *VerifyReceiptResult) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *VerifyReceiptResult) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *VerifyReceiptResult) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *VerifyReceiptResult) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *VerifyReceiptResult) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *VerifyReceiptResult) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *VerifyReceiptResult) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *VerifyReceiptResult) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *VerifyReceiptResult) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *VerifyReceiptResult) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *VerifyReceiptResult) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *VerifyReceiptResult) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *VerifyReceiptResult) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *VerifyReceiptResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *VerifyReceiptResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *VerifyReceiptResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *VerifyReceiptResult) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *VerifyReceiptResult) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *VerifyReceiptResult) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *VerifyReceiptResult) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *VerifyReceiptResult) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *VerifyReceiptResult) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *VerifyReceiptResult) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *VerifyReceiptResult) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *VerifyReceiptResult) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetDeleter

`func (o *VerifyReceiptResult) GetDeleter() IdentityUser`

GetDeleter returns the Deleter field if non-nil, zero value otherwise.

### GetDeleterOk

`func (o *VerifyReceiptResult) GetDeleterOk() (*IdentityUser, bool)`

GetDeleterOk returns a tuple with the Deleter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleter

`func (o *VerifyReceiptResult) SetDeleter(v IdentityUser)`

SetDeleter sets Deleter field to given value.

### HasDeleter

`func (o *VerifyReceiptResult) HasDeleter() bool`

HasDeleter returns a boolean if a field has been set.

### GetCreator

`func (o *VerifyReceiptResult) GetCreator() IdentityUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *VerifyReceiptResult) GetCreatorOk() (*IdentityUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *VerifyReceiptResult) SetCreator(v IdentityUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *VerifyReceiptResult) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetLastModifier

`func (o *VerifyReceiptResult) GetLastModifier() IdentityUser`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *VerifyReceiptResult) GetLastModifierOk() (*IdentityUser, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *VerifyReceiptResult) SetLastModifier(v IdentityUser)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *VerifyReceiptResult) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.

### GetOrderId

`func (o *VerifyReceiptResult) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *VerifyReceiptResult) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *VerifyReceiptResult) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *VerifyReceiptResult) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetReceiptData

`func (o *VerifyReceiptResult) GetReceiptData() string`

GetReceiptData returns the ReceiptData field if non-nil, zero value otherwise.

### GetReceiptDataOk

`func (o *VerifyReceiptResult) GetReceiptDataOk() (*string, bool)`

GetReceiptDataOk returns a tuple with the ReceiptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptData

`func (o *VerifyReceiptResult) SetReceiptData(v string)`

SetReceiptData sets ReceiptData field to given value.

### HasReceiptData

`func (o *VerifyReceiptResult) HasReceiptData() bool`

HasReceiptData returns a boolean if a field has been set.

### GetPlatform

`func (o *VerifyReceiptResult) GetPlatform() AppPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VerifyReceiptResult) GetPlatformOk() (*AppPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VerifyReceiptResult) SetPlatform(v AppPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *VerifyReceiptResult) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDeviceToken

`func (o *VerifyReceiptResult) GetDeviceToken() string`

GetDeviceToken returns the DeviceToken field if non-nil, zero value otherwise.

### GetDeviceTokenOk

`func (o *VerifyReceiptResult) GetDeviceTokenOk() (*string, bool)`

GetDeviceTokenOk returns a tuple with the DeviceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceToken

`func (o *VerifyReceiptResult) SetDeviceToken(v string)`

SetDeviceToken sets DeviceToken field to given value.

### HasDeviceToken

`func (o *VerifyReceiptResult) HasDeviceToken() bool`

HasDeviceToken returns a boolean if a field has been set.

### GetOk

`func (o *VerifyReceiptResult) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *VerifyReceiptResult) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *VerifyReceiptResult) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *VerifyReceiptResult) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetStatusCode

`func (o *VerifyReceiptResult) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *VerifyReceiptResult) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *VerifyReceiptResult) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *VerifyReceiptResult) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetMessage

`func (o *VerifyReceiptResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VerifyReceiptResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VerifyReceiptResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VerifyReceiptResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResultData

`func (o *VerifyReceiptResult) GetResultData() string`

GetResultData returns the ResultData field if non-nil, zero value otherwise.

### GetResultDataOk

`func (o *VerifyReceiptResult) GetResultDataOk() (*string, bool)`

GetResultDataOk returns a tuple with the ResultData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultData

`func (o *VerifyReceiptResult) SetResultData(v string)`

SetResultData sets ResultData field to given value.

### HasResultData

`func (o *VerifyReceiptResult) HasResultData() bool`

HasResultData returns a boolean if a field has been set.

### GetRecordId

`func (o *VerifyReceiptResult) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *VerifyReceiptResult) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *VerifyReceiptResult) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.

### HasRecordId

`func (o *VerifyReceiptResult) HasRecordId() bool`

HasRecordId returns a boolean if a field has been set.

### GetAppleVerifyReceiptResult

`func (o *VerifyReceiptResult) GetAppleVerifyReceiptResult() AppleVerifyReceiptResult`

GetAppleVerifyReceiptResult returns the AppleVerifyReceiptResult field if non-nil, zero value otherwise.

### GetAppleVerifyReceiptResultOk

`func (o *VerifyReceiptResult) GetAppleVerifyReceiptResultOk() (*AppleVerifyReceiptResult, bool)`

GetAppleVerifyReceiptResultOk returns a tuple with the AppleVerifyReceiptResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleVerifyReceiptResult

`func (o *VerifyReceiptResult) SetAppleVerifyReceiptResult(v AppleVerifyReceiptResult)`

SetAppleVerifyReceiptResult sets AppleVerifyReceiptResult field to given value.

### HasAppleVerifyReceiptResult

`func (o *VerifyReceiptResult) HasAppleVerifyReceiptResult() bool`

HasAppleVerifyReceiptResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



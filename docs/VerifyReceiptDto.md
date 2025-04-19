# VerifyReceiptDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 
**LastModifierId** | Pointer to **string** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeleterId** | Pointer to **string** |  | [optional] 
**DeletionTime** | Pointer to **time.Time** |  | [optional] 
**OrderId** | **string** |  | 
**ReceiptData** | **string** |  | 
**Platform** | [**AppPlatform**](AppPlatform.md) |  | 
**DeviceToken** | **string** |  | 

## Methods

### NewVerifyReceiptDto

`func NewVerifyReceiptDto(orderId string, receiptData string, platform AppPlatform, deviceToken string, ) *VerifyReceiptDto`

NewVerifyReceiptDto instantiates a new VerifyReceiptDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyReceiptDtoWithDefaults

`func NewVerifyReceiptDtoWithDefaults() *VerifyReceiptDto`

NewVerifyReceiptDtoWithDefaults instantiates a new VerifyReceiptDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerifyReceiptDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifyReceiptDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifyReceiptDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerifyReceiptDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *VerifyReceiptDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VerifyReceiptDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VerifyReceiptDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *VerifyReceiptDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *VerifyReceiptDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *VerifyReceiptDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *VerifyReceiptDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *VerifyReceiptDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *VerifyReceiptDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *VerifyReceiptDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *VerifyReceiptDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *VerifyReceiptDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *VerifyReceiptDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *VerifyReceiptDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *VerifyReceiptDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *VerifyReceiptDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *VerifyReceiptDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *VerifyReceiptDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *VerifyReceiptDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *VerifyReceiptDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *VerifyReceiptDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *VerifyReceiptDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *VerifyReceiptDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *VerifyReceiptDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *VerifyReceiptDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *VerifyReceiptDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *VerifyReceiptDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *VerifyReceiptDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetOrderId

`func (o *VerifyReceiptDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *VerifyReceiptDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *VerifyReceiptDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetReceiptData

`func (o *VerifyReceiptDto) GetReceiptData() string`

GetReceiptData returns the ReceiptData field if non-nil, zero value otherwise.

### GetReceiptDataOk

`func (o *VerifyReceiptDto) GetReceiptDataOk() (*string, bool)`

GetReceiptDataOk returns a tuple with the ReceiptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptData

`func (o *VerifyReceiptDto) SetReceiptData(v string)`

SetReceiptData sets ReceiptData field to given value.


### GetPlatform

`func (o *VerifyReceiptDto) GetPlatform() AppPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VerifyReceiptDto) GetPlatformOk() (*AppPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VerifyReceiptDto) SetPlatform(v AppPlatform)`

SetPlatform sets Platform field to given value.


### GetDeviceToken

`func (o *VerifyReceiptDto) GetDeviceToken() string`

GetDeviceToken returns the DeviceToken field if non-nil, zero value otherwise.

### GetDeviceTokenOk

`func (o *VerifyReceiptDto) GetDeviceTokenOk() (*string, bool)`

GetDeviceTokenOk returns a tuple with the DeviceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceToken

`func (o *VerifyReceiptDto) SetDeviceToken(v string)`

SetDeviceToken sets DeviceToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



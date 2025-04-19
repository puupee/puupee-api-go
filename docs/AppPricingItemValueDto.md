# AppPricingItemValueDto

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
**IsAvailable** | Pointer to **bool** | 是否可用 | [optional] 
**HasValue** | Pointer to **bool** | 是否有值 | [optional] 
**IntValue** | Pointer to **int64** |  | [optional] 
**StringValue** | Pointer to **string** |  | [optional] 
**BoolValue** | Pointer to **bool** |  | [optional] 
**IntValueType** | Pointer to **string** | 数字值类型, FileSize: 文件大小, Count: 数目 | [optional] 

## Methods

### NewAppPricingItemValueDto

`func NewAppPricingItemValueDto() *AppPricingItemValueDto`

NewAppPricingItemValueDto instantiates a new AppPricingItemValueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPricingItemValueDtoWithDefaults

`func NewAppPricingItemValueDtoWithDefaults() *AppPricingItemValueDto`

NewAppPricingItemValueDtoWithDefaults instantiates a new AppPricingItemValueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppPricingItemValueDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppPricingItemValueDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppPricingItemValueDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppPricingItemValueDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *AppPricingItemValueDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AppPricingItemValueDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AppPricingItemValueDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AppPricingItemValueDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *AppPricingItemValueDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AppPricingItemValueDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AppPricingItemValueDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *AppPricingItemValueDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *AppPricingItemValueDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *AppPricingItemValueDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *AppPricingItemValueDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *AppPricingItemValueDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *AppPricingItemValueDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *AppPricingItemValueDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *AppPricingItemValueDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *AppPricingItemValueDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *AppPricingItemValueDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AppPricingItemValueDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AppPricingItemValueDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AppPricingItemValueDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *AppPricingItemValueDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *AppPricingItemValueDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *AppPricingItemValueDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *AppPricingItemValueDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *AppPricingItemValueDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *AppPricingItemValueDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *AppPricingItemValueDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *AppPricingItemValueDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetIsAvailable

`func (o *AppPricingItemValueDto) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *AppPricingItemValueDto) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *AppPricingItemValueDto) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.

### HasIsAvailable

`func (o *AppPricingItemValueDto) HasIsAvailable() bool`

HasIsAvailable returns a boolean if a field has been set.

### GetHasValue

`func (o *AppPricingItemValueDto) GetHasValue() bool`

GetHasValue returns the HasValue field if non-nil, zero value otherwise.

### GetHasValueOk

`func (o *AppPricingItemValueDto) GetHasValueOk() (*bool, bool)`

GetHasValueOk returns a tuple with the HasValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValue

`func (o *AppPricingItemValueDto) SetHasValue(v bool)`

SetHasValue sets HasValue field to given value.

### HasHasValue

`func (o *AppPricingItemValueDto) HasHasValue() bool`

HasHasValue returns a boolean if a field has been set.

### GetIntValue

`func (o *AppPricingItemValueDto) GetIntValue() int64`

GetIntValue returns the IntValue field if non-nil, zero value otherwise.

### GetIntValueOk

`func (o *AppPricingItemValueDto) GetIntValueOk() (*int64, bool)`

GetIntValueOk returns a tuple with the IntValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntValue

`func (o *AppPricingItemValueDto) SetIntValue(v int64)`

SetIntValue sets IntValue field to given value.

### HasIntValue

`func (o *AppPricingItemValueDto) HasIntValue() bool`

HasIntValue returns a boolean if a field has been set.

### GetStringValue

`func (o *AppPricingItemValueDto) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *AppPricingItemValueDto) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *AppPricingItemValueDto) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *AppPricingItemValueDto) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetBoolValue

`func (o *AppPricingItemValueDto) GetBoolValue() bool`

GetBoolValue returns the BoolValue field if non-nil, zero value otherwise.

### GetBoolValueOk

`func (o *AppPricingItemValueDto) GetBoolValueOk() (*bool, bool)`

GetBoolValueOk returns a tuple with the BoolValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValue

`func (o *AppPricingItemValueDto) SetBoolValue(v bool)`

SetBoolValue sets BoolValue field to given value.

### HasBoolValue

`func (o *AppPricingItemValueDto) HasBoolValue() bool`

HasBoolValue returns a boolean if a field has been set.

### GetIntValueType

`func (o *AppPricingItemValueDto) GetIntValueType() string`

GetIntValueType returns the IntValueType field if non-nil, zero value otherwise.

### GetIntValueTypeOk

`func (o *AppPricingItemValueDto) GetIntValueTypeOk() (*string, bool)`

GetIntValueTypeOk returns a tuple with the IntValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntValueType

`func (o *AppPricingItemValueDto) SetIntValueType(v string)`

SetIntValueType sets IntValueType field to given value.

### HasIntValueType

`func (o *AppPricingItemValueDto) HasIntValueType() bool`

HasIntValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



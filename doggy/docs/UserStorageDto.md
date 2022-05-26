# UserStorageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSize** | Pointer to **int64** |  | [optional] 
**CurrentSize** | Pointer to **int64** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]UserStorageItemDto**](UserStorageItemDto.md) |  | [optional] 

## Methods

### NewUserStorageDto

`func NewUserStorageDto() *UserStorageDto`

NewUserStorageDto instantiates a new UserStorageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStorageDtoWithDefaults

`func NewUserStorageDtoWithDefaults() *UserStorageDto`

NewUserStorageDtoWithDefaults instantiates a new UserStorageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSize

`func (o *UserStorageDto) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *UserStorageDto) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *UserStorageDto) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *UserStorageDto) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetCurrentSize

`func (o *UserStorageDto) GetCurrentSize() int64`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *UserStorageDto) GetCurrentSizeOk() (*int64, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *UserStorageDto) SetCurrentSize(v int64)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *UserStorageDto) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *UserStorageDto) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UserStorageDto) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UserStorageDto) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *UserStorageDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *UserStorageDto) GetItems() []UserStorageItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserStorageDto) GetItemsOk() (*[]UserStorageItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserStorageDto) SetItems(v []UserStorageItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *UserStorageDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *UserStorageDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *UserStorageDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



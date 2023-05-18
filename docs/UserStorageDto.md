# UserStorageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**CurrentSize** | Pointer to **int64** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**SingleFileMaxSize** | Pointer to **int64** |  | [optional] 
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

### GetName

`func (o *UserStorageDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserStorageDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserStorageDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserStorageDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserStorageDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserStorageDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *UserStorageDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserStorageDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserStorageDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserStorageDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UserStorageDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UserStorageDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSize

`func (o *UserStorageDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UserStorageDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UserStorageDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *UserStorageDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

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

### GetSingleFileMaxSize

`func (o *UserStorageDto) GetSingleFileMaxSize() int64`

GetSingleFileMaxSize returns the SingleFileMaxSize field if non-nil, zero value otherwise.

### GetSingleFileMaxSizeOk

`func (o *UserStorageDto) GetSingleFileMaxSizeOk() (*int64, bool)`

GetSingleFileMaxSizeOk returns a tuple with the SingleFileMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleFileMaxSize

`func (o *UserStorageDto) SetSingleFileMaxSize(v int64)`

SetSingleFileMaxSize sets SingleFileMaxSize field to given value.

### HasSingleFileMaxSize

`func (o *UserStorageDto) HasSingleFileMaxSize() bool`

HasSingleFileMaxSize returns a boolean if a field has been set.

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



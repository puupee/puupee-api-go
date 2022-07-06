# SpecialItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ItemDto**](ItemDto.md) |  | [optional] 
**Names** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSpecialItemDto

`func NewSpecialItemDto() *SpecialItemDto`

NewSpecialItemDto instantiates a new SpecialItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecialItemDtoWithDefaults

`func NewSpecialItemDtoWithDefaults() *SpecialItemDto`

NewSpecialItemDtoWithDefaults instantiates a new SpecialItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SpecialItemDto) GetItems() []ItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SpecialItemDto) GetItemsOk() (*[]ItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SpecialItemDto) SetItems(v []ItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *SpecialItemDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *SpecialItemDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *SpecialItemDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetNames

`func (o *SpecialItemDto) GetNames() string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *SpecialItemDto) GetNamesOk() (*string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *SpecialItemDto) SetNames(v string)`

SetNames sets Names field to given value.

### HasNames

`func (o *SpecialItemDto) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNamesNil

`func (o *SpecialItemDto) SetNamesNil(b bool)`

 SetNamesNil sets the value for Names to be an explicit nil

### UnsetNames
`func (o *SpecialItemDto) UnsetNames()`

UnsetNames ensures that no value is present for Names, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



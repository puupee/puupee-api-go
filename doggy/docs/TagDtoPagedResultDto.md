# TagDtoPagedResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]TagDto**](TagDto.md) |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewTagDtoPagedResultDto

`func NewTagDtoPagedResultDto() *TagDtoPagedResultDto`

NewTagDtoPagedResultDto instantiates a new TagDtoPagedResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagDtoPagedResultDtoWithDefaults

`func NewTagDtoPagedResultDtoWithDefaults() *TagDtoPagedResultDto`

NewTagDtoPagedResultDtoWithDefaults instantiates a new TagDtoPagedResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TagDtoPagedResultDto) GetItems() []TagDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TagDtoPagedResultDto) GetItemsOk() (*[]TagDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TagDtoPagedResultDto) SetItems(v []TagDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *TagDtoPagedResultDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *TagDtoPagedResultDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *TagDtoPagedResultDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetTotalCount

`func (o *TagDtoPagedResultDto) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TagDtoPagedResultDto) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TagDtoPagedResultDto) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TagDtoPagedResultDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



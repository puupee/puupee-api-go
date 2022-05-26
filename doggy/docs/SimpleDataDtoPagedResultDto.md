# SimpleDataDtoPagedResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SimpleDataDto**](SimpleDataDto.md) |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewSimpleDataDtoPagedResultDto

`func NewSimpleDataDtoPagedResultDto() *SimpleDataDtoPagedResultDto`

NewSimpleDataDtoPagedResultDto instantiates a new SimpleDataDtoPagedResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleDataDtoPagedResultDtoWithDefaults

`func NewSimpleDataDtoPagedResultDtoWithDefaults() *SimpleDataDtoPagedResultDto`

NewSimpleDataDtoPagedResultDtoWithDefaults instantiates a new SimpleDataDtoPagedResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SimpleDataDtoPagedResultDto) GetItems() []SimpleDataDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SimpleDataDtoPagedResultDto) GetItemsOk() (*[]SimpleDataDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SimpleDataDtoPagedResultDto) SetItems(v []SimpleDataDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *SimpleDataDtoPagedResultDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *SimpleDataDtoPagedResultDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *SimpleDataDtoPagedResultDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetTotalCount

`func (o *SimpleDataDtoPagedResultDto) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SimpleDataDtoPagedResultDto) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SimpleDataDtoPagedResultDto) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SimpleDataDtoPagedResultDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



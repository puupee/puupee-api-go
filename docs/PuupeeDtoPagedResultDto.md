# PuupeeDtoPagedResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]PuupeeDto**](PuupeeDto.md) |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewPuupeeDtoPagedResultDto

`func NewPuupeeDtoPagedResultDto() *PuupeeDtoPagedResultDto`

NewPuupeeDtoPagedResultDto instantiates a new PuupeeDtoPagedResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuupeeDtoPagedResultDtoWithDefaults

`func NewPuupeeDtoPagedResultDtoWithDefaults() *PuupeeDtoPagedResultDto`

NewPuupeeDtoPagedResultDtoWithDefaults instantiates a new PuupeeDtoPagedResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PuupeeDtoPagedResultDto) GetItems() []PuupeeDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PuupeeDtoPagedResultDto) GetItemsOk() (*[]PuupeeDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PuupeeDtoPagedResultDto) SetItems(v []PuupeeDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *PuupeeDtoPagedResultDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCount

`func (o *PuupeeDtoPagedResultDto) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PuupeeDtoPagedResultDto) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PuupeeDtoPagedResultDto) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PuupeeDtoPagedResultDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



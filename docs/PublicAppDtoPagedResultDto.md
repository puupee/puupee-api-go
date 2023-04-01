# PublicAppDtoPagedResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]PublicAppDto**](PublicAppDto.md) |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewPublicAppDtoPagedResultDto

`func NewPublicAppDtoPagedResultDto() *PublicAppDtoPagedResultDto`

NewPublicAppDtoPagedResultDto instantiates a new PublicAppDtoPagedResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAppDtoPagedResultDtoWithDefaults

`func NewPublicAppDtoPagedResultDtoWithDefaults() *PublicAppDtoPagedResultDto`

NewPublicAppDtoPagedResultDtoWithDefaults instantiates a new PublicAppDtoPagedResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PublicAppDtoPagedResultDto) GetItems() []PublicAppDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PublicAppDtoPagedResultDto) GetItemsOk() (*[]PublicAppDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PublicAppDtoPagedResultDto) SetItems(v []PublicAppDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *PublicAppDtoPagedResultDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCount

`func (o *PublicAppDtoPagedResultDto) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PublicAppDtoPagedResultDto) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PublicAppDtoPagedResultDto) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PublicAppDtoPagedResultDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



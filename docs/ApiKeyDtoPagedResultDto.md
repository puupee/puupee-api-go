# ApiKeyDtoPagedResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ApiKeyDto**](ApiKeyDto.md) |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewApiKeyDtoPagedResultDto

`func NewApiKeyDtoPagedResultDto() *ApiKeyDtoPagedResultDto`

NewApiKeyDtoPagedResultDto instantiates a new ApiKeyDtoPagedResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyDtoPagedResultDtoWithDefaults

`func NewApiKeyDtoPagedResultDtoWithDefaults() *ApiKeyDtoPagedResultDto`

NewApiKeyDtoPagedResultDtoWithDefaults instantiates a new ApiKeyDtoPagedResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ApiKeyDtoPagedResultDto) GetItems() []ApiKeyDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApiKeyDtoPagedResultDto) GetItemsOk() (*[]ApiKeyDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApiKeyDtoPagedResultDto) SetItems(v []ApiKeyDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApiKeyDtoPagedResultDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCount

`func (o *ApiKeyDtoPagedResultDto) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ApiKeyDtoPagedResultDto) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ApiKeyDtoPagedResultDto) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ApiKeyDtoPagedResultDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



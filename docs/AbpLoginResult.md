# AbpLoginResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**LoginResultType**](LoginResultType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAbpLoginResult

`func NewAbpLoginResult() *AbpLoginResult`

NewAbpLoginResult instantiates a new AbpLoginResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbpLoginResultWithDefaults

`func NewAbpLoginResultWithDefaults() *AbpLoginResult`

NewAbpLoginResultWithDefaults instantiates a new AbpLoginResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *AbpLoginResult) GetResult() LoginResultType`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AbpLoginResult) GetResultOk() (*LoginResultType, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AbpLoginResult) SetResult(v LoginResultType)`

SetResult sets Result field to given value.

### HasResult

`func (o *AbpLoginResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDescription

`func (o *AbpLoginResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AbpLoginResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AbpLoginResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AbpLoginResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



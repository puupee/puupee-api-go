# AppRunStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Value** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAppRunStatus

`func NewAppRunStatus() *AppRunStatus`

NewAppRunStatus instantiates a new AppRunStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunStatusWithDefaults

`func NewAppRunStatusWithDefaults() *AppRunStatus`

NewAppRunStatusWithDefaults instantiates a new AppRunStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppRunStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRunStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRunStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppRunStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *AppRunStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppRunStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppRunStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AppRunStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



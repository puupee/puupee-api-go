# PropertyApiDescriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**JsonName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TypeSimple** | Pointer to **string** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**Minimum** | Pointer to **string** |  | [optional] 
**Maximum** | Pointer to **string** |  | [optional] 
**Regex** | Pointer to **string** |  | [optional] 

## Methods

### NewPropertyApiDescriptionModel

`func NewPropertyApiDescriptionModel() *PropertyApiDescriptionModel`

NewPropertyApiDescriptionModel instantiates a new PropertyApiDescriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyApiDescriptionModelWithDefaults

`func NewPropertyApiDescriptionModelWithDefaults() *PropertyApiDescriptionModel`

NewPropertyApiDescriptionModelWithDefaults instantiates a new PropertyApiDescriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PropertyApiDescriptionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyApiDescriptionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyApiDescriptionModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PropertyApiDescriptionModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetJsonName

`func (o *PropertyApiDescriptionModel) GetJsonName() string`

GetJsonName returns the JsonName field if non-nil, zero value otherwise.

### GetJsonNameOk

`func (o *PropertyApiDescriptionModel) GetJsonNameOk() (*string, bool)`

GetJsonNameOk returns a tuple with the JsonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonName

`func (o *PropertyApiDescriptionModel) SetJsonName(v string)`

SetJsonName sets JsonName field to given value.

### HasJsonName

`func (o *PropertyApiDescriptionModel) HasJsonName() bool`

HasJsonName returns a boolean if a field has been set.

### GetType

`func (o *PropertyApiDescriptionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyApiDescriptionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyApiDescriptionModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PropertyApiDescriptionModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeSimple

`func (o *PropertyApiDescriptionModel) GetTypeSimple() string`

GetTypeSimple returns the TypeSimple field if non-nil, zero value otherwise.

### GetTypeSimpleOk

`func (o *PropertyApiDescriptionModel) GetTypeSimpleOk() (*string, bool)`

GetTypeSimpleOk returns a tuple with the TypeSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSimple

`func (o *PropertyApiDescriptionModel) SetTypeSimple(v string)`

SetTypeSimple sets TypeSimple field to given value.

### HasTypeSimple

`func (o *PropertyApiDescriptionModel) HasTypeSimple() bool`

HasTypeSimple returns a boolean if a field has been set.

### GetIsRequired

`func (o *PropertyApiDescriptionModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *PropertyApiDescriptionModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *PropertyApiDescriptionModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *PropertyApiDescriptionModel) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetMinLength

`func (o *PropertyApiDescriptionModel) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PropertyApiDescriptionModel) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PropertyApiDescriptionModel) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *PropertyApiDescriptionModel) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *PropertyApiDescriptionModel) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *PropertyApiDescriptionModel) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *PropertyApiDescriptionModel) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *PropertyApiDescriptionModel) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinimum

`func (o *PropertyApiDescriptionModel) GetMinimum() string`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *PropertyApiDescriptionModel) GetMinimumOk() (*string, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *PropertyApiDescriptionModel) SetMinimum(v string)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *PropertyApiDescriptionModel) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetMaximum

`func (o *PropertyApiDescriptionModel) GetMaximum() string`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *PropertyApiDescriptionModel) GetMaximumOk() (*string, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *PropertyApiDescriptionModel) SetMaximum(v string)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *PropertyApiDescriptionModel) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetRegex

`func (o *PropertyApiDescriptionModel) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *PropertyApiDescriptionModel) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *PropertyApiDescriptionModel) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *PropertyApiDescriptionModel) HasRegex() bool`

HasRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IValueValidator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Properties** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewIValueValidator

`func NewIValueValidator() *IValueValidator`

NewIValueValidator instantiates a new IValueValidator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIValueValidatorWithDefaults

`func NewIValueValidatorWithDefaults() *IValueValidator`

NewIValueValidatorWithDefaults instantiates a new IValueValidator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IValueValidator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IValueValidator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IValueValidator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IValueValidator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *IValueValidator) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IValueValidator) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IValueValidator) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IValueValidator) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



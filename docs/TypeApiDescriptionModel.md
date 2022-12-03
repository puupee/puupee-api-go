# TypeApiDescriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseType** | Pointer to **string** |  | [optional] 
**IsEnum** | Pointer to **bool** |  | [optional] 
**EnumNames** | Pointer to **[]string** |  | [optional] 
**EnumValues** | Pointer to **[]map[string]interface{}** |  | [optional] 
**GenericArguments** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to [**[]PropertyApiDescriptionModel**](PropertyApiDescriptionModel.md) |  | [optional] 

## Methods

### NewTypeApiDescriptionModel

`func NewTypeApiDescriptionModel() *TypeApiDescriptionModel`

NewTypeApiDescriptionModel instantiates a new TypeApiDescriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeApiDescriptionModelWithDefaults

`func NewTypeApiDescriptionModelWithDefaults() *TypeApiDescriptionModel`

NewTypeApiDescriptionModelWithDefaults instantiates a new TypeApiDescriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseType

`func (o *TypeApiDescriptionModel) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *TypeApiDescriptionModel) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *TypeApiDescriptionModel) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *TypeApiDescriptionModel) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetIsEnum

`func (o *TypeApiDescriptionModel) GetIsEnum() bool`

GetIsEnum returns the IsEnum field if non-nil, zero value otherwise.

### GetIsEnumOk

`func (o *TypeApiDescriptionModel) GetIsEnumOk() (*bool, bool)`

GetIsEnumOk returns a tuple with the IsEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnum

`func (o *TypeApiDescriptionModel) SetIsEnum(v bool)`

SetIsEnum sets IsEnum field to given value.

### HasIsEnum

`func (o *TypeApiDescriptionModel) HasIsEnum() bool`

HasIsEnum returns a boolean if a field has been set.

### GetEnumNames

`func (o *TypeApiDescriptionModel) GetEnumNames() []string`

GetEnumNames returns the EnumNames field if non-nil, zero value otherwise.

### GetEnumNamesOk

`func (o *TypeApiDescriptionModel) GetEnumNamesOk() (*[]string, bool)`

GetEnumNamesOk returns a tuple with the EnumNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumNames

`func (o *TypeApiDescriptionModel) SetEnumNames(v []string)`

SetEnumNames sets EnumNames field to given value.

### HasEnumNames

`func (o *TypeApiDescriptionModel) HasEnumNames() bool`

HasEnumNames returns a boolean if a field has been set.

### GetEnumValues

`func (o *TypeApiDescriptionModel) GetEnumValues() []map[string]interface{}`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *TypeApiDescriptionModel) GetEnumValuesOk() (*[]map[string]interface{}, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *TypeApiDescriptionModel) SetEnumValues(v []map[string]interface{})`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *TypeApiDescriptionModel) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### GetGenericArguments

`func (o *TypeApiDescriptionModel) GetGenericArguments() []string`

GetGenericArguments returns the GenericArguments field if non-nil, zero value otherwise.

### GetGenericArgumentsOk

`func (o *TypeApiDescriptionModel) GetGenericArgumentsOk() (*[]string, bool)`

GetGenericArgumentsOk returns a tuple with the GenericArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericArguments

`func (o *TypeApiDescriptionModel) SetGenericArguments(v []string)`

SetGenericArguments sets GenericArguments field to given value.

### HasGenericArguments

`func (o *TypeApiDescriptionModel) HasGenericArguments() bool`

HasGenericArguments returns a boolean if a field has been set.

### GetProperties

`func (o *TypeApiDescriptionModel) GetProperties() []PropertyApiDescriptionModel`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TypeApiDescriptionModel) GetPropertiesOk() (*[]PropertyApiDescriptionModel, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TypeApiDescriptionModel) SetProperties(v []PropertyApiDescriptionModel)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TypeApiDescriptionModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



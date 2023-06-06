# TenantDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ConcurrencyStamp** | Pointer to **string** |  | [optional] 

## Methods

### NewTenantDto

`func NewTenantDto() *TenantDto`

NewTenantDto instantiates a new TenantDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDtoWithDefaults

`func NewTenantDtoWithDefaults() *TenantDto`

NewTenantDtoWithDefaults instantiates a new TenantDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraProperties

`func (o *TenantDto) GetExtraProperties() map[string]map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *TenantDto) GetExtraPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *TenantDto) SetExtraProperties(v map[string]map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *TenantDto) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### GetId

`func (o *TenantDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TenantDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConcurrencyStamp

`func (o *TenantDto) GetConcurrencyStamp() string`

GetConcurrencyStamp returns the ConcurrencyStamp field if non-nil, zero value otherwise.

### GetConcurrencyStampOk

`func (o *TenantDto) GetConcurrencyStampOk() (*string, bool)`

GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyStamp

`func (o *TenantDto) SetConcurrencyStamp(v string)`

SetConcurrencyStamp sets ConcurrencyStamp field to given value.

### HasConcurrencyStamp

`func (o *TenantDto) HasConcurrencyStamp() bool`

HasConcurrencyStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FeatureDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**FeatureProviderDto**](FeatureProviderDto.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ValueType** | Pointer to [**IStringValueType**](IStringValueType.md) |  | [optional] 
**Depth** | Pointer to **int32** |  | [optional] 
**ParentName** | Pointer to **string** |  | [optional] 

## Methods

### NewFeatureDto

`func NewFeatureDto() *FeatureDto`

NewFeatureDto instantiates a new FeatureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureDtoWithDefaults

`func NewFeatureDtoWithDefaults() *FeatureDto`

NewFeatureDtoWithDefaults instantiates a new FeatureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeatureDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *FeatureDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FeatureDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FeatureDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FeatureDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetValue

`func (o *FeatureDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FeatureDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FeatureDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FeatureDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetProvider

`func (o *FeatureDto) GetProvider() FeatureProviderDto`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *FeatureDto) GetProviderOk() (*FeatureProviderDto, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *FeatureDto) SetProvider(v FeatureProviderDto)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *FeatureDto) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetDescription

`func (o *FeatureDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueType

`func (o *FeatureDto) GetValueType() IStringValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *FeatureDto) GetValueTypeOk() (*IStringValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *FeatureDto) SetValueType(v IStringValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *FeatureDto) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetDepth

`func (o *FeatureDto) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *FeatureDto) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *FeatureDto) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *FeatureDto) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetParentName

`func (o *FeatureDto) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *FeatureDto) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *FeatureDto) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *FeatureDto) HasParentName() bool`

HasParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



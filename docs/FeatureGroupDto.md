# FeatureGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]FeatureDto**](FeatureDto.md) |  | [optional] 

## Methods

### NewFeatureGroupDto

`func NewFeatureGroupDto() *FeatureGroupDto`

NewFeatureGroupDto instantiates a new FeatureGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureGroupDtoWithDefaults

`func NewFeatureGroupDtoWithDefaults() *FeatureGroupDto`

NewFeatureGroupDtoWithDefaults instantiates a new FeatureGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureGroupDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeatureGroupDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *FeatureGroupDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FeatureGroupDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FeatureGroupDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FeatureGroupDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFeatures

`func (o *FeatureGroupDto) GetFeatures() []FeatureDto`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FeatureGroupDto) GetFeaturesOk() (*[]FeatureDto, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FeatureGroupDto) SetFeatures(v []FeatureDto)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FeatureGroupDto) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



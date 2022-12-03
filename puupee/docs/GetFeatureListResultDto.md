# GetFeatureListResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]FeatureGroupDto**](FeatureGroupDto.md) |  | [optional] 

## Methods

### NewGetFeatureListResultDto

`func NewGetFeatureListResultDto() *GetFeatureListResultDto`

NewGetFeatureListResultDto instantiates a new GetFeatureListResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeatureListResultDtoWithDefaults

`func NewGetFeatureListResultDtoWithDefaults() *GetFeatureListResultDto`

NewGetFeatureListResultDtoWithDefaults instantiates a new GetFeatureListResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *GetFeatureListResultDto) GetGroups() []FeatureGroupDto`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetFeatureListResultDto) GetGroupsOk() (*[]FeatureGroupDto, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetFeatureListResultDto) SetGroups(v []FeatureGroupDto)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GetFeatureListResultDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



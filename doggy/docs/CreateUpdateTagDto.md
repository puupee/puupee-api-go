# CreateUpdateTagDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ParentTagId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateUpdateTagDto

`func NewCreateUpdateTagDto() *CreateUpdateTagDto`

NewCreateUpdateTagDto instantiates a new CreateUpdateTagDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateTagDtoWithDefaults

`func NewCreateUpdateTagDtoWithDefaults() *CreateUpdateTagDto`

NewCreateUpdateTagDtoWithDefaults instantiates a new CreateUpdateTagDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUpdateTagDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpdateTagDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpdateTagDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUpdateTagDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateUpdateTagDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateUpdateTagDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentTagId

`func (o *CreateUpdateTagDto) GetParentTagId() string`

GetParentTagId returns the ParentTagId field if non-nil, zero value otherwise.

### GetParentTagIdOk

`func (o *CreateUpdateTagDto) GetParentTagIdOk() (*string, bool)`

GetParentTagIdOk returns a tuple with the ParentTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTagId

`func (o *CreateUpdateTagDto) SetParentTagId(v string)`

SetParentTagId sets ParentTagId field to given value.

### HasParentTagId

`func (o *CreateUpdateTagDto) HasParentTagId() bool`

HasParentTagId returns a boolean if a field has been set.

### SetParentTagIdNil

`func (o *CreateUpdateTagDto) SetParentTagIdNil(b bool)`

 SetParentTagIdNil sets the value for ParentTagId to be an explicit nil

### UnsetParentTagId
`func (o *CreateUpdateTagDto) UnsetParentTagId()`

UnsetParentTagId ensures that no value is present for ParentTagId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



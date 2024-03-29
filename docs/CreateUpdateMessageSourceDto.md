# CreateUpdateMessageSourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsPublished** | Pointer to **bool** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**Routes** | Pointer to [**[]CreateUpdateMessageSourceRouteSubDto**](CreateUpdateMessageSourceRouteSubDto.md) |  | [optional] 

## Methods

### NewCreateUpdateMessageSourceDto

`func NewCreateUpdateMessageSourceDto() *CreateUpdateMessageSourceDto`

NewCreateUpdateMessageSourceDto instantiates a new CreateUpdateMessageSourceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateMessageSourceDtoWithDefaults

`func NewCreateUpdateMessageSourceDtoWithDefaults() *CreateUpdateMessageSourceDto`

NewCreateUpdateMessageSourceDtoWithDefaults instantiates a new CreateUpdateMessageSourceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUpdateMessageSourceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpdateMessageSourceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpdateMessageSourceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUpdateMessageSourceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateUpdateMessageSourceDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpdateMessageSourceDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpdateMessageSourceDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUpdateMessageSourceDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublished

`func (o *CreateUpdateMessageSourceDto) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *CreateUpdateMessageSourceDto) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *CreateUpdateMessageSourceDto) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *CreateUpdateMessageSourceDto) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetIconUrl

`func (o *CreateUpdateMessageSourceDto) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *CreateUpdateMessageSourceDto) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *CreateUpdateMessageSourceDto) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *CreateUpdateMessageSourceDto) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetRoutes

`func (o *CreateUpdateMessageSourceDto) GetRoutes() []CreateUpdateMessageSourceRouteSubDto`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *CreateUpdateMessageSourceDto) GetRoutesOk() (*[]CreateUpdateMessageSourceRouteSubDto, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *CreateUpdateMessageSourceDto) SetRoutes(v []CreateUpdateMessageSourceRouteSubDto)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *CreateUpdateMessageSourceDto) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateUpdateTodoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to [**Priority**](Priority.md) |  | [optional] 
**TagIds** | Pointer to **[]string** |  | [optional] 
**IsDone** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**EndAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCreateUpdateTodoDto

`func NewCreateUpdateTodoDto() *CreateUpdateTodoDto`

NewCreateUpdateTodoDto instantiates a new CreateUpdateTodoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateTodoDtoWithDefaults

`func NewCreateUpdateTodoDtoWithDefaults() *CreateUpdateTodoDto`

NewCreateUpdateTodoDtoWithDefaults instantiates a new CreateUpdateTodoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateUpdateTodoDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateUpdateTodoDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateUpdateTodoDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateUpdateTodoDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreateUpdateTodoDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateUpdateTodoDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPriority

`func (o *CreateUpdateTodoDto) GetPriority() Priority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateUpdateTodoDto) GetPriorityOk() (*Priority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateUpdateTodoDto) SetPriority(v Priority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateUpdateTodoDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTagIds

`func (o *CreateUpdateTodoDto) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *CreateUpdateTodoDto) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *CreateUpdateTodoDto) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *CreateUpdateTodoDto) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### SetTagIdsNil

`func (o *CreateUpdateTodoDto) SetTagIdsNil(b bool)`

 SetTagIdsNil sets the value for TagIds to be an explicit nil

### UnsetTagIds
`func (o *CreateUpdateTodoDto) UnsetTagIds()`

UnsetTagIds ensures that no value is present for TagIds, not even an explicit nil
### GetIsDone

`func (o *CreateUpdateTodoDto) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *CreateUpdateTodoDto) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *CreateUpdateTodoDto) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.

### HasIsDone

`func (o *CreateUpdateTodoDto) HasIsDone() bool`

HasIsDone returns a boolean if a field has been set.

### GetParentId

`func (o *CreateUpdateTodoDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateUpdateTodoDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateUpdateTodoDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateUpdateTodoDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *CreateUpdateTodoDto) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *CreateUpdateTodoDto) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetEndAt

`func (o *CreateUpdateTodoDto) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *CreateUpdateTodoDto) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *CreateUpdateTodoDto) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *CreateUpdateTodoDto) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *CreateUpdateTodoDto) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *CreateUpdateTodoDto) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



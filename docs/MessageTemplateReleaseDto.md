# MessageTemplateReleaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 
**LastModifierId** | Pointer to **string** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeleterId** | Pointer to **string** |  | [optional] 
**DeletionTime** | Pointer to **time.Time** |  | [optional] 
**TemplateName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageTemplateReleaseDto

`func NewMessageTemplateReleaseDto() *MessageTemplateReleaseDto`

NewMessageTemplateReleaseDto instantiates a new MessageTemplateReleaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageTemplateReleaseDtoWithDefaults

`func NewMessageTemplateReleaseDtoWithDefaults() *MessageTemplateReleaseDto`

NewMessageTemplateReleaseDtoWithDefaults instantiates a new MessageTemplateReleaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageTemplateReleaseDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageTemplateReleaseDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageTemplateReleaseDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageTemplateReleaseDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *MessageTemplateReleaseDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *MessageTemplateReleaseDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *MessageTemplateReleaseDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *MessageTemplateReleaseDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *MessageTemplateReleaseDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *MessageTemplateReleaseDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *MessageTemplateReleaseDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *MessageTemplateReleaseDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *MessageTemplateReleaseDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *MessageTemplateReleaseDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *MessageTemplateReleaseDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *MessageTemplateReleaseDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *MessageTemplateReleaseDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *MessageTemplateReleaseDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *MessageTemplateReleaseDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *MessageTemplateReleaseDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *MessageTemplateReleaseDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *MessageTemplateReleaseDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *MessageTemplateReleaseDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *MessageTemplateReleaseDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *MessageTemplateReleaseDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *MessageTemplateReleaseDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *MessageTemplateReleaseDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *MessageTemplateReleaseDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *MessageTemplateReleaseDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *MessageTemplateReleaseDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *MessageTemplateReleaseDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *MessageTemplateReleaseDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetTemplateName

`func (o *MessageTemplateReleaseDto) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *MessageTemplateReleaseDto) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *MessageTemplateReleaseDto) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *MessageTemplateReleaseDto) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetVersion

`func (o *MessageTemplateReleaseDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MessageTemplateReleaseDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MessageTemplateReleaseDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MessageTemplateReleaseDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetContent

`func (o *MessageTemplateReleaseDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageTemplateReleaseDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageTemplateReleaseDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MessageTemplateReleaseDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTemplateId

`func (o *MessageTemplateReleaseDto) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *MessageTemplateReleaseDto) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *MessageTemplateReleaseDto) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *MessageTemplateReleaseDto) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



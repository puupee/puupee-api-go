# NoteSpecDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**AuthorId** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 
**Extension** | Pointer to **NullableString** |  | [optional] 
**ContentFormating** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNoteSpecDto

`func NewNoteSpecDto() *NoteSpecDto`

NewNoteSpecDto instantiates a new NoteSpecDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteSpecDtoWithDefaults

`func NewNoteSpecDtoWithDefaults() *NoteSpecDto`

NewNoteSpecDtoWithDefaults instantiates a new NoteSpecDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoteSpecDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteSpecDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteSpecDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NoteSpecDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *NoteSpecDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NoteSpecDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NoteSpecDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NoteSpecDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *NoteSpecDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *NoteSpecDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *NoteSpecDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteSpecDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteSpecDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *NoteSpecDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *NoteSpecDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *NoteSpecDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetVersion

`func (o *NoteSpecDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NoteSpecDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NoteSpecDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NoteSpecDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAuthor

`func (o *NoteSpecDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NoteSpecDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NoteSpecDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *NoteSpecDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *NoteSpecDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *NoteSpecDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetAuthorId

`func (o *NoteSpecDto) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *NoteSpecDto) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *NoteSpecDto) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *NoteSpecDto) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreationTime

`func (o *NoteSpecDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *NoteSpecDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *NoteSpecDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *NoteSpecDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *NoteSpecDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *NoteSpecDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *NoteSpecDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *NoteSpecDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetExtension

`func (o *NoteSpecDto) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *NoteSpecDto) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *NoteSpecDto) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *NoteSpecDto) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### SetExtensionNil

`func (o *NoteSpecDto) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *NoteSpecDto) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil
### GetContentFormating

`func (o *NoteSpecDto) GetContentFormating() string`

GetContentFormating returns the ContentFormating field if non-nil, zero value otherwise.

### GetContentFormatingOk

`func (o *NoteSpecDto) GetContentFormatingOk() (*string, bool)`

GetContentFormatingOk returns a tuple with the ContentFormating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFormating

`func (o *NoteSpecDto) SetContentFormating(v string)`

SetContentFormating sets ContentFormating field to given value.

### HasContentFormating

`func (o *NoteSpecDto) HasContentFormating() bool`

HasContentFormating returns a boolean if a field has been set.

### SetContentFormatingNil

`func (o *NoteSpecDto) SetContentFormatingNil(b bool)`

 SetContentFormatingNil sets the value for ContentFormating to be an explicit nil

### UnsetContentFormating
`func (o *NoteSpecDto) UnsetContentFormating()`

UnsetContentFormating ensures that no value is present for ContentFormating, not even an explicit nil
### GetName

`func (o *NoteSpecDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NoteSpecDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NoteSpecDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NoteSpecDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NoteSpecDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NoteSpecDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *NoteSpecDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NoteSpecDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NoteSpecDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NoteSpecDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NoteSpecDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NoteSpecDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsite

`func (o *NoteSpecDto) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *NoteSpecDto) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *NoteSpecDto) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *NoteSpecDto) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *NoteSpecDto) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *NoteSpecDto) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



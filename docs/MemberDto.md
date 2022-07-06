# MemberDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**TodoExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**NoteExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**AlbumExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**FileExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**Level** | Pointer to **NullableString** |  | [optional] [readonly] 
**TodoLevel** | Pointer to **NullableString** |  | [optional] [readonly] 
**NoteLevel** | Pointer to **NullableString** |  | [optional] [readonly] 
**AlbumLevel** | Pointer to **NullableString** |  | [optional] [readonly] 
**FileLevel** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewMemberDto

`func NewMemberDto() *MemberDto`

NewMemberDto instantiates a new MemberDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDtoWithDefaults

`func NewMemberDtoWithDefaults() *MemberDto`

NewMemberDtoWithDefaults instantiates a new MemberDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiredAt

`func (o *MemberDto) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *MemberDto) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *MemberDto) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *MemberDto) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *MemberDto) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *MemberDto) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetTodoExpiredAt

`func (o *MemberDto) GetTodoExpiredAt() time.Time`

GetTodoExpiredAt returns the TodoExpiredAt field if non-nil, zero value otherwise.

### GetTodoExpiredAtOk

`func (o *MemberDto) GetTodoExpiredAtOk() (*time.Time, bool)`

GetTodoExpiredAtOk returns a tuple with the TodoExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoExpiredAt

`func (o *MemberDto) SetTodoExpiredAt(v time.Time)`

SetTodoExpiredAt sets TodoExpiredAt field to given value.

### HasTodoExpiredAt

`func (o *MemberDto) HasTodoExpiredAt() bool`

HasTodoExpiredAt returns a boolean if a field has been set.

### SetTodoExpiredAtNil

`func (o *MemberDto) SetTodoExpiredAtNil(b bool)`

 SetTodoExpiredAtNil sets the value for TodoExpiredAt to be an explicit nil

### UnsetTodoExpiredAt
`func (o *MemberDto) UnsetTodoExpiredAt()`

UnsetTodoExpiredAt ensures that no value is present for TodoExpiredAt, not even an explicit nil
### GetNoteExpiredAt

`func (o *MemberDto) GetNoteExpiredAt() time.Time`

GetNoteExpiredAt returns the NoteExpiredAt field if non-nil, zero value otherwise.

### GetNoteExpiredAtOk

`func (o *MemberDto) GetNoteExpiredAtOk() (*time.Time, bool)`

GetNoteExpiredAtOk returns a tuple with the NoteExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteExpiredAt

`func (o *MemberDto) SetNoteExpiredAt(v time.Time)`

SetNoteExpiredAt sets NoteExpiredAt field to given value.

### HasNoteExpiredAt

`func (o *MemberDto) HasNoteExpiredAt() bool`

HasNoteExpiredAt returns a boolean if a field has been set.

### SetNoteExpiredAtNil

`func (o *MemberDto) SetNoteExpiredAtNil(b bool)`

 SetNoteExpiredAtNil sets the value for NoteExpiredAt to be an explicit nil

### UnsetNoteExpiredAt
`func (o *MemberDto) UnsetNoteExpiredAt()`

UnsetNoteExpiredAt ensures that no value is present for NoteExpiredAt, not even an explicit nil
### GetAlbumExpiredAt

`func (o *MemberDto) GetAlbumExpiredAt() time.Time`

GetAlbumExpiredAt returns the AlbumExpiredAt field if non-nil, zero value otherwise.

### GetAlbumExpiredAtOk

`func (o *MemberDto) GetAlbumExpiredAtOk() (*time.Time, bool)`

GetAlbumExpiredAtOk returns a tuple with the AlbumExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumExpiredAt

`func (o *MemberDto) SetAlbumExpiredAt(v time.Time)`

SetAlbumExpiredAt sets AlbumExpiredAt field to given value.

### HasAlbumExpiredAt

`func (o *MemberDto) HasAlbumExpiredAt() bool`

HasAlbumExpiredAt returns a boolean if a field has been set.

### SetAlbumExpiredAtNil

`func (o *MemberDto) SetAlbumExpiredAtNil(b bool)`

 SetAlbumExpiredAtNil sets the value for AlbumExpiredAt to be an explicit nil

### UnsetAlbumExpiredAt
`func (o *MemberDto) UnsetAlbumExpiredAt()`

UnsetAlbumExpiredAt ensures that no value is present for AlbumExpiredAt, not even an explicit nil
### GetFileExpiredAt

`func (o *MemberDto) GetFileExpiredAt() time.Time`

GetFileExpiredAt returns the FileExpiredAt field if non-nil, zero value otherwise.

### GetFileExpiredAtOk

`func (o *MemberDto) GetFileExpiredAtOk() (*time.Time, bool)`

GetFileExpiredAtOk returns a tuple with the FileExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExpiredAt

`func (o *MemberDto) SetFileExpiredAt(v time.Time)`

SetFileExpiredAt sets FileExpiredAt field to given value.

### HasFileExpiredAt

`func (o *MemberDto) HasFileExpiredAt() bool`

HasFileExpiredAt returns a boolean if a field has been set.

### SetFileExpiredAtNil

`func (o *MemberDto) SetFileExpiredAtNil(b bool)`

 SetFileExpiredAtNil sets the value for FileExpiredAt to be an explicit nil

### UnsetFileExpiredAt
`func (o *MemberDto) UnsetFileExpiredAt()`

UnsetFileExpiredAt ensures that no value is present for FileExpiredAt, not even an explicit nil
### GetLevel

`func (o *MemberDto) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MemberDto) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MemberDto) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *MemberDto) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *MemberDto) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *MemberDto) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetTodoLevel

`func (o *MemberDto) GetTodoLevel() string`

GetTodoLevel returns the TodoLevel field if non-nil, zero value otherwise.

### GetTodoLevelOk

`func (o *MemberDto) GetTodoLevelOk() (*string, bool)`

GetTodoLevelOk returns a tuple with the TodoLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoLevel

`func (o *MemberDto) SetTodoLevel(v string)`

SetTodoLevel sets TodoLevel field to given value.

### HasTodoLevel

`func (o *MemberDto) HasTodoLevel() bool`

HasTodoLevel returns a boolean if a field has been set.

### SetTodoLevelNil

`func (o *MemberDto) SetTodoLevelNil(b bool)`

 SetTodoLevelNil sets the value for TodoLevel to be an explicit nil

### UnsetTodoLevel
`func (o *MemberDto) UnsetTodoLevel()`

UnsetTodoLevel ensures that no value is present for TodoLevel, not even an explicit nil
### GetNoteLevel

`func (o *MemberDto) GetNoteLevel() string`

GetNoteLevel returns the NoteLevel field if non-nil, zero value otherwise.

### GetNoteLevelOk

`func (o *MemberDto) GetNoteLevelOk() (*string, bool)`

GetNoteLevelOk returns a tuple with the NoteLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteLevel

`func (o *MemberDto) SetNoteLevel(v string)`

SetNoteLevel sets NoteLevel field to given value.

### HasNoteLevel

`func (o *MemberDto) HasNoteLevel() bool`

HasNoteLevel returns a boolean if a field has been set.

### SetNoteLevelNil

`func (o *MemberDto) SetNoteLevelNil(b bool)`

 SetNoteLevelNil sets the value for NoteLevel to be an explicit nil

### UnsetNoteLevel
`func (o *MemberDto) UnsetNoteLevel()`

UnsetNoteLevel ensures that no value is present for NoteLevel, not even an explicit nil
### GetAlbumLevel

`func (o *MemberDto) GetAlbumLevel() string`

GetAlbumLevel returns the AlbumLevel field if non-nil, zero value otherwise.

### GetAlbumLevelOk

`func (o *MemberDto) GetAlbumLevelOk() (*string, bool)`

GetAlbumLevelOk returns a tuple with the AlbumLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumLevel

`func (o *MemberDto) SetAlbumLevel(v string)`

SetAlbumLevel sets AlbumLevel field to given value.

### HasAlbumLevel

`func (o *MemberDto) HasAlbumLevel() bool`

HasAlbumLevel returns a boolean if a field has been set.

### SetAlbumLevelNil

`func (o *MemberDto) SetAlbumLevelNil(b bool)`

 SetAlbumLevelNil sets the value for AlbumLevel to be an explicit nil

### UnsetAlbumLevel
`func (o *MemberDto) UnsetAlbumLevel()`

UnsetAlbumLevel ensures that no value is present for AlbumLevel, not even an explicit nil
### GetFileLevel

`func (o *MemberDto) GetFileLevel() string`

GetFileLevel returns the FileLevel field if non-nil, zero value otherwise.

### GetFileLevelOk

`func (o *MemberDto) GetFileLevelOk() (*string, bool)`

GetFileLevelOk returns a tuple with the FileLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLevel

`func (o *MemberDto) SetFileLevel(v string)`

SetFileLevel sets FileLevel field to given value.

### HasFileLevel

`func (o *MemberDto) HasFileLevel() bool`

HasFileLevel returns a boolean if a field has been set.

### SetFileLevelNil

`func (o *MemberDto) SetFileLevelNil(b bool)`

 SetFileLevelNil sets the value for FileLevel to be an explicit nil

### UnsetFileLevel
`func (o *MemberDto) UnsetFileLevel()`

UnsetFileLevel ensures that no value is present for FileLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



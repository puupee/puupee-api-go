# AppRunDto

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
**AppId** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **map[string]interface{}** |  | [optional] 
**Envs** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAppRunDto

`func NewAppRunDto() *AppRunDto`

NewAppRunDto instantiates a new AppRunDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunDtoWithDefaults

`func NewAppRunDtoWithDefaults() *AppRunDto`

NewAppRunDtoWithDefaults instantiates a new AppRunDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppRunDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRunDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRunDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppRunDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *AppRunDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AppRunDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AppRunDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AppRunDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *AppRunDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AppRunDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AppRunDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *AppRunDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *AppRunDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *AppRunDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *AppRunDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *AppRunDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *AppRunDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *AppRunDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *AppRunDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *AppRunDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *AppRunDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AppRunDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AppRunDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AppRunDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *AppRunDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *AppRunDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *AppRunDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *AppRunDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *AppRunDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *AppRunDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *AppRunDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *AppRunDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetAppId

`func (o *AppRunDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppRunDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppRunDto) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppRunDto) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppName

`func (o *AppRunDto) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AppRunDto) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AppRunDto) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *AppRunDto) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetArgs

`func (o *AppRunDto) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *AppRunDto) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *AppRunDto) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *AppRunDto) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetEnvs

`func (o *AppRunDto) GetEnvs() map[string]interface{}`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *AppRunDto) GetEnvsOk() (*map[string]interface{}, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *AppRunDto) SetEnvs(v map[string]interface{})`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *AppRunDto) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



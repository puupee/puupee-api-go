# AppRunRecordDto

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
**Args** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Envs** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**AppRunStatus**](AppRunStatus.md) |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**FinishAt** | Pointer to **time.Time** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 

## Methods

### NewAppRunRecordDto

`func NewAppRunRecordDto() *AppRunRecordDto`

NewAppRunRecordDto instantiates a new AppRunRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunRecordDtoWithDefaults

`func NewAppRunRecordDtoWithDefaults() *AppRunRecordDto`

NewAppRunRecordDtoWithDefaults instantiates a new AppRunRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppRunRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRunRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRunRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppRunRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *AppRunRecordDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AppRunRecordDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AppRunRecordDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AppRunRecordDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *AppRunRecordDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AppRunRecordDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AppRunRecordDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *AppRunRecordDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *AppRunRecordDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *AppRunRecordDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *AppRunRecordDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *AppRunRecordDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *AppRunRecordDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *AppRunRecordDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *AppRunRecordDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *AppRunRecordDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *AppRunRecordDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AppRunRecordDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AppRunRecordDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AppRunRecordDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *AppRunRecordDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *AppRunRecordDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *AppRunRecordDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *AppRunRecordDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *AppRunRecordDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *AppRunRecordDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *AppRunRecordDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *AppRunRecordDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetAppId

`func (o *AppRunRecordDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppRunRecordDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppRunRecordDto) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppRunRecordDto) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppName

`func (o *AppRunRecordDto) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AppRunRecordDto) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AppRunRecordDto) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *AppRunRecordDto) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetArgs

`func (o *AppRunRecordDto) GetArgs() map[string]map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *AppRunRecordDto) GetArgsOk() (*map[string]map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *AppRunRecordDto) SetArgs(v map[string]map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *AppRunRecordDto) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetEnvs

`func (o *AppRunRecordDto) GetEnvs() map[string]map[string]interface{}`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *AppRunRecordDto) GetEnvsOk() (*map[string]map[string]interface{}, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *AppRunRecordDto) SetEnvs(v map[string]map[string]interface{})`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *AppRunRecordDto) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetStatus

`func (o *AppRunRecordDto) GetStatus() AppRunStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppRunRecordDto) GetStatusOk() (*AppRunStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppRunRecordDto) SetStatus(v AppRunStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppRunRecordDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *AppRunRecordDto) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AppRunRecordDto) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AppRunRecordDto) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AppRunRecordDto) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *AppRunRecordDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AppRunRecordDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AppRunRecordDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AppRunRecordDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinishAt

`func (o *AppRunRecordDto) GetFinishAt() time.Time`

GetFinishAt returns the FinishAt field if non-nil, zero value otherwise.

### GetFinishAtOk

`func (o *AppRunRecordDto) GetFinishAtOk() (*time.Time, bool)`

GetFinishAtOk returns a tuple with the FinishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishAt

`func (o *AppRunRecordDto) SetFinishAt(v time.Time)`

SetFinishAt sets FinishAt field to given value.

### HasFinishAt

`func (o *AppRunRecordDto) HasFinishAt() bool`

HasFinishAt returns a boolean if a field has been set.

### GetOutput

`func (o *AppRunRecordDto) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AppRunRecordDto) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AppRunRecordDto) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AppRunRecordDto) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



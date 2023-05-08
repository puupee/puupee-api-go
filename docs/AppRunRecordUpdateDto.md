# AppRunRecordUpdateDto

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
**Status** | [**AppRunStatus**](AppRunStatus.md) |  | 
**Result** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**FinishAt** | Pointer to **time.Time** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 
**WorkerId** | **string** |  | 
**WorkerName** | **string** |  | 

## Methods

### NewAppRunRecordUpdateDto

`func NewAppRunRecordUpdateDto(status AppRunStatus, workerId string, workerName string, ) *AppRunRecordUpdateDto`

NewAppRunRecordUpdateDto instantiates a new AppRunRecordUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunRecordUpdateDtoWithDefaults

`func NewAppRunRecordUpdateDtoWithDefaults() *AppRunRecordUpdateDto`

NewAppRunRecordUpdateDtoWithDefaults instantiates a new AppRunRecordUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppRunRecordUpdateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRunRecordUpdateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRunRecordUpdateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppRunRecordUpdateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *AppRunRecordUpdateDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AppRunRecordUpdateDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AppRunRecordUpdateDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AppRunRecordUpdateDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *AppRunRecordUpdateDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AppRunRecordUpdateDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AppRunRecordUpdateDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *AppRunRecordUpdateDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *AppRunRecordUpdateDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *AppRunRecordUpdateDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *AppRunRecordUpdateDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *AppRunRecordUpdateDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *AppRunRecordUpdateDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *AppRunRecordUpdateDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *AppRunRecordUpdateDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *AppRunRecordUpdateDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *AppRunRecordUpdateDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AppRunRecordUpdateDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AppRunRecordUpdateDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AppRunRecordUpdateDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *AppRunRecordUpdateDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *AppRunRecordUpdateDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *AppRunRecordUpdateDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *AppRunRecordUpdateDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *AppRunRecordUpdateDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *AppRunRecordUpdateDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *AppRunRecordUpdateDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *AppRunRecordUpdateDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetStatus

`func (o *AppRunRecordUpdateDto) GetStatus() AppRunStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppRunRecordUpdateDto) GetStatusOk() (*AppRunStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppRunRecordUpdateDto) SetStatus(v AppRunStatus)`

SetStatus sets Status field to given value.


### GetResult

`func (o *AppRunRecordUpdateDto) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AppRunRecordUpdateDto) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AppRunRecordUpdateDto) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AppRunRecordUpdateDto) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *AppRunRecordUpdateDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AppRunRecordUpdateDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AppRunRecordUpdateDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AppRunRecordUpdateDto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinishAt

`func (o *AppRunRecordUpdateDto) GetFinishAt() time.Time`

GetFinishAt returns the FinishAt field if non-nil, zero value otherwise.

### GetFinishAtOk

`func (o *AppRunRecordUpdateDto) GetFinishAtOk() (*time.Time, bool)`

GetFinishAtOk returns a tuple with the FinishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishAt

`func (o *AppRunRecordUpdateDto) SetFinishAt(v time.Time)`

SetFinishAt sets FinishAt field to given value.

### HasFinishAt

`func (o *AppRunRecordUpdateDto) HasFinishAt() bool`

HasFinishAt returns a boolean if a field has been set.

### GetOutput

`func (o *AppRunRecordUpdateDto) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AppRunRecordUpdateDto) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AppRunRecordUpdateDto) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AppRunRecordUpdateDto) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetWorkerId

`func (o *AppRunRecordUpdateDto) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *AppRunRecordUpdateDto) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *AppRunRecordUpdateDto) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetWorkerName

`func (o *AppRunRecordUpdateDto) GetWorkerName() string`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *AppRunRecordUpdateDto) GetWorkerNameOk() (*string, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *AppRunRecordUpdateDto) SetWorkerName(v string)`

SetWorkerName sets WorkerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



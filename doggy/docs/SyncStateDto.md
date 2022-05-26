# SyncStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncAt** | Pointer to **time.Time** |  | [optional] 
**TagVersion** | Pointer to **int64** |  | [optional] 
**ItemVersion** | Pointer to **int64** |  | [optional] 
**TodoVersion** | Pointer to **int64** |  | [optional] 

## Methods

### NewSyncStateDto

`func NewSyncStateDto() *SyncStateDto`

NewSyncStateDto instantiates a new SyncStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncStateDtoWithDefaults

`func NewSyncStateDtoWithDefaults() *SyncStateDto`

NewSyncStateDtoWithDefaults instantiates a new SyncStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSyncAt

`func (o *SyncStateDto) GetLastSyncAt() time.Time`

GetLastSyncAt returns the LastSyncAt field if non-nil, zero value otherwise.

### GetLastSyncAtOk

`func (o *SyncStateDto) GetLastSyncAtOk() (*time.Time, bool)`

GetLastSyncAtOk returns a tuple with the LastSyncAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncAt

`func (o *SyncStateDto) SetLastSyncAt(v time.Time)`

SetLastSyncAt sets LastSyncAt field to given value.

### HasLastSyncAt

`func (o *SyncStateDto) HasLastSyncAt() bool`

HasLastSyncAt returns a boolean if a field has been set.

### GetTagVersion

`func (o *SyncStateDto) GetTagVersion() int64`

GetTagVersion returns the TagVersion field if non-nil, zero value otherwise.

### GetTagVersionOk

`func (o *SyncStateDto) GetTagVersionOk() (*int64, bool)`

GetTagVersionOk returns a tuple with the TagVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagVersion

`func (o *SyncStateDto) SetTagVersion(v int64)`

SetTagVersion sets TagVersion field to given value.

### HasTagVersion

`func (o *SyncStateDto) HasTagVersion() bool`

HasTagVersion returns a boolean if a field has been set.

### GetItemVersion

`func (o *SyncStateDto) GetItemVersion() int64`

GetItemVersion returns the ItemVersion field if non-nil, zero value otherwise.

### GetItemVersionOk

`func (o *SyncStateDto) GetItemVersionOk() (*int64, bool)`

GetItemVersionOk returns a tuple with the ItemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersion

`func (o *SyncStateDto) SetItemVersion(v int64)`

SetItemVersion sets ItemVersion field to given value.

### HasItemVersion

`func (o *SyncStateDto) HasItemVersion() bool`

HasItemVersion returns a boolean if a field has been set.

### GetTodoVersion

`func (o *SyncStateDto) GetTodoVersion() int64`

GetTodoVersion returns the TodoVersion field if non-nil, zero value otherwise.

### GetTodoVersionOk

`func (o *SyncStateDto) GetTodoVersionOk() (*int64, bool)`

GetTodoVersionOk returns a tuple with the TodoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoVersion

`func (o *SyncStateDto) SetTodoVersion(v int64)`

SetTodoVersion sets TodoVersion field to given value.

### HasTodoVersion

`func (o *SyncStateDto) HasTodoVersion() bool`

HasTodoVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



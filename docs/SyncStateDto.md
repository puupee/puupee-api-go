# SyncStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncAt** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

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

### GetVersion

`func (o *SyncStateDto) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SyncStateDto) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SyncStateDto) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SyncStateDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



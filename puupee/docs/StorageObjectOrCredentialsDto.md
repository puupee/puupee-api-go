# StorageObjectOrCredentialsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageObject** | Pointer to [**StorageObjectDto**](StorageObjectDto.md) |  | [optional] 
**Credentials** | Pointer to [**StorageObjectCredentials**](StorageObjectCredentials.md) |  | [optional] 

## Methods

### NewStorageObjectOrCredentialsDto

`func NewStorageObjectOrCredentialsDto() *StorageObjectOrCredentialsDto`

NewStorageObjectOrCredentialsDto instantiates a new StorageObjectOrCredentialsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageObjectOrCredentialsDtoWithDefaults

`func NewStorageObjectOrCredentialsDtoWithDefaults() *StorageObjectOrCredentialsDto`

NewStorageObjectOrCredentialsDtoWithDefaults instantiates a new StorageObjectOrCredentialsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageObject

`func (o *StorageObjectOrCredentialsDto) GetStorageObject() StorageObjectDto`

GetStorageObject returns the StorageObject field if non-nil, zero value otherwise.

### GetStorageObjectOk

`func (o *StorageObjectOrCredentialsDto) GetStorageObjectOk() (*StorageObjectDto, bool)`

GetStorageObjectOk returns a tuple with the StorageObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageObject

`func (o *StorageObjectOrCredentialsDto) SetStorageObject(v StorageObjectDto)`

SetStorageObject sets StorageObject field to given value.

### HasStorageObject

`func (o *StorageObjectOrCredentialsDto) HasStorageObject() bool`

HasStorageObject returns a boolean if a field has been set.

### GetCredentials

`func (o *StorageObjectOrCredentialsDto) GetCredentials() StorageObjectCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *StorageObjectOrCredentialsDto) GetCredentialsOk() (*StorageObjectCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *StorageObjectOrCredentialsDto) SetCredentials(v StorageObjectCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *StorageObjectOrCredentialsDto) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



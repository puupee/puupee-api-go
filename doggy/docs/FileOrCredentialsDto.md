# FileOrCredentialsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**FileDto**](FileDto.md) |  | [optional] 
**Credentials** | Pointer to [**UploadCredentials**](UploadCredentials.md) |  | [optional] 

## Methods

### NewFileOrCredentialsDto

`func NewFileOrCredentialsDto() *FileOrCredentialsDto`

NewFileOrCredentialsDto instantiates a new FileOrCredentialsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileOrCredentialsDtoWithDefaults

`func NewFileOrCredentialsDtoWithDefaults() *FileOrCredentialsDto`

NewFileOrCredentialsDtoWithDefaults instantiates a new FileOrCredentialsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *FileOrCredentialsDto) GetFile() FileDto`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FileOrCredentialsDto) GetFileOk() (*FileDto, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FileOrCredentialsDto) SetFile(v FileDto)`

SetFile sets File field to given value.

### HasFile

`func (o *FileOrCredentialsDto) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetCredentials

`func (o *FileOrCredentialsDto) GetCredentials() UploadCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *FileOrCredentialsDto) GetCredentialsOk() (*UploadCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *FileOrCredentialsDto) SetCredentials(v UploadCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *FileOrCredentialsDto) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



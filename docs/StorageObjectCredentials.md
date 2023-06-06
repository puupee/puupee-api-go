# StorageObjectCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageClass** | Pointer to **string** |  | [optional] 
**EndPoint** | Pointer to **string** |  | [optional] 
**Protocal** | Pointer to **string** |  | [optional] 
**BucketName** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**SecurityToken** | Pointer to **string** |  | [optional] 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**AccessKeySecret** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **string** |  | [optional] 
**ExpiredTime** | Pointer to **int64** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 

## Methods

### NewStorageObjectCredentials

`func NewStorageObjectCredentials() *StorageObjectCredentials`

NewStorageObjectCredentials instantiates a new StorageObjectCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageObjectCredentialsWithDefaults

`func NewStorageObjectCredentialsWithDefaults() *StorageObjectCredentials`

NewStorageObjectCredentialsWithDefaults instantiates a new StorageObjectCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageClass

`func (o *StorageObjectCredentials) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *StorageObjectCredentials) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *StorageObjectCredentials) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *StorageObjectCredentials) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetEndPoint

`func (o *StorageObjectCredentials) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *StorageObjectCredentials) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *StorageObjectCredentials) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.

### HasEndPoint

`func (o *StorageObjectCredentials) HasEndPoint() bool`

HasEndPoint returns a boolean if a field has been set.

### GetProtocal

`func (o *StorageObjectCredentials) GetProtocal() string`

GetProtocal returns the Protocal field if non-nil, zero value otherwise.

### GetProtocalOk

`func (o *StorageObjectCredentials) GetProtocalOk() (*string, bool)`

GetProtocalOk returns a tuple with the Protocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocal

`func (o *StorageObjectCredentials) SetProtocal(v string)`

SetProtocal sets Protocal field to given value.

### HasProtocal

`func (o *StorageObjectCredentials) HasProtocal() bool`

HasProtocal returns a boolean if a field has been set.

### GetBucketName

`func (o *StorageObjectCredentials) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *StorageObjectCredentials) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *StorageObjectCredentials) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *StorageObjectCredentials) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetRegionId

`func (o *StorageObjectCredentials) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *StorageObjectCredentials) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *StorageObjectCredentials) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *StorageObjectCredentials) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetSecurityToken

`func (o *StorageObjectCredentials) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *StorageObjectCredentials) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *StorageObjectCredentials) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *StorageObjectCredentials) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *StorageObjectCredentials) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *StorageObjectCredentials) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *StorageObjectCredentials) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *StorageObjectCredentials) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAccessKeySecret

`func (o *StorageObjectCredentials) GetAccessKeySecret() string`

GetAccessKeySecret returns the AccessKeySecret field if non-nil, zero value otherwise.

### GetAccessKeySecretOk

`func (o *StorageObjectCredentials) GetAccessKeySecretOk() (*string, bool)`

GetAccessKeySecretOk returns a tuple with the AccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecret

`func (o *StorageObjectCredentials) SetAccessKeySecret(v string)`

SetAccessKeySecret sets AccessKeySecret field to given value.

### HasAccessKeySecret

`func (o *StorageObjectCredentials) HasAccessKeySecret() bool`

HasAccessKeySecret returns a boolean if a field has been set.

### GetExpiration

`func (o *StorageObjectCredentials) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *StorageObjectCredentials) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *StorageObjectCredentials) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *StorageObjectCredentials) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetExpiredTime

`func (o *StorageObjectCredentials) GetExpiredTime() int64`

GetExpiredTime returns the ExpiredTime field if non-nil, zero value otherwise.

### GetExpiredTimeOk

`func (o *StorageObjectCredentials) GetExpiredTimeOk() (*int64, bool)`

GetExpiredTimeOk returns a tuple with the ExpiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTime

`func (o *StorageObjectCredentials) SetExpiredTime(v int64)`

SetExpiredTime sets ExpiredTime field to given value.

### HasExpiredTime

`func (o *StorageObjectCredentials) HasExpiredTime() bool`

HasExpiredTime returns a boolean if a field has been set.

### GetAppId

`func (o *StorageObjectCredentials) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *StorageObjectCredentials) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *StorageObjectCredentials) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *StorageObjectCredentials) HasAppId() bool`

HasAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



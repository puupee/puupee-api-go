/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// checks if the StorageObjectCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageObjectCredentials{}

// StorageObjectCredentials struct for StorageObjectCredentials
type StorageObjectCredentials struct {
	StorageClass *string `json:"storageClass,omitempty"`
	EndPoint *string `json:"endPoint,omitempty"`
	Protocal *string `json:"protocal,omitempty"`
	BucketName *string `json:"bucketName,omitempty"`
	RegionId *string `json:"regionId,omitempty"`
	SecurityToken *string `json:"securityToken,omitempty"`
	AccessKeyId *string `json:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty"`
	Expiration *string `json:"expiration,omitempty"`
	ExpiredTime *int64 `json:"expiredTime,omitempty"`
	AppId *string `json:"appId,omitempty"`
}

// NewStorageObjectCredentials instantiates a new StorageObjectCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageObjectCredentials() *StorageObjectCredentials {
	this := StorageObjectCredentials{}
	return &this
}

// NewStorageObjectCredentialsWithDefaults instantiates a new StorageObjectCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageObjectCredentialsWithDefaults() *StorageObjectCredentials {
	this := StorageObjectCredentials{}
	return &this
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetStorageClass() string {
	if o == nil || IsNil(o.StorageClass) {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetStorageClassOk() (*string, bool) {
	if o == nil || IsNil(o.StorageClass) {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasStorageClass() bool {
	if o != nil && !IsNil(o.StorageClass) {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *StorageObjectCredentials) SetStorageClass(v string) {
	o.StorageClass = &v
}

// GetEndPoint returns the EndPoint field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetEndPoint() string {
	if o == nil || IsNil(o.EndPoint) {
		var ret string
		return ret
	}
	return *o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetEndPointOk() (*string, bool) {
	if o == nil || IsNil(o.EndPoint) {
		return nil, false
	}
	return o.EndPoint, true
}

// HasEndPoint returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasEndPoint() bool {
	if o != nil && !IsNil(o.EndPoint) {
		return true
	}

	return false
}

// SetEndPoint gets a reference to the given string and assigns it to the EndPoint field.
func (o *StorageObjectCredentials) SetEndPoint(v string) {
	o.EndPoint = &v
}

// GetProtocal returns the Protocal field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetProtocal() string {
	if o == nil || IsNil(o.Protocal) {
		var ret string
		return ret
	}
	return *o.Protocal
}

// GetProtocalOk returns a tuple with the Protocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetProtocalOk() (*string, bool) {
	if o == nil || IsNil(o.Protocal) {
		return nil, false
	}
	return o.Protocal, true
}

// HasProtocal returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasProtocal() bool {
	if o != nil && !IsNil(o.Protocal) {
		return true
	}

	return false
}

// SetProtocal gets a reference to the given string and assigns it to the Protocal field.
func (o *StorageObjectCredentials) SetProtocal(v string) {
	o.Protocal = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *StorageObjectCredentials) SetBucketName(v string) {
	o.BucketName = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetRegionId() string {
	if o == nil || IsNil(o.RegionId) {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetRegionIdOk() (*string, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *StorageObjectCredentials) SetRegionId(v string) {
	o.RegionId = &v
}

// GetSecurityToken returns the SecurityToken field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetSecurityToken() string {
	if o == nil || IsNil(o.SecurityToken) {
		var ret string
		return ret
	}
	return *o.SecurityToken
}

// GetSecurityTokenOk returns a tuple with the SecurityToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetSecurityTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityToken) {
		return nil, false
	}
	return o.SecurityToken, true
}

// HasSecurityToken returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasSecurityToken() bool {
	if o != nil && !IsNil(o.SecurityToken) {
		return true
	}

	return false
}

// SetSecurityToken gets a reference to the given string and assigns it to the SecurityToken field.
func (o *StorageObjectCredentials) SetSecurityToken(v string) {
	o.SecurityToken = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetAccessKeyId() string {
	if o == nil || IsNil(o.AccessKeyId) {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKeyId) {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasAccessKeyId() bool {
	if o != nil && !IsNil(o.AccessKeyId) {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *StorageObjectCredentials) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetAccessKeySecret returns the AccessKeySecret field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetAccessKeySecret() string {
	if o == nil || IsNil(o.AccessKeySecret) {
		var ret string
		return ret
	}
	return *o.AccessKeySecret
}

// GetAccessKeySecretOk returns a tuple with the AccessKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetAccessKeySecretOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKeySecret) {
		return nil, false
	}
	return o.AccessKeySecret, true
}

// HasAccessKeySecret returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasAccessKeySecret() bool {
	if o != nil && !IsNil(o.AccessKeySecret) {
		return true
	}

	return false
}

// SetAccessKeySecret gets a reference to the given string and assigns it to the AccessKeySecret field.
func (o *StorageObjectCredentials) SetAccessKeySecret(v string) {
	o.AccessKeySecret = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetExpiration() string {
	if o == nil || IsNil(o.Expiration) {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *StorageObjectCredentials) SetExpiration(v string) {
	o.Expiration = &v
}

// GetExpiredTime returns the ExpiredTime field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetExpiredTime() int64 {
	if o == nil || IsNil(o.ExpiredTime) {
		var ret int64
		return ret
	}
	return *o.ExpiredTime
}

// GetExpiredTimeOk returns a tuple with the ExpiredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetExpiredTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiredTime) {
		return nil, false
	}
	return o.ExpiredTime, true
}

// HasExpiredTime returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasExpiredTime() bool {
	if o != nil && !IsNil(o.ExpiredTime) {
		return true
	}

	return false
}

// SetExpiredTime gets a reference to the given int64 and assigns it to the ExpiredTime field.
func (o *StorageObjectCredentials) SetExpiredTime(v int64) {
	o.ExpiredTime = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *StorageObjectCredentials) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageObjectCredentials) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *StorageObjectCredentials) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *StorageObjectCredentials) SetAppId(v string) {
	o.AppId = &v
}

func (o StorageObjectCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageObjectCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageClass) {
		toSerialize["storageClass"] = o.StorageClass
	}
	if !IsNil(o.EndPoint) {
		toSerialize["endPoint"] = o.EndPoint
	}
	if !IsNil(o.Protocal) {
		toSerialize["protocal"] = o.Protocal
	}
	if !IsNil(o.BucketName) {
		toSerialize["bucketName"] = o.BucketName
	}
	if !IsNil(o.RegionId) {
		toSerialize["regionId"] = o.RegionId
	}
	if !IsNil(o.SecurityToken) {
		toSerialize["securityToken"] = o.SecurityToken
	}
	if !IsNil(o.AccessKeyId) {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if !IsNil(o.AccessKeySecret) {
		toSerialize["accessKeySecret"] = o.AccessKeySecret
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.ExpiredTime) {
		toSerialize["expiredTime"] = o.ExpiredTime
	}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	return toSerialize, nil
}

type NullableStorageObjectCredentials struct {
	value *StorageObjectCredentials
	isSet bool
}

func (v NullableStorageObjectCredentials) Get() *StorageObjectCredentials {
	return v.value
}

func (v *NullableStorageObjectCredentials) Set(val *StorageObjectCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageObjectCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageObjectCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageObjectCredentials(val *StorageObjectCredentials) *NullableStorageObjectCredentials {
	return &NullableStorageObjectCredentials{value: val, isSet: true}
}

func (v NullableStorageObjectCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageObjectCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



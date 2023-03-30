/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// SyncStateDto struct for SyncStateDto
type SyncStateDto struct {
	LastSyncAt *time.Time `json:"lastSyncAt,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewSyncStateDto instantiates a new SyncStateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncStateDto() *SyncStateDto {
	this := SyncStateDto{}
	return &this
}

// NewSyncStateDtoWithDefaults instantiates a new SyncStateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncStateDtoWithDefaults() *SyncStateDto {
	this := SyncStateDto{}
	return &this
}

// GetLastSyncAt returns the LastSyncAt field value if set, zero value otherwise.
func (o *SyncStateDto) GetLastSyncAt() time.Time {
	if o == nil || isNil(o.LastSyncAt) {
		var ret time.Time
		return ret
	}
	return *o.LastSyncAt
}

// GetLastSyncAtOk returns a tuple with the LastSyncAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncStateDto) GetLastSyncAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastSyncAt) {
    return nil, false
	}
	return o.LastSyncAt, true
}

// HasLastSyncAt returns a boolean if a field has been set.
func (o *SyncStateDto) HasLastSyncAt() bool {
	if o != nil && !isNil(o.LastSyncAt) {
		return true
	}

	return false
}

// SetLastSyncAt gets a reference to the given time.Time and assigns it to the LastSyncAt field.
func (o *SyncStateDto) SetLastSyncAt(v time.Time) {
	o.LastSyncAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SyncStateDto) GetVersion() int64 {
	if o == nil || isNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncStateDto) GetVersionOk() (*int64, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SyncStateDto) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SyncStateDto) SetVersion(v int64) {
	o.Version = &v
}

func (o SyncStateDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LastSyncAt) {
		toSerialize["lastSyncAt"] = o.LastSyncAt
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSyncStateDto struct {
	value *SyncStateDto
	isSet bool
}

func (v NullableSyncStateDto) Get() *SyncStateDto {
	return v.value
}

func (v *NullableSyncStateDto) Set(val *SyncStateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncStateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncStateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncStateDto(val *SyncStateDto) *NullableSyncStateDto {
	return &NullableSyncStateDto{value: val, isSet: true}
}

func (v NullableSyncStateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncStateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



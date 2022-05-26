/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
)

// RefreshDeviceStatusDto struct for RefreshDeviceStatusDto
type RefreshDeviceStatusDto struct {
	Id NullableString `json:"id,omitempty"`
	Status *DeviceStatus `json:"status,omitempty"`
}

// NewRefreshDeviceStatusDto instantiates a new RefreshDeviceStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshDeviceStatusDto() *RefreshDeviceStatusDto {
	this := RefreshDeviceStatusDto{}
	return &this
}

// NewRefreshDeviceStatusDtoWithDefaults instantiates a new RefreshDeviceStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshDeviceStatusDtoWithDefaults() *RefreshDeviceStatusDto {
	this := RefreshDeviceStatusDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshDeviceStatusDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshDeviceStatusDto) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RefreshDeviceStatusDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *RefreshDeviceStatusDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RefreshDeviceStatusDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RefreshDeviceStatusDto) UnsetId() {
	o.Id.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RefreshDeviceStatusDto) GetStatus() DeviceStatus {
	if o == nil || o.Status == nil {
		var ret DeviceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshDeviceStatusDto) GetStatusOk() (*DeviceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RefreshDeviceStatusDto) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeviceStatus and assigns it to the Status field.
func (o *RefreshDeviceStatusDto) SetStatus(v DeviceStatus) {
	o.Status = &v
}

func (o RefreshDeviceStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshDeviceStatusDto struct {
	value *RefreshDeviceStatusDto
	isSet bool
}

func (v NullableRefreshDeviceStatusDto) Get() *RefreshDeviceStatusDto {
	return v.value
}

func (v *NullableRefreshDeviceStatusDto) Set(val *RefreshDeviceStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshDeviceStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshDeviceStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshDeviceStatusDto(val *RefreshDeviceStatusDto) *NullableRefreshDeviceStatusDto {
	return &NullableRefreshDeviceStatusDto{value: val, isSet: true}
}

func (v NullableRefreshDeviceStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshDeviceStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



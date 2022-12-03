/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// DeviceDtoPagedResultDto struct for DeviceDtoPagedResultDto
type DeviceDtoPagedResultDto struct {
	Items []DeviceDto `json:"items,omitempty"`
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// NewDeviceDtoPagedResultDto instantiates a new DeviceDtoPagedResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceDtoPagedResultDto() *DeviceDtoPagedResultDto {
	this := DeviceDtoPagedResultDto{}
	return &this
}

// NewDeviceDtoPagedResultDtoWithDefaults instantiates a new DeviceDtoPagedResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceDtoPagedResultDtoWithDefaults() *DeviceDtoPagedResultDto {
	this := DeviceDtoPagedResultDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DeviceDtoPagedResultDto) GetItems() []DeviceDto {
	if o == nil || isNil(o.Items) {
		var ret []DeviceDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDtoPagedResultDto) GetItemsOk() ([]DeviceDto, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DeviceDtoPagedResultDto) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DeviceDto and assigns it to the Items field.
func (o *DeviceDtoPagedResultDto) SetItems(v []DeviceDto) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *DeviceDtoPagedResultDto) GetTotalCount() int64 {
	if o == nil || isNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDtoPagedResultDto) GetTotalCountOk() (*int64, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *DeviceDtoPagedResultDto) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *DeviceDtoPagedResultDto) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o DeviceDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceDtoPagedResultDto struct {
	value *DeviceDtoPagedResultDto
	isSet bool
}

func (v NullableDeviceDtoPagedResultDto) Get() *DeviceDtoPagedResultDto {
	return v.value
}

func (v *NullableDeviceDtoPagedResultDto) Set(val *DeviceDtoPagedResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceDtoPagedResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceDtoPagedResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceDtoPagedResultDto(val *DeviceDtoPagedResultDto) *NullableDeviceDtoPagedResultDto {
	return &NullableDeviceDtoPagedResultDto{value: val, isSet: true}
}

func (v NullableDeviceDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceDtoPagedResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



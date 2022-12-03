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

// AppReleaseDtoPagedResultDto struct for AppReleaseDtoPagedResultDto
type AppReleaseDtoPagedResultDto struct {
	Items []AppReleaseDto `json:"items,omitempty"`
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// NewAppReleaseDtoPagedResultDto instantiates a new AppReleaseDtoPagedResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppReleaseDtoPagedResultDto() *AppReleaseDtoPagedResultDto {
	this := AppReleaseDtoPagedResultDto{}
	return &this
}

// NewAppReleaseDtoPagedResultDtoWithDefaults instantiates a new AppReleaseDtoPagedResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppReleaseDtoPagedResultDtoWithDefaults() *AppReleaseDtoPagedResultDto {
	this := AppReleaseDtoPagedResultDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AppReleaseDtoPagedResultDto) GetItems() []AppReleaseDto {
	if o == nil || isNil(o.Items) {
		var ret []AppReleaseDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDtoPagedResultDto) GetItemsOk() ([]AppReleaseDto, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AppReleaseDtoPagedResultDto) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AppReleaseDto and assigns it to the Items field.
func (o *AppReleaseDtoPagedResultDto) SetItems(v []AppReleaseDto) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AppReleaseDtoPagedResultDto) GetTotalCount() int64 {
	if o == nil || isNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDtoPagedResultDto) GetTotalCountOk() (*int64, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AppReleaseDtoPagedResultDto) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *AppReleaseDtoPagedResultDto) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o AppReleaseDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableAppReleaseDtoPagedResultDto struct {
	value *AppReleaseDtoPagedResultDto
	isSet bool
}

func (v NullableAppReleaseDtoPagedResultDto) Get() *AppReleaseDtoPagedResultDto {
	return v.value
}

func (v *NullableAppReleaseDtoPagedResultDto) Set(val *AppReleaseDtoPagedResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppReleaseDtoPagedResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppReleaseDtoPagedResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppReleaseDtoPagedResultDto(val *AppReleaseDtoPagedResultDto) *NullableAppReleaseDtoPagedResultDto {
	return &NullableAppReleaseDtoPagedResultDto{value: val, isSet: true}
}

func (v NullableAppReleaseDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppReleaseDtoPagedResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



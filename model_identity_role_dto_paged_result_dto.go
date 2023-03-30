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

// IdentityRoleDtoPagedResultDto struct for IdentityRoleDtoPagedResultDto
type IdentityRoleDtoPagedResultDto struct {
	Items []IdentityRoleDto `json:"items,omitempty"`
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// NewIdentityRoleDtoPagedResultDto instantiates a new IdentityRoleDtoPagedResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRoleDtoPagedResultDto() *IdentityRoleDtoPagedResultDto {
	this := IdentityRoleDtoPagedResultDto{}
	return &this
}

// NewIdentityRoleDtoPagedResultDtoWithDefaults instantiates a new IdentityRoleDtoPagedResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityRoleDtoPagedResultDtoWithDefaults() *IdentityRoleDtoPagedResultDto {
	this := IdentityRoleDtoPagedResultDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *IdentityRoleDtoPagedResultDto) GetItems() []IdentityRoleDto {
	if o == nil || isNil(o.Items) {
		var ret []IdentityRoleDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDtoPagedResultDto) GetItemsOk() ([]IdentityRoleDto, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *IdentityRoleDtoPagedResultDto) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []IdentityRoleDto and assigns it to the Items field.
func (o *IdentityRoleDtoPagedResultDto) SetItems(v []IdentityRoleDto) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *IdentityRoleDtoPagedResultDto) GetTotalCount() int64 {
	if o == nil || isNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDtoPagedResultDto) GetTotalCountOk() (*int64, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *IdentityRoleDtoPagedResultDto) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *IdentityRoleDtoPagedResultDto) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o IdentityRoleDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityRoleDtoPagedResultDto struct {
	value *IdentityRoleDtoPagedResultDto
	isSet bool
}

func (v NullableIdentityRoleDtoPagedResultDto) Get() *IdentityRoleDtoPagedResultDto {
	return v.value
}

func (v *NullableIdentityRoleDtoPagedResultDto) Set(val *IdentityRoleDtoPagedResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRoleDtoPagedResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRoleDtoPagedResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRoleDtoPagedResultDto(val *IdentityRoleDtoPagedResultDto) *NullableIdentityRoleDtoPagedResultDto {
	return &NullableIdentityRoleDtoPagedResultDto{value: val, isSet: true}
}

func (v NullableIdentityRoleDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRoleDtoPagedResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



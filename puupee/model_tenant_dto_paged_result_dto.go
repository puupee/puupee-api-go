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

// TenantDtoPagedResultDto struct for TenantDtoPagedResultDto
type TenantDtoPagedResultDto struct {
	Items []TenantDto `json:"items,omitempty"`
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// NewTenantDtoPagedResultDto instantiates a new TenantDtoPagedResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantDtoPagedResultDto() *TenantDtoPagedResultDto {
	this := TenantDtoPagedResultDto{}
	return &this
}

// NewTenantDtoPagedResultDtoWithDefaults instantiates a new TenantDtoPagedResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantDtoPagedResultDtoWithDefaults() *TenantDtoPagedResultDto {
	this := TenantDtoPagedResultDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *TenantDtoPagedResultDto) GetItems() []TenantDto {
	if o == nil || isNil(o.Items) {
		var ret []TenantDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDtoPagedResultDto) GetItemsOk() ([]TenantDto, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TenantDtoPagedResultDto) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TenantDto and assigns it to the Items field.
func (o *TenantDtoPagedResultDto) SetItems(v []TenantDto) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *TenantDtoPagedResultDto) GetTotalCount() int64 {
	if o == nil || isNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDtoPagedResultDto) GetTotalCountOk() (*int64, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *TenantDtoPagedResultDto) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *TenantDtoPagedResultDto) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o TenantDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableTenantDtoPagedResultDto struct {
	value *TenantDtoPagedResultDto
	isSet bool
}

func (v NullableTenantDtoPagedResultDto) Get() *TenantDtoPagedResultDto {
	return v.value
}

func (v *NullableTenantDtoPagedResultDto) Set(val *TenantDtoPagedResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantDtoPagedResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantDtoPagedResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantDtoPagedResultDto(val *TenantDtoPagedResultDto) *NullableTenantDtoPagedResultDto {
	return &NullableTenantDtoPagedResultDto{value: val, isSet: true}
}

func (v NullableTenantDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantDtoPagedResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



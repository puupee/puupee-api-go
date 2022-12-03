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

// PuupeeDtoPagedResultDto struct for PuupeeDtoPagedResultDto
type PuupeeDtoPagedResultDto struct {
	Items []PuupeeDto `json:"items,omitempty"`
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// NewPuupeeDtoPagedResultDto instantiates a new PuupeeDtoPagedResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPuupeeDtoPagedResultDto() *PuupeeDtoPagedResultDto {
	this := PuupeeDtoPagedResultDto{}
	return &this
}

// NewPuupeeDtoPagedResultDtoWithDefaults instantiates a new PuupeeDtoPagedResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPuupeeDtoPagedResultDtoWithDefaults() *PuupeeDtoPagedResultDto {
	this := PuupeeDtoPagedResultDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PuupeeDtoPagedResultDto) GetItems() []PuupeeDto {
	if o == nil || isNil(o.Items) {
		var ret []PuupeeDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PuupeeDtoPagedResultDto) GetItemsOk() ([]PuupeeDto, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PuupeeDtoPagedResultDto) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PuupeeDto and assigns it to the Items field.
func (o *PuupeeDtoPagedResultDto) SetItems(v []PuupeeDto) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PuupeeDtoPagedResultDto) GetTotalCount() int64 {
	if o == nil || isNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PuupeeDtoPagedResultDto) GetTotalCountOk() (*int64, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PuupeeDtoPagedResultDto) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *PuupeeDtoPagedResultDto) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o PuupeeDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullablePuupeeDtoPagedResultDto struct {
	value *PuupeeDtoPagedResultDto
	isSet bool
}

func (v NullablePuupeeDtoPagedResultDto) Get() *PuupeeDtoPagedResultDto {
	return v.value
}

func (v *NullablePuupeeDtoPagedResultDto) Set(val *PuupeeDtoPagedResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePuupeeDtoPagedResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePuupeeDtoPagedResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePuupeeDtoPagedResultDto(val *PuupeeDtoPagedResultDto) *NullablePuupeeDtoPagedResultDto {
	return &NullablePuupeeDtoPagedResultDto{value: val, isSet: true}
}

func (v NullablePuupeeDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePuupeeDtoPagedResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


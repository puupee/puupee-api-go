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

// checks if the StringSetKeyValueDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringSetKeyValueDto{}

// StringSetKeyValueDto struct for StringSetKeyValueDto
type StringSetKeyValueDto struct {
	Value *string `json:"value,omitempty"`
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`
}

// NewStringSetKeyValueDto instantiates a new StringSetKeyValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringSetKeyValueDto() *StringSetKeyValueDto {
	this := StringSetKeyValueDto{}
	return &this
}

// NewStringSetKeyValueDtoWithDefaults instantiates a new StringSetKeyValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringSetKeyValueDtoWithDefaults() *StringSetKeyValueDto {
	this := StringSetKeyValueDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StringSetKeyValueDto) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSetKeyValueDto) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StringSetKeyValueDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *StringSetKeyValueDto) SetValue(v string) {
	o.Value = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *StringSetKeyValueDto) GetDurationSeconds() float64 {
	if o == nil || IsNil(o.DurationSeconds) {
		var ret float64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringSetKeyValueDto) GetDurationSecondsOk() (*float64, bool) {
	if o == nil || IsNil(o.DurationSeconds) {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *StringSetKeyValueDto) HasDurationSeconds() bool {
	if o != nil && !IsNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given float64 and assigns it to the DurationSeconds field.
func (o *StringSetKeyValueDto) SetDurationSeconds(v float64) {
	o.DurationSeconds = &v
}

func (o StringSetKeyValueDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringSetKeyValueDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.DurationSeconds) {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	return toSerialize, nil
}

type NullableStringSetKeyValueDto struct {
	value *StringSetKeyValueDto
	isSet bool
}

func (v NullableStringSetKeyValueDto) Get() *StringSetKeyValueDto {
	return v.value
}

func (v *NullableStringSetKeyValueDto) Set(val *StringSetKeyValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStringSetKeyValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStringSetKeyValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringSetKeyValueDto(val *StringSetKeyValueDto) *NullableStringSetKeyValueDto {
	return &NullableStringSetKeyValueDto{value: val, isSet: true}
}

func (v NullableStringSetKeyValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringSetKeyValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



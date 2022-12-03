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

// DecimalSetKeyValueDto struct for DecimalSetKeyValueDto
type DecimalSetKeyValueDto struct {
	Value *float64 `json:"value,omitempty"`
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`
}

// NewDecimalSetKeyValueDto instantiates a new DecimalSetKeyValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecimalSetKeyValueDto() *DecimalSetKeyValueDto {
	this := DecimalSetKeyValueDto{}
	return &this
}

// NewDecimalSetKeyValueDtoWithDefaults instantiates a new DecimalSetKeyValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecimalSetKeyValueDtoWithDefaults() *DecimalSetKeyValueDto {
	this := DecimalSetKeyValueDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DecimalSetKeyValueDto) GetValue() float64 {
	if o == nil || isNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecimalSetKeyValueDto) GetValueOk() (*float64, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DecimalSetKeyValueDto) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *DecimalSetKeyValueDto) SetValue(v float64) {
	o.Value = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *DecimalSetKeyValueDto) GetDurationSeconds() float64 {
	if o == nil || isNil(o.DurationSeconds) {
		var ret float64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecimalSetKeyValueDto) GetDurationSecondsOk() (*float64, bool) {
	if o == nil || isNil(o.DurationSeconds) {
    return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *DecimalSetKeyValueDto) HasDurationSeconds() bool {
	if o != nil && !isNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given float64 and assigns it to the DurationSeconds field.
func (o *DecimalSetKeyValueDto) SetDurationSeconds(v float64) {
	o.DurationSeconds = &v
}

func (o DecimalSetKeyValueDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.DurationSeconds) {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableDecimalSetKeyValueDto struct {
	value *DecimalSetKeyValueDto
	isSet bool
}

func (v NullableDecimalSetKeyValueDto) Get() *DecimalSetKeyValueDto {
	return v.value
}

func (v *NullableDecimalSetKeyValueDto) Set(val *DecimalSetKeyValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDecimalSetKeyValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDecimalSetKeyValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecimalSetKeyValueDto(val *DecimalSetKeyValueDto) *NullableDecimalSetKeyValueDto {
	return &NullableDecimalSetKeyValueDto{value: val, isSet: true}
}

func (v NullableDecimalSetKeyValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecimalSetKeyValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



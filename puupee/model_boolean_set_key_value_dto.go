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

// BooleanSetKeyValueDto struct for BooleanSetKeyValueDto
type BooleanSetKeyValueDto struct {
	Value *bool `json:"value,omitempty"`
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`
}

// NewBooleanSetKeyValueDto instantiates a new BooleanSetKeyValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBooleanSetKeyValueDto() *BooleanSetKeyValueDto {
	this := BooleanSetKeyValueDto{}
	return &this
}

// NewBooleanSetKeyValueDtoWithDefaults instantiates a new BooleanSetKeyValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBooleanSetKeyValueDtoWithDefaults() *BooleanSetKeyValueDto {
	this := BooleanSetKeyValueDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BooleanSetKeyValueDto) GetValue() bool {
	if o == nil || isNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BooleanSetKeyValueDto) GetValueOk() (*bool, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BooleanSetKeyValueDto) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *BooleanSetKeyValueDto) SetValue(v bool) {
	o.Value = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *BooleanSetKeyValueDto) GetDurationSeconds() float64 {
	if o == nil || isNil(o.DurationSeconds) {
		var ret float64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BooleanSetKeyValueDto) GetDurationSecondsOk() (*float64, bool) {
	if o == nil || isNil(o.DurationSeconds) {
    return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *BooleanSetKeyValueDto) HasDurationSeconds() bool {
	if o != nil && !isNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given float64 and assigns it to the DurationSeconds field.
func (o *BooleanSetKeyValueDto) SetDurationSeconds(v float64) {
	o.DurationSeconds = &v
}

func (o BooleanSetKeyValueDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.DurationSeconds) {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableBooleanSetKeyValueDto struct {
	value *BooleanSetKeyValueDto
	isSet bool
}

func (v NullableBooleanSetKeyValueDto) Get() *BooleanSetKeyValueDto {
	return v.value
}

func (v *NullableBooleanSetKeyValueDto) Set(val *BooleanSetKeyValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBooleanSetKeyValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBooleanSetKeyValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBooleanSetKeyValueDto(val *BooleanSetKeyValueDto) *NullableBooleanSetKeyValueDto {
	return &NullableBooleanSetKeyValueDto{value: val, isSet: true}
}

func (v NullableBooleanSetKeyValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBooleanSetKeyValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



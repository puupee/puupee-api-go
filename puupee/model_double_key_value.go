/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// DoubleKeyValue struct for DoubleKeyValue
type DoubleKeyValue struct {
	Value *float64 `json:"value,omitempty"`
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewDoubleKeyValue instantiates a new DoubleKeyValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDoubleKeyValue() *DoubleKeyValue {
	this := DoubleKeyValue{}
	return &this
}

// NewDoubleKeyValueWithDefaults instantiates a new DoubleKeyValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDoubleKeyValueWithDefaults() *DoubleKeyValue {
	this := DoubleKeyValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DoubleKeyValue) GetValue() float64 {
	if o == nil || isNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoubleKeyValue) GetValueOk() (*float64, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DoubleKeyValue) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *DoubleKeyValue) SetValue(v float64) {
	o.Value = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *DoubleKeyValue) GetDurationSeconds() float64 {
	if o == nil || isNil(o.DurationSeconds) {
		var ret float64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoubleKeyValue) GetDurationSecondsOk() (*float64, bool) {
	if o == nil || isNil(o.DurationSeconds) {
    return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *DoubleKeyValue) HasDurationSeconds() bool {
	if o != nil && !isNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given float64 and assigns it to the DurationSeconds field.
func (o *DoubleKeyValue) SetDurationSeconds(v float64) {
	o.DurationSeconds = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *DoubleKeyValue) GetExpiredAt() time.Time {
	if o == nil || isNil(o.ExpiredAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoubleKeyValue) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiredAt) {
    return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *DoubleKeyValue) HasExpiredAt() bool {
	if o != nil && !isNil(o.ExpiredAt) {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *DoubleKeyValue) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DoubleKeyValue) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DoubleKeyValue) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DoubleKeyValue) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DoubleKeyValue) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o DoubleKeyValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.DurationSeconds) {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	if !isNil(o.ExpiredAt) {
		toSerialize["expiredAt"] = o.ExpiredAt
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableDoubleKeyValue struct {
	value *DoubleKeyValue
	isSet bool
}

func (v NullableDoubleKeyValue) Get() *DoubleKeyValue {
	return v.value
}

func (v *NullableDoubleKeyValue) Set(val *DoubleKeyValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDoubleKeyValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDoubleKeyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDoubleKeyValue(val *DoubleKeyValue) *NullableDoubleKeyValue {
	return &NullableDoubleKeyValue{value: val, isSet: true}
}

func (v NullableDoubleKeyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDoubleKeyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



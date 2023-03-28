/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// Int32KeyValue struct for Int32KeyValue
type Int32KeyValue struct {
	Value *int32 `json:"value,omitempty"`
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewInt32KeyValue instantiates a new Int32KeyValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInt32KeyValue() *Int32KeyValue {
	this := Int32KeyValue{}
	return &this
}

// NewInt32KeyValueWithDefaults instantiates a new Int32KeyValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInt32KeyValueWithDefaults() *Int32KeyValue {
	this := Int32KeyValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Int32KeyValue) GetValue() int32 {
	if o == nil || isNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Int32KeyValue) GetValueOk() (*int32, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Int32KeyValue) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *Int32KeyValue) SetValue(v int32) {
	o.Value = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *Int32KeyValue) GetDurationSeconds() float64 {
	if o == nil || isNil(o.DurationSeconds) {
		var ret float64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Int32KeyValue) GetDurationSecondsOk() (*float64, bool) {
	if o == nil || isNil(o.DurationSeconds) {
    return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *Int32KeyValue) HasDurationSeconds() bool {
	if o != nil && !isNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given float64 and assigns it to the DurationSeconds field.
func (o *Int32KeyValue) SetDurationSeconds(v float64) {
	o.DurationSeconds = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *Int32KeyValue) GetExpiredAt() time.Time {
	if o == nil || isNil(o.ExpiredAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Int32KeyValue) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiredAt) {
    return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *Int32KeyValue) HasExpiredAt() bool {
	if o != nil && !isNil(o.ExpiredAt) {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *Int32KeyValue) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Int32KeyValue) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Int32KeyValue) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Int32KeyValue) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Int32KeyValue) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o Int32KeyValue) MarshalJSON() ([]byte, error) {
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

type NullableInt32KeyValue struct {
	value *Int32KeyValue
	isSet bool
}

func (v NullableInt32KeyValue) Get() *Int32KeyValue {
	return v.value
}

func (v *NullableInt32KeyValue) Set(val *Int32KeyValue) {
	v.value = val
	v.isSet = true
}

func (v NullableInt32KeyValue) IsSet() bool {
	return v.isSet
}

func (v *NullableInt32KeyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt32KeyValue(val *Int32KeyValue) *NullableInt32KeyValue {
	return &NullableInt32KeyValue{value: val, isSet: true}
}

func (v NullableInt32KeyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt32KeyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



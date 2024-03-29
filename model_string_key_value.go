/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// checks if the StringKeyValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringKeyValue{}

// StringKeyValue struct for StringKeyValue
type StringKeyValue struct {
	Value *string `json:"value,omitempty"`
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewStringKeyValue instantiates a new StringKeyValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringKeyValue() *StringKeyValue {
	this := StringKeyValue{}
	return &this
}

// NewStringKeyValueWithDefaults instantiates a new StringKeyValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringKeyValueWithDefaults() *StringKeyValue {
	this := StringKeyValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StringKeyValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringKeyValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StringKeyValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *StringKeyValue) SetValue(v string) {
	o.Value = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *StringKeyValue) GetDurationSeconds() float64 {
	if o == nil || IsNil(o.DurationSeconds) {
		var ret float64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringKeyValue) GetDurationSecondsOk() (*float64, bool) {
	if o == nil || IsNil(o.DurationSeconds) {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *StringKeyValue) HasDurationSeconds() bool {
	if o != nil && !IsNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given float64 and assigns it to the DurationSeconds field.
func (o *StringKeyValue) SetDurationSeconds(v float64) {
	o.DurationSeconds = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *StringKeyValue) GetExpiredAt() time.Time {
	if o == nil || IsNil(o.ExpiredAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringKeyValue) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiredAt) {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *StringKeyValue) HasExpiredAt() bool {
	if o != nil && !IsNil(o.ExpiredAt) {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *StringKeyValue) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StringKeyValue) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringKeyValue) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StringKeyValue) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *StringKeyValue) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o StringKeyValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringKeyValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.DurationSeconds) {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	if !IsNil(o.ExpiredAt) {
		toSerialize["expiredAt"] = o.ExpiredAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableStringKeyValue struct {
	value *StringKeyValue
	isSet bool
}

func (v NullableStringKeyValue) Get() *StringKeyValue {
	return v.value
}

func (v *NullableStringKeyValue) Set(val *StringKeyValue) {
	v.value = val
	v.isSet = true
}

func (v NullableStringKeyValue) IsSet() bool {
	return v.isSet
}

func (v *NullableStringKeyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringKeyValue(val *StringKeyValue) *NullableStringKeyValue {
	return &NullableStringKeyValue{value: val, isSet: true}
}

func (v NullableStringKeyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringKeyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



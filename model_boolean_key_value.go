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

// checks if the BooleanKeyValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BooleanKeyValue{}

// BooleanKeyValue struct for BooleanKeyValue
type BooleanKeyValue struct {
	Value *bool `json:"value,omitempty"`
	DurationSeconds NullableFloat64 `json:"durationSeconds,omitempty"`
	ExpiredAt NullableTime `json:"expiredAt,omitempty"`
	CreatedAt NullableTime `json:"createdAt,omitempty"`
}

// NewBooleanKeyValue instantiates a new BooleanKeyValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBooleanKeyValue() *BooleanKeyValue {
	this := BooleanKeyValue{}
	return &this
}

// NewBooleanKeyValueWithDefaults instantiates a new BooleanKeyValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBooleanKeyValueWithDefaults() *BooleanKeyValue {
	this := BooleanKeyValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BooleanKeyValue) GetValue() bool {
	if o == nil || IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BooleanKeyValue) GetValueOk() (*bool, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BooleanKeyValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *BooleanKeyValue) SetValue(v bool) {
	o.Value = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BooleanKeyValue) GetDurationSeconds() float64 {
	if o == nil || IsNil(o.DurationSeconds.Get()) {
		var ret float64
		return ret
	}
	return *o.DurationSeconds.Get()
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BooleanKeyValue) GetDurationSecondsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationSeconds.Get(), o.DurationSeconds.IsSet()
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *BooleanKeyValue) HasDurationSeconds() bool {
	if o != nil && o.DurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given NullableFloat64 and assigns it to the DurationSeconds field.
func (o *BooleanKeyValue) SetDurationSeconds(v float64) {
	o.DurationSeconds.Set(&v)
}
// SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil
func (o *BooleanKeyValue) SetDurationSecondsNil() {
	o.DurationSeconds.Set(nil)
}

// UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
func (o *BooleanKeyValue) UnsetDurationSeconds() {
	o.DurationSeconds.Unset()
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BooleanKeyValue) GetExpiredAt() time.Time {
	if o == nil || IsNil(o.ExpiredAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BooleanKeyValue) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *BooleanKeyValue) HasExpiredAt() bool {
	if o != nil && o.ExpiredAt.IsSet() {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given NullableTime and assigns it to the ExpiredAt field.
func (o *BooleanKeyValue) SetExpiredAt(v time.Time) {
	o.ExpiredAt.Set(&v)
}
// SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil
func (o *BooleanKeyValue) SetExpiredAtNil() {
	o.ExpiredAt.Set(nil)
}

// UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
func (o *BooleanKeyValue) UnsetExpiredAt() {
	o.ExpiredAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BooleanKeyValue) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BooleanKeyValue) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BooleanKeyValue) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *BooleanKeyValue) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *BooleanKeyValue) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *BooleanKeyValue) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

func (o BooleanKeyValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BooleanKeyValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if o.DurationSeconds.IsSet() {
		toSerialize["durationSeconds"] = o.DurationSeconds.Get()
	}
	if o.ExpiredAt.IsSet() {
		toSerialize["expiredAt"] = o.ExpiredAt.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	return toSerialize, nil
}

type NullableBooleanKeyValue struct {
	value *BooleanKeyValue
	isSet bool
}

func (v NullableBooleanKeyValue) Get() *BooleanKeyValue {
	return v.value
}

func (v *NullableBooleanKeyValue) Set(val *BooleanKeyValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBooleanKeyValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBooleanKeyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBooleanKeyValue(val *BooleanKeyValue) *NullableBooleanKeyValue {
	return &NullableBooleanKeyValue{value: val, isSet: true}
}

func (v NullableBooleanKeyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBooleanKeyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the DateTimeSetKeyValueDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateTimeSetKeyValueDto{}

// DateTimeSetKeyValueDto struct for DateTimeSetKeyValueDto
type DateTimeSetKeyValueDto struct {
	Value *time.Time `json:"value,omitempty"`
	DurationSeconds *float64 `json:"durationSeconds,omitempty"`
}

// NewDateTimeSetKeyValueDto instantiates a new DateTimeSetKeyValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeSetKeyValueDto() *DateTimeSetKeyValueDto {
	this := DateTimeSetKeyValueDto{}
	return &this
}

// NewDateTimeSetKeyValueDtoWithDefaults instantiates a new DateTimeSetKeyValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeSetKeyValueDtoWithDefaults() *DateTimeSetKeyValueDto {
	this := DateTimeSetKeyValueDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DateTimeSetKeyValueDto) GetValue() time.Time {
	if o == nil || IsNil(o.Value) {
		var ret time.Time
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeSetKeyValueDto) GetValueOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DateTimeSetKeyValueDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given time.Time and assigns it to the Value field.
func (o *DateTimeSetKeyValueDto) SetValue(v time.Time) {
	o.Value = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *DateTimeSetKeyValueDto) GetDurationSeconds() float64 {
	if o == nil || IsNil(o.DurationSeconds) {
		var ret float64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeSetKeyValueDto) GetDurationSecondsOk() (*float64, bool) {
	if o == nil || IsNil(o.DurationSeconds) {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *DateTimeSetKeyValueDto) HasDurationSeconds() bool {
	if o != nil && !IsNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given float64 and assigns it to the DurationSeconds field.
func (o *DateTimeSetKeyValueDto) SetDurationSeconds(v float64) {
	o.DurationSeconds = &v
}

func (o DateTimeSetKeyValueDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateTimeSetKeyValueDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.DurationSeconds) {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	return toSerialize, nil
}

type NullableDateTimeSetKeyValueDto struct {
	value *DateTimeSetKeyValueDto
	isSet bool
}

func (v NullableDateTimeSetKeyValueDto) Get() *DateTimeSetKeyValueDto {
	return v.value
}

func (v *NullableDateTimeSetKeyValueDto) Set(val *DateTimeSetKeyValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeSetKeyValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeSetKeyValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeSetKeyValueDto(val *DateTimeSetKeyValueDto) *NullableDateTimeSetKeyValueDto {
	return &NullableDateTimeSetKeyValueDto{value: val, isSet: true}
}

func (v NullableDateTimeSetKeyValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateTimeSetKeyValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



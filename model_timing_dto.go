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

// TimingDto struct for TimingDto
type TimingDto struct {
	TimeZone *TimeZone `json:"timeZone,omitempty"`
}

// NewTimingDto instantiates a new TimingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimingDto() *TimingDto {
	this := TimingDto{}
	return &this
}

// NewTimingDtoWithDefaults instantiates a new TimingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimingDtoWithDefaults() *TimingDto {
	this := TimingDto{}
	return &this
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *TimingDto) GetTimeZone() TimeZone {
	if o == nil || isNil(o.TimeZone) {
		var ret TimeZone
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimingDto) GetTimeZoneOk() (*TimeZone, bool) {
	if o == nil || isNil(o.TimeZone) {
    return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *TimingDto) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given TimeZone and assigns it to the TimeZone field.
func (o *TimingDto) SetTimeZone(v TimeZone) {
	o.TimeZone = &v
}

func (o TimingDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return json.Marshal(toSerialize)
}

type NullableTimingDto struct {
	value *TimingDto
	isSet bool
}

func (v NullableTimingDto) Get() *TimingDto {
	return v.value
}

func (v *NullableTimingDto) Set(val *TimingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTimingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTimingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimingDto(val *TimingDto) *NullableTimingDto {
	return &NullableTimingDto{value: val, isSet: true}
}

func (v NullableTimingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



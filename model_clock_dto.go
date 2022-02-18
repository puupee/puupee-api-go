/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
)

// ClockDto struct for ClockDto
type ClockDto struct {
	Kind NullableString `json:"kind,omitempty"`
}

// NewClockDto instantiates a new ClockDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClockDto() *ClockDto {
	this := ClockDto{}
	return &this
}

// NewClockDtoWithDefaults instantiates a new ClockDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClockDtoWithDefaults() *ClockDto {
	this := ClockDto{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClockDto) GetKind() string {
	if o == nil || o.Kind.Get() == nil {
		var ret string
		return ret
	}
	return *o.Kind.Get()
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClockDto) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Kind.Get(), o.Kind.IsSet()
}

// HasKind returns a boolean if a field has been set.
func (o *ClockDto) HasKind() bool {
	if o != nil && o.Kind.IsSet() {
		return true
	}

	return false
}

// SetKind gets a reference to the given NullableString and assigns it to the Kind field.
func (o *ClockDto) SetKind(v string) {
	o.Kind.Set(&v)
}
// SetKindNil sets the value for Kind to be an explicit nil
func (o *ClockDto) SetKindNil() {
	o.Kind.Set(nil)
}

// UnsetKind ensures that no value is present for Kind, not even an explicit nil
func (o *ClockDto) UnsetKind() {
	o.Kind.Unset()
}

func (o ClockDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind.IsSet() {
		toSerialize["kind"] = o.Kind.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableClockDto struct {
	value *ClockDto
	isSet bool
}

func (v NullableClockDto) Get() *ClockDto {
	return v.value
}

func (v *NullableClockDto) Set(val *ClockDto) {
	v.value = val
	v.isSet = true
}

func (v NullableClockDto) IsSet() bool {
	return v.isSet
}

func (v *NullableClockDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClockDto(val *ClockDto) *NullableClockDto {
	return &NullableClockDto{value: val, isSet: true}
}

func (v NullableClockDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClockDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



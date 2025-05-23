/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.17.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// checks if the IanaTimeZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IanaTimeZone{}

// IanaTimeZone struct for IanaTimeZone
type IanaTimeZone struct {
	TimeZoneName *string `json:"timeZoneName,omitempty"`
}

// NewIanaTimeZone instantiates a new IanaTimeZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIanaTimeZone() *IanaTimeZone {
	this := IanaTimeZone{}
	return &this
}

// NewIanaTimeZoneWithDefaults instantiates a new IanaTimeZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIanaTimeZoneWithDefaults() *IanaTimeZone {
	this := IanaTimeZone{}
	return &this
}

// GetTimeZoneName returns the TimeZoneName field value if set, zero value otherwise.
func (o *IanaTimeZone) GetTimeZoneName() string {
	if o == nil || IsNil(o.TimeZoneName) {
		var ret string
		return ret
	}
	return *o.TimeZoneName
}

// GetTimeZoneNameOk returns a tuple with the TimeZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IanaTimeZone) GetTimeZoneNameOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZoneName) {
		return nil, false
	}
	return o.TimeZoneName, true
}

// HasTimeZoneName returns a boolean if a field has been set.
func (o *IanaTimeZone) HasTimeZoneName() bool {
	if o != nil && !IsNil(o.TimeZoneName) {
		return true
	}

	return false
}

// SetTimeZoneName gets a reference to the given string and assigns it to the TimeZoneName field.
func (o *IanaTimeZone) SetTimeZoneName(v string) {
	o.TimeZoneName = &v
}

func (o IanaTimeZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IanaTimeZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeZoneName) {
		toSerialize["timeZoneName"] = o.TimeZoneName
	}
	return toSerialize, nil
}

type NullableIanaTimeZone struct {
	value *IanaTimeZone
	isSet bool
}

func (v NullableIanaTimeZone) Get() *IanaTimeZone {
	return v.value
}

func (v *NullableIanaTimeZone) Set(val *IanaTimeZone) {
	v.value = val
	v.isSet = true
}

func (v NullableIanaTimeZone) IsSet() bool {
	return v.isSet
}

func (v *NullableIanaTimeZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIanaTimeZone(val *IanaTimeZone) *NullableIanaTimeZone {
	return &NullableIanaTimeZone{value: val, isSet: true}
}

func (v NullableIanaTimeZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIanaTimeZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



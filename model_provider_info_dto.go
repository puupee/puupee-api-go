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

// checks if the ProviderInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderInfoDto{}

// ProviderInfoDto struct for ProviderInfoDto
type ProviderInfoDto struct {
	ProviderName *string `json:"providerName,omitempty"`
	ProviderKey *string `json:"providerKey,omitempty"`
}

// NewProviderInfoDto instantiates a new ProviderInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderInfoDto() *ProviderInfoDto {
	this := ProviderInfoDto{}
	return &this
}

// NewProviderInfoDtoWithDefaults instantiates a new ProviderInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderInfoDtoWithDefaults() *ProviderInfoDto {
	this := ProviderInfoDto{}
	return &this
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *ProviderInfoDto) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderInfoDto) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *ProviderInfoDto) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *ProviderInfoDto) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise.
func (o *ProviderInfoDto) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey) {
		var ret string
		return ret
	}
	return *o.ProviderKey
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderInfoDto) GetProviderKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderKey) {
		return nil, false
	}
	return o.ProviderKey, true
}

// HasProviderKey returns a boolean if a field has been set.
func (o *ProviderInfoDto) HasProviderKey() bool {
	if o != nil && !IsNil(o.ProviderKey) {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given string and assigns it to the ProviderKey field.
func (o *ProviderInfoDto) SetProviderKey(v string) {
	o.ProviderKey = &v
}

func (o ProviderInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !IsNil(o.ProviderKey) {
		toSerialize["providerKey"] = o.ProviderKey
	}
	return toSerialize, nil
}

type NullableProviderInfoDto struct {
	value *ProviderInfoDto
	isSet bool
}

func (v NullableProviderInfoDto) Get() *ProviderInfoDto {
	return v.value
}

func (v *NullableProviderInfoDto) Set(val *ProviderInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderInfoDto(val *ProviderInfoDto) *NullableProviderInfoDto {
	return &NullableProviderInfoDto{value: val, isSet: true}
}

func (v NullableProviderInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



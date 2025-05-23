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

// checks if the FeatureProviderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureProviderDto{}

// FeatureProviderDto struct for FeatureProviderDto
type FeatureProviderDto struct {
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewFeatureProviderDto instantiates a new FeatureProviderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureProviderDto() *FeatureProviderDto {
	this := FeatureProviderDto{}
	return &this
}

// NewFeatureProviderDtoWithDefaults instantiates a new FeatureProviderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureProviderDtoWithDefaults() *FeatureProviderDto {
	this := FeatureProviderDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeatureProviderDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureProviderDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeatureProviderDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeatureProviderDto) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FeatureProviderDto) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureProviderDto) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FeatureProviderDto) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FeatureProviderDto) SetKey(v string) {
	o.Key = &v
}

func (o FeatureProviderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureProviderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableFeatureProviderDto struct {
	value *FeatureProviderDto
	isSet bool
}

func (v NullableFeatureProviderDto) Get() *FeatureProviderDto {
	return v.value
}

func (v *NullableFeatureProviderDto) Set(val *FeatureProviderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureProviderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureProviderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureProviderDto(val *FeatureProviderDto) *NullableFeatureProviderDto {
	return &NullableFeatureProviderDto{value: val, isSet: true}
}

func (v NullableFeatureProviderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureProviderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



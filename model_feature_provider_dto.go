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

// checks if the FeatureProviderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureProviderDto{}

// FeatureProviderDto struct for FeatureProviderDto
type FeatureProviderDto struct {
	Name NullableString `json:"name,omitempty"`
	Key NullableString `json:"key,omitempty"`
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

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureProviderDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureProviderDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FeatureProviderDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FeatureProviderDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FeatureProviderDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FeatureProviderDto) UnsetName() {
	o.Name.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureProviderDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureProviderDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *FeatureProviderDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *FeatureProviderDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *FeatureProviderDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *FeatureProviderDto) UnsetKey() {
	o.Key.Unset()
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
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
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



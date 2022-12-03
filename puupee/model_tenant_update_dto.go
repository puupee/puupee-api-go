/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// TenantUpdateDto struct for TenantUpdateDto
type TenantUpdateDto struct {
	ExtraProperties map[string]map[string]interface{} `json:"extraProperties,omitempty"`
	Name string `json:"name"`
	ConcurrencyStamp *string `json:"concurrencyStamp,omitempty"`
}

// NewTenantUpdateDto instantiates a new TenantUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantUpdateDto(name string) *TenantUpdateDto {
	this := TenantUpdateDto{}
	this.Name = name
	return &this
}

// NewTenantUpdateDtoWithDefaults instantiates a new TenantUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantUpdateDtoWithDefaults() *TenantUpdateDto {
	this := TenantUpdateDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise.
func (o *TenantUpdateDto) GetExtraProperties() map[string]map[string]interface{} {
	if o == nil || isNil(o.ExtraProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantUpdateDto) GetExtraPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.ExtraProperties) {
    return map[string]map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *TenantUpdateDto) HasExtraProperties() bool {
	if o != nil && !isNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtraProperties field.
func (o *TenantUpdateDto) SetExtraProperties(v map[string]map[string]interface{}) {
	o.ExtraProperties = v
}

// GetName returns the Name field value
func (o *TenantUpdateDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TenantUpdateDto) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TenantUpdateDto) SetName(v string) {
	o.Name = v
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise.
func (o *TenantUpdateDto) GetConcurrencyStamp() string {
	if o == nil || isNil(o.ConcurrencyStamp) {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantUpdateDto) GetConcurrencyStampOk() (*string, bool) {
	if o == nil || isNil(o.ConcurrencyStamp) {
    return nil, false
	}
	return o.ConcurrencyStamp, true
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *TenantUpdateDto) HasConcurrencyStamp() bool {
	if o != nil && !isNil(o.ConcurrencyStamp) {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given string and assigns it to the ConcurrencyStamp field.
func (o *TenantUpdateDto) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp = &v
}

func (o TenantUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExtraProperties) {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ConcurrencyStamp) {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp
	}
	return json.Marshal(toSerialize)
}

type NullableTenantUpdateDto struct {
	value *TenantUpdateDto
	isSet bool
}

func (v NullableTenantUpdateDto) Get() *TenantUpdateDto {
	return v.value
}

func (v *NullableTenantUpdateDto) Set(val *TenantUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantUpdateDto(val *TenantUpdateDto) *NullableTenantUpdateDto {
	return &NullableTenantUpdateDto{value: val, isSet: true}
}

func (v NullableTenantUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


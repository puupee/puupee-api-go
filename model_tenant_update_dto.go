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

// checks if the TenantUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantUpdateDto{}

// TenantUpdateDto struct for TenantUpdateDto
type TenantUpdateDto struct {
	ExtraProperties map[string]interface{} `json:"extraProperties,omitempty"`
	Name string `json:"name"`
	ConcurrencyStamp NullableString `json:"concurrencyStamp,omitempty"`
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

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantUpdateDto) GetExtraProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantUpdateDto) GetExtraPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtraProperties) {
		return map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *TenantUpdateDto) HasExtraProperties() bool {
	if o != nil && IsNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]interface{} and assigns it to the ExtraProperties field.
func (o *TenantUpdateDto) SetExtraProperties(v map[string]interface{}) {
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

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantUpdateDto) GetConcurrencyStamp() string {
	if o == nil || IsNil(o.ConcurrencyStamp.Get()) {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp.Get()
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantUpdateDto) GetConcurrencyStampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConcurrencyStamp.Get(), o.ConcurrencyStamp.IsSet()
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *TenantUpdateDto) HasConcurrencyStamp() bool {
	if o != nil && o.ConcurrencyStamp.IsSet() {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given NullableString and assigns it to the ConcurrencyStamp field.
func (o *TenantUpdateDto) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp.Set(&v)
}
// SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil
func (o *TenantUpdateDto) SetConcurrencyStampNil() {
	o.ConcurrencyStamp.Set(nil)
}

// UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
func (o *TenantUpdateDto) UnsetConcurrencyStamp() {
	o.ConcurrencyStamp.Unset()
}

func (o TenantUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtraProperties != nil {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	toSerialize["name"] = o.Name
	if o.ConcurrencyStamp.IsSet() {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp.Get()
	}
	return toSerialize, nil
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



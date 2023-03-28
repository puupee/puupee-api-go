/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// TenantDto struct for TenantDto
type TenantDto struct {
	ExtraProperties map[string]map[string]interface{} `json:"extraProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ConcurrencyStamp *string `json:"concurrencyStamp,omitempty"`
}

// NewTenantDto instantiates a new TenantDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantDto() *TenantDto {
	this := TenantDto{}
	return &this
}

// NewTenantDtoWithDefaults instantiates a new TenantDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantDtoWithDefaults() *TenantDto {
	this := TenantDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise.
func (o *TenantDto) GetExtraProperties() map[string]map[string]interface{} {
	if o == nil || isNil(o.ExtraProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetExtraPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.ExtraProperties) {
    return map[string]map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *TenantDto) HasExtraProperties() bool {
	if o != nil && !isNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtraProperties field.
func (o *TenantDto) SetExtraProperties(v map[string]map[string]interface{}) {
	o.ExtraProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TenantDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TenantDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TenantDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TenantDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TenantDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TenantDto) SetName(v string) {
	o.Name = &v
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise.
func (o *TenantDto) GetConcurrencyStamp() string {
	if o == nil || isNil(o.ConcurrencyStamp) {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDto) GetConcurrencyStampOk() (*string, bool) {
	if o == nil || isNil(o.ConcurrencyStamp) {
    return nil, false
	}
	return o.ConcurrencyStamp, true
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *TenantDto) HasConcurrencyStamp() bool {
	if o != nil && !isNil(o.ConcurrencyStamp) {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given string and assigns it to the ConcurrencyStamp field.
func (o *TenantDto) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp = &v
}

func (o TenantDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExtraProperties) {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ConcurrencyStamp) {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp
	}
	return json.Marshal(toSerialize)
}

type NullableTenantDto struct {
	value *TenantDto
	isSet bool
}

func (v NullableTenantDto) Get() *TenantDto {
	return v.value
}

func (v *NullableTenantDto) Set(val *TenantDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantDto(val *TenantDto) *NullableTenantDto {
	return &NullableTenantDto{value: val, isSet: true}
}

func (v NullableTenantDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



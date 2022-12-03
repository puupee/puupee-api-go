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

// IdentityRoleUpdateDto struct for IdentityRoleUpdateDto
type IdentityRoleUpdateDto struct {
	ExtraProperties map[string]map[string]interface{} `json:"extraProperties,omitempty"`
	Name string `json:"name"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
	ConcurrencyStamp *string `json:"concurrencyStamp,omitempty"`
}

// NewIdentityRoleUpdateDto instantiates a new IdentityRoleUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRoleUpdateDto(name string) *IdentityRoleUpdateDto {
	this := IdentityRoleUpdateDto{}
	this.Name = name
	return &this
}

// NewIdentityRoleUpdateDtoWithDefaults instantiates a new IdentityRoleUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityRoleUpdateDtoWithDefaults() *IdentityRoleUpdateDto {
	this := IdentityRoleUpdateDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise.
func (o *IdentityRoleUpdateDto) GetExtraProperties() map[string]map[string]interface{} {
	if o == nil || isNil(o.ExtraProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleUpdateDto) GetExtraPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.ExtraProperties) {
    return map[string]map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *IdentityRoleUpdateDto) HasExtraProperties() bool {
	if o != nil && !isNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtraProperties field.
func (o *IdentityRoleUpdateDto) SetExtraProperties(v map[string]map[string]interface{}) {
	o.ExtraProperties = v
}

// GetName returns the Name field value
func (o *IdentityRoleUpdateDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityRoleUpdateDto) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityRoleUpdateDto) SetName(v string) {
	o.Name = v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *IdentityRoleUpdateDto) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleUpdateDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsDefault) {
    return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *IdentityRoleUpdateDto) HasIsDefault() bool {
	if o != nil && !isNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *IdentityRoleUpdateDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *IdentityRoleUpdateDto) GetIsPublic() bool {
	if o == nil || isNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleUpdateDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || isNil(o.IsPublic) {
    return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *IdentityRoleUpdateDto) HasIsPublic() bool {
	if o != nil && !isNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *IdentityRoleUpdateDto) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise.
func (o *IdentityRoleUpdateDto) GetConcurrencyStamp() string {
	if o == nil || isNil(o.ConcurrencyStamp) {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleUpdateDto) GetConcurrencyStampOk() (*string, bool) {
	if o == nil || isNil(o.ConcurrencyStamp) {
    return nil, false
	}
	return o.ConcurrencyStamp, true
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *IdentityRoleUpdateDto) HasConcurrencyStamp() bool {
	if o != nil && !isNil(o.ConcurrencyStamp) {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given string and assigns it to the ConcurrencyStamp field.
func (o *IdentityRoleUpdateDto) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp = &v
}

func (o IdentityRoleUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExtraProperties) {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !isNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	if !isNil(o.ConcurrencyStamp) {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityRoleUpdateDto struct {
	value *IdentityRoleUpdateDto
	isSet bool
}

func (v NullableIdentityRoleUpdateDto) Get() *IdentityRoleUpdateDto {
	return v.value
}

func (v *NullableIdentityRoleUpdateDto) Set(val *IdentityRoleUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRoleUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRoleUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRoleUpdateDto(val *IdentityRoleUpdateDto) *NullableIdentityRoleUpdateDto {
	return &NullableIdentityRoleUpdateDto{value: val, isSet: true}
}

func (v NullableIdentityRoleUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRoleUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



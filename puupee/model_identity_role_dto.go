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

// IdentityRoleDto struct for IdentityRoleDto
type IdentityRoleDto struct {
	ExtraProperties map[string]map[string]interface{} `json:"extraProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsStatic *bool `json:"isStatic,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
	ConcurrencyStamp *string `json:"concurrencyStamp,omitempty"`
}

// NewIdentityRoleDto instantiates a new IdentityRoleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRoleDto() *IdentityRoleDto {
	this := IdentityRoleDto{}
	return &this
}

// NewIdentityRoleDtoWithDefaults instantiates a new IdentityRoleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityRoleDtoWithDefaults() *IdentityRoleDto {
	this := IdentityRoleDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetExtraProperties() map[string]map[string]interface{} {
	if o == nil || isNil(o.ExtraProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetExtraPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.ExtraProperties) {
    return map[string]map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasExtraProperties() bool {
	if o != nil && !isNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtraProperties field.
func (o *IdentityRoleDto) SetExtraProperties(v map[string]map[string]interface{}) {
	o.ExtraProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityRoleDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityRoleDto) SetName(v string) {
	o.Name = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsDefault) {
    return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasIsDefault() bool {
	if o != nil && !isNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *IdentityRoleDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsStatic returns the IsStatic field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetIsStatic() bool {
	if o == nil || isNil(o.IsStatic) {
		var ret bool
		return ret
	}
	return *o.IsStatic
}

// GetIsStaticOk returns a tuple with the IsStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetIsStaticOk() (*bool, bool) {
	if o == nil || isNil(o.IsStatic) {
    return nil, false
	}
	return o.IsStatic, true
}

// HasIsStatic returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasIsStatic() bool {
	if o != nil && !isNil(o.IsStatic) {
		return true
	}

	return false
}

// SetIsStatic gets a reference to the given bool and assigns it to the IsStatic field.
func (o *IdentityRoleDto) SetIsStatic(v bool) {
	o.IsStatic = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetIsPublic() bool {
	if o == nil || isNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || isNil(o.IsPublic) {
    return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasIsPublic() bool {
	if o != nil && !isNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *IdentityRoleDto) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise.
func (o *IdentityRoleDto) GetConcurrencyStamp() string {
	if o == nil || isNil(o.ConcurrencyStamp) {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleDto) GetConcurrencyStampOk() (*string, bool) {
	if o == nil || isNil(o.ConcurrencyStamp) {
    return nil, false
	}
	return o.ConcurrencyStamp, true
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *IdentityRoleDto) HasConcurrencyStamp() bool {
	if o != nil && !isNil(o.ConcurrencyStamp) {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given string and assigns it to the ConcurrencyStamp field.
func (o *IdentityRoleDto) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp = &v
}

func (o IdentityRoleDto) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !isNil(o.IsStatic) {
		toSerialize["isStatic"] = o.IsStatic
	}
	if !isNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	if !isNil(o.ConcurrencyStamp) {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityRoleDto struct {
	value *IdentityRoleDto
	isSet bool
}

func (v NullableIdentityRoleDto) Get() *IdentityRoleDto {
	return v.value
}

func (v *NullableIdentityRoleDto) Set(val *IdentityRoleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRoleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRoleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRoleDto(val *IdentityRoleDto) *NullableIdentityRoleDto {
	return &NullableIdentityRoleDto{value: val, isSet: true}
}

func (v NullableIdentityRoleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRoleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



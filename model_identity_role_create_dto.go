/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.17.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the IdentityRoleCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityRoleCreateDto{}

// IdentityRoleCreateDto struct for IdentityRoleCreateDto
type IdentityRoleCreateDto struct {
	ExtraProperties map[string]map[string]interface{} `json:"extraProperties,omitempty"`
	Name string `json:"name"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
}

type _IdentityRoleCreateDto IdentityRoleCreateDto

// NewIdentityRoleCreateDto instantiates a new IdentityRoleCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRoleCreateDto(name string) *IdentityRoleCreateDto {
	this := IdentityRoleCreateDto{}
	this.Name = name
	return &this
}

// NewIdentityRoleCreateDtoWithDefaults instantiates a new IdentityRoleCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityRoleCreateDtoWithDefaults() *IdentityRoleCreateDto {
	this := IdentityRoleCreateDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise.
func (o *IdentityRoleCreateDto) GetExtraProperties() map[string]map[string]interface{} {
	if o == nil || IsNil(o.ExtraProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleCreateDto) GetExtraPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtraProperties) {
		return map[string]map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *IdentityRoleCreateDto) HasExtraProperties() bool {
	if o != nil && !IsNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtraProperties field.
func (o *IdentityRoleCreateDto) SetExtraProperties(v map[string]map[string]interface{}) {
	o.ExtraProperties = v
}

// GetName returns the Name field value
func (o *IdentityRoleCreateDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityRoleCreateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityRoleCreateDto) SetName(v string) {
	o.Name = v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *IdentityRoleCreateDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleCreateDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *IdentityRoleCreateDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *IdentityRoleCreateDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *IdentityRoleCreateDto) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityRoleCreateDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *IdentityRoleCreateDto) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *IdentityRoleCreateDto) SetIsPublic(v bool) {
	o.IsPublic = &v
}

func (o IdentityRoleCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityRoleCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtraProperties) {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	return toSerialize, nil
}

func (o *IdentityRoleCreateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIdentityRoleCreateDto := _IdentityRoleCreateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentityRoleCreateDto)

	if err != nil {
		return err
	}

	*o = IdentityRoleCreateDto(varIdentityRoleCreateDto)

	return err
}

type NullableIdentityRoleCreateDto struct {
	value *IdentityRoleCreateDto
	isSet bool
}

func (v NullableIdentityRoleCreateDto) Get() *IdentityRoleCreateDto {
	return v.value
}

func (v *NullableIdentityRoleCreateDto) Set(val *IdentityRoleCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRoleCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRoleCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRoleCreateDto(val *IdentityRoleCreateDto) *NullableIdentityRoleCreateDto {
	return &NullableIdentityRoleCreateDto{value: val, isSet: true}
}

func (v NullableIdentityRoleCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRoleCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



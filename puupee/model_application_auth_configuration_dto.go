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

// ApplicationAuthConfigurationDto struct for ApplicationAuthConfigurationDto
type ApplicationAuthConfigurationDto struct {
	GrantedPolicies *map[string]bool `json:"grantedPolicies,omitempty"`
}

// NewApplicationAuthConfigurationDto instantiates a new ApplicationAuthConfigurationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationAuthConfigurationDto() *ApplicationAuthConfigurationDto {
	this := ApplicationAuthConfigurationDto{}
	return &this
}

// NewApplicationAuthConfigurationDtoWithDefaults instantiates a new ApplicationAuthConfigurationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationAuthConfigurationDtoWithDefaults() *ApplicationAuthConfigurationDto {
	this := ApplicationAuthConfigurationDto{}
	return &this
}

// GetGrantedPolicies returns the GrantedPolicies field value if set, zero value otherwise.
func (o *ApplicationAuthConfigurationDto) GetGrantedPolicies() map[string]bool {
	if o == nil || isNil(o.GrantedPolicies) {
		var ret map[string]bool
		return ret
	}
	return *o.GrantedPolicies
}

// GetGrantedPoliciesOk returns a tuple with the GrantedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAuthConfigurationDto) GetGrantedPoliciesOk() (*map[string]bool, bool) {
	if o == nil || isNil(o.GrantedPolicies) {
    return nil, false
	}
	return o.GrantedPolicies, true
}

// HasGrantedPolicies returns a boolean if a field has been set.
func (o *ApplicationAuthConfigurationDto) HasGrantedPolicies() bool {
	if o != nil && !isNil(o.GrantedPolicies) {
		return true
	}

	return false
}

// SetGrantedPolicies gets a reference to the given map[string]bool and assigns it to the GrantedPolicies field.
func (o *ApplicationAuthConfigurationDto) SetGrantedPolicies(v map[string]bool) {
	o.GrantedPolicies = &v
}

func (o ApplicationAuthConfigurationDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GrantedPolicies) {
		toSerialize["grantedPolicies"] = o.GrantedPolicies
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationAuthConfigurationDto struct {
	value *ApplicationAuthConfigurationDto
	isSet bool
}

func (v NullableApplicationAuthConfigurationDto) Get() *ApplicationAuthConfigurationDto {
	return v.value
}

func (v *NullableApplicationAuthConfigurationDto) Set(val *ApplicationAuthConfigurationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationAuthConfigurationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationAuthConfigurationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationAuthConfigurationDto(val *ApplicationAuthConfigurationDto) *NullableApplicationAuthConfigurationDto {
	return &NullableApplicationAuthConfigurationDto{value: val, isSet: true}
}

func (v NullableApplicationAuthConfigurationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationAuthConfigurationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
)

// GetFeatureListResultDto struct for GetFeatureListResultDto
type GetFeatureListResultDto struct {
	Groups []FeatureGroupDto `json:"groups,omitempty"`
}

// NewGetFeatureListResultDto instantiates a new GetFeatureListResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeatureListResultDto() *GetFeatureListResultDto {
	this := GetFeatureListResultDto{}
	return &this
}

// NewGetFeatureListResultDtoWithDefaults instantiates a new GetFeatureListResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeatureListResultDtoWithDefaults() *GetFeatureListResultDto {
	this := GetFeatureListResultDto{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetFeatureListResultDto) GetGroups() []FeatureGroupDto {
	if o == nil  {
		var ret []FeatureGroupDto
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetFeatureListResultDto) GetGroupsOk() (*[]FeatureGroupDto, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GetFeatureListResultDto) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []FeatureGroupDto and assigns it to the Groups field.
func (o *GetFeatureListResultDto) SetGroups(v []FeatureGroupDto) {
	o.Groups = v
}

func (o GetFeatureListResultDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableGetFeatureListResultDto struct {
	value *GetFeatureListResultDto
	isSet bool
}

func (v NullableGetFeatureListResultDto) Get() *GetFeatureListResultDto {
	return v.value
}

func (v *NullableGetFeatureListResultDto) Set(val *GetFeatureListResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeatureListResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeatureListResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeatureListResultDto(val *GetFeatureListResultDto) *NullableGetFeatureListResultDto {
	return &NullableGetFeatureListResultDto{value: val, isSet: true}
}

func (v NullableGetFeatureListResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFeatureListResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



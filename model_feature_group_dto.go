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

// checks if the FeatureGroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureGroupDto{}

// FeatureGroupDto struct for FeatureGroupDto
type FeatureGroupDto struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Features []FeatureDto `json:"features,omitempty"`
}

// NewFeatureGroupDto instantiates a new FeatureGroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureGroupDto() *FeatureGroupDto {
	this := FeatureGroupDto{}
	return &this
}

// NewFeatureGroupDtoWithDefaults instantiates a new FeatureGroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureGroupDtoWithDefaults() *FeatureGroupDto {
	this := FeatureGroupDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeatureGroupDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureGroupDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeatureGroupDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeatureGroupDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *FeatureGroupDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureGroupDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FeatureGroupDto) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FeatureGroupDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FeatureGroupDto) GetFeatures() []FeatureDto {
	if o == nil || IsNil(o.Features) {
		var ret []FeatureDto
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureGroupDto) GetFeaturesOk() ([]FeatureDto, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FeatureGroupDto) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []FeatureDto and assigns it to the Features field.
func (o *FeatureGroupDto) SetFeatures(v []FeatureDto) {
	o.Features = v
}

func (o FeatureGroupDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureGroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return toSerialize, nil
}

type NullableFeatureGroupDto struct {
	value *FeatureGroupDto
	isSet bool
}

func (v NullableFeatureGroupDto) Get() *FeatureGroupDto {
	return v.value
}

func (v *NullableFeatureGroupDto) Set(val *FeatureGroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureGroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureGroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureGroupDto(val *FeatureGroupDto) *NullableFeatureGroupDto {
	return &NullableFeatureGroupDto{value: val, isSet: true}
}

func (v NullableFeatureGroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureGroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



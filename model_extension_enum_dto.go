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

// ExtensionEnumDto struct for ExtensionEnumDto
type ExtensionEnumDto struct {
	Fields []ExtensionEnumFieldDto `json:"fields,omitempty"`
	LocalizationResource *string `json:"localizationResource,omitempty"`
}

// NewExtensionEnumDto instantiates a new ExtensionEnumDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionEnumDto() *ExtensionEnumDto {
	this := ExtensionEnumDto{}
	return &this
}

// NewExtensionEnumDtoWithDefaults instantiates a new ExtensionEnumDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionEnumDtoWithDefaults() *ExtensionEnumDto {
	this := ExtensionEnumDto{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ExtensionEnumDto) GetFields() []ExtensionEnumFieldDto {
	if o == nil || isNil(o.Fields) {
		var ret []ExtensionEnumFieldDto
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionEnumDto) GetFieldsOk() ([]ExtensionEnumFieldDto, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ExtensionEnumDto) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []ExtensionEnumFieldDto and assigns it to the Fields field.
func (o *ExtensionEnumDto) SetFields(v []ExtensionEnumFieldDto) {
	o.Fields = v
}

// GetLocalizationResource returns the LocalizationResource field value if set, zero value otherwise.
func (o *ExtensionEnumDto) GetLocalizationResource() string {
	if o == nil || isNil(o.LocalizationResource) {
		var ret string
		return ret
	}
	return *o.LocalizationResource
}

// GetLocalizationResourceOk returns a tuple with the LocalizationResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionEnumDto) GetLocalizationResourceOk() (*string, bool) {
	if o == nil || isNil(o.LocalizationResource) {
    return nil, false
	}
	return o.LocalizationResource, true
}

// HasLocalizationResource returns a boolean if a field has been set.
func (o *ExtensionEnumDto) HasLocalizationResource() bool {
	if o != nil && !isNil(o.LocalizationResource) {
		return true
	}

	return false
}

// SetLocalizationResource gets a reference to the given string and assigns it to the LocalizationResource field.
func (o *ExtensionEnumDto) SetLocalizationResource(v string) {
	o.LocalizationResource = &v
}

func (o ExtensionEnumDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.LocalizationResource) {
		toSerialize["localizationResource"] = o.LocalizationResource
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionEnumDto struct {
	value *ExtensionEnumDto
	isSet bool
}

func (v NullableExtensionEnumDto) Get() *ExtensionEnumDto {
	return v.value
}

func (v *NullableExtensionEnumDto) Set(val *ExtensionEnumDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionEnumDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionEnumDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionEnumDto(val *ExtensionEnumDto) *NullableExtensionEnumDto {
	return &NullableExtensionEnumDto{value: val, isSet: true}
}

func (v NullableExtensionEnumDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionEnumDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// IStringValueType struct for IStringValueType
type IStringValueType struct {
	Name *string `json:"name,omitempty"`
	Properties map[string]map[string]interface{} `json:"properties,omitempty"`
	Validator *IValueValidator `json:"validator,omitempty"`
}

// NewIStringValueType instantiates a new IStringValueType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIStringValueType() *IStringValueType {
	this := IStringValueType{}
	return &this
}

// NewIStringValueTypeWithDefaults instantiates a new IStringValueType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIStringValueTypeWithDefaults() *IStringValueType {
	this := IStringValueType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IStringValueType) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IStringValueType) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IStringValueType) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IStringValueType) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *IStringValueType) GetProperties() map[string]map[string]interface{} {
	if o == nil || isNil(o.Properties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IStringValueType) GetPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.Properties) {
    return map[string]map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *IStringValueType) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the Properties field.
func (o *IStringValueType) SetProperties(v map[string]map[string]interface{}) {
	o.Properties = v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *IStringValueType) GetValidator() IValueValidator {
	if o == nil || isNil(o.Validator) {
		var ret IValueValidator
		return ret
	}
	return *o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IStringValueType) GetValidatorOk() (*IValueValidator, bool) {
	if o == nil || isNil(o.Validator) {
    return nil, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *IStringValueType) HasValidator() bool {
	if o != nil && !isNil(o.Validator) {
		return true
	}

	return false
}

// SetValidator gets a reference to the given IValueValidator and assigns it to the Validator field.
func (o *IStringValueType) SetValidator(v IValueValidator) {
	o.Validator = &v
}

func (o IStringValueType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !isNil(o.Validator) {
		toSerialize["validator"] = o.Validator
	}
	return json.Marshal(toSerialize)
}

type NullableIStringValueType struct {
	value *IStringValueType
	isSet bool
}

func (v NullableIStringValueType) Get() *IStringValueType {
	return v.value
}

func (v *NullableIStringValueType) Set(val *IStringValueType) {
	v.value = val
	v.isSet = true
}

func (v NullableIStringValueType) IsSet() bool {
	return v.isSet
}

func (v *NullableIStringValueType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIStringValueType(val *IStringValueType) *NullableIStringValueType {
	return &NullableIStringValueType{value: val, isSet: true}
}

func (v NullableIStringValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIStringValueType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



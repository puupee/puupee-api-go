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

// checks if the ChangePasswordInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangePasswordInput{}

// ChangePasswordInput struct for ChangePasswordInput
type ChangePasswordInput struct {
	CurrentPassword *string `json:"currentPassword,omitempty"`
	NewPassword string `json:"newPassword"`
}

type _ChangePasswordInput ChangePasswordInput

// NewChangePasswordInput instantiates a new ChangePasswordInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangePasswordInput(newPassword string) *ChangePasswordInput {
	this := ChangePasswordInput{}
	this.NewPassword = newPassword
	return &this
}

// NewChangePasswordInputWithDefaults instantiates a new ChangePasswordInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangePasswordInputWithDefaults() *ChangePasswordInput {
	this := ChangePasswordInput{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *ChangePasswordInput) GetCurrentPassword() string {
	if o == nil || IsNil(o.CurrentPassword) {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangePasswordInput) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentPassword) {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *ChangePasswordInput) HasCurrentPassword() bool {
	if o != nil && !IsNil(o.CurrentPassword) {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *ChangePasswordInput) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetNewPassword returns the NewPassword field value
func (o *ChangePasswordInput) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *ChangePasswordInput) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *ChangePasswordInput) SetNewPassword(v string) {
	o.NewPassword = v
}

func (o ChangePasswordInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangePasswordInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentPassword) {
		toSerialize["currentPassword"] = o.CurrentPassword
	}
	toSerialize["newPassword"] = o.NewPassword
	return toSerialize, nil
}

func (o *ChangePasswordInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newPassword",
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

	varChangePasswordInput := _ChangePasswordInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChangePasswordInput)

	if err != nil {
		return err
	}

	*o = ChangePasswordInput(varChangePasswordInput)

	return err
}

type NullableChangePasswordInput struct {
	value *ChangePasswordInput
	isSet bool
}

func (v NullableChangePasswordInput) Get() *ChangePasswordInput {
	return v.value
}

func (v *NullableChangePasswordInput) Set(val *ChangePasswordInput) {
	v.value = val
	v.isSet = true
}

func (v NullableChangePasswordInput) IsSet() bool {
	return v.isSet
}

func (v *NullableChangePasswordInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangePasswordInput(val *ChangePasswordInput) *NullableChangePasswordInput {
	return &NullableChangePasswordInput{value: val, isSet: true}
}

func (v NullableChangePasswordInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangePasswordInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



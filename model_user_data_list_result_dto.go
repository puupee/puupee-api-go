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

// checks if the UserDataListResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataListResultDto{}

// UserDataListResultDto struct for UserDataListResultDto
type UserDataListResultDto struct {
	Items []UserData `json:"items,omitempty"`
}

// NewUserDataListResultDto instantiates a new UserDataListResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataListResultDto() *UserDataListResultDto {
	this := UserDataListResultDto{}
	return &this
}

// NewUserDataListResultDtoWithDefaults instantiates a new UserDataListResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataListResultDtoWithDefaults() *UserDataListResultDto {
	this := UserDataListResultDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UserDataListResultDto) GetItems() []UserData {
	if o == nil || IsNil(o.Items) {
		var ret []UserData
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataListResultDto) GetItemsOk() ([]UserData, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UserDataListResultDto) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UserData and assigns it to the Items field.
func (o *UserDataListResultDto) SetItems(v []UserData) {
	o.Items = v
}

func (o UserDataListResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataListResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableUserDataListResultDto struct {
	value *UserDataListResultDto
	isSet bool
}

func (v NullableUserDataListResultDto) Get() *UserDataListResultDto {
	return v.value
}

func (v *NullableUserDataListResultDto) Set(val *UserDataListResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataListResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataListResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataListResultDto(val *UserDataListResultDto) *NullableUserDataListResultDto {
	return &NullableUserDataListResultDto{value: val, isSet: true}
}

func (v NullableUserDataListResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataListResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the RemoteServiceValidationErrorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteServiceValidationErrorInfo{}

// RemoteServiceValidationErrorInfo struct for RemoteServiceValidationErrorInfo
type RemoteServiceValidationErrorInfo struct {
	Message *string `json:"message,omitempty"`
	Members []string `json:"members,omitempty"`
}

// NewRemoteServiceValidationErrorInfo instantiates a new RemoteServiceValidationErrorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteServiceValidationErrorInfo() *RemoteServiceValidationErrorInfo {
	this := RemoteServiceValidationErrorInfo{}
	return &this
}

// NewRemoteServiceValidationErrorInfoWithDefaults instantiates a new RemoteServiceValidationErrorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteServiceValidationErrorInfoWithDefaults() *RemoteServiceValidationErrorInfo {
	this := RemoteServiceValidationErrorInfo{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RemoteServiceValidationErrorInfo) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteServiceValidationErrorInfo) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RemoteServiceValidationErrorInfo) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RemoteServiceValidationErrorInfo) SetMessage(v string) {
	o.Message = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *RemoteServiceValidationErrorInfo) GetMembers() []string {
	if o == nil || IsNil(o.Members) {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteServiceValidationErrorInfo) GetMembersOk() ([]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *RemoteServiceValidationErrorInfo) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *RemoteServiceValidationErrorInfo) SetMembers(v []string) {
	o.Members = v
}

func (o RemoteServiceValidationErrorInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteServiceValidationErrorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	return toSerialize, nil
}

type NullableRemoteServiceValidationErrorInfo struct {
	value *RemoteServiceValidationErrorInfo
	isSet bool
}

func (v NullableRemoteServiceValidationErrorInfo) Get() *RemoteServiceValidationErrorInfo {
	return v.value
}

func (v *NullableRemoteServiceValidationErrorInfo) Set(val *RemoteServiceValidationErrorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteServiceValidationErrorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteServiceValidationErrorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteServiceValidationErrorInfo(val *RemoteServiceValidationErrorInfo) *NullableRemoteServiceValidationErrorInfo {
	return &NullableRemoteServiceValidationErrorInfo{value: val, isSet: true}
}

func (v NullableRemoteServiceValidationErrorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteServiceValidationErrorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



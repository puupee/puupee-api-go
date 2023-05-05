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

// checks if the MessageSubscribeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageSubscribeDto{}

// MessageSubscribeDto struct for MessageSubscribeDto
type MessageSubscribeDto struct {
	AppId *string `json:"appId,omitempty"`
}

// NewMessageSubscribeDto instantiates a new MessageSubscribeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSubscribeDto() *MessageSubscribeDto {
	this := MessageSubscribeDto{}
	return &this
}

// NewMessageSubscribeDtoWithDefaults instantiates a new MessageSubscribeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageSubscribeDtoWithDefaults() *MessageSubscribeDto {
	this := MessageSubscribeDto{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *MessageSubscribeDto) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSubscribeDto) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *MessageSubscribeDto) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *MessageSubscribeDto) SetAppId(v string) {
	o.AppId = &v
}

func (o MessageSubscribeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageSubscribeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	return toSerialize, nil
}

type NullableMessageSubscribeDto struct {
	value *MessageSubscribeDto
	isSet bool
}

func (v NullableMessageSubscribeDto) Get() *MessageSubscribeDto {
	return v.value
}

func (v *NullableMessageSubscribeDto) Set(val *MessageSubscribeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSubscribeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSubscribeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSubscribeDto(val *MessageSubscribeDto) *NullableMessageSubscribeDto {
	return &NullableMessageSubscribeDto{value: val, isSet: true}
}

func (v NullableMessageSubscribeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSubscribeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



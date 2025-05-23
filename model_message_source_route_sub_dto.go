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

// checks if the MessageSourceRouteSubDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageSourceRouteSubDto{}

// MessageSourceRouteSubDto struct for MessageSourceRouteSubDto
type MessageSourceRouteSubDto struct {
	RouteId *string `json:"routeId,omitempty"`
	Path *string `json:"path,omitempty"`
	Values map[string]interface{} `json:"values,omitempty"`
}

// NewMessageSourceRouteSubDto instantiates a new MessageSourceRouteSubDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSourceRouteSubDto() *MessageSourceRouteSubDto {
	this := MessageSourceRouteSubDto{}
	return &this
}

// NewMessageSourceRouteSubDtoWithDefaults instantiates a new MessageSourceRouteSubDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageSourceRouteSubDtoWithDefaults() *MessageSourceRouteSubDto {
	this := MessageSourceRouteSubDto{}
	return &this
}

// GetRouteId returns the RouteId field value if set, zero value otherwise.
func (o *MessageSourceRouteSubDto) GetRouteId() string {
	if o == nil || IsNil(o.RouteId) {
		var ret string
		return ret
	}
	return *o.RouteId
}

// GetRouteIdOk returns a tuple with the RouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSourceRouteSubDto) GetRouteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RouteId) {
		return nil, false
	}
	return o.RouteId, true
}

// HasRouteId returns a boolean if a field has been set.
func (o *MessageSourceRouteSubDto) HasRouteId() bool {
	if o != nil && !IsNil(o.RouteId) {
		return true
	}

	return false
}

// SetRouteId gets a reference to the given string and assigns it to the RouteId field.
func (o *MessageSourceRouteSubDto) SetRouteId(v string) {
	o.RouteId = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *MessageSourceRouteSubDto) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSourceRouteSubDto) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *MessageSourceRouteSubDto) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *MessageSourceRouteSubDto) SetPath(v string) {
	o.Path = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MessageSourceRouteSubDto) GetValues() map[string]interface{} {
	if o == nil || IsNil(o.Values) {
		var ret map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageSourceRouteSubDto) GetValuesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Values) {
		return map[string]interface{}{}, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MessageSourceRouteSubDto) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *MessageSourceRouteSubDto) SetValues(v map[string]interface{}) {
	o.Values = v
}

func (o MessageSourceRouteSubDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageSourceRouteSubDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RouteId) {
		toSerialize["routeId"] = o.RouteId
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableMessageSourceRouteSubDto struct {
	value *MessageSourceRouteSubDto
	isSet bool
}

func (v NullableMessageSourceRouteSubDto) Get() *MessageSourceRouteSubDto {
	return v.value
}

func (v *NullableMessageSourceRouteSubDto) Set(val *MessageSourceRouteSubDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSourceRouteSubDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSourceRouteSubDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSourceRouteSubDto(val *MessageSourceRouteSubDto) *NullableMessageSourceRouteSubDto {
	return &NullableMessageSourceRouteSubDto{value: val, isSet: true}
}

func (v NullableMessageSourceRouteSubDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSourceRouteSubDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



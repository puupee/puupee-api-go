/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
)

// PuupeeChangedEto struct for PuupeeChangedEto
type PuupeeChangedEto struct {
	Data *PuupeeDto `json:"data,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewPuupeeChangedEto instantiates a new PuupeeChangedEto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPuupeeChangedEto() *PuupeeChangedEto {
	this := PuupeeChangedEto{}
	return &this
}

// NewPuupeeChangedEtoWithDefaults instantiates a new PuupeeChangedEto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPuupeeChangedEtoWithDefaults() *PuupeeChangedEto {
	this := PuupeeChangedEto{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PuupeeChangedEto) GetData() PuupeeDto {
	if o == nil || isNil(o.Data) {
		var ret PuupeeDto
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PuupeeChangedEto) GetDataOk() (*PuupeeDto, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PuupeeChangedEto) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PuupeeDto and assigns it to the Data field.
func (o *PuupeeChangedEto) SetData(v PuupeeDto) {
	o.Data = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PuupeeChangedEto) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PuupeeChangedEto) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
    return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PuupeeChangedEto) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PuupeeChangedEto) SetUserId(v string) {
	o.UserId = &v
}

func (o PuupeeChangedEto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullablePuupeeChangedEto struct {
	value *PuupeeChangedEto
	isSet bool
}

func (v NullablePuupeeChangedEto) Get() *PuupeeChangedEto {
	return v.value
}

func (v *NullablePuupeeChangedEto) Set(val *PuupeeChangedEto) {
	v.value = val
	v.isSet = true
}

func (v NullablePuupeeChangedEto) IsSet() bool {
	return v.isSet
}

func (v *NullablePuupeeChangedEto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePuupeeChangedEto(val *PuupeeChangedEto) *NullablePuupeeChangedEto {
	return &NullablePuupeeChangedEto{value: val, isSet: true}
}

func (v NullablePuupeeChangedEto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePuupeeChangedEto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



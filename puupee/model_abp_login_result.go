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

// AbpLoginResult struct for AbpLoginResult
type AbpLoginResult struct {
	Result *LoginResultType `json:"result,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewAbpLoginResult instantiates a new AbpLoginResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbpLoginResult() *AbpLoginResult {
	this := AbpLoginResult{}
	return &this
}

// NewAbpLoginResultWithDefaults instantiates a new AbpLoginResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbpLoginResultWithDefaults() *AbpLoginResult {
	this := AbpLoginResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AbpLoginResult) GetResult() LoginResultType {
	if o == nil || isNil(o.Result) {
		var ret LoginResultType
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbpLoginResult) GetResultOk() (*LoginResultType, bool) {
	if o == nil || isNil(o.Result) {
    return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AbpLoginResult) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given LoginResultType and assigns it to the Result field.
func (o *AbpLoginResult) SetResult(v LoginResultType) {
	o.Result = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AbpLoginResult) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbpLoginResult) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AbpLoginResult) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AbpLoginResult) SetDescription(v string) {
	o.Description = &v
}

func (o AbpLoginResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAbpLoginResult struct {
	value *AbpLoginResult
	isSet bool
}

func (v NullableAbpLoginResult) Get() *AbpLoginResult {
	return v.value
}

func (v *NullableAbpLoginResult) Set(val *AbpLoginResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAbpLoginResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAbpLoginResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbpLoginResult(val *AbpLoginResult) *NullableAbpLoginResult {
	return &NullableAbpLoginResult{value: val, isSet: true}
}

func (v NullableAbpLoginResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbpLoginResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// VerifyReceiptResult struct for VerifyReceiptResult
type VerifyReceiptResult struct {
	Ok *bool `json:"ok,omitempty"`
	StatusCode *string `json:"statusCode,omitempty"`
	Message *string `json:"message,omitempty"`
	AppleVerifyReceiptResult *AppleVerifyReceiptResult `json:"appleVerifyReceiptResult,omitempty"`
}

// NewVerifyReceiptResult instantiates a new VerifyReceiptResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyReceiptResult() *VerifyReceiptResult {
	this := VerifyReceiptResult{}
	return &this
}

// NewVerifyReceiptResultWithDefaults instantiates a new VerifyReceiptResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyReceiptResultWithDefaults() *VerifyReceiptResult {
	this := VerifyReceiptResult{}
	return &this
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *VerifyReceiptResult) GetOk() bool {
	if o == nil || isNil(o.Ok) {
		var ret bool
		return ret
	}
	return *o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyReceiptResult) GetOkOk() (*bool, bool) {
	if o == nil || isNil(o.Ok) {
    return nil, false
	}
	return o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *VerifyReceiptResult) HasOk() bool {
	if o != nil && !isNil(o.Ok) {
		return true
	}

	return false
}

// SetOk gets a reference to the given bool and assigns it to the Ok field.
func (o *VerifyReceiptResult) SetOk(v bool) {
	o.Ok = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *VerifyReceiptResult) GetStatusCode() string {
	if o == nil || isNil(o.StatusCode) {
		var ret string
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyReceiptResult) GetStatusCodeOk() (*string, bool) {
	if o == nil || isNil(o.StatusCode) {
    return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *VerifyReceiptResult) HasStatusCode() bool {
	if o != nil && !isNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given string and assigns it to the StatusCode field.
func (o *VerifyReceiptResult) SetStatusCode(v string) {
	o.StatusCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *VerifyReceiptResult) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyReceiptResult) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *VerifyReceiptResult) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *VerifyReceiptResult) SetMessage(v string) {
	o.Message = &v
}

// GetAppleVerifyReceiptResult returns the AppleVerifyReceiptResult field value if set, zero value otherwise.
func (o *VerifyReceiptResult) GetAppleVerifyReceiptResult() AppleVerifyReceiptResult {
	if o == nil || isNil(o.AppleVerifyReceiptResult) {
		var ret AppleVerifyReceiptResult
		return ret
	}
	return *o.AppleVerifyReceiptResult
}

// GetAppleVerifyReceiptResultOk returns a tuple with the AppleVerifyReceiptResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyReceiptResult) GetAppleVerifyReceiptResultOk() (*AppleVerifyReceiptResult, bool) {
	if o == nil || isNil(o.AppleVerifyReceiptResult) {
    return nil, false
	}
	return o.AppleVerifyReceiptResult, true
}

// HasAppleVerifyReceiptResult returns a boolean if a field has been set.
func (o *VerifyReceiptResult) HasAppleVerifyReceiptResult() bool {
	if o != nil && !isNil(o.AppleVerifyReceiptResult) {
		return true
	}

	return false
}

// SetAppleVerifyReceiptResult gets a reference to the given AppleVerifyReceiptResult and assigns it to the AppleVerifyReceiptResult field.
func (o *VerifyReceiptResult) SetAppleVerifyReceiptResult(v AppleVerifyReceiptResult) {
	o.AppleVerifyReceiptResult = &v
}

func (o VerifyReceiptResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ok) {
		toSerialize["ok"] = o.Ok
	}
	if !isNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.AppleVerifyReceiptResult) {
		toSerialize["appleVerifyReceiptResult"] = o.AppleVerifyReceiptResult
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyReceiptResult struct {
	value *VerifyReceiptResult
	isSet bool
}

func (v NullableVerifyReceiptResult) Get() *VerifyReceiptResult {
	return v.value
}

func (v *NullableVerifyReceiptResult) Set(val *VerifyReceiptResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyReceiptResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyReceiptResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyReceiptResult(val *VerifyReceiptResult) *NullableVerifyReceiptResult {
	return &NullableVerifyReceiptResult{value: val, isSet: true}
}

func (v NullableVerifyReceiptResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyReceiptResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


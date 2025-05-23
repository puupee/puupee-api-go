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

// checks if the SendVerificationCodeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendVerificationCodeDto{}

// SendVerificationCodeDto struct for SendVerificationCodeDto
type SendVerificationCodeDto struct {
	// 验证码发送器 暂时支持: SMS: 手机短信验证码
	CodeSender *string `json:"codeSender,omitempty"`
	// 验证码接受者, 用户账户
	Account *string `json:"account,omitempty"`
	// 验证码用途
	Purpose *string `json:"purpose,omitempty"`
}

// NewSendVerificationCodeDto instantiates a new SendVerificationCodeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendVerificationCodeDto() *SendVerificationCodeDto {
	this := SendVerificationCodeDto{}
	return &this
}

// NewSendVerificationCodeDtoWithDefaults instantiates a new SendVerificationCodeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendVerificationCodeDtoWithDefaults() *SendVerificationCodeDto {
	this := SendVerificationCodeDto{}
	return &this
}

// GetCodeSender returns the CodeSender field value if set, zero value otherwise.
func (o *SendVerificationCodeDto) GetCodeSender() string {
	if o == nil || IsNil(o.CodeSender) {
		var ret string
		return ret
	}
	return *o.CodeSender
}

// GetCodeSenderOk returns a tuple with the CodeSender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVerificationCodeDto) GetCodeSenderOk() (*string, bool) {
	if o == nil || IsNil(o.CodeSender) {
		return nil, false
	}
	return o.CodeSender, true
}

// HasCodeSender returns a boolean if a field has been set.
func (o *SendVerificationCodeDto) HasCodeSender() bool {
	if o != nil && !IsNil(o.CodeSender) {
		return true
	}

	return false
}

// SetCodeSender gets a reference to the given string and assigns it to the CodeSender field.
func (o *SendVerificationCodeDto) SetCodeSender(v string) {
	o.CodeSender = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SendVerificationCodeDto) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVerificationCodeDto) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SendVerificationCodeDto) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *SendVerificationCodeDto) SetAccount(v string) {
	o.Account = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *SendVerificationCodeDto) GetPurpose() string {
	if o == nil || IsNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendVerificationCodeDto) GetPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *SendVerificationCodeDto) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *SendVerificationCodeDto) SetPurpose(v string) {
	o.Purpose = &v
}

func (o SendVerificationCodeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendVerificationCodeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodeSender) {
		toSerialize["codeSender"] = o.CodeSender
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	return toSerialize, nil
}

type NullableSendVerificationCodeDto struct {
	value *SendVerificationCodeDto
	isSet bool
}

func (v NullableSendVerificationCodeDto) Get() *SendVerificationCodeDto {
	return v.value
}

func (v *NullableSendVerificationCodeDto) Set(val *SendVerificationCodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSendVerificationCodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSendVerificationCodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendVerificationCodeDto(val *SendVerificationCodeDto) *NullableSendVerificationCodeDto {
	return &NullableSendVerificationCodeDto{value: val, isSet: true}
}

func (v NullableSendVerificationCodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendVerificationCodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the SendVerificationCodeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendVerificationCodeDto{}

// SendVerificationCodeDto struct for SendVerificationCodeDto
type SendVerificationCodeDto struct {
	CodeSender NullableString `json:"codeSender,omitempty"`
	Account NullableString `json:"account,omitempty"`
	Purpose NullableString `json:"purpose,omitempty"`
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

// GetCodeSender returns the CodeSender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendVerificationCodeDto) GetCodeSender() string {
	if o == nil || IsNil(o.CodeSender.Get()) {
		var ret string
		return ret
	}
	return *o.CodeSender.Get()
}

// GetCodeSenderOk returns a tuple with the CodeSender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendVerificationCodeDto) GetCodeSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeSender.Get(), o.CodeSender.IsSet()
}

// HasCodeSender returns a boolean if a field has been set.
func (o *SendVerificationCodeDto) HasCodeSender() bool {
	if o != nil && o.CodeSender.IsSet() {
		return true
	}

	return false
}

// SetCodeSender gets a reference to the given NullableString and assigns it to the CodeSender field.
func (o *SendVerificationCodeDto) SetCodeSender(v string) {
	o.CodeSender.Set(&v)
}
// SetCodeSenderNil sets the value for CodeSender to be an explicit nil
func (o *SendVerificationCodeDto) SetCodeSenderNil() {
	o.CodeSender.Set(nil)
}

// UnsetCodeSender ensures that no value is present for CodeSender, not even an explicit nil
func (o *SendVerificationCodeDto) UnsetCodeSender() {
	o.CodeSender.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendVerificationCodeDto) GetAccount() string {
	if o == nil || IsNil(o.Account.Get()) {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendVerificationCodeDto) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *SendVerificationCodeDto) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *SendVerificationCodeDto) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *SendVerificationCodeDto) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *SendVerificationCodeDto) UnsetAccount() {
	o.Account.Unset()
}

// GetPurpose returns the Purpose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendVerificationCodeDto) GetPurpose() string {
	if o == nil || IsNil(o.Purpose.Get()) {
		var ret string
		return ret
	}
	return *o.Purpose.Get()
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendVerificationCodeDto) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Purpose.Get(), o.Purpose.IsSet()
}

// HasPurpose returns a boolean if a field has been set.
func (o *SendVerificationCodeDto) HasPurpose() bool {
	if o != nil && o.Purpose.IsSet() {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given NullableString and assigns it to the Purpose field.
func (o *SendVerificationCodeDto) SetPurpose(v string) {
	o.Purpose.Set(&v)
}
// SetPurposeNil sets the value for Purpose to be an explicit nil
func (o *SendVerificationCodeDto) SetPurposeNil() {
	o.Purpose.Set(nil)
}

// UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
func (o *SendVerificationCodeDto) UnsetPurpose() {
	o.Purpose.Unset()
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
	if o.CodeSender.IsSet() {
		toSerialize["codeSender"] = o.CodeSender.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Purpose.IsSet() {
		toSerialize["purpose"] = o.Purpose.Get()
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



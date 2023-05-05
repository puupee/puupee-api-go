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

// checks if the SendPasswordResetCodeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendPasswordResetCodeDto{}

// SendPasswordResetCodeDto struct for SendPasswordResetCodeDto
type SendPasswordResetCodeDto struct {
	Email string `json:"email"`
	AppName string `json:"appName"`
	ReturnUrl *string `json:"returnUrl,omitempty"`
	ReturnUrlHash *string `json:"returnUrlHash,omitempty"`
}

// NewSendPasswordResetCodeDto instantiates a new SendPasswordResetCodeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendPasswordResetCodeDto(email string, appName string) *SendPasswordResetCodeDto {
	this := SendPasswordResetCodeDto{}
	this.Email = email
	this.AppName = appName
	return &this
}

// NewSendPasswordResetCodeDtoWithDefaults instantiates a new SendPasswordResetCodeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendPasswordResetCodeDtoWithDefaults() *SendPasswordResetCodeDto {
	this := SendPasswordResetCodeDto{}
	return &this
}

// GetEmail returns the Email field value
func (o *SendPasswordResetCodeDto) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SendPasswordResetCodeDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SendPasswordResetCodeDto) SetEmail(v string) {
	o.Email = v
}

// GetAppName returns the AppName field value
func (o *SendPasswordResetCodeDto) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *SendPasswordResetCodeDto) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *SendPasswordResetCodeDto) SetAppName(v string) {
	o.AppName = v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *SendPasswordResetCodeDto) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPasswordResetCodeDto) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *SendPasswordResetCodeDto) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *SendPasswordResetCodeDto) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetReturnUrlHash returns the ReturnUrlHash field value if set, zero value otherwise.
func (o *SendPasswordResetCodeDto) GetReturnUrlHash() string {
	if o == nil || IsNil(o.ReturnUrlHash) {
		var ret string
		return ret
	}
	return *o.ReturnUrlHash
}

// GetReturnUrlHashOk returns a tuple with the ReturnUrlHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendPasswordResetCodeDto) GetReturnUrlHashOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrlHash) {
		return nil, false
	}
	return o.ReturnUrlHash, true
}

// HasReturnUrlHash returns a boolean if a field has been set.
func (o *SendPasswordResetCodeDto) HasReturnUrlHash() bool {
	if o != nil && !IsNil(o.ReturnUrlHash) {
		return true
	}

	return false
}

// SetReturnUrlHash gets a reference to the given string and assigns it to the ReturnUrlHash field.
func (o *SendPasswordResetCodeDto) SetReturnUrlHash(v string) {
	o.ReturnUrlHash = &v
}

func (o SendPasswordResetCodeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendPasswordResetCodeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["appName"] = o.AppName
	if !IsNil(o.ReturnUrl) {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if !IsNil(o.ReturnUrlHash) {
		toSerialize["returnUrlHash"] = o.ReturnUrlHash
	}
	return toSerialize, nil
}

type NullableSendPasswordResetCodeDto struct {
	value *SendPasswordResetCodeDto
	isSet bool
}

func (v NullableSendPasswordResetCodeDto) Get() *SendPasswordResetCodeDto {
	return v.value
}

func (v *NullableSendPasswordResetCodeDto) Set(val *SendPasswordResetCodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSendPasswordResetCodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSendPasswordResetCodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendPasswordResetCodeDto(val *SendPasswordResetCodeDto) *NullableSendPasswordResetCodeDto {
	return &NullableSendPasswordResetCodeDto{value: val, isSet: true}
}

func (v NullableSendPasswordResetCodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendPasswordResetCodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



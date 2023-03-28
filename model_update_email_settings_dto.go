/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// UpdateEmailSettingsDto struct for UpdateEmailSettingsDto
type UpdateEmailSettingsDto struct {
	SmtpHost *string `json:"smtpHost,omitempty"`
	SmtpPort *int32 `json:"smtpPort,omitempty"`
	SmtpUserName *string `json:"smtpUserName,omitempty"`
	SmtpPassword *string `json:"smtpPassword,omitempty"`
	SmtpDomain *string `json:"smtpDomain,omitempty"`
	SmtpEnableSsl *bool `json:"smtpEnableSsl,omitempty"`
	SmtpUseDefaultCredentials *bool `json:"smtpUseDefaultCredentials,omitempty"`
	DefaultFromAddress string `json:"defaultFromAddress"`
	DefaultFromDisplayName string `json:"defaultFromDisplayName"`
}

// NewUpdateEmailSettingsDto instantiates a new UpdateEmailSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEmailSettingsDto(defaultFromAddress string, defaultFromDisplayName string) *UpdateEmailSettingsDto {
	this := UpdateEmailSettingsDto{}
	this.DefaultFromAddress = defaultFromAddress
	this.DefaultFromDisplayName = defaultFromDisplayName
	return &this
}

// NewUpdateEmailSettingsDtoWithDefaults instantiates a new UpdateEmailSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEmailSettingsDtoWithDefaults() *UpdateEmailSettingsDto {
	this := UpdateEmailSettingsDto{}
	return &this
}

// GetSmtpHost returns the SmtpHost field value if set, zero value otherwise.
func (o *UpdateEmailSettingsDto) GetSmtpHost() string {
	if o == nil || isNil(o.SmtpHost) {
		var ret string
		return ret
	}
	return *o.SmtpHost
}

// GetSmtpHostOk returns a tuple with the SmtpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmailSettingsDto) GetSmtpHostOk() (*string, bool) {
	if o == nil || isNil(o.SmtpHost) {
    return nil, false
	}
	return o.SmtpHost, true
}

// HasSmtpHost returns a boolean if a field has been set.
func (o *UpdateEmailSettingsDto) HasSmtpHost() bool {
	if o != nil && !isNil(o.SmtpHost) {
		return true
	}

	return false
}

// SetSmtpHost gets a reference to the given string and assigns it to the SmtpHost field.
func (o *UpdateEmailSettingsDto) SetSmtpHost(v string) {
	o.SmtpHost = &v
}

// GetSmtpPort returns the SmtpPort field value if set, zero value otherwise.
func (o *UpdateEmailSettingsDto) GetSmtpPort() int32 {
	if o == nil || isNil(o.SmtpPort) {
		var ret int32
		return ret
	}
	return *o.SmtpPort
}

// GetSmtpPortOk returns a tuple with the SmtpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmailSettingsDto) GetSmtpPortOk() (*int32, bool) {
	if o == nil || isNil(o.SmtpPort) {
    return nil, false
	}
	return o.SmtpPort, true
}

// HasSmtpPort returns a boolean if a field has been set.
func (o *UpdateEmailSettingsDto) HasSmtpPort() bool {
	if o != nil && !isNil(o.SmtpPort) {
		return true
	}

	return false
}

// SetSmtpPort gets a reference to the given int32 and assigns it to the SmtpPort field.
func (o *UpdateEmailSettingsDto) SetSmtpPort(v int32) {
	o.SmtpPort = &v
}

// GetSmtpUserName returns the SmtpUserName field value if set, zero value otherwise.
func (o *UpdateEmailSettingsDto) GetSmtpUserName() string {
	if o == nil || isNil(o.SmtpUserName) {
		var ret string
		return ret
	}
	return *o.SmtpUserName
}

// GetSmtpUserNameOk returns a tuple with the SmtpUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmailSettingsDto) GetSmtpUserNameOk() (*string, bool) {
	if o == nil || isNil(o.SmtpUserName) {
    return nil, false
	}
	return o.SmtpUserName, true
}

// HasSmtpUserName returns a boolean if a field has been set.
func (o *UpdateEmailSettingsDto) HasSmtpUserName() bool {
	if o != nil && !isNil(o.SmtpUserName) {
		return true
	}

	return false
}

// SetSmtpUserName gets a reference to the given string and assigns it to the SmtpUserName field.
func (o *UpdateEmailSettingsDto) SetSmtpUserName(v string) {
	o.SmtpUserName = &v
}

// GetSmtpPassword returns the SmtpPassword field value if set, zero value otherwise.
func (o *UpdateEmailSettingsDto) GetSmtpPassword() string {
	if o == nil || isNil(o.SmtpPassword) {
		var ret string
		return ret
	}
	return *o.SmtpPassword
}

// GetSmtpPasswordOk returns a tuple with the SmtpPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmailSettingsDto) GetSmtpPasswordOk() (*string, bool) {
	if o == nil || isNil(o.SmtpPassword) {
    return nil, false
	}
	return o.SmtpPassword, true
}

// HasSmtpPassword returns a boolean if a field has been set.
func (o *UpdateEmailSettingsDto) HasSmtpPassword() bool {
	if o != nil && !isNil(o.SmtpPassword) {
		return true
	}

	return false
}

// SetSmtpPassword gets a reference to the given string and assigns it to the SmtpPassword field.
func (o *UpdateEmailSettingsDto) SetSmtpPassword(v string) {
	o.SmtpPassword = &v
}

// GetSmtpDomain returns the SmtpDomain field value if set, zero value otherwise.
func (o *UpdateEmailSettingsDto) GetSmtpDomain() string {
	if o == nil || isNil(o.SmtpDomain) {
		var ret string
		return ret
	}
	return *o.SmtpDomain
}

// GetSmtpDomainOk returns a tuple with the SmtpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmailSettingsDto) GetSmtpDomainOk() (*string, bool) {
	if o == nil || isNil(o.SmtpDomain) {
    return nil, false
	}
	return o.SmtpDomain, true
}

// HasSmtpDomain returns a boolean if a field has been set.
func (o *UpdateEmailSettingsDto) HasSmtpDomain() bool {
	if o != nil && !isNil(o.SmtpDomain) {
		return true
	}

	return false
}

// SetSmtpDomain gets a reference to the given string and assigns it to the SmtpDomain field.
func (o *UpdateEmailSettingsDto) SetSmtpDomain(v string) {
	o.SmtpDomain = &v
}

// GetSmtpEnableSsl returns the SmtpEnableSsl field value if set, zero value otherwise.
func (o *UpdateEmailSettingsDto) GetSmtpEnableSsl() bool {
	if o == nil || isNil(o.SmtpEnableSsl) {
		var ret bool
		return ret
	}
	return *o.SmtpEnableSsl
}

// GetSmtpEnableSslOk returns a tuple with the SmtpEnableSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmailSettingsDto) GetSmtpEnableSslOk() (*bool, bool) {
	if o == nil || isNil(o.SmtpEnableSsl) {
    return nil, false
	}
	return o.SmtpEnableSsl, true
}

// HasSmtpEnableSsl returns a boolean if a field has been set.
func (o *UpdateEmailSettingsDto) HasSmtpEnableSsl() bool {
	if o != nil && !isNil(o.SmtpEnableSsl) {
		return true
	}

	return false
}

// SetSmtpEnableSsl gets a reference to the given bool and assigns it to the SmtpEnableSsl field.
func (o *UpdateEmailSettingsDto) SetSmtpEnableSsl(v bool) {
	o.SmtpEnableSsl = &v
}

// GetSmtpUseDefaultCredentials returns the SmtpUseDefaultCredentials field value if set, zero value otherwise.
func (o *UpdateEmailSettingsDto) GetSmtpUseDefaultCredentials() bool {
	if o == nil || isNil(o.SmtpUseDefaultCredentials) {
		var ret bool
		return ret
	}
	return *o.SmtpUseDefaultCredentials
}

// GetSmtpUseDefaultCredentialsOk returns a tuple with the SmtpUseDefaultCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmailSettingsDto) GetSmtpUseDefaultCredentialsOk() (*bool, bool) {
	if o == nil || isNil(o.SmtpUseDefaultCredentials) {
    return nil, false
	}
	return o.SmtpUseDefaultCredentials, true
}

// HasSmtpUseDefaultCredentials returns a boolean if a field has been set.
func (o *UpdateEmailSettingsDto) HasSmtpUseDefaultCredentials() bool {
	if o != nil && !isNil(o.SmtpUseDefaultCredentials) {
		return true
	}

	return false
}

// SetSmtpUseDefaultCredentials gets a reference to the given bool and assigns it to the SmtpUseDefaultCredentials field.
func (o *UpdateEmailSettingsDto) SetSmtpUseDefaultCredentials(v bool) {
	o.SmtpUseDefaultCredentials = &v
}

// GetDefaultFromAddress returns the DefaultFromAddress field value
func (o *UpdateEmailSettingsDto) GetDefaultFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultFromAddress
}

// GetDefaultFromAddressOk returns a tuple with the DefaultFromAddress field value
// and a boolean to check if the value has been set.
func (o *UpdateEmailSettingsDto) GetDefaultFromAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DefaultFromAddress, true
}

// SetDefaultFromAddress sets field value
func (o *UpdateEmailSettingsDto) SetDefaultFromAddress(v string) {
	o.DefaultFromAddress = v
}

// GetDefaultFromDisplayName returns the DefaultFromDisplayName field value
func (o *UpdateEmailSettingsDto) GetDefaultFromDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultFromDisplayName
}

// GetDefaultFromDisplayNameOk returns a tuple with the DefaultFromDisplayName field value
// and a boolean to check if the value has been set.
func (o *UpdateEmailSettingsDto) GetDefaultFromDisplayNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DefaultFromDisplayName, true
}

// SetDefaultFromDisplayName sets field value
func (o *UpdateEmailSettingsDto) SetDefaultFromDisplayName(v string) {
	o.DefaultFromDisplayName = v
}

func (o UpdateEmailSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SmtpHost) {
		toSerialize["smtpHost"] = o.SmtpHost
	}
	if !isNil(o.SmtpPort) {
		toSerialize["smtpPort"] = o.SmtpPort
	}
	if !isNil(o.SmtpUserName) {
		toSerialize["smtpUserName"] = o.SmtpUserName
	}
	if !isNil(o.SmtpPassword) {
		toSerialize["smtpPassword"] = o.SmtpPassword
	}
	if !isNil(o.SmtpDomain) {
		toSerialize["smtpDomain"] = o.SmtpDomain
	}
	if !isNil(o.SmtpEnableSsl) {
		toSerialize["smtpEnableSsl"] = o.SmtpEnableSsl
	}
	if !isNil(o.SmtpUseDefaultCredentials) {
		toSerialize["smtpUseDefaultCredentials"] = o.SmtpUseDefaultCredentials
	}
	if true {
		toSerialize["defaultFromAddress"] = o.DefaultFromAddress
	}
	if true {
		toSerialize["defaultFromDisplayName"] = o.DefaultFromDisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateEmailSettingsDto struct {
	value *UpdateEmailSettingsDto
	isSet bool
}

func (v NullableUpdateEmailSettingsDto) Get() *UpdateEmailSettingsDto {
	return v.value
}

func (v *NullableUpdateEmailSettingsDto) Set(val *UpdateEmailSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEmailSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEmailSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEmailSettingsDto(val *UpdateEmailSettingsDto) *NullableUpdateEmailSettingsDto {
	return &NullableUpdateEmailSettingsDto{value: val, isSet: true}
}

func (v NullableUpdateEmailSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEmailSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



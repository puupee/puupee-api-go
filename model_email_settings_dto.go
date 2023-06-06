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

// checks if the EmailSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSettingsDto{}

// EmailSettingsDto struct for EmailSettingsDto
type EmailSettingsDto struct {
	SmtpHost *string `json:"smtpHost,omitempty"`
	SmtpPort *int32 `json:"smtpPort,omitempty"`
	SmtpUserName *string `json:"smtpUserName,omitempty"`
	SmtpPassword *string `json:"smtpPassword,omitempty"`
	SmtpDomain *string `json:"smtpDomain,omitempty"`
	SmtpEnableSsl *bool `json:"smtpEnableSsl,omitempty"`
	SmtpUseDefaultCredentials *bool `json:"smtpUseDefaultCredentials,omitempty"`
	DefaultFromAddress *string `json:"defaultFromAddress,omitempty"`
	DefaultFromDisplayName *string `json:"defaultFromDisplayName,omitempty"`
}

// NewEmailSettingsDto instantiates a new EmailSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSettingsDto() *EmailSettingsDto {
	this := EmailSettingsDto{}
	return &this
}

// NewEmailSettingsDtoWithDefaults instantiates a new EmailSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSettingsDtoWithDefaults() *EmailSettingsDto {
	this := EmailSettingsDto{}
	return &this
}

// GetSmtpHost returns the SmtpHost field value if set, zero value otherwise.
func (o *EmailSettingsDto) GetSmtpHost() string {
	if o == nil || IsNil(o.SmtpHost) {
		var ret string
		return ret
	}
	return *o.SmtpHost
}

// GetSmtpHostOk returns a tuple with the SmtpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsDto) GetSmtpHostOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpHost) {
		return nil, false
	}
	return o.SmtpHost, true
}

// HasSmtpHost returns a boolean if a field has been set.
func (o *EmailSettingsDto) HasSmtpHost() bool {
	if o != nil && !IsNil(o.SmtpHost) {
		return true
	}

	return false
}

// SetSmtpHost gets a reference to the given string and assigns it to the SmtpHost field.
func (o *EmailSettingsDto) SetSmtpHost(v string) {
	o.SmtpHost = &v
}

// GetSmtpPort returns the SmtpPort field value if set, zero value otherwise.
func (o *EmailSettingsDto) GetSmtpPort() int32 {
	if o == nil || IsNil(o.SmtpPort) {
		var ret int32
		return ret
	}
	return *o.SmtpPort
}

// GetSmtpPortOk returns a tuple with the SmtpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsDto) GetSmtpPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SmtpPort) {
		return nil, false
	}
	return o.SmtpPort, true
}

// HasSmtpPort returns a boolean if a field has been set.
func (o *EmailSettingsDto) HasSmtpPort() bool {
	if o != nil && !IsNil(o.SmtpPort) {
		return true
	}

	return false
}

// SetSmtpPort gets a reference to the given int32 and assigns it to the SmtpPort field.
func (o *EmailSettingsDto) SetSmtpPort(v int32) {
	o.SmtpPort = &v
}

// GetSmtpUserName returns the SmtpUserName field value if set, zero value otherwise.
func (o *EmailSettingsDto) GetSmtpUserName() string {
	if o == nil || IsNil(o.SmtpUserName) {
		var ret string
		return ret
	}
	return *o.SmtpUserName
}

// GetSmtpUserNameOk returns a tuple with the SmtpUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsDto) GetSmtpUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpUserName) {
		return nil, false
	}
	return o.SmtpUserName, true
}

// HasSmtpUserName returns a boolean if a field has been set.
func (o *EmailSettingsDto) HasSmtpUserName() bool {
	if o != nil && !IsNil(o.SmtpUserName) {
		return true
	}

	return false
}

// SetSmtpUserName gets a reference to the given string and assigns it to the SmtpUserName field.
func (o *EmailSettingsDto) SetSmtpUserName(v string) {
	o.SmtpUserName = &v
}

// GetSmtpPassword returns the SmtpPassword field value if set, zero value otherwise.
func (o *EmailSettingsDto) GetSmtpPassword() string {
	if o == nil || IsNil(o.SmtpPassword) {
		var ret string
		return ret
	}
	return *o.SmtpPassword
}

// GetSmtpPasswordOk returns a tuple with the SmtpPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsDto) GetSmtpPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpPassword) {
		return nil, false
	}
	return o.SmtpPassword, true
}

// HasSmtpPassword returns a boolean if a field has been set.
func (o *EmailSettingsDto) HasSmtpPassword() bool {
	if o != nil && !IsNil(o.SmtpPassword) {
		return true
	}

	return false
}

// SetSmtpPassword gets a reference to the given string and assigns it to the SmtpPassword field.
func (o *EmailSettingsDto) SetSmtpPassword(v string) {
	o.SmtpPassword = &v
}

// GetSmtpDomain returns the SmtpDomain field value if set, zero value otherwise.
func (o *EmailSettingsDto) GetSmtpDomain() string {
	if o == nil || IsNil(o.SmtpDomain) {
		var ret string
		return ret
	}
	return *o.SmtpDomain
}

// GetSmtpDomainOk returns a tuple with the SmtpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsDto) GetSmtpDomainOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpDomain) {
		return nil, false
	}
	return o.SmtpDomain, true
}

// HasSmtpDomain returns a boolean if a field has been set.
func (o *EmailSettingsDto) HasSmtpDomain() bool {
	if o != nil && !IsNil(o.SmtpDomain) {
		return true
	}

	return false
}

// SetSmtpDomain gets a reference to the given string and assigns it to the SmtpDomain field.
func (o *EmailSettingsDto) SetSmtpDomain(v string) {
	o.SmtpDomain = &v
}

// GetSmtpEnableSsl returns the SmtpEnableSsl field value if set, zero value otherwise.
func (o *EmailSettingsDto) GetSmtpEnableSsl() bool {
	if o == nil || IsNil(o.SmtpEnableSsl) {
		var ret bool
		return ret
	}
	return *o.SmtpEnableSsl
}

// GetSmtpEnableSslOk returns a tuple with the SmtpEnableSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsDto) GetSmtpEnableSslOk() (*bool, bool) {
	if o == nil || IsNil(o.SmtpEnableSsl) {
		return nil, false
	}
	return o.SmtpEnableSsl, true
}

// HasSmtpEnableSsl returns a boolean if a field has been set.
func (o *EmailSettingsDto) HasSmtpEnableSsl() bool {
	if o != nil && !IsNil(o.SmtpEnableSsl) {
		return true
	}

	return false
}

// SetSmtpEnableSsl gets a reference to the given bool and assigns it to the SmtpEnableSsl field.
func (o *EmailSettingsDto) SetSmtpEnableSsl(v bool) {
	o.SmtpEnableSsl = &v
}

// GetSmtpUseDefaultCredentials returns the SmtpUseDefaultCredentials field value if set, zero value otherwise.
func (o *EmailSettingsDto) GetSmtpUseDefaultCredentials() bool {
	if o == nil || IsNil(o.SmtpUseDefaultCredentials) {
		var ret bool
		return ret
	}
	return *o.SmtpUseDefaultCredentials
}

// GetSmtpUseDefaultCredentialsOk returns a tuple with the SmtpUseDefaultCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsDto) GetSmtpUseDefaultCredentialsOk() (*bool, bool) {
	if o == nil || IsNil(o.SmtpUseDefaultCredentials) {
		return nil, false
	}
	return o.SmtpUseDefaultCredentials, true
}

// HasSmtpUseDefaultCredentials returns a boolean if a field has been set.
func (o *EmailSettingsDto) HasSmtpUseDefaultCredentials() bool {
	if o != nil && !IsNil(o.SmtpUseDefaultCredentials) {
		return true
	}

	return false
}

// SetSmtpUseDefaultCredentials gets a reference to the given bool and assigns it to the SmtpUseDefaultCredentials field.
func (o *EmailSettingsDto) SetSmtpUseDefaultCredentials(v bool) {
	o.SmtpUseDefaultCredentials = &v
}

// GetDefaultFromAddress returns the DefaultFromAddress field value if set, zero value otherwise.
func (o *EmailSettingsDto) GetDefaultFromAddress() string {
	if o == nil || IsNil(o.DefaultFromAddress) {
		var ret string
		return ret
	}
	return *o.DefaultFromAddress
}

// GetDefaultFromAddressOk returns a tuple with the DefaultFromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsDto) GetDefaultFromAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultFromAddress) {
		return nil, false
	}
	return o.DefaultFromAddress, true
}

// HasDefaultFromAddress returns a boolean if a field has been set.
func (o *EmailSettingsDto) HasDefaultFromAddress() bool {
	if o != nil && !IsNil(o.DefaultFromAddress) {
		return true
	}

	return false
}

// SetDefaultFromAddress gets a reference to the given string and assigns it to the DefaultFromAddress field.
func (o *EmailSettingsDto) SetDefaultFromAddress(v string) {
	o.DefaultFromAddress = &v
}

// GetDefaultFromDisplayName returns the DefaultFromDisplayName field value if set, zero value otherwise.
func (o *EmailSettingsDto) GetDefaultFromDisplayName() string {
	if o == nil || IsNil(o.DefaultFromDisplayName) {
		var ret string
		return ret
	}
	return *o.DefaultFromDisplayName
}

// GetDefaultFromDisplayNameOk returns a tuple with the DefaultFromDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsDto) GetDefaultFromDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultFromDisplayName) {
		return nil, false
	}
	return o.DefaultFromDisplayName, true
}

// HasDefaultFromDisplayName returns a boolean if a field has been set.
func (o *EmailSettingsDto) HasDefaultFromDisplayName() bool {
	if o != nil && !IsNil(o.DefaultFromDisplayName) {
		return true
	}

	return false
}

// SetDefaultFromDisplayName gets a reference to the given string and assigns it to the DefaultFromDisplayName field.
func (o *EmailSettingsDto) SetDefaultFromDisplayName(v string) {
	o.DefaultFromDisplayName = &v
}

func (o EmailSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SmtpHost) {
		toSerialize["smtpHost"] = o.SmtpHost
	}
	if !IsNil(o.SmtpPort) {
		toSerialize["smtpPort"] = o.SmtpPort
	}
	if !IsNil(o.SmtpUserName) {
		toSerialize["smtpUserName"] = o.SmtpUserName
	}
	if !IsNil(o.SmtpPassword) {
		toSerialize["smtpPassword"] = o.SmtpPassword
	}
	if !IsNil(o.SmtpDomain) {
		toSerialize["smtpDomain"] = o.SmtpDomain
	}
	if !IsNil(o.SmtpEnableSsl) {
		toSerialize["smtpEnableSsl"] = o.SmtpEnableSsl
	}
	if !IsNil(o.SmtpUseDefaultCredentials) {
		toSerialize["smtpUseDefaultCredentials"] = o.SmtpUseDefaultCredentials
	}
	if !IsNil(o.DefaultFromAddress) {
		toSerialize["defaultFromAddress"] = o.DefaultFromAddress
	}
	if !IsNil(o.DefaultFromDisplayName) {
		toSerialize["defaultFromDisplayName"] = o.DefaultFromDisplayName
	}
	return toSerialize, nil
}

type NullableEmailSettingsDto struct {
	value *EmailSettingsDto
	isSet bool
}

func (v NullableEmailSettingsDto) Get() *EmailSettingsDto {
	return v.value
}

func (v *NullableEmailSettingsDto) Set(val *EmailSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSettingsDto(val *EmailSettingsDto) *NullableEmailSettingsDto {
	return &NullableEmailSettingsDto{value: val, isSet: true}
}

func (v NullableEmailSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



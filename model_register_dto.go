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

// checks if the RegisterDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterDto{}

// RegisterDto struct for RegisterDto
type RegisterDto struct {
	ExtraProperties map[string]map[string]interface{} `json:"extraProperties,omitempty"`
	UserName string `json:"userName"`
	EmailAddress string `json:"emailAddress"`
	Password string `json:"password"`
	AppName string `json:"appName"`
}

// NewRegisterDto instantiates a new RegisterDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterDto(userName string, emailAddress string, password string, appName string) *RegisterDto {
	this := RegisterDto{}
	this.UserName = userName
	this.EmailAddress = emailAddress
	this.Password = password
	this.AppName = appName
	return &this
}

// NewRegisterDtoWithDefaults instantiates a new RegisterDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterDtoWithDefaults() *RegisterDto {
	this := RegisterDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise.
func (o *RegisterDto) GetExtraProperties() map[string]map[string]interface{} {
	if o == nil || IsNil(o.ExtraProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDto) GetExtraPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtraProperties) {
		return map[string]map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *RegisterDto) HasExtraProperties() bool {
	if o != nil && !IsNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtraProperties field.
func (o *RegisterDto) SetExtraProperties(v map[string]map[string]interface{}) {
	o.ExtraProperties = v
}

// GetUserName returns the UserName field value
func (o *RegisterDto) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *RegisterDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *RegisterDto) SetUserName(v string) {
	o.UserName = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *RegisterDto) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *RegisterDto) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *RegisterDto) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetPassword returns the Password field value
func (o *RegisterDto) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *RegisterDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *RegisterDto) SetPassword(v string) {
	o.Password = v
}

// GetAppName returns the AppName field value
func (o *RegisterDto) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *RegisterDto) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *RegisterDto) SetAppName(v string) {
	o.AppName = v
}

func (o RegisterDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: extraProperties is readOnly
	toSerialize["userName"] = o.UserName
	toSerialize["emailAddress"] = o.EmailAddress
	toSerialize["password"] = o.Password
	toSerialize["appName"] = o.AppName
	return toSerialize, nil
}

type NullableRegisterDto struct {
	value *RegisterDto
	isSet bool
}

func (v NullableRegisterDto) Get() *RegisterDto {
	return v.value
}

func (v *NullableRegisterDto) Set(val *RegisterDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterDto(val *RegisterDto) *NullableRegisterDto {
	return &NullableRegisterDto{value: val, isSet: true}
}

func (v NullableRegisterDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



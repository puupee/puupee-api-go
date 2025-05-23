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

// checks if the IdentityUserLogin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityUserLogin{}

// IdentityUserLogin struct for IdentityUserLogin
type IdentityUserLogin struct {
	TenantId *string `json:"tenantId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	LoginProvider *string `json:"loginProvider,omitempty"`
	ProviderKey *string `json:"providerKey,omitempty"`
	ProviderDisplayName *string `json:"providerDisplayName,omitempty"`
}

// NewIdentityUserLogin instantiates a new IdentityUserLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUserLogin() *IdentityUserLogin {
	this := IdentityUserLogin{}
	return &this
}

// NewIdentityUserLoginWithDefaults instantiates a new IdentityUserLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUserLoginWithDefaults() *IdentityUserLogin {
	this := IdentityUserLogin{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IdentityUserLogin) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserLogin) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IdentityUserLogin) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *IdentityUserLogin) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IdentityUserLogin) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserLogin) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IdentityUserLogin) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IdentityUserLogin) SetUserId(v string) {
	o.UserId = &v
}

// GetLoginProvider returns the LoginProvider field value if set, zero value otherwise.
func (o *IdentityUserLogin) GetLoginProvider() string {
	if o == nil || IsNil(o.LoginProvider) {
		var ret string
		return ret
	}
	return *o.LoginProvider
}

// GetLoginProviderOk returns a tuple with the LoginProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserLogin) GetLoginProviderOk() (*string, bool) {
	if o == nil || IsNil(o.LoginProvider) {
		return nil, false
	}
	return o.LoginProvider, true
}

// HasLoginProvider returns a boolean if a field has been set.
func (o *IdentityUserLogin) HasLoginProvider() bool {
	if o != nil && !IsNil(o.LoginProvider) {
		return true
	}

	return false
}

// SetLoginProvider gets a reference to the given string and assigns it to the LoginProvider field.
func (o *IdentityUserLogin) SetLoginProvider(v string) {
	o.LoginProvider = &v
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise.
func (o *IdentityUserLogin) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey) {
		var ret string
		return ret
	}
	return *o.ProviderKey
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserLogin) GetProviderKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderKey) {
		return nil, false
	}
	return o.ProviderKey, true
}

// HasProviderKey returns a boolean if a field has been set.
func (o *IdentityUserLogin) HasProviderKey() bool {
	if o != nil && !IsNil(o.ProviderKey) {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given string and assigns it to the ProviderKey field.
func (o *IdentityUserLogin) SetProviderKey(v string) {
	o.ProviderKey = &v
}

// GetProviderDisplayName returns the ProviderDisplayName field value if set, zero value otherwise.
func (o *IdentityUserLogin) GetProviderDisplayName() string {
	if o == nil || IsNil(o.ProviderDisplayName) {
		var ret string
		return ret
	}
	return *o.ProviderDisplayName
}

// GetProviderDisplayNameOk returns a tuple with the ProviderDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserLogin) GetProviderDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderDisplayName) {
		return nil, false
	}
	return o.ProviderDisplayName, true
}

// HasProviderDisplayName returns a boolean if a field has been set.
func (o *IdentityUserLogin) HasProviderDisplayName() bool {
	if o != nil && !IsNil(o.ProviderDisplayName) {
		return true
	}

	return false
}

// SetProviderDisplayName gets a reference to the given string and assigns it to the ProviderDisplayName field.
func (o *IdentityUserLogin) SetProviderDisplayName(v string) {
	o.ProviderDisplayName = &v
}

func (o IdentityUserLogin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityUserLogin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.LoginProvider) {
		toSerialize["loginProvider"] = o.LoginProvider
	}
	if !IsNil(o.ProviderKey) {
		toSerialize["providerKey"] = o.ProviderKey
	}
	if !IsNil(o.ProviderDisplayName) {
		toSerialize["providerDisplayName"] = o.ProviderDisplayName
	}
	return toSerialize, nil
}

type NullableIdentityUserLogin struct {
	value *IdentityUserLogin
	isSet bool
}

func (v NullableIdentityUserLogin) Get() *IdentityUserLogin {
	return v.value
}

func (v *NullableIdentityUserLogin) Set(val *IdentityUserLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserLogin(val *IdentityUserLogin) *NullableIdentityUserLogin {
	return &NullableIdentityUserLogin{value: val, isSet: true}
}

func (v NullableIdentityUserLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



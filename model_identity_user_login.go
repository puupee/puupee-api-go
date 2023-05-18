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

// checks if the IdentityUserLogin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityUserLogin{}

// IdentityUserLogin struct for IdentityUserLogin
type IdentityUserLogin struct {
	TenantId NullableString `json:"tenantId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	LoginProvider NullableString `json:"loginProvider,omitempty"`
	ProviderKey NullableString `json:"providerKey,omitempty"`
	ProviderDisplayName NullableString `json:"providerDisplayName,omitempty"`
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

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserLogin) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserLogin) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *IdentityUserLogin) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *IdentityUserLogin) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *IdentityUserLogin) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *IdentityUserLogin) UnsetTenantId() {
	o.TenantId.Unset()
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

// GetLoginProvider returns the LoginProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserLogin) GetLoginProvider() string {
	if o == nil || IsNil(o.LoginProvider.Get()) {
		var ret string
		return ret
	}
	return *o.LoginProvider.Get()
}

// GetLoginProviderOk returns a tuple with the LoginProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserLogin) GetLoginProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginProvider.Get(), o.LoginProvider.IsSet()
}

// HasLoginProvider returns a boolean if a field has been set.
func (o *IdentityUserLogin) HasLoginProvider() bool {
	if o != nil && o.LoginProvider.IsSet() {
		return true
	}

	return false
}

// SetLoginProvider gets a reference to the given NullableString and assigns it to the LoginProvider field.
func (o *IdentityUserLogin) SetLoginProvider(v string) {
	o.LoginProvider.Set(&v)
}
// SetLoginProviderNil sets the value for LoginProvider to be an explicit nil
func (o *IdentityUserLogin) SetLoginProviderNil() {
	o.LoginProvider.Set(nil)
}

// UnsetLoginProvider ensures that no value is present for LoginProvider, not even an explicit nil
func (o *IdentityUserLogin) UnsetLoginProvider() {
	o.LoginProvider.Unset()
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserLogin) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderKey.Get()
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserLogin) GetProviderKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderKey.Get(), o.ProviderKey.IsSet()
}

// HasProviderKey returns a boolean if a field has been set.
func (o *IdentityUserLogin) HasProviderKey() bool {
	if o != nil && o.ProviderKey.IsSet() {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given NullableString and assigns it to the ProviderKey field.
func (o *IdentityUserLogin) SetProviderKey(v string) {
	o.ProviderKey.Set(&v)
}
// SetProviderKeyNil sets the value for ProviderKey to be an explicit nil
func (o *IdentityUserLogin) SetProviderKeyNil() {
	o.ProviderKey.Set(nil)
}

// UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
func (o *IdentityUserLogin) UnsetProviderKey() {
	o.ProviderKey.Unset()
}

// GetProviderDisplayName returns the ProviderDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserLogin) GetProviderDisplayName() string {
	if o == nil || IsNil(o.ProviderDisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderDisplayName.Get()
}

// GetProviderDisplayNameOk returns a tuple with the ProviderDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserLogin) GetProviderDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderDisplayName.Get(), o.ProviderDisplayName.IsSet()
}

// HasProviderDisplayName returns a boolean if a field has been set.
func (o *IdentityUserLogin) HasProviderDisplayName() bool {
	if o != nil && o.ProviderDisplayName.IsSet() {
		return true
	}

	return false
}

// SetProviderDisplayName gets a reference to the given NullableString and assigns it to the ProviderDisplayName field.
func (o *IdentityUserLogin) SetProviderDisplayName(v string) {
	o.ProviderDisplayName.Set(&v)
}
// SetProviderDisplayNameNil sets the value for ProviderDisplayName to be an explicit nil
func (o *IdentityUserLogin) SetProviderDisplayNameNil() {
	o.ProviderDisplayName.Set(nil)
}

// UnsetProviderDisplayName ensures that no value is present for ProviderDisplayName, not even an explicit nil
func (o *IdentityUserLogin) UnsetProviderDisplayName() {
	o.ProviderDisplayName.Unset()
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
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	// skip: userId is readOnly
	if o.LoginProvider.IsSet() {
		toSerialize["loginProvider"] = o.LoginProvider.Get()
	}
	if o.ProviderKey.IsSet() {
		toSerialize["providerKey"] = o.ProviderKey.Get()
	}
	if o.ProviderDisplayName.IsSet() {
		toSerialize["providerDisplayName"] = o.ProviderDisplayName.Get()
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



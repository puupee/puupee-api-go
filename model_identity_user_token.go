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

// checks if the IdentityUserToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityUserToken{}

// IdentityUserToken struct for IdentityUserToken
type IdentityUserToken struct {
	TenantId *string `json:"tenantId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	LoginProvider *string `json:"loginProvider,omitempty"`
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewIdentityUserToken instantiates a new IdentityUserToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUserToken() *IdentityUserToken {
	this := IdentityUserToken{}
	return &this
}

// NewIdentityUserTokenWithDefaults instantiates a new IdentityUserToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUserTokenWithDefaults() *IdentityUserToken {
	this := IdentityUserToken{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IdentityUserToken) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserToken) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IdentityUserToken) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *IdentityUserToken) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IdentityUserToken) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserToken) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IdentityUserToken) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IdentityUserToken) SetUserId(v string) {
	o.UserId = &v
}

// GetLoginProvider returns the LoginProvider field value if set, zero value otherwise.
func (o *IdentityUserToken) GetLoginProvider() string {
	if o == nil || IsNil(o.LoginProvider) {
		var ret string
		return ret
	}
	return *o.LoginProvider
}

// GetLoginProviderOk returns a tuple with the LoginProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserToken) GetLoginProviderOk() (*string, bool) {
	if o == nil || IsNil(o.LoginProvider) {
		return nil, false
	}
	return o.LoginProvider, true
}

// HasLoginProvider returns a boolean if a field has been set.
func (o *IdentityUserToken) HasLoginProvider() bool {
	if o != nil && !IsNil(o.LoginProvider) {
		return true
	}

	return false
}

// SetLoginProvider gets a reference to the given string and assigns it to the LoginProvider field.
func (o *IdentityUserToken) SetLoginProvider(v string) {
	o.LoginProvider = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityUserToken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserToken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityUserToken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityUserToken) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IdentityUserToken) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserToken) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IdentityUserToken) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IdentityUserToken) SetValue(v string) {
	o.Value = &v
}

func (o IdentityUserToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityUserToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: tenantId is readOnly
	// skip: userId is readOnly
	// skip: loginProvider is readOnly
	// skip: name is readOnly
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableIdentityUserToken struct {
	value *IdentityUserToken
	isSet bool
}

func (v NullableIdentityUserToken) Get() *IdentityUserToken {
	return v.value
}

func (v *NullableIdentityUserToken) Set(val *IdentityUserToken) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserToken) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserToken(val *IdentityUserToken) *NullableIdentityUserToken {
	return &NullableIdentityUserToken{value: val, isSet: true}
}

func (v NullableIdentityUserToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


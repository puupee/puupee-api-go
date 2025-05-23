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

// checks if the IdentityUserClaim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityUserClaim{}

// IdentityUserClaim struct for IdentityUserClaim
type IdentityUserClaim struct {
	Id *string `json:"id,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	ClaimType *string `json:"claimType,omitempty"`
	ClaimValue *string `json:"claimValue,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewIdentityUserClaim instantiates a new IdentityUserClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUserClaim() *IdentityUserClaim {
	this := IdentityUserClaim{}
	return &this
}

// NewIdentityUserClaimWithDefaults instantiates a new IdentityUserClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUserClaimWithDefaults() *IdentityUserClaim {
	this := IdentityUserClaim{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityUserClaim) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserClaim) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityUserClaim) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityUserClaim) SetId(v string) {
	o.Id = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *IdentityUserClaim) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserClaim) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *IdentityUserClaim) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *IdentityUserClaim) SetTenantId(v string) {
	o.TenantId = &v
}

// GetClaimType returns the ClaimType field value if set, zero value otherwise.
func (o *IdentityUserClaim) GetClaimType() string {
	if o == nil || IsNil(o.ClaimType) {
		var ret string
		return ret
	}
	return *o.ClaimType
}

// GetClaimTypeOk returns a tuple with the ClaimType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserClaim) GetClaimTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimType) {
		return nil, false
	}
	return o.ClaimType, true
}

// HasClaimType returns a boolean if a field has been set.
func (o *IdentityUserClaim) HasClaimType() bool {
	if o != nil && !IsNil(o.ClaimType) {
		return true
	}

	return false
}

// SetClaimType gets a reference to the given string and assigns it to the ClaimType field.
func (o *IdentityUserClaim) SetClaimType(v string) {
	o.ClaimType = &v
}

// GetClaimValue returns the ClaimValue field value if set, zero value otherwise.
func (o *IdentityUserClaim) GetClaimValue() string {
	if o == nil || IsNil(o.ClaimValue) {
		var ret string
		return ret
	}
	return *o.ClaimValue
}

// GetClaimValueOk returns a tuple with the ClaimValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserClaim) GetClaimValueOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimValue) {
		return nil, false
	}
	return o.ClaimValue, true
}

// HasClaimValue returns a boolean if a field has been set.
func (o *IdentityUserClaim) HasClaimValue() bool {
	if o != nil && !IsNil(o.ClaimValue) {
		return true
	}

	return false
}

// SetClaimValue gets a reference to the given string and assigns it to the ClaimValue field.
func (o *IdentityUserClaim) SetClaimValue(v string) {
	o.ClaimValue = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IdentityUserClaim) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserClaim) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IdentityUserClaim) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IdentityUserClaim) SetUserId(v string) {
	o.UserId = &v
}

func (o IdentityUserClaim) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityUserClaim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ClaimType) {
		toSerialize["claimType"] = o.ClaimType
	}
	if !IsNil(o.ClaimValue) {
		toSerialize["claimValue"] = o.ClaimValue
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableIdentityUserClaim struct {
	value *IdentityUserClaim
	isSet bool
}

func (v NullableIdentityUserClaim) Get() *IdentityUserClaim {
	return v.value
}

func (v *NullableIdentityUserClaim) Set(val *IdentityUserClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserClaim(val *IdentityUserClaim) *NullableIdentityUserClaim {
	return &NullableIdentityUserClaim{value: val, isSet: true}
}

func (v NullableIdentityUserClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



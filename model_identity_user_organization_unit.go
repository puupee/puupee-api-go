/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// checks if the IdentityUserOrganizationUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityUserOrganizationUnit{}

// IdentityUserOrganizationUnit struct for IdentityUserOrganizationUnit
type IdentityUserOrganizationUnit struct {
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId NullableString `json:"creatorId,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	OrganizationUnitId *string `json:"organizationUnitId,omitempty"`
}

// NewIdentityUserOrganizationUnit instantiates a new IdentityUserOrganizationUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUserOrganizationUnit() *IdentityUserOrganizationUnit {
	this := IdentityUserOrganizationUnit{}
	return &this
}

// NewIdentityUserOrganizationUnitWithDefaults instantiates a new IdentityUserOrganizationUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUserOrganizationUnitWithDefaults() *IdentityUserOrganizationUnit {
	this := IdentityUserOrganizationUnit{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *IdentityUserOrganizationUnit) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserOrganizationUnit) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *IdentityUserOrganizationUnit) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *IdentityUserOrganizationUnit) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserOrganizationUnit) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserOrganizationUnit) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *IdentityUserOrganizationUnit) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableString and assigns it to the CreatorId field.
func (o *IdentityUserOrganizationUnit) SetCreatorId(v string) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *IdentityUserOrganizationUnit) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *IdentityUserOrganizationUnit) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserOrganizationUnit) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserOrganizationUnit) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *IdentityUserOrganizationUnit) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *IdentityUserOrganizationUnit) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *IdentityUserOrganizationUnit) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *IdentityUserOrganizationUnit) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IdentityUserOrganizationUnit) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserOrganizationUnit) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IdentityUserOrganizationUnit) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IdentityUserOrganizationUnit) SetUserId(v string) {
	o.UserId = &v
}

// GetOrganizationUnitId returns the OrganizationUnitId field value if set, zero value otherwise.
func (o *IdentityUserOrganizationUnit) GetOrganizationUnitId() string {
	if o == nil || IsNil(o.OrganizationUnitId) {
		var ret string
		return ret
	}
	return *o.OrganizationUnitId
}

// GetOrganizationUnitIdOk returns a tuple with the OrganizationUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserOrganizationUnit) GetOrganizationUnitIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationUnitId) {
		return nil, false
	}
	return o.OrganizationUnitId, true
}

// HasOrganizationUnitId returns a boolean if a field has been set.
func (o *IdentityUserOrganizationUnit) HasOrganizationUnitId() bool {
	if o != nil && !IsNil(o.OrganizationUnitId) {
		return true
	}

	return false
}

// SetOrganizationUnitId gets a reference to the given string and assigns it to the OrganizationUnitId field.
func (o *IdentityUserOrganizationUnit) SetOrganizationUnitId(v string) {
	o.OrganizationUnitId = &v
}

func (o IdentityUserOrganizationUnit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityUserOrganizationUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: creationTime is readOnly
	if o.CreatorId.IsSet() {
		toSerialize["creatorId"] = o.CreatorId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.OrganizationUnitId) {
		toSerialize["organizationUnitId"] = o.OrganizationUnitId
	}
	return toSerialize, nil
}

type NullableIdentityUserOrganizationUnit struct {
	value *IdentityUserOrganizationUnit
	isSet bool
}

func (v NullableIdentityUserOrganizationUnit) Get() *IdentityUserOrganizationUnit {
	return v.value
}

func (v *NullableIdentityUserOrganizationUnit) Set(val *IdentityUserOrganizationUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserOrganizationUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserOrganizationUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserOrganizationUnit(val *IdentityUserOrganizationUnit) *NullableIdentityUserOrganizationUnit {
	return &NullableIdentityUserOrganizationUnit{value: val, isSet: true}
}

func (v NullableIdentityUserOrganizationUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserOrganizationUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



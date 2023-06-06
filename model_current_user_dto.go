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

// checks if the CurrentUserDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentUserDto{}

// CurrentUserDto struct for CurrentUserDto
type CurrentUserDto struct {
	IsAuthenticated *bool `json:"isAuthenticated,omitempty"`
	Id *string `json:"id,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	ImpersonatorUserId *string `json:"impersonatorUserId,omitempty"`
	ImpersonatorTenantId *string `json:"impersonatorTenantId,omitempty"`
	ImpersonatorUserName *string `json:"impersonatorUserName,omitempty"`
	ImpersonatorTenantName *string `json:"impersonatorTenantName,omitempty"`
	UserName *string `json:"userName,omitempty"`
	Name *string `json:"name,omitempty"`
	SurName *string `json:"surName,omitempty"`
	Email *string `json:"email,omitempty"`
	EmailVerified *bool `json:"emailVerified,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	PhoneNumberVerified *bool `json:"phoneNumberVerified,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

// NewCurrentUserDto instantiates a new CurrentUserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentUserDto() *CurrentUserDto {
	this := CurrentUserDto{}
	return &this
}

// NewCurrentUserDtoWithDefaults instantiates a new CurrentUserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentUserDtoWithDefaults() *CurrentUserDto {
	this := CurrentUserDto{}
	return &this
}

// GetIsAuthenticated returns the IsAuthenticated field value if set, zero value otherwise.
func (o *CurrentUserDto) GetIsAuthenticated() bool {
	if o == nil || IsNil(o.IsAuthenticated) {
		var ret bool
		return ret
	}
	return *o.IsAuthenticated
}

// GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetIsAuthenticatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAuthenticated) {
		return nil, false
	}
	return o.IsAuthenticated, true
}

// HasIsAuthenticated returns a boolean if a field has been set.
func (o *CurrentUserDto) HasIsAuthenticated() bool {
	if o != nil && !IsNil(o.IsAuthenticated) {
		return true
	}

	return false
}

// SetIsAuthenticated gets a reference to the given bool and assigns it to the IsAuthenticated field.
func (o *CurrentUserDto) SetIsAuthenticated(v bool) {
	o.IsAuthenticated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CurrentUserDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CurrentUserDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CurrentUserDto) SetId(v string) {
	o.Id = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CurrentUserDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CurrentUserDto) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CurrentUserDto) SetTenantId(v string) {
	o.TenantId = &v
}

// GetImpersonatorUserId returns the ImpersonatorUserId field value if set, zero value otherwise.
func (o *CurrentUserDto) GetImpersonatorUserId() string {
	if o == nil || IsNil(o.ImpersonatorUserId) {
		var ret string
		return ret
	}
	return *o.ImpersonatorUserId
}

// GetImpersonatorUserIdOk returns a tuple with the ImpersonatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetImpersonatorUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImpersonatorUserId) {
		return nil, false
	}
	return o.ImpersonatorUserId, true
}

// HasImpersonatorUserId returns a boolean if a field has been set.
func (o *CurrentUserDto) HasImpersonatorUserId() bool {
	if o != nil && !IsNil(o.ImpersonatorUserId) {
		return true
	}

	return false
}

// SetImpersonatorUserId gets a reference to the given string and assigns it to the ImpersonatorUserId field.
func (o *CurrentUserDto) SetImpersonatorUserId(v string) {
	o.ImpersonatorUserId = &v
}

// GetImpersonatorTenantId returns the ImpersonatorTenantId field value if set, zero value otherwise.
func (o *CurrentUserDto) GetImpersonatorTenantId() string {
	if o == nil || IsNil(o.ImpersonatorTenantId) {
		var ret string
		return ret
	}
	return *o.ImpersonatorTenantId
}

// GetImpersonatorTenantIdOk returns a tuple with the ImpersonatorTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetImpersonatorTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImpersonatorTenantId) {
		return nil, false
	}
	return o.ImpersonatorTenantId, true
}

// HasImpersonatorTenantId returns a boolean if a field has been set.
func (o *CurrentUserDto) HasImpersonatorTenantId() bool {
	if o != nil && !IsNil(o.ImpersonatorTenantId) {
		return true
	}

	return false
}

// SetImpersonatorTenantId gets a reference to the given string and assigns it to the ImpersonatorTenantId field.
func (o *CurrentUserDto) SetImpersonatorTenantId(v string) {
	o.ImpersonatorTenantId = &v
}

// GetImpersonatorUserName returns the ImpersonatorUserName field value if set, zero value otherwise.
func (o *CurrentUserDto) GetImpersonatorUserName() string {
	if o == nil || IsNil(o.ImpersonatorUserName) {
		var ret string
		return ret
	}
	return *o.ImpersonatorUserName
}

// GetImpersonatorUserNameOk returns a tuple with the ImpersonatorUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetImpersonatorUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImpersonatorUserName) {
		return nil, false
	}
	return o.ImpersonatorUserName, true
}

// HasImpersonatorUserName returns a boolean if a field has been set.
func (o *CurrentUserDto) HasImpersonatorUserName() bool {
	if o != nil && !IsNil(o.ImpersonatorUserName) {
		return true
	}

	return false
}

// SetImpersonatorUserName gets a reference to the given string and assigns it to the ImpersonatorUserName field.
func (o *CurrentUserDto) SetImpersonatorUserName(v string) {
	o.ImpersonatorUserName = &v
}

// GetImpersonatorTenantName returns the ImpersonatorTenantName field value if set, zero value otherwise.
func (o *CurrentUserDto) GetImpersonatorTenantName() string {
	if o == nil || IsNil(o.ImpersonatorTenantName) {
		var ret string
		return ret
	}
	return *o.ImpersonatorTenantName
}

// GetImpersonatorTenantNameOk returns a tuple with the ImpersonatorTenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetImpersonatorTenantNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImpersonatorTenantName) {
		return nil, false
	}
	return o.ImpersonatorTenantName, true
}

// HasImpersonatorTenantName returns a boolean if a field has been set.
func (o *CurrentUserDto) HasImpersonatorTenantName() bool {
	if o != nil && !IsNil(o.ImpersonatorTenantName) {
		return true
	}

	return false
}

// SetImpersonatorTenantName gets a reference to the given string and assigns it to the ImpersonatorTenantName field.
func (o *CurrentUserDto) SetImpersonatorTenantName(v string) {
	o.ImpersonatorTenantName = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *CurrentUserDto) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *CurrentUserDto) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *CurrentUserDto) SetUserName(v string) {
	o.UserName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CurrentUserDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CurrentUserDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CurrentUserDto) SetName(v string) {
	o.Name = &v
}

// GetSurName returns the SurName field value if set, zero value otherwise.
func (o *CurrentUserDto) GetSurName() string {
	if o == nil || IsNil(o.SurName) {
		var ret string
		return ret
	}
	return *o.SurName
}

// GetSurNameOk returns a tuple with the SurName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetSurNameOk() (*string, bool) {
	if o == nil || IsNil(o.SurName) {
		return nil, false
	}
	return o.SurName, true
}

// HasSurName returns a boolean if a field has been set.
func (o *CurrentUserDto) HasSurName() bool {
	if o != nil && !IsNil(o.SurName) {
		return true
	}

	return false
}

// SetSurName gets a reference to the given string and assigns it to the SurName field.
func (o *CurrentUserDto) SetSurName(v string) {
	o.SurName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CurrentUserDto) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CurrentUserDto) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CurrentUserDto) SetEmail(v string) {
	o.Email = &v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *CurrentUserDto) GetEmailVerified() bool {
	if o == nil || IsNil(o.EmailVerified) {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailVerified) {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *CurrentUserDto) HasEmailVerified() bool {
	if o != nil && !IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *CurrentUserDto) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CurrentUserDto) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CurrentUserDto) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CurrentUserDto) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneNumberVerified returns the PhoneNumberVerified field value if set, zero value otherwise.
func (o *CurrentUserDto) GetPhoneNumberVerified() bool {
	if o == nil || IsNil(o.PhoneNumberVerified) {
		var ret bool
		return ret
	}
	return *o.PhoneNumberVerified
}

// GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetPhoneNumberVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.PhoneNumberVerified) {
		return nil, false
	}
	return o.PhoneNumberVerified, true
}

// HasPhoneNumberVerified returns a boolean if a field has been set.
func (o *CurrentUserDto) HasPhoneNumberVerified() bool {
	if o != nil && !IsNil(o.PhoneNumberVerified) {
		return true
	}

	return false
}

// SetPhoneNumberVerified gets a reference to the given bool and assigns it to the PhoneNumberVerified field.
func (o *CurrentUserDto) SetPhoneNumberVerified(v bool) {
	o.PhoneNumberVerified = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CurrentUserDto) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUserDto) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CurrentUserDto) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CurrentUserDto) SetRoles(v []string) {
	o.Roles = v
}

func (o CurrentUserDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentUserDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAuthenticated) {
		toSerialize["isAuthenticated"] = o.IsAuthenticated
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ImpersonatorUserId) {
		toSerialize["impersonatorUserId"] = o.ImpersonatorUserId
	}
	if !IsNil(o.ImpersonatorTenantId) {
		toSerialize["impersonatorTenantId"] = o.ImpersonatorTenantId
	}
	if !IsNil(o.ImpersonatorUserName) {
		toSerialize["impersonatorUserName"] = o.ImpersonatorUserName
	}
	if !IsNil(o.ImpersonatorTenantName) {
		toSerialize["impersonatorTenantName"] = o.ImpersonatorTenantName
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SurName) {
		toSerialize["surName"] = o.SurName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EmailVerified) {
		toSerialize["emailVerified"] = o.EmailVerified
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.PhoneNumberVerified) {
		toSerialize["phoneNumberVerified"] = o.PhoneNumberVerified
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableCurrentUserDto struct {
	value *CurrentUserDto
	isSet bool
}

func (v NullableCurrentUserDto) Get() *CurrentUserDto {
	return v.value
}

func (v *NullableCurrentUserDto) Set(val *CurrentUserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentUserDto(val *CurrentUserDto) *NullableCurrentUserDto {
	return &NullableCurrentUserDto{value: val, isSet: true}
}

func (v NullableCurrentUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



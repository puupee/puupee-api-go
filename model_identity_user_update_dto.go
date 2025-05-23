/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.17.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the IdentityUserUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityUserUpdateDto{}

// IdentityUserUpdateDto struct for IdentityUserUpdateDto
type IdentityUserUpdateDto struct {
	ExtraProperties map[string]map[string]interface{} `json:"extraProperties,omitempty"`
	UserName string `json:"userName"`
	Name *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
	Email string `json:"email"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	LockoutEnabled *bool `json:"lockoutEnabled,omitempty"`
	RoleNames []string `json:"roleNames,omitempty"`
	Password *string `json:"password,omitempty"`
	ConcurrencyStamp *string `json:"concurrencyStamp,omitempty"`
}

type _IdentityUserUpdateDto IdentityUserUpdateDto

// NewIdentityUserUpdateDto instantiates a new IdentityUserUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUserUpdateDto(userName string, email string) *IdentityUserUpdateDto {
	this := IdentityUserUpdateDto{}
	this.UserName = userName
	this.Email = email
	return &this
}

// NewIdentityUserUpdateDtoWithDefaults instantiates a new IdentityUserUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUserUpdateDtoWithDefaults() *IdentityUserUpdateDto {
	this := IdentityUserUpdateDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetExtraProperties() map[string]map[string]interface{} {
	if o == nil || IsNil(o.ExtraProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetExtraPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtraProperties) {
		return map[string]map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasExtraProperties() bool {
	if o != nil && !IsNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtraProperties field.
func (o *IdentityUserUpdateDto) SetExtraProperties(v map[string]map[string]interface{}) {
	o.ExtraProperties = v
}

// GetUserName returns the UserName field value
func (o *IdentityUserUpdateDto) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *IdentityUserUpdateDto) SetUserName(v string) {
	o.UserName = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityUserUpdateDto) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *IdentityUserUpdateDto) SetSurname(v string) {
	o.Surname = &v
}

// GetEmail returns the Email field value
func (o *IdentityUserUpdateDto) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *IdentityUserUpdateDto) SetEmail(v string) {
	o.Email = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *IdentityUserUpdateDto) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *IdentityUserUpdateDto) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLockoutEnabled returns the LockoutEnabled field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetLockoutEnabled() bool {
	if o == nil || IsNil(o.LockoutEnabled) {
		var ret bool
		return ret
	}
	return *o.LockoutEnabled
}

// GetLockoutEnabledOk returns a tuple with the LockoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetLockoutEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LockoutEnabled) {
		return nil, false
	}
	return o.LockoutEnabled, true
}

// HasLockoutEnabled returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasLockoutEnabled() bool {
	if o != nil && !IsNil(o.LockoutEnabled) {
		return true
	}

	return false
}

// SetLockoutEnabled gets a reference to the given bool and assigns it to the LockoutEnabled field.
func (o *IdentityUserUpdateDto) SetLockoutEnabled(v bool) {
	o.LockoutEnabled = &v
}

// GetRoleNames returns the RoleNames field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetRoleNames() []string {
	if o == nil || IsNil(o.RoleNames) {
		var ret []string
		return ret
	}
	return o.RoleNames
}

// GetRoleNamesOk returns a tuple with the RoleNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetRoleNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleNames) {
		return nil, false
	}
	return o.RoleNames, true
}

// HasRoleNames returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasRoleNames() bool {
	if o != nil && !IsNil(o.RoleNames) {
		return true
	}

	return false
}

// SetRoleNames gets a reference to the given []string and assigns it to the RoleNames field.
func (o *IdentityUserUpdateDto) SetRoleNames(v []string) {
	o.RoleNames = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IdentityUserUpdateDto) SetPassword(v string) {
	o.Password = &v
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise.
func (o *IdentityUserUpdateDto) GetConcurrencyStamp() string {
	if o == nil || IsNil(o.ConcurrencyStamp) {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUserUpdateDto) GetConcurrencyStampOk() (*string, bool) {
	if o == nil || IsNil(o.ConcurrencyStamp) {
		return nil, false
	}
	return o.ConcurrencyStamp, true
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *IdentityUserUpdateDto) HasConcurrencyStamp() bool {
	if o != nil && !IsNil(o.ConcurrencyStamp) {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given string and assigns it to the ConcurrencyStamp field.
func (o *IdentityUserUpdateDto) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp = &v
}

func (o IdentityUserUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityUserUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtraProperties) {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	toSerialize["userName"] = o.UserName
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.LockoutEnabled) {
		toSerialize["lockoutEnabled"] = o.LockoutEnabled
	}
	if !IsNil(o.RoleNames) {
		toSerialize["roleNames"] = o.RoleNames
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.ConcurrencyStamp) {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp
	}
	return toSerialize, nil
}

func (o *IdentityUserUpdateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userName",
		"email",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIdentityUserUpdateDto := _IdentityUserUpdateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentityUserUpdateDto)

	if err != nil {
		return err
	}

	*o = IdentityUserUpdateDto(varIdentityUserUpdateDto)

	return err
}

type NullableIdentityUserUpdateDto struct {
	value *IdentityUserUpdateDto
	isSet bool
}

func (v NullableIdentityUserUpdateDto) Get() *IdentityUserUpdateDto {
	return v.value
}

func (v *NullableIdentityUserUpdateDto) Set(val *IdentityUserUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserUpdateDto(val *IdentityUserUpdateDto) *NullableIdentityUserUpdateDto {
	return &NullableIdentityUserUpdateDto{value: val, isSet: true}
}

func (v NullableIdentityUserUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



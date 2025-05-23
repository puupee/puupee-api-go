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

// checks if the UserData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserData{}

// UserData struct for UserData
type UserData struct {
	Id *string `json:"id,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	UserName *string `json:"userName,omitempty"`
	Name *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	Email *string `json:"email,omitempty"`
	EmailConfirmed *bool `json:"emailConfirmed,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	PhoneNumberConfirmed *bool `json:"phoneNumberConfirmed,omitempty"`
	ExtraProperties map[string]map[string]interface{} `json:"extraProperties,omitempty"`
}

// NewUserData instantiates a new UserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserData() *UserData {
	this := UserData{}
	return &this
}

// NewUserDataWithDefaults instantiates a new UserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataWithDefaults() *UserData {
	this := UserData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserData) SetId(v string) {
	o.Id = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *UserData) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *UserData) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *UserData) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UserData) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UserData) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UserData) SetUserName(v string) {
	o.UserName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserData) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *UserData) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *UserData) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *UserData) SetSurname(v string) {
	o.Surname = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UserData) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UserData) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UserData) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserData) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserData) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserData) SetEmail(v string) {
	o.Email = &v
}

// GetEmailConfirmed returns the EmailConfirmed field value if set, zero value otherwise.
func (o *UserData) GetEmailConfirmed() bool {
	if o == nil || IsNil(o.EmailConfirmed) {
		var ret bool
		return ret
	}
	return *o.EmailConfirmed
}

// GetEmailConfirmedOk returns a tuple with the EmailConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetEmailConfirmedOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailConfirmed) {
		return nil, false
	}
	return o.EmailConfirmed, true
}

// HasEmailConfirmed returns a boolean if a field has been set.
func (o *UserData) HasEmailConfirmed() bool {
	if o != nil && !IsNil(o.EmailConfirmed) {
		return true
	}

	return false
}

// SetEmailConfirmed gets a reference to the given bool and assigns it to the EmailConfirmed field.
func (o *UserData) SetEmailConfirmed(v bool) {
	o.EmailConfirmed = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserData) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserData) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserData) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneNumberConfirmed returns the PhoneNumberConfirmed field value if set, zero value otherwise.
func (o *UserData) GetPhoneNumberConfirmed() bool {
	if o == nil || IsNil(o.PhoneNumberConfirmed) {
		var ret bool
		return ret
	}
	return *o.PhoneNumberConfirmed
}

// GetPhoneNumberConfirmedOk returns a tuple with the PhoneNumberConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetPhoneNumberConfirmedOk() (*bool, bool) {
	if o == nil || IsNil(o.PhoneNumberConfirmed) {
		return nil, false
	}
	return o.PhoneNumberConfirmed, true
}

// HasPhoneNumberConfirmed returns a boolean if a field has been set.
func (o *UserData) HasPhoneNumberConfirmed() bool {
	if o != nil && !IsNil(o.PhoneNumberConfirmed) {
		return true
	}

	return false
}

// SetPhoneNumberConfirmed gets a reference to the given bool and assigns it to the PhoneNumberConfirmed field.
func (o *UserData) SetPhoneNumberConfirmed(v bool) {
	o.PhoneNumberConfirmed = &v
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise.
func (o *UserData) GetExtraProperties() map[string]map[string]interface{} {
	if o == nil || IsNil(o.ExtraProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetExtraPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtraProperties) {
		return map[string]map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *UserData) HasExtraProperties() bool {
	if o != nil && !IsNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtraProperties field.
func (o *UserData) SetExtraProperties(v map[string]map[string]interface{}) {
	o.ExtraProperties = v
}

func (o UserData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EmailConfirmed) {
		toSerialize["emailConfirmed"] = o.EmailConfirmed
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.PhoneNumberConfirmed) {
		toSerialize["phoneNumberConfirmed"] = o.PhoneNumberConfirmed
	}
	if !IsNil(o.ExtraProperties) {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	return toSerialize, nil
}

type NullableUserData struct {
	value *UserData
	isSet bool
}

func (v NullableUserData) Get() *UserData {
	return v.value
}

func (v *NullableUserData) Set(val *UserData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserData(val *UserData) *NullableUserData {
	return &NullableUserData{value: val, isSet: true}
}

func (v NullableUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the CurrentTenantDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentTenantDto{}

// CurrentTenantDto struct for CurrentTenantDto
type CurrentTenantDto struct {
	Id NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	IsAvailable *bool `json:"isAvailable,omitempty"`
}

// NewCurrentTenantDto instantiates a new CurrentTenantDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentTenantDto() *CurrentTenantDto {
	this := CurrentTenantDto{}
	return &this
}

// NewCurrentTenantDtoWithDefaults instantiates a new CurrentTenantDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentTenantDtoWithDefaults() *CurrentTenantDto {
	this := CurrentTenantDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentTenantDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentTenantDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CurrentTenantDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CurrentTenantDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CurrentTenantDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CurrentTenantDto) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentTenantDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentTenantDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CurrentTenantDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CurrentTenantDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CurrentTenantDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CurrentTenantDto) UnsetName() {
	o.Name.Unset()
}

// GetIsAvailable returns the IsAvailable field value if set, zero value otherwise.
func (o *CurrentTenantDto) GetIsAvailable() bool {
	if o == nil || IsNil(o.IsAvailable) {
		var ret bool
		return ret
	}
	return *o.IsAvailable
}

// GetIsAvailableOk returns a tuple with the IsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentTenantDto) GetIsAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAvailable) {
		return nil, false
	}
	return o.IsAvailable, true
}

// HasIsAvailable returns a boolean if a field has been set.
func (o *CurrentTenantDto) HasIsAvailable() bool {
	if o != nil && !IsNil(o.IsAvailable) {
		return true
	}

	return false
}

// SetIsAvailable gets a reference to the given bool and assigns it to the IsAvailable field.
func (o *CurrentTenantDto) SetIsAvailable(v bool) {
	o.IsAvailable = &v
}

func (o CurrentTenantDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentTenantDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.IsAvailable) {
		toSerialize["isAvailable"] = o.IsAvailable
	}
	return toSerialize, nil
}

type NullableCurrentTenantDto struct {
	value *CurrentTenantDto
	isSet bool
}

func (v NullableCurrentTenantDto) Get() *CurrentTenantDto {
	return v.value
}

func (v *NullableCurrentTenantDto) Set(val *CurrentTenantDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentTenantDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentTenantDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentTenantDto(val *CurrentTenantDto) *NullableCurrentTenantDto {
	return &NullableCurrentTenantDto{value: val, isSet: true}
}

func (v NullableCurrentTenantDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentTenantDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



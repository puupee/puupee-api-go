/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
)

// ExtensionPropertyUiDto struct for ExtensionPropertyUiDto
type ExtensionPropertyUiDto struct {
	OnTable *ExtensionPropertyUiTableDto `json:"onTable,omitempty"`
	OnCreateForm *ExtensionPropertyUiFormDto `json:"onCreateForm,omitempty"`
	OnEditForm *ExtensionPropertyUiFormDto `json:"onEditForm,omitempty"`
	Lookup *ExtensionPropertyUiLookupDto `json:"lookup,omitempty"`
}

// NewExtensionPropertyUiDto instantiates a new ExtensionPropertyUiDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionPropertyUiDto() *ExtensionPropertyUiDto {
	this := ExtensionPropertyUiDto{}
	return &this
}

// NewExtensionPropertyUiDtoWithDefaults instantiates a new ExtensionPropertyUiDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionPropertyUiDtoWithDefaults() *ExtensionPropertyUiDto {
	this := ExtensionPropertyUiDto{}
	return &this
}

// GetOnTable returns the OnTable field value if set, zero value otherwise.
func (o *ExtensionPropertyUiDto) GetOnTable() ExtensionPropertyUiTableDto {
	if o == nil || o.OnTable == nil {
		var ret ExtensionPropertyUiTableDto
		return ret
	}
	return *o.OnTable
}

// GetOnTableOk returns a tuple with the OnTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyUiDto) GetOnTableOk() (*ExtensionPropertyUiTableDto, bool) {
	if o == nil || o.OnTable == nil {
		return nil, false
	}
	return o.OnTable, true
}

// HasOnTable returns a boolean if a field has been set.
func (o *ExtensionPropertyUiDto) HasOnTable() bool {
	if o != nil && o.OnTable != nil {
		return true
	}

	return false
}

// SetOnTable gets a reference to the given ExtensionPropertyUiTableDto and assigns it to the OnTable field.
func (o *ExtensionPropertyUiDto) SetOnTable(v ExtensionPropertyUiTableDto) {
	o.OnTable = &v
}

// GetOnCreateForm returns the OnCreateForm field value if set, zero value otherwise.
func (o *ExtensionPropertyUiDto) GetOnCreateForm() ExtensionPropertyUiFormDto {
	if o == nil || o.OnCreateForm == nil {
		var ret ExtensionPropertyUiFormDto
		return ret
	}
	return *o.OnCreateForm
}

// GetOnCreateFormOk returns a tuple with the OnCreateForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyUiDto) GetOnCreateFormOk() (*ExtensionPropertyUiFormDto, bool) {
	if o == nil || o.OnCreateForm == nil {
		return nil, false
	}
	return o.OnCreateForm, true
}

// HasOnCreateForm returns a boolean if a field has been set.
func (o *ExtensionPropertyUiDto) HasOnCreateForm() bool {
	if o != nil && o.OnCreateForm != nil {
		return true
	}

	return false
}

// SetOnCreateForm gets a reference to the given ExtensionPropertyUiFormDto and assigns it to the OnCreateForm field.
func (o *ExtensionPropertyUiDto) SetOnCreateForm(v ExtensionPropertyUiFormDto) {
	o.OnCreateForm = &v
}

// GetOnEditForm returns the OnEditForm field value if set, zero value otherwise.
func (o *ExtensionPropertyUiDto) GetOnEditForm() ExtensionPropertyUiFormDto {
	if o == nil || o.OnEditForm == nil {
		var ret ExtensionPropertyUiFormDto
		return ret
	}
	return *o.OnEditForm
}

// GetOnEditFormOk returns a tuple with the OnEditForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyUiDto) GetOnEditFormOk() (*ExtensionPropertyUiFormDto, bool) {
	if o == nil || o.OnEditForm == nil {
		return nil, false
	}
	return o.OnEditForm, true
}

// HasOnEditForm returns a boolean if a field has been set.
func (o *ExtensionPropertyUiDto) HasOnEditForm() bool {
	if o != nil && o.OnEditForm != nil {
		return true
	}

	return false
}

// SetOnEditForm gets a reference to the given ExtensionPropertyUiFormDto and assigns it to the OnEditForm field.
func (o *ExtensionPropertyUiDto) SetOnEditForm(v ExtensionPropertyUiFormDto) {
	o.OnEditForm = &v
}

// GetLookup returns the Lookup field value if set, zero value otherwise.
func (o *ExtensionPropertyUiDto) GetLookup() ExtensionPropertyUiLookupDto {
	if o == nil || o.Lookup == nil {
		var ret ExtensionPropertyUiLookupDto
		return ret
	}
	return *o.Lookup
}

// GetLookupOk returns a tuple with the Lookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyUiDto) GetLookupOk() (*ExtensionPropertyUiLookupDto, bool) {
	if o == nil || o.Lookup == nil {
		return nil, false
	}
	return o.Lookup, true
}

// HasLookup returns a boolean if a field has been set.
func (o *ExtensionPropertyUiDto) HasLookup() bool {
	if o != nil && o.Lookup != nil {
		return true
	}

	return false
}

// SetLookup gets a reference to the given ExtensionPropertyUiLookupDto and assigns it to the Lookup field.
func (o *ExtensionPropertyUiDto) SetLookup(v ExtensionPropertyUiLookupDto) {
	o.Lookup = &v
}

func (o ExtensionPropertyUiDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OnTable != nil {
		toSerialize["onTable"] = o.OnTable
	}
	if o.OnCreateForm != nil {
		toSerialize["onCreateForm"] = o.OnCreateForm
	}
	if o.OnEditForm != nil {
		toSerialize["onEditForm"] = o.OnEditForm
	}
	if o.Lookup != nil {
		toSerialize["lookup"] = o.Lookup
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionPropertyUiDto struct {
	value *ExtensionPropertyUiDto
	isSet bool
}

func (v NullableExtensionPropertyUiDto) Get() *ExtensionPropertyUiDto {
	return v.value
}

func (v *NullableExtensionPropertyUiDto) Set(val *ExtensionPropertyUiDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionPropertyUiDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionPropertyUiDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionPropertyUiDto(val *ExtensionPropertyUiDto) *NullableExtensionPropertyUiDto {
	return &NullableExtensionPropertyUiDto{value: val, isSet: true}
}

func (v NullableExtensionPropertyUiDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionPropertyUiDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


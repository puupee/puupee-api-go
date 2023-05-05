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

// checks if the ExtensionPropertyApiDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionPropertyApiDto{}

// ExtensionPropertyApiDto struct for ExtensionPropertyApiDto
type ExtensionPropertyApiDto struct {
	OnGet *ExtensionPropertyApiGetDto `json:"onGet,omitempty"`
	OnCreate *ExtensionPropertyApiCreateDto `json:"onCreate,omitempty"`
	OnUpdate *ExtensionPropertyApiUpdateDto `json:"onUpdate,omitempty"`
}

// NewExtensionPropertyApiDto instantiates a new ExtensionPropertyApiDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionPropertyApiDto() *ExtensionPropertyApiDto {
	this := ExtensionPropertyApiDto{}
	return &this
}

// NewExtensionPropertyApiDtoWithDefaults instantiates a new ExtensionPropertyApiDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionPropertyApiDtoWithDefaults() *ExtensionPropertyApiDto {
	this := ExtensionPropertyApiDto{}
	return &this
}

// GetOnGet returns the OnGet field value if set, zero value otherwise.
func (o *ExtensionPropertyApiDto) GetOnGet() ExtensionPropertyApiGetDto {
	if o == nil || IsNil(o.OnGet) {
		var ret ExtensionPropertyApiGetDto
		return ret
	}
	return *o.OnGet
}

// GetOnGetOk returns a tuple with the OnGet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyApiDto) GetOnGetOk() (*ExtensionPropertyApiGetDto, bool) {
	if o == nil || IsNil(o.OnGet) {
		return nil, false
	}
	return o.OnGet, true
}

// HasOnGet returns a boolean if a field has been set.
func (o *ExtensionPropertyApiDto) HasOnGet() bool {
	if o != nil && !IsNil(o.OnGet) {
		return true
	}

	return false
}

// SetOnGet gets a reference to the given ExtensionPropertyApiGetDto and assigns it to the OnGet field.
func (o *ExtensionPropertyApiDto) SetOnGet(v ExtensionPropertyApiGetDto) {
	o.OnGet = &v
}

// GetOnCreate returns the OnCreate field value if set, zero value otherwise.
func (o *ExtensionPropertyApiDto) GetOnCreate() ExtensionPropertyApiCreateDto {
	if o == nil || IsNil(o.OnCreate) {
		var ret ExtensionPropertyApiCreateDto
		return ret
	}
	return *o.OnCreate
}

// GetOnCreateOk returns a tuple with the OnCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyApiDto) GetOnCreateOk() (*ExtensionPropertyApiCreateDto, bool) {
	if o == nil || IsNil(o.OnCreate) {
		return nil, false
	}
	return o.OnCreate, true
}

// HasOnCreate returns a boolean if a field has been set.
func (o *ExtensionPropertyApiDto) HasOnCreate() bool {
	if o != nil && !IsNil(o.OnCreate) {
		return true
	}

	return false
}

// SetOnCreate gets a reference to the given ExtensionPropertyApiCreateDto and assigns it to the OnCreate field.
func (o *ExtensionPropertyApiDto) SetOnCreate(v ExtensionPropertyApiCreateDto) {
	o.OnCreate = &v
}

// GetOnUpdate returns the OnUpdate field value if set, zero value otherwise.
func (o *ExtensionPropertyApiDto) GetOnUpdate() ExtensionPropertyApiUpdateDto {
	if o == nil || IsNil(o.OnUpdate) {
		var ret ExtensionPropertyApiUpdateDto
		return ret
	}
	return *o.OnUpdate
}

// GetOnUpdateOk returns a tuple with the OnUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionPropertyApiDto) GetOnUpdateOk() (*ExtensionPropertyApiUpdateDto, bool) {
	if o == nil || IsNil(o.OnUpdate) {
		return nil, false
	}
	return o.OnUpdate, true
}

// HasOnUpdate returns a boolean if a field has been set.
func (o *ExtensionPropertyApiDto) HasOnUpdate() bool {
	if o != nil && !IsNil(o.OnUpdate) {
		return true
	}

	return false
}

// SetOnUpdate gets a reference to the given ExtensionPropertyApiUpdateDto and assigns it to the OnUpdate field.
func (o *ExtensionPropertyApiDto) SetOnUpdate(v ExtensionPropertyApiUpdateDto) {
	o.OnUpdate = &v
}

func (o ExtensionPropertyApiDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionPropertyApiDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnGet) {
		toSerialize["onGet"] = o.OnGet
	}
	if !IsNil(o.OnCreate) {
		toSerialize["onCreate"] = o.OnCreate
	}
	if !IsNil(o.OnUpdate) {
		toSerialize["onUpdate"] = o.OnUpdate
	}
	return toSerialize, nil
}

type NullableExtensionPropertyApiDto struct {
	value *ExtensionPropertyApiDto
	isSet bool
}

func (v NullableExtensionPropertyApiDto) Get() *ExtensionPropertyApiDto {
	return v.value
}

func (v *NullableExtensionPropertyApiDto) Set(val *ExtensionPropertyApiDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionPropertyApiDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionPropertyApiDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionPropertyApiDto(val *ExtensionPropertyApiDto) *NullableExtensionPropertyApiDto {
	return &NullableExtensionPropertyApiDto{value: val, isSet: true}
}

func (v NullableExtensionPropertyApiDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionPropertyApiDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



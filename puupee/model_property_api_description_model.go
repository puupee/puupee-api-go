/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// PropertyApiDescriptionModel struct for PropertyApiDescriptionModel
type PropertyApiDescriptionModel struct {
	Name *string `json:"name,omitempty"`
	JsonName *string `json:"jsonName,omitempty"`
	Type *string `json:"type,omitempty"`
	TypeSimple *string `json:"typeSimple,omitempty"`
	IsRequired *bool `json:"isRequired,omitempty"`
	MinLength *int32 `json:"minLength,omitempty"`
	MaxLength *int32 `json:"maxLength,omitempty"`
	Minimum *string `json:"minimum,omitempty"`
	Maximum *string `json:"maximum,omitempty"`
	Regex *string `json:"regex,omitempty"`
}

// NewPropertyApiDescriptionModel instantiates a new PropertyApiDescriptionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyApiDescriptionModel() *PropertyApiDescriptionModel {
	this := PropertyApiDescriptionModel{}
	return &this
}

// NewPropertyApiDescriptionModelWithDefaults instantiates a new PropertyApiDescriptionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyApiDescriptionModelWithDefaults() *PropertyApiDescriptionModel {
	this := PropertyApiDescriptionModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PropertyApiDescriptionModel) SetName(v string) {
	o.Name = &v
}

// GetJsonName returns the JsonName field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetJsonName() string {
	if o == nil || isNil(o.JsonName) {
		var ret string
		return ret
	}
	return *o.JsonName
}

// GetJsonNameOk returns a tuple with the JsonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetJsonNameOk() (*string, bool) {
	if o == nil || isNil(o.JsonName) {
    return nil, false
	}
	return o.JsonName, true
}

// HasJsonName returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasJsonName() bool {
	if o != nil && !isNil(o.JsonName) {
		return true
	}

	return false
}

// SetJsonName gets a reference to the given string and assigns it to the JsonName field.
func (o *PropertyApiDescriptionModel) SetJsonName(v string) {
	o.JsonName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PropertyApiDescriptionModel) SetType(v string) {
	o.Type = &v
}

// GetTypeSimple returns the TypeSimple field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetTypeSimple() string {
	if o == nil || isNil(o.TypeSimple) {
		var ret string
		return ret
	}
	return *o.TypeSimple
}

// GetTypeSimpleOk returns a tuple with the TypeSimple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetTypeSimpleOk() (*string, bool) {
	if o == nil || isNil(o.TypeSimple) {
    return nil, false
	}
	return o.TypeSimple, true
}

// HasTypeSimple returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasTypeSimple() bool {
	if o != nil && !isNil(o.TypeSimple) {
		return true
	}

	return false
}

// SetTypeSimple gets a reference to the given string and assigns it to the TypeSimple field.
func (o *PropertyApiDescriptionModel) SetTypeSimple(v string) {
	o.TypeSimple = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetIsRequired() bool {
	if o == nil || isNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetIsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.IsRequired) {
    return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasIsRequired() bool {
	if o != nil && !isNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *PropertyApiDescriptionModel) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetMinLength() int32 {
	if o == nil || isNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetMinLengthOk() (*int32, bool) {
	if o == nil || isNil(o.MinLength) {
    return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasMinLength() bool {
	if o != nil && !isNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *PropertyApiDescriptionModel) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetMaxLength() int32 {
	if o == nil || isNil(o.MaxLength) {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetMaxLengthOk() (*int32, bool) {
	if o == nil || isNil(o.MaxLength) {
    return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasMaxLength() bool {
	if o != nil && !isNil(o.MaxLength) {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *PropertyApiDescriptionModel) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetMinimum() string {
	if o == nil || isNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetMinimumOk() (*string, bool) {
	if o == nil || isNil(o.Minimum) {
    return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasMinimum() bool {
	if o != nil && !isNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *PropertyApiDescriptionModel) SetMinimum(v string) {
	o.Minimum = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetMaximum() string {
	if o == nil || isNil(o.Maximum) {
		var ret string
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetMaximumOk() (*string, bool) {
	if o == nil || isNil(o.Maximum) {
    return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasMaximum() bool {
	if o != nil && !isNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given string and assigns it to the Maximum field.
func (o *PropertyApiDescriptionModel) SetMaximum(v string) {
	o.Maximum = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *PropertyApiDescriptionModel) GetRegex() string {
	if o == nil || isNil(o.Regex) {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyApiDescriptionModel) GetRegexOk() (*string, bool) {
	if o == nil || isNil(o.Regex) {
    return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *PropertyApiDescriptionModel) HasRegex() bool {
	if o != nil && !isNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *PropertyApiDescriptionModel) SetRegex(v string) {
	o.Regex = &v
}

func (o PropertyApiDescriptionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.JsonName) {
		toSerialize["jsonName"] = o.JsonName
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.TypeSimple) {
		toSerialize["typeSimple"] = o.TypeSimple
	}
	if !isNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !isNil(o.MinLength) {
		toSerialize["minLength"] = o.MinLength
	}
	if !isNil(o.MaxLength) {
		toSerialize["maxLength"] = o.MaxLength
	}
	if !isNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}
	if !isNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	if !isNil(o.Regex) {
		toSerialize["regex"] = o.Regex
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyApiDescriptionModel struct {
	value *PropertyApiDescriptionModel
	isSet bool
}

func (v NullablePropertyApiDescriptionModel) Get() *PropertyApiDescriptionModel {
	return v.value
}

func (v *NullablePropertyApiDescriptionModel) Set(val *PropertyApiDescriptionModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyApiDescriptionModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyApiDescriptionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyApiDescriptionModel(val *PropertyApiDescriptionModel) *NullablePropertyApiDescriptionModel {
	return &NullablePropertyApiDescriptionModel{value: val, isSet: true}
}

func (v NullablePropertyApiDescriptionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyApiDescriptionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



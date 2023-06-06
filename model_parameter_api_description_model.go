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

// checks if the ParameterApiDescriptionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterApiDescriptionModel{}

// ParameterApiDescriptionModel struct for ParameterApiDescriptionModel
type ParameterApiDescriptionModel struct {
	NameOnMethod *string `json:"nameOnMethod,omitempty"`
	Name *string `json:"name,omitempty"`
	JsonName *string `json:"jsonName,omitempty"`
	Type *string `json:"type,omitempty"`
	TypeSimple *string `json:"typeSimple,omitempty"`
	IsOptional *bool `json:"isOptional,omitempty"`
	DefaultValue map[string]interface{} `json:"defaultValue,omitempty"`
	ConstraintTypes []string `json:"constraintTypes,omitempty"`
	BindingSourceId *string `json:"bindingSourceId,omitempty"`
	DescriptorName *string `json:"descriptorName,omitempty"`
}

// NewParameterApiDescriptionModel instantiates a new ParameterApiDescriptionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterApiDescriptionModel() *ParameterApiDescriptionModel {
	this := ParameterApiDescriptionModel{}
	return &this
}

// NewParameterApiDescriptionModelWithDefaults instantiates a new ParameterApiDescriptionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterApiDescriptionModelWithDefaults() *ParameterApiDescriptionModel {
	this := ParameterApiDescriptionModel{}
	return &this
}

// GetNameOnMethod returns the NameOnMethod field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetNameOnMethod() string {
	if o == nil || IsNil(o.NameOnMethod) {
		var ret string
		return ret
	}
	return *o.NameOnMethod
}

// GetNameOnMethodOk returns a tuple with the NameOnMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetNameOnMethodOk() (*string, bool) {
	if o == nil || IsNil(o.NameOnMethod) {
		return nil, false
	}
	return o.NameOnMethod, true
}

// HasNameOnMethod returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasNameOnMethod() bool {
	if o != nil && !IsNil(o.NameOnMethod) {
		return true
	}

	return false
}

// SetNameOnMethod gets a reference to the given string and assigns it to the NameOnMethod field.
func (o *ParameterApiDescriptionModel) SetNameOnMethod(v string) {
	o.NameOnMethod = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ParameterApiDescriptionModel) SetName(v string) {
	o.Name = &v
}

// GetJsonName returns the JsonName field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetJsonName() string {
	if o == nil || IsNil(o.JsonName) {
		var ret string
		return ret
	}
	return *o.JsonName
}

// GetJsonNameOk returns a tuple with the JsonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetJsonNameOk() (*string, bool) {
	if o == nil || IsNil(o.JsonName) {
		return nil, false
	}
	return o.JsonName, true
}

// HasJsonName returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasJsonName() bool {
	if o != nil && !IsNil(o.JsonName) {
		return true
	}

	return false
}

// SetJsonName gets a reference to the given string and assigns it to the JsonName field.
func (o *ParameterApiDescriptionModel) SetJsonName(v string) {
	o.JsonName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ParameterApiDescriptionModel) SetType(v string) {
	o.Type = &v
}

// GetTypeSimple returns the TypeSimple field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetTypeSimple() string {
	if o == nil || IsNil(o.TypeSimple) {
		var ret string
		return ret
	}
	return *o.TypeSimple
}

// GetTypeSimpleOk returns a tuple with the TypeSimple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetTypeSimpleOk() (*string, bool) {
	if o == nil || IsNil(o.TypeSimple) {
		return nil, false
	}
	return o.TypeSimple, true
}

// HasTypeSimple returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasTypeSimple() bool {
	if o != nil && !IsNil(o.TypeSimple) {
		return true
	}

	return false
}

// SetTypeSimple gets a reference to the given string and assigns it to the TypeSimple field.
func (o *ParameterApiDescriptionModel) SetTypeSimple(v string) {
	o.TypeSimple = &v
}

// GetIsOptional returns the IsOptional field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetIsOptional() bool {
	if o == nil || IsNil(o.IsOptional) {
		var ret bool
		return ret
	}
	return *o.IsOptional
}

// GetIsOptionalOk returns a tuple with the IsOptional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetIsOptionalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOptional) {
		return nil, false
	}
	return o.IsOptional, true
}

// HasIsOptional returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasIsOptional() bool {
	if o != nil && !IsNil(o.IsOptional) {
		return true
	}

	return false
}

// SetIsOptional gets a reference to the given bool and assigns it to the IsOptional field.
func (o *ParameterApiDescriptionModel) SetIsOptional(v bool) {
	o.IsOptional = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetDefaultValue() map[string]interface{} {
	if o == nil || IsNil(o.DefaultValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetDefaultValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return map[string]interface{}{}, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given map[string]interface{} and assigns it to the DefaultValue field.
func (o *ParameterApiDescriptionModel) SetDefaultValue(v map[string]interface{}) {
	o.DefaultValue = v
}

// GetConstraintTypes returns the ConstraintTypes field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetConstraintTypes() []string {
	if o == nil || IsNil(o.ConstraintTypes) {
		var ret []string
		return ret
	}
	return o.ConstraintTypes
}

// GetConstraintTypesOk returns a tuple with the ConstraintTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetConstraintTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ConstraintTypes) {
		return nil, false
	}
	return o.ConstraintTypes, true
}

// HasConstraintTypes returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasConstraintTypes() bool {
	if o != nil && !IsNil(o.ConstraintTypes) {
		return true
	}

	return false
}

// SetConstraintTypes gets a reference to the given []string and assigns it to the ConstraintTypes field.
func (o *ParameterApiDescriptionModel) SetConstraintTypes(v []string) {
	o.ConstraintTypes = v
}

// GetBindingSourceId returns the BindingSourceId field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetBindingSourceId() string {
	if o == nil || IsNil(o.BindingSourceId) {
		var ret string
		return ret
	}
	return *o.BindingSourceId
}

// GetBindingSourceIdOk returns a tuple with the BindingSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetBindingSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.BindingSourceId) {
		return nil, false
	}
	return o.BindingSourceId, true
}

// HasBindingSourceId returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasBindingSourceId() bool {
	if o != nil && !IsNil(o.BindingSourceId) {
		return true
	}

	return false
}

// SetBindingSourceId gets a reference to the given string and assigns it to the BindingSourceId field.
func (o *ParameterApiDescriptionModel) SetBindingSourceId(v string) {
	o.BindingSourceId = &v
}

// GetDescriptorName returns the DescriptorName field value if set, zero value otherwise.
func (o *ParameterApiDescriptionModel) GetDescriptorName() string {
	if o == nil || IsNil(o.DescriptorName) {
		var ret string
		return ret
	}
	return *o.DescriptorName
}

// GetDescriptorNameOk returns a tuple with the DescriptorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterApiDescriptionModel) GetDescriptorNameOk() (*string, bool) {
	if o == nil || IsNil(o.DescriptorName) {
		return nil, false
	}
	return o.DescriptorName, true
}

// HasDescriptorName returns a boolean if a field has been set.
func (o *ParameterApiDescriptionModel) HasDescriptorName() bool {
	if o != nil && !IsNil(o.DescriptorName) {
		return true
	}

	return false
}

// SetDescriptorName gets a reference to the given string and assigns it to the DescriptorName field.
func (o *ParameterApiDescriptionModel) SetDescriptorName(v string) {
	o.DescriptorName = &v
}

func (o ParameterApiDescriptionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterApiDescriptionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NameOnMethod) {
		toSerialize["nameOnMethod"] = o.NameOnMethod
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.JsonName) {
		toSerialize["jsonName"] = o.JsonName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeSimple) {
		toSerialize["typeSimple"] = o.TypeSimple
	}
	if !IsNil(o.IsOptional) {
		toSerialize["isOptional"] = o.IsOptional
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.ConstraintTypes) {
		toSerialize["constraintTypes"] = o.ConstraintTypes
	}
	if !IsNil(o.BindingSourceId) {
		toSerialize["bindingSourceId"] = o.BindingSourceId
	}
	if !IsNil(o.DescriptorName) {
		toSerialize["descriptorName"] = o.DescriptorName
	}
	return toSerialize, nil
}

type NullableParameterApiDescriptionModel struct {
	value *ParameterApiDescriptionModel
	isSet bool
}

func (v NullableParameterApiDescriptionModel) Get() *ParameterApiDescriptionModel {
	return v.value
}

func (v *NullableParameterApiDescriptionModel) Set(val *ParameterApiDescriptionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterApiDescriptionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterApiDescriptionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterApiDescriptionModel(val *ParameterApiDescriptionModel) *NullableParameterApiDescriptionModel {
	return &NullableParameterApiDescriptionModel{value: val, isSet: true}
}

func (v NullableParameterApiDescriptionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterApiDescriptionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



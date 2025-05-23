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

// checks if the FeatureDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureDto{}

// FeatureDto struct for FeatureDto
type FeatureDto struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Value *string `json:"value,omitempty"`
	Provider *FeatureProviderDto `json:"provider,omitempty"`
	Description *string `json:"description,omitempty"`
	ValueType *IStringValueType `json:"valueType,omitempty"`
	Depth *int32 `json:"depth,omitempty"`
	ParentName *string `json:"parentName,omitempty"`
}

// NewFeatureDto instantiates a new FeatureDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureDto() *FeatureDto {
	this := FeatureDto{}
	return &this
}

// NewFeatureDtoWithDefaults instantiates a new FeatureDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureDtoWithDefaults() *FeatureDto {
	this := FeatureDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeatureDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeatureDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeatureDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *FeatureDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FeatureDto) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FeatureDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FeatureDto) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDto) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FeatureDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FeatureDto) SetValue(v string) {
	o.Value = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *FeatureDto) GetProvider() FeatureProviderDto {
	if o == nil || IsNil(o.Provider) {
		var ret FeatureProviderDto
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDto) GetProviderOk() (*FeatureProviderDto, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *FeatureDto) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given FeatureProviderDto and assigns it to the Provider field.
func (o *FeatureDto) SetProvider(v FeatureProviderDto) {
	o.Provider = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FeatureDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FeatureDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FeatureDto) SetDescription(v string) {
	o.Description = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *FeatureDto) GetValueType() IStringValueType {
	if o == nil || IsNil(o.ValueType) {
		var ret IStringValueType
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDto) GetValueTypeOk() (*IStringValueType, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *FeatureDto) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given IStringValueType and assigns it to the ValueType field.
func (o *FeatureDto) SetValueType(v IStringValueType) {
	o.ValueType = &v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *FeatureDto) GetDepth() int32 {
	if o == nil || IsNil(o.Depth) {
		var ret int32
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDto) GetDepthOk() (*int32, bool) {
	if o == nil || IsNil(o.Depth) {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *FeatureDto) HasDepth() bool {
	if o != nil && !IsNil(o.Depth) {
		return true
	}

	return false
}

// SetDepth gets a reference to the given int32 and assigns it to the Depth field.
func (o *FeatureDto) SetDepth(v int32) {
	o.Depth = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *FeatureDto) GetParentName() string {
	if o == nil || IsNil(o.ParentName) {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDto) GetParentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentName) {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *FeatureDto) HasParentName() bool {
	if o != nil && !IsNil(o.ParentName) {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *FeatureDto) SetParentName(v string) {
	o.ParentName = &v
}

func (o FeatureDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	if !IsNil(o.Depth) {
		toSerialize["depth"] = o.Depth
	}
	if !IsNil(o.ParentName) {
		toSerialize["parentName"] = o.ParentName
	}
	return toSerialize, nil
}

type NullableFeatureDto struct {
	value *FeatureDto
	isSet bool
}

func (v NullableFeatureDto) Get() *FeatureDto {
	return v.value
}

func (v *NullableFeatureDto) Set(val *FeatureDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureDto(val *FeatureDto) *NullableFeatureDto {
	return &NullableFeatureDto{value: val, isSet: true}
}

func (v NullableFeatureDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



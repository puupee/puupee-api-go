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

// checks if the CreateUpdateMessageSourceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUpdateMessageSourceDto{}

// CreateUpdateMessageSourceDto struct for CreateUpdateMessageSourceDto
type CreateUpdateMessageSourceDto struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPublished *bool `json:"isPublished,omitempty"`
	IconUrl *string `json:"iconUrl,omitempty"`
	Routes []CreateUpdateMessageSourceRouteSubDto `json:"routes,omitempty"`
}

// NewCreateUpdateMessageSourceDto instantiates a new CreateUpdateMessageSourceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpdateMessageSourceDto() *CreateUpdateMessageSourceDto {
	this := CreateUpdateMessageSourceDto{}
	return &this
}

// NewCreateUpdateMessageSourceDtoWithDefaults instantiates a new CreateUpdateMessageSourceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpdateMessageSourceDtoWithDefaults() *CreateUpdateMessageSourceDto {
	this := CreateUpdateMessageSourceDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateUpdateMessageSourceDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateUpdateMessageSourceDto) SetDescription(v string) {
	o.Description = &v
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceDto) GetIsPublished() bool {
	if o == nil || IsNil(o.IsPublished) {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceDto) GetIsPublishedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublished) {
		return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceDto) HasIsPublished() bool {
	if o != nil && !IsNil(o.IsPublished) {
		return true
	}

	return false
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *CreateUpdateMessageSourceDto) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceDto) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceDto) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceDto) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *CreateUpdateMessageSourceDto) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceDto) GetRoutes() []CreateUpdateMessageSourceRouteSubDto {
	if o == nil || IsNil(o.Routes) {
		var ret []CreateUpdateMessageSourceRouteSubDto
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceDto) GetRoutesOk() ([]CreateUpdateMessageSourceRouteSubDto, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceDto) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []CreateUpdateMessageSourceRouteSubDto and assigns it to the Routes field.
func (o *CreateUpdateMessageSourceDto) SetRoutes(v []CreateUpdateMessageSourceRouteSubDto) {
	o.Routes = v
}

func (o CreateUpdateMessageSourceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUpdateMessageSourceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsPublished) {
		toSerialize["isPublished"] = o.IsPublished
	}
	if !IsNil(o.IconUrl) {
		toSerialize["iconUrl"] = o.IconUrl
	}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	return toSerialize, nil
}

type NullableCreateUpdateMessageSourceDto struct {
	value *CreateUpdateMessageSourceDto
	isSet bool
}

func (v NullableCreateUpdateMessageSourceDto) Get() *CreateUpdateMessageSourceDto {
	return v.value
}

func (v *NullableCreateUpdateMessageSourceDto) Set(val *CreateUpdateMessageSourceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpdateMessageSourceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpdateMessageSourceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpdateMessageSourceDto(val *CreateUpdateMessageSourceDto) *NullableCreateUpdateMessageSourceDto {
	return &NullableCreateUpdateMessageSourceDto{value: val, isSet: true}
}

func (v NullableCreateUpdateMessageSourceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpdateMessageSourceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



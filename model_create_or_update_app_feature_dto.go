/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// CreateOrUpdateAppFeatureDto struct for CreateOrUpdateAppFeatureDto
type CreateOrUpdateAppFeatureDto struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Description *string `json:"description,omitempty"`
	Details *string `json:"details,omitempty"`
	ScreenshotKeys *string `json:"screenshotKeys,omitempty"`
}

// NewCreateOrUpdateAppFeatureDto instantiates a new CreateOrUpdateAppFeatureDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateAppFeatureDto() *CreateOrUpdateAppFeatureDto {
	this := CreateOrUpdateAppFeatureDto{}
	return &this
}

// NewCreateOrUpdateAppFeatureDtoWithDefaults instantiates a new CreateOrUpdateAppFeatureDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateAppFeatureDtoWithDefaults() *CreateOrUpdateAppFeatureDto {
	this := CreateOrUpdateAppFeatureDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrUpdateAppFeatureDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppFeatureDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrUpdateAppFeatureDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrUpdateAppFeatureDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateOrUpdateAppFeatureDto) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppFeatureDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateOrUpdateAppFeatureDto) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateOrUpdateAppFeatureDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdateAppFeatureDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppFeatureDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateAppFeatureDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrUpdateAppFeatureDto) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CreateOrUpdateAppFeatureDto) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppFeatureDto) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CreateOrUpdateAppFeatureDto) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *CreateOrUpdateAppFeatureDto) SetDetails(v string) {
	o.Details = &v
}

// GetScreenshotKeys returns the ScreenshotKeys field value if set, zero value otherwise.
func (o *CreateOrUpdateAppFeatureDto) GetScreenshotKeys() string {
	if o == nil || isNil(o.ScreenshotKeys) {
		var ret string
		return ret
	}
	return *o.ScreenshotKeys
}

// GetScreenshotKeysOk returns a tuple with the ScreenshotKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppFeatureDto) GetScreenshotKeysOk() (*string, bool) {
	if o == nil || isNil(o.ScreenshotKeys) {
    return nil, false
	}
	return o.ScreenshotKeys, true
}

// HasScreenshotKeys returns a boolean if a field has been set.
func (o *CreateOrUpdateAppFeatureDto) HasScreenshotKeys() bool {
	if o != nil && !isNil(o.ScreenshotKeys) {
		return true
	}

	return false
}

// SetScreenshotKeys gets a reference to the given string and assigns it to the ScreenshotKeys field.
func (o *CreateOrUpdateAppFeatureDto) SetScreenshotKeys(v string) {
	o.ScreenshotKeys = &v
}

func (o CreateOrUpdateAppFeatureDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !isNil(o.ScreenshotKeys) {
		toSerialize["screenshotKeys"] = o.ScreenshotKeys
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrUpdateAppFeatureDto struct {
	value *CreateOrUpdateAppFeatureDto
	isSet bool
}

func (v NullableCreateOrUpdateAppFeatureDto) Get() *CreateOrUpdateAppFeatureDto {
	return v.value
}

func (v *NullableCreateOrUpdateAppFeatureDto) Set(val *CreateOrUpdateAppFeatureDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateAppFeatureDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateAppFeatureDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateAppFeatureDto(val *CreateOrUpdateAppFeatureDto) *NullableCreateOrUpdateAppFeatureDto {
	return &NullableCreateOrUpdateAppFeatureDto{value: val, isSet: true}
}

func (v NullableCreateOrUpdateAppFeatureDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateAppFeatureDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



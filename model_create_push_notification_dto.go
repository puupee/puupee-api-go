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

// checks if the CreatePushNotificationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePushNotificationDto{}

// CreatePushNotificationDto struct for CreatePushNotificationDto
type CreatePushNotificationDto struct {
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	PuupeeId *string `json:"puupeeId,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	App *string `json:"app,omitempty"`
}

// NewCreatePushNotificationDto instantiates a new CreatePushNotificationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePushNotificationDto() *CreatePushNotificationDto {
	this := CreatePushNotificationDto{}
	return &this
}

// NewCreatePushNotificationDtoWithDefaults instantiates a new CreatePushNotificationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePushNotificationDtoWithDefaults() *CreatePushNotificationDto {
	this := CreatePushNotificationDto{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreatePushNotificationDto) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushNotificationDto) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreatePushNotificationDto) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreatePushNotificationDto) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreatePushNotificationDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushNotificationDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreatePushNotificationDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreatePushNotificationDto) SetDescription(v string) {
	o.Description = &v
}

// GetPuupeeId returns the PuupeeId field value if set, zero value otherwise.
func (o *CreatePushNotificationDto) GetPuupeeId() string {
	if o == nil || IsNil(o.PuupeeId) {
		var ret string
		return ret
	}
	return *o.PuupeeId
}

// GetPuupeeIdOk returns a tuple with the PuupeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushNotificationDto) GetPuupeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PuupeeId) {
		return nil, false
	}
	return o.PuupeeId, true
}

// HasPuupeeId returns a boolean if a field has been set.
func (o *CreatePushNotificationDto) HasPuupeeId() bool {
	if o != nil && !IsNil(o.PuupeeId) {
		return true
	}

	return false
}

// SetPuupeeId gets a reference to the given string and assigns it to the PuupeeId field.
func (o *CreatePushNotificationDto) SetPuupeeId(v string) {
	o.PuupeeId = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *CreatePushNotificationDto) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushNotificationDto) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *CreatePushNotificationDto) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *CreatePushNotificationDto) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *CreatePushNotificationDto) GetApp() string {
	if o == nil || IsNil(o.App) {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushNotificationDto) GetAppOk() (*string, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *CreatePushNotificationDto) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *CreatePushNotificationDto) SetApp(v string) {
	o.App = &v
}

func (o CreatePushNotificationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePushNotificationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PuupeeId) {
		toSerialize["puupeeId"] = o.PuupeeId
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	return toSerialize, nil
}

type NullableCreatePushNotificationDto struct {
	value *CreatePushNotificationDto
	isSet bool
}

func (v NullableCreatePushNotificationDto) Get() *CreatePushNotificationDto {
	return v.value
}

func (v *NullableCreatePushNotificationDto) Set(val *CreatePushNotificationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePushNotificationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePushNotificationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePushNotificationDto(val *CreatePushNotificationDto) *NullableCreatePushNotificationDto {
	return &NullableCreatePushNotificationDto{value: val, isSet: true}
}

func (v NullableCreatePushNotificationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePushNotificationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



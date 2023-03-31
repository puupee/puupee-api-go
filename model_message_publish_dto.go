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

// MessagePublishDto struct for MessagePublishDto
type MessagePublishDto struct {
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	AppId *string `json:"appId,omitempty"`
	Template *string `json:"template,omitempty"`
	Data map[string]map[string]interface{} `json:"data,omitempty"`
}

// NewMessagePublishDto instantiates a new MessagePublishDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagePublishDto() *MessagePublishDto {
	this := MessagePublishDto{}
	return &this
}

// NewMessagePublishDtoWithDefaults instantiates a new MessagePublishDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagePublishDtoWithDefaults() *MessagePublishDto {
	this := MessagePublishDto{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MessagePublishDto) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagePublishDto) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MessagePublishDto) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MessagePublishDto) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MessagePublishDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagePublishDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MessagePublishDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MessagePublishDto) SetDescription(v string) {
	o.Description = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *MessagePublishDto) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagePublishDto) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *MessagePublishDto) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *MessagePublishDto) SetAppId(v string) {
	o.AppId = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *MessagePublishDto) GetTemplate() string {
	if o == nil || isNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagePublishDto) GetTemplateOk() (*string, bool) {
	if o == nil || isNil(o.Template) {
    return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *MessagePublishDto) HasTemplate() bool {
	if o != nil && !isNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *MessagePublishDto) SetTemplate(v string) {
	o.Template = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MessagePublishDto) GetData() map[string]map[string]interface{} {
	if o == nil || isNil(o.Data) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagePublishDto) GetDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.Data) {
    return map[string]map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MessagePublishDto) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]map[string]interface{} and assigns it to the Data field.
func (o *MessagePublishDto) SetData(v map[string]map[string]interface{}) {
	o.Data = v
}

func (o MessagePublishDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMessagePublishDto struct {
	value *MessagePublishDto
	isSet bool
}

func (v NullableMessagePublishDto) Get() *MessagePublishDto {
	return v.value
}

func (v *NullableMessagePublishDto) Set(val *MessagePublishDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagePublishDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagePublishDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagePublishDto(val *MessagePublishDto) *NullableMessagePublishDto {
	return &NullableMessagePublishDto{value: val, isSet: true}
}

func (v NullableMessagePublishDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessagePublishDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



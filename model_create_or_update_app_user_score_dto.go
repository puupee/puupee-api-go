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

// checks if the CreateOrUpdateAppUserScoreDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateAppUserScoreDto{}

// CreateOrUpdateAppUserScoreDto struct for CreateOrUpdateAppUserScoreDto
type CreateOrUpdateAppUserScoreDto struct {
	AppId NullableString `json:"appId,omitempty"`
	Score *int32 `json:"score,omitempty"`
	Comment NullableString `json:"comment,omitempty"`
}

// NewCreateOrUpdateAppUserScoreDto instantiates a new CreateOrUpdateAppUserScoreDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateAppUserScoreDto() *CreateOrUpdateAppUserScoreDto {
	this := CreateOrUpdateAppUserScoreDto{}
	return &this
}

// NewCreateOrUpdateAppUserScoreDtoWithDefaults instantiates a new CreateOrUpdateAppUserScoreDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateAppUserScoreDtoWithDefaults() *CreateOrUpdateAppUserScoreDto {
	this := CreateOrUpdateAppUserScoreDto{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppUserScoreDto) GetAppId() string {
	if o == nil || IsNil(o.AppId.Get()) {
		var ret string
		return ret
	}
	return *o.AppId.Get()
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppUserScoreDto) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppId.Get(), o.AppId.IsSet()
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateOrUpdateAppUserScoreDto) HasAppId() bool {
	if o != nil && o.AppId.IsSet() {
		return true
	}

	return false
}

// SetAppId gets a reference to the given NullableString and assigns it to the AppId field.
func (o *CreateOrUpdateAppUserScoreDto) SetAppId(v string) {
	o.AppId.Set(&v)
}
// SetAppIdNil sets the value for AppId to be an explicit nil
func (o *CreateOrUpdateAppUserScoreDto) SetAppIdNil() {
	o.AppId.Set(nil)
}

// UnsetAppId ensures that no value is present for AppId, not even an explicit nil
func (o *CreateOrUpdateAppUserScoreDto) UnsetAppId() {
	o.AppId.Unset()
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *CreateOrUpdateAppUserScoreDto) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppUserScoreDto) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *CreateOrUpdateAppUserScoreDto) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *CreateOrUpdateAppUserScoreDto) SetScore(v int32) {
	o.Score = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateAppUserScoreDto) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateAppUserScoreDto) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *CreateOrUpdateAppUserScoreDto) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *CreateOrUpdateAppUserScoreDto) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *CreateOrUpdateAppUserScoreDto) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *CreateOrUpdateAppUserScoreDto) UnsetComment() {
	o.Comment.Unset()
}

func (o CreateOrUpdateAppUserScoreDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateAppUserScoreDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId.IsSet() {
		toSerialize["appId"] = o.AppId.Get()
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return toSerialize, nil
}

type NullableCreateOrUpdateAppUserScoreDto struct {
	value *CreateOrUpdateAppUserScoreDto
	isSet bool
}

func (v NullableCreateOrUpdateAppUserScoreDto) Get() *CreateOrUpdateAppUserScoreDto {
	return v.value
}

func (v *NullableCreateOrUpdateAppUserScoreDto) Set(val *CreateOrUpdateAppUserScoreDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateAppUserScoreDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateAppUserScoreDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateAppUserScoreDto(val *CreateOrUpdateAppUserScoreDto) *NullableCreateOrUpdateAppUserScoreDto {
	return &NullableCreateOrUpdateAppUserScoreDto{value: val, isSet: true}
}

func (v NullableCreateOrUpdateAppUserScoreDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateAppUserScoreDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



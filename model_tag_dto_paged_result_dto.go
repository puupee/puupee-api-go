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

// TagDtoPagedResultDto struct for TagDtoPagedResultDto
type TagDtoPagedResultDto struct {
	Items []TagDto `json:"items,omitempty"`
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// NewTagDtoPagedResultDto instantiates a new TagDtoPagedResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagDtoPagedResultDto() *TagDtoPagedResultDto {
	this := TagDtoPagedResultDto{}
	return &this
}

// NewTagDtoPagedResultDtoWithDefaults instantiates a new TagDtoPagedResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagDtoPagedResultDtoWithDefaults() *TagDtoPagedResultDto {
	this := TagDtoPagedResultDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDtoPagedResultDto) GetItems() []TagDto {
	if o == nil  {
		var ret []TagDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDtoPagedResultDto) GetItemsOk() (*[]TagDto, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TagDtoPagedResultDto) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TagDto and assigns it to the Items field.
func (o *TagDtoPagedResultDto) SetItems(v []TagDto) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *TagDtoPagedResultDto) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDtoPagedResultDto) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *TagDtoPagedResultDto) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *TagDtoPagedResultDto) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o TagDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableTagDtoPagedResultDto struct {
	value *TagDtoPagedResultDto
	isSet bool
}

func (v NullableTagDtoPagedResultDto) Get() *TagDtoPagedResultDto {
	return v.value
}

func (v *NullableTagDtoPagedResultDto) Set(val *TagDtoPagedResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTagDtoPagedResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTagDtoPagedResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagDtoPagedResultDto(val *TagDtoPagedResultDto) *NullableTagDtoPagedResultDto {
	return &NullableTagDtoPagedResultDto{value: val, isSet: true}
}

func (v NullableTagDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagDtoPagedResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



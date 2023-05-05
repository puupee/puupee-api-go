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

// checks if the NotificationInfoDtoPagedResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationInfoDtoPagedResultDto{}

// NotificationInfoDtoPagedResultDto struct for NotificationInfoDtoPagedResultDto
type NotificationInfoDtoPagedResultDto struct {
	Items []NotificationInfoDto `json:"items,omitempty"`
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// NewNotificationInfoDtoPagedResultDto instantiates a new NotificationInfoDtoPagedResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationInfoDtoPagedResultDto() *NotificationInfoDtoPagedResultDto {
	this := NotificationInfoDtoPagedResultDto{}
	return &this
}

// NewNotificationInfoDtoPagedResultDtoWithDefaults instantiates a new NotificationInfoDtoPagedResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationInfoDtoPagedResultDtoWithDefaults() *NotificationInfoDtoPagedResultDto {
	this := NotificationInfoDtoPagedResultDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *NotificationInfoDtoPagedResultDto) GetItems() []NotificationInfoDto {
	if o == nil || IsNil(o.Items) {
		var ret []NotificationInfoDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationInfoDtoPagedResultDto) GetItemsOk() ([]NotificationInfoDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *NotificationInfoDtoPagedResultDto) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []NotificationInfoDto and assigns it to the Items field.
func (o *NotificationInfoDtoPagedResultDto) SetItems(v []NotificationInfoDto) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *NotificationInfoDtoPagedResultDto) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationInfoDtoPagedResultDto) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *NotificationInfoDtoPagedResultDto) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *NotificationInfoDtoPagedResultDto) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o NotificationInfoDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationInfoDtoPagedResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableNotificationInfoDtoPagedResultDto struct {
	value *NotificationInfoDtoPagedResultDto
	isSet bool
}

func (v NullableNotificationInfoDtoPagedResultDto) Get() *NotificationInfoDtoPagedResultDto {
	return v.value
}

func (v *NullableNotificationInfoDtoPagedResultDto) Set(val *NotificationInfoDtoPagedResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationInfoDtoPagedResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationInfoDtoPagedResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationInfoDtoPagedResultDto(val *NotificationInfoDtoPagedResultDto) *NullableNotificationInfoDtoPagedResultDto {
	return &NullableNotificationInfoDtoPagedResultDto{value: val, isSet: true}
}

func (v NullableNotificationInfoDtoPagedResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationInfoDtoPagedResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



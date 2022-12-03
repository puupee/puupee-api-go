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

// UserStorageDto struct for UserStorageDto
type UserStorageDto struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	MaxSize *int64 `json:"maxSize,omitempty"`
	CurrentSize *int64 `json:"currentSize,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
	Items []UserStorageItemDto `json:"items,omitempty"`
}

// NewUserStorageDto instantiates a new UserStorageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStorageDto() *UserStorageDto {
	this := UserStorageDto{}
	return &this
}

// NewUserStorageDtoWithDefaults instantiates a new UserStorageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStorageDtoWithDefaults() *UserStorageDto {
	this := UserStorageDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserStorageDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserStorageDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserStorageDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserStorageDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserStorageDto) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserStorageDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserStorageDto) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserStorageDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *UserStorageDto) GetMaxSize() int64 {
	if o == nil || isNil(o.MaxSize) {
		var ret int64
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserStorageDto) GetMaxSizeOk() (*int64, bool) {
	if o == nil || isNil(o.MaxSize) {
    return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *UserStorageDto) HasMaxSize() bool {
	if o != nil && !isNil(o.MaxSize) {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given int64 and assigns it to the MaxSize field.
func (o *UserStorageDto) SetMaxSize(v int64) {
	o.MaxSize = &v
}

// GetCurrentSize returns the CurrentSize field value if set, zero value otherwise.
func (o *UserStorageDto) GetCurrentSize() int64 {
	if o == nil || isNil(o.CurrentSize) {
		var ret int64
		return ret
	}
	return *o.CurrentSize
}

// GetCurrentSizeOk returns a tuple with the CurrentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserStorageDto) GetCurrentSizeOk() (*int64, bool) {
	if o == nil || isNil(o.CurrentSize) {
    return nil, false
	}
	return o.CurrentSize, true
}

// HasCurrentSize returns a boolean if a field has been set.
func (o *UserStorageDto) HasCurrentSize() bool {
	if o != nil && !isNil(o.CurrentSize) {
		return true
	}

	return false
}

// SetCurrentSize gets a reference to the given int64 and assigns it to the CurrentSize field.
func (o *UserStorageDto) SetCurrentSize(v int64) {
	o.CurrentSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *UserStorageDto) GetTotalCount() int32 {
	if o == nil || isNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserStorageDto) GetTotalCountOk() (*int32, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *UserStorageDto) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *UserStorageDto) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UserStorageDto) GetItems() []UserStorageItemDto {
	if o == nil || isNil(o.Items) {
		var ret []UserStorageItemDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserStorageDto) GetItemsOk() ([]UserStorageItemDto, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UserStorageDto) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UserStorageItemDto and assigns it to the Items field.
func (o *UserStorageDto) SetItems(v []UserStorageItemDto) {
	o.Items = v
}

func (o UserStorageDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.MaxSize) {
		toSerialize["maxSize"] = o.MaxSize
	}
	if !isNil(o.CurrentSize) {
		toSerialize["currentSize"] = o.CurrentSize
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableUserStorageDto struct {
	value *UserStorageDto
	isSet bool
}

func (v NullableUserStorageDto) Get() *UserStorageDto {
	return v.value
}

func (v *NullableUserStorageDto) Set(val *UserStorageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStorageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStorageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStorageDto(val *UserStorageDto) *NullableUserStorageDto {
	return &NullableUserStorageDto{value: val, isSet: true}
}

func (v NullableUserStorageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStorageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


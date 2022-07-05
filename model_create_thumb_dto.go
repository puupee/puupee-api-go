/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
	"time"
)

// CreateThumbDto struct for CreateThumbDto
type CreateThumbDto struct {
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Key NullableString `json:"key,omitempty"`
	Md5 NullableString `json:"md5,omitempty"`
	ContentType NullableString `json:"contentType,omitempty"`
	Extension NullableString `json:"extension,omitempty"`
	StorageClass NullableString `json:"storageClass,omitempty"`
	FileCreatedAt NullableTime `json:"fileCreatedAt,omitempty"`
	FileUpdatedAt NullableTime `json:"fileUpdatedAt,omitempty"`
}

// NewCreateThumbDto instantiates a new CreateThumbDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateThumbDto() *CreateThumbDto {
	this := CreateThumbDto{}
	return &this
}

// NewCreateThumbDtoWithDefaults instantiates a new CreateThumbDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateThumbDtoWithDefaults() *CreateThumbDto {
	this := CreateThumbDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateThumbDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateThumbDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateThumbDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateThumbDto) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbDto) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateThumbDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateThumbDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateThumbDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateThumbDto) UnsetDescription() {
	o.Description.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbDto) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *CreateThumbDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *CreateThumbDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *CreateThumbDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *CreateThumbDto) UnsetKey() {
	o.Key.Unset()
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbDto) GetMd5() string {
	if o == nil || o.Md5.Get() == nil {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbDto) GetMd5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *CreateThumbDto) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *CreateThumbDto) SetMd5(v string) {
	o.Md5.Set(&v)
}
// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *CreateThumbDto) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *CreateThumbDto) UnsetMd5() {
	o.Md5.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbDto) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbDto) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *CreateThumbDto) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *CreateThumbDto) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *CreateThumbDto) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *CreateThumbDto) UnsetContentType() {
	o.ContentType.Unset()
}

// GetExtension returns the Extension field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbDto) GetExtension() string {
	if o == nil || o.Extension.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extension.Get()
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbDto) GetExtensionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Extension.Get(), o.Extension.IsSet()
}

// HasExtension returns a boolean if a field has been set.
func (o *CreateThumbDto) HasExtension() bool {
	if o != nil && o.Extension.IsSet() {
		return true
	}

	return false
}

// SetExtension gets a reference to the given NullableString and assigns it to the Extension field.
func (o *CreateThumbDto) SetExtension(v string) {
	o.Extension.Set(&v)
}
// SetExtensionNil sets the value for Extension to be an explicit nil
func (o *CreateThumbDto) SetExtensionNil() {
	o.Extension.Set(nil)
}

// UnsetExtension ensures that no value is present for Extension, not even an explicit nil
func (o *CreateThumbDto) UnsetExtension() {
	o.Extension.Unset()
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbDto) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbDto) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// HasStorageClass returns a boolean if a field has been set.
func (o *CreateThumbDto) HasStorageClass() bool {
	if o != nil && o.StorageClass.IsSet() {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given NullableString and assigns it to the StorageClass field.
func (o *CreateThumbDto) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}
// SetStorageClassNil sets the value for StorageClass to be an explicit nil
func (o *CreateThumbDto) SetStorageClassNil() {
	o.StorageClass.Set(nil)
}

// UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
func (o *CreateThumbDto) UnsetStorageClass() {
	o.StorageClass.Unset()
}

// GetFileCreatedAt returns the FileCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbDto) GetFileCreatedAt() time.Time {
	if o == nil || o.FileCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FileCreatedAt.Get()
}

// GetFileCreatedAtOk returns a tuple with the FileCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbDto) GetFileCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileCreatedAt.Get(), o.FileCreatedAt.IsSet()
}

// HasFileCreatedAt returns a boolean if a field has been set.
func (o *CreateThumbDto) HasFileCreatedAt() bool {
	if o != nil && o.FileCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetFileCreatedAt gets a reference to the given NullableTime and assigns it to the FileCreatedAt field.
func (o *CreateThumbDto) SetFileCreatedAt(v time.Time) {
	o.FileCreatedAt.Set(&v)
}
// SetFileCreatedAtNil sets the value for FileCreatedAt to be an explicit nil
func (o *CreateThumbDto) SetFileCreatedAtNil() {
	o.FileCreatedAt.Set(nil)
}

// UnsetFileCreatedAt ensures that no value is present for FileCreatedAt, not even an explicit nil
func (o *CreateThumbDto) UnsetFileCreatedAt() {
	o.FileCreatedAt.Unset()
}

// GetFileUpdatedAt returns the FileUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThumbDto) GetFileUpdatedAt() time.Time {
	if o == nil || o.FileUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FileUpdatedAt.Get()
}

// GetFileUpdatedAtOk returns a tuple with the FileUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThumbDto) GetFileUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileUpdatedAt.Get(), o.FileUpdatedAt.IsSet()
}

// HasFileUpdatedAt returns a boolean if a field has been set.
func (o *CreateThumbDto) HasFileUpdatedAt() bool {
	if o != nil && o.FileUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetFileUpdatedAt gets a reference to the given NullableTime and assigns it to the FileUpdatedAt field.
func (o *CreateThumbDto) SetFileUpdatedAt(v time.Time) {
	o.FileUpdatedAt.Set(&v)
}
// SetFileUpdatedAtNil sets the value for FileUpdatedAt to be an explicit nil
func (o *CreateThumbDto) SetFileUpdatedAtNil() {
	o.FileUpdatedAt.Set(nil)
}

// UnsetFileUpdatedAt ensures that no value is present for FileUpdatedAt, not even an explicit nil
func (o *CreateThumbDto) UnsetFileUpdatedAt() {
	o.FileUpdatedAt.Unset()
}

func (o CreateThumbDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Md5.IsSet() {
		toSerialize["md5"] = o.Md5.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	if o.Extension.IsSet() {
		toSerialize["extension"] = o.Extension.Get()
	}
	if o.StorageClass.IsSet() {
		toSerialize["storageClass"] = o.StorageClass.Get()
	}
	if o.FileCreatedAt.IsSet() {
		toSerialize["fileCreatedAt"] = o.FileCreatedAt.Get()
	}
	if o.FileUpdatedAt.IsSet() {
		toSerialize["fileUpdatedAt"] = o.FileUpdatedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateThumbDto struct {
	value *CreateThumbDto
	isSet bool
}

func (v NullableCreateThumbDto) Get() *CreateThumbDto {
	return v.value
}

func (v *NullableCreateThumbDto) Set(val *CreateThumbDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateThumbDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateThumbDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateThumbDto(val *CreateThumbDto) *NullableCreateThumbDto {
	return &NullableCreateThumbDto{value: val, isSet: true}
}

func (v NullableCreateThumbDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateThumbDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


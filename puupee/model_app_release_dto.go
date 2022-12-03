/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// AppReleaseDto struct for AppReleaseDto
type AppReleaseDto struct {
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	LastModificationTime *time.Time `json:"lastModificationTime,omitempty"`
	LastModifierId *string `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId *string `json:"deleterId,omitempty"`
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	Version *string `json:"version,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Key *string `json:"key,omitempty"`
	RapidCode *string `json:"rapidCode,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Md5 *string `json:"md5,omitempty"`
	SliceMd5 *string `json:"sliceMd5,omitempty"`
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	ProductType *string `json:"productType,omitempty"`
	IsForceUpdate *bool `json:"isForceUpdate,omitempty"`
	AppId *string `json:"appId,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	Channel *string `json:"channel,omitempty"`
	Environment *string `json:"environment,omitempty"`
}

// NewAppReleaseDto instantiates a new AppReleaseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppReleaseDto() *AppReleaseDto {
	this := AppReleaseDto{}
	return &this
}

// NewAppReleaseDtoWithDefaults instantiates a new AppReleaseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppReleaseDtoWithDefaults() *AppReleaseDto {
	this := AppReleaseDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppReleaseDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppReleaseDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppReleaseDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AppReleaseDto) GetCreationTime() time.Time {
	if o == nil || isNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreationTime) {
    return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AppReleaseDto) HasCreationTime() bool {
	if o != nil && !isNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *AppReleaseDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *AppReleaseDto) GetCreatorId() string {
	if o == nil || isNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetCreatorIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatorId) {
    return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *AppReleaseDto) HasCreatorId() bool {
	if o != nil && !isNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *AppReleaseDto) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *AppReleaseDto) GetLastModificationTime() time.Time {
	if o == nil || isNil(o.LastModificationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastModificationTime) {
    return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *AppReleaseDto) HasLastModificationTime() bool {
	if o != nil && !isNil(o.LastModificationTime) {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given time.Time and assigns it to the LastModificationTime field.
func (o *AppReleaseDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *AppReleaseDto) GetLastModifierId() string {
	if o == nil || isNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil || isNil(o.LastModifierId) {
    return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *AppReleaseDto) HasLastModifierId() bool {
	if o != nil && !isNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *AppReleaseDto) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AppReleaseDto) GetIsDeleted() bool {
	if o == nil || isNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDeleted) {
    return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AppReleaseDto) HasIsDeleted() bool {
	if o != nil && !isNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AppReleaseDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise.
func (o *AppReleaseDto) GetDeleterId() string {
	if o == nil || isNil(o.DeleterId) {
		var ret string
		return ret
	}
	return *o.DeleterId
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetDeleterIdOk() (*string, bool) {
	if o == nil || isNil(o.DeleterId) {
    return nil, false
	}
	return o.DeleterId, true
}

// HasDeleterId returns a boolean if a field has been set.
func (o *AppReleaseDto) HasDeleterId() bool {
	if o != nil && !isNil(o.DeleterId) {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given string and assigns it to the DeleterId field.
func (o *AppReleaseDto) SetDeleterId(v string) {
	o.DeleterId = &v
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise.
func (o *AppReleaseDto) GetDeletionTime() time.Time {
	if o == nil || isNil(o.DeletionTime) {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletionTime) {
    return nil, false
	}
	return o.DeletionTime, true
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *AppReleaseDto) HasDeletionTime() bool {
	if o != nil && !isNil(o.DeletionTime) {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given time.Time and assigns it to the DeletionTime field.
func (o *AppReleaseDto) SetDeletionTime(v time.Time) {
	o.DeletionTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AppReleaseDto) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AppReleaseDto) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AppReleaseDto) SetVersion(v string) {
	o.Version = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AppReleaseDto) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AppReleaseDto) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AppReleaseDto) SetNotes(v string) {
	o.Notes = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AppReleaseDto) GetPlatform() string {
	if o == nil || isNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetPlatformOk() (*string, bool) {
	if o == nil || isNil(o.Platform) {
    return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AppReleaseDto) HasPlatform() bool {
	if o != nil && !isNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *AppReleaseDto) SetPlatform(v string) {
	o.Platform = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AppReleaseDto) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AppReleaseDto) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AppReleaseDto) SetKey(v string) {
	o.Key = &v
}

// GetRapidCode returns the RapidCode field value if set, zero value otherwise.
func (o *AppReleaseDto) GetRapidCode() string {
	if o == nil || isNil(o.RapidCode) {
		var ret string
		return ret
	}
	return *o.RapidCode
}

// GetRapidCodeOk returns a tuple with the RapidCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetRapidCodeOk() (*string, bool) {
	if o == nil || isNil(o.RapidCode) {
    return nil, false
	}
	return o.RapidCode, true
}

// HasRapidCode returns a boolean if a field has been set.
func (o *AppReleaseDto) HasRapidCode() bool {
	if o != nil && !isNil(o.RapidCode) {
		return true
	}

	return false
}

// SetRapidCode gets a reference to the given string and assigns it to the RapidCode field.
func (o *AppReleaseDto) SetRapidCode(v string) {
	o.RapidCode = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *AppReleaseDto) GetSize() int64 {
	if o == nil || isNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetSizeOk() (*int64, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *AppReleaseDto) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *AppReleaseDto) SetSize(v int64) {
	o.Size = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *AppReleaseDto) GetMd5() string {
	if o == nil || isNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetMd5Ok() (*string, bool) {
	if o == nil || isNil(o.Md5) {
    return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *AppReleaseDto) HasMd5() bool {
	if o != nil && !isNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *AppReleaseDto) SetMd5(v string) {
	o.Md5 = &v
}

// GetSliceMd5 returns the SliceMd5 field value if set, zero value otherwise.
func (o *AppReleaseDto) GetSliceMd5() string {
	if o == nil || isNil(o.SliceMd5) {
		var ret string
		return ret
	}
	return *o.SliceMd5
}

// GetSliceMd5Ok returns a tuple with the SliceMd5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetSliceMd5Ok() (*string, bool) {
	if o == nil || isNil(o.SliceMd5) {
    return nil, false
	}
	return o.SliceMd5, true
}

// HasSliceMd5 returns a boolean if a field has been set.
func (o *AppReleaseDto) HasSliceMd5() bool {
	if o != nil && !isNil(o.SliceMd5) {
		return true
	}

	return false
}

// SetSliceMd5 gets a reference to the given string and assigns it to the SliceMd5 field.
func (o *AppReleaseDto) SetSliceMd5(v string) {
	o.SliceMd5 = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *AppReleaseDto) GetDownloadUrl() string {
	if o == nil || isNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetDownloadUrlOk() (*string, bool) {
	if o == nil || isNil(o.DownloadUrl) {
    return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *AppReleaseDto) HasDownloadUrl() bool {
	if o != nil && !isNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *AppReleaseDto) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *AppReleaseDto) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *AppReleaseDto) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *AppReleaseDto) SetProductType(v string) {
	o.ProductType = &v
}

// GetIsForceUpdate returns the IsForceUpdate field value if set, zero value otherwise.
func (o *AppReleaseDto) GetIsForceUpdate() bool {
	if o == nil || isNil(o.IsForceUpdate) {
		var ret bool
		return ret
	}
	return *o.IsForceUpdate
}

// GetIsForceUpdateOk returns a tuple with the IsForceUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetIsForceUpdateOk() (*bool, bool) {
	if o == nil || isNil(o.IsForceUpdate) {
    return nil, false
	}
	return o.IsForceUpdate, true
}

// HasIsForceUpdate returns a boolean if a field has been set.
func (o *AppReleaseDto) HasIsForceUpdate() bool {
	if o != nil && !isNil(o.IsForceUpdate) {
		return true
	}

	return false
}

// SetIsForceUpdate gets a reference to the given bool and assigns it to the IsForceUpdate field.
func (o *AppReleaseDto) SetIsForceUpdate(v bool) {
	o.IsForceUpdate = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AppReleaseDto) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AppReleaseDto) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *AppReleaseDto) SetAppId(v string) {
	o.AppId = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *AppReleaseDto) GetIsEnabled() bool {
	if o == nil || isNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsEnabled) {
    return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *AppReleaseDto) HasIsEnabled() bool {
	if o != nil && !isNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *AppReleaseDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *AppReleaseDto) GetChannel() string {
	if o == nil || isNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetChannelOk() (*string, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *AppReleaseDto) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *AppReleaseDto) SetChannel(v string) {
	o.Channel = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AppReleaseDto) GetEnvironment() string {
	if o == nil || isNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppReleaseDto) GetEnvironmentOk() (*string, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AppReleaseDto) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *AppReleaseDto) SetEnvironment(v string) {
	o.Environment = &v
}

func (o AppReleaseDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !isNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if !isNil(o.LastModificationTime) {
		toSerialize["lastModificationTime"] = o.LastModificationTime
	}
	if !isNil(o.LastModifierId) {
		toSerialize["lastModifierId"] = o.LastModifierId
	}
	if !isNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !isNil(o.DeleterId) {
		toSerialize["deleterId"] = o.DeleterId
	}
	if !isNil(o.DeletionTime) {
		toSerialize["deletionTime"] = o.DeletionTime
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.RapidCode) {
		toSerialize["rapidCode"] = o.RapidCode
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !isNil(o.SliceMd5) {
		toSerialize["sliceMd5"] = o.SliceMd5
	}
	if !isNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.IsForceUpdate) {
		toSerialize["isForceUpdate"] = o.IsForceUpdate
	}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableAppReleaseDto struct {
	value *AppReleaseDto
	isSet bool
}

func (v NullableAppReleaseDto) Get() *AppReleaseDto {
	return v.value
}

func (v *NullableAppReleaseDto) Set(val *AppReleaseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppReleaseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppReleaseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppReleaseDto(val *AppReleaseDto) *NullableAppReleaseDto {
	return &NullableAppReleaseDto{value: val, isSet: true}
}

func (v NullableAppReleaseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppReleaseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


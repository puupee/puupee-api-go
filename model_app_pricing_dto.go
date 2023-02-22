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

// AppPricingDto struct for AppPricingDto
type AppPricingDto struct {
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	LastModificationTime *time.Time `json:"lastModificationTime,omitempty"`
	LastModifierId *string `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId *string `json:"deleterId,omitempty"`
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	Naming *string `json:"naming,omitempty"`
	Description *string `json:"description,omitempty"`
	AppId *string `json:"appId,omitempty"`
	MonthPrice *float64 `json:"monthPrice,omitempty"`
	MonthDiscount *float64 `json:"monthDiscount,omitempty"`
	MonthDiscountPrice *float64 `json:"monthDiscountPrice,omitempty"`
	MonthDiscountStartAt *time.Time `json:"monthDiscountStartAt,omitempty"`
	MonthDiscountEndAt *time.Time `json:"monthDiscountEndAt,omitempty"`
	YearPrice *float64 `json:"yearPrice,omitempty"`
	YearDiscount *float64 `json:"yearDiscount,omitempty"`
	YearDiscountPrice *float64 `json:"yearDiscountPrice,omitempty"`
	YearDiscountStartAt *time.Time `json:"yearDiscountStartAt,omitempty"`
	YearDiscountEndAt *time.Time `json:"yearDiscountEndAt,omitempty"`
	Items []AppPricingItemDto `json:"items,omitempty"`
}

// NewAppPricingDto instantiates a new AppPricingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricingDto() *AppPricingDto {
	this := AppPricingDto{}
	return &this
}

// NewAppPricingDtoWithDefaults instantiates a new AppPricingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricingDtoWithDefaults() *AppPricingDto {
	this := AppPricingDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppPricingDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppPricingDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppPricingDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AppPricingDto) GetCreationTime() time.Time {
	if o == nil || isNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreationTime) {
    return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AppPricingDto) HasCreationTime() bool {
	if o != nil && !isNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *AppPricingDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *AppPricingDto) GetCreatorId() string {
	if o == nil || isNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetCreatorIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatorId) {
    return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *AppPricingDto) HasCreatorId() bool {
	if o != nil && !isNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *AppPricingDto) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *AppPricingDto) GetLastModificationTime() time.Time {
	if o == nil || isNil(o.LastModificationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastModificationTime) {
    return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *AppPricingDto) HasLastModificationTime() bool {
	if o != nil && !isNil(o.LastModificationTime) {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given time.Time and assigns it to the LastModificationTime field.
func (o *AppPricingDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *AppPricingDto) GetLastModifierId() string {
	if o == nil || isNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil || isNil(o.LastModifierId) {
    return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *AppPricingDto) HasLastModifierId() bool {
	if o != nil && !isNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *AppPricingDto) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AppPricingDto) GetIsDeleted() bool {
	if o == nil || isNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDeleted) {
    return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AppPricingDto) HasIsDeleted() bool {
	if o != nil && !isNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AppPricingDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise.
func (o *AppPricingDto) GetDeleterId() string {
	if o == nil || isNil(o.DeleterId) {
		var ret string
		return ret
	}
	return *o.DeleterId
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetDeleterIdOk() (*string, bool) {
	if o == nil || isNil(o.DeleterId) {
    return nil, false
	}
	return o.DeleterId, true
}

// HasDeleterId returns a boolean if a field has been set.
func (o *AppPricingDto) HasDeleterId() bool {
	if o != nil && !isNil(o.DeleterId) {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given string and assigns it to the DeleterId field.
func (o *AppPricingDto) SetDeleterId(v string) {
	o.DeleterId = &v
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise.
func (o *AppPricingDto) GetDeletionTime() time.Time {
	if o == nil || isNil(o.DeletionTime) {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletionTime) {
    return nil, false
	}
	return o.DeletionTime, true
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *AppPricingDto) HasDeletionTime() bool {
	if o != nil && !isNil(o.DeletionTime) {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given time.Time and assigns it to the DeletionTime field.
func (o *AppPricingDto) SetDeletionTime(v time.Time) {
	o.DeletionTime = &v
}

// GetNaming returns the Naming field value if set, zero value otherwise.
func (o *AppPricingDto) GetNaming() string {
	if o == nil || isNil(o.Naming) {
		var ret string
		return ret
	}
	return *o.Naming
}

// GetNamingOk returns a tuple with the Naming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetNamingOk() (*string, bool) {
	if o == nil || isNil(o.Naming) {
    return nil, false
	}
	return o.Naming, true
}

// HasNaming returns a boolean if a field has been set.
func (o *AppPricingDto) HasNaming() bool {
	if o != nil && !isNil(o.Naming) {
		return true
	}

	return false
}

// SetNaming gets a reference to the given string and assigns it to the Naming field.
func (o *AppPricingDto) SetNaming(v string) {
	o.Naming = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppPricingDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppPricingDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppPricingDto) SetDescription(v string) {
	o.Description = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AppPricingDto) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AppPricingDto) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *AppPricingDto) SetAppId(v string) {
	o.AppId = &v
}

// GetMonthPrice returns the MonthPrice field value if set, zero value otherwise.
func (o *AppPricingDto) GetMonthPrice() float64 {
	if o == nil || isNil(o.MonthPrice) {
		var ret float64
		return ret
	}
	return *o.MonthPrice
}

// GetMonthPriceOk returns a tuple with the MonthPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetMonthPriceOk() (*float64, bool) {
	if o == nil || isNil(o.MonthPrice) {
    return nil, false
	}
	return o.MonthPrice, true
}

// HasMonthPrice returns a boolean if a field has been set.
func (o *AppPricingDto) HasMonthPrice() bool {
	if o != nil && !isNil(o.MonthPrice) {
		return true
	}

	return false
}

// SetMonthPrice gets a reference to the given float64 and assigns it to the MonthPrice field.
func (o *AppPricingDto) SetMonthPrice(v float64) {
	o.MonthPrice = &v
}

// GetMonthDiscount returns the MonthDiscount field value if set, zero value otherwise.
func (o *AppPricingDto) GetMonthDiscount() float64 {
	if o == nil || isNil(o.MonthDiscount) {
		var ret float64
		return ret
	}
	return *o.MonthDiscount
}

// GetMonthDiscountOk returns a tuple with the MonthDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetMonthDiscountOk() (*float64, bool) {
	if o == nil || isNil(o.MonthDiscount) {
    return nil, false
	}
	return o.MonthDiscount, true
}

// HasMonthDiscount returns a boolean if a field has been set.
func (o *AppPricingDto) HasMonthDiscount() bool {
	if o != nil && !isNil(o.MonthDiscount) {
		return true
	}

	return false
}

// SetMonthDiscount gets a reference to the given float64 and assigns it to the MonthDiscount field.
func (o *AppPricingDto) SetMonthDiscount(v float64) {
	o.MonthDiscount = &v
}

// GetMonthDiscountPrice returns the MonthDiscountPrice field value if set, zero value otherwise.
func (o *AppPricingDto) GetMonthDiscountPrice() float64 {
	if o == nil || isNil(o.MonthDiscountPrice) {
		var ret float64
		return ret
	}
	return *o.MonthDiscountPrice
}

// GetMonthDiscountPriceOk returns a tuple with the MonthDiscountPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetMonthDiscountPriceOk() (*float64, bool) {
	if o == nil || isNil(o.MonthDiscountPrice) {
    return nil, false
	}
	return o.MonthDiscountPrice, true
}

// HasMonthDiscountPrice returns a boolean if a field has been set.
func (o *AppPricingDto) HasMonthDiscountPrice() bool {
	if o != nil && !isNil(o.MonthDiscountPrice) {
		return true
	}

	return false
}

// SetMonthDiscountPrice gets a reference to the given float64 and assigns it to the MonthDiscountPrice field.
func (o *AppPricingDto) SetMonthDiscountPrice(v float64) {
	o.MonthDiscountPrice = &v
}

// GetMonthDiscountStartAt returns the MonthDiscountStartAt field value if set, zero value otherwise.
func (o *AppPricingDto) GetMonthDiscountStartAt() time.Time {
	if o == nil || isNil(o.MonthDiscountStartAt) {
		var ret time.Time
		return ret
	}
	return *o.MonthDiscountStartAt
}

// GetMonthDiscountStartAtOk returns a tuple with the MonthDiscountStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetMonthDiscountStartAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.MonthDiscountStartAt) {
    return nil, false
	}
	return o.MonthDiscountStartAt, true
}

// HasMonthDiscountStartAt returns a boolean if a field has been set.
func (o *AppPricingDto) HasMonthDiscountStartAt() bool {
	if o != nil && !isNil(o.MonthDiscountStartAt) {
		return true
	}

	return false
}

// SetMonthDiscountStartAt gets a reference to the given time.Time and assigns it to the MonthDiscountStartAt field.
func (o *AppPricingDto) SetMonthDiscountStartAt(v time.Time) {
	o.MonthDiscountStartAt = &v
}

// GetMonthDiscountEndAt returns the MonthDiscountEndAt field value if set, zero value otherwise.
func (o *AppPricingDto) GetMonthDiscountEndAt() time.Time {
	if o == nil || isNil(o.MonthDiscountEndAt) {
		var ret time.Time
		return ret
	}
	return *o.MonthDiscountEndAt
}

// GetMonthDiscountEndAtOk returns a tuple with the MonthDiscountEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetMonthDiscountEndAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.MonthDiscountEndAt) {
    return nil, false
	}
	return o.MonthDiscountEndAt, true
}

// HasMonthDiscountEndAt returns a boolean if a field has been set.
func (o *AppPricingDto) HasMonthDiscountEndAt() bool {
	if o != nil && !isNil(o.MonthDiscountEndAt) {
		return true
	}

	return false
}

// SetMonthDiscountEndAt gets a reference to the given time.Time and assigns it to the MonthDiscountEndAt field.
func (o *AppPricingDto) SetMonthDiscountEndAt(v time.Time) {
	o.MonthDiscountEndAt = &v
}

// GetYearPrice returns the YearPrice field value if set, zero value otherwise.
func (o *AppPricingDto) GetYearPrice() float64 {
	if o == nil || isNil(o.YearPrice) {
		var ret float64
		return ret
	}
	return *o.YearPrice
}

// GetYearPriceOk returns a tuple with the YearPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetYearPriceOk() (*float64, bool) {
	if o == nil || isNil(o.YearPrice) {
    return nil, false
	}
	return o.YearPrice, true
}

// HasYearPrice returns a boolean if a field has been set.
func (o *AppPricingDto) HasYearPrice() bool {
	if o != nil && !isNil(o.YearPrice) {
		return true
	}

	return false
}

// SetYearPrice gets a reference to the given float64 and assigns it to the YearPrice field.
func (o *AppPricingDto) SetYearPrice(v float64) {
	o.YearPrice = &v
}

// GetYearDiscount returns the YearDiscount field value if set, zero value otherwise.
func (o *AppPricingDto) GetYearDiscount() float64 {
	if o == nil || isNil(o.YearDiscount) {
		var ret float64
		return ret
	}
	return *o.YearDiscount
}

// GetYearDiscountOk returns a tuple with the YearDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetYearDiscountOk() (*float64, bool) {
	if o == nil || isNil(o.YearDiscount) {
    return nil, false
	}
	return o.YearDiscount, true
}

// HasYearDiscount returns a boolean if a field has been set.
func (o *AppPricingDto) HasYearDiscount() bool {
	if o != nil && !isNil(o.YearDiscount) {
		return true
	}

	return false
}

// SetYearDiscount gets a reference to the given float64 and assigns it to the YearDiscount field.
func (o *AppPricingDto) SetYearDiscount(v float64) {
	o.YearDiscount = &v
}

// GetYearDiscountPrice returns the YearDiscountPrice field value if set, zero value otherwise.
func (o *AppPricingDto) GetYearDiscountPrice() float64 {
	if o == nil || isNil(o.YearDiscountPrice) {
		var ret float64
		return ret
	}
	return *o.YearDiscountPrice
}

// GetYearDiscountPriceOk returns a tuple with the YearDiscountPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetYearDiscountPriceOk() (*float64, bool) {
	if o == nil || isNil(o.YearDiscountPrice) {
    return nil, false
	}
	return o.YearDiscountPrice, true
}

// HasYearDiscountPrice returns a boolean if a field has been set.
func (o *AppPricingDto) HasYearDiscountPrice() bool {
	if o != nil && !isNil(o.YearDiscountPrice) {
		return true
	}

	return false
}

// SetYearDiscountPrice gets a reference to the given float64 and assigns it to the YearDiscountPrice field.
func (o *AppPricingDto) SetYearDiscountPrice(v float64) {
	o.YearDiscountPrice = &v
}

// GetYearDiscountStartAt returns the YearDiscountStartAt field value if set, zero value otherwise.
func (o *AppPricingDto) GetYearDiscountStartAt() time.Time {
	if o == nil || isNil(o.YearDiscountStartAt) {
		var ret time.Time
		return ret
	}
	return *o.YearDiscountStartAt
}

// GetYearDiscountStartAtOk returns a tuple with the YearDiscountStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetYearDiscountStartAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.YearDiscountStartAt) {
    return nil, false
	}
	return o.YearDiscountStartAt, true
}

// HasYearDiscountStartAt returns a boolean if a field has been set.
func (o *AppPricingDto) HasYearDiscountStartAt() bool {
	if o != nil && !isNil(o.YearDiscountStartAt) {
		return true
	}

	return false
}

// SetYearDiscountStartAt gets a reference to the given time.Time and assigns it to the YearDiscountStartAt field.
func (o *AppPricingDto) SetYearDiscountStartAt(v time.Time) {
	o.YearDiscountStartAt = &v
}

// GetYearDiscountEndAt returns the YearDiscountEndAt field value if set, zero value otherwise.
func (o *AppPricingDto) GetYearDiscountEndAt() time.Time {
	if o == nil || isNil(o.YearDiscountEndAt) {
		var ret time.Time
		return ret
	}
	return *o.YearDiscountEndAt
}

// GetYearDiscountEndAtOk returns a tuple with the YearDiscountEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetYearDiscountEndAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.YearDiscountEndAt) {
    return nil, false
	}
	return o.YearDiscountEndAt, true
}

// HasYearDiscountEndAt returns a boolean if a field has been set.
func (o *AppPricingDto) HasYearDiscountEndAt() bool {
	if o != nil && !isNil(o.YearDiscountEndAt) {
		return true
	}

	return false
}

// SetYearDiscountEndAt gets a reference to the given time.Time and assigns it to the YearDiscountEndAt field.
func (o *AppPricingDto) SetYearDiscountEndAt(v time.Time) {
	o.YearDiscountEndAt = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AppPricingDto) GetItems() []AppPricingItemDto {
	if o == nil || isNil(o.Items) {
		var ret []AppPricingItemDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricingDto) GetItemsOk() ([]AppPricingItemDto, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AppPricingDto) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AppPricingItemDto and assigns it to the Items field.
func (o *AppPricingDto) SetItems(v []AppPricingItemDto) {
	o.Items = v
}

func (o AppPricingDto) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Naming) {
		toSerialize["naming"] = o.Naming
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.MonthPrice) {
		toSerialize["monthPrice"] = o.MonthPrice
	}
	if !isNil(o.MonthDiscount) {
		toSerialize["monthDiscount"] = o.MonthDiscount
	}
	if !isNil(o.MonthDiscountPrice) {
		toSerialize["monthDiscountPrice"] = o.MonthDiscountPrice
	}
	if !isNil(o.MonthDiscountStartAt) {
		toSerialize["monthDiscountStartAt"] = o.MonthDiscountStartAt
	}
	if !isNil(o.MonthDiscountEndAt) {
		toSerialize["monthDiscountEndAt"] = o.MonthDiscountEndAt
	}
	if !isNil(o.YearPrice) {
		toSerialize["yearPrice"] = o.YearPrice
	}
	if !isNil(o.YearDiscount) {
		toSerialize["yearDiscount"] = o.YearDiscount
	}
	if !isNil(o.YearDiscountPrice) {
		toSerialize["yearDiscountPrice"] = o.YearDiscountPrice
	}
	if !isNil(o.YearDiscountStartAt) {
		toSerialize["yearDiscountStartAt"] = o.YearDiscountStartAt
	}
	if !isNil(o.YearDiscountEndAt) {
		toSerialize["yearDiscountEndAt"] = o.YearDiscountEndAt
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableAppPricingDto struct {
	value *AppPricingDto
	isSet bool
}

func (v NullableAppPricingDto) Get() *AppPricingDto {
	return v.value
}

func (v *NullableAppPricingDto) Set(val *AppPricingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricingDto(val *AppPricingDto) *NullableAppPricingDto {
	return &NullableAppPricingDto{value: val, isSet: true}
}

func (v NullableAppPricingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


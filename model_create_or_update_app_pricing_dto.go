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

// CreateOrUpdateAppPricingDto struct for CreateOrUpdateAppPricingDto
type CreateOrUpdateAppPricingDto struct {
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

// NewCreateOrUpdateAppPricingDto instantiates a new CreateOrUpdateAppPricingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateAppPricingDto() *CreateOrUpdateAppPricingDto {
	this := CreateOrUpdateAppPricingDto{}
	return &this
}

// NewCreateOrUpdateAppPricingDtoWithDefaults instantiates a new CreateOrUpdateAppPricingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateAppPricingDtoWithDefaults() *CreateOrUpdateAppPricingDto {
	this := CreateOrUpdateAppPricingDto{}
	return &this
}

// GetNaming returns the Naming field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetNaming() string {
	if o == nil || isNil(o.Naming) {
		var ret string
		return ret
	}
	return *o.Naming
}

// GetNamingOk returns a tuple with the Naming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetNamingOk() (*string, bool) {
	if o == nil || isNil(o.Naming) {
    return nil, false
	}
	return o.Naming, true
}

// HasNaming returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasNaming() bool {
	if o != nil && !isNil(o.Naming) {
		return true
	}

	return false
}

// SetNaming gets a reference to the given string and assigns it to the Naming field.
func (o *CreateOrUpdateAppPricingDto) SetNaming(v string) {
	o.Naming = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrUpdateAppPricingDto) SetDescription(v string) {
	o.Description = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *CreateOrUpdateAppPricingDto) SetAppId(v string) {
	o.AppId = &v
}

// GetMonthPrice returns the MonthPrice field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetMonthPrice() float64 {
	if o == nil || isNil(o.MonthPrice) {
		var ret float64
		return ret
	}
	return *o.MonthPrice
}

// GetMonthPriceOk returns a tuple with the MonthPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetMonthPriceOk() (*float64, bool) {
	if o == nil || isNil(o.MonthPrice) {
    return nil, false
	}
	return o.MonthPrice, true
}

// HasMonthPrice returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasMonthPrice() bool {
	if o != nil && !isNil(o.MonthPrice) {
		return true
	}

	return false
}

// SetMonthPrice gets a reference to the given float64 and assigns it to the MonthPrice field.
func (o *CreateOrUpdateAppPricingDto) SetMonthPrice(v float64) {
	o.MonthPrice = &v
}

// GetMonthDiscount returns the MonthDiscount field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetMonthDiscount() float64 {
	if o == nil || isNil(o.MonthDiscount) {
		var ret float64
		return ret
	}
	return *o.MonthDiscount
}

// GetMonthDiscountOk returns a tuple with the MonthDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountOk() (*float64, bool) {
	if o == nil || isNil(o.MonthDiscount) {
    return nil, false
	}
	return o.MonthDiscount, true
}

// HasMonthDiscount returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasMonthDiscount() bool {
	if o != nil && !isNil(o.MonthDiscount) {
		return true
	}

	return false
}

// SetMonthDiscount gets a reference to the given float64 and assigns it to the MonthDiscount field.
func (o *CreateOrUpdateAppPricingDto) SetMonthDiscount(v float64) {
	o.MonthDiscount = &v
}

// GetMonthDiscountPrice returns the MonthDiscountPrice field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountPrice() float64 {
	if o == nil || isNil(o.MonthDiscountPrice) {
		var ret float64
		return ret
	}
	return *o.MonthDiscountPrice
}

// GetMonthDiscountPriceOk returns a tuple with the MonthDiscountPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountPriceOk() (*float64, bool) {
	if o == nil || isNil(o.MonthDiscountPrice) {
    return nil, false
	}
	return o.MonthDiscountPrice, true
}

// HasMonthDiscountPrice returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasMonthDiscountPrice() bool {
	if o != nil && !isNil(o.MonthDiscountPrice) {
		return true
	}

	return false
}

// SetMonthDiscountPrice gets a reference to the given float64 and assigns it to the MonthDiscountPrice field.
func (o *CreateOrUpdateAppPricingDto) SetMonthDiscountPrice(v float64) {
	o.MonthDiscountPrice = &v
}

// GetMonthDiscountStartAt returns the MonthDiscountStartAt field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountStartAt() time.Time {
	if o == nil || isNil(o.MonthDiscountStartAt) {
		var ret time.Time
		return ret
	}
	return *o.MonthDiscountStartAt
}

// GetMonthDiscountStartAtOk returns a tuple with the MonthDiscountStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountStartAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.MonthDiscountStartAt) {
    return nil, false
	}
	return o.MonthDiscountStartAt, true
}

// HasMonthDiscountStartAt returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasMonthDiscountStartAt() bool {
	if o != nil && !isNil(o.MonthDiscountStartAt) {
		return true
	}

	return false
}

// SetMonthDiscountStartAt gets a reference to the given time.Time and assigns it to the MonthDiscountStartAt field.
func (o *CreateOrUpdateAppPricingDto) SetMonthDiscountStartAt(v time.Time) {
	o.MonthDiscountStartAt = &v
}

// GetMonthDiscountEndAt returns the MonthDiscountEndAt field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountEndAt() time.Time {
	if o == nil || isNil(o.MonthDiscountEndAt) {
		var ret time.Time
		return ret
	}
	return *o.MonthDiscountEndAt
}

// GetMonthDiscountEndAtOk returns a tuple with the MonthDiscountEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountEndAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.MonthDiscountEndAt) {
    return nil, false
	}
	return o.MonthDiscountEndAt, true
}

// HasMonthDiscountEndAt returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasMonthDiscountEndAt() bool {
	if o != nil && !isNil(o.MonthDiscountEndAt) {
		return true
	}

	return false
}

// SetMonthDiscountEndAt gets a reference to the given time.Time and assigns it to the MonthDiscountEndAt field.
func (o *CreateOrUpdateAppPricingDto) SetMonthDiscountEndAt(v time.Time) {
	o.MonthDiscountEndAt = &v
}

// GetYearPrice returns the YearPrice field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetYearPrice() float64 {
	if o == nil || isNil(o.YearPrice) {
		var ret float64
		return ret
	}
	return *o.YearPrice
}

// GetYearPriceOk returns a tuple with the YearPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetYearPriceOk() (*float64, bool) {
	if o == nil || isNil(o.YearPrice) {
    return nil, false
	}
	return o.YearPrice, true
}

// HasYearPrice returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasYearPrice() bool {
	if o != nil && !isNil(o.YearPrice) {
		return true
	}

	return false
}

// SetYearPrice gets a reference to the given float64 and assigns it to the YearPrice field.
func (o *CreateOrUpdateAppPricingDto) SetYearPrice(v float64) {
	o.YearPrice = &v
}

// GetYearDiscount returns the YearDiscount field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetYearDiscount() float64 {
	if o == nil || isNil(o.YearDiscount) {
		var ret float64
		return ret
	}
	return *o.YearDiscount
}

// GetYearDiscountOk returns a tuple with the YearDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetYearDiscountOk() (*float64, bool) {
	if o == nil || isNil(o.YearDiscount) {
    return nil, false
	}
	return o.YearDiscount, true
}

// HasYearDiscount returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasYearDiscount() bool {
	if o != nil && !isNil(o.YearDiscount) {
		return true
	}

	return false
}

// SetYearDiscount gets a reference to the given float64 and assigns it to the YearDiscount field.
func (o *CreateOrUpdateAppPricingDto) SetYearDiscount(v float64) {
	o.YearDiscount = &v
}

// GetYearDiscountPrice returns the YearDiscountPrice field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetYearDiscountPrice() float64 {
	if o == nil || isNil(o.YearDiscountPrice) {
		var ret float64
		return ret
	}
	return *o.YearDiscountPrice
}

// GetYearDiscountPriceOk returns a tuple with the YearDiscountPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetYearDiscountPriceOk() (*float64, bool) {
	if o == nil || isNil(o.YearDiscountPrice) {
    return nil, false
	}
	return o.YearDiscountPrice, true
}

// HasYearDiscountPrice returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasYearDiscountPrice() bool {
	if o != nil && !isNil(o.YearDiscountPrice) {
		return true
	}

	return false
}

// SetYearDiscountPrice gets a reference to the given float64 and assigns it to the YearDiscountPrice field.
func (o *CreateOrUpdateAppPricingDto) SetYearDiscountPrice(v float64) {
	o.YearDiscountPrice = &v
}

// GetYearDiscountStartAt returns the YearDiscountStartAt field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetYearDiscountStartAt() time.Time {
	if o == nil || isNil(o.YearDiscountStartAt) {
		var ret time.Time
		return ret
	}
	return *o.YearDiscountStartAt
}

// GetYearDiscountStartAtOk returns a tuple with the YearDiscountStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetYearDiscountStartAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.YearDiscountStartAt) {
    return nil, false
	}
	return o.YearDiscountStartAt, true
}

// HasYearDiscountStartAt returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasYearDiscountStartAt() bool {
	if o != nil && !isNil(o.YearDiscountStartAt) {
		return true
	}

	return false
}

// SetYearDiscountStartAt gets a reference to the given time.Time and assigns it to the YearDiscountStartAt field.
func (o *CreateOrUpdateAppPricingDto) SetYearDiscountStartAt(v time.Time) {
	o.YearDiscountStartAt = &v
}

// GetYearDiscountEndAt returns the YearDiscountEndAt field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetYearDiscountEndAt() time.Time {
	if o == nil || isNil(o.YearDiscountEndAt) {
		var ret time.Time
		return ret
	}
	return *o.YearDiscountEndAt
}

// GetYearDiscountEndAtOk returns a tuple with the YearDiscountEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetYearDiscountEndAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.YearDiscountEndAt) {
    return nil, false
	}
	return o.YearDiscountEndAt, true
}

// HasYearDiscountEndAt returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasYearDiscountEndAt() bool {
	if o != nil && !isNil(o.YearDiscountEndAt) {
		return true
	}

	return false
}

// SetYearDiscountEndAt gets a reference to the given time.Time and assigns it to the YearDiscountEndAt field.
func (o *CreateOrUpdateAppPricingDto) SetYearDiscountEndAt(v time.Time) {
	o.YearDiscountEndAt = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *CreateOrUpdateAppPricingDto) GetItems() []AppPricingItemDto {
	if o == nil || isNil(o.Items) {
		var ret []AppPricingItemDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppPricingDto) GetItemsOk() ([]AppPricingItemDto, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *CreateOrUpdateAppPricingDto) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AppPricingItemDto and assigns it to the Items field.
func (o *CreateOrUpdateAppPricingDto) SetItems(v []AppPricingItemDto) {
	o.Items = v
}

func (o CreateOrUpdateAppPricingDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableCreateOrUpdateAppPricingDto struct {
	value *CreateOrUpdateAppPricingDto
	isSet bool
}

func (v NullableCreateOrUpdateAppPricingDto) Get() *CreateOrUpdateAppPricingDto {
	return v.value
}

func (v *NullableCreateOrUpdateAppPricingDto) Set(val *CreateOrUpdateAppPricingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateAppPricingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateAppPricingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateAppPricingDto(val *CreateOrUpdateAppPricingDto) *NullableCreateOrUpdateAppPricingDto {
	return &NullableCreateOrUpdateAppPricingDto{value: val, isSet: true}
}

func (v NullableCreateOrUpdateAppPricingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateAppPricingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


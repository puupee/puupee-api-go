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

// CurrentCultureDto struct for CurrentCultureDto
type CurrentCultureDto struct {
	DisplayName *string `json:"displayName,omitempty"`
	EnglishName *string `json:"englishName,omitempty"`
	ThreeLetterIsoLanguageName *string `json:"threeLetterIsoLanguageName,omitempty"`
	TwoLetterIsoLanguageName *string `json:"twoLetterIsoLanguageName,omitempty"`
	IsRightToLeft *bool `json:"isRightToLeft,omitempty"`
	CultureName *string `json:"cultureName,omitempty"`
	Name *string `json:"name,omitempty"`
	NativeName *string `json:"nativeName,omitempty"`
	DateTimeFormat *DateTimeFormatDto `json:"dateTimeFormat,omitempty"`
}

// NewCurrentCultureDto instantiates a new CurrentCultureDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentCultureDto() *CurrentCultureDto {
	this := CurrentCultureDto{}
	return &this
}

// NewCurrentCultureDtoWithDefaults instantiates a new CurrentCultureDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentCultureDtoWithDefaults() *CurrentCultureDto {
	this := CurrentCultureDto{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CurrentCultureDto) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCultureDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CurrentCultureDto) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CurrentCultureDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnglishName returns the EnglishName field value if set, zero value otherwise.
func (o *CurrentCultureDto) GetEnglishName() string {
	if o == nil || isNil(o.EnglishName) {
		var ret string
		return ret
	}
	return *o.EnglishName
}

// GetEnglishNameOk returns a tuple with the EnglishName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCultureDto) GetEnglishNameOk() (*string, bool) {
	if o == nil || isNil(o.EnglishName) {
    return nil, false
	}
	return o.EnglishName, true
}

// HasEnglishName returns a boolean if a field has been set.
func (o *CurrentCultureDto) HasEnglishName() bool {
	if o != nil && !isNil(o.EnglishName) {
		return true
	}

	return false
}

// SetEnglishName gets a reference to the given string and assigns it to the EnglishName field.
func (o *CurrentCultureDto) SetEnglishName(v string) {
	o.EnglishName = &v
}

// GetThreeLetterIsoLanguageName returns the ThreeLetterIsoLanguageName field value if set, zero value otherwise.
func (o *CurrentCultureDto) GetThreeLetterIsoLanguageName() string {
	if o == nil || isNil(o.ThreeLetterIsoLanguageName) {
		var ret string
		return ret
	}
	return *o.ThreeLetterIsoLanguageName
}

// GetThreeLetterIsoLanguageNameOk returns a tuple with the ThreeLetterIsoLanguageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCultureDto) GetThreeLetterIsoLanguageNameOk() (*string, bool) {
	if o == nil || isNil(o.ThreeLetterIsoLanguageName) {
    return nil, false
	}
	return o.ThreeLetterIsoLanguageName, true
}

// HasThreeLetterIsoLanguageName returns a boolean if a field has been set.
func (o *CurrentCultureDto) HasThreeLetterIsoLanguageName() bool {
	if o != nil && !isNil(o.ThreeLetterIsoLanguageName) {
		return true
	}

	return false
}

// SetThreeLetterIsoLanguageName gets a reference to the given string and assigns it to the ThreeLetterIsoLanguageName field.
func (o *CurrentCultureDto) SetThreeLetterIsoLanguageName(v string) {
	o.ThreeLetterIsoLanguageName = &v
}

// GetTwoLetterIsoLanguageName returns the TwoLetterIsoLanguageName field value if set, zero value otherwise.
func (o *CurrentCultureDto) GetTwoLetterIsoLanguageName() string {
	if o == nil || isNil(o.TwoLetterIsoLanguageName) {
		var ret string
		return ret
	}
	return *o.TwoLetterIsoLanguageName
}

// GetTwoLetterIsoLanguageNameOk returns a tuple with the TwoLetterIsoLanguageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCultureDto) GetTwoLetterIsoLanguageNameOk() (*string, bool) {
	if o == nil || isNil(o.TwoLetterIsoLanguageName) {
    return nil, false
	}
	return o.TwoLetterIsoLanguageName, true
}

// HasTwoLetterIsoLanguageName returns a boolean if a field has been set.
func (o *CurrentCultureDto) HasTwoLetterIsoLanguageName() bool {
	if o != nil && !isNil(o.TwoLetterIsoLanguageName) {
		return true
	}

	return false
}

// SetTwoLetterIsoLanguageName gets a reference to the given string and assigns it to the TwoLetterIsoLanguageName field.
func (o *CurrentCultureDto) SetTwoLetterIsoLanguageName(v string) {
	o.TwoLetterIsoLanguageName = &v
}

// GetIsRightToLeft returns the IsRightToLeft field value if set, zero value otherwise.
func (o *CurrentCultureDto) GetIsRightToLeft() bool {
	if o == nil || isNil(o.IsRightToLeft) {
		var ret bool
		return ret
	}
	return *o.IsRightToLeft
}

// GetIsRightToLeftOk returns a tuple with the IsRightToLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCultureDto) GetIsRightToLeftOk() (*bool, bool) {
	if o == nil || isNil(o.IsRightToLeft) {
    return nil, false
	}
	return o.IsRightToLeft, true
}

// HasIsRightToLeft returns a boolean if a field has been set.
func (o *CurrentCultureDto) HasIsRightToLeft() bool {
	if o != nil && !isNil(o.IsRightToLeft) {
		return true
	}

	return false
}

// SetIsRightToLeft gets a reference to the given bool and assigns it to the IsRightToLeft field.
func (o *CurrentCultureDto) SetIsRightToLeft(v bool) {
	o.IsRightToLeft = &v
}

// GetCultureName returns the CultureName field value if set, zero value otherwise.
func (o *CurrentCultureDto) GetCultureName() string {
	if o == nil || isNil(o.CultureName) {
		var ret string
		return ret
	}
	return *o.CultureName
}

// GetCultureNameOk returns a tuple with the CultureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCultureDto) GetCultureNameOk() (*string, bool) {
	if o == nil || isNil(o.CultureName) {
    return nil, false
	}
	return o.CultureName, true
}

// HasCultureName returns a boolean if a field has been set.
func (o *CurrentCultureDto) HasCultureName() bool {
	if o != nil && !isNil(o.CultureName) {
		return true
	}

	return false
}

// SetCultureName gets a reference to the given string and assigns it to the CultureName field.
func (o *CurrentCultureDto) SetCultureName(v string) {
	o.CultureName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CurrentCultureDto) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCultureDto) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CurrentCultureDto) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CurrentCultureDto) SetName(v string) {
	o.Name = &v
}

// GetNativeName returns the NativeName field value if set, zero value otherwise.
func (o *CurrentCultureDto) GetNativeName() string {
	if o == nil || isNil(o.NativeName) {
		var ret string
		return ret
	}
	return *o.NativeName
}

// GetNativeNameOk returns a tuple with the NativeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCultureDto) GetNativeNameOk() (*string, bool) {
	if o == nil || isNil(o.NativeName) {
    return nil, false
	}
	return o.NativeName, true
}

// HasNativeName returns a boolean if a field has been set.
func (o *CurrentCultureDto) HasNativeName() bool {
	if o != nil && !isNil(o.NativeName) {
		return true
	}

	return false
}

// SetNativeName gets a reference to the given string and assigns it to the NativeName field.
func (o *CurrentCultureDto) SetNativeName(v string) {
	o.NativeName = &v
}

// GetDateTimeFormat returns the DateTimeFormat field value if set, zero value otherwise.
func (o *CurrentCultureDto) GetDateTimeFormat() DateTimeFormatDto {
	if o == nil || isNil(o.DateTimeFormat) {
		var ret DateTimeFormatDto
		return ret
	}
	return *o.DateTimeFormat
}

// GetDateTimeFormatOk returns a tuple with the DateTimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentCultureDto) GetDateTimeFormatOk() (*DateTimeFormatDto, bool) {
	if o == nil || isNil(o.DateTimeFormat) {
    return nil, false
	}
	return o.DateTimeFormat, true
}

// HasDateTimeFormat returns a boolean if a field has been set.
func (o *CurrentCultureDto) HasDateTimeFormat() bool {
	if o != nil && !isNil(o.DateTimeFormat) {
		return true
	}

	return false
}

// SetDateTimeFormat gets a reference to the given DateTimeFormatDto and assigns it to the DateTimeFormat field.
func (o *CurrentCultureDto) SetDateTimeFormat(v DateTimeFormatDto) {
	o.DateTimeFormat = &v
}

func (o CurrentCultureDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.EnglishName) {
		toSerialize["englishName"] = o.EnglishName
	}
	if !isNil(o.ThreeLetterIsoLanguageName) {
		toSerialize["threeLetterIsoLanguageName"] = o.ThreeLetterIsoLanguageName
	}
	if !isNil(o.TwoLetterIsoLanguageName) {
		toSerialize["twoLetterIsoLanguageName"] = o.TwoLetterIsoLanguageName
	}
	if !isNil(o.IsRightToLeft) {
		toSerialize["isRightToLeft"] = o.IsRightToLeft
	}
	if !isNil(o.CultureName) {
		toSerialize["cultureName"] = o.CultureName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NativeName) {
		toSerialize["nativeName"] = o.NativeName
	}
	if !isNil(o.DateTimeFormat) {
		toSerialize["dateTimeFormat"] = o.DateTimeFormat
	}
	return json.Marshal(toSerialize)
}

type NullableCurrentCultureDto struct {
	value *CurrentCultureDto
	isSet bool
}

func (v NullableCurrentCultureDto) Get() *CurrentCultureDto {
	return v.value
}

func (v *NullableCurrentCultureDto) Set(val *CurrentCultureDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentCultureDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentCultureDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentCultureDto(val *CurrentCultureDto) *NullableCurrentCultureDto {
	return &NullableCurrentCultureDto{value: val, isSet: true}
}

func (v NullableCurrentCultureDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentCultureDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



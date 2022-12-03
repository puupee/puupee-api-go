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

// AppTheme struct for AppTheme
type AppTheme struct {
	SourceColor *string `json:"sourceColor,omitempty"`
	ThemeMode *AppThemeMode `json:"themeMode,omitempty"`
}

// NewAppTheme instantiates a new AppTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppTheme() *AppTheme {
	this := AppTheme{}
	return &this
}

// NewAppThemeWithDefaults instantiates a new AppTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppThemeWithDefaults() *AppTheme {
	this := AppTheme{}
	return &this
}

// GetSourceColor returns the SourceColor field value if set, zero value otherwise.
func (o *AppTheme) GetSourceColor() string {
	if o == nil || isNil(o.SourceColor) {
		var ret string
		return ret
	}
	return *o.SourceColor
}

// GetSourceColorOk returns a tuple with the SourceColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppTheme) GetSourceColorOk() (*string, bool) {
	if o == nil || isNil(o.SourceColor) {
    return nil, false
	}
	return o.SourceColor, true
}

// HasSourceColor returns a boolean if a field has been set.
func (o *AppTheme) HasSourceColor() bool {
	if o != nil && !isNil(o.SourceColor) {
		return true
	}

	return false
}

// SetSourceColor gets a reference to the given string and assigns it to the SourceColor field.
func (o *AppTheme) SetSourceColor(v string) {
	o.SourceColor = &v
}

// GetThemeMode returns the ThemeMode field value if set, zero value otherwise.
func (o *AppTheme) GetThemeMode() AppThemeMode {
	if o == nil || isNil(o.ThemeMode) {
		var ret AppThemeMode
		return ret
	}
	return *o.ThemeMode
}

// GetThemeModeOk returns a tuple with the ThemeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppTheme) GetThemeModeOk() (*AppThemeMode, bool) {
	if o == nil || isNil(o.ThemeMode) {
    return nil, false
	}
	return o.ThemeMode, true
}

// HasThemeMode returns a boolean if a field has been set.
func (o *AppTheme) HasThemeMode() bool {
	if o != nil && !isNil(o.ThemeMode) {
		return true
	}

	return false
}

// SetThemeMode gets a reference to the given AppThemeMode and assigns it to the ThemeMode field.
func (o *AppTheme) SetThemeMode(v AppThemeMode) {
	o.ThemeMode = &v
}

func (o AppTheme) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SourceColor) {
		toSerialize["sourceColor"] = o.SourceColor
	}
	if !isNil(o.ThemeMode) {
		toSerialize["themeMode"] = o.ThemeMode
	}
	return json.Marshal(toSerialize)
}

type NullableAppTheme struct {
	value *AppTheme
	isSet bool
}

func (v NullableAppTheme) Get() *AppTheme {
	return v.value
}

func (v *NullableAppTheme) Set(val *AppTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableAppTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableAppTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppTheme(val *AppTheme) *NullableAppTheme {
	return &NullableAppTheme{value: val, isSet: true}
}

func (v NullableAppTheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the DateTimeFormatDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateTimeFormatDto{}

// DateTimeFormatDto struct for DateTimeFormatDto
type DateTimeFormatDto struct {
	CalendarAlgorithmType NullableString `json:"calendarAlgorithmType,omitempty"`
	DateTimeFormatLong NullableString `json:"dateTimeFormatLong,omitempty"`
	ShortDatePattern NullableString `json:"shortDatePattern,omitempty"`
	FullDateTimePattern NullableString `json:"fullDateTimePattern,omitempty"`
	DateSeparator NullableString `json:"dateSeparator,omitempty"`
	ShortTimePattern NullableString `json:"shortTimePattern,omitempty"`
	LongTimePattern NullableString `json:"longTimePattern,omitempty"`
}

// NewDateTimeFormatDto instantiates a new DateTimeFormatDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeFormatDto() *DateTimeFormatDto {
	this := DateTimeFormatDto{}
	return &this
}

// NewDateTimeFormatDtoWithDefaults instantiates a new DateTimeFormatDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeFormatDtoWithDefaults() *DateTimeFormatDto {
	this := DateTimeFormatDto{}
	return &this
}

// GetCalendarAlgorithmType returns the CalendarAlgorithmType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFormatDto) GetCalendarAlgorithmType() string {
	if o == nil || IsNil(o.CalendarAlgorithmType.Get()) {
		var ret string
		return ret
	}
	return *o.CalendarAlgorithmType.Get()
}

// GetCalendarAlgorithmTypeOk returns a tuple with the CalendarAlgorithmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFormatDto) GetCalendarAlgorithmTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CalendarAlgorithmType.Get(), o.CalendarAlgorithmType.IsSet()
}

// HasCalendarAlgorithmType returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasCalendarAlgorithmType() bool {
	if o != nil && o.CalendarAlgorithmType.IsSet() {
		return true
	}

	return false
}

// SetCalendarAlgorithmType gets a reference to the given NullableString and assigns it to the CalendarAlgorithmType field.
func (o *DateTimeFormatDto) SetCalendarAlgorithmType(v string) {
	o.CalendarAlgorithmType.Set(&v)
}
// SetCalendarAlgorithmTypeNil sets the value for CalendarAlgorithmType to be an explicit nil
func (o *DateTimeFormatDto) SetCalendarAlgorithmTypeNil() {
	o.CalendarAlgorithmType.Set(nil)
}

// UnsetCalendarAlgorithmType ensures that no value is present for CalendarAlgorithmType, not even an explicit nil
func (o *DateTimeFormatDto) UnsetCalendarAlgorithmType() {
	o.CalendarAlgorithmType.Unset()
}

// GetDateTimeFormatLong returns the DateTimeFormatLong field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFormatDto) GetDateTimeFormatLong() string {
	if o == nil || IsNil(o.DateTimeFormatLong.Get()) {
		var ret string
		return ret
	}
	return *o.DateTimeFormatLong.Get()
}

// GetDateTimeFormatLongOk returns a tuple with the DateTimeFormatLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFormatDto) GetDateTimeFormatLongOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateTimeFormatLong.Get(), o.DateTimeFormatLong.IsSet()
}

// HasDateTimeFormatLong returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasDateTimeFormatLong() bool {
	if o != nil && o.DateTimeFormatLong.IsSet() {
		return true
	}

	return false
}

// SetDateTimeFormatLong gets a reference to the given NullableString and assigns it to the DateTimeFormatLong field.
func (o *DateTimeFormatDto) SetDateTimeFormatLong(v string) {
	o.DateTimeFormatLong.Set(&v)
}
// SetDateTimeFormatLongNil sets the value for DateTimeFormatLong to be an explicit nil
func (o *DateTimeFormatDto) SetDateTimeFormatLongNil() {
	o.DateTimeFormatLong.Set(nil)
}

// UnsetDateTimeFormatLong ensures that no value is present for DateTimeFormatLong, not even an explicit nil
func (o *DateTimeFormatDto) UnsetDateTimeFormatLong() {
	o.DateTimeFormatLong.Unset()
}

// GetShortDatePattern returns the ShortDatePattern field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFormatDto) GetShortDatePattern() string {
	if o == nil || IsNil(o.ShortDatePattern.Get()) {
		var ret string
		return ret
	}
	return *o.ShortDatePattern.Get()
}

// GetShortDatePatternOk returns a tuple with the ShortDatePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFormatDto) GetShortDatePatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortDatePattern.Get(), o.ShortDatePattern.IsSet()
}

// HasShortDatePattern returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasShortDatePattern() bool {
	if o != nil && o.ShortDatePattern.IsSet() {
		return true
	}

	return false
}

// SetShortDatePattern gets a reference to the given NullableString and assigns it to the ShortDatePattern field.
func (o *DateTimeFormatDto) SetShortDatePattern(v string) {
	o.ShortDatePattern.Set(&v)
}
// SetShortDatePatternNil sets the value for ShortDatePattern to be an explicit nil
func (o *DateTimeFormatDto) SetShortDatePatternNil() {
	o.ShortDatePattern.Set(nil)
}

// UnsetShortDatePattern ensures that no value is present for ShortDatePattern, not even an explicit nil
func (o *DateTimeFormatDto) UnsetShortDatePattern() {
	o.ShortDatePattern.Unset()
}

// GetFullDateTimePattern returns the FullDateTimePattern field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFormatDto) GetFullDateTimePattern() string {
	if o == nil || IsNil(o.FullDateTimePattern.Get()) {
		var ret string
		return ret
	}
	return *o.FullDateTimePattern.Get()
}

// GetFullDateTimePatternOk returns a tuple with the FullDateTimePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFormatDto) GetFullDateTimePatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullDateTimePattern.Get(), o.FullDateTimePattern.IsSet()
}

// HasFullDateTimePattern returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasFullDateTimePattern() bool {
	if o != nil && o.FullDateTimePattern.IsSet() {
		return true
	}

	return false
}

// SetFullDateTimePattern gets a reference to the given NullableString and assigns it to the FullDateTimePattern field.
func (o *DateTimeFormatDto) SetFullDateTimePattern(v string) {
	o.FullDateTimePattern.Set(&v)
}
// SetFullDateTimePatternNil sets the value for FullDateTimePattern to be an explicit nil
func (o *DateTimeFormatDto) SetFullDateTimePatternNil() {
	o.FullDateTimePattern.Set(nil)
}

// UnsetFullDateTimePattern ensures that no value is present for FullDateTimePattern, not even an explicit nil
func (o *DateTimeFormatDto) UnsetFullDateTimePattern() {
	o.FullDateTimePattern.Unset()
}

// GetDateSeparator returns the DateSeparator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFormatDto) GetDateSeparator() string {
	if o == nil || IsNil(o.DateSeparator.Get()) {
		var ret string
		return ret
	}
	return *o.DateSeparator.Get()
}

// GetDateSeparatorOk returns a tuple with the DateSeparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFormatDto) GetDateSeparatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateSeparator.Get(), o.DateSeparator.IsSet()
}

// HasDateSeparator returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasDateSeparator() bool {
	if o != nil && o.DateSeparator.IsSet() {
		return true
	}

	return false
}

// SetDateSeparator gets a reference to the given NullableString and assigns it to the DateSeparator field.
func (o *DateTimeFormatDto) SetDateSeparator(v string) {
	o.DateSeparator.Set(&v)
}
// SetDateSeparatorNil sets the value for DateSeparator to be an explicit nil
func (o *DateTimeFormatDto) SetDateSeparatorNil() {
	o.DateSeparator.Set(nil)
}

// UnsetDateSeparator ensures that no value is present for DateSeparator, not even an explicit nil
func (o *DateTimeFormatDto) UnsetDateSeparator() {
	o.DateSeparator.Unset()
}

// GetShortTimePattern returns the ShortTimePattern field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFormatDto) GetShortTimePattern() string {
	if o == nil || IsNil(o.ShortTimePattern.Get()) {
		var ret string
		return ret
	}
	return *o.ShortTimePattern.Get()
}

// GetShortTimePatternOk returns a tuple with the ShortTimePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFormatDto) GetShortTimePatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortTimePattern.Get(), o.ShortTimePattern.IsSet()
}

// HasShortTimePattern returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasShortTimePattern() bool {
	if o != nil && o.ShortTimePattern.IsSet() {
		return true
	}

	return false
}

// SetShortTimePattern gets a reference to the given NullableString and assigns it to the ShortTimePattern field.
func (o *DateTimeFormatDto) SetShortTimePattern(v string) {
	o.ShortTimePattern.Set(&v)
}
// SetShortTimePatternNil sets the value for ShortTimePattern to be an explicit nil
func (o *DateTimeFormatDto) SetShortTimePatternNil() {
	o.ShortTimePattern.Set(nil)
}

// UnsetShortTimePattern ensures that no value is present for ShortTimePattern, not even an explicit nil
func (o *DateTimeFormatDto) UnsetShortTimePattern() {
	o.ShortTimePattern.Unset()
}

// GetLongTimePattern returns the LongTimePattern field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DateTimeFormatDto) GetLongTimePattern() string {
	if o == nil || IsNil(o.LongTimePattern.Get()) {
		var ret string
		return ret
	}
	return *o.LongTimePattern.Get()
}

// GetLongTimePatternOk returns a tuple with the LongTimePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DateTimeFormatDto) GetLongTimePatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LongTimePattern.Get(), o.LongTimePattern.IsSet()
}

// HasLongTimePattern returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasLongTimePattern() bool {
	if o != nil && o.LongTimePattern.IsSet() {
		return true
	}

	return false
}

// SetLongTimePattern gets a reference to the given NullableString and assigns it to the LongTimePattern field.
func (o *DateTimeFormatDto) SetLongTimePattern(v string) {
	o.LongTimePattern.Set(&v)
}
// SetLongTimePatternNil sets the value for LongTimePattern to be an explicit nil
func (o *DateTimeFormatDto) SetLongTimePatternNil() {
	o.LongTimePattern.Set(nil)
}

// UnsetLongTimePattern ensures that no value is present for LongTimePattern, not even an explicit nil
func (o *DateTimeFormatDto) UnsetLongTimePattern() {
	o.LongTimePattern.Unset()
}

func (o DateTimeFormatDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateTimeFormatDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CalendarAlgorithmType.IsSet() {
		toSerialize["calendarAlgorithmType"] = o.CalendarAlgorithmType.Get()
	}
	if o.DateTimeFormatLong.IsSet() {
		toSerialize["dateTimeFormatLong"] = o.DateTimeFormatLong.Get()
	}
	if o.ShortDatePattern.IsSet() {
		toSerialize["shortDatePattern"] = o.ShortDatePattern.Get()
	}
	if o.FullDateTimePattern.IsSet() {
		toSerialize["fullDateTimePattern"] = o.FullDateTimePattern.Get()
	}
	if o.DateSeparator.IsSet() {
		toSerialize["dateSeparator"] = o.DateSeparator.Get()
	}
	if o.ShortTimePattern.IsSet() {
		toSerialize["shortTimePattern"] = o.ShortTimePattern.Get()
	}
	if o.LongTimePattern.IsSet() {
		toSerialize["longTimePattern"] = o.LongTimePattern.Get()
	}
	return toSerialize, nil
}

type NullableDateTimeFormatDto struct {
	value *DateTimeFormatDto
	isSet bool
}

func (v NullableDateTimeFormatDto) Get() *DateTimeFormatDto {
	return v.value
}

func (v *NullableDateTimeFormatDto) Set(val *DateTimeFormatDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeFormatDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeFormatDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeFormatDto(val *DateTimeFormatDto) *NullableDateTimeFormatDto {
	return &NullableDateTimeFormatDto{value: val, isSet: true}
}

func (v NullableDateTimeFormatDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateTimeFormatDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



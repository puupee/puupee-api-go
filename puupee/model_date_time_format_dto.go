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

// DateTimeFormatDto struct for DateTimeFormatDto
type DateTimeFormatDto struct {
	CalendarAlgorithmType *string `json:"calendarAlgorithmType,omitempty"`
	DateTimeFormatLong *string `json:"dateTimeFormatLong,omitempty"`
	ShortDatePattern *string `json:"shortDatePattern,omitempty"`
	FullDateTimePattern *string `json:"fullDateTimePattern,omitempty"`
	DateSeparator *string `json:"dateSeparator,omitempty"`
	ShortTimePattern *string `json:"shortTimePattern,omitempty"`
	LongTimePattern *string `json:"longTimePattern,omitempty"`
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

// GetCalendarAlgorithmType returns the CalendarAlgorithmType field value if set, zero value otherwise.
func (o *DateTimeFormatDto) GetCalendarAlgorithmType() string {
	if o == nil || isNil(o.CalendarAlgorithmType) {
		var ret string
		return ret
	}
	return *o.CalendarAlgorithmType
}

// GetCalendarAlgorithmTypeOk returns a tuple with the CalendarAlgorithmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFormatDto) GetCalendarAlgorithmTypeOk() (*string, bool) {
	if o == nil || isNil(o.CalendarAlgorithmType) {
    return nil, false
	}
	return o.CalendarAlgorithmType, true
}

// HasCalendarAlgorithmType returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasCalendarAlgorithmType() bool {
	if o != nil && !isNil(o.CalendarAlgorithmType) {
		return true
	}

	return false
}

// SetCalendarAlgorithmType gets a reference to the given string and assigns it to the CalendarAlgorithmType field.
func (o *DateTimeFormatDto) SetCalendarAlgorithmType(v string) {
	o.CalendarAlgorithmType = &v
}

// GetDateTimeFormatLong returns the DateTimeFormatLong field value if set, zero value otherwise.
func (o *DateTimeFormatDto) GetDateTimeFormatLong() string {
	if o == nil || isNil(o.DateTimeFormatLong) {
		var ret string
		return ret
	}
	return *o.DateTimeFormatLong
}

// GetDateTimeFormatLongOk returns a tuple with the DateTimeFormatLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFormatDto) GetDateTimeFormatLongOk() (*string, bool) {
	if o == nil || isNil(o.DateTimeFormatLong) {
    return nil, false
	}
	return o.DateTimeFormatLong, true
}

// HasDateTimeFormatLong returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasDateTimeFormatLong() bool {
	if o != nil && !isNil(o.DateTimeFormatLong) {
		return true
	}

	return false
}

// SetDateTimeFormatLong gets a reference to the given string and assigns it to the DateTimeFormatLong field.
func (o *DateTimeFormatDto) SetDateTimeFormatLong(v string) {
	o.DateTimeFormatLong = &v
}

// GetShortDatePattern returns the ShortDatePattern field value if set, zero value otherwise.
func (o *DateTimeFormatDto) GetShortDatePattern() string {
	if o == nil || isNil(o.ShortDatePattern) {
		var ret string
		return ret
	}
	return *o.ShortDatePattern
}

// GetShortDatePatternOk returns a tuple with the ShortDatePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFormatDto) GetShortDatePatternOk() (*string, bool) {
	if o == nil || isNil(o.ShortDatePattern) {
    return nil, false
	}
	return o.ShortDatePattern, true
}

// HasShortDatePattern returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasShortDatePattern() bool {
	if o != nil && !isNil(o.ShortDatePattern) {
		return true
	}

	return false
}

// SetShortDatePattern gets a reference to the given string and assigns it to the ShortDatePattern field.
func (o *DateTimeFormatDto) SetShortDatePattern(v string) {
	o.ShortDatePattern = &v
}

// GetFullDateTimePattern returns the FullDateTimePattern field value if set, zero value otherwise.
func (o *DateTimeFormatDto) GetFullDateTimePattern() string {
	if o == nil || isNil(o.FullDateTimePattern) {
		var ret string
		return ret
	}
	return *o.FullDateTimePattern
}

// GetFullDateTimePatternOk returns a tuple with the FullDateTimePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFormatDto) GetFullDateTimePatternOk() (*string, bool) {
	if o == nil || isNil(o.FullDateTimePattern) {
    return nil, false
	}
	return o.FullDateTimePattern, true
}

// HasFullDateTimePattern returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasFullDateTimePattern() bool {
	if o != nil && !isNil(o.FullDateTimePattern) {
		return true
	}

	return false
}

// SetFullDateTimePattern gets a reference to the given string and assigns it to the FullDateTimePattern field.
func (o *DateTimeFormatDto) SetFullDateTimePattern(v string) {
	o.FullDateTimePattern = &v
}

// GetDateSeparator returns the DateSeparator field value if set, zero value otherwise.
func (o *DateTimeFormatDto) GetDateSeparator() string {
	if o == nil || isNil(o.DateSeparator) {
		var ret string
		return ret
	}
	return *o.DateSeparator
}

// GetDateSeparatorOk returns a tuple with the DateSeparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFormatDto) GetDateSeparatorOk() (*string, bool) {
	if o == nil || isNil(o.DateSeparator) {
    return nil, false
	}
	return o.DateSeparator, true
}

// HasDateSeparator returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasDateSeparator() bool {
	if o != nil && !isNil(o.DateSeparator) {
		return true
	}

	return false
}

// SetDateSeparator gets a reference to the given string and assigns it to the DateSeparator field.
func (o *DateTimeFormatDto) SetDateSeparator(v string) {
	o.DateSeparator = &v
}

// GetShortTimePattern returns the ShortTimePattern field value if set, zero value otherwise.
func (o *DateTimeFormatDto) GetShortTimePattern() string {
	if o == nil || isNil(o.ShortTimePattern) {
		var ret string
		return ret
	}
	return *o.ShortTimePattern
}

// GetShortTimePatternOk returns a tuple with the ShortTimePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFormatDto) GetShortTimePatternOk() (*string, bool) {
	if o == nil || isNil(o.ShortTimePattern) {
    return nil, false
	}
	return o.ShortTimePattern, true
}

// HasShortTimePattern returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasShortTimePattern() bool {
	if o != nil && !isNil(o.ShortTimePattern) {
		return true
	}

	return false
}

// SetShortTimePattern gets a reference to the given string and assigns it to the ShortTimePattern field.
func (o *DateTimeFormatDto) SetShortTimePattern(v string) {
	o.ShortTimePattern = &v
}

// GetLongTimePattern returns the LongTimePattern field value if set, zero value otherwise.
func (o *DateTimeFormatDto) GetLongTimePattern() string {
	if o == nil || isNil(o.LongTimePattern) {
		var ret string
		return ret
	}
	return *o.LongTimePattern
}

// GetLongTimePatternOk returns a tuple with the LongTimePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeFormatDto) GetLongTimePatternOk() (*string, bool) {
	if o == nil || isNil(o.LongTimePattern) {
    return nil, false
	}
	return o.LongTimePattern, true
}

// HasLongTimePattern returns a boolean if a field has been set.
func (o *DateTimeFormatDto) HasLongTimePattern() bool {
	if o != nil && !isNil(o.LongTimePattern) {
		return true
	}

	return false
}

// SetLongTimePattern gets a reference to the given string and assigns it to the LongTimePattern field.
func (o *DateTimeFormatDto) SetLongTimePattern(v string) {
	o.LongTimePattern = &v
}

func (o DateTimeFormatDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CalendarAlgorithmType) {
		toSerialize["calendarAlgorithmType"] = o.CalendarAlgorithmType
	}
	if !isNil(o.DateTimeFormatLong) {
		toSerialize["dateTimeFormatLong"] = o.DateTimeFormatLong
	}
	if !isNil(o.ShortDatePattern) {
		toSerialize["shortDatePattern"] = o.ShortDatePattern
	}
	if !isNil(o.FullDateTimePattern) {
		toSerialize["fullDateTimePattern"] = o.FullDateTimePattern
	}
	if !isNil(o.DateSeparator) {
		toSerialize["dateSeparator"] = o.DateSeparator
	}
	if !isNil(o.ShortTimePattern) {
		toSerialize["shortTimePattern"] = o.ShortTimePattern
	}
	if !isNil(o.LongTimePattern) {
		toSerialize["longTimePattern"] = o.LongTimePattern
	}
	return json.Marshal(toSerialize)
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



/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.17.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ApiKeyUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyUpdateDto{}

// ApiKeyUpdateDto struct for ApiKeyUpdateDto
type ApiKeyUpdateDto struct {
	Name string `json:"name"`
	Active *bool `json:"active,omitempty"`
	ExpireAt *time.Time `json:"expireAt,omitempty"`
}

type _ApiKeyUpdateDto ApiKeyUpdateDto

// NewApiKeyUpdateDto instantiates a new ApiKeyUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyUpdateDto(name string) *ApiKeyUpdateDto {
	this := ApiKeyUpdateDto{}
	this.Name = name
	return &this
}

// NewApiKeyUpdateDtoWithDefaults instantiates a new ApiKeyUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyUpdateDtoWithDefaults() *ApiKeyUpdateDto {
	this := ApiKeyUpdateDto{}
	return &this
}

// GetName returns the Name field value
func (o *ApiKeyUpdateDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiKeyUpdateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiKeyUpdateDto) SetName(v string) {
	o.Name = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApiKeyUpdateDto) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyUpdateDto) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApiKeyUpdateDto) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApiKeyUpdateDto) SetActive(v bool) {
	o.Active = &v
}

// GetExpireAt returns the ExpireAt field value if set, zero value otherwise.
func (o *ApiKeyUpdateDto) GetExpireAt() time.Time {
	if o == nil || IsNil(o.ExpireAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpireAt
}

// GetExpireAtOk returns a tuple with the ExpireAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyUpdateDto) GetExpireAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpireAt) {
		return nil, false
	}
	return o.ExpireAt, true
}

// HasExpireAt returns a boolean if a field has been set.
func (o *ApiKeyUpdateDto) HasExpireAt() bool {
	if o != nil && !IsNil(o.ExpireAt) {
		return true
	}

	return false
}

// SetExpireAt gets a reference to the given time.Time and assigns it to the ExpireAt field.
func (o *ApiKeyUpdateDto) SetExpireAt(v time.Time) {
	o.ExpireAt = &v
}

func (o ApiKeyUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ExpireAt) {
		toSerialize["expireAt"] = o.ExpireAt
	}
	return toSerialize, nil
}

func (o *ApiKeyUpdateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiKeyUpdateDto := _ApiKeyUpdateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiKeyUpdateDto)

	if err != nil {
		return err
	}

	*o = ApiKeyUpdateDto(varApiKeyUpdateDto)

	return err
}

type NullableApiKeyUpdateDto struct {
	value *ApiKeyUpdateDto
	isSet bool
}

func (v NullableApiKeyUpdateDto) Get() *ApiKeyUpdateDto {
	return v.value
}

func (v *NullableApiKeyUpdateDto) Set(val *ApiKeyUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyUpdateDto(val *ApiKeyUpdateDto) *NullableApiKeyUpdateDto {
	return &NullableApiKeyUpdateDto{value: val, isSet: true}
}

func (v NullableApiKeyUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



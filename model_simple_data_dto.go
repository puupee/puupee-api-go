/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// SimpleDataDto struct for SimpleDataDto
type SimpleDataDto struct {
	ExtraProperties map[string]map[string]interface{} `json:"extraProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	Collection *string `json:"collection,omitempty"`
}

// NewSimpleDataDto instantiates a new SimpleDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleDataDto() *SimpleDataDto {
	this := SimpleDataDto{}
	return &this
}

// NewSimpleDataDtoWithDefaults instantiates a new SimpleDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleDataDtoWithDefaults() *SimpleDataDto {
	this := SimpleDataDto{}
	return &this
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise.
func (o *SimpleDataDto) GetExtraProperties() map[string]map[string]interface{} {
	if o == nil || isNil(o.ExtraProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleDataDto) GetExtraPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.ExtraProperties) {
    return map[string]map[string]interface{}{}, false
	}
	return o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *SimpleDataDto) HasExtraProperties() bool {
	if o != nil && !isNil(o.ExtraProperties) {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the ExtraProperties field.
func (o *SimpleDataDto) SetExtraProperties(v map[string]map[string]interface{}) {
	o.ExtraProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimpleDataDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleDataDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimpleDataDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SimpleDataDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *SimpleDataDto) GetCreationTime() time.Time {
	if o == nil || isNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleDataDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreationTime) {
    return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *SimpleDataDto) HasCreationTime() bool {
	if o != nil && !isNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *SimpleDataDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *SimpleDataDto) GetCreatorId() string {
	if o == nil || isNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleDataDto) GetCreatorIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatorId) {
    return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *SimpleDataDto) HasCreatorId() bool {
	if o != nil && !isNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *SimpleDataDto) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *SimpleDataDto) GetCollection() string {
	if o == nil || isNil(o.Collection) {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleDataDto) GetCollectionOk() (*string, bool) {
	if o == nil || isNil(o.Collection) {
    return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *SimpleDataDto) HasCollection() bool {
	if o != nil && !isNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *SimpleDataDto) SetCollection(v string) {
	o.Collection = &v
}

func (o SimpleDataDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExtraProperties) {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !isNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if !isNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleDataDto struct {
	value *SimpleDataDto
	isSet bool
}

func (v NullableSimpleDataDto) Get() *SimpleDataDto {
	return v.value
}

func (v *NullableSimpleDataDto) Set(val *SimpleDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleDataDto(val *SimpleDataDto) *NullableSimpleDataDto {
	return &NullableSimpleDataDto{value: val, isSet: true}
}

func (v NullableSimpleDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



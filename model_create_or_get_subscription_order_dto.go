/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
	"time"
)

// CreateOrGetSubscriptionOrderDto struct for CreateOrGetSubscriptionOrderDto
type CreateOrGetSubscriptionOrderDto struct {
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	LastModificationTime *time.Time `json:"lastModificationTime,omitempty"`
	LastModifierId *string `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId *string `json:"deleterId,omitempty"`
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	AppId *string `json:"appId,omitempty"`
	PricingId *string `json:"pricingId,omitempty"`
}

// NewCreateOrGetSubscriptionOrderDto instantiates a new CreateOrGetSubscriptionOrderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrGetSubscriptionOrderDto() *CreateOrGetSubscriptionOrderDto {
	this := CreateOrGetSubscriptionOrderDto{}
	return &this
}

// NewCreateOrGetSubscriptionOrderDtoWithDefaults instantiates a new CreateOrGetSubscriptionOrderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrGetSubscriptionOrderDtoWithDefaults() *CreateOrGetSubscriptionOrderDto {
	this := CreateOrGetSubscriptionOrderDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateOrGetSubscriptionOrderDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetCreationTime() time.Time {
	if o == nil || isNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreationTime) {
    return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasCreationTime() bool {
	if o != nil && !isNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *CreateOrGetSubscriptionOrderDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetCreatorId() string {
	if o == nil || isNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetCreatorIdOk() (*string, bool) {
	if o == nil || isNil(o.CreatorId) {
    return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasCreatorId() bool {
	if o != nil && !isNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *CreateOrGetSubscriptionOrderDto) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetLastModificationTime() time.Time {
	if o == nil || isNil(o.LastModificationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastModificationTime) {
    return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasLastModificationTime() bool {
	if o != nil && !isNil(o.LastModificationTime) {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given time.Time and assigns it to the LastModificationTime field.
func (o *CreateOrGetSubscriptionOrderDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetLastModifierId() string {
	if o == nil || isNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil || isNil(o.LastModifierId) {
    return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasLastModifierId() bool {
	if o != nil && !isNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *CreateOrGetSubscriptionOrderDto) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetIsDeleted() bool {
	if o == nil || isNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDeleted) {
    return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasIsDeleted() bool {
	if o != nil && !isNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *CreateOrGetSubscriptionOrderDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetDeleterId() string {
	if o == nil || isNil(o.DeleterId) {
		var ret string
		return ret
	}
	return *o.DeleterId
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetDeleterIdOk() (*string, bool) {
	if o == nil || isNil(o.DeleterId) {
    return nil, false
	}
	return o.DeleterId, true
}

// HasDeleterId returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasDeleterId() bool {
	if o != nil && !isNil(o.DeleterId) {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given string and assigns it to the DeleterId field.
func (o *CreateOrGetSubscriptionOrderDto) SetDeleterId(v string) {
	o.DeleterId = &v
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetDeletionTime() time.Time {
	if o == nil || isNil(o.DeletionTime) {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeletionTime) {
    return nil, false
	}
	return o.DeletionTime, true
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasDeletionTime() bool {
	if o != nil && !isNil(o.DeletionTime) {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given time.Time and assigns it to the DeletionTime field.
func (o *CreateOrGetSubscriptionOrderDto) SetDeletionTime(v time.Time) {
	o.DeletionTime = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *CreateOrGetSubscriptionOrderDto) SetAppId(v string) {
	o.AppId = &v
}

// GetPricingId returns the PricingId field value if set, zero value otherwise.
func (o *CreateOrGetSubscriptionOrderDto) GetPricingId() string {
	if o == nil || isNil(o.PricingId) {
		var ret string
		return ret
	}
	return *o.PricingId
}

// GetPricingIdOk returns a tuple with the PricingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrGetSubscriptionOrderDto) GetPricingIdOk() (*string, bool) {
	if o == nil || isNil(o.PricingId) {
    return nil, false
	}
	return o.PricingId, true
}

// HasPricingId returns a boolean if a field has been set.
func (o *CreateOrGetSubscriptionOrderDto) HasPricingId() bool {
	if o != nil && !isNil(o.PricingId) {
		return true
	}

	return false
}

// SetPricingId gets a reference to the given string and assigns it to the PricingId field.
func (o *CreateOrGetSubscriptionOrderDto) SetPricingId(v string) {
	o.PricingId = &v
}

func (o CreateOrGetSubscriptionOrderDto) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.PricingId) {
		toSerialize["pricingId"] = o.PricingId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrGetSubscriptionOrderDto struct {
	value *CreateOrGetSubscriptionOrderDto
	isSet bool
}

func (v NullableCreateOrGetSubscriptionOrderDto) Get() *CreateOrGetSubscriptionOrderDto {
	return v.value
}

func (v *NullableCreateOrGetSubscriptionOrderDto) Set(val *CreateOrGetSubscriptionOrderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrGetSubscriptionOrderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrGetSubscriptionOrderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrGetSubscriptionOrderDto(val *CreateOrGetSubscriptionOrderDto) *NullableCreateOrGetSubscriptionOrderDto {
	return &NullableCreateOrGetSubscriptionOrderDto{value: val, isSet: true}
}

func (v NullableCreateOrGetSubscriptionOrderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrGetSubscriptionOrderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



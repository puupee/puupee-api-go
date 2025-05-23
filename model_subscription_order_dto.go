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
)

// checks if the SubscriptionOrderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOrderDto{}

// SubscriptionOrderDto struct for SubscriptionOrderDto
type SubscriptionOrderDto struct {
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	LastModificationTime *time.Time `json:"lastModificationTime,omitempty"`
	LastModifierId *string `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId *string `json:"deleterId,omitempty"`
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	Type *SubscriptionOrderType `json:"type,omitempty"`
	Status *SubscriptionOrderStatus `json:"status,omitempty"`
	AppId *string `json:"appId,omitempty"`
	PricingId *string `json:"pricingId,omitempty"`
	ProductId *string `json:"productId,omitempty"`
}

// NewSubscriptionOrderDto instantiates a new SubscriptionOrderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOrderDto() *SubscriptionOrderDto {
	this := SubscriptionOrderDto{}
	return &this
}

// NewSubscriptionOrderDtoWithDefaults instantiates a new SubscriptionOrderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOrderDtoWithDefaults() *SubscriptionOrderDto {
	this := SubscriptionOrderDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionOrderDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *SubscriptionOrderDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *SubscriptionOrderDto) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetLastModificationTime() time.Time {
	if o == nil || IsNil(o.LastModificationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModificationTime) {
		return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasLastModificationTime() bool {
	if o != nil && !IsNil(o.LastModificationTime) {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given time.Time and assigns it to the LastModificationTime field.
func (o *SubscriptionOrderDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime = &v
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetLastModifierId() string {
	if o == nil || IsNil(o.LastModifierId) {
		var ret string
		return ret
	}
	return *o.LastModifierId
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifierId) {
		return nil, false
	}
	return o.LastModifierId, true
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasLastModifierId() bool {
	if o != nil && !IsNil(o.LastModifierId) {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given string and assigns it to the LastModifierId field.
func (o *SubscriptionOrderDto) SetLastModifierId(v string) {
	o.LastModifierId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *SubscriptionOrderDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetDeleterId() string {
	if o == nil || IsNil(o.DeleterId) {
		var ret string
		return ret
	}
	return *o.DeleterId
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetDeleterIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeleterId) {
		return nil, false
	}
	return o.DeleterId, true
}

// HasDeleterId returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasDeleterId() bool {
	if o != nil && !IsNil(o.DeleterId) {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given string and assigns it to the DeleterId field.
func (o *SubscriptionOrderDto) SetDeleterId(v string) {
	o.DeleterId = &v
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetDeletionTime() time.Time {
	if o == nil || IsNil(o.DeletionTime) {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletionTime) {
		return nil, false
	}
	return o.DeletionTime, true
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasDeletionTime() bool {
	if o != nil && !IsNil(o.DeletionTime) {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given time.Time and assigns it to the DeletionTime field.
func (o *SubscriptionOrderDto) SetDeletionTime(v time.Time) {
	o.DeletionTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetType() SubscriptionOrderType {
	if o == nil || IsNil(o.Type) {
		var ret SubscriptionOrderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetTypeOk() (*SubscriptionOrderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SubscriptionOrderType and assigns it to the Type field.
func (o *SubscriptionOrderDto) SetType(v SubscriptionOrderType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetStatus() SubscriptionOrderStatus {
	if o == nil || IsNil(o.Status) {
		var ret SubscriptionOrderStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetStatusOk() (*SubscriptionOrderStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SubscriptionOrderStatus and assigns it to the Status field.
func (o *SubscriptionOrderDto) SetStatus(v SubscriptionOrderStatus) {
	o.Status = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *SubscriptionOrderDto) SetAppId(v string) {
	o.AppId = &v
}

// GetPricingId returns the PricingId field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetPricingId() string {
	if o == nil || IsNil(o.PricingId) {
		var ret string
		return ret
	}
	return *o.PricingId
}

// GetPricingIdOk returns a tuple with the PricingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetPricingIdOk() (*string, bool) {
	if o == nil || IsNil(o.PricingId) {
		return nil, false
	}
	return o.PricingId, true
}

// HasPricingId returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasPricingId() bool {
	if o != nil && !IsNil(o.PricingId) {
		return true
	}

	return false
}

// SetPricingId gets a reference to the given string and assigns it to the PricingId field.
func (o *SubscriptionOrderDto) SetPricingId(v string) {
	o.PricingId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *SubscriptionOrderDto) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOrderDto) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *SubscriptionOrderDto) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *SubscriptionOrderDto) SetProductId(v string) {
	o.ProductId = &v
}

func (o SubscriptionOrderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOrderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if !IsNil(o.LastModificationTime) {
		toSerialize["lastModificationTime"] = o.LastModificationTime
	}
	if !IsNil(o.LastModifierId) {
		toSerialize["lastModifierId"] = o.LastModifierId
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.DeleterId) {
		toSerialize["deleterId"] = o.DeleterId
	}
	if !IsNil(o.DeletionTime) {
		toSerialize["deletionTime"] = o.DeletionTime
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.PricingId) {
		toSerialize["pricingId"] = o.PricingId
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	return toSerialize, nil
}

type NullableSubscriptionOrderDto struct {
	value *SubscriptionOrderDto
	isSet bool
}

func (v NullableSubscriptionOrderDto) Get() *SubscriptionOrderDto {
	return v.value
}

func (v *NullableSubscriptionOrderDto) Set(val *SubscriptionOrderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOrderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOrderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOrderDto(val *SubscriptionOrderDto) *NullableSubscriptionOrderDto {
	return &NullableSubscriptionOrderDto{value: val, isSet: true}
}

func (v NullableSubscriptionOrderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOrderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



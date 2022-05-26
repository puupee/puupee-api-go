/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
	"time"
)

// NotificationDto struct for NotificationDto
type NotificationDto struct {
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId NullableString `json:"creatorId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	NotificationInfoId *string `json:"notificationInfoId,omitempty"`
	NotificationMethod NullableString `json:"notificationMethod,omitempty"`
	Success NullableBool `json:"success,omitempty"`
	CompletionTime NullableTime `json:"completionTime,omitempty"`
	FailureReason NullableString `json:"failureReason,omitempty"`
	RetryNotificationId NullableString `json:"retryNotificationId,omitempty"`
}

// NewNotificationDto instantiates a new NotificationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationDto() *NotificationDto {
	this := NotificationDto{}
	return &this
}

// NewNotificationDtoWithDefaults instantiates a new NotificationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationDtoWithDefaults() *NotificationDto {
	this := NotificationDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NotificationDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *NotificationDto) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *NotificationDto) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *NotificationDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationDto) GetCreatorId() string {
	if o == nil || o.CreatorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationDto) GetCreatorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *NotificationDto) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableString and assigns it to the CreatorId field.
func (o *NotificationDto) SetCreatorId(v string) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *NotificationDto) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *NotificationDto) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *NotificationDto) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDto) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *NotificationDto) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *NotificationDto) SetUserId(v string) {
	o.UserId = &v
}

// GetNotificationInfoId returns the NotificationInfoId field value if set, zero value otherwise.
func (o *NotificationDto) GetNotificationInfoId() string {
	if o == nil || o.NotificationInfoId == nil {
		var ret string
		return ret
	}
	return *o.NotificationInfoId
}

// GetNotificationInfoIdOk returns a tuple with the NotificationInfoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDto) GetNotificationInfoIdOk() (*string, bool) {
	if o == nil || o.NotificationInfoId == nil {
		return nil, false
	}
	return o.NotificationInfoId, true
}

// HasNotificationInfoId returns a boolean if a field has been set.
func (o *NotificationDto) HasNotificationInfoId() bool {
	if o != nil && o.NotificationInfoId != nil {
		return true
	}

	return false
}

// SetNotificationInfoId gets a reference to the given string and assigns it to the NotificationInfoId field.
func (o *NotificationDto) SetNotificationInfoId(v string) {
	o.NotificationInfoId = &v
}

// GetNotificationMethod returns the NotificationMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationDto) GetNotificationMethod() string {
	if o == nil || o.NotificationMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.NotificationMethod.Get()
}

// GetNotificationMethodOk returns a tuple with the NotificationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationDto) GetNotificationMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NotificationMethod.Get(), o.NotificationMethod.IsSet()
}

// HasNotificationMethod returns a boolean if a field has been set.
func (o *NotificationDto) HasNotificationMethod() bool {
	if o != nil && o.NotificationMethod.IsSet() {
		return true
	}

	return false
}

// SetNotificationMethod gets a reference to the given NullableString and assigns it to the NotificationMethod field.
func (o *NotificationDto) SetNotificationMethod(v string) {
	o.NotificationMethod.Set(&v)
}
// SetNotificationMethodNil sets the value for NotificationMethod to be an explicit nil
func (o *NotificationDto) SetNotificationMethodNil() {
	o.NotificationMethod.Set(nil)
}

// UnsetNotificationMethod ensures that no value is present for NotificationMethod, not even an explicit nil
func (o *NotificationDto) UnsetNotificationMethod() {
	o.NotificationMethod.Unset()
}

// GetSuccess returns the Success field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationDto) GetSuccess() bool {
	if o == nil || o.Success.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Success.Get()
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationDto) GetSuccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Success.Get(), o.Success.IsSet()
}

// HasSuccess returns a boolean if a field has been set.
func (o *NotificationDto) HasSuccess() bool {
	if o != nil && o.Success.IsSet() {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given NullableBool and assigns it to the Success field.
func (o *NotificationDto) SetSuccess(v bool) {
	o.Success.Set(&v)
}
// SetSuccessNil sets the value for Success to be an explicit nil
func (o *NotificationDto) SetSuccessNil() {
	o.Success.Set(nil)
}

// UnsetSuccess ensures that no value is present for Success, not even an explicit nil
func (o *NotificationDto) UnsetSuccess() {
	o.Success.Unset()
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationDto) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime.Get()
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationDto) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompletionTime.Get(), o.CompletionTime.IsSet()
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *NotificationDto) HasCompletionTime() bool {
	if o != nil && o.CompletionTime.IsSet() {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given NullableTime and assigns it to the CompletionTime field.
func (o *NotificationDto) SetCompletionTime(v time.Time) {
	o.CompletionTime.Set(&v)
}
// SetCompletionTimeNil sets the value for CompletionTime to be an explicit nil
func (o *NotificationDto) SetCompletionTimeNil() {
	o.CompletionTime.Set(nil)
}

// UnsetCompletionTime ensures that no value is present for CompletionTime, not even an explicit nil
func (o *NotificationDto) UnsetCompletionTime() {
	o.CompletionTime.Unset()
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationDto) GetFailureReason() string {
	if o == nil || o.FailureReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationDto) GetFailureReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// HasFailureReason returns a boolean if a field has been set.
func (o *NotificationDto) HasFailureReason() bool {
	if o != nil && o.FailureReason.IsSet() {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given NullableString and assigns it to the FailureReason field.
func (o *NotificationDto) SetFailureReason(v string) {
	o.FailureReason.Set(&v)
}
// SetFailureReasonNil sets the value for FailureReason to be an explicit nil
func (o *NotificationDto) SetFailureReasonNil() {
	o.FailureReason.Set(nil)
}

// UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
func (o *NotificationDto) UnsetFailureReason() {
	o.FailureReason.Unset()
}

// GetRetryNotificationId returns the RetryNotificationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationDto) GetRetryNotificationId() string {
	if o == nil || o.RetryNotificationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RetryNotificationId.Get()
}

// GetRetryNotificationIdOk returns a tuple with the RetryNotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationDto) GetRetryNotificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RetryNotificationId.Get(), o.RetryNotificationId.IsSet()
}

// HasRetryNotificationId returns a boolean if a field has been set.
func (o *NotificationDto) HasRetryNotificationId() bool {
	if o != nil && o.RetryNotificationId.IsSet() {
		return true
	}

	return false
}

// SetRetryNotificationId gets a reference to the given NullableString and assigns it to the RetryNotificationId field.
func (o *NotificationDto) SetRetryNotificationId(v string) {
	o.RetryNotificationId.Set(&v)
}
// SetRetryNotificationIdNil sets the value for RetryNotificationId to be an explicit nil
func (o *NotificationDto) SetRetryNotificationIdNil() {
	o.RetryNotificationId.Set(nil)
}

// UnsetRetryNotificationId ensures that no value is present for RetryNotificationId, not even an explicit nil
func (o *NotificationDto) UnsetRetryNotificationId() {
	o.RetryNotificationId.Unset()
}

func (o NotificationDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.CreatorId.IsSet() {
		toSerialize["creatorId"] = o.CreatorId.Get()
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.NotificationInfoId != nil {
		toSerialize["notificationInfoId"] = o.NotificationInfoId
	}
	if o.NotificationMethod.IsSet() {
		toSerialize["notificationMethod"] = o.NotificationMethod.Get()
	}
	if o.Success.IsSet() {
		toSerialize["success"] = o.Success.Get()
	}
	if o.CompletionTime.IsSet() {
		toSerialize["completionTime"] = o.CompletionTime.Get()
	}
	if o.FailureReason.IsSet() {
		toSerialize["failureReason"] = o.FailureReason.Get()
	}
	if o.RetryNotificationId.IsSet() {
		toSerialize["retryNotificationId"] = o.RetryNotificationId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationDto struct {
	value *NotificationDto
	isSet bool
}

func (v NullableNotificationDto) Get() *NotificationDto {
	return v.value
}

func (v *NullableNotificationDto) Set(val *NotificationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationDto(val *NotificationDto) *NullableNotificationDto {
	return &NullableNotificationDto{value: val, isSet: true}
}

func (v NullableNotificationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



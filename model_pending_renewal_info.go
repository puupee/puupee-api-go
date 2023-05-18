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

// checks if the PendingRenewalInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingRenewalInfo{}

// PendingRenewalInfo struct for PendingRenewalInfo
type PendingRenewalInfo struct {
	AutoRenewProductId NullableString `json:"auto_renew_product_id,omitempty"`
	AutoRenewStatus NullableString `json:"auto_renew_status,omitempty"`
	IsInBillingRetryPeriod NullableString `json:"is_in_billing_retry_period,omitempty"`
	OriginalTransactionId NullableString `json:"original_transaction_id,omitempty"`
	ProductId NullableString `json:"product_id,omitempty"`
	ExpirationIntent NullableString `json:"expiration_intent,omitempty"`
	PriceConsentStatus NullableString `json:"price_consent_status,omitempty"`
	GracePeriodExpiresDate NullableString `json:"grace_period_expires_date,omitempty"`
	GracePeriodExpiresDateMs NullableString `json:"grace_period_expires_date_ms,omitempty"`
	GracePeriodExpiresDatePst NullableString `json:"grace_period_expires_date_pst,omitempty"`
}

// NewPendingRenewalInfo instantiates a new PendingRenewalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingRenewalInfo() *PendingRenewalInfo {
	this := PendingRenewalInfo{}
	return &this
}

// NewPendingRenewalInfoWithDefaults instantiates a new PendingRenewalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingRenewalInfoWithDefaults() *PendingRenewalInfo {
	this := PendingRenewalInfo{}
	return &this
}

// GetAutoRenewProductId returns the AutoRenewProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetAutoRenewProductId() string {
	if o == nil || IsNil(o.AutoRenewProductId.Get()) {
		var ret string
		return ret
	}
	return *o.AutoRenewProductId.Get()
}

// GetAutoRenewProductIdOk returns a tuple with the AutoRenewProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetAutoRenewProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoRenewProductId.Get(), o.AutoRenewProductId.IsSet()
}

// HasAutoRenewProductId returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasAutoRenewProductId() bool {
	if o != nil && o.AutoRenewProductId.IsSet() {
		return true
	}

	return false
}

// SetAutoRenewProductId gets a reference to the given NullableString and assigns it to the AutoRenewProductId field.
func (o *PendingRenewalInfo) SetAutoRenewProductId(v string) {
	o.AutoRenewProductId.Set(&v)
}
// SetAutoRenewProductIdNil sets the value for AutoRenewProductId to be an explicit nil
func (o *PendingRenewalInfo) SetAutoRenewProductIdNil() {
	o.AutoRenewProductId.Set(nil)
}

// UnsetAutoRenewProductId ensures that no value is present for AutoRenewProductId, not even an explicit nil
func (o *PendingRenewalInfo) UnsetAutoRenewProductId() {
	o.AutoRenewProductId.Unset()
}

// GetAutoRenewStatus returns the AutoRenewStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetAutoRenewStatus() string {
	if o == nil || IsNil(o.AutoRenewStatus.Get()) {
		var ret string
		return ret
	}
	return *o.AutoRenewStatus.Get()
}

// GetAutoRenewStatusOk returns a tuple with the AutoRenewStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetAutoRenewStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoRenewStatus.Get(), o.AutoRenewStatus.IsSet()
}

// HasAutoRenewStatus returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasAutoRenewStatus() bool {
	if o != nil && o.AutoRenewStatus.IsSet() {
		return true
	}

	return false
}

// SetAutoRenewStatus gets a reference to the given NullableString and assigns it to the AutoRenewStatus field.
func (o *PendingRenewalInfo) SetAutoRenewStatus(v string) {
	o.AutoRenewStatus.Set(&v)
}
// SetAutoRenewStatusNil sets the value for AutoRenewStatus to be an explicit nil
func (o *PendingRenewalInfo) SetAutoRenewStatusNil() {
	o.AutoRenewStatus.Set(nil)
}

// UnsetAutoRenewStatus ensures that no value is present for AutoRenewStatus, not even an explicit nil
func (o *PendingRenewalInfo) UnsetAutoRenewStatus() {
	o.AutoRenewStatus.Unset()
}

// GetIsInBillingRetryPeriod returns the IsInBillingRetryPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetIsInBillingRetryPeriod() string {
	if o == nil || IsNil(o.IsInBillingRetryPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.IsInBillingRetryPeriod.Get()
}

// GetIsInBillingRetryPeriodOk returns a tuple with the IsInBillingRetryPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetIsInBillingRetryPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInBillingRetryPeriod.Get(), o.IsInBillingRetryPeriod.IsSet()
}

// HasIsInBillingRetryPeriod returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasIsInBillingRetryPeriod() bool {
	if o != nil && o.IsInBillingRetryPeriod.IsSet() {
		return true
	}

	return false
}

// SetIsInBillingRetryPeriod gets a reference to the given NullableString and assigns it to the IsInBillingRetryPeriod field.
func (o *PendingRenewalInfo) SetIsInBillingRetryPeriod(v string) {
	o.IsInBillingRetryPeriod.Set(&v)
}
// SetIsInBillingRetryPeriodNil sets the value for IsInBillingRetryPeriod to be an explicit nil
func (o *PendingRenewalInfo) SetIsInBillingRetryPeriodNil() {
	o.IsInBillingRetryPeriod.Set(nil)
}

// UnsetIsInBillingRetryPeriod ensures that no value is present for IsInBillingRetryPeriod, not even an explicit nil
func (o *PendingRenewalInfo) UnsetIsInBillingRetryPeriod() {
	o.IsInBillingRetryPeriod.Unset()
}

// GetOriginalTransactionId returns the OriginalTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetOriginalTransactionId() string {
	if o == nil || IsNil(o.OriginalTransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalTransactionId.Get()
}

// GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetOriginalTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalTransactionId.Get(), o.OriginalTransactionId.IsSet()
}

// HasOriginalTransactionId returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasOriginalTransactionId() bool {
	if o != nil && o.OriginalTransactionId.IsSet() {
		return true
	}

	return false
}

// SetOriginalTransactionId gets a reference to the given NullableString and assigns it to the OriginalTransactionId field.
func (o *PendingRenewalInfo) SetOriginalTransactionId(v string) {
	o.OriginalTransactionId.Set(&v)
}
// SetOriginalTransactionIdNil sets the value for OriginalTransactionId to be an explicit nil
func (o *PendingRenewalInfo) SetOriginalTransactionIdNil() {
	o.OriginalTransactionId.Set(nil)
}

// UnsetOriginalTransactionId ensures that no value is present for OriginalTransactionId, not even an explicit nil
func (o *PendingRenewalInfo) UnsetOriginalTransactionId() {
	o.OriginalTransactionId.Unset()
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetProductId() string {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret string
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableString and assigns it to the ProductId field.
func (o *PendingRenewalInfo) SetProductId(v string) {
	o.ProductId.Set(&v)
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *PendingRenewalInfo) SetProductIdNil() {
	o.ProductId.Set(nil)
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *PendingRenewalInfo) UnsetProductId() {
	o.ProductId.Unset()
}

// GetExpirationIntent returns the ExpirationIntent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetExpirationIntent() string {
	if o == nil || IsNil(o.ExpirationIntent.Get()) {
		var ret string
		return ret
	}
	return *o.ExpirationIntent.Get()
}

// GetExpirationIntentOk returns a tuple with the ExpirationIntent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetExpirationIntentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationIntent.Get(), o.ExpirationIntent.IsSet()
}

// HasExpirationIntent returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasExpirationIntent() bool {
	if o != nil && o.ExpirationIntent.IsSet() {
		return true
	}

	return false
}

// SetExpirationIntent gets a reference to the given NullableString and assigns it to the ExpirationIntent field.
func (o *PendingRenewalInfo) SetExpirationIntent(v string) {
	o.ExpirationIntent.Set(&v)
}
// SetExpirationIntentNil sets the value for ExpirationIntent to be an explicit nil
func (o *PendingRenewalInfo) SetExpirationIntentNil() {
	o.ExpirationIntent.Set(nil)
}

// UnsetExpirationIntent ensures that no value is present for ExpirationIntent, not even an explicit nil
func (o *PendingRenewalInfo) UnsetExpirationIntent() {
	o.ExpirationIntent.Unset()
}

// GetPriceConsentStatus returns the PriceConsentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetPriceConsentStatus() string {
	if o == nil || IsNil(o.PriceConsentStatus.Get()) {
		var ret string
		return ret
	}
	return *o.PriceConsentStatus.Get()
}

// GetPriceConsentStatusOk returns a tuple with the PriceConsentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetPriceConsentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceConsentStatus.Get(), o.PriceConsentStatus.IsSet()
}

// HasPriceConsentStatus returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasPriceConsentStatus() bool {
	if o != nil && o.PriceConsentStatus.IsSet() {
		return true
	}

	return false
}

// SetPriceConsentStatus gets a reference to the given NullableString and assigns it to the PriceConsentStatus field.
func (o *PendingRenewalInfo) SetPriceConsentStatus(v string) {
	o.PriceConsentStatus.Set(&v)
}
// SetPriceConsentStatusNil sets the value for PriceConsentStatus to be an explicit nil
func (o *PendingRenewalInfo) SetPriceConsentStatusNil() {
	o.PriceConsentStatus.Set(nil)
}

// UnsetPriceConsentStatus ensures that no value is present for PriceConsentStatus, not even an explicit nil
func (o *PendingRenewalInfo) UnsetPriceConsentStatus() {
	o.PriceConsentStatus.Unset()
}

// GetGracePeriodExpiresDate returns the GracePeriodExpiresDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetGracePeriodExpiresDate() string {
	if o == nil || IsNil(o.GracePeriodExpiresDate.Get()) {
		var ret string
		return ret
	}
	return *o.GracePeriodExpiresDate.Get()
}

// GetGracePeriodExpiresDateOk returns a tuple with the GracePeriodExpiresDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetGracePeriodExpiresDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GracePeriodExpiresDate.Get(), o.GracePeriodExpiresDate.IsSet()
}

// HasGracePeriodExpiresDate returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasGracePeriodExpiresDate() bool {
	if o != nil && o.GracePeriodExpiresDate.IsSet() {
		return true
	}

	return false
}

// SetGracePeriodExpiresDate gets a reference to the given NullableString and assigns it to the GracePeriodExpiresDate field.
func (o *PendingRenewalInfo) SetGracePeriodExpiresDate(v string) {
	o.GracePeriodExpiresDate.Set(&v)
}
// SetGracePeriodExpiresDateNil sets the value for GracePeriodExpiresDate to be an explicit nil
func (o *PendingRenewalInfo) SetGracePeriodExpiresDateNil() {
	o.GracePeriodExpiresDate.Set(nil)
}

// UnsetGracePeriodExpiresDate ensures that no value is present for GracePeriodExpiresDate, not even an explicit nil
func (o *PendingRenewalInfo) UnsetGracePeriodExpiresDate() {
	o.GracePeriodExpiresDate.Unset()
}

// GetGracePeriodExpiresDateMs returns the GracePeriodExpiresDateMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetGracePeriodExpiresDateMs() string {
	if o == nil || IsNil(o.GracePeriodExpiresDateMs.Get()) {
		var ret string
		return ret
	}
	return *o.GracePeriodExpiresDateMs.Get()
}

// GetGracePeriodExpiresDateMsOk returns a tuple with the GracePeriodExpiresDateMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetGracePeriodExpiresDateMsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GracePeriodExpiresDateMs.Get(), o.GracePeriodExpiresDateMs.IsSet()
}

// HasGracePeriodExpiresDateMs returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasGracePeriodExpiresDateMs() bool {
	if o != nil && o.GracePeriodExpiresDateMs.IsSet() {
		return true
	}

	return false
}

// SetGracePeriodExpiresDateMs gets a reference to the given NullableString and assigns it to the GracePeriodExpiresDateMs field.
func (o *PendingRenewalInfo) SetGracePeriodExpiresDateMs(v string) {
	o.GracePeriodExpiresDateMs.Set(&v)
}
// SetGracePeriodExpiresDateMsNil sets the value for GracePeriodExpiresDateMs to be an explicit nil
func (o *PendingRenewalInfo) SetGracePeriodExpiresDateMsNil() {
	o.GracePeriodExpiresDateMs.Set(nil)
}

// UnsetGracePeriodExpiresDateMs ensures that no value is present for GracePeriodExpiresDateMs, not even an explicit nil
func (o *PendingRenewalInfo) UnsetGracePeriodExpiresDateMs() {
	o.GracePeriodExpiresDateMs.Unset()
}

// GetGracePeriodExpiresDatePst returns the GracePeriodExpiresDatePst field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingRenewalInfo) GetGracePeriodExpiresDatePst() string {
	if o == nil || IsNil(o.GracePeriodExpiresDatePst.Get()) {
		var ret string
		return ret
	}
	return *o.GracePeriodExpiresDatePst.Get()
}

// GetGracePeriodExpiresDatePstOk returns a tuple with the GracePeriodExpiresDatePst field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingRenewalInfo) GetGracePeriodExpiresDatePstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GracePeriodExpiresDatePst.Get(), o.GracePeriodExpiresDatePst.IsSet()
}

// HasGracePeriodExpiresDatePst returns a boolean if a field has been set.
func (o *PendingRenewalInfo) HasGracePeriodExpiresDatePst() bool {
	if o != nil && o.GracePeriodExpiresDatePst.IsSet() {
		return true
	}

	return false
}

// SetGracePeriodExpiresDatePst gets a reference to the given NullableString and assigns it to the GracePeriodExpiresDatePst field.
func (o *PendingRenewalInfo) SetGracePeriodExpiresDatePst(v string) {
	o.GracePeriodExpiresDatePst.Set(&v)
}
// SetGracePeriodExpiresDatePstNil sets the value for GracePeriodExpiresDatePst to be an explicit nil
func (o *PendingRenewalInfo) SetGracePeriodExpiresDatePstNil() {
	o.GracePeriodExpiresDatePst.Set(nil)
}

// UnsetGracePeriodExpiresDatePst ensures that no value is present for GracePeriodExpiresDatePst, not even an explicit nil
func (o *PendingRenewalInfo) UnsetGracePeriodExpiresDatePst() {
	o.GracePeriodExpiresDatePst.Unset()
}

func (o PendingRenewalInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingRenewalInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoRenewProductId.IsSet() {
		toSerialize["auto_renew_product_id"] = o.AutoRenewProductId.Get()
	}
	if o.AutoRenewStatus.IsSet() {
		toSerialize["auto_renew_status"] = o.AutoRenewStatus.Get()
	}
	if o.IsInBillingRetryPeriod.IsSet() {
		toSerialize["is_in_billing_retry_period"] = o.IsInBillingRetryPeriod.Get()
	}
	if o.OriginalTransactionId.IsSet() {
		toSerialize["original_transaction_id"] = o.OriginalTransactionId.Get()
	}
	if o.ProductId.IsSet() {
		toSerialize["product_id"] = o.ProductId.Get()
	}
	if o.ExpirationIntent.IsSet() {
		toSerialize["expiration_intent"] = o.ExpirationIntent.Get()
	}
	if o.PriceConsentStatus.IsSet() {
		toSerialize["price_consent_status"] = o.PriceConsentStatus.Get()
	}
	if o.GracePeriodExpiresDate.IsSet() {
		toSerialize["grace_period_expires_date"] = o.GracePeriodExpiresDate.Get()
	}
	if o.GracePeriodExpiresDateMs.IsSet() {
		toSerialize["grace_period_expires_date_ms"] = o.GracePeriodExpiresDateMs.Get()
	}
	if o.GracePeriodExpiresDatePst.IsSet() {
		toSerialize["grace_period_expires_date_pst"] = o.GracePeriodExpiresDatePst.Get()
	}
	return toSerialize, nil
}

type NullablePendingRenewalInfo struct {
	value *PendingRenewalInfo
	isSet bool
}

func (v NullablePendingRenewalInfo) Get() *PendingRenewalInfo {
	return v.value
}

func (v *NullablePendingRenewalInfo) Set(val *PendingRenewalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingRenewalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingRenewalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingRenewalInfo(val *PendingRenewalInfo) *NullablePendingRenewalInfo {
	return &NullablePendingRenewalInfo{value: val, isSet: true}
}

func (v NullablePendingRenewalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingRenewalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



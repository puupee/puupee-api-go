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

// AppleVerifyReceiptResult struct for AppleVerifyReceiptResult
type AppleVerifyReceiptResult struct {
	Environment *string `json:"environment,omitempty"`
	IsRetryable *bool `json:"is_retryable,omitempty"`
	Status *int32 `json:"status,omitempty"`
	LatestReceiptInfo []LatestReceiptInfo `json:"latest_receipt_info,omitempty"`
	LatestReceipt *string `json:"latest_receipt,omitempty"`
	PendingRenewalInfo []PendingRenewalInfo `json:"pending_renewal_info,omitempty"`
	Receipt *Receipt `json:"receipt,omitempty"`
}

// NewAppleVerifyReceiptResult instantiates a new AppleVerifyReceiptResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppleVerifyReceiptResult() *AppleVerifyReceiptResult {
	this := AppleVerifyReceiptResult{}
	return &this
}

// NewAppleVerifyReceiptResultWithDefaults instantiates a new AppleVerifyReceiptResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleVerifyReceiptResultWithDefaults() *AppleVerifyReceiptResult {
	this := AppleVerifyReceiptResult{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AppleVerifyReceiptResult) GetEnvironment() string {
	if o == nil || isNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleVerifyReceiptResult) GetEnvironmentOk() (*string, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AppleVerifyReceiptResult) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *AppleVerifyReceiptResult) SetEnvironment(v string) {
	o.Environment = &v
}

// GetIsRetryable returns the IsRetryable field value if set, zero value otherwise.
func (o *AppleVerifyReceiptResult) GetIsRetryable() bool {
	if o == nil || isNil(o.IsRetryable) {
		var ret bool
		return ret
	}
	return *o.IsRetryable
}

// GetIsRetryableOk returns a tuple with the IsRetryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleVerifyReceiptResult) GetIsRetryableOk() (*bool, bool) {
	if o == nil || isNil(o.IsRetryable) {
    return nil, false
	}
	return o.IsRetryable, true
}

// HasIsRetryable returns a boolean if a field has been set.
func (o *AppleVerifyReceiptResult) HasIsRetryable() bool {
	if o != nil && !isNil(o.IsRetryable) {
		return true
	}

	return false
}

// SetIsRetryable gets a reference to the given bool and assigns it to the IsRetryable field.
func (o *AppleVerifyReceiptResult) SetIsRetryable(v bool) {
	o.IsRetryable = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppleVerifyReceiptResult) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleVerifyReceiptResult) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppleVerifyReceiptResult) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *AppleVerifyReceiptResult) SetStatus(v int32) {
	o.Status = &v
}

// GetLatestReceiptInfo returns the LatestReceiptInfo field value if set, zero value otherwise.
func (o *AppleVerifyReceiptResult) GetLatestReceiptInfo() []LatestReceiptInfo {
	if o == nil || isNil(o.LatestReceiptInfo) {
		var ret []LatestReceiptInfo
		return ret
	}
	return o.LatestReceiptInfo
}

// GetLatestReceiptInfoOk returns a tuple with the LatestReceiptInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleVerifyReceiptResult) GetLatestReceiptInfoOk() ([]LatestReceiptInfo, bool) {
	if o == nil || isNil(o.LatestReceiptInfo) {
    return nil, false
	}
	return o.LatestReceiptInfo, true
}

// HasLatestReceiptInfo returns a boolean if a field has been set.
func (o *AppleVerifyReceiptResult) HasLatestReceiptInfo() bool {
	if o != nil && !isNil(o.LatestReceiptInfo) {
		return true
	}

	return false
}

// SetLatestReceiptInfo gets a reference to the given []LatestReceiptInfo and assigns it to the LatestReceiptInfo field.
func (o *AppleVerifyReceiptResult) SetLatestReceiptInfo(v []LatestReceiptInfo) {
	o.LatestReceiptInfo = v
}

// GetLatestReceipt returns the LatestReceipt field value if set, zero value otherwise.
func (o *AppleVerifyReceiptResult) GetLatestReceipt() string {
	if o == nil || isNil(o.LatestReceipt) {
		var ret string
		return ret
	}
	return *o.LatestReceipt
}

// GetLatestReceiptOk returns a tuple with the LatestReceipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleVerifyReceiptResult) GetLatestReceiptOk() (*string, bool) {
	if o == nil || isNil(o.LatestReceipt) {
    return nil, false
	}
	return o.LatestReceipt, true
}

// HasLatestReceipt returns a boolean if a field has been set.
func (o *AppleVerifyReceiptResult) HasLatestReceipt() bool {
	if o != nil && !isNil(o.LatestReceipt) {
		return true
	}

	return false
}

// SetLatestReceipt gets a reference to the given string and assigns it to the LatestReceipt field.
func (o *AppleVerifyReceiptResult) SetLatestReceipt(v string) {
	o.LatestReceipt = &v
}

// GetPendingRenewalInfo returns the PendingRenewalInfo field value if set, zero value otherwise.
func (o *AppleVerifyReceiptResult) GetPendingRenewalInfo() []PendingRenewalInfo {
	if o == nil || isNil(o.PendingRenewalInfo) {
		var ret []PendingRenewalInfo
		return ret
	}
	return o.PendingRenewalInfo
}

// GetPendingRenewalInfoOk returns a tuple with the PendingRenewalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleVerifyReceiptResult) GetPendingRenewalInfoOk() ([]PendingRenewalInfo, bool) {
	if o == nil || isNil(o.PendingRenewalInfo) {
    return nil, false
	}
	return o.PendingRenewalInfo, true
}

// HasPendingRenewalInfo returns a boolean if a field has been set.
func (o *AppleVerifyReceiptResult) HasPendingRenewalInfo() bool {
	if o != nil && !isNil(o.PendingRenewalInfo) {
		return true
	}

	return false
}

// SetPendingRenewalInfo gets a reference to the given []PendingRenewalInfo and assigns it to the PendingRenewalInfo field.
func (o *AppleVerifyReceiptResult) SetPendingRenewalInfo(v []PendingRenewalInfo) {
	o.PendingRenewalInfo = v
}

// GetReceipt returns the Receipt field value if set, zero value otherwise.
func (o *AppleVerifyReceiptResult) GetReceipt() Receipt {
	if o == nil || isNil(o.Receipt) {
		var ret Receipt
		return ret
	}
	return *o.Receipt
}

// GetReceiptOk returns a tuple with the Receipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleVerifyReceiptResult) GetReceiptOk() (*Receipt, bool) {
	if o == nil || isNil(o.Receipt) {
    return nil, false
	}
	return o.Receipt, true
}

// HasReceipt returns a boolean if a field has been set.
func (o *AppleVerifyReceiptResult) HasReceipt() bool {
	if o != nil && !isNil(o.Receipt) {
		return true
	}

	return false
}

// SetReceipt gets a reference to the given Receipt and assigns it to the Receipt field.
func (o *AppleVerifyReceiptResult) SetReceipt(v Receipt) {
	o.Receipt = &v
}

func (o AppleVerifyReceiptResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !isNil(o.IsRetryable) {
		toSerialize["is_retryable"] = o.IsRetryable
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.LatestReceiptInfo) {
		toSerialize["latest_receipt_info"] = o.LatestReceiptInfo
	}
	if !isNil(o.LatestReceipt) {
		toSerialize["latest_receipt"] = o.LatestReceipt
	}
	if !isNil(o.PendingRenewalInfo) {
		toSerialize["pending_renewal_info"] = o.PendingRenewalInfo
	}
	if !isNil(o.Receipt) {
		toSerialize["receipt"] = o.Receipt
	}
	return json.Marshal(toSerialize)
}

type NullableAppleVerifyReceiptResult struct {
	value *AppleVerifyReceiptResult
	isSet bool
}

func (v NullableAppleVerifyReceiptResult) Get() *AppleVerifyReceiptResult {
	return v.value
}

func (v *NullableAppleVerifyReceiptResult) Set(val *AppleVerifyReceiptResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAppleVerifyReceiptResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAppleVerifyReceiptResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppleVerifyReceiptResult(val *AppleVerifyReceiptResult) *NullableAppleVerifyReceiptResult {
	return &NullableAppleVerifyReceiptResult{value: val, isSet: true}
}

func (v NullableAppleVerifyReceiptResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppleVerifyReceiptResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


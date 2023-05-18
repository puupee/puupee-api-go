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

// checks if the RemoteServiceErrorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteServiceErrorInfo{}

// RemoteServiceErrorInfo struct for RemoteServiceErrorInfo
type RemoteServiceErrorInfo struct {
	Code NullableString `json:"code,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Details NullableString `json:"details,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	ValidationErrors []RemoteServiceValidationErrorInfo `json:"validationErrors,omitempty"`
}

// NewRemoteServiceErrorInfo instantiates a new RemoteServiceErrorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteServiceErrorInfo() *RemoteServiceErrorInfo {
	this := RemoteServiceErrorInfo{}
	return &this
}

// NewRemoteServiceErrorInfoWithDefaults instantiates a new RemoteServiceErrorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteServiceErrorInfoWithDefaults() *RemoteServiceErrorInfo {
	this := RemoteServiceErrorInfo{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteServiceErrorInfo) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteServiceErrorInfo) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *RemoteServiceErrorInfo) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *RemoteServiceErrorInfo) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *RemoteServiceErrorInfo) UnsetCode() {
	o.Code.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteServiceErrorInfo) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteServiceErrorInfo) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *RemoteServiceErrorInfo) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *RemoteServiceErrorInfo) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *RemoteServiceErrorInfo) UnsetMessage() {
	o.Message.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteServiceErrorInfo) GetDetails() string {
	if o == nil || IsNil(o.Details.Get()) {
		var ret string
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteServiceErrorInfo) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableString and assigns it to the Details field.
func (o *RemoteServiceErrorInfo) SetDetails(v string) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *RemoteServiceErrorInfo) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *RemoteServiceErrorInfo) UnsetDetails() {
	o.Details.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteServiceErrorInfo) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteServiceErrorInfo) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *RemoteServiceErrorInfo) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteServiceErrorInfo) GetValidationErrors() []RemoteServiceValidationErrorInfo {
	if o == nil {
		var ret []RemoteServiceValidationErrorInfo
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteServiceErrorInfo) GetValidationErrorsOk() ([]RemoteServiceValidationErrorInfo, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasValidationErrors() bool {
	if o != nil && IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []RemoteServiceValidationErrorInfo and assigns it to the ValidationErrors field.
func (o *RemoteServiceErrorInfo) SetValidationErrors(v []RemoteServiceValidationErrorInfo) {
	o.ValidationErrors = v
}

func (o RemoteServiceErrorInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteServiceErrorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableRemoteServiceErrorInfo struct {
	value *RemoteServiceErrorInfo
	isSet bool
}

func (v NullableRemoteServiceErrorInfo) Get() *RemoteServiceErrorInfo {
	return v.value
}

func (v *NullableRemoteServiceErrorInfo) Set(val *RemoteServiceErrorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteServiceErrorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteServiceErrorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteServiceErrorInfo(val *RemoteServiceErrorInfo) *NullableRemoteServiceErrorInfo {
	return &NullableRemoteServiceErrorInfo{value: val, isSet: true}
}

func (v NullableRemoteServiceErrorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteServiceErrorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



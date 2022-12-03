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

// RemoteServiceErrorInfo struct for RemoteServiceErrorInfo
type RemoteServiceErrorInfo struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Details *string `json:"details,omitempty"`
	Data map[string]map[string]interface{} `json:"data,omitempty"`
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

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RemoteServiceErrorInfo) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteServiceErrorInfo) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *RemoteServiceErrorInfo) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RemoteServiceErrorInfo) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteServiceErrorInfo) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RemoteServiceErrorInfo) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RemoteServiceErrorInfo) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteServiceErrorInfo) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *RemoteServiceErrorInfo) SetDetails(v string) {
	o.Details = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RemoteServiceErrorInfo) GetData() map[string]map[string]interface{} {
	if o == nil || isNil(o.Data) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteServiceErrorInfo) GetDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.Data) {
    return map[string]map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]map[string]interface{} and assigns it to the Data field.
func (o *RemoteServiceErrorInfo) SetData(v map[string]map[string]interface{}) {
	o.Data = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *RemoteServiceErrorInfo) GetValidationErrors() []RemoteServiceValidationErrorInfo {
	if o == nil || isNil(o.ValidationErrors) {
		var ret []RemoteServiceValidationErrorInfo
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteServiceErrorInfo) GetValidationErrorsOk() ([]RemoteServiceValidationErrorInfo, bool) {
	if o == nil || isNil(o.ValidationErrors) {
    return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *RemoteServiceErrorInfo) HasValidationErrors() bool {
	if o != nil && !isNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []RemoteServiceValidationErrorInfo and assigns it to the ValidationErrors field.
func (o *RemoteServiceErrorInfo) SetValidationErrors(v []RemoteServiceValidationErrorInfo) {
	o.ValidationErrors = v
}

func (o RemoteServiceErrorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}
	return json.Marshal(toSerialize)
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


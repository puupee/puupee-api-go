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

// ControllerApiDescriptionModel struct for ControllerApiDescriptionModel
type ControllerApiDescriptionModel struct {
	ControllerName *string `json:"controllerName,omitempty"`
	ControllerGroupName *string `json:"controllerGroupName,omitempty"`
	IsRemoteService *bool `json:"isRemoteService,omitempty"`
	IsIntegrationService *bool `json:"isIntegrationService,omitempty"`
	ApiVersion *string `json:"apiVersion,omitempty"`
	Type *string `json:"type,omitempty"`
	Interfaces []ControllerInterfaceApiDescriptionModel `json:"interfaces,omitempty"`
	Actions *map[string]ActionApiDescriptionModel `json:"actions,omitempty"`
}

// NewControllerApiDescriptionModel instantiates a new ControllerApiDescriptionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerApiDescriptionModel() *ControllerApiDescriptionModel {
	this := ControllerApiDescriptionModel{}
	return &this
}

// NewControllerApiDescriptionModelWithDefaults instantiates a new ControllerApiDescriptionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerApiDescriptionModelWithDefaults() *ControllerApiDescriptionModel {
	this := ControllerApiDescriptionModel{}
	return &this
}

// GetControllerName returns the ControllerName field value if set, zero value otherwise.
func (o *ControllerApiDescriptionModel) GetControllerName() string {
	if o == nil || isNil(o.ControllerName) {
		var ret string
		return ret
	}
	return *o.ControllerName
}

// GetControllerNameOk returns a tuple with the ControllerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerApiDescriptionModel) GetControllerNameOk() (*string, bool) {
	if o == nil || isNil(o.ControllerName) {
    return nil, false
	}
	return o.ControllerName, true
}

// HasControllerName returns a boolean if a field has been set.
func (o *ControllerApiDescriptionModel) HasControllerName() bool {
	if o != nil && !isNil(o.ControllerName) {
		return true
	}

	return false
}

// SetControllerName gets a reference to the given string and assigns it to the ControllerName field.
func (o *ControllerApiDescriptionModel) SetControllerName(v string) {
	o.ControllerName = &v
}

// GetControllerGroupName returns the ControllerGroupName field value if set, zero value otherwise.
func (o *ControllerApiDescriptionModel) GetControllerGroupName() string {
	if o == nil || isNil(o.ControllerGroupName) {
		var ret string
		return ret
	}
	return *o.ControllerGroupName
}

// GetControllerGroupNameOk returns a tuple with the ControllerGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerApiDescriptionModel) GetControllerGroupNameOk() (*string, bool) {
	if o == nil || isNil(o.ControllerGroupName) {
    return nil, false
	}
	return o.ControllerGroupName, true
}

// HasControllerGroupName returns a boolean if a field has been set.
func (o *ControllerApiDescriptionModel) HasControllerGroupName() bool {
	if o != nil && !isNil(o.ControllerGroupName) {
		return true
	}

	return false
}

// SetControllerGroupName gets a reference to the given string and assigns it to the ControllerGroupName field.
func (o *ControllerApiDescriptionModel) SetControllerGroupName(v string) {
	o.ControllerGroupName = &v
}

// GetIsRemoteService returns the IsRemoteService field value if set, zero value otherwise.
func (o *ControllerApiDescriptionModel) GetIsRemoteService() bool {
	if o == nil || isNil(o.IsRemoteService) {
		var ret bool
		return ret
	}
	return *o.IsRemoteService
}

// GetIsRemoteServiceOk returns a tuple with the IsRemoteService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerApiDescriptionModel) GetIsRemoteServiceOk() (*bool, bool) {
	if o == nil || isNil(o.IsRemoteService) {
    return nil, false
	}
	return o.IsRemoteService, true
}

// HasIsRemoteService returns a boolean if a field has been set.
func (o *ControllerApiDescriptionModel) HasIsRemoteService() bool {
	if o != nil && !isNil(o.IsRemoteService) {
		return true
	}

	return false
}

// SetIsRemoteService gets a reference to the given bool and assigns it to the IsRemoteService field.
func (o *ControllerApiDescriptionModel) SetIsRemoteService(v bool) {
	o.IsRemoteService = &v
}

// GetIsIntegrationService returns the IsIntegrationService field value if set, zero value otherwise.
func (o *ControllerApiDescriptionModel) GetIsIntegrationService() bool {
	if o == nil || isNil(o.IsIntegrationService) {
		var ret bool
		return ret
	}
	return *o.IsIntegrationService
}

// GetIsIntegrationServiceOk returns a tuple with the IsIntegrationService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerApiDescriptionModel) GetIsIntegrationServiceOk() (*bool, bool) {
	if o == nil || isNil(o.IsIntegrationService) {
    return nil, false
	}
	return o.IsIntegrationService, true
}

// HasIsIntegrationService returns a boolean if a field has been set.
func (o *ControllerApiDescriptionModel) HasIsIntegrationService() bool {
	if o != nil && !isNil(o.IsIntegrationService) {
		return true
	}

	return false
}

// SetIsIntegrationService gets a reference to the given bool and assigns it to the IsIntegrationService field.
func (o *ControllerApiDescriptionModel) SetIsIntegrationService(v bool) {
	o.IsIntegrationService = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ControllerApiDescriptionModel) GetApiVersion() string {
	if o == nil || isNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerApiDescriptionModel) GetApiVersionOk() (*string, bool) {
	if o == nil || isNil(o.ApiVersion) {
    return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ControllerApiDescriptionModel) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ControllerApiDescriptionModel) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ControllerApiDescriptionModel) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerApiDescriptionModel) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ControllerApiDescriptionModel) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ControllerApiDescriptionModel) SetType(v string) {
	o.Type = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *ControllerApiDescriptionModel) GetInterfaces() []ControllerInterfaceApiDescriptionModel {
	if o == nil || isNil(o.Interfaces) {
		var ret []ControllerInterfaceApiDescriptionModel
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerApiDescriptionModel) GetInterfacesOk() ([]ControllerInterfaceApiDescriptionModel, bool) {
	if o == nil || isNil(o.Interfaces) {
    return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *ControllerApiDescriptionModel) HasInterfaces() bool {
	if o != nil && !isNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []ControllerInterfaceApiDescriptionModel and assigns it to the Interfaces field.
func (o *ControllerApiDescriptionModel) SetInterfaces(v []ControllerInterfaceApiDescriptionModel) {
	o.Interfaces = v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ControllerApiDescriptionModel) GetActions() map[string]ActionApiDescriptionModel {
	if o == nil || isNil(o.Actions) {
		var ret map[string]ActionApiDescriptionModel
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerApiDescriptionModel) GetActionsOk() (*map[string]ActionApiDescriptionModel, bool) {
	if o == nil || isNil(o.Actions) {
    return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ControllerApiDescriptionModel) HasActions() bool {
	if o != nil && !isNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given map[string]ActionApiDescriptionModel and assigns it to the Actions field.
func (o *ControllerApiDescriptionModel) SetActions(v map[string]ActionApiDescriptionModel) {
	o.Actions = &v
}

func (o ControllerApiDescriptionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ControllerName) {
		toSerialize["controllerName"] = o.ControllerName
	}
	if !isNil(o.ControllerGroupName) {
		toSerialize["controllerGroupName"] = o.ControllerGroupName
	}
	if !isNil(o.IsRemoteService) {
		toSerialize["isRemoteService"] = o.IsRemoteService
	}
	if !isNil(o.IsIntegrationService) {
		toSerialize["isIntegrationService"] = o.IsIntegrationService
	}
	if !isNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	if !isNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableControllerApiDescriptionModel struct {
	value *ControllerApiDescriptionModel
	isSet bool
}

func (v NullableControllerApiDescriptionModel) Get() *ControllerApiDescriptionModel {
	return v.value
}

func (v *NullableControllerApiDescriptionModel) Set(val *ControllerApiDescriptionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerApiDescriptionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerApiDescriptionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerApiDescriptionModel(val *ControllerApiDescriptionModel) *NullableControllerApiDescriptionModel {
	return &NullableControllerApiDescriptionModel{value: val, isSet: true}
}

func (v NullableControllerApiDescriptionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerApiDescriptionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



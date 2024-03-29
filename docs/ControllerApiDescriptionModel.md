# ControllerApiDescriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerName** | Pointer to **string** |  | [optional] 
**ControllerGroupName** | Pointer to **string** |  | [optional] 
**IsRemoteService** | Pointer to **bool** |  | [optional] 
**IsIntegrationService** | Pointer to **bool** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Interfaces** | Pointer to [**[]ControllerInterfaceApiDescriptionModel**](ControllerInterfaceApiDescriptionModel.md) |  | [optional] 
**Actions** | Pointer to [**map[string]ActionApiDescriptionModel**](ActionApiDescriptionModel.md) |  | [optional] 

## Methods

### NewControllerApiDescriptionModel

`func NewControllerApiDescriptionModel() *ControllerApiDescriptionModel`

NewControllerApiDescriptionModel instantiates a new ControllerApiDescriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerApiDescriptionModelWithDefaults

`func NewControllerApiDescriptionModelWithDefaults() *ControllerApiDescriptionModel`

NewControllerApiDescriptionModelWithDefaults instantiates a new ControllerApiDescriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerName

`func (o *ControllerApiDescriptionModel) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *ControllerApiDescriptionModel) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *ControllerApiDescriptionModel) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *ControllerApiDescriptionModel) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetControllerGroupName

`func (o *ControllerApiDescriptionModel) GetControllerGroupName() string`

GetControllerGroupName returns the ControllerGroupName field if non-nil, zero value otherwise.

### GetControllerGroupNameOk

`func (o *ControllerApiDescriptionModel) GetControllerGroupNameOk() (*string, bool)`

GetControllerGroupNameOk returns a tuple with the ControllerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerGroupName

`func (o *ControllerApiDescriptionModel) SetControllerGroupName(v string)`

SetControllerGroupName sets ControllerGroupName field to given value.

### HasControllerGroupName

`func (o *ControllerApiDescriptionModel) HasControllerGroupName() bool`

HasControllerGroupName returns a boolean if a field has been set.

### GetIsRemoteService

`func (o *ControllerApiDescriptionModel) GetIsRemoteService() bool`

GetIsRemoteService returns the IsRemoteService field if non-nil, zero value otherwise.

### GetIsRemoteServiceOk

`func (o *ControllerApiDescriptionModel) GetIsRemoteServiceOk() (*bool, bool)`

GetIsRemoteServiceOk returns a tuple with the IsRemoteService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteService

`func (o *ControllerApiDescriptionModel) SetIsRemoteService(v bool)`

SetIsRemoteService sets IsRemoteService field to given value.

### HasIsRemoteService

`func (o *ControllerApiDescriptionModel) HasIsRemoteService() bool`

HasIsRemoteService returns a boolean if a field has been set.

### GetIsIntegrationService

`func (o *ControllerApiDescriptionModel) GetIsIntegrationService() bool`

GetIsIntegrationService returns the IsIntegrationService field if non-nil, zero value otherwise.

### GetIsIntegrationServiceOk

`func (o *ControllerApiDescriptionModel) GetIsIntegrationServiceOk() (*bool, bool)`

GetIsIntegrationServiceOk returns a tuple with the IsIntegrationService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIntegrationService

`func (o *ControllerApiDescriptionModel) SetIsIntegrationService(v bool)`

SetIsIntegrationService sets IsIntegrationService field to given value.

### HasIsIntegrationService

`func (o *ControllerApiDescriptionModel) HasIsIntegrationService() bool`

HasIsIntegrationService returns a boolean if a field has been set.

### GetApiVersion

`func (o *ControllerApiDescriptionModel) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ControllerApiDescriptionModel) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ControllerApiDescriptionModel) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ControllerApiDescriptionModel) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetType

`func (o *ControllerApiDescriptionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ControllerApiDescriptionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ControllerApiDescriptionModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ControllerApiDescriptionModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInterfaces

`func (o *ControllerApiDescriptionModel) GetInterfaces() []ControllerInterfaceApiDescriptionModel`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ControllerApiDescriptionModel) GetInterfacesOk() (*[]ControllerInterfaceApiDescriptionModel, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ControllerApiDescriptionModel) SetInterfaces(v []ControllerInterfaceApiDescriptionModel)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ControllerApiDescriptionModel) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetActions

`func (o *ControllerApiDescriptionModel) GetActions() map[string]ActionApiDescriptionModel`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ControllerApiDescriptionModel) GetActionsOk() (*map[string]ActionApiDescriptionModel, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ControllerApiDescriptionModel) SetActions(v map[string]ActionApiDescriptionModel)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ControllerApiDescriptionModel) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



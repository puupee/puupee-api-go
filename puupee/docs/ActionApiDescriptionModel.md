# ActionApiDescriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**HttpMethod** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SupportedVersions** | Pointer to **[]string** |  | [optional] 
**ParametersOnMethod** | Pointer to [**[]MethodParameterApiDescriptionModel**](MethodParameterApiDescriptionModel.md) |  | [optional] 
**Parameters** | Pointer to [**[]ParameterApiDescriptionModel**](ParameterApiDescriptionModel.md) |  | [optional] 
**ReturnValue** | Pointer to [**ReturnValueApiDescriptionModel**](ReturnValueApiDescriptionModel.md) |  | [optional] 
**AllowAnonymous** | Pointer to **bool** |  | [optional] 
**ImplementFrom** | Pointer to **string** |  | [optional] 

## Methods

### NewActionApiDescriptionModel

`func NewActionApiDescriptionModel() *ActionApiDescriptionModel`

NewActionApiDescriptionModel instantiates a new ActionApiDescriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionApiDescriptionModelWithDefaults

`func NewActionApiDescriptionModelWithDefaults() *ActionApiDescriptionModel`

NewActionApiDescriptionModelWithDefaults instantiates a new ActionApiDescriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueName

`func (o *ActionApiDescriptionModel) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ActionApiDescriptionModel) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ActionApiDescriptionModel) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ActionApiDescriptionModel) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetName

`func (o *ActionApiDescriptionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionApiDescriptionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionApiDescriptionModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionApiDescriptionModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHttpMethod

`func (o *ActionApiDescriptionModel) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *ActionApiDescriptionModel) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *ActionApiDescriptionModel) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *ActionApiDescriptionModel) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetUrl

`func (o *ActionApiDescriptionModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ActionApiDescriptionModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ActionApiDescriptionModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ActionApiDescriptionModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSupportedVersions

`func (o *ActionApiDescriptionModel) GetSupportedVersions() []string`

GetSupportedVersions returns the SupportedVersions field if non-nil, zero value otherwise.

### GetSupportedVersionsOk

`func (o *ActionApiDescriptionModel) GetSupportedVersionsOk() (*[]string, bool)`

GetSupportedVersionsOk returns a tuple with the SupportedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVersions

`func (o *ActionApiDescriptionModel) SetSupportedVersions(v []string)`

SetSupportedVersions sets SupportedVersions field to given value.

### HasSupportedVersions

`func (o *ActionApiDescriptionModel) HasSupportedVersions() bool`

HasSupportedVersions returns a boolean if a field has been set.

### GetParametersOnMethod

`func (o *ActionApiDescriptionModel) GetParametersOnMethod() []MethodParameterApiDescriptionModel`

GetParametersOnMethod returns the ParametersOnMethod field if non-nil, zero value otherwise.

### GetParametersOnMethodOk

`func (o *ActionApiDescriptionModel) GetParametersOnMethodOk() (*[]MethodParameterApiDescriptionModel, bool)`

GetParametersOnMethodOk returns a tuple with the ParametersOnMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametersOnMethod

`func (o *ActionApiDescriptionModel) SetParametersOnMethod(v []MethodParameterApiDescriptionModel)`

SetParametersOnMethod sets ParametersOnMethod field to given value.

### HasParametersOnMethod

`func (o *ActionApiDescriptionModel) HasParametersOnMethod() bool`

HasParametersOnMethod returns a boolean if a field has been set.

### GetParameters

`func (o *ActionApiDescriptionModel) GetParameters() []ParameterApiDescriptionModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ActionApiDescriptionModel) GetParametersOk() (*[]ParameterApiDescriptionModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ActionApiDescriptionModel) SetParameters(v []ParameterApiDescriptionModel)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ActionApiDescriptionModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetReturnValue

`func (o *ActionApiDescriptionModel) GetReturnValue() ReturnValueApiDescriptionModel`

GetReturnValue returns the ReturnValue field if non-nil, zero value otherwise.

### GetReturnValueOk

`func (o *ActionApiDescriptionModel) GetReturnValueOk() (*ReturnValueApiDescriptionModel, bool)`

GetReturnValueOk returns a tuple with the ReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnValue

`func (o *ActionApiDescriptionModel) SetReturnValue(v ReturnValueApiDescriptionModel)`

SetReturnValue sets ReturnValue field to given value.

### HasReturnValue

`func (o *ActionApiDescriptionModel) HasReturnValue() bool`

HasReturnValue returns a boolean if a field has been set.

### GetAllowAnonymous

`func (o *ActionApiDescriptionModel) GetAllowAnonymous() bool`

GetAllowAnonymous returns the AllowAnonymous field if non-nil, zero value otherwise.

### GetAllowAnonymousOk

`func (o *ActionApiDescriptionModel) GetAllowAnonymousOk() (*bool, bool)`

GetAllowAnonymousOk returns a tuple with the AllowAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymous

`func (o *ActionApiDescriptionModel) SetAllowAnonymous(v bool)`

SetAllowAnonymous sets AllowAnonymous field to given value.

### HasAllowAnonymous

`func (o *ActionApiDescriptionModel) HasAllowAnonymous() bool`

HasAllowAnonymous returns a boolean if a field has been set.

### GetImplementFrom

`func (o *ActionApiDescriptionModel) GetImplementFrom() string`

GetImplementFrom returns the ImplementFrom field if non-nil, zero value otherwise.

### GetImplementFromOk

`func (o *ActionApiDescriptionModel) GetImplementFromOk() (*string, bool)`

GetImplementFromOk returns a tuple with the ImplementFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementFrom

`func (o *ActionApiDescriptionModel) SetImplementFrom(v string)`

SetImplementFrom sets ImplementFrom field to given value.

### HasImplementFrom

`func (o *ActionApiDescriptionModel) HasImplementFrom() bool`

HasImplementFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



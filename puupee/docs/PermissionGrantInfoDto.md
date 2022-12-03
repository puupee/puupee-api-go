# PermissionGrantInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ParentName** | Pointer to **string** |  | [optional] 
**IsGranted** | Pointer to **bool** |  | [optional] 
**AllowedProviders** | Pointer to **[]string** |  | [optional] 
**GrantedProviders** | Pointer to [**[]ProviderInfoDto**](ProviderInfoDto.md) |  | [optional] 

## Methods

### NewPermissionGrantInfoDto

`func NewPermissionGrantInfoDto() *PermissionGrantInfoDto`

NewPermissionGrantInfoDto instantiates a new PermissionGrantInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGrantInfoDtoWithDefaults

`func NewPermissionGrantInfoDtoWithDefaults() *PermissionGrantInfoDto`

NewPermissionGrantInfoDtoWithDefaults instantiates a new PermissionGrantInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PermissionGrantInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionGrantInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionGrantInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionGrantInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *PermissionGrantInfoDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PermissionGrantInfoDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PermissionGrantInfoDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PermissionGrantInfoDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetParentName

`func (o *PermissionGrantInfoDto) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *PermissionGrantInfoDto) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *PermissionGrantInfoDto) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *PermissionGrantInfoDto) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### GetIsGranted

`func (o *PermissionGrantInfoDto) GetIsGranted() bool`

GetIsGranted returns the IsGranted field if non-nil, zero value otherwise.

### GetIsGrantedOk

`func (o *PermissionGrantInfoDto) GetIsGrantedOk() (*bool, bool)`

GetIsGrantedOk returns a tuple with the IsGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGranted

`func (o *PermissionGrantInfoDto) SetIsGranted(v bool)`

SetIsGranted sets IsGranted field to given value.

### HasIsGranted

`func (o *PermissionGrantInfoDto) HasIsGranted() bool`

HasIsGranted returns a boolean if a field has been set.

### GetAllowedProviders

`func (o *PermissionGrantInfoDto) GetAllowedProviders() []string`

GetAllowedProviders returns the AllowedProviders field if non-nil, zero value otherwise.

### GetAllowedProvidersOk

`func (o *PermissionGrantInfoDto) GetAllowedProvidersOk() (*[]string, bool)`

GetAllowedProvidersOk returns a tuple with the AllowedProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedProviders

`func (o *PermissionGrantInfoDto) SetAllowedProviders(v []string)`

SetAllowedProviders sets AllowedProviders field to given value.

### HasAllowedProviders

`func (o *PermissionGrantInfoDto) HasAllowedProviders() bool`

HasAllowedProviders returns a boolean if a field has been set.

### GetGrantedProviders

`func (o *PermissionGrantInfoDto) GetGrantedProviders() []ProviderInfoDto`

GetGrantedProviders returns the GrantedProviders field if non-nil, zero value otherwise.

### GetGrantedProvidersOk

`func (o *PermissionGrantInfoDto) GetGrantedProvidersOk() (*[]ProviderInfoDto, bool)`

GetGrantedProvidersOk returns a tuple with the GrantedProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedProviders

`func (o *PermissionGrantInfoDto) SetGrantedProviders(v []ProviderInfoDto)`

SetGrantedProviders sets GrantedProviders field to given value.

### HasGrantedProviders

`func (o *PermissionGrantInfoDto) HasGrantedProviders() bool`

HasGrantedProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



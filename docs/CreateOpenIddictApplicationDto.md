# CreateOpenIddictApplicationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DisplayNames** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **string** |  | [optional] 
**RedirectUris** | Pointer to **string** |  | [optional] 
**Requirements** | Pointer to **string** |  | [optional] 
**ClientUri** | Pointer to **string** |  | [optional] 
**LogoUri** | Pointer to **string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateOpenIddictApplicationDto

`func NewCreateOpenIddictApplicationDto() *CreateOpenIddictApplicationDto`

NewCreateOpenIddictApplicationDto instantiates a new CreateOpenIddictApplicationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOpenIddictApplicationDtoWithDefaults

`func NewCreateOpenIddictApplicationDtoWithDefaults() *CreateOpenIddictApplicationDto`

NewCreateOpenIddictApplicationDtoWithDefaults instantiates a new CreateOpenIddictApplicationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOpenIddictApplicationDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOpenIddictApplicationDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOpenIddictApplicationDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateOpenIddictApplicationDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateOpenIddictApplicationDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateOpenIddictApplicationDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateOpenIddictApplicationDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateOpenIddictApplicationDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayNames

`func (o *CreateOpenIddictApplicationDto) GetDisplayNames() string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *CreateOpenIddictApplicationDto) GetDisplayNamesOk() (*string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *CreateOpenIddictApplicationDto) SetDisplayNames(v string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *CreateOpenIddictApplicationDto) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateOpenIddictApplicationDto) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateOpenIddictApplicationDto) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateOpenIddictApplicationDto) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateOpenIddictApplicationDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPostLogoutRedirectUris

`func (o *CreateOpenIddictApplicationDto) GetPostLogoutRedirectUris() string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *CreateOpenIddictApplicationDto) GetPostLogoutRedirectUrisOk() (*string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *CreateOpenIddictApplicationDto) SetPostLogoutRedirectUris(v string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *CreateOpenIddictApplicationDto) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetProperties

`func (o *CreateOpenIddictApplicationDto) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateOpenIddictApplicationDto) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateOpenIddictApplicationDto) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateOpenIddictApplicationDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRedirectUris

`func (o *CreateOpenIddictApplicationDto) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *CreateOpenIddictApplicationDto) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *CreateOpenIddictApplicationDto) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *CreateOpenIddictApplicationDto) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRequirements

`func (o *CreateOpenIddictApplicationDto) GetRequirements() string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *CreateOpenIddictApplicationDto) GetRequirementsOk() (*string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *CreateOpenIddictApplicationDto) SetRequirements(v string)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *CreateOpenIddictApplicationDto) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetClientUri

`func (o *CreateOpenIddictApplicationDto) GetClientUri() string`

GetClientUri returns the ClientUri field if non-nil, zero value otherwise.

### GetClientUriOk

`func (o *CreateOpenIddictApplicationDto) GetClientUriOk() (*string, bool)`

GetClientUriOk returns a tuple with the ClientUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUri

`func (o *CreateOpenIddictApplicationDto) SetClientUri(v string)`

SetClientUri sets ClientUri field to given value.

### HasClientUri

`func (o *CreateOpenIddictApplicationDto) HasClientUri() bool`

HasClientUri returns a boolean if a field has been set.

### GetLogoUri

`func (o *CreateOpenIddictApplicationDto) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *CreateOpenIddictApplicationDto) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *CreateOpenIddictApplicationDto) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *CreateOpenIddictApplicationDto) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### GetGrantTypes

`func (o *CreateOpenIddictApplicationDto) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *CreateOpenIddictApplicationDto) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *CreateOpenIddictApplicationDto) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *CreateOpenIddictApplicationDto) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetScopes

`func (o *CreateOpenIddictApplicationDto) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateOpenIddictApplicationDto) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateOpenIddictApplicationDto) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateOpenIddictApplicationDto) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



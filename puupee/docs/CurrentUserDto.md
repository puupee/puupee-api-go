# CurrentUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAuthenticated** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**ImpersonatorUserId** | Pointer to **string** |  | [optional] 
**ImpersonatorTenantId** | Pointer to **string** |  | [optional] 
**ImpersonatorUserName** | Pointer to **string** |  | [optional] 
**ImpersonatorTenantName** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SurName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailVerified** | Pointer to **bool** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PhoneNumberVerified** | Pointer to **bool** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCurrentUserDto

`func NewCurrentUserDto() *CurrentUserDto`

NewCurrentUserDto instantiates a new CurrentUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserDtoWithDefaults

`func NewCurrentUserDtoWithDefaults() *CurrentUserDto`

NewCurrentUserDtoWithDefaults instantiates a new CurrentUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAuthenticated

`func (o *CurrentUserDto) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *CurrentUserDto) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *CurrentUserDto) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.

### HasIsAuthenticated

`func (o *CurrentUserDto) HasIsAuthenticated() bool`

HasIsAuthenticated returns a boolean if a field has been set.

### GetId

`func (o *CurrentUserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentUserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentUserDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CurrentUserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantId

`func (o *CurrentUserDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CurrentUserDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CurrentUserDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CurrentUserDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetImpersonatorUserId

`func (o *CurrentUserDto) GetImpersonatorUserId() string`

GetImpersonatorUserId returns the ImpersonatorUserId field if non-nil, zero value otherwise.

### GetImpersonatorUserIdOk

`func (o *CurrentUserDto) GetImpersonatorUserIdOk() (*string, bool)`

GetImpersonatorUserIdOk returns a tuple with the ImpersonatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonatorUserId

`func (o *CurrentUserDto) SetImpersonatorUserId(v string)`

SetImpersonatorUserId sets ImpersonatorUserId field to given value.

### HasImpersonatorUserId

`func (o *CurrentUserDto) HasImpersonatorUserId() bool`

HasImpersonatorUserId returns a boolean if a field has been set.

### GetImpersonatorTenantId

`func (o *CurrentUserDto) GetImpersonatorTenantId() string`

GetImpersonatorTenantId returns the ImpersonatorTenantId field if non-nil, zero value otherwise.

### GetImpersonatorTenantIdOk

`func (o *CurrentUserDto) GetImpersonatorTenantIdOk() (*string, bool)`

GetImpersonatorTenantIdOk returns a tuple with the ImpersonatorTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonatorTenantId

`func (o *CurrentUserDto) SetImpersonatorTenantId(v string)`

SetImpersonatorTenantId sets ImpersonatorTenantId field to given value.

### HasImpersonatorTenantId

`func (o *CurrentUserDto) HasImpersonatorTenantId() bool`

HasImpersonatorTenantId returns a boolean if a field has been set.

### GetImpersonatorUserName

`func (o *CurrentUserDto) GetImpersonatorUserName() string`

GetImpersonatorUserName returns the ImpersonatorUserName field if non-nil, zero value otherwise.

### GetImpersonatorUserNameOk

`func (o *CurrentUserDto) GetImpersonatorUserNameOk() (*string, bool)`

GetImpersonatorUserNameOk returns a tuple with the ImpersonatorUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonatorUserName

`func (o *CurrentUserDto) SetImpersonatorUserName(v string)`

SetImpersonatorUserName sets ImpersonatorUserName field to given value.

### HasImpersonatorUserName

`func (o *CurrentUserDto) HasImpersonatorUserName() bool`

HasImpersonatorUserName returns a boolean if a field has been set.

### GetImpersonatorTenantName

`func (o *CurrentUserDto) GetImpersonatorTenantName() string`

GetImpersonatorTenantName returns the ImpersonatorTenantName field if non-nil, zero value otherwise.

### GetImpersonatorTenantNameOk

`func (o *CurrentUserDto) GetImpersonatorTenantNameOk() (*string, bool)`

GetImpersonatorTenantNameOk returns a tuple with the ImpersonatorTenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonatorTenantName

`func (o *CurrentUserDto) SetImpersonatorTenantName(v string)`

SetImpersonatorTenantName sets ImpersonatorTenantName field to given value.

### HasImpersonatorTenantName

`func (o *CurrentUserDto) HasImpersonatorTenantName() bool`

HasImpersonatorTenantName returns a boolean if a field has been set.

### GetUserName

`func (o *CurrentUserDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CurrentUserDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CurrentUserDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *CurrentUserDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetName

`func (o *CurrentUserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentUserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentUserDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrentUserDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurName

`func (o *CurrentUserDto) GetSurName() string`

GetSurName returns the SurName field if non-nil, zero value otherwise.

### GetSurNameOk

`func (o *CurrentUserDto) GetSurNameOk() (*string, bool)`

GetSurNameOk returns a tuple with the SurName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurName

`func (o *CurrentUserDto) SetSurName(v string)`

SetSurName sets SurName field to given value.

### HasSurName

`func (o *CurrentUserDto) HasSurName() bool`

HasSurName returns a boolean if a field has been set.

### GetEmail

`func (o *CurrentUserDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CurrentUserDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CurrentUserDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CurrentUserDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *CurrentUserDto) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *CurrentUserDto) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *CurrentUserDto) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *CurrentUserDto) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CurrentUserDto) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CurrentUserDto) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CurrentUserDto) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CurrentUserDto) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumberVerified

`func (o *CurrentUserDto) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *CurrentUserDto) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *CurrentUserDto) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *CurrentUserDto) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### GetRoles

`func (o *CurrentUserDto) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CurrentUserDto) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CurrentUserDto) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CurrentUserDto) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



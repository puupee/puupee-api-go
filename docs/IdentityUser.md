# IdentityUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ExtraProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**ConcurrencyStamp** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 
**LastModifierId** | Pointer to **string** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeleterId** | Pointer to **string** |  | [optional] 
**DeletionTime** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**NormalizedUserName** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**NormalizedEmail** | Pointer to **string** |  | [optional] [readonly] 
**EmailConfirmed** | Pointer to **bool** |  | [optional] [readonly] 
**PasswordHash** | Pointer to **string** |  | [optional] [readonly] 
**SecurityStamp** | Pointer to **string** |  | [optional] [readonly] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] [readonly] 
**PhoneNumberConfirmed** | Pointer to **bool** |  | [optional] [readonly] 
**IsActive** | Pointer to **bool** |  | [optional] [readonly] 
**TwoFactorEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**LockoutEnd** | Pointer to **time.Time** |  | [optional] [readonly] 
**LockoutEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**AccessFailedCount** | Pointer to **int32** |  | [optional] [readonly] 
**ShouldChangePasswordOnNextLogin** | Pointer to **bool** |  | [optional] [readonly] 
**EntityVersion** | Pointer to **int32** |  | [optional] [readonly] 
**LastPasswordChangeTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Roles** | Pointer to [**[]IdentityUserRole**](IdentityUserRole.md) |  | [optional] [readonly] 
**Claims** | Pointer to [**[]IdentityUserClaim**](IdentityUserClaim.md) |  | [optional] [readonly] 
**Logins** | Pointer to [**[]IdentityUserLogin**](IdentityUserLogin.md) |  | [optional] [readonly] 
**Tokens** | Pointer to [**[]IdentityUserToken**](IdentityUserToken.md) |  | [optional] [readonly] 
**OrganizationUnits** | Pointer to [**[]IdentityUserOrganizationUnit**](IdentityUserOrganizationUnit.md) |  | [optional] [readonly] 

## Methods

### NewIdentityUser

`func NewIdentityUser() *IdentityUser`

NewIdentityUser instantiates a new IdentityUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserWithDefaults

`func NewIdentityUserWithDefaults() *IdentityUser`

NewIdentityUserWithDefaults instantiates a new IdentityUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtraProperties

`func (o *IdentityUser) GetExtraProperties() map[string]map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *IdentityUser) GetExtraPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *IdentityUser) SetExtraProperties(v map[string]map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *IdentityUser) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### GetConcurrencyStamp

`func (o *IdentityUser) GetConcurrencyStamp() string`

GetConcurrencyStamp returns the ConcurrencyStamp field if non-nil, zero value otherwise.

### GetConcurrencyStampOk

`func (o *IdentityUser) GetConcurrencyStampOk() (*string, bool)`

GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyStamp

`func (o *IdentityUser) SetConcurrencyStamp(v string)`

SetConcurrencyStamp sets ConcurrencyStamp field to given value.

### HasConcurrencyStamp

`func (o *IdentityUser) HasConcurrencyStamp() bool`

HasConcurrencyStamp returns a boolean if a field has been set.

### GetCreationTime

`func (o *IdentityUser) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *IdentityUser) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *IdentityUser) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *IdentityUser) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *IdentityUser) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *IdentityUser) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *IdentityUser) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *IdentityUser) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *IdentityUser) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *IdentityUser) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *IdentityUser) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *IdentityUser) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *IdentityUser) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *IdentityUser) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *IdentityUser) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *IdentityUser) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *IdentityUser) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *IdentityUser) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *IdentityUser) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *IdentityUser) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *IdentityUser) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *IdentityUser) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *IdentityUser) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *IdentityUser) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *IdentityUser) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *IdentityUser) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *IdentityUser) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *IdentityUser) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetTenantId

`func (o *IdentityUser) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IdentityUser) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IdentityUser) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IdentityUser) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUserName

`func (o *IdentityUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IdentityUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IdentityUser) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *IdentityUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetNormalizedUserName

`func (o *IdentityUser) GetNormalizedUserName() string`

GetNormalizedUserName returns the NormalizedUserName field if non-nil, zero value otherwise.

### GetNormalizedUserNameOk

`func (o *IdentityUser) GetNormalizedUserNameOk() (*string, bool)`

GetNormalizedUserNameOk returns a tuple with the NormalizedUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedUserName

`func (o *IdentityUser) SetNormalizedUserName(v string)`

SetNormalizedUserName sets NormalizedUserName field to given value.

### HasNormalizedUserName

`func (o *IdentityUser) HasNormalizedUserName() bool`

HasNormalizedUserName returns a boolean if a field has been set.

### GetName

`func (o *IdentityUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *IdentityUser) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *IdentityUser) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *IdentityUser) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *IdentityUser) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetEmail

`func (o *IdentityUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdentityUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNormalizedEmail

`func (o *IdentityUser) GetNormalizedEmail() string`

GetNormalizedEmail returns the NormalizedEmail field if non-nil, zero value otherwise.

### GetNormalizedEmailOk

`func (o *IdentityUser) GetNormalizedEmailOk() (*string, bool)`

GetNormalizedEmailOk returns a tuple with the NormalizedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedEmail

`func (o *IdentityUser) SetNormalizedEmail(v string)`

SetNormalizedEmail sets NormalizedEmail field to given value.

### HasNormalizedEmail

`func (o *IdentityUser) HasNormalizedEmail() bool`

HasNormalizedEmail returns a boolean if a field has been set.

### GetEmailConfirmed

`func (o *IdentityUser) GetEmailConfirmed() bool`

GetEmailConfirmed returns the EmailConfirmed field if non-nil, zero value otherwise.

### GetEmailConfirmedOk

`func (o *IdentityUser) GetEmailConfirmedOk() (*bool, bool)`

GetEmailConfirmedOk returns a tuple with the EmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmed

`func (o *IdentityUser) SetEmailConfirmed(v bool)`

SetEmailConfirmed sets EmailConfirmed field to given value.

### HasEmailConfirmed

`func (o *IdentityUser) HasEmailConfirmed() bool`

HasEmailConfirmed returns a boolean if a field has been set.

### GetPasswordHash

`func (o *IdentityUser) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *IdentityUser) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *IdentityUser) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *IdentityUser) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetSecurityStamp

`func (o *IdentityUser) GetSecurityStamp() string`

GetSecurityStamp returns the SecurityStamp field if non-nil, zero value otherwise.

### GetSecurityStampOk

`func (o *IdentityUser) GetSecurityStampOk() (*string, bool)`

GetSecurityStampOk returns a tuple with the SecurityStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStamp

`func (o *IdentityUser) SetSecurityStamp(v string)`

SetSecurityStamp sets SecurityStamp field to given value.

### HasSecurityStamp

`func (o *IdentityUser) HasSecurityStamp() bool`

HasSecurityStamp returns a boolean if a field has been set.

### GetIsExternal

`func (o *IdentityUser) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *IdentityUser) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *IdentityUser) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *IdentityUser) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *IdentityUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IdentityUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IdentityUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IdentityUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumberConfirmed

`func (o *IdentityUser) GetPhoneNumberConfirmed() bool`

GetPhoneNumberConfirmed returns the PhoneNumberConfirmed field if non-nil, zero value otherwise.

### GetPhoneNumberConfirmedOk

`func (o *IdentityUser) GetPhoneNumberConfirmedOk() (*bool, bool)`

GetPhoneNumberConfirmedOk returns a tuple with the PhoneNumberConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberConfirmed

`func (o *IdentityUser) SetPhoneNumberConfirmed(v bool)`

SetPhoneNumberConfirmed sets PhoneNumberConfirmed field to given value.

### HasPhoneNumberConfirmed

`func (o *IdentityUser) HasPhoneNumberConfirmed() bool`

HasPhoneNumberConfirmed returns a boolean if a field has been set.

### GetIsActive

`func (o *IdentityUser) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *IdentityUser) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *IdentityUser) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *IdentityUser) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetTwoFactorEnabled

`func (o *IdentityUser) GetTwoFactorEnabled() bool`

GetTwoFactorEnabled returns the TwoFactorEnabled field if non-nil, zero value otherwise.

### GetTwoFactorEnabledOk

`func (o *IdentityUser) GetTwoFactorEnabledOk() (*bool, bool)`

GetTwoFactorEnabledOk returns a tuple with the TwoFactorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorEnabled

`func (o *IdentityUser) SetTwoFactorEnabled(v bool)`

SetTwoFactorEnabled sets TwoFactorEnabled field to given value.

### HasTwoFactorEnabled

`func (o *IdentityUser) HasTwoFactorEnabled() bool`

HasTwoFactorEnabled returns a boolean if a field has been set.

### GetLockoutEnd

`func (o *IdentityUser) GetLockoutEnd() time.Time`

GetLockoutEnd returns the LockoutEnd field if non-nil, zero value otherwise.

### GetLockoutEndOk

`func (o *IdentityUser) GetLockoutEndOk() (*time.Time, bool)`

GetLockoutEndOk returns a tuple with the LockoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnd

`func (o *IdentityUser) SetLockoutEnd(v time.Time)`

SetLockoutEnd sets LockoutEnd field to given value.

### HasLockoutEnd

`func (o *IdentityUser) HasLockoutEnd() bool`

HasLockoutEnd returns a boolean if a field has been set.

### GetLockoutEnabled

`func (o *IdentityUser) GetLockoutEnabled() bool`

GetLockoutEnabled returns the LockoutEnabled field if non-nil, zero value otherwise.

### GetLockoutEnabledOk

`func (o *IdentityUser) GetLockoutEnabledOk() (*bool, bool)`

GetLockoutEnabledOk returns a tuple with the LockoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnabled

`func (o *IdentityUser) SetLockoutEnabled(v bool)`

SetLockoutEnabled sets LockoutEnabled field to given value.

### HasLockoutEnabled

`func (o *IdentityUser) HasLockoutEnabled() bool`

HasLockoutEnabled returns a boolean if a field has been set.

### GetAccessFailedCount

`func (o *IdentityUser) GetAccessFailedCount() int32`

GetAccessFailedCount returns the AccessFailedCount field if non-nil, zero value otherwise.

### GetAccessFailedCountOk

`func (o *IdentityUser) GetAccessFailedCountOk() (*int32, bool)`

GetAccessFailedCountOk returns a tuple with the AccessFailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessFailedCount

`func (o *IdentityUser) SetAccessFailedCount(v int32)`

SetAccessFailedCount sets AccessFailedCount field to given value.

### HasAccessFailedCount

`func (o *IdentityUser) HasAccessFailedCount() bool`

HasAccessFailedCount returns a boolean if a field has been set.

### GetShouldChangePasswordOnNextLogin

`func (o *IdentityUser) GetShouldChangePasswordOnNextLogin() bool`

GetShouldChangePasswordOnNextLogin returns the ShouldChangePasswordOnNextLogin field if non-nil, zero value otherwise.

### GetShouldChangePasswordOnNextLoginOk

`func (o *IdentityUser) GetShouldChangePasswordOnNextLoginOk() (*bool, bool)`

GetShouldChangePasswordOnNextLoginOk returns a tuple with the ShouldChangePasswordOnNextLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldChangePasswordOnNextLogin

`func (o *IdentityUser) SetShouldChangePasswordOnNextLogin(v bool)`

SetShouldChangePasswordOnNextLogin sets ShouldChangePasswordOnNextLogin field to given value.

### HasShouldChangePasswordOnNextLogin

`func (o *IdentityUser) HasShouldChangePasswordOnNextLogin() bool`

HasShouldChangePasswordOnNextLogin returns a boolean if a field has been set.

### GetEntityVersion

`func (o *IdentityUser) GetEntityVersion() int32`

GetEntityVersion returns the EntityVersion field if non-nil, zero value otherwise.

### GetEntityVersionOk

`func (o *IdentityUser) GetEntityVersionOk() (*int32, bool)`

GetEntityVersionOk returns a tuple with the EntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityVersion

`func (o *IdentityUser) SetEntityVersion(v int32)`

SetEntityVersion sets EntityVersion field to given value.

### HasEntityVersion

`func (o *IdentityUser) HasEntityVersion() bool`

HasEntityVersion returns a boolean if a field has been set.

### GetLastPasswordChangeTime

`func (o *IdentityUser) GetLastPasswordChangeTime() time.Time`

GetLastPasswordChangeTime returns the LastPasswordChangeTime field if non-nil, zero value otherwise.

### GetLastPasswordChangeTimeOk

`func (o *IdentityUser) GetLastPasswordChangeTimeOk() (*time.Time, bool)`

GetLastPasswordChangeTimeOk returns a tuple with the LastPasswordChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChangeTime

`func (o *IdentityUser) SetLastPasswordChangeTime(v time.Time)`

SetLastPasswordChangeTime sets LastPasswordChangeTime field to given value.

### HasLastPasswordChangeTime

`func (o *IdentityUser) HasLastPasswordChangeTime() bool`

HasLastPasswordChangeTime returns a boolean if a field has been set.

### GetRoles

`func (o *IdentityUser) GetRoles() []IdentityUserRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IdentityUser) GetRolesOk() (*[]IdentityUserRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IdentityUser) SetRoles(v []IdentityUserRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IdentityUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetClaims

`func (o *IdentityUser) GetClaims() []IdentityUserClaim`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *IdentityUser) GetClaimsOk() (*[]IdentityUserClaim, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *IdentityUser) SetClaims(v []IdentityUserClaim)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *IdentityUser) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetLogins

`func (o *IdentityUser) GetLogins() []IdentityUserLogin`

GetLogins returns the Logins field if non-nil, zero value otherwise.

### GetLoginsOk

`func (o *IdentityUser) GetLoginsOk() (*[]IdentityUserLogin, bool)`

GetLoginsOk returns a tuple with the Logins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogins

`func (o *IdentityUser) SetLogins(v []IdentityUserLogin)`

SetLogins sets Logins field to given value.

### HasLogins

`func (o *IdentityUser) HasLogins() bool`

HasLogins returns a boolean if a field has been set.

### GetTokens

`func (o *IdentityUser) GetTokens() []IdentityUserToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *IdentityUser) GetTokensOk() (*[]IdentityUserToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *IdentityUser) SetTokens(v []IdentityUserToken)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *IdentityUser) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetOrganizationUnits

`func (o *IdentityUser) GetOrganizationUnits() []IdentityUserOrganizationUnit`

GetOrganizationUnits returns the OrganizationUnits field if non-nil, zero value otherwise.

### GetOrganizationUnitsOk

`func (o *IdentityUser) GetOrganizationUnitsOk() (*[]IdentityUserOrganizationUnit, bool)`

GetOrganizationUnitsOk returns a tuple with the OrganizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnits

`func (o *IdentityUser) SetOrganizationUnits(v []IdentityUserOrganizationUnit)`

SetOrganizationUnits sets OrganizationUnits field to given value.

### HasOrganizationUnits

`func (o *IdentityUser) HasOrganizationUnits() bool`

HasOrganizationUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



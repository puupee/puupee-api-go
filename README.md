# Go API client for puupee

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0.0
- Package version: 1.0.0+4
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import puupee "github.com/puupee/puupee-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), puupee.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), puupee.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), puupee.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), puupee.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AbpApiDefinitionApi* | [**ApiAbpApiDefinitionGet**](docs/AbpApiDefinitionApi.md#apiabpapidefinitionget) | **Get** /api/abp/api-definition | 
*AbpApplicationConfigurationApi* | [**ApiAbpApplicationConfigurationGet**](docs/AbpApplicationConfigurationApi.md#apiabpapplicationconfigurationget) | **Get** /api/abp/application-configuration | 
*AbpApplicationLocalizationApi* | [**ApiAbpApplicationLocalizationGet**](docs/AbpApplicationLocalizationApi.md#apiabpapplicationlocalizationget) | **Get** /api/abp/application-localization | 
*AbpTenantApi* | [**ApiAbpMultiTenancyTenantsByIdIdGet**](docs/AbpTenantApi.md#apiabpmultitenancytenantsbyididget) | **Get** /api/abp/multi-tenancy/tenants/by-id/{id} | 
*AbpTenantApi* | [**ApiAbpMultiTenancyTenantsByNameNameGet**](docs/AbpTenantApi.md#apiabpmultitenancytenantsbynamenameget) | **Get** /api/abp/multi-tenancy/tenants/by-name/{name} | 
*AccountApi* | [**ApiAccountRegisterPost**](docs/AccountApi.md#apiaccountregisterpost) | **Post** /api/account/register | 
*AccountApi* | [**ApiAccountResetPasswordPost**](docs/AccountApi.md#apiaccountresetpasswordpost) | **Post** /api/account/reset-password | 
*AccountApi* | [**ApiAccountSendPasswordResetCodePost**](docs/AccountApi.md#apiaccountsendpasswordresetcodepost) | **Post** /api/account/send-password-reset-code | 
*AccountApi* | [**ApiAccountVerifyPasswordResetTokenPost**](docs/AccountApi.md#apiaccountverifypasswordresettokenpost) | **Post** /api/account/verify-password-reset-token | 
*AccountApi* | [**ApiAppAccountDelete**](docs/AccountApi.md#apiappaccountdelete) | **Delete** /api/app/account | 
*AppApi* | [**ApiAppAppByDeveloperAllGet**](docs/AppApi.md#apiappappbydeveloperallget) | **Get** /api/app/app/by-developer-all | 
*AppApi* | [**ApiAppAppByDeveloperGet**](docs/AppApi.md#apiappappbydeveloperget) | **Get** /api/app/app/by-developer | 
*AppApi* | [**ApiAppAppByNameGet**](docs/AppApi.md#apiappappbynameget) | **Get** /api/app/app/by-name | 
*AppApi* | [**ApiAppAppGet**](docs/AppApi.md#apiappappget) | **Get** /api/app/app | 
*AppApi* | [**ApiAppAppIdDelete**](docs/AppApi.md#apiappappiddelete) | **Delete** /api/app/app/{id} | 
*AppApi* | [**ApiAppAppIdGet**](docs/AppApi.md#apiappappidget) | **Get** /api/app/app/{id} | 
*AppApi* | [**ApiAppAppIdPut**](docs/AppApi.md#apiappappidput) | **Put** /api/app/app/{id} | 
*AppApi* | [**ApiAppAppIdWithUserGet**](docs/AppApi.md#apiappappidwithuserget) | **Get** /api/app/app/{id}/with-user | 
*AppApi* | [**ApiAppAppPost**](docs/AppApi.md#apiappapppost) | **Post** /api/app/app | 
*AppApi* | [**ApiAppAppPublicGet**](docs/AppApi.md#apiappapppublicget) | **Get** /api/app/app/public | 
*AppApi* | [**ApiAppAppUploadCredentialsGet**](docs/AppApi.md#apiappappuploadcredentialsget) | **Get** /api/app/app/upload-credentials | 
*AppFeatureApi* | [**ApiAppAppFeatureGet**](docs/AppFeatureApi.md#apiappappfeatureget) | **Get** /api/app/app-feature | 
*AppFeatureApi* | [**ApiAppAppFeatureIdDelete**](docs/AppFeatureApi.md#apiappappfeatureiddelete) | **Delete** /api/app/app-feature/{id} | 
*AppFeatureApi* | [**ApiAppAppFeatureIdPut**](docs/AppFeatureApi.md#apiappappfeatureidput) | **Put** /api/app/app-feature/{id} | 
*AppFeatureApi* | [**ApiAppAppFeaturePost**](docs/AppFeatureApi.md#apiappappfeaturepost) | **Post** /api/app/app-feature | 
*AppPricingApi* | [**ApiAppAppPricingByAppIdAppIdGet**](docs/AppPricingApi.md#apiappapppricingbyappidappidget) | **Get** /api/app/app-pricing/by-app-id/{appId} | 
*AppPricingApi* | [**ApiAppAppPricingGet**](docs/AppPricingApi.md#apiappapppricingget) | **Get** /api/app/app-pricing | 
*AppPricingApi* | [**ApiAppAppPricingIdDelete**](docs/AppPricingApi.md#apiappapppricingiddelete) | **Delete** /api/app/app-pricing/{id} | 
*AppPricingApi* | [**ApiAppAppPricingIdGet**](docs/AppPricingApi.md#apiappapppricingidget) | **Get** /api/app/app-pricing/{id} | 
*AppPricingApi* | [**ApiAppAppPricingIdPut**](docs/AppPricingApi.md#apiappapppricingidput) | **Put** /api/app/app-pricing/{id} | 
*AppPricingApi* | [**ApiAppAppPricingPost**](docs/AppPricingApi.md#apiappapppricingpost) | **Post** /api/app/app-pricing | 
*AppPricingItemApi* | [**ApiAppAppPricingItemGet**](docs/AppPricingItemApi.md#apiappapppricingitemget) | **Get** /api/app/app-pricing-item | 
*AppPricingItemApi* | [**ApiAppAppPricingItemIdDelete**](docs/AppPricingItemApi.md#apiappapppricingitemiddelete) | **Delete** /api/app/app-pricing-item/{id} | 
*AppPricingItemApi* | [**ApiAppAppPricingItemIdGet**](docs/AppPricingItemApi.md#apiappapppricingitemidget) | **Get** /api/app/app-pricing-item/{id} | 
*AppPricingItemApi* | [**ApiAppAppPricingItemIdPut**](docs/AppPricingItemApi.md#apiappapppricingitemidput) | **Put** /api/app/app-pricing-item/{id} | 
*AppPricingItemApi* | [**ApiAppAppPricingItemPost**](docs/AppPricingItemApi.md#apiappapppricingitempost) | **Post** /api/app/app-pricing-item | 
*AppReleaseApi* | [**ApiAppAppReleaseGet**](docs/AppReleaseApi.md#apiappappreleaseget) | **Get** /api/app/app-release | 
*AppReleaseApi* | [**ApiAppAppReleaseIdDelete**](docs/AppReleaseApi.md#apiappappreleaseiddelete) | **Delete** /api/app/app-release/{id} | 
*AppReleaseApi* | [**ApiAppAppReleaseIdGet**](docs/AppReleaseApi.md#apiappappreleaseidget) | **Get** /api/app/app-release/{id} | 
*AppReleaseApi* | [**ApiAppAppReleaseIdPut**](docs/AppReleaseApi.md#apiappappreleaseidput) | **Put** /api/app/app-release/{id} | 
*AppReleaseApi* | [**ApiAppAppReleaseLatestGet**](docs/AppReleaseApi.md#apiappappreleaselatestget) | **Get** /api/app/app-release/latest | 
*AppReleaseApi* | [**ApiAppAppReleasePost**](docs/AppReleaseApi.md#apiappappreleasepost) | **Post** /api/app/app-release | 
*AppSdkApi* | [**ApiAppAppSdkGet**](docs/AppSdkApi.md#apiappappsdkget) | **Get** /api/app/app-sdk | 
*AppSdkApi* | [**ApiAppAppSdkIdDelete**](docs/AppSdkApi.md#apiappappsdkiddelete) | **Delete** /api/app/app-sdk/{id} | 
*AppSdkApi* | [**ApiAppAppSdkIdPut**](docs/AppSdkApi.md#apiappappsdkidput) | **Put** /api/app/app-sdk/{id} | 
*AppSdkApi* | [**ApiAppAppSdkPost**](docs/AppSdkApi.md#apiappappsdkpost) | **Post** /api/app/app-sdk | 
*AppUserScoreApi* | [**ApiAppAppUserScorePost**](docs/AppUserScoreApi.md#apiappappuserscorepost) | **Post** /api/app/app-user-score | 
*DeviceApi* | [**ApiAppDeviceBindPost**](docs/DeviceApi.md#apiappdevicebindpost) | **Post** /api/app/device/bind | 
*DeviceApi* | [**ApiAppDeviceDelete**](docs/DeviceApi.md#apiappdevicedelete) | **Delete** /api/app/device | 
*DeviceApi* | [**ApiAppDeviceGet**](docs/DeviceApi.md#apiappdeviceget) | **Get** /api/app/device | 
*DeviceApi* | [**ApiAppDeviceRefreshPost**](docs/DeviceApi.md#apiappdevicerefreshpost) | **Post** /api/app/device/refresh | 
*EmailSettingsApi* | [**ApiSettingManagementEmailingGet**](docs/EmailSettingsApi.md#apisettingmanagementemailingget) | **Get** /api/setting-management/emailing | 
*EmailSettingsApi* | [**ApiSettingManagementEmailingPost**](docs/EmailSettingsApi.md#apisettingmanagementemailingpost) | **Post** /api/setting-management/emailing | 
*EmailSettingsApi* | [**ApiSettingManagementEmailingSendTestEmailPost**](docs/EmailSettingsApi.md#apisettingmanagementemailingsendtestemailpost) | **Post** /api/setting-management/emailing/send-test-email | 
*FeaturesApi* | [**ApiFeatureManagementFeaturesDelete**](docs/FeaturesApi.md#apifeaturemanagementfeaturesdelete) | **Delete** /api/feature-management/features | 
*FeaturesApi* | [**ApiFeatureManagementFeaturesGet**](docs/FeaturesApi.md#apifeaturemanagementfeaturesget) | **Get** /api/feature-management/features | 
*FeaturesApi* | [**ApiFeatureManagementFeaturesPut**](docs/FeaturesApi.md#apifeaturemanagementfeaturesput) | **Put** /api/feature-management/features | 
*KeyValueApi* | [**ApiAppKeyValueBoolGet**](docs/KeyValueApi.md#apiappkeyvalueboolget) | **Get** /api/app/key-value/bool | 
*KeyValueApi* | [**ApiAppKeyValueDateTimeGet**](docs/KeyValueApi.md#apiappkeyvaluedatetimeget) | **Get** /api/app/key-value/date-time | 
*KeyValueApi* | [**ApiAppKeyValueDecimalGet**](docs/KeyValueApi.md#apiappkeyvaluedecimalget) | **Get** /api/app/key-value/decimal | 
*KeyValueApi* | [**ApiAppKeyValueDoubleGet**](docs/KeyValueApi.md#apiappkeyvaluedoubleget) | **Get** /api/app/key-value/double | 
*KeyValueApi* | [**ApiAppKeyValueIntGet**](docs/KeyValueApi.md#apiappkeyvalueintget) | **Get** /api/app/key-value/int | 
*KeyValueApi* | [**ApiAppKeyValueSetBoolPost**](docs/KeyValueApi.md#apiappkeyvaluesetboolpost) | **Post** /api/app/key-value/set-bool | 
*KeyValueApi* | [**ApiAppKeyValueSetDateTimePost**](docs/KeyValueApi.md#apiappkeyvaluesetdatetimepost) | **Post** /api/app/key-value/set-date-time | 
*KeyValueApi* | [**ApiAppKeyValueSetDecimalPost**](docs/KeyValueApi.md#apiappkeyvaluesetdecimalpost) | **Post** /api/app/key-value/set-decimal | 
*KeyValueApi* | [**ApiAppKeyValueSetDoublePost**](docs/KeyValueApi.md#apiappkeyvaluesetdoublepost) | **Post** /api/app/key-value/set-double | 
*KeyValueApi* | [**ApiAppKeyValueSetIntPost**](docs/KeyValueApi.md#apiappkeyvaluesetintpost) | **Post** /api/app/key-value/set-int | 
*KeyValueApi* | [**ApiAppKeyValueSetStringPost**](docs/KeyValueApi.md#apiappkeyvaluesetstringpost) | **Post** /api/app/key-value/set-string | 
*KeyValueApi* | [**ApiAppKeyValueStringGet**](docs/KeyValueApi.md#apiappkeyvaluestringget) | **Get** /api/app/key-value/string | 
*LoginApi* | [**ApiAccountCheckPasswordPost**](docs/LoginApi.md#apiaccountcheckpasswordpost) | **Post** /api/account/check-password | 
*LoginApi* | [**ApiAccountLoginPost**](docs/LoginApi.md#apiaccountloginpost) | **Post** /api/account/login | 
*LoginApi* | [**ApiAccountLogoutGet**](docs/LoginApi.md#apiaccountlogoutget) | **Get** /api/account/logout | 
*MessageApi* | [**ApiAppMessagePublishPost**](docs/MessageApi.md#apiappmessagepublishpost) | **Post** /api/app/message/publish | 
*MessageApi* | [**ApiAppMessageRecallPost**](docs/MessageApi.md#apiappmessagerecallpost) | **Post** /api/app/message/recall | 
*MessageApi* | [**ApiAppMessageSubscribePost**](docs/MessageApi.md#apiappmessagesubscribepost) | **Post** /api/app/message/subscribe | 
*MessageApi* | [**ApiAppMessageUnsubscribePost**](docs/MessageApi.md#apiappmessageunsubscribepost) | **Post** /api/app/message/unsubscribe | 
*MessageTemplateApi* | [**ApiAppMessageTemplateGet**](docs/MessageTemplateApi.md#apiappmessagetemplateget) | **Get** /api/app/message-template | 
*MessageTemplateApi* | [**ApiAppMessageTemplateIdDelete**](docs/MessageTemplateApi.md#apiappmessagetemplateiddelete) | **Delete** /api/app/message-template/{id} | 
*MessageTemplateApi* | [**ApiAppMessageTemplateIdGet**](docs/MessageTemplateApi.md#apiappmessagetemplateidget) | **Get** /api/app/message-template/{id} | 
*MessageTemplateApi* | [**ApiAppMessageTemplateIdPut**](docs/MessageTemplateApi.md#apiappmessagetemplateidput) | **Put** /api/app/message-template/{id} | 
*MessageTemplateApi* | [**ApiAppMessageTemplatePost**](docs/MessageTemplateApi.md#apiappmessagetemplatepost) | **Post** /api/app/message-template | 
*MessageTemplateReleaseApi* | [**ApiAppMessageTemplateReleaseGet**](docs/MessageTemplateReleaseApi.md#apiappmessagetemplatereleaseget) | **Get** /api/app/message-template-release | 
*MessageTemplateReleaseApi* | [**ApiAppMessageTemplateReleaseIdGet**](docs/MessageTemplateReleaseApi.md#apiappmessagetemplatereleaseidget) | **Get** /api/app/message-template-release/{id} | 
*MessageTemplateReleaseApi* | [**ApiAppMessageTemplateReleasePost**](docs/MessageTemplateReleaseApi.md#apiappmessagetemplatereleasepost) | **Post** /api/app/message-template-release | 
*NotificationApi* | [**ApiAppNotificationBarkApiKeyMessageGet**](docs/NotificationApi.md#apiappnotificationbarkapikeymessageget) | **Get** /api/app/notification/bark/{apiKey}/{message} | 
*NotificationApi* | [**ApiAppNotificationGet**](docs/NotificationApi.md#apiappnotificationget) | **Get** /api/app/notification | 
*NotificationApi* | [**ApiAppNotificationPushPost**](docs/NotificationApi.md#apiappnotificationpushpost) | **Post** /api/app/notification/push | 
*PermissionsApi* | [**ApiPermissionManagementPermissionsGet**](docs/PermissionsApi.md#apipermissionmanagementpermissionsget) | **Get** /api/permission-management/permissions | 
*PermissionsApi* | [**ApiPermissionManagementPermissionsPut**](docs/PermissionsApi.md#apipermissionmanagementpermissionsput) | **Put** /api/permission-management/permissions | 
*ProfileApi* | [**ApiAccountMyProfileChangePasswordPost**](docs/ProfileApi.md#apiaccountmyprofilechangepasswordpost) | **Post** /api/account/my-profile/change-password | 
*ProfileApi* | [**ApiAccountMyProfileGet**](docs/ProfileApi.md#apiaccountmyprofileget) | **Get** /api/account/my-profile | 
*ProfileApi* | [**ApiAccountMyProfilePut**](docs/ProfileApi.md#apiaccountmyprofileput) | **Put** /api/account/my-profile | 
*PuupeeApi* | [**ApiAppPuupeePullGet**](docs/PuupeeApi.md#apiapppuupeepullget) | **Get** /api/app/puupee/pull | 
*PuupeeApi* | [**ApiAppPuupeePushPost**](docs/PuupeeApi.md#apiapppuupeepushpost) | **Post** /api/app/puupee/push | 
*RoleApi* | [**ApiIdentityRolesAllGet**](docs/RoleApi.md#apiidentityrolesallget) | **Get** /api/identity/roles/all | 
*RoleApi* | [**ApiIdentityRolesGet**](docs/RoleApi.md#apiidentityrolesget) | **Get** /api/identity/roles | 
*RoleApi* | [**ApiIdentityRolesIdDelete**](docs/RoleApi.md#apiidentityrolesiddelete) | **Delete** /api/identity/roles/{id} | 
*RoleApi* | [**ApiIdentityRolesIdGet**](docs/RoleApi.md#apiidentityrolesidget) | **Get** /api/identity/roles/{id} | 
*RoleApi* | [**ApiIdentityRolesIdPut**](docs/RoleApi.md#apiidentityrolesidput) | **Put** /api/identity/roles/{id} | 
*RoleApi* | [**ApiIdentityRolesPost**](docs/RoleApi.md#apiidentityrolespost) | **Post** /api/identity/roles | 
*SettingsApi* | [**ApiAppSettingsGet**](docs/SettingsApi.md#apiappsettingsget) | **Get** /api/app/settings | 
*SettingsApi* | [**ApiAppSettingsSetPost**](docs/SettingsApi.md#apiappsettingssetpost) | **Post** /api/app/settings/set | 
*SimpleDataApi* | [**ApiAppSimpleDataGet**](docs/SimpleDataApi.md#apiappsimpledataget) | **Get** /api/app/simple-data | 
*SimpleDataApi* | [**ApiAppSimpleDataIdDelete**](docs/SimpleDataApi.md#apiappsimpledataiddelete) | **Delete** /api/app/simple-data/{id} | 
*SimpleDataApi* | [**ApiAppSimpleDataIdGet**](docs/SimpleDataApi.md#apiappsimpledataidget) | **Get** /api/app/simple-data/{id} | 
*SimpleDataApi* | [**ApiAppSimpleDataSavePost**](docs/SimpleDataApi.md#apiappsimpledatasavepost) | **Post** /api/app/simple-data/save | 
*StorageObjectApi* | [**ApiAppStorageObjectFileGet**](docs/StorageObjectApi.md#apiappstorageobjectfileget) | **Get** /api/app/storage-object/file | 
*StorageObjectApi* | [**ApiAppStorageObjectFileOrCredentialsGet**](docs/StorageObjectApi.md#apiappstorageobjectfileorcredentialsget) | **Get** /api/app/storage-object/file-or-credentials | 
*StorageObjectApi* | [**ApiAppStorageObjectPreSignUrlPost**](docs/StorageObjectApi.md#apiappstorageobjectpresignurlpost) | **Post** /api/app/storage-object/pre-sign-url | 
*StorageObjectApi* | [**ApiAppStorageObjectThumbGet**](docs/StorageObjectApi.md#apiappstorageobjectthumbget) | **Get** /api/app/storage-object/thumb | 
*SubscriptionApi* | [**ApiAppSubscriptionVerifyApplePost**](docs/SubscriptionApi.md#apiappsubscriptionverifyapplepost) | **Post** /api/app/subscription/verify-apple | 
*SyncStateApi* | [**ApiAppSyncStateGet**](docs/SyncStateApi.md#apiappsyncstateget) | **Get** /api/app/sync-state | 
*SyncStateApi* | [**ApiAppSyncStatePuupeeChangedEtoPost**](docs/SyncStateApi.md#apiappsyncstatepuupeechangedetopost) | **Post** /api/app/sync-state/puupee-changed-eto | 
*TenantApi* | [**ApiMultiTenancyTenantsGet**](docs/TenantApi.md#apimultitenancytenantsget) | **Get** /api/multi-tenancy/tenants | 
*TenantApi* | [**ApiMultiTenancyTenantsIdDefaultConnectionStringDelete**](docs/TenantApi.md#apimultitenancytenantsiddefaultconnectionstringdelete) | **Delete** /api/multi-tenancy/tenants/{id}/default-connection-string | 
*TenantApi* | [**ApiMultiTenancyTenantsIdDefaultConnectionStringGet**](docs/TenantApi.md#apimultitenancytenantsiddefaultconnectionstringget) | **Get** /api/multi-tenancy/tenants/{id}/default-connection-string | 
*TenantApi* | [**ApiMultiTenancyTenantsIdDefaultConnectionStringPut**](docs/TenantApi.md#apimultitenancytenantsiddefaultconnectionstringput) | **Put** /api/multi-tenancy/tenants/{id}/default-connection-string | 
*TenantApi* | [**ApiMultiTenancyTenantsIdDelete**](docs/TenantApi.md#apimultitenancytenantsiddelete) | **Delete** /api/multi-tenancy/tenants/{id} | 
*TenantApi* | [**ApiMultiTenancyTenantsIdGet**](docs/TenantApi.md#apimultitenancytenantsidget) | **Get** /api/multi-tenancy/tenants/{id} | 
*TenantApi* | [**ApiMultiTenancyTenantsIdPut**](docs/TenantApi.md#apimultitenancytenantsidput) | **Put** /api/multi-tenancy/tenants/{id} | 
*TenantApi* | [**ApiMultiTenancyTenantsPost**](docs/TenantApi.md#apimultitenancytenantspost) | **Post** /api/multi-tenancy/tenants | 
*TestApi* | [**ApiTestDatetimeGet**](docs/TestApi.md#apitestdatetimeget) | **Get** /api/Test/datetime | 
*UserApi* | [**ApiIdentityUsersAssignableRolesGet**](docs/UserApi.md#apiidentityusersassignablerolesget) | **Get** /api/identity/users/assignable-roles | 
*UserApi* | [**ApiIdentityUsersByEmailEmailGet**](docs/UserApi.md#apiidentityusersbyemailemailget) | **Get** /api/identity/users/by-email/{email} | 
*UserApi* | [**ApiIdentityUsersByUsernameUserNameGet**](docs/UserApi.md#apiidentityusersbyusernameusernameget) | **Get** /api/identity/users/by-username/{userName} | 
*UserApi* | [**ApiIdentityUsersGet**](docs/UserApi.md#apiidentityusersget) | **Get** /api/identity/users | 
*UserApi* | [**ApiIdentityUsersIdDelete**](docs/UserApi.md#apiidentityusersiddelete) | **Delete** /api/identity/users/{id} | 
*UserApi* | [**ApiIdentityUsersIdGet**](docs/UserApi.md#apiidentityusersidget) | **Get** /api/identity/users/{id} | 
*UserApi* | [**ApiIdentityUsersIdPut**](docs/UserApi.md#apiidentityusersidput) | **Put** /api/identity/users/{id} | 
*UserApi* | [**ApiIdentityUsersIdRolesGet**](docs/UserApi.md#apiidentityusersidrolesget) | **Get** /api/identity/users/{id}/roles | 
*UserApi* | [**ApiIdentityUsersIdRolesPut**](docs/UserApi.md#apiidentityusersidrolesput) | **Put** /api/identity/users/{id}/roles | 
*UserApi* | [**ApiIdentityUsersPost**](docs/UserApi.md#apiidentityuserspost) | **Post** /api/identity/users | 
*UserLookupApi* | [**ApiIdentityUsersLookupByUsernameUserNameGet**](docs/UserLookupApi.md#apiidentityuserslookupbyusernameusernameget) | **Get** /api/identity/users/lookup/by-username/{userName} | 
*UserLookupApi* | [**ApiIdentityUsersLookupCountGet**](docs/UserLookupApi.md#apiidentityuserslookupcountget) | **Get** /api/identity/users/lookup/count | 
*UserLookupApi* | [**ApiIdentityUsersLookupIdGet**](docs/UserLookupApi.md#apiidentityuserslookupidget) | **Get** /api/identity/users/lookup/{id} | 
*UserLookupApi* | [**ApiIdentityUsersLookupSearchGet**](docs/UserLookupApi.md#apiidentityuserslookupsearchget) | **Get** /api/identity/users/lookup/search | 
*UserStorageApi* | [**ApiAppUserStorageGet**](docs/UserStorageApi.md#apiappuserstorageget) | **Get** /api/app/user-storage | 
*VerificationApi* | [**ApiAppVerificationSendCodePost**](docs/VerificationApi.md#apiappverificationsendcodepost) | **Post** /api/app/verification/send-code | 


## Documentation For Models

 - [AbpLoginResult](docs/AbpLoginResult.md)
 - [ActionApiDescriptionModel](docs/ActionApiDescriptionModel.md)
 - [AppDto](docs/AppDto.md)
 - [AppDtoPagedResultDto](docs/AppDtoPagedResultDto.md)
 - [AppFeatureDto](docs/AppFeatureDto.md)
 - [AppPricingDto](docs/AppPricingDto.md)
 - [AppPricingDtoPagedResultDto](docs/AppPricingDtoPagedResultDto.md)
 - [AppPricingItemDto](docs/AppPricingItemDto.md)
 - [AppReleaseDto](docs/AppReleaseDto.md)
 - [AppReleaseDtoPagedResultDto](docs/AppReleaseDtoPagedResultDto.md)
 - [AppSdkDto](docs/AppSdkDto.md)
 - [AppTheme](docs/AppTheme.md)
 - [AppThemeMode](docs/AppThemeMode.md)
 - [AppUserScoreDto](docs/AppUserScoreDto.md)
 - [AppWithUserDto](docs/AppWithUserDto.md)
 - [AppWithUserDtoPagedResultDto](docs/AppWithUserDtoPagedResultDto.md)
 - [ApplicationApiDescriptionModel](docs/ApplicationApiDescriptionModel.md)
 - [ApplicationAuthConfigurationDto](docs/ApplicationAuthConfigurationDto.md)
 - [ApplicationConfigurationDto](docs/ApplicationConfigurationDto.md)
 - [ApplicationFeatureConfigurationDto](docs/ApplicationFeatureConfigurationDto.md)
 - [ApplicationGlobalFeatureConfigurationDto](docs/ApplicationGlobalFeatureConfigurationDto.md)
 - [ApplicationLocalizationConfigurationDto](docs/ApplicationLocalizationConfigurationDto.md)
 - [ApplicationLocalizationDto](docs/ApplicationLocalizationDto.md)
 - [ApplicationLocalizationResourceDto](docs/ApplicationLocalizationResourceDto.md)
 - [ApplicationSettingConfigurationDto](docs/ApplicationSettingConfigurationDto.md)
 - [BindDeviceDto](docs/BindDeviceDto.md)
 - [BooleanKeyValue](docs/BooleanKeyValue.md)
 - [BooleanSetKeyValueDto](docs/BooleanSetKeyValueDto.md)
 - [ChangePasswordInput](docs/ChangePasswordInput.md)
 - [ClockDto](docs/ClockDto.md)
 - [ControllerApiDescriptionModel](docs/ControllerApiDescriptionModel.md)
 - [ControllerInterfaceApiDescriptionModel](docs/ControllerInterfaceApiDescriptionModel.md)
 - [CreateMessageTemplateReleaseDto](docs/CreateMessageTemplateReleaseDto.md)
 - [CreateOpenIddictApplicationDto](docs/CreateOpenIddictApplicationDto.md)
 - [CreateOrUpdateAppDto](docs/CreateOrUpdateAppDto.md)
 - [CreateOrUpdateAppFeatureDto](docs/CreateOrUpdateAppFeatureDto.md)
 - [CreateOrUpdateAppPricingDto](docs/CreateOrUpdateAppPricingDto.md)
 - [CreateOrUpdateAppPricingItemDto](docs/CreateOrUpdateAppPricingItemDto.md)
 - [CreateOrUpdateAppReleaseDto](docs/CreateOrUpdateAppReleaseDto.md)
 - [CreateOrUpdateAppSdkDto](docs/CreateOrUpdateAppSdkDto.md)
 - [CreateOrUpdateAppUserScoreDto](docs/CreateOrUpdateAppUserScoreDto.md)
 - [CreateOrUpdateMessageTemplateDto](docs/CreateOrUpdateMessageTemplateDto.md)
 - [CreateOrUpdatePuupeeDto](docs/CreateOrUpdatePuupeeDto.md)
 - [CreatePushNotificationDto](docs/CreatePushNotificationDto.md)
 - [CurrentCultureDto](docs/CurrentCultureDto.md)
 - [CurrentTenantDto](docs/CurrentTenantDto.md)
 - [CurrentUserDto](docs/CurrentUserDto.md)
 - [DateTimeFormatDto](docs/DateTimeFormatDto.md)
 - [DateTimeKeyValue](docs/DateTimeKeyValue.md)
 - [DateTimeSetKeyValueDto](docs/DateTimeSetKeyValueDto.md)
 - [DecimalKeyValue](docs/DecimalKeyValue.md)
 - [DecimalSetKeyValueDto](docs/DecimalSetKeyValueDto.md)
 - [DeviceDto](docs/DeviceDto.md)
 - [DeviceDtoPagedResultDto](docs/DeviceDtoPagedResultDto.md)
 - [DoubleKeyValue](docs/DoubleKeyValue.md)
 - [DoubleSetKeyValueDto](docs/DoubleSetKeyValueDto.md)
 - [EmailSettingsDto](docs/EmailSettingsDto.md)
 - [EntityExtensionDto](docs/EntityExtensionDto.md)
 - [ExtensionEnumDto](docs/ExtensionEnumDto.md)
 - [ExtensionEnumFieldDto](docs/ExtensionEnumFieldDto.md)
 - [ExtensionPropertyApiCreateDto](docs/ExtensionPropertyApiCreateDto.md)
 - [ExtensionPropertyApiDto](docs/ExtensionPropertyApiDto.md)
 - [ExtensionPropertyApiGetDto](docs/ExtensionPropertyApiGetDto.md)
 - [ExtensionPropertyApiUpdateDto](docs/ExtensionPropertyApiUpdateDto.md)
 - [ExtensionPropertyAttributeDto](docs/ExtensionPropertyAttributeDto.md)
 - [ExtensionPropertyDto](docs/ExtensionPropertyDto.md)
 - [ExtensionPropertyUiDto](docs/ExtensionPropertyUiDto.md)
 - [ExtensionPropertyUiFormDto](docs/ExtensionPropertyUiFormDto.md)
 - [ExtensionPropertyUiLookupDto](docs/ExtensionPropertyUiLookupDto.md)
 - [ExtensionPropertyUiTableDto](docs/ExtensionPropertyUiTableDto.md)
 - [FeatureDto](docs/FeatureDto.md)
 - [FeatureGroupDto](docs/FeatureGroupDto.md)
 - [FeatureProviderDto](docs/FeatureProviderDto.md)
 - [FindTenantResultDto](docs/FindTenantResultDto.md)
 - [GetFeatureListResultDto](docs/GetFeatureListResultDto.md)
 - [GetPermissionListResultDto](docs/GetPermissionListResultDto.md)
 - [IStringValueType](docs/IStringValueType.md)
 - [IValueValidator](docs/IValueValidator.md)
 - [IanaTimeZone](docs/IanaTimeZone.md)
 - [IdentityRoleCreateDto](docs/IdentityRoleCreateDto.md)
 - [IdentityRoleDto](docs/IdentityRoleDto.md)
 - [IdentityRoleDtoListResultDto](docs/IdentityRoleDtoListResultDto.md)
 - [IdentityRoleDtoPagedResultDto](docs/IdentityRoleDtoPagedResultDto.md)
 - [IdentityRoleUpdateDto](docs/IdentityRoleUpdateDto.md)
 - [IdentityUserCreateDto](docs/IdentityUserCreateDto.md)
 - [IdentityUserDto](docs/IdentityUserDto.md)
 - [IdentityUserDtoPagedResultDto](docs/IdentityUserDtoPagedResultDto.md)
 - [IdentityUserUpdateDto](docs/IdentityUserUpdateDto.md)
 - [IdentityUserUpdateRolesDto](docs/IdentityUserUpdateRolesDto.md)
 - [Int32KeyValue](docs/Int32KeyValue.md)
 - [Int32SetKeyValueDto](docs/Int32SetKeyValueDto.md)
 - [InterfaceMethodApiDescriptionModel](docs/InterfaceMethodApiDescriptionModel.md)
 - [LanguageInfo](docs/LanguageInfo.md)
 - [LocalizableStringDto](docs/LocalizableStringDto.md)
 - [LoginResultType](docs/LoginResultType.md)
 - [MessagePublishDto](docs/MessagePublishDto.md)
 - [MessageRecallDto](docs/MessageRecallDto.md)
 - [MessageSubscribeDto](docs/MessageSubscribeDto.md)
 - [MessageTemplateDto](docs/MessageTemplateDto.md)
 - [MessageTemplateReleaseDto](docs/MessageTemplateReleaseDto.md)
 - [MessageUnsubscribeDto](docs/MessageUnsubscribeDto.md)
 - [MethodParameterApiDescriptionModel](docs/MethodParameterApiDescriptionModel.md)
 - [ModuleApiDescriptionModel](docs/ModuleApiDescriptionModel.md)
 - [ModuleExtensionDto](docs/ModuleExtensionDto.md)
 - [MultiTenancyInfoDto](docs/MultiTenancyInfoDto.md)
 - [NameValue](docs/NameValue.md)
 - [NotificationInfoDto](docs/NotificationInfoDto.md)
 - [NotificationInfoDtoPagedResultDto](docs/NotificationInfoDtoPagedResultDto.md)
 - [ObjectExtensionsDto](docs/ObjectExtensionsDto.md)
 - [ParameterApiDescriptionModel](docs/ParameterApiDescriptionModel.md)
 - [PermissionGrantInfoDto](docs/PermissionGrantInfoDto.md)
 - [PermissionGroupDto](docs/PermissionGroupDto.md)
 - [ProfileDto](docs/ProfileDto.md)
 - [PropertyApiDescriptionModel](docs/PropertyApiDescriptionModel.md)
 - [ProviderInfoDto](docs/ProviderInfoDto.md)
 - [PuupeeChangedEto](docs/PuupeeChangedEto.md)
 - [PuupeeDto](docs/PuupeeDto.md)
 - [PuupeeDtoPagedResultDto](docs/PuupeeDtoPagedResultDto.md)
 - [RefreshDeviceStatusDto](docs/RefreshDeviceStatusDto.md)
 - [RegisterDto](docs/RegisterDto.md)
 - [RemoteServiceErrorInfo](docs/RemoteServiceErrorInfo.md)
 - [RemoteServiceErrorResponse](docs/RemoteServiceErrorResponse.md)
 - [RemoteServiceValidationErrorInfo](docs/RemoteServiceValidationErrorInfo.md)
 - [ResetPasswordDto](docs/ResetPasswordDto.md)
 - [ReturnValueApiDescriptionModel](docs/ReturnValueApiDescriptionModel.md)
 - [SendPasswordResetCodeDto](docs/SendPasswordResetCodeDto.md)
 - [SendTestEmailInput](docs/SendTestEmailInput.md)
 - [SendVerificationCodeDto](docs/SendVerificationCodeDto.md)
 - [SettingsDto](docs/SettingsDto.md)
 - [SimpleDataDto](docs/SimpleDataDto.md)
 - [SimpleDataDtoPagedResultDto](docs/SimpleDataDtoPagedResultDto.md)
 - [StorageObjectCredentials](docs/StorageObjectCredentials.md)
 - [StorageObjectDto](docs/StorageObjectDto.md)
 - [StorageObjectOrCredentialsDto](docs/StorageObjectOrCredentialsDto.md)
 - [StringKeyValue](docs/StringKeyValue.md)
 - [StringSetKeyValueDto](docs/StringSetKeyValueDto.md)
 - [SyncStateDto](docs/SyncStateDto.md)
 - [TenantCreateDto](docs/TenantCreateDto.md)
 - [TenantDto](docs/TenantDto.md)
 - [TenantDtoPagedResultDto](docs/TenantDtoPagedResultDto.md)
 - [TenantUpdateDto](docs/TenantUpdateDto.md)
 - [TestDateTime](docs/TestDateTime.md)
 - [TimeZone](docs/TimeZone.md)
 - [TimingDto](docs/TimingDto.md)
 - [TodoOrderBy](docs/TodoOrderBy.md)
 - [TodoSettingsDto](docs/TodoSettingsDto.md)
 - [TypeApiDescriptionModel](docs/TypeApiDescriptionModel.md)
 - [UpdateEmailSettingsDto](docs/UpdateEmailSettingsDto.md)
 - [UpdateFeatureDto](docs/UpdateFeatureDto.md)
 - [UpdateFeaturesDto](docs/UpdateFeaturesDto.md)
 - [UpdatePermissionDto](docs/UpdatePermissionDto.md)
 - [UpdatePermissionsDto](docs/UpdatePermissionsDto.md)
 - [UpdateProfileDto](docs/UpdateProfileDto.md)
 - [UserData](docs/UserData.md)
 - [UserDataListResultDto](docs/UserDataListResultDto.md)
 - [UserLoginInfo](docs/UserLoginInfo.md)
 - [UserStorageDto](docs/UserStorageDto.md)
 - [UserStorageItemDto](docs/UserStorageItemDto.md)
 - [VerifyPasswordResetTokenInput](docs/VerifyPasswordResetTokenInput.md)
 - [WindowsTimeZone](docs/WindowsTimeZone.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://localhost:44355/connect/authorize
- **Scopes**: 
 - **Puupees**: Puupee API

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




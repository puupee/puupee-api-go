# CreatePushNotificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BodyType** | Pointer to [**NotificationBodyType**](NotificationBodyType.md) |  | [optional] 
**Body** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreatePushNotificationDto

`func NewCreatePushNotificationDto() *CreatePushNotificationDto`

NewCreatePushNotificationDto instantiates a new CreatePushNotificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePushNotificationDtoWithDefaults

`func NewCreatePushNotificationDtoWithDefaults() *CreatePushNotificationDto`

NewCreatePushNotificationDtoWithDefaults instantiates a new CreatePushNotificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *CreatePushNotificationDto) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *CreatePushNotificationDto) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *CreatePushNotificationDto) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *CreatePushNotificationDto) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *CreatePushNotificationDto) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *CreatePushNotificationDto) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetTitle

`func (o *CreatePushNotificationDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreatePushNotificationDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreatePushNotificationDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreatePushNotificationDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreatePushNotificationDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreatePushNotificationDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CreatePushNotificationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePushNotificationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePushNotificationDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePushNotificationDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreatePushNotificationDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreatePushNotificationDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBodyType

`func (o *CreatePushNotificationDto) GetBodyType() NotificationBodyType`

GetBodyType returns the BodyType field if non-nil, zero value otherwise.

### GetBodyTypeOk

`func (o *CreatePushNotificationDto) GetBodyTypeOk() (*NotificationBodyType, bool)`

GetBodyTypeOk returns a tuple with the BodyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyType

`func (o *CreatePushNotificationDto) SetBodyType(v NotificationBodyType)`

SetBodyType sets BodyType field to given value.

### HasBodyType

`func (o *CreatePushNotificationDto) HasBodyType() bool`

HasBodyType returns a boolean if a field has been set.

### GetBody

`func (o *CreatePushNotificationDto) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreatePushNotificationDto) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreatePushNotificationDto) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreatePushNotificationDto) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *CreatePushNotificationDto) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *CreatePushNotificationDto) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetUrl

`func (o *CreatePushNotificationDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreatePushNotificationDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreatePushNotificationDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreatePushNotificationDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CreatePushNotificationDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CreatePushNotificationDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



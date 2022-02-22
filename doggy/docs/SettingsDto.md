# SettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppThemes** | Pointer to [**AppThemePlatformSettings**](AppThemePlatformSettings.md) |  | [optional] 

## Methods

### NewSettingsDto

`func NewSettingsDto() *SettingsDto`

NewSettingsDto instantiates a new SettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsDtoWithDefaults

`func NewSettingsDtoWithDefaults() *SettingsDto`

NewSettingsDtoWithDefaults instantiates a new SettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppThemes

`func (o *SettingsDto) GetAppThemes() AppThemePlatformSettings`

GetAppThemes returns the AppThemes field if non-nil, zero value otherwise.

### GetAppThemesOk

`func (o *SettingsDto) GetAppThemesOk() (*AppThemePlatformSettings, bool)`

GetAppThemesOk returns a tuple with the AppThemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppThemes

`func (o *SettingsDto) SetAppThemes(v AppThemePlatformSettings)`

SetAppThemes sets AppThemes field to given value.

### HasAppThemes

`func (o *SettingsDto) HasAppThemes() bool`

HasAppThemes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



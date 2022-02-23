# SettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppTheme** | Pointer to [**AppTheme**](AppTheme.md) |  | [optional] 

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

### GetAppTheme

`func (o *SettingsDto) GetAppTheme() AppTheme`

GetAppTheme returns the AppTheme field if non-nil, zero value otherwise.

### GetAppThemeOk

`func (o *SettingsDto) GetAppThemeOk() (*AppTheme, bool)`

GetAppThemeOk returns a tuple with the AppTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTheme

`func (o *SettingsDto) SetAppTheme(v AppTheme)`

SetAppTheme sets AppTheme field to given value.

### HasAppTheme

`func (o *SettingsDto) HasAppTheme() bool`

HasAppTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



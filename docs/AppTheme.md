# AppTheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceColor** | Pointer to **NullableString** |  | [optional] 
**ThemeMode** | Pointer to [**AppThemeMode**](AppThemeMode.md) |  | [optional] 

## Methods

### NewAppTheme

`func NewAppTheme() *AppTheme`

NewAppTheme instantiates a new AppTheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppThemeWithDefaults

`func NewAppThemeWithDefaults() *AppTheme`

NewAppThemeWithDefaults instantiates a new AppTheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceColor

`func (o *AppTheme) GetSourceColor() string`

GetSourceColor returns the SourceColor field if non-nil, zero value otherwise.

### GetSourceColorOk

`func (o *AppTheme) GetSourceColorOk() (*string, bool)`

GetSourceColorOk returns a tuple with the SourceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceColor

`func (o *AppTheme) SetSourceColor(v string)`

SetSourceColor sets SourceColor field to given value.

### HasSourceColor

`func (o *AppTheme) HasSourceColor() bool`

HasSourceColor returns a boolean if a field has been set.

### SetSourceColorNil

`func (o *AppTheme) SetSourceColorNil(b bool)`

 SetSourceColorNil sets the value for SourceColor to be an explicit nil

### UnsetSourceColor
`func (o *AppTheme) UnsetSourceColor()`

UnsetSourceColor ensures that no value is present for SourceColor, not even an explicit nil
### GetThemeMode

`func (o *AppTheme) GetThemeMode() AppThemeMode`

GetThemeMode returns the ThemeMode field if non-nil, zero value otherwise.

### GetThemeModeOk

`func (o *AppTheme) GetThemeModeOk() (*AppThemeMode, bool)`

GetThemeModeOk returns a tuple with the ThemeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeMode

`func (o *AppTheme) SetThemeMode(v AppThemeMode)`

SetThemeMode sets ThemeMode field to given value.

### HasThemeMode

`func (o *AppTheme) HasThemeMode() bool`

HasThemeMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



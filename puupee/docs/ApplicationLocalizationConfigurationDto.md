# ApplicationLocalizationConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to **map[string]map[string]string** |  | [optional] 
**Resources** | Pointer to [**map[string]ApplicationLocalizationResourceDto**](ApplicationLocalizationResourceDto.md) |  | [optional] 
**Languages** | Pointer to [**[]LanguageInfo**](LanguageInfo.md) |  | [optional] 
**CurrentCulture** | Pointer to [**CurrentCultureDto**](CurrentCultureDto.md) |  | [optional] 
**DefaultResourceName** | Pointer to **string** |  | [optional] 
**LanguagesMap** | Pointer to [**map[string][]NameValue**](array.md) |  | [optional] 
**LanguageFilesMap** | Pointer to [**map[string][]NameValue**](array.md) |  | [optional] 

## Methods

### NewApplicationLocalizationConfigurationDto

`func NewApplicationLocalizationConfigurationDto() *ApplicationLocalizationConfigurationDto`

NewApplicationLocalizationConfigurationDto instantiates a new ApplicationLocalizationConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLocalizationConfigurationDtoWithDefaults

`func NewApplicationLocalizationConfigurationDtoWithDefaults() *ApplicationLocalizationConfigurationDto`

NewApplicationLocalizationConfigurationDtoWithDefaults instantiates a new ApplicationLocalizationConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *ApplicationLocalizationConfigurationDto) GetValues() map[string]map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ApplicationLocalizationConfigurationDto) GetValuesOk() (*map[string]map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ApplicationLocalizationConfigurationDto) SetValues(v map[string]map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ApplicationLocalizationConfigurationDto) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetResources

`func (o *ApplicationLocalizationConfigurationDto) GetResources() map[string]ApplicationLocalizationResourceDto`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ApplicationLocalizationConfigurationDto) GetResourcesOk() (*map[string]ApplicationLocalizationResourceDto, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ApplicationLocalizationConfigurationDto) SetResources(v map[string]ApplicationLocalizationResourceDto)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ApplicationLocalizationConfigurationDto) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetLanguages

`func (o *ApplicationLocalizationConfigurationDto) GetLanguages() []LanguageInfo`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ApplicationLocalizationConfigurationDto) GetLanguagesOk() (*[]LanguageInfo, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ApplicationLocalizationConfigurationDto) SetLanguages(v []LanguageInfo)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ApplicationLocalizationConfigurationDto) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetCurrentCulture

`func (o *ApplicationLocalizationConfigurationDto) GetCurrentCulture() CurrentCultureDto`

GetCurrentCulture returns the CurrentCulture field if non-nil, zero value otherwise.

### GetCurrentCultureOk

`func (o *ApplicationLocalizationConfigurationDto) GetCurrentCultureOk() (*CurrentCultureDto, bool)`

GetCurrentCultureOk returns a tuple with the CurrentCulture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCulture

`func (o *ApplicationLocalizationConfigurationDto) SetCurrentCulture(v CurrentCultureDto)`

SetCurrentCulture sets CurrentCulture field to given value.

### HasCurrentCulture

`func (o *ApplicationLocalizationConfigurationDto) HasCurrentCulture() bool`

HasCurrentCulture returns a boolean if a field has been set.

### GetDefaultResourceName

`func (o *ApplicationLocalizationConfigurationDto) GetDefaultResourceName() string`

GetDefaultResourceName returns the DefaultResourceName field if non-nil, zero value otherwise.

### GetDefaultResourceNameOk

`func (o *ApplicationLocalizationConfigurationDto) GetDefaultResourceNameOk() (*string, bool)`

GetDefaultResourceNameOk returns a tuple with the DefaultResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultResourceName

`func (o *ApplicationLocalizationConfigurationDto) SetDefaultResourceName(v string)`

SetDefaultResourceName sets DefaultResourceName field to given value.

### HasDefaultResourceName

`func (o *ApplicationLocalizationConfigurationDto) HasDefaultResourceName() bool`

HasDefaultResourceName returns a boolean if a field has been set.

### GetLanguagesMap

`func (o *ApplicationLocalizationConfigurationDto) GetLanguagesMap() map[string][]NameValue`

GetLanguagesMap returns the LanguagesMap field if non-nil, zero value otherwise.

### GetLanguagesMapOk

`func (o *ApplicationLocalizationConfigurationDto) GetLanguagesMapOk() (*map[string][]NameValue, bool)`

GetLanguagesMapOk returns a tuple with the LanguagesMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesMap

`func (o *ApplicationLocalizationConfigurationDto) SetLanguagesMap(v map[string][]NameValue)`

SetLanguagesMap sets LanguagesMap field to given value.

### HasLanguagesMap

`func (o *ApplicationLocalizationConfigurationDto) HasLanguagesMap() bool`

HasLanguagesMap returns a boolean if a field has been set.

### GetLanguageFilesMap

`func (o *ApplicationLocalizationConfigurationDto) GetLanguageFilesMap() map[string][]NameValue`

GetLanguageFilesMap returns the LanguageFilesMap field if non-nil, zero value otherwise.

### GetLanguageFilesMapOk

`func (o *ApplicationLocalizationConfigurationDto) GetLanguageFilesMapOk() (*map[string][]NameValue, bool)`

GetLanguageFilesMapOk returns a tuple with the LanguageFilesMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageFilesMap

`func (o *ApplicationLocalizationConfigurationDto) SetLanguageFilesMap(v map[string][]NameValue)`

SetLanguageFilesMap sets LanguageFilesMap field to given value.

### HasLanguageFilesMap

`func (o *ApplicationLocalizationConfigurationDto) HasLanguageFilesMap() bool`

HasLanguageFilesMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



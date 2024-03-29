# CurrentCultureDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**EnglishName** | Pointer to **string** |  | [optional] 
**ThreeLetterIsoLanguageName** | Pointer to **string** |  | [optional] 
**TwoLetterIsoLanguageName** | Pointer to **string** |  | [optional] 
**IsRightToLeft** | Pointer to **bool** |  | [optional] 
**CultureName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NativeName** | Pointer to **string** |  | [optional] 
**DateTimeFormat** | Pointer to [**DateTimeFormatDto**](DateTimeFormatDto.md) |  | [optional] 

## Methods

### NewCurrentCultureDto

`func NewCurrentCultureDto() *CurrentCultureDto`

NewCurrentCultureDto instantiates a new CurrentCultureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentCultureDtoWithDefaults

`func NewCurrentCultureDtoWithDefaults() *CurrentCultureDto`

NewCurrentCultureDtoWithDefaults instantiates a new CurrentCultureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CurrentCultureDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CurrentCultureDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CurrentCultureDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CurrentCultureDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnglishName

`func (o *CurrentCultureDto) GetEnglishName() string`

GetEnglishName returns the EnglishName field if non-nil, zero value otherwise.

### GetEnglishNameOk

`func (o *CurrentCultureDto) GetEnglishNameOk() (*string, bool)`

GetEnglishNameOk returns a tuple with the EnglishName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnglishName

`func (o *CurrentCultureDto) SetEnglishName(v string)`

SetEnglishName sets EnglishName field to given value.

### HasEnglishName

`func (o *CurrentCultureDto) HasEnglishName() bool`

HasEnglishName returns a boolean if a field has been set.

### GetThreeLetterIsoLanguageName

`func (o *CurrentCultureDto) GetThreeLetterIsoLanguageName() string`

GetThreeLetterIsoLanguageName returns the ThreeLetterIsoLanguageName field if non-nil, zero value otherwise.

### GetThreeLetterIsoLanguageNameOk

`func (o *CurrentCultureDto) GetThreeLetterIsoLanguageNameOk() (*string, bool)`

GetThreeLetterIsoLanguageNameOk returns a tuple with the ThreeLetterIsoLanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeLetterIsoLanguageName

`func (o *CurrentCultureDto) SetThreeLetterIsoLanguageName(v string)`

SetThreeLetterIsoLanguageName sets ThreeLetterIsoLanguageName field to given value.

### HasThreeLetterIsoLanguageName

`func (o *CurrentCultureDto) HasThreeLetterIsoLanguageName() bool`

HasThreeLetterIsoLanguageName returns a boolean if a field has been set.

### GetTwoLetterIsoLanguageName

`func (o *CurrentCultureDto) GetTwoLetterIsoLanguageName() string`

GetTwoLetterIsoLanguageName returns the TwoLetterIsoLanguageName field if non-nil, zero value otherwise.

### GetTwoLetterIsoLanguageNameOk

`func (o *CurrentCultureDto) GetTwoLetterIsoLanguageNameOk() (*string, bool)`

GetTwoLetterIsoLanguageNameOk returns a tuple with the TwoLetterIsoLanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoLetterIsoLanguageName

`func (o *CurrentCultureDto) SetTwoLetterIsoLanguageName(v string)`

SetTwoLetterIsoLanguageName sets TwoLetterIsoLanguageName field to given value.

### HasTwoLetterIsoLanguageName

`func (o *CurrentCultureDto) HasTwoLetterIsoLanguageName() bool`

HasTwoLetterIsoLanguageName returns a boolean if a field has been set.

### GetIsRightToLeft

`func (o *CurrentCultureDto) GetIsRightToLeft() bool`

GetIsRightToLeft returns the IsRightToLeft field if non-nil, zero value otherwise.

### GetIsRightToLeftOk

`func (o *CurrentCultureDto) GetIsRightToLeftOk() (*bool, bool)`

GetIsRightToLeftOk returns a tuple with the IsRightToLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRightToLeft

`func (o *CurrentCultureDto) SetIsRightToLeft(v bool)`

SetIsRightToLeft sets IsRightToLeft field to given value.

### HasIsRightToLeft

`func (o *CurrentCultureDto) HasIsRightToLeft() bool`

HasIsRightToLeft returns a boolean if a field has been set.

### GetCultureName

`func (o *CurrentCultureDto) GetCultureName() string`

GetCultureName returns the CultureName field if non-nil, zero value otherwise.

### GetCultureNameOk

`func (o *CurrentCultureDto) GetCultureNameOk() (*string, bool)`

GetCultureNameOk returns a tuple with the CultureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureName

`func (o *CurrentCultureDto) SetCultureName(v string)`

SetCultureName sets CultureName field to given value.

### HasCultureName

`func (o *CurrentCultureDto) HasCultureName() bool`

HasCultureName returns a boolean if a field has been set.

### GetName

`func (o *CurrentCultureDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentCultureDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentCultureDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrentCultureDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeName

`func (o *CurrentCultureDto) GetNativeName() string`

GetNativeName returns the NativeName field if non-nil, zero value otherwise.

### GetNativeNameOk

`func (o *CurrentCultureDto) GetNativeNameOk() (*string, bool)`

GetNativeNameOk returns a tuple with the NativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeName

`func (o *CurrentCultureDto) SetNativeName(v string)`

SetNativeName sets NativeName field to given value.

### HasNativeName

`func (o *CurrentCultureDto) HasNativeName() bool`

HasNativeName returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *CurrentCultureDto) GetDateTimeFormat() DateTimeFormatDto`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *CurrentCultureDto) GetDateTimeFormatOk() (*DateTimeFormatDto, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *CurrentCultureDto) SetDateTimeFormat(v DateTimeFormatDto)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *CurrentCultureDto) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



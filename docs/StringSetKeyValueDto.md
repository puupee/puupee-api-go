# StringSetKeyValueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableString** |  | [optional] 
**DurationSeconds** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewStringSetKeyValueDto

`func NewStringSetKeyValueDto() *StringSetKeyValueDto`

NewStringSetKeyValueDto instantiates a new StringSetKeyValueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringSetKeyValueDtoWithDefaults

`func NewStringSetKeyValueDtoWithDefaults() *StringSetKeyValueDto`

NewStringSetKeyValueDtoWithDefaults instantiates a new StringSetKeyValueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *StringSetKeyValueDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringSetKeyValueDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringSetKeyValueDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StringSetKeyValueDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *StringSetKeyValueDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StringSetKeyValueDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDurationSeconds

`func (o *StringSetKeyValueDto) GetDurationSeconds() float64`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *StringSetKeyValueDto) GetDurationSecondsOk() (*float64, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *StringSetKeyValueDto) SetDurationSeconds(v float64)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *StringSetKeyValueDto) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *StringSetKeyValueDto) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *StringSetKeyValueDto) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



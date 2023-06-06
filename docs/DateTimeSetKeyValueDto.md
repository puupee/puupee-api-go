# DateTimeSetKeyValueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **time.Time** |  | [optional] 
**DurationSeconds** | Pointer to **float64** |  | [optional] 

## Methods

### NewDateTimeSetKeyValueDto

`func NewDateTimeSetKeyValueDto() *DateTimeSetKeyValueDto`

NewDateTimeSetKeyValueDto instantiates a new DateTimeSetKeyValueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeSetKeyValueDtoWithDefaults

`func NewDateTimeSetKeyValueDtoWithDefaults() *DateTimeSetKeyValueDto`

NewDateTimeSetKeyValueDtoWithDefaults instantiates a new DateTimeSetKeyValueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DateTimeSetKeyValueDto) GetValue() time.Time`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DateTimeSetKeyValueDto) GetValueOk() (*time.Time, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DateTimeSetKeyValueDto) SetValue(v time.Time)`

SetValue sets Value field to given value.

### HasValue

`func (o *DateTimeSetKeyValueDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *DateTimeSetKeyValueDto) GetDurationSeconds() float64`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *DateTimeSetKeyValueDto) GetDurationSecondsOk() (*float64, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *DateTimeSetKeyValueDto) SetDurationSeconds(v float64)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *DateTimeSetKeyValueDto) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



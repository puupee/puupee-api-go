# DateTimeKeyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **time.Time** |  | [optional] 
**DurationSeconds** | Pointer to **float64** |  | [optional] 
**ExpiredAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDateTimeKeyValue

`func NewDateTimeKeyValue() *DateTimeKeyValue`

NewDateTimeKeyValue instantiates a new DateTimeKeyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeKeyValueWithDefaults

`func NewDateTimeKeyValueWithDefaults() *DateTimeKeyValue`

NewDateTimeKeyValueWithDefaults instantiates a new DateTimeKeyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DateTimeKeyValue) GetValue() time.Time`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DateTimeKeyValue) GetValueOk() (*time.Time, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DateTimeKeyValue) SetValue(v time.Time)`

SetValue sets Value field to given value.

### HasValue

`func (o *DateTimeKeyValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *DateTimeKeyValue) GetDurationSeconds() float64`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *DateTimeKeyValue) GetDurationSecondsOk() (*float64, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *DateTimeKeyValue) SetDurationSeconds(v float64)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *DateTimeKeyValue) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetExpiredAt

`func (o *DateTimeKeyValue) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *DateTimeKeyValue) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *DateTimeKeyValue) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *DateTimeKeyValue) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DateTimeKeyValue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DateTimeKeyValue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DateTimeKeyValue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DateTimeKeyValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



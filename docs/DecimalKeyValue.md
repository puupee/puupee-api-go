# DecimalKeyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float64** |  | [optional] 
**DurationSeconds** | Pointer to **float64** |  | [optional] 
**ExpiredAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDecimalKeyValue

`func NewDecimalKeyValue() *DecimalKeyValue`

NewDecimalKeyValue instantiates a new DecimalKeyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecimalKeyValueWithDefaults

`func NewDecimalKeyValueWithDefaults() *DecimalKeyValue`

NewDecimalKeyValueWithDefaults instantiates a new DecimalKeyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DecimalKeyValue) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DecimalKeyValue) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DecimalKeyValue) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *DecimalKeyValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *DecimalKeyValue) GetDurationSeconds() float64`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *DecimalKeyValue) GetDurationSecondsOk() (*float64, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *DecimalKeyValue) SetDurationSeconds(v float64)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *DecimalKeyValue) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetExpiredAt

`func (o *DecimalKeyValue) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *DecimalKeyValue) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *DecimalKeyValue) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *DecimalKeyValue) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DecimalKeyValue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DecimalKeyValue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DecimalKeyValue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DecimalKeyValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StringKeyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**DurationSeconds** | Pointer to **float64** |  | [optional] 
**ExpiredAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStringKeyValue

`func NewStringKeyValue() *StringKeyValue`

NewStringKeyValue instantiates a new StringKeyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringKeyValueWithDefaults

`func NewStringKeyValueWithDefaults() *StringKeyValue`

NewStringKeyValueWithDefaults instantiates a new StringKeyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *StringKeyValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringKeyValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringKeyValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StringKeyValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *StringKeyValue) GetDurationSeconds() float64`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *StringKeyValue) GetDurationSecondsOk() (*float64, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *StringKeyValue) SetDurationSeconds(v float64)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *StringKeyValue) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetExpiredAt

`func (o *StringKeyValue) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *StringKeyValue) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *StringKeyValue) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *StringKeyValue) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StringKeyValue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StringKeyValue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StringKeyValue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StringKeyValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TestDateTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Puupee** | Pointer to [**PuupeeDto**](PuupeeDto.md) |  | [optional] 

## Methods

### NewTestDateTime

`func NewTestDateTime() *TestDateTime`

NewTestDateTime instantiates a new TestDateTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestDateTimeWithDefaults

`func NewTestDateTimeWithDefaults() *TestDateTime`

NewTestDateTimeWithDefaults instantiates a new TestDateTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TestDateTime) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestDateTime) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestDateTime) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TestDateTime) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPuupee

`func (o *TestDateTime) GetPuupee() PuupeeDto`

GetPuupee returns the Puupee field if non-nil, zero value otherwise.

### GetPuupeeOk

`func (o *TestDateTime) GetPuupeeOk() (*PuupeeDto, bool)`

GetPuupeeOk returns a tuple with the Puupee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuupee

`func (o *TestDateTime) SetPuupee(v PuupeeDto)`

SetPuupee sets Puupee field to given value.

### HasPuupee

`func (o *TestDateTime) HasPuupee() bool`

HasPuupee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



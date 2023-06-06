# ApiKeyCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Key** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**ExpireAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewApiKeyCreateDto

`func NewApiKeyCreateDto(name string, key string, ) *ApiKeyCreateDto`

NewApiKeyCreateDto instantiates a new ApiKeyCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyCreateDtoWithDefaults

`func NewApiKeyCreateDtoWithDefaults() *ApiKeyCreateDto`

NewApiKeyCreateDtoWithDefaults instantiates a new ApiKeyCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiKeyCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ApiKeyCreateDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKeyCreateDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiKeyCreateDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetActive

`func (o *ApiKeyCreateDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiKeyCreateDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiKeyCreateDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiKeyCreateDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetExpireAt

`func (o *ApiKeyCreateDto) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *ApiKeyCreateDto) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *ApiKeyCreateDto) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *ApiKeyCreateDto) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



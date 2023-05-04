# CreateOrGetSubscriptionOrderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**LastModificationTime** | Pointer to **time.Time** |  | [optional] 
**LastModifierId** | Pointer to **string** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**DeleterId** | Pointer to **string** |  | [optional] 
**DeletionTime** | Pointer to **time.Time** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**PricingId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOrGetSubscriptionOrderDto

`func NewCreateOrGetSubscriptionOrderDto() *CreateOrGetSubscriptionOrderDto`

NewCreateOrGetSubscriptionOrderDto instantiates a new CreateOrGetSubscriptionOrderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrGetSubscriptionOrderDtoWithDefaults

`func NewCreateOrGetSubscriptionOrderDtoWithDefaults() *CreateOrGetSubscriptionOrderDto`

NewCreateOrGetSubscriptionOrderDtoWithDefaults instantiates a new CreateOrGetSubscriptionOrderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOrGetSubscriptionOrderDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrGetSubscriptionOrderDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrGetSubscriptionOrderDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrGetSubscriptionOrderDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *CreateOrGetSubscriptionOrderDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreateOrGetSubscriptionOrderDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreateOrGetSubscriptionOrderDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CreateOrGetSubscriptionOrderDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *CreateOrGetSubscriptionOrderDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CreateOrGetSubscriptionOrderDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CreateOrGetSubscriptionOrderDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *CreateOrGetSubscriptionOrderDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *CreateOrGetSubscriptionOrderDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *CreateOrGetSubscriptionOrderDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *CreateOrGetSubscriptionOrderDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *CreateOrGetSubscriptionOrderDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *CreateOrGetSubscriptionOrderDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *CreateOrGetSubscriptionOrderDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *CreateOrGetSubscriptionOrderDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *CreateOrGetSubscriptionOrderDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *CreateOrGetSubscriptionOrderDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CreateOrGetSubscriptionOrderDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CreateOrGetSubscriptionOrderDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CreateOrGetSubscriptionOrderDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *CreateOrGetSubscriptionOrderDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *CreateOrGetSubscriptionOrderDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *CreateOrGetSubscriptionOrderDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *CreateOrGetSubscriptionOrderDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *CreateOrGetSubscriptionOrderDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *CreateOrGetSubscriptionOrderDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *CreateOrGetSubscriptionOrderDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *CreateOrGetSubscriptionOrderDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetAppId

`func (o *CreateOrGetSubscriptionOrderDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateOrGetSubscriptionOrderDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateOrGetSubscriptionOrderDto) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CreateOrGetSubscriptionOrderDto) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetPricingId

`func (o *CreateOrGetSubscriptionOrderDto) GetPricingId() string`

GetPricingId returns the PricingId field if non-nil, zero value otherwise.

### GetPricingIdOk

`func (o *CreateOrGetSubscriptionOrderDto) GetPricingIdOk() (*string, bool)`

GetPricingIdOk returns a tuple with the PricingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingId

`func (o *CreateOrGetSubscriptionOrderDto) SetPricingId(v string)`

SetPricingId sets PricingId field to given value.

### HasPricingId

`func (o *CreateOrGetSubscriptionOrderDto) HasPricingId() bool`

HasPricingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



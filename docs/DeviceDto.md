# DeviceDto

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
**Token** | Pointer to **string** |  | [optional] 
**TpnsToken** | Pointer to **string** |  | [optional] 
**IsPhysicalDevice** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to [**AppPlatform**](AppPlatform.md) |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**SystemVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceDto

`func NewDeviceDto() *DeviceDto`

NewDeviceDto instantiates a new DeviceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceDtoWithDefaults

`func NewDeviceDtoWithDefaults() *DeviceDto`

NewDeviceDtoWithDefaults instantiates a new DeviceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *DeviceDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DeviceDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DeviceDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *DeviceDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *DeviceDto) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *DeviceDto) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *DeviceDto) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *DeviceDto) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *DeviceDto) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *DeviceDto) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *DeviceDto) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *DeviceDto) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLastModifierId

`func (o *DeviceDto) GetLastModifierId() string`

GetLastModifierId returns the LastModifierId field if non-nil, zero value otherwise.

### GetLastModifierIdOk

`func (o *DeviceDto) GetLastModifierIdOk() (*string, bool)`

GetLastModifierIdOk returns a tuple with the LastModifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifierId

`func (o *DeviceDto) SetLastModifierId(v string)`

SetLastModifierId sets LastModifierId field to given value.

### HasLastModifierId

`func (o *DeviceDto) HasLastModifierId() bool`

HasLastModifierId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *DeviceDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DeviceDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DeviceDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *DeviceDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeleterId

`func (o *DeviceDto) GetDeleterId() string`

GetDeleterId returns the DeleterId field if non-nil, zero value otherwise.

### GetDeleterIdOk

`func (o *DeviceDto) GetDeleterIdOk() (*string, bool)`

GetDeleterIdOk returns a tuple with the DeleterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleterId

`func (o *DeviceDto) SetDeleterId(v string)`

SetDeleterId sets DeleterId field to given value.

### HasDeleterId

`func (o *DeviceDto) HasDeleterId() bool`

HasDeleterId returns a boolean if a field has been set.

### GetDeletionTime

`func (o *DeviceDto) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *DeviceDto) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *DeviceDto) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *DeviceDto) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetToken

`func (o *DeviceDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeviceDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeviceDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeviceDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTpnsToken

`func (o *DeviceDto) GetTpnsToken() string`

GetTpnsToken returns the TpnsToken field if non-nil, zero value otherwise.

### GetTpnsTokenOk

`func (o *DeviceDto) GetTpnsTokenOk() (*string, bool)`

GetTpnsTokenOk returns a tuple with the TpnsToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpnsToken

`func (o *DeviceDto) SetTpnsToken(v string)`

SetTpnsToken sets TpnsToken field to given value.

### HasTpnsToken

`func (o *DeviceDto) HasTpnsToken() bool`

HasTpnsToken returns a boolean if a field has been set.

### GetIsPhysicalDevice

`func (o *DeviceDto) GetIsPhysicalDevice() bool`

GetIsPhysicalDevice returns the IsPhysicalDevice field if non-nil, zero value otherwise.

### GetIsPhysicalDeviceOk

`func (o *DeviceDto) GetIsPhysicalDeviceOk() (*bool, bool)`

GetIsPhysicalDeviceOk returns a tuple with the IsPhysicalDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPhysicalDevice

`func (o *DeviceDto) SetIsPhysicalDevice(v bool)`

SetIsPhysicalDevice sets IsPhysicalDevice field to given value.

### HasIsPhysicalDevice

`func (o *DeviceDto) HasIsPhysicalDevice() bool`

HasIsPhysicalDevice returns a boolean if a field has been set.

### GetName

`func (o *DeviceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceDto) GetPlatform() AppPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceDto) GetPlatformOk() (*AppPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceDto) SetPlatform(v AppPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceDto) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetBrand

`func (o *DeviceDto) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *DeviceDto) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *DeviceDto) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *DeviceDto) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetSystemVersion

`func (o *DeviceDto) GetSystemVersion() string`

GetSystemVersion returns the SystemVersion field if non-nil, zero value otherwise.

### GetSystemVersionOk

`func (o *DeviceDto) GetSystemVersionOk() (*string, bool)`

GetSystemVersionOk returns a tuple with the SystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemVersion

`func (o *DeviceDto) SetSystemVersion(v string)`

SetSystemVersion sets SystemVersion field to given value.

### HasSystemVersion

`func (o *DeviceDto) HasSystemVersion() bool`

HasSystemVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



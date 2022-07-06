# RefreshDeviceStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **map[string]interface{}** | None, Unknow, Online, Offline | [optional] 

## Methods

### NewRefreshDeviceStatusDto

`func NewRefreshDeviceStatusDto() *RefreshDeviceStatusDto`

NewRefreshDeviceStatusDto instantiates a new RefreshDeviceStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshDeviceStatusDtoWithDefaults

`func NewRefreshDeviceStatusDtoWithDefaults() *RefreshDeviceStatusDto`

NewRefreshDeviceStatusDtoWithDefaults instantiates a new RefreshDeviceStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RefreshDeviceStatusDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RefreshDeviceStatusDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RefreshDeviceStatusDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RefreshDeviceStatusDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RefreshDeviceStatusDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RefreshDeviceStatusDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStatus

`func (o *RefreshDeviceStatusDto) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RefreshDeviceStatusDto) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RefreshDeviceStatusDto) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RefreshDeviceStatusDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



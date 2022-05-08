# MemberDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**Level** | Pointer to [**MemberLevel**](MemberLevel.md) |  | [optional] 

## Methods

### NewMemberDto

`func NewMemberDto() *MemberDto`

NewMemberDto instantiates a new MemberDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDtoWithDefaults

`func NewMemberDtoWithDefaults() *MemberDto`

NewMemberDtoWithDefaults instantiates a new MemberDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiredAt

`func (o *MemberDto) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *MemberDto) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *MemberDto) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *MemberDto) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *MemberDto) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *MemberDto) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetLevel

`func (o *MemberDto) GetLevel() MemberLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MemberDto) GetLevelOk() (*MemberLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MemberDto) SetLevel(v MemberLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *MemberDto) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



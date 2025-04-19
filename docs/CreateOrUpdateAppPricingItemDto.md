# CreateOrUpdateAppPricingItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | 名称: 坐席 | [optional] 
**Description** | Pointer to **string** | 描述, 使用 Markdown 格式, 允许包含图片 | [optional] 
**LinkUrl** | Pointer to **string** | 链接地址 | [optional] 
**Display** | Pointer to **string** | 显示模板: 包括{0}个坐席 | [optional] 
**SortIndex** | Pointer to **int32** | 排序 | [optional] 

## Methods

### NewCreateOrUpdateAppPricingItemDto

`func NewCreateOrUpdateAppPricingItemDto() *CreateOrUpdateAppPricingItemDto`

NewCreateOrUpdateAppPricingItemDto instantiates a new CreateOrUpdateAppPricingItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateAppPricingItemDtoWithDefaults

`func NewCreateOrUpdateAppPricingItemDtoWithDefaults() *CreateOrUpdateAppPricingItemDto`

NewCreateOrUpdateAppPricingItemDtoWithDefaults instantiates a new CreateOrUpdateAppPricingItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrUpdateAppPricingItemDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateAppPricingItemDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateAppPricingItemDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrUpdateAppPricingItemDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrUpdateAppPricingItemDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateAppPricingItemDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateAppPricingItemDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateAppPricingItemDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLinkUrl

`func (o *CreateOrUpdateAppPricingItemDto) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *CreateOrUpdateAppPricingItemDto) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *CreateOrUpdateAppPricingItemDto) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *CreateOrUpdateAppPricingItemDto) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *CreateOrUpdateAppPricingItemDto) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CreateOrUpdateAppPricingItemDto) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CreateOrUpdateAppPricingItemDto) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CreateOrUpdateAppPricingItemDto) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetSortIndex

`func (o *CreateOrUpdateAppPricingItemDto) GetSortIndex() int32`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *CreateOrUpdateAppPricingItemDto) GetSortIndexOk() (*int32, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *CreateOrUpdateAppPricingItemDto) SetSortIndex(v int32)`

SetSortIndex sets SortIndex field to given value.

### HasSortIndex

`func (o *CreateOrUpdateAppPricingItemDto) HasSortIndex() bool`

HasSortIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



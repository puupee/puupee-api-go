# CreateOrUpdateAppPricingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Naming** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**MonthPrice** | Pointer to **float64** |  | [optional] 
**MonthDiscount** | Pointer to **float64** |  | [optional] 
**MonthDiscountPrice** | Pointer to **float64** |  | [optional] 
**MonthDiscountStartAt** | Pointer to **time.Time** |  | [optional] 
**MonthDiscountEndAt** | Pointer to **time.Time** |  | [optional] 
**YearPrice** | Pointer to **float64** |  | [optional] 
**YearDiscount** | Pointer to **float64** |  | [optional] 
**YearDiscountPrice** | Pointer to **float64** |  | [optional] 
**YearDiscountStartAt** | Pointer to **time.Time** |  | [optional] 
**YearDiscountEndAt** | Pointer to **time.Time** |  | [optional] 
**Items** | Pointer to [**[]AppPricingItemDto**](AppPricingItemDto.md) |  | [optional] 

## Methods

### NewCreateOrUpdateAppPricingDto

`func NewCreateOrUpdateAppPricingDto() *CreateOrUpdateAppPricingDto`

NewCreateOrUpdateAppPricingDto instantiates a new CreateOrUpdateAppPricingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateAppPricingDtoWithDefaults

`func NewCreateOrUpdateAppPricingDtoWithDefaults() *CreateOrUpdateAppPricingDto`

NewCreateOrUpdateAppPricingDtoWithDefaults instantiates a new CreateOrUpdateAppPricingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNaming

`func (o *CreateOrUpdateAppPricingDto) GetNaming() string`

GetNaming returns the Naming field if non-nil, zero value otherwise.

### GetNamingOk

`func (o *CreateOrUpdateAppPricingDto) GetNamingOk() (*string, bool)`

GetNamingOk returns a tuple with the Naming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaming

`func (o *CreateOrUpdateAppPricingDto) SetNaming(v string)`

SetNaming sets Naming field to given value.

### HasNaming

`func (o *CreateOrUpdateAppPricingDto) HasNaming() bool`

HasNaming returns a boolean if a field has been set.

### GetProductId

`func (o *CreateOrUpdateAppPricingDto) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateOrUpdateAppPricingDto) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateOrUpdateAppPricingDto) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CreateOrUpdateAppPricingDto) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrUpdateAppPricingDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateAppPricingDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateAppPricingDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateAppPricingDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAppId

`func (o *CreateOrUpdateAppPricingDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateOrUpdateAppPricingDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateOrUpdateAppPricingDto) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CreateOrUpdateAppPricingDto) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetMonthPrice

`func (o *CreateOrUpdateAppPricingDto) GetMonthPrice() float64`

GetMonthPrice returns the MonthPrice field if non-nil, zero value otherwise.

### GetMonthPriceOk

`func (o *CreateOrUpdateAppPricingDto) GetMonthPriceOk() (*float64, bool)`

GetMonthPriceOk returns a tuple with the MonthPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthPrice

`func (o *CreateOrUpdateAppPricingDto) SetMonthPrice(v float64)`

SetMonthPrice sets MonthPrice field to given value.

### HasMonthPrice

`func (o *CreateOrUpdateAppPricingDto) HasMonthPrice() bool`

HasMonthPrice returns a boolean if a field has been set.

### GetMonthDiscount

`func (o *CreateOrUpdateAppPricingDto) GetMonthDiscount() float64`

GetMonthDiscount returns the MonthDiscount field if non-nil, zero value otherwise.

### GetMonthDiscountOk

`func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountOk() (*float64, bool)`

GetMonthDiscountOk returns a tuple with the MonthDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDiscount

`func (o *CreateOrUpdateAppPricingDto) SetMonthDiscount(v float64)`

SetMonthDiscount sets MonthDiscount field to given value.

### HasMonthDiscount

`func (o *CreateOrUpdateAppPricingDto) HasMonthDiscount() bool`

HasMonthDiscount returns a boolean if a field has been set.

### GetMonthDiscountPrice

`func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountPrice() float64`

GetMonthDiscountPrice returns the MonthDiscountPrice field if non-nil, zero value otherwise.

### GetMonthDiscountPriceOk

`func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountPriceOk() (*float64, bool)`

GetMonthDiscountPriceOk returns a tuple with the MonthDiscountPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDiscountPrice

`func (o *CreateOrUpdateAppPricingDto) SetMonthDiscountPrice(v float64)`

SetMonthDiscountPrice sets MonthDiscountPrice field to given value.

### HasMonthDiscountPrice

`func (o *CreateOrUpdateAppPricingDto) HasMonthDiscountPrice() bool`

HasMonthDiscountPrice returns a boolean if a field has been set.

### GetMonthDiscountStartAt

`func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountStartAt() time.Time`

GetMonthDiscountStartAt returns the MonthDiscountStartAt field if non-nil, zero value otherwise.

### GetMonthDiscountStartAtOk

`func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountStartAtOk() (*time.Time, bool)`

GetMonthDiscountStartAtOk returns a tuple with the MonthDiscountStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDiscountStartAt

`func (o *CreateOrUpdateAppPricingDto) SetMonthDiscountStartAt(v time.Time)`

SetMonthDiscountStartAt sets MonthDiscountStartAt field to given value.

### HasMonthDiscountStartAt

`func (o *CreateOrUpdateAppPricingDto) HasMonthDiscountStartAt() bool`

HasMonthDiscountStartAt returns a boolean if a field has been set.

### GetMonthDiscountEndAt

`func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountEndAt() time.Time`

GetMonthDiscountEndAt returns the MonthDiscountEndAt field if non-nil, zero value otherwise.

### GetMonthDiscountEndAtOk

`func (o *CreateOrUpdateAppPricingDto) GetMonthDiscountEndAtOk() (*time.Time, bool)`

GetMonthDiscountEndAtOk returns a tuple with the MonthDiscountEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDiscountEndAt

`func (o *CreateOrUpdateAppPricingDto) SetMonthDiscountEndAt(v time.Time)`

SetMonthDiscountEndAt sets MonthDiscountEndAt field to given value.

### HasMonthDiscountEndAt

`func (o *CreateOrUpdateAppPricingDto) HasMonthDiscountEndAt() bool`

HasMonthDiscountEndAt returns a boolean if a field has been set.

### GetYearPrice

`func (o *CreateOrUpdateAppPricingDto) GetYearPrice() float64`

GetYearPrice returns the YearPrice field if non-nil, zero value otherwise.

### GetYearPriceOk

`func (o *CreateOrUpdateAppPricingDto) GetYearPriceOk() (*float64, bool)`

GetYearPriceOk returns a tuple with the YearPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearPrice

`func (o *CreateOrUpdateAppPricingDto) SetYearPrice(v float64)`

SetYearPrice sets YearPrice field to given value.

### HasYearPrice

`func (o *CreateOrUpdateAppPricingDto) HasYearPrice() bool`

HasYearPrice returns a boolean if a field has been set.

### GetYearDiscount

`func (o *CreateOrUpdateAppPricingDto) GetYearDiscount() float64`

GetYearDiscount returns the YearDiscount field if non-nil, zero value otherwise.

### GetYearDiscountOk

`func (o *CreateOrUpdateAppPricingDto) GetYearDiscountOk() (*float64, bool)`

GetYearDiscountOk returns a tuple with the YearDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearDiscount

`func (o *CreateOrUpdateAppPricingDto) SetYearDiscount(v float64)`

SetYearDiscount sets YearDiscount field to given value.

### HasYearDiscount

`func (o *CreateOrUpdateAppPricingDto) HasYearDiscount() bool`

HasYearDiscount returns a boolean if a field has been set.

### GetYearDiscountPrice

`func (o *CreateOrUpdateAppPricingDto) GetYearDiscountPrice() float64`

GetYearDiscountPrice returns the YearDiscountPrice field if non-nil, zero value otherwise.

### GetYearDiscountPriceOk

`func (o *CreateOrUpdateAppPricingDto) GetYearDiscountPriceOk() (*float64, bool)`

GetYearDiscountPriceOk returns a tuple with the YearDiscountPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearDiscountPrice

`func (o *CreateOrUpdateAppPricingDto) SetYearDiscountPrice(v float64)`

SetYearDiscountPrice sets YearDiscountPrice field to given value.

### HasYearDiscountPrice

`func (o *CreateOrUpdateAppPricingDto) HasYearDiscountPrice() bool`

HasYearDiscountPrice returns a boolean if a field has been set.

### GetYearDiscountStartAt

`func (o *CreateOrUpdateAppPricingDto) GetYearDiscountStartAt() time.Time`

GetYearDiscountStartAt returns the YearDiscountStartAt field if non-nil, zero value otherwise.

### GetYearDiscountStartAtOk

`func (o *CreateOrUpdateAppPricingDto) GetYearDiscountStartAtOk() (*time.Time, bool)`

GetYearDiscountStartAtOk returns a tuple with the YearDiscountStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearDiscountStartAt

`func (o *CreateOrUpdateAppPricingDto) SetYearDiscountStartAt(v time.Time)`

SetYearDiscountStartAt sets YearDiscountStartAt field to given value.

### HasYearDiscountStartAt

`func (o *CreateOrUpdateAppPricingDto) HasYearDiscountStartAt() bool`

HasYearDiscountStartAt returns a boolean if a field has been set.

### GetYearDiscountEndAt

`func (o *CreateOrUpdateAppPricingDto) GetYearDiscountEndAt() time.Time`

GetYearDiscountEndAt returns the YearDiscountEndAt field if non-nil, zero value otherwise.

### GetYearDiscountEndAtOk

`func (o *CreateOrUpdateAppPricingDto) GetYearDiscountEndAtOk() (*time.Time, bool)`

GetYearDiscountEndAtOk returns a tuple with the YearDiscountEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearDiscountEndAt

`func (o *CreateOrUpdateAppPricingDto) SetYearDiscountEndAt(v time.Time)`

SetYearDiscountEndAt sets YearDiscountEndAt field to given value.

### HasYearDiscountEndAt

`func (o *CreateOrUpdateAppPricingDto) HasYearDiscountEndAt() bool`

HasYearDiscountEndAt returns a boolean if a field has been set.

### GetItems

`func (o *CreateOrUpdateAppPricingDto) GetItems() []AppPricingItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateOrUpdateAppPricingDto) GetItemsOk() (*[]AppPricingItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateOrUpdateAppPricingDto) SetItems(v []AppPricingItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *CreateOrUpdateAppPricingDto) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



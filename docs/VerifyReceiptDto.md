# VerifyReceiptDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** |  | [optional] 
**ReceiptData** | Pointer to **string** |  | [optional] 

## Methods

### NewVerifyReceiptDto

`func NewVerifyReceiptDto() *VerifyReceiptDto`

NewVerifyReceiptDto instantiates a new VerifyReceiptDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyReceiptDtoWithDefaults

`func NewVerifyReceiptDtoWithDefaults() *VerifyReceiptDto`

NewVerifyReceiptDtoWithDefaults instantiates a new VerifyReceiptDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *VerifyReceiptDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *VerifyReceiptDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *VerifyReceiptDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *VerifyReceiptDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetReceiptData

`func (o *VerifyReceiptDto) GetReceiptData() string`

GetReceiptData returns the ReceiptData field if non-nil, zero value otherwise.

### GetReceiptDataOk

`func (o *VerifyReceiptDto) GetReceiptDataOk() (*string, bool)`

GetReceiptDataOk returns a tuple with the ReceiptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptData

`func (o *VerifyReceiptDto) SetReceiptData(v string)`

SetReceiptData sets ReceiptData field to given value.

### HasReceiptData

`func (o *VerifyReceiptDto) HasReceiptData() bool`

HasReceiptData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



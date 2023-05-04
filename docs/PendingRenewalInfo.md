# PendingRenewalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRenewProductId** | Pointer to **string** |  | [optional] 
**AutoRenewStatus** | Pointer to **string** |  | [optional] 
**IsInBillingRetryPeriod** | Pointer to **string** |  | [optional] 
**OriginalTransactionId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**ExpirationIntent** | Pointer to **string** |  | [optional] 
**PriceConsentStatus** | Pointer to **string** |  | [optional] 
**GracePeriodExpiresDate** | Pointer to **string** |  | [optional] 
**GracePeriodExpiresDateMs** | Pointer to **string** |  | [optional] 
**GracePeriodExpiresDatePst** | Pointer to **string** |  | [optional] 

## Methods

### NewPendingRenewalInfo

`func NewPendingRenewalInfo() *PendingRenewalInfo`

NewPendingRenewalInfo instantiates a new PendingRenewalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingRenewalInfoWithDefaults

`func NewPendingRenewalInfoWithDefaults() *PendingRenewalInfo`

NewPendingRenewalInfoWithDefaults instantiates a new PendingRenewalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRenewProductId

`func (o *PendingRenewalInfo) GetAutoRenewProductId() string`

GetAutoRenewProductId returns the AutoRenewProductId field if non-nil, zero value otherwise.

### GetAutoRenewProductIdOk

`func (o *PendingRenewalInfo) GetAutoRenewProductIdOk() (*string, bool)`

GetAutoRenewProductIdOk returns a tuple with the AutoRenewProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewProductId

`func (o *PendingRenewalInfo) SetAutoRenewProductId(v string)`

SetAutoRenewProductId sets AutoRenewProductId field to given value.

### HasAutoRenewProductId

`func (o *PendingRenewalInfo) HasAutoRenewProductId() bool`

HasAutoRenewProductId returns a boolean if a field has been set.

### GetAutoRenewStatus

`func (o *PendingRenewalInfo) GetAutoRenewStatus() string`

GetAutoRenewStatus returns the AutoRenewStatus field if non-nil, zero value otherwise.

### GetAutoRenewStatusOk

`func (o *PendingRenewalInfo) GetAutoRenewStatusOk() (*string, bool)`

GetAutoRenewStatusOk returns a tuple with the AutoRenewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewStatus

`func (o *PendingRenewalInfo) SetAutoRenewStatus(v string)`

SetAutoRenewStatus sets AutoRenewStatus field to given value.

### HasAutoRenewStatus

`func (o *PendingRenewalInfo) HasAutoRenewStatus() bool`

HasAutoRenewStatus returns a boolean if a field has been set.

### GetIsInBillingRetryPeriod

`func (o *PendingRenewalInfo) GetIsInBillingRetryPeriod() string`

GetIsInBillingRetryPeriod returns the IsInBillingRetryPeriod field if non-nil, zero value otherwise.

### GetIsInBillingRetryPeriodOk

`func (o *PendingRenewalInfo) GetIsInBillingRetryPeriodOk() (*string, bool)`

GetIsInBillingRetryPeriodOk returns a tuple with the IsInBillingRetryPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInBillingRetryPeriod

`func (o *PendingRenewalInfo) SetIsInBillingRetryPeriod(v string)`

SetIsInBillingRetryPeriod sets IsInBillingRetryPeriod field to given value.

### HasIsInBillingRetryPeriod

`func (o *PendingRenewalInfo) HasIsInBillingRetryPeriod() bool`

HasIsInBillingRetryPeriod returns a boolean if a field has been set.

### GetOriginalTransactionId

`func (o *PendingRenewalInfo) GetOriginalTransactionId() string`

GetOriginalTransactionId returns the OriginalTransactionId field if non-nil, zero value otherwise.

### GetOriginalTransactionIdOk

`func (o *PendingRenewalInfo) GetOriginalTransactionIdOk() (*string, bool)`

GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTransactionId

`func (o *PendingRenewalInfo) SetOriginalTransactionId(v string)`

SetOriginalTransactionId sets OriginalTransactionId field to given value.

### HasOriginalTransactionId

`func (o *PendingRenewalInfo) HasOriginalTransactionId() bool`

HasOriginalTransactionId returns a boolean if a field has been set.

### GetProductId

`func (o *PendingRenewalInfo) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PendingRenewalInfo) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PendingRenewalInfo) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PendingRenewalInfo) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetExpirationIntent

`func (o *PendingRenewalInfo) GetExpirationIntent() string`

GetExpirationIntent returns the ExpirationIntent field if non-nil, zero value otherwise.

### GetExpirationIntentOk

`func (o *PendingRenewalInfo) GetExpirationIntentOk() (*string, bool)`

GetExpirationIntentOk returns a tuple with the ExpirationIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationIntent

`func (o *PendingRenewalInfo) SetExpirationIntent(v string)`

SetExpirationIntent sets ExpirationIntent field to given value.

### HasExpirationIntent

`func (o *PendingRenewalInfo) HasExpirationIntent() bool`

HasExpirationIntent returns a boolean if a field has been set.

### GetPriceConsentStatus

`func (o *PendingRenewalInfo) GetPriceConsentStatus() string`

GetPriceConsentStatus returns the PriceConsentStatus field if non-nil, zero value otherwise.

### GetPriceConsentStatusOk

`func (o *PendingRenewalInfo) GetPriceConsentStatusOk() (*string, bool)`

GetPriceConsentStatusOk returns a tuple with the PriceConsentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceConsentStatus

`func (o *PendingRenewalInfo) SetPriceConsentStatus(v string)`

SetPriceConsentStatus sets PriceConsentStatus field to given value.

### HasPriceConsentStatus

`func (o *PendingRenewalInfo) HasPriceConsentStatus() bool`

HasPriceConsentStatus returns a boolean if a field has been set.

### GetGracePeriodExpiresDate

`func (o *PendingRenewalInfo) GetGracePeriodExpiresDate() string`

GetGracePeriodExpiresDate returns the GracePeriodExpiresDate field if non-nil, zero value otherwise.

### GetGracePeriodExpiresDateOk

`func (o *PendingRenewalInfo) GetGracePeriodExpiresDateOk() (*string, bool)`

GetGracePeriodExpiresDateOk returns a tuple with the GracePeriodExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodExpiresDate

`func (o *PendingRenewalInfo) SetGracePeriodExpiresDate(v string)`

SetGracePeriodExpiresDate sets GracePeriodExpiresDate field to given value.

### HasGracePeriodExpiresDate

`func (o *PendingRenewalInfo) HasGracePeriodExpiresDate() bool`

HasGracePeriodExpiresDate returns a boolean if a field has been set.

### GetGracePeriodExpiresDateMs

`func (o *PendingRenewalInfo) GetGracePeriodExpiresDateMs() string`

GetGracePeriodExpiresDateMs returns the GracePeriodExpiresDateMs field if non-nil, zero value otherwise.

### GetGracePeriodExpiresDateMsOk

`func (o *PendingRenewalInfo) GetGracePeriodExpiresDateMsOk() (*string, bool)`

GetGracePeriodExpiresDateMsOk returns a tuple with the GracePeriodExpiresDateMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodExpiresDateMs

`func (o *PendingRenewalInfo) SetGracePeriodExpiresDateMs(v string)`

SetGracePeriodExpiresDateMs sets GracePeriodExpiresDateMs field to given value.

### HasGracePeriodExpiresDateMs

`func (o *PendingRenewalInfo) HasGracePeriodExpiresDateMs() bool`

HasGracePeriodExpiresDateMs returns a boolean if a field has been set.

### GetGracePeriodExpiresDatePst

`func (o *PendingRenewalInfo) GetGracePeriodExpiresDatePst() string`

GetGracePeriodExpiresDatePst returns the GracePeriodExpiresDatePst field if non-nil, zero value otherwise.

### GetGracePeriodExpiresDatePstOk

`func (o *PendingRenewalInfo) GetGracePeriodExpiresDatePstOk() (*string, bool)`

GetGracePeriodExpiresDatePstOk returns a tuple with the GracePeriodExpiresDatePst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodExpiresDatePst

`func (o *PendingRenewalInfo) SetGracePeriodExpiresDatePst(v string)`

SetGracePeriodExpiresDatePst sets GracePeriodExpiresDatePst field to given value.

### HasGracePeriodExpiresDatePst

`func (o *PendingRenewalInfo) HasGracePeriodExpiresDatePst() bool`

HasGracePeriodExpiresDatePst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



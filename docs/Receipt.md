# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceiptType** | Pointer to **string** |  | [optional] 
**AdamId** | Pointer to **int32** |  | [optional] 
**AppItemId** | Pointer to **int32** |  | [optional] 
**BundleId** | Pointer to **string** |  | [optional] 
**ApplicationVersion** | Pointer to **string** |  | [optional] 
**DownloadId** | Pointer to **int32** |  | [optional] 
**VersionExternalIdentifier** | Pointer to **int32** |  | [optional] 
**ReceiptCreationDate** | Pointer to **string** |  | [optional] 
**ReceiptCreationDateMs** | Pointer to **string** |  | [optional] 
**ReceiptCreationDatePst** | Pointer to **string** |  | [optional] 
**RequestDate** | Pointer to **string** |  | [optional] 
**RequestDateMs** | Pointer to **string** |  | [optional] 
**RequestDatePst** | Pointer to **string** |  | [optional] 
**OriginalPurchaseDate** | Pointer to **string** |  | [optional] 
**OriginalPurchaseDateMs** | Pointer to **string** |  | [optional] 
**OriginalPurchaseDatePst** | Pointer to **string** |  | [optional] 
**OriginalApplicationVersion** | Pointer to **string** |  | [optional] 
**InApp** | Pointer to [**[]InApp**](InApp.md) |  | [optional] 

## Methods

### NewReceipt

`func NewReceipt() *Receipt`

NewReceipt instantiates a new Receipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptWithDefaults

`func NewReceiptWithDefaults() *Receipt`

NewReceiptWithDefaults instantiates a new Receipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiptType

`func (o *Receipt) GetReceiptType() string`

GetReceiptType returns the ReceiptType field if non-nil, zero value otherwise.

### GetReceiptTypeOk

`func (o *Receipt) GetReceiptTypeOk() (*string, bool)`

GetReceiptTypeOk returns a tuple with the ReceiptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptType

`func (o *Receipt) SetReceiptType(v string)`

SetReceiptType sets ReceiptType field to given value.

### HasReceiptType

`func (o *Receipt) HasReceiptType() bool`

HasReceiptType returns a boolean if a field has been set.

### GetAdamId

`func (o *Receipt) GetAdamId() int32`

GetAdamId returns the AdamId field if non-nil, zero value otherwise.

### GetAdamIdOk

`func (o *Receipt) GetAdamIdOk() (*int32, bool)`

GetAdamIdOk returns a tuple with the AdamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdamId

`func (o *Receipt) SetAdamId(v int32)`

SetAdamId sets AdamId field to given value.

### HasAdamId

`func (o *Receipt) HasAdamId() bool`

HasAdamId returns a boolean if a field has been set.

### GetAppItemId

`func (o *Receipt) GetAppItemId() int32`

GetAppItemId returns the AppItemId field if non-nil, zero value otherwise.

### GetAppItemIdOk

`func (o *Receipt) GetAppItemIdOk() (*int32, bool)`

GetAppItemIdOk returns a tuple with the AppItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppItemId

`func (o *Receipt) SetAppItemId(v int32)`

SetAppItemId sets AppItemId field to given value.

### HasAppItemId

`func (o *Receipt) HasAppItemId() bool`

HasAppItemId returns a boolean if a field has been set.

### GetBundleId

`func (o *Receipt) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *Receipt) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *Receipt) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *Receipt) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetApplicationVersion

`func (o *Receipt) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *Receipt) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *Receipt) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *Receipt) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### GetDownloadId

`func (o *Receipt) GetDownloadId() int32`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *Receipt) GetDownloadIdOk() (*int32, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *Receipt) SetDownloadId(v int32)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *Receipt) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### GetVersionExternalIdentifier

`func (o *Receipt) GetVersionExternalIdentifier() int32`

GetVersionExternalIdentifier returns the VersionExternalIdentifier field if non-nil, zero value otherwise.

### GetVersionExternalIdentifierOk

`func (o *Receipt) GetVersionExternalIdentifierOk() (*int32, bool)`

GetVersionExternalIdentifierOk returns a tuple with the VersionExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionExternalIdentifier

`func (o *Receipt) SetVersionExternalIdentifier(v int32)`

SetVersionExternalIdentifier sets VersionExternalIdentifier field to given value.

### HasVersionExternalIdentifier

`func (o *Receipt) HasVersionExternalIdentifier() bool`

HasVersionExternalIdentifier returns a boolean if a field has been set.

### GetReceiptCreationDate

`func (o *Receipt) GetReceiptCreationDate() string`

GetReceiptCreationDate returns the ReceiptCreationDate field if non-nil, zero value otherwise.

### GetReceiptCreationDateOk

`func (o *Receipt) GetReceiptCreationDateOk() (*string, bool)`

GetReceiptCreationDateOk returns a tuple with the ReceiptCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptCreationDate

`func (o *Receipt) SetReceiptCreationDate(v string)`

SetReceiptCreationDate sets ReceiptCreationDate field to given value.

### HasReceiptCreationDate

`func (o *Receipt) HasReceiptCreationDate() bool`

HasReceiptCreationDate returns a boolean if a field has been set.

### GetReceiptCreationDateMs

`func (o *Receipt) GetReceiptCreationDateMs() string`

GetReceiptCreationDateMs returns the ReceiptCreationDateMs field if non-nil, zero value otherwise.

### GetReceiptCreationDateMsOk

`func (o *Receipt) GetReceiptCreationDateMsOk() (*string, bool)`

GetReceiptCreationDateMsOk returns a tuple with the ReceiptCreationDateMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptCreationDateMs

`func (o *Receipt) SetReceiptCreationDateMs(v string)`

SetReceiptCreationDateMs sets ReceiptCreationDateMs field to given value.

### HasReceiptCreationDateMs

`func (o *Receipt) HasReceiptCreationDateMs() bool`

HasReceiptCreationDateMs returns a boolean if a field has been set.

### GetReceiptCreationDatePst

`func (o *Receipt) GetReceiptCreationDatePst() string`

GetReceiptCreationDatePst returns the ReceiptCreationDatePst field if non-nil, zero value otherwise.

### GetReceiptCreationDatePstOk

`func (o *Receipt) GetReceiptCreationDatePstOk() (*string, bool)`

GetReceiptCreationDatePstOk returns a tuple with the ReceiptCreationDatePst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptCreationDatePst

`func (o *Receipt) SetReceiptCreationDatePst(v string)`

SetReceiptCreationDatePst sets ReceiptCreationDatePst field to given value.

### HasReceiptCreationDatePst

`func (o *Receipt) HasReceiptCreationDatePst() bool`

HasReceiptCreationDatePst returns a boolean if a field has been set.

### GetRequestDate

`func (o *Receipt) GetRequestDate() string`

GetRequestDate returns the RequestDate field if non-nil, zero value otherwise.

### GetRequestDateOk

`func (o *Receipt) GetRequestDateOk() (*string, bool)`

GetRequestDateOk returns a tuple with the RequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDate

`func (o *Receipt) SetRequestDate(v string)`

SetRequestDate sets RequestDate field to given value.

### HasRequestDate

`func (o *Receipt) HasRequestDate() bool`

HasRequestDate returns a boolean if a field has been set.

### GetRequestDateMs

`func (o *Receipt) GetRequestDateMs() string`

GetRequestDateMs returns the RequestDateMs field if non-nil, zero value otherwise.

### GetRequestDateMsOk

`func (o *Receipt) GetRequestDateMsOk() (*string, bool)`

GetRequestDateMsOk returns a tuple with the RequestDateMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDateMs

`func (o *Receipt) SetRequestDateMs(v string)`

SetRequestDateMs sets RequestDateMs field to given value.

### HasRequestDateMs

`func (o *Receipt) HasRequestDateMs() bool`

HasRequestDateMs returns a boolean if a field has been set.

### GetRequestDatePst

`func (o *Receipt) GetRequestDatePst() string`

GetRequestDatePst returns the RequestDatePst field if non-nil, zero value otherwise.

### GetRequestDatePstOk

`func (o *Receipt) GetRequestDatePstOk() (*string, bool)`

GetRequestDatePstOk returns a tuple with the RequestDatePst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDatePst

`func (o *Receipt) SetRequestDatePst(v string)`

SetRequestDatePst sets RequestDatePst field to given value.

### HasRequestDatePst

`func (o *Receipt) HasRequestDatePst() bool`

HasRequestDatePst returns a boolean if a field has been set.

### GetOriginalPurchaseDate

`func (o *Receipt) GetOriginalPurchaseDate() string`

GetOriginalPurchaseDate returns the OriginalPurchaseDate field if non-nil, zero value otherwise.

### GetOriginalPurchaseDateOk

`func (o *Receipt) GetOriginalPurchaseDateOk() (*string, bool)`

GetOriginalPurchaseDateOk returns a tuple with the OriginalPurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPurchaseDate

`func (o *Receipt) SetOriginalPurchaseDate(v string)`

SetOriginalPurchaseDate sets OriginalPurchaseDate field to given value.

### HasOriginalPurchaseDate

`func (o *Receipt) HasOriginalPurchaseDate() bool`

HasOriginalPurchaseDate returns a boolean if a field has been set.

### GetOriginalPurchaseDateMs

`func (o *Receipt) GetOriginalPurchaseDateMs() string`

GetOriginalPurchaseDateMs returns the OriginalPurchaseDateMs field if non-nil, zero value otherwise.

### GetOriginalPurchaseDateMsOk

`func (o *Receipt) GetOriginalPurchaseDateMsOk() (*string, bool)`

GetOriginalPurchaseDateMsOk returns a tuple with the OriginalPurchaseDateMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPurchaseDateMs

`func (o *Receipt) SetOriginalPurchaseDateMs(v string)`

SetOriginalPurchaseDateMs sets OriginalPurchaseDateMs field to given value.

### HasOriginalPurchaseDateMs

`func (o *Receipt) HasOriginalPurchaseDateMs() bool`

HasOriginalPurchaseDateMs returns a boolean if a field has been set.

### GetOriginalPurchaseDatePst

`func (o *Receipt) GetOriginalPurchaseDatePst() string`

GetOriginalPurchaseDatePst returns the OriginalPurchaseDatePst field if non-nil, zero value otherwise.

### GetOriginalPurchaseDatePstOk

`func (o *Receipt) GetOriginalPurchaseDatePstOk() (*string, bool)`

GetOriginalPurchaseDatePstOk returns a tuple with the OriginalPurchaseDatePst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPurchaseDatePst

`func (o *Receipt) SetOriginalPurchaseDatePst(v string)`

SetOriginalPurchaseDatePst sets OriginalPurchaseDatePst field to given value.

### HasOriginalPurchaseDatePst

`func (o *Receipt) HasOriginalPurchaseDatePst() bool`

HasOriginalPurchaseDatePst returns a boolean if a field has been set.

### GetOriginalApplicationVersion

`func (o *Receipt) GetOriginalApplicationVersion() string`

GetOriginalApplicationVersion returns the OriginalApplicationVersion field if non-nil, zero value otherwise.

### GetOriginalApplicationVersionOk

`func (o *Receipt) GetOriginalApplicationVersionOk() (*string, bool)`

GetOriginalApplicationVersionOk returns a tuple with the OriginalApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalApplicationVersion

`func (o *Receipt) SetOriginalApplicationVersion(v string)`

SetOriginalApplicationVersion sets OriginalApplicationVersion field to given value.

### HasOriginalApplicationVersion

`func (o *Receipt) HasOriginalApplicationVersion() bool`

HasOriginalApplicationVersion returns a boolean if a field has been set.

### GetInApp

`func (o *Receipt) GetInApp() []InApp`

GetInApp returns the InApp field if non-nil, zero value otherwise.

### GetInAppOk

`func (o *Receipt) GetInAppOk() (*[]InApp, bool)`

GetInAppOk returns a tuple with the InApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInApp

`func (o *Receipt) SetInApp(v []InApp)`

SetInApp sets InApp field to given value.

### HasInApp

`func (o *Receipt) HasInApp() bool`

HasInApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



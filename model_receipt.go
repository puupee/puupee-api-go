/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// Receipt struct for Receipt
type Receipt struct {
	ReceiptType *string `json:"receipt_type,omitempty"`
	AdamId *int32 `json:"adam_id,omitempty"`
	AppItemId *int32 `json:"app_item_id,omitempty"`
	BundleId *string `json:"bundle_id,omitempty"`
	ApplicationVersion *string `json:"application_version,omitempty"`
	DownloadId *int32 `json:"download_id,omitempty"`
	VersionExternalIdentifier *int32 `json:"version_external_identifier,omitempty"`
	ReceiptCreationDate *string `json:"receipt_creation_date,omitempty"`
	ReceiptCreationDateMs *string `json:"receipt_creation_date_ms,omitempty"`
	ReceiptCreationDatePst *string `json:"receipt_creation_date_pst,omitempty"`
	RequestDate *string `json:"request_date,omitempty"`
	RequestDateMs *string `json:"request_date_ms,omitempty"`
	RequestDatePst *string `json:"request_date_pst,omitempty"`
	OriginalPurchaseDate *string `json:"original_purchase_date,omitempty"`
	OriginalPurchaseDateMs *string `json:"original_purchase_date_ms,omitempty"`
	OriginalPurchaseDatePst *string `json:"original_purchase_date_pst,omitempty"`
	OriginalApplicationVersion *string `json:"original_application_version,omitempty"`
	InApp []InApp `json:"in_app,omitempty"`
}

// NewReceipt instantiates a new Receipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceipt() *Receipt {
	this := Receipt{}
	return &this
}

// NewReceiptWithDefaults instantiates a new Receipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptWithDefaults() *Receipt {
	this := Receipt{}
	return &this
}

// GetReceiptType returns the ReceiptType field value if set, zero value otherwise.
func (o *Receipt) GetReceiptType() string {
	if o == nil || isNil(o.ReceiptType) {
		var ret string
		return ret
	}
	return *o.ReceiptType
}

// GetReceiptTypeOk returns a tuple with the ReceiptType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetReceiptTypeOk() (*string, bool) {
	if o == nil || isNil(o.ReceiptType) {
    return nil, false
	}
	return o.ReceiptType, true
}

// HasReceiptType returns a boolean if a field has been set.
func (o *Receipt) HasReceiptType() bool {
	if o != nil && !isNil(o.ReceiptType) {
		return true
	}

	return false
}

// SetReceiptType gets a reference to the given string and assigns it to the ReceiptType field.
func (o *Receipt) SetReceiptType(v string) {
	o.ReceiptType = &v
}

// GetAdamId returns the AdamId field value if set, zero value otherwise.
func (o *Receipt) GetAdamId() int32 {
	if o == nil || isNil(o.AdamId) {
		var ret int32
		return ret
	}
	return *o.AdamId
}

// GetAdamIdOk returns a tuple with the AdamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetAdamIdOk() (*int32, bool) {
	if o == nil || isNil(o.AdamId) {
    return nil, false
	}
	return o.AdamId, true
}

// HasAdamId returns a boolean if a field has been set.
func (o *Receipt) HasAdamId() bool {
	if o != nil && !isNil(o.AdamId) {
		return true
	}

	return false
}

// SetAdamId gets a reference to the given int32 and assigns it to the AdamId field.
func (o *Receipt) SetAdamId(v int32) {
	o.AdamId = &v
}

// GetAppItemId returns the AppItemId field value if set, zero value otherwise.
func (o *Receipt) GetAppItemId() int32 {
	if o == nil || isNil(o.AppItemId) {
		var ret int32
		return ret
	}
	return *o.AppItemId
}

// GetAppItemIdOk returns a tuple with the AppItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetAppItemIdOk() (*int32, bool) {
	if o == nil || isNil(o.AppItemId) {
    return nil, false
	}
	return o.AppItemId, true
}

// HasAppItemId returns a boolean if a field has been set.
func (o *Receipt) HasAppItemId() bool {
	if o != nil && !isNil(o.AppItemId) {
		return true
	}

	return false
}

// SetAppItemId gets a reference to the given int32 and assigns it to the AppItemId field.
func (o *Receipt) SetAppItemId(v int32) {
	o.AppItemId = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *Receipt) GetBundleId() string {
	if o == nil || isNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetBundleIdOk() (*string, bool) {
	if o == nil || isNil(o.BundleId) {
    return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *Receipt) HasBundleId() bool {
	if o != nil && !isNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *Receipt) SetBundleId(v string) {
	o.BundleId = &v
}

// GetApplicationVersion returns the ApplicationVersion field value if set, zero value otherwise.
func (o *Receipt) GetApplicationVersion() string {
	if o == nil || isNil(o.ApplicationVersion) {
		var ret string
		return ret
	}
	return *o.ApplicationVersion
}

// GetApplicationVersionOk returns a tuple with the ApplicationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetApplicationVersionOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationVersion) {
    return nil, false
	}
	return o.ApplicationVersion, true
}

// HasApplicationVersion returns a boolean if a field has been set.
func (o *Receipt) HasApplicationVersion() bool {
	if o != nil && !isNil(o.ApplicationVersion) {
		return true
	}

	return false
}

// SetApplicationVersion gets a reference to the given string and assigns it to the ApplicationVersion field.
func (o *Receipt) SetApplicationVersion(v string) {
	o.ApplicationVersion = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise.
func (o *Receipt) GetDownloadId() int32 {
	if o == nil || isNil(o.DownloadId) {
		var ret int32
		return ret
	}
	return *o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetDownloadIdOk() (*int32, bool) {
	if o == nil || isNil(o.DownloadId) {
    return nil, false
	}
	return o.DownloadId, true
}

// HasDownloadId returns a boolean if a field has been set.
func (o *Receipt) HasDownloadId() bool {
	if o != nil && !isNil(o.DownloadId) {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given int32 and assigns it to the DownloadId field.
func (o *Receipt) SetDownloadId(v int32) {
	o.DownloadId = &v
}

// GetVersionExternalIdentifier returns the VersionExternalIdentifier field value if set, zero value otherwise.
func (o *Receipt) GetVersionExternalIdentifier() int32 {
	if o == nil || isNil(o.VersionExternalIdentifier) {
		var ret int32
		return ret
	}
	return *o.VersionExternalIdentifier
}

// GetVersionExternalIdentifierOk returns a tuple with the VersionExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetVersionExternalIdentifierOk() (*int32, bool) {
	if o == nil || isNil(o.VersionExternalIdentifier) {
    return nil, false
	}
	return o.VersionExternalIdentifier, true
}

// HasVersionExternalIdentifier returns a boolean if a field has been set.
func (o *Receipt) HasVersionExternalIdentifier() bool {
	if o != nil && !isNil(o.VersionExternalIdentifier) {
		return true
	}

	return false
}

// SetVersionExternalIdentifier gets a reference to the given int32 and assigns it to the VersionExternalIdentifier field.
func (o *Receipt) SetVersionExternalIdentifier(v int32) {
	o.VersionExternalIdentifier = &v
}

// GetReceiptCreationDate returns the ReceiptCreationDate field value if set, zero value otherwise.
func (o *Receipt) GetReceiptCreationDate() string {
	if o == nil || isNil(o.ReceiptCreationDate) {
		var ret string
		return ret
	}
	return *o.ReceiptCreationDate
}

// GetReceiptCreationDateOk returns a tuple with the ReceiptCreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetReceiptCreationDateOk() (*string, bool) {
	if o == nil || isNil(o.ReceiptCreationDate) {
    return nil, false
	}
	return o.ReceiptCreationDate, true
}

// HasReceiptCreationDate returns a boolean if a field has been set.
func (o *Receipt) HasReceiptCreationDate() bool {
	if o != nil && !isNil(o.ReceiptCreationDate) {
		return true
	}

	return false
}

// SetReceiptCreationDate gets a reference to the given string and assigns it to the ReceiptCreationDate field.
func (o *Receipt) SetReceiptCreationDate(v string) {
	o.ReceiptCreationDate = &v
}

// GetReceiptCreationDateMs returns the ReceiptCreationDateMs field value if set, zero value otherwise.
func (o *Receipt) GetReceiptCreationDateMs() string {
	if o == nil || isNil(o.ReceiptCreationDateMs) {
		var ret string
		return ret
	}
	return *o.ReceiptCreationDateMs
}

// GetReceiptCreationDateMsOk returns a tuple with the ReceiptCreationDateMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetReceiptCreationDateMsOk() (*string, bool) {
	if o == nil || isNil(o.ReceiptCreationDateMs) {
    return nil, false
	}
	return o.ReceiptCreationDateMs, true
}

// HasReceiptCreationDateMs returns a boolean if a field has been set.
func (o *Receipt) HasReceiptCreationDateMs() bool {
	if o != nil && !isNil(o.ReceiptCreationDateMs) {
		return true
	}

	return false
}

// SetReceiptCreationDateMs gets a reference to the given string and assigns it to the ReceiptCreationDateMs field.
func (o *Receipt) SetReceiptCreationDateMs(v string) {
	o.ReceiptCreationDateMs = &v
}

// GetReceiptCreationDatePst returns the ReceiptCreationDatePst field value if set, zero value otherwise.
func (o *Receipt) GetReceiptCreationDatePst() string {
	if o == nil || isNil(o.ReceiptCreationDatePst) {
		var ret string
		return ret
	}
	return *o.ReceiptCreationDatePst
}

// GetReceiptCreationDatePstOk returns a tuple with the ReceiptCreationDatePst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetReceiptCreationDatePstOk() (*string, bool) {
	if o == nil || isNil(o.ReceiptCreationDatePst) {
    return nil, false
	}
	return o.ReceiptCreationDatePst, true
}

// HasReceiptCreationDatePst returns a boolean if a field has been set.
func (o *Receipt) HasReceiptCreationDatePst() bool {
	if o != nil && !isNil(o.ReceiptCreationDatePst) {
		return true
	}

	return false
}

// SetReceiptCreationDatePst gets a reference to the given string and assigns it to the ReceiptCreationDatePst field.
func (o *Receipt) SetReceiptCreationDatePst(v string) {
	o.ReceiptCreationDatePst = &v
}

// GetRequestDate returns the RequestDate field value if set, zero value otherwise.
func (o *Receipt) GetRequestDate() string {
	if o == nil || isNil(o.RequestDate) {
		var ret string
		return ret
	}
	return *o.RequestDate
}

// GetRequestDateOk returns a tuple with the RequestDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetRequestDateOk() (*string, bool) {
	if o == nil || isNil(o.RequestDate) {
    return nil, false
	}
	return o.RequestDate, true
}

// HasRequestDate returns a boolean if a field has been set.
func (o *Receipt) HasRequestDate() bool {
	if o != nil && !isNil(o.RequestDate) {
		return true
	}

	return false
}

// SetRequestDate gets a reference to the given string and assigns it to the RequestDate field.
func (o *Receipt) SetRequestDate(v string) {
	o.RequestDate = &v
}

// GetRequestDateMs returns the RequestDateMs field value if set, zero value otherwise.
func (o *Receipt) GetRequestDateMs() string {
	if o == nil || isNil(o.RequestDateMs) {
		var ret string
		return ret
	}
	return *o.RequestDateMs
}

// GetRequestDateMsOk returns a tuple with the RequestDateMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetRequestDateMsOk() (*string, bool) {
	if o == nil || isNil(o.RequestDateMs) {
    return nil, false
	}
	return o.RequestDateMs, true
}

// HasRequestDateMs returns a boolean if a field has been set.
func (o *Receipt) HasRequestDateMs() bool {
	if o != nil && !isNil(o.RequestDateMs) {
		return true
	}

	return false
}

// SetRequestDateMs gets a reference to the given string and assigns it to the RequestDateMs field.
func (o *Receipt) SetRequestDateMs(v string) {
	o.RequestDateMs = &v
}

// GetRequestDatePst returns the RequestDatePst field value if set, zero value otherwise.
func (o *Receipt) GetRequestDatePst() string {
	if o == nil || isNil(o.RequestDatePst) {
		var ret string
		return ret
	}
	return *o.RequestDatePst
}

// GetRequestDatePstOk returns a tuple with the RequestDatePst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetRequestDatePstOk() (*string, bool) {
	if o == nil || isNil(o.RequestDatePst) {
    return nil, false
	}
	return o.RequestDatePst, true
}

// HasRequestDatePst returns a boolean if a field has been set.
func (o *Receipt) HasRequestDatePst() bool {
	if o != nil && !isNil(o.RequestDatePst) {
		return true
	}

	return false
}

// SetRequestDatePst gets a reference to the given string and assigns it to the RequestDatePst field.
func (o *Receipt) SetRequestDatePst(v string) {
	o.RequestDatePst = &v
}

// GetOriginalPurchaseDate returns the OriginalPurchaseDate field value if set, zero value otherwise.
func (o *Receipt) GetOriginalPurchaseDate() string {
	if o == nil || isNil(o.OriginalPurchaseDate) {
		var ret string
		return ret
	}
	return *o.OriginalPurchaseDate
}

// GetOriginalPurchaseDateOk returns a tuple with the OriginalPurchaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetOriginalPurchaseDateOk() (*string, bool) {
	if o == nil || isNil(o.OriginalPurchaseDate) {
    return nil, false
	}
	return o.OriginalPurchaseDate, true
}

// HasOriginalPurchaseDate returns a boolean if a field has been set.
func (o *Receipt) HasOriginalPurchaseDate() bool {
	if o != nil && !isNil(o.OriginalPurchaseDate) {
		return true
	}

	return false
}

// SetOriginalPurchaseDate gets a reference to the given string and assigns it to the OriginalPurchaseDate field.
func (o *Receipt) SetOriginalPurchaseDate(v string) {
	o.OriginalPurchaseDate = &v
}

// GetOriginalPurchaseDateMs returns the OriginalPurchaseDateMs field value if set, zero value otherwise.
func (o *Receipt) GetOriginalPurchaseDateMs() string {
	if o == nil || isNil(o.OriginalPurchaseDateMs) {
		var ret string
		return ret
	}
	return *o.OriginalPurchaseDateMs
}

// GetOriginalPurchaseDateMsOk returns a tuple with the OriginalPurchaseDateMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetOriginalPurchaseDateMsOk() (*string, bool) {
	if o == nil || isNil(o.OriginalPurchaseDateMs) {
    return nil, false
	}
	return o.OriginalPurchaseDateMs, true
}

// HasOriginalPurchaseDateMs returns a boolean if a field has been set.
func (o *Receipt) HasOriginalPurchaseDateMs() bool {
	if o != nil && !isNil(o.OriginalPurchaseDateMs) {
		return true
	}

	return false
}

// SetOriginalPurchaseDateMs gets a reference to the given string and assigns it to the OriginalPurchaseDateMs field.
func (o *Receipt) SetOriginalPurchaseDateMs(v string) {
	o.OriginalPurchaseDateMs = &v
}

// GetOriginalPurchaseDatePst returns the OriginalPurchaseDatePst field value if set, zero value otherwise.
func (o *Receipt) GetOriginalPurchaseDatePst() string {
	if o == nil || isNil(o.OriginalPurchaseDatePst) {
		var ret string
		return ret
	}
	return *o.OriginalPurchaseDatePst
}

// GetOriginalPurchaseDatePstOk returns a tuple with the OriginalPurchaseDatePst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetOriginalPurchaseDatePstOk() (*string, bool) {
	if o == nil || isNil(o.OriginalPurchaseDatePst) {
    return nil, false
	}
	return o.OriginalPurchaseDatePst, true
}

// HasOriginalPurchaseDatePst returns a boolean if a field has been set.
func (o *Receipt) HasOriginalPurchaseDatePst() bool {
	if o != nil && !isNil(o.OriginalPurchaseDatePst) {
		return true
	}

	return false
}

// SetOriginalPurchaseDatePst gets a reference to the given string and assigns it to the OriginalPurchaseDatePst field.
func (o *Receipt) SetOriginalPurchaseDatePst(v string) {
	o.OriginalPurchaseDatePst = &v
}

// GetOriginalApplicationVersion returns the OriginalApplicationVersion field value if set, zero value otherwise.
func (o *Receipt) GetOriginalApplicationVersion() string {
	if o == nil || isNil(o.OriginalApplicationVersion) {
		var ret string
		return ret
	}
	return *o.OriginalApplicationVersion
}

// GetOriginalApplicationVersionOk returns a tuple with the OriginalApplicationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetOriginalApplicationVersionOk() (*string, bool) {
	if o == nil || isNil(o.OriginalApplicationVersion) {
    return nil, false
	}
	return o.OriginalApplicationVersion, true
}

// HasOriginalApplicationVersion returns a boolean if a field has been set.
func (o *Receipt) HasOriginalApplicationVersion() bool {
	if o != nil && !isNil(o.OriginalApplicationVersion) {
		return true
	}

	return false
}

// SetOriginalApplicationVersion gets a reference to the given string and assigns it to the OriginalApplicationVersion field.
func (o *Receipt) SetOriginalApplicationVersion(v string) {
	o.OriginalApplicationVersion = &v
}

// GetInApp returns the InApp field value if set, zero value otherwise.
func (o *Receipt) GetInApp() []InApp {
	if o == nil || isNil(o.InApp) {
		var ret []InApp
		return ret
	}
	return o.InApp
}

// GetInAppOk returns a tuple with the InApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetInAppOk() ([]InApp, bool) {
	if o == nil || isNil(o.InApp) {
    return nil, false
	}
	return o.InApp, true
}

// HasInApp returns a boolean if a field has been set.
func (o *Receipt) HasInApp() bool {
	if o != nil && !isNil(o.InApp) {
		return true
	}

	return false
}

// SetInApp gets a reference to the given []InApp and assigns it to the InApp field.
func (o *Receipt) SetInApp(v []InApp) {
	o.InApp = v
}

func (o Receipt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReceiptType) {
		toSerialize["receipt_type"] = o.ReceiptType
	}
	if !isNil(o.AdamId) {
		toSerialize["adam_id"] = o.AdamId
	}
	if !isNil(o.AppItemId) {
		toSerialize["app_item_id"] = o.AppItemId
	}
	if !isNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !isNil(o.ApplicationVersion) {
		toSerialize["application_version"] = o.ApplicationVersion
	}
	if !isNil(o.DownloadId) {
		toSerialize["download_id"] = o.DownloadId
	}
	if !isNil(o.VersionExternalIdentifier) {
		toSerialize["version_external_identifier"] = o.VersionExternalIdentifier
	}
	if !isNil(o.ReceiptCreationDate) {
		toSerialize["receipt_creation_date"] = o.ReceiptCreationDate
	}
	if !isNil(o.ReceiptCreationDateMs) {
		toSerialize["receipt_creation_date_ms"] = o.ReceiptCreationDateMs
	}
	if !isNil(o.ReceiptCreationDatePst) {
		toSerialize["receipt_creation_date_pst"] = o.ReceiptCreationDatePst
	}
	if !isNil(o.RequestDate) {
		toSerialize["request_date"] = o.RequestDate
	}
	if !isNil(o.RequestDateMs) {
		toSerialize["request_date_ms"] = o.RequestDateMs
	}
	if !isNil(o.RequestDatePst) {
		toSerialize["request_date_pst"] = o.RequestDatePst
	}
	if !isNil(o.OriginalPurchaseDate) {
		toSerialize["original_purchase_date"] = o.OriginalPurchaseDate
	}
	if !isNil(o.OriginalPurchaseDateMs) {
		toSerialize["original_purchase_date_ms"] = o.OriginalPurchaseDateMs
	}
	if !isNil(o.OriginalPurchaseDatePst) {
		toSerialize["original_purchase_date_pst"] = o.OriginalPurchaseDatePst
	}
	if !isNil(o.OriginalApplicationVersion) {
		toSerialize["original_application_version"] = o.OriginalApplicationVersion
	}
	if !isNil(o.InApp) {
		toSerialize["in_app"] = o.InApp
	}
	return json.Marshal(toSerialize)
}

type NullableReceipt struct {
	value *Receipt
	isSet bool
}

func (v NullableReceipt) Get() *Receipt {
	return v.value
}

func (v *NullableReceipt) Set(val *Receipt) {
	v.value = val
	v.isSet = true
}

func (v NullableReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceipt(val *Receipt) *NullableReceipt {
	return &NullableReceipt{value: val, isSet: true}
}

func (v NullableReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



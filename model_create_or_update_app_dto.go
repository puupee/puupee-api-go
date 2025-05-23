/*
Puupee API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.17.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package puupee

import (
	"encoding/json"
)

// checks if the CreateOrUpdateAppDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateAppDto{}

// CreateOrUpdateAppDto struct for CreateOrUpdateAppDto
type CreateOrUpdateAppDto struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Framework *AppFramework `json:"framework,omitempty"`
	AppType *AppType `json:"appType,omitempty"`
	Description *string `json:"description,omitempty"`
	Icon *string `json:"icon,omitempty"`
	// 产品首页
	HomePage *string `json:"homePage,omitempty"`
	// 显示排序
	SortIndex *int32 `json:"sortIndex,omitempty"`
	GitRepository *string `json:"gitRepository,omitempty"`
	GitRepositoryType *GitRepositoryType `json:"gitRepositoryType,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Webhook Url 各种事件回调地址
	WebhookUrl *string `json:"webhookUrl,omitempty"`
	// 业务域名
	BusinessDomain *string `json:"businessDomain,omitempty"`
	// 业务地址
	BusinessUrl *string `json:"businessUrl,omitempty"`
	// 可以订阅的平台 Platform 枚举, 并用\",\"分割
	SubscriptionPlatforms *string `json:"subscriptionPlatforms,omitempty"`
	// 暂时免费的平台, 付费功能免费用的平台, 用\",\"分割
	FreePlatforms *string `json:"freePlatforms,omitempty"`
	// 声明格式
	SpecJsonSchema *string `json:"specJsonSchema,omitempty"`
	// 默认存储空间大小
	DefaultStorageSize *int64 `json:"defaultStorageSize,omitempty"`
	// 默认单文件最大大小
	DefaultSingleFileMaxSize *int64 `json:"defaultSingleFileMaxSize,omitempty"`
	// 是否已经发布, 决定了是否给终端用户看见, 主要有一些 APP 自己使用
	IsPublished *bool `json:"isPublished,omitempty"`
	OpenClient *CreateOpenIddictApplicationDto `json:"openClient,omitempty"`
}

// NewCreateOrUpdateAppDto instantiates a new CreateOrUpdateAppDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateAppDto() *CreateOrUpdateAppDto {
	this := CreateOrUpdateAppDto{}
	return &this
}

// NewCreateOrUpdateAppDtoWithDefaults instantiates a new CreateOrUpdateAppDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateAppDtoWithDefaults() *CreateOrUpdateAppDto {
	this := CreateOrUpdateAppDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrUpdateAppDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateOrUpdateAppDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetFramework() AppFramework {
	if o == nil || IsNil(o.Framework) {
		var ret AppFramework
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetFrameworkOk() (*AppFramework, bool) {
	if o == nil || IsNil(o.Framework) {
		return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasFramework() bool {
	if o != nil && !IsNil(o.Framework) {
		return true
	}

	return false
}

// SetFramework gets a reference to the given AppFramework and assigns it to the Framework field.
func (o *CreateOrUpdateAppDto) SetFramework(v AppFramework) {
	o.Framework = &v
}

// GetAppType returns the AppType field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetAppType() AppType {
	if o == nil || IsNil(o.AppType) {
		var ret AppType
		return ret
	}
	return *o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetAppTypeOk() (*AppType, bool) {
	if o == nil || IsNil(o.AppType) {
		return nil, false
	}
	return o.AppType, true
}

// HasAppType returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasAppType() bool {
	if o != nil && !IsNil(o.AppType) {
		return true
	}

	return false
}

// SetAppType gets a reference to the given AppType and assigns it to the AppType field.
func (o *CreateOrUpdateAppDto) SetAppType(v AppType) {
	o.AppType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrUpdateAppDto) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *CreateOrUpdateAppDto) SetIcon(v string) {
	o.Icon = &v
}

// GetHomePage returns the HomePage field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetHomePage() string {
	if o == nil || IsNil(o.HomePage) {
		var ret string
		return ret
	}
	return *o.HomePage
}

// GetHomePageOk returns a tuple with the HomePage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetHomePageOk() (*string, bool) {
	if o == nil || IsNil(o.HomePage) {
		return nil, false
	}
	return o.HomePage, true
}

// HasHomePage returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasHomePage() bool {
	if o != nil && !IsNil(o.HomePage) {
		return true
	}

	return false
}

// SetHomePage gets a reference to the given string and assigns it to the HomePage field.
func (o *CreateOrUpdateAppDto) SetHomePage(v string) {
	o.HomePage = &v
}

// GetSortIndex returns the SortIndex field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetSortIndex() int32 {
	if o == nil || IsNil(o.SortIndex) {
		var ret int32
		return ret
	}
	return *o.SortIndex
}

// GetSortIndexOk returns a tuple with the SortIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetSortIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.SortIndex) {
		return nil, false
	}
	return o.SortIndex, true
}

// HasSortIndex returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasSortIndex() bool {
	if o != nil && !IsNil(o.SortIndex) {
		return true
	}

	return false
}

// SetSortIndex gets a reference to the given int32 and assigns it to the SortIndex field.
func (o *CreateOrUpdateAppDto) SetSortIndex(v int32) {
	o.SortIndex = &v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetGitRepository() string {
	if o == nil || IsNil(o.GitRepository) {
		var ret string
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetGitRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.GitRepository) {
		return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasGitRepository() bool {
	if o != nil && !IsNil(o.GitRepository) {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given string and assigns it to the GitRepository field.
func (o *CreateOrUpdateAppDto) SetGitRepository(v string) {
	o.GitRepository = &v
}

// GetGitRepositoryType returns the GitRepositoryType field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetGitRepositoryType() GitRepositoryType {
	if o == nil || IsNil(o.GitRepositoryType) {
		var ret GitRepositoryType
		return ret
	}
	return *o.GitRepositoryType
}

// GetGitRepositoryTypeOk returns a tuple with the GitRepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetGitRepositoryTypeOk() (*GitRepositoryType, bool) {
	if o == nil || IsNil(o.GitRepositoryType) {
		return nil, false
	}
	return o.GitRepositoryType, true
}

// HasGitRepositoryType returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasGitRepositoryType() bool {
	if o != nil && !IsNil(o.GitRepositoryType) {
		return true
	}

	return false
}

// SetGitRepositoryType gets a reference to the given GitRepositoryType and assigns it to the GitRepositoryType field.
func (o *CreateOrUpdateAppDto) SetGitRepositoryType(v GitRepositoryType) {
	o.GitRepositoryType = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *CreateOrUpdateAppDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *CreateOrUpdateAppDto) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetBusinessDomain returns the BusinessDomain field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetBusinessDomain() string {
	if o == nil || IsNil(o.BusinessDomain) {
		var ret string
		return ret
	}
	return *o.BusinessDomain
}

// GetBusinessDomainOk returns a tuple with the BusinessDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetBusinessDomainOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessDomain) {
		return nil, false
	}
	return o.BusinessDomain, true
}

// HasBusinessDomain returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasBusinessDomain() bool {
	if o != nil && !IsNil(o.BusinessDomain) {
		return true
	}

	return false
}

// SetBusinessDomain gets a reference to the given string and assigns it to the BusinessDomain field.
func (o *CreateOrUpdateAppDto) SetBusinessDomain(v string) {
	o.BusinessDomain = &v
}

// GetBusinessUrl returns the BusinessUrl field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetBusinessUrl() string {
	if o == nil || IsNil(o.BusinessUrl) {
		var ret string
		return ret
	}
	return *o.BusinessUrl
}

// GetBusinessUrlOk returns a tuple with the BusinessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetBusinessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessUrl) {
		return nil, false
	}
	return o.BusinessUrl, true
}

// HasBusinessUrl returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasBusinessUrl() bool {
	if o != nil && !IsNil(o.BusinessUrl) {
		return true
	}

	return false
}

// SetBusinessUrl gets a reference to the given string and assigns it to the BusinessUrl field.
func (o *CreateOrUpdateAppDto) SetBusinessUrl(v string) {
	o.BusinessUrl = &v
}

// GetSubscriptionPlatforms returns the SubscriptionPlatforms field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetSubscriptionPlatforms() string {
	if o == nil || IsNil(o.SubscriptionPlatforms) {
		var ret string
		return ret
	}
	return *o.SubscriptionPlatforms
}

// GetSubscriptionPlatformsOk returns a tuple with the SubscriptionPlatforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetSubscriptionPlatformsOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionPlatforms) {
		return nil, false
	}
	return o.SubscriptionPlatforms, true
}

// HasSubscriptionPlatforms returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasSubscriptionPlatforms() bool {
	if o != nil && !IsNil(o.SubscriptionPlatforms) {
		return true
	}

	return false
}

// SetSubscriptionPlatforms gets a reference to the given string and assigns it to the SubscriptionPlatforms field.
func (o *CreateOrUpdateAppDto) SetSubscriptionPlatforms(v string) {
	o.SubscriptionPlatforms = &v
}

// GetFreePlatforms returns the FreePlatforms field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetFreePlatforms() string {
	if o == nil || IsNil(o.FreePlatforms) {
		var ret string
		return ret
	}
	return *o.FreePlatforms
}

// GetFreePlatformsOk returns a tuple with the FreePlatforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetFreePlatformsOk() (*string, bool) {
	if o == nil || IsNil(o.FreePlatforms) {
		return nil, false
	}
	return o.FreePlatforms, true
}

// HasFreePlatforms returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasFreePlatforms() bool {
	if o != nil && !IsNil(o.FreePlatforms) {
		return true
	}

	return false
}

// SetFreePlatforms gets a reference to the given string and assigns it to the FreePlatforms field.
func (o *CreateOrUpdateAppDto) SetFreePlatforms(v string) {
	o.FreePlatforms = &v
}

// GetSpecJsonSchema returns the SpecJsonSchema field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetSpecJsonSchema() string {
	if o == nil || IsNil(o.SpecJsonSchema) {
		var ret string
		return ret
	}
	return *o.SpecJsonSchema
}

// GetSpecJsonSchemaOk returns a tuple with the SpecJsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetSpecJsonSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.SpecJsonSchema) {
		return nil, false
	}
	return o.SpecJsonSchema, true
}

// HasSpecJsonSchema returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasSpecJsonSchema() bool {
	if o != nil && !IsNil(o.SpecJsonSchema) {
		return true
	}

	return false
}

// SetSpecJsonSchema gets a reference to the given string and assigns it to the SpecJsonSchema field.
func (o *CreateOrUpdateAppDto) SetSpecJsonSchema(v string) {
	o.SpecJsonSchema = &v
}

// GetDefaultStorageSize returns the DefaultStorageSize field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetDefaultStorageSize() int64 {
	if o == nil || IsNil(o.DefaultStorageSize) {
		var ret int64
		return ret
	}
	return *o.DefaultStorageSize
}

// GetDefaultStorageSizeOk returns a tuple with the DefaultStorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetDefaultStorageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultStorageSize) {
		return nil, false
	}
	return o.DefaultStorageSize, true
}

// HasDefaultStorageSize returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasDefaultStorageSize() bool {
	if o != nil && !IsNil(o.DefaultStorageSize) {
		return true
	}

	return false
}

// SetDefaultStorageSize gets a reference to the given int64 and assigns it to the DefaultStorageSize field.
func (o *CreateOrUpdateAppDto) SetDefaultStorageSize(v int64) {
	o.DefaultStorageSize = &v
}

// GetDefaultSingleFileMaxSize returns the DefaultSingleFileMaxSize field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetDefaultSingleFileMaxSize() int64 {
	if o == nil || IsNil(o.DefaultSingleFileMaxSize) {
		var ret int64
		return ret
	}
	return *o.DefaultSingleFileMaxSize
}

// GetDefaultSingleFileMaxSizeOk returns a tuple with the DefaultSingleFileMaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetDefaultSingleFileMaxSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultSingleFileMaxSize) {
		return nil, false
	}
	return o.DefaultSingleFileMaxSize, true
}

// HasDefaultSingleFileMaxSize returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasDefaultSingleFileMaxSize() bool {
	if o != nil && !IsNil(o.DefaultSingleFileMaxSize) {
		return true
	}

	return false
}

// SetDefaultSingleFileMaxSize gets a reference to the given int64 and assigns it to the DefaultSingleFileMaxSize field.
func (o *CreateOrUpdateAppDto) SetDefaultSingleFileMaxSize(v int64) {
	o.DefaultSingleFileMaxSize = &v
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetIsPublished() bool {
	if o == nil || IsNil(o.IsPublished) {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetIsPublishedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublished) {
		return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasIsPublished() bool {
	if o != nil && !IsNil(o.IsPublished) {
		return true
	}

	return false
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *CreateOrUpdateAppDto) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetOpenClient returns the OpenClient field value if set, zero value otherwise.
func (o *CreateOrUpdateAppDto) GetOpenClient() CreateOpenIddictApplicationDto {
	if o == nil || IsNil(o.OpenClient) {
		var ret CreateOpenIddictApplicationDto
		return ret
	}
	return *o.OpenClient
}

// GetOpenClientOk returns a tuple with the OpenClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppDto) GetOpenClientOk() (*CreateOpenIddictApplicationDto, bool) {
	if o == nil || IsNil(o.OpenClient) {
		return nil, false
	}
	return o.OpenClient, true
}

// HasOpenClient returns a boolean if a field has been set.
func (o *CreateOrUpdateAppDto) HasOpenClient() bool {
	if o != nil && !IsNil(o.OpenClient) {
		return true
	}

	return false
}

// SetOpenClient gets a reference to the given CreateOpenIddictApplicationDto and assigns it to the OpenClient field.
func (o *CreateOrUpdateAppDto) SetOpenClient(v CreateOpenIddictApplicationDto) {
	o.OpenClient = &v
}

func (o CreateOrUpdateAppDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateAppDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Framework) {
		toSerialize["framework"] = o.Framework
	}
	if !IsNil(o.AppType) {
		toSerialize["appType"] = o.AppType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.HomePage) {
		toSerialize["homePage"] = o.HomePage
	}
	if !IsNil(o.SortIndex) {
		toSerialize["sortIndex"] = o.SortIndex
	}
	if !IsNil(o.GitRepository) {
		toSerialize["gitRepository"] = o.GitRepository
	}
	if !IsNil(o.GitRepositoryType) {
		toSerialize["gitRepositoryType"] = o.GitRepositoryType
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	if !IsNil(o.BusinessDomain) {
		toSerialize["businessDomain"] = o.BusinessDomain
	}
	if !IsNil(o.BusinessUrl) {
		toSerialize["businessUrl"] = o.BusinessUrl
	}
	if !IsNil(o.SubscriptionPlatforms) {
		toSerialize["subscriptionPlatforms"] = o.SubscriptionPlatforms
	}
	if !IsNil(o.FreePlatforms) {
		toSerialize["freePlatforms"] = o.FreePlatforms
	}
	if !IsNil(o.SpecJsonSchema) {
		toSerialize["specJsonSchema"] = o.SpecJsonSchema
	}
	if !IsNil(o.DefaultStorageSize) {
		toSerialize["defaultStorageSize"] = o.DefaultStorageSize
	}
	if !IsNil(o.DefaultSingleFileMaxSize) {
		toSerialize["defaultSingleFileMaxSize"] = o.DefaultSingleFileMaxSize
	}
	if !IsNil(o.IsPublished) {
		toSerialize["isPublished"] = o.IsPublished
	}
	if !IsNil(o.OpenClient) {
		toSerialize["openClient"] = o.OpenClient
	}
	return toSerialize, nil
}

type NullableCreateOrUpdateAppDto struct {
	value *CreateOrUpdateAppDto
	isSet bool
}

func (v NullableCreateOrUpdateAppDto) Get() *CreateOrUpdateAppDto {
	return v.value
}

func (v *NullableCreateOrUpdateAppDto) Set(val *CreateOrUpdateAppDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateAppDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateAppDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateAppDto(val *CreateOrUpdateAppDto) *NullableCreateOrUpdateAppDto {
	return &NullableCreateOrUpdateAppDto{value: val, isSet: true}
}

func (v NullableCreateOrUpdateAppDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateAppDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



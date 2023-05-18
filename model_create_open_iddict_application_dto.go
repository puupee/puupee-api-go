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

// checks if the CreateOpenIddictApplicationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOpenIddictApplicationDto{}

// CreateOpenIddictApplicationDto struct for CreateOpenIddictApplicationDto
type CreateOpenIddictApplicationDto struct {
	Type NullableString `json:"type,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	DisplayNames NullableString `json:"displayNames,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	PostLogoutRedirectUris NullableString `json:"postLogoutRedirectUris,omitempty"`
	Properties NullableString `json:"properties,omitempty"`
	RedirectUris NullableString `json:"redirectUris,omitempty"`
	Requirements NullableString `json:"requirements,omitempty"`
	ClientUri NullableString `json:"clientUri,omitempty"`
	LogoUri NullableString `json:"logoUri,omitempty"`
	GrantTypes []string `json:"grantTypes,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

// NewCreateOpenIddictApplicationDto instantiates a new CreateOpenIddictApplicationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOpenIddictApplicationDto() *CreateOpenIddictApplicationDto {
	this := CreateOpenIddictApplicationDto{}
	return &this
}

// NewCreateOpenIddictApplicationDtoWithDefaults instantiates a new CreateOpenIddictApplicationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOpenIddictApplicationDtoWithDefaults() *CreateOpenIddictApplicationDto {
	this := CreateOpenIddictApplicationDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CreateOpenIddictApplicationDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *CreateOpenIddictApplicationDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CreateOpenIddictApplicationDto) UnsetType() {
	o.Type.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *CreateOpenIddictApplicationDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *CreateOpenIddictApplicationDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *CreateOpenIddictApplicationDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetDisplayNames returns the DisplayNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetDisplayNames() string {
	if o == nil || IsNil(o.DisplayNames.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayNames.Get()
}

// GetDisplayNamesOk returns a tuple with the DisplayNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetDisplayNamesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayNames.Get(), o.DisplayNames.IsSet()
}

// HasDisplayNames returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasDisplayNames() bool {
	if o != nil && o.DisplayNames.IsSet() {
		return true
	}

	return false
}

// SetDisplayNames gets a reference to the given NullableString and assigns it to the DisplayNames field.
func (o *CreateOpenIddictApplicationDto) SetDisplayNames(v string) {
	o.DisplayNames.Set(&v)
}
// SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil
func (o *CreateOpenIddictApplicationDto) SetDisplayNamesNil() {
	o.DisplayNames.Set(nil)
}

// UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
func (o *CreateOpenIddictApplicationDto) UnsetDisplayNames() {
	o.DisplayNames.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasPermissions() bool {
	if o != nil && IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *CreateOpenIddictApplicationDto) SetPermissions(v []string) {
	o.Permissions = v
}

// GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetPostLogoutRedirectUris() string {
	if o == nil || IsNil(o.PostLogoutRedirectUris.Get()) {
		var ret string
		return ret
	}
	return *o.PostLogoutRedirectUris.Get()
}

// GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetPostLogoutRedirectUrisOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostLogoutRedirectUris.Get(), o.PostLogoutRedirectUris.IsSet()
}

// HasPostLogoutRedirectUris returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasPostLogoutRedirectUris() bool {
	if o != nil && o.PostLogoutRedirectUris.IsSet() {
		return true
	}

	return false
}

// SetPostLogoutRedirectUris gets a reference to the given NullableString and assigns it to the PostLogoutRedirectUris field.
func (o *CreateOpenIddictApplicationDto) SetPostLogoutRedirectUris(v string) {
	o.PostLogoutRedirectUris.Set(&v)
}
// SetPostLogoutRedirectUrisNil sets the value for PostLogoutRedirectUris to be an explicit nil
func (o *CreateOpenIddictApplicationDto) SetPostLogoutRedirectUrisNil() {
	o.PostLogoutRedirectUris.Set(nil)
}

// UnsetPostLogoutRedirectUris ensures that no value is present for PostLogoutRedirectUris, not even an explicit nil
func (o *CreateOpenIddictApplicationDto) UnsetPostLogoutRedirectUris() {
	o.PostLogoutRedirectUris.Unset()
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetProperties() string {
	if o == nil || IsNil(o.Properties.Get()) {
		var ret string
		return ret
	}
	return *o.Properties.Get()
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetPropertiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Get(), o.Properties.IsSet()
}

// HasProperties returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasProperties() bool {
	if o != nil && o.Properties.IsSet() {
		return true
	}

	return false
}

// SetProperties gets a reference to the given NullableString and assigns it to the Properties field.
func (o *CreateOpenIddictApplicationDto) SetProperties(v string) {
	o.Properties.Set(&v)
}
// SetPropertiesNil sets the value for Properties to be an explicit nil
func (o *CreateOpenIddictApplicationDto) SetPropertiesNil() {
	o.Properties.Set(nil)
}

// UnsetProperties ensures that no value is present for Properties, not even an explicit nil
func (o *CreateOpenIddictApplicationDto) UnsetProperties() {
	o.Properties.Unset()
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetRedirectUris() string {
	if o == nil || IsNil(o.RedirectUris.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUris.Get()
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetRedirectUrisOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUris.Get(), o.RedirectUris.IsSet()
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasRedirectUris() bool {
	if o != nil && o.RedirectUris.IsSet() {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given NullableString and assigns it to the RedirectUris field.
func (o *CreateOpenIddictApplicationDto) SetRedirectUris(v string) {
	o.RedirectUris.Set(&v)
}
// SetRedirectUrisNil sets the value for RedirectUris to be an explicit nil
func (o *CreateOpenIddictApplicationDto) SetRedirectUrisNil() {
	o.RedirectUris.Set(nil)
}

// UnsetRedirectUris ensures that no value is present for RedirectUris, not even an explicit nil
func (o *CreateOpenIddictApplicationDto) UnsetRedirectUris() {
	o.RedirectUris.Unset()
}

// GetRequirements returns the Requirements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetRequirements() string {
	if o == nil || IsNil(o.Requirements.Get()) {
		var ret string
		return ret
	}
	return *o.Requirements.Get()
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetRequirementsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requirements.Get(), o.Requirements.IsSet()
}

// HasRequirements returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasRequirements() bool {
	if o != nil && o.Requirements.IsSet() {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given NullableString and assigns it to the Requirements field.
func (o *CreateOpenIddictApplicationDto) SetRequirements(v string) {
	o.Requirements.Set(&v)
}
// SetRequirementsNil sets the value for Requirements to be an explicit nil
func (o *CreateOpenIddictApplicationDto) SetRequirementsNil() {
	o.Requirements.Set(nil)
}

// UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
func (o *CreateOpenIddictApplicationDto) UnsetRequirements() {
	o.Requirements.Unset()
}

// GetClientUri returns the ClientUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetClientUri() string {
	if o == nil || IsNil(o.ClientUri.Get()) {
		var ret string
		return ret
	}
	return *o.ClientUri.Get()
}

// GetClientUriOk returns a tuple with the ClientUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetClientUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientUri.Get(), o.ClientUri.IsSet()
}

// HasClientUri returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasClientUri() bool {
	if o != nil && o.ClientUri.IsSet() {
		return true
	}

	return false
}

// SetClientUri gets a reference to the given NullableString and assigns it to the ClientUri field.
func (o *CreateOpenIddictApplicationDto) SetClientUri(v string) {
	o.ClientUri.Set(&v)
}
// SetClientUriNil sets the value for ClientUri to be an explicit nil
func (o *CreateOpenIddictApplicationDto) SetClientUriNil() {
	o.ClientUri.Set(nil)
}

// UnsetClientUri ensures that no value is present for ClientUri, not even an explicit nil
func (o *CreateOpenIddictApplicationDto) UnsetClientUri() {
	o.ClientUri.Unset()
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetLogoUri() string {
	if o == nil || IsNil(o.LogoUri.Get()) {
		var ret string
		return ret
	}
	return *o.LogoUri.Get()
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetLogoUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUri.Get(), o.LogoUri.IsSet()
}

// HasLogoUri returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasLogoUri() bool {
	if o != nil && o.LogoUri.IsSet() {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given NullableString and assigns it to the LogoUri field.
func (o *CreateOpenIddictApplicationDto) SetLogoUri(v string) {
	o.LogoUri.Set(&v)
}
// SetLogoUriNil sets the value for LogoUri to be an explicit nil
func (o *CreateOpenIddictApplicationDto) SetLogoUriNil() {
	o.LogoUri.Set(nil)
}

// UnsetLogoUri ensures that no value is present for LogoUri, not even an explicit nil
func (o *CreateOpenIddictApplicationDto) UnsetLogoUri() {
	o.LogoUri.Unset()
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetGrantTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetGrantTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.GrantTypes) {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasGrantTypes() bool {
	if o != nil && IsNil(o.GrantTypes) {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *CreateOpenIddictApplicationDto) SetGrantTypes(v []string) {
	o.GrantTypes = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpenIddictApplicationDto) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpenIddictApplicationDto) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *CreateOpenIddictApplicationDto) HasScopes() bool {
	if o != nil && IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *CreateOpenIddictApplicationDto) SetScopes(v []string) {
	o.Scopes = v
}

func (o CreateOpenIddictApplicationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOpenIddictApplicationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.DisplayNames.IsSet() {
		toSerialize["displayNames"] = o.DisplayNames.Get()
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.PostLogoutRedirectUris.IsSet() {
		toSerialize["postLogoutRedirectUris"] = o.PostLogoutRedirectUris.Get()
	}
	if o.Properties.IsSet() {
		toSerialize["properties"] = o.Properties.Get()
	}
	if o.RedirectUris.IsSet() {
		toSerialize["redirectUris"] = o.RedirectUris.Get()
	}
	if o.Requirements.IsSet() {
		toSerialize["requirements"] = o.Requirements.Get()
	}
	if o.ClientUri.IsSet() {
		toSerialize["clientUri"] = o.ClientUri.Get()
	}
	if o.LogoUri.IsSet() {
		toSerialize["logoUri"] = o.LogoUri.Get()
	}
	if o.GrantTypes != nil {
		toSerialize["grantTypes"] = o.GrantTypes
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableCreateOpenIddictApplicationDto struct {
	value *CreateOpenIddictApplicationDto
	isSet bool
}

func (v NullableCreateOpenIddictApplicationDto) Get() *CreateOpenIddictApplicationDto {
	return v.value
}

func (v *NullableCreateOpenIddictApplicationDto) Set(val *CreateOpenIddictApplicationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOpenIddictApplicationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOpenIddictApplicationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOpenIddictApplicationDto(val *CreateOpenIddictApplicationDto) *NullableCreateOpenIddictApplicationDto {
	return &NullableCreateOpenIddictApplicationDto{value: val, isSet: true}
}

func (v NullableCreateOpenIddictApplicationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOpenIddictApplicationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



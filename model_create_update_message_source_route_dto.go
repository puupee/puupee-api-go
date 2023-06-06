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

// checks if the CreateUpdateMessageSourceRouteDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUpdateMessageSourceRouteDto{}

// CreateUpdateMessageSourceRouteDto struct for CreateUpdateMessageSourceRouteDto
type CreateUpdateMessageSourceRouteDto struct {
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Path *string `json:"path,omitempty"`
	SourceId *string `json:"sourceId,omitempty"`
	Extra *string `json:"extra,omitempty"`
	Anticrawler *bool `json:"anticrawler,omitempty"`
	Radar *bool `json:"radar,omitempty"`
	Rssbud *bool `json:"rssbud,omitempty"`
	IsPublished *bool `json:"isPublished,omitempty"`
	IconUrl *string `json:"iconUrl,omitempty"`
}

// NewCreateUpdateMessageSourceRouteDto instantiates a new CreateUpdateMessageSourceRouteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpdateMessageSourceRouteDto() *CreateUpdateMessageSourceRouteDto {
	this := CreateUpdateMessageSourceRouteDto{}
	return &this
}

// NewCreateUpdateMessageSourceRouteDtoWithDefaults instantiates a new CreateUpdateMessageSourceRouteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpdateMessageSourceRouteDtoWithDefaults() *CreateUpdateMessageSourceRouteDto {
	this := CreateUpdateMessageSourceRouteDto{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateUpdateMessageSourceRouteDto) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateUpdateMessageSourceRouteDto) SetDescription(v string) {
	o.Description = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CreateUpdateMessageSourceRouteDto) SetPath(v string) {
	o.Path = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *CreateUpdateMessageSourceRouteDto) SetSourceId(v string) {
	o.SourceId = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetExtra() string {
	if o == nil || IsNil(o.Extra) {
		var ret string
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetExtraOk() (*string, bool) {
	if o == nil || IsNil(o.Extra) {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given string and assigns it to the Extra field.
func (o *CreateUpdateMessageSourceRouteDto) SetExtra(v string) {
	o.Extra = &v
}

// GetAnticrawler returns the Anticrawler field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetAnticrawler() bool {
	if o == nil || IsNil(o.Anticrawler) {
		var ret bool
		return ret
	}
	return *o.Anticrawler
}

// GetAnticrawlerOk returns a tuple with the Anticrawler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetAnticrawlerOk() (*bool, bool) {
	if o == nil || IsNil(o.Anticrawler) {
		return nil, false
	}
	return o.Anticrawler, true
}

// HasAnticrawler returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasAnticrawler() bool {
	if o != nil && !IsNil(o.Anticrawler) {
		return true
	}

	return false
}

// SetAnticrawler gets a reference to the given bool and assigns it to the Anticrawler field.
func (o *CreateUpdateMessageSourceRouteDto) SetAnticrawler(v bool) {
	o.Anticrawler = &v
}

// GetRadar returns the Radar field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetRadar() bool {
	if o == nil || IsNil(o.Radar) {
		var ret bool
		return ret
	}
	return *o.Radar
}

// GetRadarOk returns a tuple with the Radar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetRadarOk() (*bool, bool) {
	if o == nil || IsNil(o.Radar) {
		return nil, false
	}
	return o.Radar, true
}

// HasRadar returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasRadar() bool {
	if o != nil && !IsNil(o.Radar) {
		return true
	}

	return false
}

// SetRadar gets a reference to the given bool and assigns it to the Radar field.
func (o *CreateUpdateMessageSourceRouteDto) SetRadar(v bool) {
	o.Radar = &v
}

// GetRssbud returns the Rssbud field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetRssbud() bool {
	if o == nil || IsNil(o.Rssbud) {
		var ret bool
		return ret
	}
	return *o.Rssbud
}

// GetRssbudOk returns a tuple with the Rssbud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetRssbudOk() (*bool, bool) {
	if o == nil || IsNil(o.Rssbud) {
		return nil, false
	}
	return o.Rssbud, true
}

// HasRssbud returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasRssbud() bool {
	if o != nil && !IsNil(o.Rssbud) {
		return true
	}

	return false
}

// SetRssbud gets a reference to the given bool and assigns it to the Rssbud field.
func (o *CreateUpdateMessageSourceRouteDto) SetRssbud(v bool) {
	o.Rssbud = &v
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetIsPublished() bool {
	if o == nil || IsNil(o.IsPublished) {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetIsPublishedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublished) {
		return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasIsPublished() bool {
	if o != nil && !IsNil(o.IsPublished) {
		return true
	}

	return false
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *CreateUpdateMessageSourceRouteDto) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *CreateUpdateMessageSourceRouteDto) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateMessageSourceRouteDto) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *CreateUpdateMessageSourceRouteDto) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *CreateUpdateMessageSourceRouteDto) SetIconUrl(v string) {
	o.IconUrl = &v
}

func (o CreateUpdateMessageSourceRouteDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUpdateMessageSourceRouteDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	if !IsNil(o.Anticrawler) {
		toSerialize["anticrawler"] = o.Anticrawler
	}
	if !IsNil(o.Radar) {
		toSerialize["radar"] = o.Radar
	}
	if !IsNil(o.Rssbud) {
		toSerialize["rssbud"] = o.Rssbud
	}
	if !IsNil(o.IsPublished) {
		toSerialize["isPublished"] = o.IsPublished
	}
	if !IsNil(o.IconUrl) {
		toSerialize["iconUrl"] = o.IconUrl
	}
	return toSerialize, nil
}

type NullableCreateUpdateMessageSourceRouteDto struct {
	value *CreateUpdateMessageSourceRouteDto
	isSet bool
}

func (v NullableCreateUpdateMessageSourceRouteDto) Get() *CreateUpdateMessageSourceRouteDto {
	return v.value
}

func (v *NullableCreateUpdateMessageSourceRouteDto) Set(val *CreateUpdateMessageSourceRouteDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpdateMessageSourceRouteDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpdateMessageSourceRouteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpdateMessageSourceRouteDto(val *CreateUpdateMessageSourceRouteDto) *NullableCreateUpdateMessageSourceRouteDto {
	return &NullableCreateUpdateMessageSourceRouteDto{value: val, isSet: true}
}

func (v NullableCreateUpdateMessageSourceRouteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpdateMessageSourceRouteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



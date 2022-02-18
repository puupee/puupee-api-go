/*
Doggy API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package doggy

import (
	"encoding/json"
	"time"
)

// TodoDto struct for TodoDto
type TodoDto struct {
	Id *string `json:"id,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId NullableString `json:"creatorId,omitempty"`
	LastModificationTime NullableTime `json:"lastModificationTime,omitempty"`
	LastModifierId NullableString `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId NullableString `json:"deleterId,omitempty"`
	DeletionTime NullableTime `json:"deletionTime,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Priority *Priority `json:"priority,omitempty"`
	Tags []TagDto `json:"tags,omitempty"`
	DoneAt NullableTime `json:"doneAt,omitempty"`
	IsDone *bool `json:"isDone,omitempty"`
	Children []TodoDto `json:"children,omitempty"`
}

// NewTodoDto instantiates a new TodoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTodoDto() *TodoDto {
	this := TodoDto{}
	return &this
}

// NewTodoDtoWithDefaults instantiates a new TodoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTodoDtoWithDefaults() *TodoDto {
	this := TodoDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TodoDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TodoDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TodoDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TodoDto) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *TodoDto) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TodoDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *TodoDto) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *TodoDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TodoDto) GetCreatorId() string {
	if o == nil || o.CreatorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TodoDto) GetCreatorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *TodoDto) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableString and assigns it to the CreatorId field.
func (o *TodoDto) SetCreatorId(v string) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *TodoDto) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *TodoDto) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TodoDto) GetLastModificationTime() time.Time {
	if o == nil || o.LastModificationTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime.Get()
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TodoDto) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModificationTime.Get(), o.LastModificationTime.IsSet()
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *TodoDto) HasLastModificationTime() bool {
	if o != nil && o.LastModificationTime.IsSet() {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given NullableTime and assigns it to the LastModificationTime field.
func (o *TodoDto) SetLastModificationTime(v time.Time) {
	o.LastModificationTime.Set(&v)
}
// SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil
func (o *TodoDto) SetLastModificationTimeNil() {
	o.LastModificationTime.Set(nil)
}

// UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
func (o *TodoDto) UnsetLastModificationTime() {
	o.LastModificationTime.Unset()
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TodoDto) GetLastModifierId() string {
	if o == nil || o.LastModifierId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastModifierId.Get()
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TodoDto) GetLastModifierIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifierId.Get(), o.LastModifierId.IsSet()
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *TodoDto) HasLastModifierId() bool {
	if o != nil && o.LastModifierId.IsSet() {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given NullableString and assigns it to the LastModifierId field.
func (o *TodoDto) SetLastModifierId(v string) {
	o.LastModifierId.Set(&v)
}
// SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil
func (o *TodoDto) SetLastModifierIdNil() {
	o.LastModifierId.Set(nil)
}

// UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
func (o *TodoDto) UnsetLastModifierId() {
	o.LastModifierId.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *TodoDto) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TodoDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *TodoDto) HasIsDeleted() bool {
	if o != nil && o.IsDeleted != nil {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *TodoDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TodoDto) GetDeleterId() string {
	if o == nil || o.DeleterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeleterId.Get()
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TodoDto) GetDeleterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeleterId.Get(), o.DeleterId.IsSet()
}

// HasDeleterId returns a boolean if a field has been set.
func (o *TodoDto) HasDeleterId() bool {
	if o != nil && o.DeleterId.IsSet() {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given NullableString and assigns it to the DeleterId field.
func (o *TodoDto) SetDeleterId(v string) {
	o.DeleterId.Set(&v)
}
// SetDeleterIdNil sets the value for DeleterId to be an explicit nil
func (o *TodoDto) SetDeleterIdNil() {
	o.DeleterId.Set(nil)
}

// UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
func (o *TodoDto) UnsetDeleterId() {
	o.DeleterId.Unset()
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TodoDto) GetDeletionTime() time.Time {
	if o == nil || o.DeletionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime.Get()
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TodoDto) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletionTime.Get(), o.DeletionTime.IsSet()
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *TodoDto) HasDeletionTime() bool {
	if o != nil && o.DeletionTime.IsSet() {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given NullableTime and assigns it to the DeletionTime field.
func (o *TodoDto) SetDeletionTime(v time.Time) {
	o.DeletionTime.Set(&v)
}
// SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil
func (o *TodoDto) SetDeletionTimeNil() {
	o.DeletionTime.Set(nil)
}

// UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
func (o *TodoDto) UnsetDeletionTime() {
	o.DeletionTime.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TodoDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TodoDto) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TodoDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TodoDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TodoDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TodoDto) UnsetTitle() {
	o.Title.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TodoDto) GetPriority() Priority {
	if o == nil || o.Priority == nil {
		var ret Priority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TodoDto) GetPriorityOk() (*Priority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TodoDto) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given Priority and assigns it to the Priority field.
func (o *TodoDto) SetPriority(v Priority) {
	o.Priority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TodoDto) GetTags() []TagDto {
	if o == nil  {
		var ret []TagDto
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TodoDto) GetTagsOk() (*[]TagDto, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TodoDto) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagDto and assigns it to the Tags field.
func (o *TodoDto) SetTags(v []TagDto) {
	o.Tags = v
}

// GetDoneAt returns the DoneAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TodoDto) GetDoneAt() time.Time {
	if o == nil || o.DoneAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DoneAt.Get()
}

// GetDoneAtOk returns a tuple with the DoneAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TodoDto) GetDoneAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DoneAt.Get(), o.DoneAt.IsSet()
}

// HasDoneAt returns a boolean if a field has been set.
func (o *TodoDto) HasDoneAt() bool {
	if o != nil && o.DoneAt.IsSet() {
		return true
	}

	return false
}

// SetDoneAt gets a reference to the given NullableTime and assigns it to the DoneAt field.
func (o *TodoDto) SetDoneAt(v time.Time) {
	o.DoneAt.Set(&v)
}
// SetDoneAtNil sets the value for DoneAt to be an explicit nil
func (o *TodoDto) SetDoneAtNil() {
	o.DoneAt.Set(nil)
}

// UnsetDoneAt ensures that no value is present for DoneAt, not even an explicit nil
func (o *TodoDto) UnsetDoneAt() {
	o.DoneAt.Unset()
}

// GetIsDone returns the IsDone field value if set, zero value otherwise.
func (o *TodoDto) GetIsDone() bool {
	if o == nil || o.IsDone == nil {
		var ret bool
		return ret
	}
	return *o.IsDone
}

// GetIsDoneOk returns a tuple with the IsDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TodoDto) GetIsDoneOk() (*bool, bool) {
	if o == nil || o.IsDone == nil {
		return nil, false
	}
	return o.IsDone, true
}

// HasIsDone returns a boolean if a field has been set.
func (o *TodoDto) HasIsDone() bool {
	if o != nil && o.IsDone != nil {
		return true
	}

	return false
}

// SetIsDone gets a reference to the given bool and assigns it to the IsDone field.
func (o *TodoDto) SetIsDone(v bool) {
	o.IsDone = &v
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TodoDto) GetChildren() []TodoDto {
	if o == nil  {
		var ret []TodoDto
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TodoDto) GetChildrenOk() (*[]TodoDto, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *TodoDto) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []TodoDto and assigns it to the Children field.
func (o *TodoDto) SetChildren(v []TodoDto) {
	o.Children = v
}

func (o TodoDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.CreatorId.IsSet() {
		toSerialize["creatorId"] = o.CreatorId.Get()
	}
	if o.LastModificationTime.IsSet() {
		toSerialize["lastModificationTime"] = o.LastModificationTime.Get()
	}
	if o.LastModifierId.IsSet() {
		toSerialize["lastModifierId"] = o.LastModifierId.Get()
	}
	if o.IsDeleted != nil {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if o.DeleterId.IsSet() {
		toSerialize["deleterId"] = o.DeleterId.Get()
	}
	if o.DeletionTime.IsSet() {
		toSerialize["deletionTime"] = o.DeletionTime.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.DoneAt.IsSet() {
		toSerialize["doneAt"] = o.DoneAt.Get()
	}
	if o.IsDone != nil {
		toSerialize["isDone"] = o.IsDone
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	return json.Marshal(toSerialize)
}

type NullableTodoDto struct {
	value *TodoDto
	isSet bool
}

func (v NullableTodoDto) Get() *TodoDto {
	return v.value
}

func (v *NullableTodoDto) Set(val *TodoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTodoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTodoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTodoDto(val *TodoDto) *NullableTodoDto {
	return &NullableTodoDto{value: val, isSet: true}
}

func (v NullableTodoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTodoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



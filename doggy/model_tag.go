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

// Tag struct for Tag
type Tag struct {
	Id *string `json:"id,omitempty"`
	ExtraProperties map[string]interface{} `json:"extraProperties,omitempty"`
	ConcurrencyStamp NullableString `json:"concurrencyStamp,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	CreatorId NullableString `json:"creatorId,omitempty"`
	LastModificationTime NullableTime `json:"lastModificationTime,omitempty"`
	LastModifierId NullableString `json:"lastModifierId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	DeleterId NullableString `json:"deleterId,omitempty"`
	DeletionTime NullableTime `json:"deletionTime,omitempty"`
	Deleter *IdentityUser `json:"deleter,omitempty"`
	Creator *IdentityUser `json:"creator,omitempty"`
	LastModifier *IdentityUser `json:"lastModifier,omitempty"`
	Name NullableString `json:"name,omitempty"`
	TagCount *int32 `json:"tagCount,omitempty"`
	ParentTagId NullableString `json:"parentTagId,omitempty"`
	ParentTag *Tag `json:"parentTag,omitempty"`
	Children []Tag `json:"children,omitempty"`
	Todos []Todo `json:"todos,omitempty"`
	TodoTags []TodoTag `json:"todoTags,omitempty"`
	Items []Item `json:"items,omitempty"`
	ItemTags []ItemTag `json:"itemTags,omitempty"`
	FullPath NullableString `json:"fullPath,omitempty"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag() *Tag {
	this := Tag{}
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tag) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Tag) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Tag) SetId(v string) {
	o.Id = &v
}

// GetExtraProperties returns the ExtraProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetExtraProperties() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraProperties
}

// GetExtraPropertiesOk returns a tuple with the ExtraProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetExtraPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtraProperties == nil {
		return nil, false
	}
	return &o.ExtraProperties, true
}

// HasExtraProperties returns a boolean if a field has been set.
func (o *Tag) HasExtraProperties() bool {
	if o != nil && o.ExtraProperties != nil {
		return true
	}

	return false
}

// SetExtraProperties gets a reference to the given map[string]interface{} and assigns it to the ExtraProperties field.
func (o *Tag) SetExtraProperties(v map[string]interface{}) {
	o.ExtraProperties = v
}

// GetConcurrencyStamp returns the ConcurrencyStamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetConcurrencyStamp() string {
	if o == nil || o.ConcurrencyStamp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConcurrencyStamp.Get()
}

// GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetConcurrencyStampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConcurrencyStamp.Get(), o.ConcurrencyStamp.IsSet()
}

// HasConcurrencyStamp returns a boolean if a field has been set.
func (o *Tag) HasConcurrencyStamp() bool {
	if o != nil && o.ConcurrencyStamp.IsSet() {
		return true
	}

	return false
}

// SetConcurrencyStamp gets a reference to the given NullableString and assigns it to the ConcurrencyStamp field.
func (o *Tag) SetConcurrencyStamp(v string) {
	o.ConcurrencyStamp.Set(&v)
}
// SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil
func (o *Tag) SetConcurrencyStampNil() {
	o.ConcurrencyStamp.Set(nil)
}

// UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
func (o *Tag) UnsetConcurrencyStamp() {
	o.ConcurrencyStamp.Unset()
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Tag) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Tag) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Tag) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetCreatorId() string {
	if o == nil || o.CreatorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetCreatorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *Tag) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableString and assigns it to the CreatorId field.
func (o *Tag) SetCreatorId(v string) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *Tag) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *Tag) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetLastModificationTime() time.Time {
	if o == nil || o.LastModificationTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModificationTime.Get()
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetLastModificationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModificationTime.Get(), o.LastModificationTime.IsSet()
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *Tag) HasLastModificationTime() bool {
	if o != nil && o.LastModificationTime.IsSet() {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given NullableTime and assigns it to the LastModificationTime field.
func (o *Tag) SetLastModificationTime(v time.Time) {
	o.LastModificationTime.Set(&v)
}
// SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil
func (o *Tag) SetLastModificationTimeNil() {
	o.LastModificationTime.Set(nil)
}

// UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
func (o *Tag) UnsetLastModificationTime() {
	o.LastModificationTime.Unset()
}

// GetLastModifierId returns the LastModifierId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetLastModifierId() string {
	if o == nil || o.LastModifierId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastModifierId.Get()
}

// GetLastModifierIdOk returns a tuple with the LastModifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetLastModifierIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifierId.Get(), o.LastModifierId.IsSet()
}

// HasLastModifierId returns a boolean if a field has been set.
func (o *Tag) HasLastModifierId() bool {
	if o != nil && o.LastModifierId.IsSet() {
		return true
	}

	return false
}

// SetLastModifierId gets a reference to the given NullableString and assigns it to the LastModifierId field.
func (o *Tag) SetLastModifierId(v string) {
	o.LastModifierId.Set(&v)
}
// SetLastModifierIdNil sets the value for LastModifierId to be an explicit nil
func (o *Tag) SetLastModifierIdNil() {
	o.LastModifierId.Set(nil)
}

// UnsetLastModifierId ensures that no value is present for LastModifierId, not even an explicit nil
func (o *Tag) UnsetLastModifierId() {
	o.LastModifierId.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Tag) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Tag) HasIsDeleted() bool {
	if o != nil && o.IsDeleted != nil {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Tag) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetDeleterId returns the DeleterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetDeleterId() string {
	if o == nil || o.DeleterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeleterId.Get()
}

// GetDeleterIdOk returns a tuple with the DeleterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetDeleterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeleterId.Get(), o.DeleterId.IsSet()
}

// HasDeleterId returns a boolean if a field has been set.
func (o *Tag) HasDeleterId() bool {
	if o != nil && o.DeleterId.IsSet() {
		return true
	}

	return false
}

// SetDeleterId gets a reference to the given NullableString and assigns it to the DeleterId field.
func (o *Tag) SetDeleterId(v string) {
	o.DeleterId.Set(&v)
}
// SetDeleterIdNil sets the value for DeleterId to be an explicit nil
func (o *Tag) SetDeleterIdNil() {
	o.DeleterId.Set(nil)
}

// UnsetDeleterId ensures that no value is present for DeleterId, not even an explicit nil
func (o *Tag) UnsetDeleterId() {
	o.DeleterId.Unset()
}

// GetDeletionTime returns the DeletionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetDeletionTime() time.Time {
	if o == nil || o.DeletionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletionTime.Get()
}

// GetDeletionTimeOk returns a tuple with the DeletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetDeletionTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletionTime.Get(), o.DeletionTime.IsSet()
}

// HasDeletionTime returns a boolean if a field has been set.
func (o *Tag) HasDeletionTime() bool {
	if o != nil && o.DeletionTime.IsSet() {
		return true
	}

	return false
}

// SetDeletionTime gets a reference to the given NullableTime and assigns it to the DeletionTime field.
func (o *Tag) SetDeletionTime(v time.Time) {
	o.DeletionTime.Set(&v)
}
// SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil
func (o *Tag) SetDeletionTimeNil() {
	o.DeletionTime.Set(nil)
}

// UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
func (o *Tag) UnsetDeletionTime() {
	o.DeletionTime.Unset()
}

// GetDeleter returns the Deleter field value if set, zero value otherwise.
func (o *Tag) GetDeleter() IdentityUser {
	if o == nil || o.Deleter == nil {
		var ret IdentityUser
		return ret
	}
	return *o.Deleter
}

// GetDeleterOk returns a tuple with the Deleter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetDeleterOk() (*IdentityUser, bool) {
	if o == nil || o.Deleter == nil {
		return nil, false
	}
	return o.Deleter, true
}

// HasDeleter returns a boolean if a field has been set.
func (o *Tag) HasDeleter() bool {
	if o != nil && o.Deleter != nil {
		return true
	}

	return false
}

// SetDeleter gets a reference to the given IdentityUser and assigns it to the Deleter field.
func (o *Tag) SetDeleter(v IdentityUser) {
	o.Deleter = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Tag) GetCreator() IdentityUser {
	if o == nil || o.Creator == nil {
		var ret IdentityUser
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetCreatorOk() (*IdentityUser, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Tag) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given IdentityUser and assigns it to the Creator field.
func (o *Tag) SetCreator(v IdentityUser) {
	o.Creator = &v
}

// GetLastModifier returns the LastModifier field value if set, zero value otherwise.
func (o *Tag) GetLastModifier() IdentityUser {
	if o == nil || o.LastModifier == nil {
		var ret IdentityUser
		return ret
	}
	return *o.LastModifier
}

// GetLastModifierOk returns a tuple with the LastModifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetLastModifierOk() (*IdentityUser, bool) {
	if o == nil || o.LastModifier == nil {
		return nil, false
	}
	return o.LastModifier, true
}

// HasLastModifier returns a boolean if a field has been set.
func (o *Tag) HasLastModifier() bool {
	if o != nil && o.LastModifier != nil {
		return true
	}

	return false
}

// SetLastModifier gets a reference to the given IdentityUser and assigns it to the LastModifier field.
func (o *Tag) SetLastModifier(v IdentityUser) {
	o.LastModifier = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Tag) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Tag) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Tag) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Tag) UnsetName() {
	o.Name.Unset()
}

// GetTagCount returns the TagCount field value if set, zero value otherwise.
func (o *Tag) GetTagCount() int32 {
	if o == nil || o.TagCount == nil {
		var ret int32
		return ret
	}
	return *o.TagCount
}

// GetTagCountOk returns a tuple with the TagCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetTagCountOk() (*int32, bool) {
	if o == nil || o.TagCount == nil {
		return nil, false
	}
	return o.TagCount, true
}

// HasTagCount returns a boolean if a field has been set.
func (o *Tag) HasTagCount() bool {
	if o != nil && o.TagCount != nil {
		return true
	}

	return false
}

// SetTagCount gets a reference to the given int32 and assigns it to the TagCount field.
func (o *Tag) SetTagCount(v int32) {
	o.TagCount = &v
}

// GetParentTagId returns the ParentTagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetParentTagId() string {
	if o == nil || o.ParentTagId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentTagId.Get()
}

// GetParentTagIdOk returns a tuple with the ParentTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetParentTagIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentTagId.Get(), o.ParentTagId.IsSet()
}

// HasParentTagId returns a boolean if a field has been set.
func (o *Tag) HasParentTagId() bool {
	if o != nil && o.ParentTagId.IsSet() {
		return true
	}

	return false
}

// SetParentTagId gets a reference to the given NullableString and assigns it to the ParentTagId field.
func (o *Tag) SetParentTagId(v string) {
	o.ParentTagId.Set(&v)
}
// SetParentTagIdNil sets the value for ParentTagId to be an explicit nil
func (o *Tag) SetParentTagIdNil() {
	o.ParentTagId.Set(nil)
}

// UnsetParentTagId ensures that no value is present for ParentTagId, not even an explicit nil
func (o *Tag) UnsetParentTagId() {
	o.ParentTagId.Unset()
}

// GetParentTag returns the ParentTag field value if set, zero value otherwise.
func (o *Tag) GetParentTag() Tag {
	if o == nil || o.ParentTag == nil {
		var ret Tag
		return ret
	}
	return *o.ParentTag
}

// GetParentTagOk returns a tuple with the ParentTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetParentTagOk() (*Tag, bool) {
	if o == nil || o.ParentTag == nil {
		return nil, false
	}
	return o.ParentTag, true
}

// HasParentTag returns a boolean if a field has been set.
func (o *Tag) HasParentTag() bool {
	if o != nil && o.ParentTag != nil {
		return true
	}

	return false
}

// SetParentTag gets a reference to the given Tag and assigns it to the ParentTag field.
func (o *Tag) SetParentTag(v Tag) {
	o.ParentTag = &v
}

// GetChildren returns the Children field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetChildren() []Tag {
	if o == nil  {
		var ret []Tag
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetChildrenOk() (*[]Tag, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Tag) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []Tag and assigns it to the Children field.
func (o *Tag) SetChildren(v []Tag) {
	o.Children = v
}

// GetTodos returns the Todos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetTodos() []Todo {
	if o == nil  {
		var ret []Todo
		return ret
	}
	return o.Todos
}

// GetTodosOk returns a tuple with the Todos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetTodosOk() (*[]Todo, bool) {
	if o == nil || o.Todos == nil {
		return nil, false
	}
	return &o.Todos, true
}

// HasTodos returns a boolean if a field has been set.
func (o *Tag) HasTodos() bool {
	if o != nil && o.Todos != nil {
		return true
	}

	return false
}

// SetTodos gets a reference to the given []Todo and assigns it to the Todos field.
func (o *Tag) SetTodos(v []Todo) {
	o.Todos = v
}

// GetTodoTags returns the TodoTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetTodoTags() []TodoTag {
	if o == nil  {
		var ret []TodoTag
		return ret
	}
	return o.TodoTags
}

// GetTodoTagsOk returns a tuple with the TodoTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetTodoTagsOk() (*[]TodoTag, bool) {
	if o == nil || o.TodoTags == nil {
		return nil, false
	}
	return &o.TodoTags, true
}

// HasTodoTags returns a boolean if a field has been set.
func (o *Tag) HasTodoTags() bool {
	if o != nil && o.TodoTags != nil {
		return true
	}

	return false
}

// SetTodoTags gets a reference to the given []TodoTag and assigns it to the TodoTags field.
func (o *Tag) SetTodoTags(v []TodoTag) {
	o.TodoTags = v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetItems() []Item {
	if o == nil  {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetItemsOk() (*[]Item, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Tag) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *Tag) SetItems(v []Item) {
	o.Items = v
}

// GetItemTags returns the ItemTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetItemTags() []ItemTag {
	if o == nil  {
		var ret []ItemTag
		return ret
	}
	return o.ItemTags
}

// GetItemTagsOk returns a tuple with the ItemTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetItemTagsOk() (*[]ItemTag, bool) {
	if o == nil || o.ItemTags == nil {
		return nil, false
	}
	return &o.ItemTags, true
}

// HasItemTags returns a boolean if a field has been set.
func (o *Tag) HasItemTags() bool {
	if o != nil && o.ItemTags != nil {
		return true
	}

	return false
}

// SetItemTags gets a reference to the given []ItemTag and assigns it to the ItemTags field.
func (o *Tag) SetItemTags(v []ItemTag) {
	o.ItemTags = v
}

// GetFullPath returns the FullPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetFullPath() string {
	if o == nil || o.FullPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.FullPath.Get()
}

// GetFullPathOk returns a tuple with the FullPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetFullPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FullPath.Get(), o.FullPath.IsSet()
}

// HasFullPath returns a boolean if a field has been set.
func (o *Tag) HasFullPath() bool {
	if o != nil && o.FullPath.IsSet() {
		return true
	}

	return false
}

// SetFullPath gets a reference to the given NullableString and assigns it to the FullPath field.
func (o *Tag) SetFullPath(v string) {
	o.FullPath.Set(&v)
}
// SetFullPathNil sets the value for FullPath to be an explicit nil
func (o *Tag) SetFullPathNil() {
	o.FullPath.Set(nil)
}

// UnsetFullPath ensures that no value is present for FullPath, not even an explicit nil
func (o *Tag) UnsetFullPath() {
	o.FullPath.Unset()
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ExtraProperties != nil {
		toSerialize["extraProperties"] = o.ExtraProperties
	}
	if o.ConcurrencyStamp.IsSet() {
		toSerialize["concurrencyStamp"] = o.ConcurrencyStamp.Get()
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
	if o.Deleter != nil {
		toSerialize["deleter"] = o.Deleter
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.LastModifier != nil {
		toSerialize["lastModifier"] = o.LastModifier
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TagCount != nil {
		toSerialize["tagCount"] = o.TagCount
	}
	if o.ParentTagId.IsSet() {
		toSerialize["parentTagId"] = o.ParentTagId.Get()
	}
	if o.ParentTag != nil {
		toSerialize["parentTag"] = o.ParentTag
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.Todos != nil {
		toSerialize["todos"] = o.Todos
	}
	if o.TodoTags != nil {
		toSerialize["todoTags"] = o.TodoTags
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.ItemTags != nil {
		toSerialize["itemTags"] = o.ItemTags
	}
	if o.FullPath.IsSet() {
		toSerialize["fullPath"] = o.FullPath.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



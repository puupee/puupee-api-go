# TodoSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowCompleted** | Pointer to **bool** |  | [optional] 
**ShowDetails** | Pointer to **bool** |  | [optional] 
**OrderBy** | Pointer to [**TodoOrderBy**](TodoOrderBy.md) |  | [optional] 

## Methods

### NewTodoSettingsDto

`func NewTodoSettingsDto() *TodoSettingsDto`

NewTodoSettingsDto instantiates a new TodoSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTodoSettingsDtoWithDefaults

`func NewTodoSettingsDtoWithDefaults() *TodoSettingsDto`

NewTodoSettingsDtoWithDefaults instantiates a new TodoSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowCompleted

`func (o *TodoSettingsDto) GetShowCompleted() bool`

GetShowCompleted returns the ShowCompleted field if non-nil, zero value otherwise.

### GetShowCompletedOk

`func (o *TodoSettingsDto) GetShowCompletedOk() (*bool, bool)`

GetShowCompletedOk returns a tuple with the ShowCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCompleted

`func (o *TodoSettingsDto) SetShowCompleted(v bool)`

SetShowCompleted sets ShowCompleted field to given value.

### HasShowCompleted

`func (o *TodoSettingsDto) HasShowCompleted() bool`

HasShowCompleted returns a boolean if a field has been set.

### GetShowDetails

`func (o *TodoSettingsDto) GetShowDetails() bool`

GetShowDetails returns the ShowDetails field if non-nil, zero value otherwise.

### GetShowDetailsOk

`func (o *TodoSettingsDto) GetShowDetailsOk() (*bool, bool)`

GetShowDetailsOk returns a tuple with the ShowDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDetails

`func (o *TodoSettingsDto) SetShowDetails(v bool)`

SetShowDetails sets ShowDetails field to given value.

### HasShowDetails

`func (o *TodoSettingsDto) HasShowDetails() bool`

HasShowDetails returns a boolean if a field has been set.

### GetOrderBy

`func (o *TodoSettingsDto) GetOrderBy() TodoOrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *TodoSettingsDto) GetOrderByOk() (*TodoOrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *TodoSettingsDto) SetOrderBy(v TodoOrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *TodoSettingsDto) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



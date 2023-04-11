/*
lobstr API v1

Lobstr API is an easy-to-implement API to launch already-made data crawlers, handle accounts and schedules, and collect results.  ## Authentication  All requests must include the `Authorization` headers. You can obtain these from the settings menu.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TaskDeactivated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskDeactivated{}

// TaskDeactivated struct for TaskDeactivated
type TaskDeactivated struct {
	Id string `json:"id"`
	Object *string `json:"object,omitempty"`
	// date of creation (ISO8601 format)
	CreatedAt string `json:"created_at"`
	IsActive bool `json:"is_active"`
	Params TaskCommonParams `json:"params"`
	Status *string `json:"status,omitempty"`
}

// NewTaskDeactivated instantiates a new TaskDeactivated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskDeactivated(id string, createdAt string, isActive bool, params TaskCommonParams) *TaskDeactivated {
	this := TaskDeactivated{}
	this.Id = id
	this.CreatedAt = createdAt
	this.IsActive = isActive
	this.Params = params
	return &this
}

// NewTaskDeactivatedWithDefaults instantiates a new TaskDeactivated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskDeactivatedWithDefaults() *TaskDeactivated {
	this := TaskDeactivated{}
	return &this
}

// GetId returns the Id field value
func (o *TaskDeactivated) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskDeactivated) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskDeactivated) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *TaskDeactivated) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDeactivated) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *TaskDeactivated) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *TaskDeactivated) SetObject(v string) {
	o.Object = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TaskDeactivated) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskDeactivated) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TaskDeactivated) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetIsActive returns the IsActive field value
func (o *TaskDeactivated) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *TaskDeactivated) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *TaskDeactivated) SetIsActive(v bool) {
	o.IsActive = v
}

// GetParams returns the Params field value
func (o *TaskDeactivated) GetParams() TaskCommonParams {
	if o == nil {
		var ret TaskCommonParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *TaskDeactivated) GetParamsOk() (*TaskCommonParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *TaskDeactivated) SetParams(v TaskCommonParams) {
	o.Params = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TaskDeactivated) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDeactivated) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TaskDeactivated) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TaskDeactivated) SetStatus(v string) {
	o.Status = &v
}

func (o TaskDeactivated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskDeactivated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["is_active"] = o.IsActive
	toSerialize["params"] = o.Params
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableTaskDeactivated struct {
	value *TaskDeactivated
	isSet bool
}

func (v NullableTaskDeactivated) Get() *TaskDeactivated {
	return v.value
}

func (v *NullableTaskDeactivated) Set(val *TaskDeactivated) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskDeactivated) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskDeactivated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskDeactivated(val *TaskDeactivated) *NullableTaskDeactivated {
	return &NullableTaskDeactivated{value: val, isSet: true}
}

func (v NullableTaskDeactivated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskDeactivated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



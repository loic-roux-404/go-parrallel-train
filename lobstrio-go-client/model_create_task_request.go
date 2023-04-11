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

// checks if the CreateTaskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTaskRequest{}

// CreateTaskRequest struct for CreateTaskRequest
type CreateTaskRequest struct {
	Cluster string `json:"cluster"`
	Tasks []interface{} `json:"tasks,omitempty"`
}

// NewCreateTaskRequest instantiates a new CreateTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTaskRequest(cluster string) *CreateTaskRequest {
	this := CreateTaskRequest{}
	this.Cluster = cluster
	return &this
}

// NewCreateTaskRequestWithDefaults instantiates a new CreateTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTaskRequestWithDefaults() *CreateTaskRequest {
	this := CreateTaskRequest{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *CreateTaskRequest) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *CreateTaskRequest) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *CreateTaskRequest) SetCluster(v string) {
	o.Cluster = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *CreateTaskRequest) GetTasks() []interface{} {
	if o == nil || IsNil(o.Tasks) {
		var ret []interface{}
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTaskRequest) GetTasksOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *CreateTaskRequest) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []interface{} and assigns it to the Tasks field.
func (o *CreateTaskRequest) SetTasks(v []interface{}) {
	o.Tasks = v
}

func (o CreateTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

type NullableCreateTaskRequest struct {
	value *CreateTaskRequest
	isSet bool
}

func (v NullableCreateTaskRequest) Get() *CreateTaskRequest {
	return v.value
}

func (v *NullableCreateTaskRequest) Set(val *CreateTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTaskRequest(val *CreateTaskRequest) *NullableCreateTaskRequest {
	return &NullableCreateTaskRequest{value: val, isSet: true}
}

func (v NullableCreateTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


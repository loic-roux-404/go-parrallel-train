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

// checks if the CreateTaskParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTaskParams{}

// CreateTaskParams struct for CreateTaskParams
type CreateTaskParams struct {
	Params TaskCommonParams `json:"params"`
}

// NewCreateTaskParams instantiates a new CreateTaskParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTaskParams(params TaskCommonParams) *CreateTaskParams {
	this := CreateTaskParams{}
	this.Params = params
	return &this
}

// NewCreateTaskParamsWithDefaults instantiates a new CreateTaskParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTaskParamsWithDefaults() *CreateTaskParams {
	this := CreateTaskParams{}
	return &this
}

// GetParams returns the Params field value
func (o *CreateTaskParams) GetParams() TaskCommonParams {
	if o == nil {
		var ret TaskCommonParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CreateTaskParams) GetParamsOk() (*TaskCommonParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CreateTaskParams) SetParams(v TaskCommonParams) {
	o.Params = v
}

func (o CreateTaskParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTaskParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableCreateTaskParams struct {
	value *CreateTaskParams
	isSet bool
}

func (v NullableCreateTaskParams) Get() *CreateTaskParams {
	return v.value
}

func (v *NullableCreateTaskParams) Set(val *CreateTaskParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTaskParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTaskParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTaskParams(val *CreateTaskParams) *NullableCreateTaskParams {
	return &NullableCreateTaskParams{value: val, isSet: true}
}

func (v NullableCreateTaskParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTaskParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



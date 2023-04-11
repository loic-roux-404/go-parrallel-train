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

// checks if the ScheduleDeleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleDeleted{}

// ScheduleDeleted struct for ScheduleDeleted
type ScheduleDeleted struct {
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
}

// NewScheduleDeleted instantiates a new ScheduleDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleDeleted() *ScheduleDeleted {
	this := ScheduleDeleted{}
	return &this
}

// NewScheduleDeletedWithDefaults instantiates a new ScheduleDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleDeletedWithDefaults() *ScheduleDeleted {
	this := ScheduleDeleted{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScheduleDeleted) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleDeleted) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScheduleDeleted) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScheduleDeleted) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ScheduleDeleted) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleDeleted) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ScheduleDeleted) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ScheduleDeleted) SetObject(v string) {
	o.Object = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *ScheduleDeleted) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleDeleted) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *ScheduleDeleted) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *ScheduleDeleted) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o ScheduleDeleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleDeleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

type NullableScheduleDeleted struct {
	value *ScheduleDeleted
	isSet bool
}

func (v NullableScheduleDeleted) Get() *ScheduleDeleted {
	return v.value
}

func (v *NullableScheduleDeleted) Set(val *ScheduleDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleDeleted(val *ScheduleDeleted) *NullableScheduleDeleted {
	return &NullableScheduleDeleted{value: val, isSet: true}
}

func (v NullableScheduleDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


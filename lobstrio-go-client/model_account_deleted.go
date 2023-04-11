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

// checks if the AccountDeleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountDeleted{}

// AccountDeleted struct for AccountDeleted
type AccountDeleted struct {
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
}

// NewAccountDeleted instantiates a new AccountDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDeleted() *AccountDeleted {
	this := AccountDeleted{}
	return &this
}

// NewAccountDeletedWithDefaults instantiates a new AccountDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDeletedWithDefaults() *AccountDeleted {
	this := AccountDeleted{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountDeleted) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDeleted) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountDeleted) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountDeleted) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *AccountDeleted) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDeleted) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *AccountDeleted) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *AccountDeleted) SetObject(v string) {
	o.Object = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *AccountDeleted) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDeleted) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *AccountDeleted) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *AccountDeleted) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o AccountDeleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountDeleted) ToMap() (map[string]interface{}, error) {
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

type NullableAccountDeleted struct {
	value *AccountDeleted
	isSet bool
}

func (v NullableAccountDeleted) Get() *AccountDeleted {
	return v.value
}

func (v *NullableAccountDeleted) Set(val *AccountDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDeleted(val *AccountDeleted) *NullableAccountDeleted {
	return &NullableAccountDeleted{value: val, isSet: true}
}

func (v NullableAccountDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



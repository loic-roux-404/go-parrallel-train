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

// checks if the ListClusters200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListClusters200Response{}

// ListClusters200Response struct for ListClusters200Response
type ListClusters200Response struct {
	Count *float32 `json:"count,omitempty"`
	// page number
	Page float32 `json:"page"`
	// number of items returned max
	Limit float32 `json:"limit"`
	Data []interface{} `json:"data,omitempty"`
}

// NewListClusters200Response instantiates a new ListClusters200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClusters200Response(page float32, limit float32) *ListClusters200Response {
	this := ListClusters200Response{}
	this.Page = page
	this.Limit = limit
	return &this
}

// NewListClusters200ResponseWithDefaults instantiates a new ListClusters200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClusters200ResponseWithDefaults() *ListClusters200Response {
	this := ListClusters200Response{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListClusters200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusters200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListClusters200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ListClusters200Response) SetCount(v float32) {
	o.Count = &v
}

// GetPage returns the Page field value
func (o *ListClusters200Response) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ListClusters200Response) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ListClusters200Response) SetPage(v float32) {
	o.Page = v
}

// GetLimit returns the Limit field value
func (o *ListClusters200Response) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListClusters200Response) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListClusters200Response) SetLimit(v float32) {
	o.Limit = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListClusters200Response) GetData() []interface{} {
	if o == nil || IsNil(o.Data) {
		var ret []interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusters200Response) GetDataOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListClusters200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []interface{} and assigns it to the Data field.
func (o *ListClusters200Response) SetData(v []interface{}) {
	o.Data = v
}

func (o ListClusters200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListClusters200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["page"] = o.Page
	toSerialize["limit"] = o.Limit
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListClusters200Response struct {
	value *ListClusters200Response
	isSet bool
}

func (v NullableListClusters200Response) Get() *ListClusters200Response {
	return v.value
}

func (v *NullableListClusters200Response) Set(val *ListClusters200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListClusters200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListClusters200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListClusters200Response(val *ListClusters200Response) *NullableListClusters200Response {
	return &NullableListClusters200Response{value: val, isSet: true}
}

func (v NullableListClusters200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListClusters200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


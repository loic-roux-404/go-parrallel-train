/*
Lobstr

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lobstrio

import (
	"encoding/json"
)

// checks if the UpdateClusterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateClusterRequest{}

// UpdateClusterRequest struct for UpdateClusterRequest
type UpdateClusterRequest struct {
	Name string `json:"name"`
	Concurrency int32 `json:"concurrency"`
	ExportUniqueResults bool `json:"export_unique_results"`
	NoLineBreaks bool `json:"no_line_breaks"`
	ToComplete bool `json:"to_complete"`
	Params UpdateClusterRequestParams `json:"params"`
	Accounts NullableString `json:"accounts"`
	RunNotify string `json:"run_notify"`
}

// NewUpdateClusterRequest instantiates a new UpdateClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateClusterRequest(name string, concurrency int32, exportUniqueResults bool, noLineBreaks bool, toComplete bool, params UpdateClusterRequestParams, accounts NullableString, runNotify string) *UpdateClusterRequest {
	this := UpdateClusterRequest{}
	this.Name = name
	this.Concurrency = concurrency
	this.ExportUniqueResults = exportUniqueResults
	this.NoLineBreaks = noLineBreaks
	this.ToComplete = toComplete
	this.Params = params
	this.Accounts = accounts
	this.RunNotify = runNotify
	return &this
}

// NewUpdateClusterRequestWithDefaults instantiates a new UpdateClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateClusterRequestWithDefaults() *UpdateClusterRequest {
	this := UpdateClusterRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateClusterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateClusterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateClusterRequest) SetName(v string) {
	o.Name = v
}

// GetConcurrency returns the Concurrency field value
func (o *UpdateClusterRequest) GetConcurrency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Concurrency
}

// GetConcurrencyOk returns a tuple with the Concurrency field value
// and a boolean to check if the value has been set.
func (o *UpdateClusterRequest) GetConcurrencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Concurrency, true
}

// SetConcurrency sets field value
func (o *UpdateClusterRequest) SetConcurrency(v int32) {
	o.Concurrency = v
}

// GetExportUniqueResults returns the ExportUniqueResults field value
func (o *UpdateClusterRequest) GetExportUniqueResults() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExportUniqueResults
}

// GetExportUniqueResultsOk returns a tuple with the ExportUniqueResults field value
// and a boolean to check if the value has been set.
func (o *UpdateClusterRequest) GetExportUniqueResultsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportUniqueResults, true
}

// SetExportUniqueResults sets field value
func (o *UpdateClusterRequest) SetExportUniqueResults(v bool) {
	o.ExportUniqueResults = v
}

// GetNoLineBreaks returns the NoLineBreaks field value
func (o *UpdateClusterRequest) GetNoLineBreaks() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NoLineBreaks
}

// GetNoLineBreaksOk returns a tuple with the NoLineBreaks field value
// and a boolean to check if the value has been set.
func (o *UpdateClusterRequest) GetNoLineBreaksOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoLineBreaks, true
}

// SetNoLineBreaks sets field value
func (o *UpdateClusterRequest) SetNoLineBreaks(v bool) {
	o.NoLineBreaks = v
}

// GetToComplete returns the ToComplete field value
func (o *UpdateClusterRequest) GetToComplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ToComplete
}

// GetToCompleteOk returns a tuple with the ToComplete field value
// and a boolean to check if the value has been set.
func (o *UpdateClusterRequest) GetToCompleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToComplete, true
}

// SetToComplete sets field value
func (o *UpdateClusterRequest) SetToComplete(v bool) {
	o.ToComplete = v
}

// GetParams returns the Params field value
func (o *UpdateClusterRequest) GetParams() UpdateClusterRequestParams {
	if o == nil {
		var ret UpdateClusterRequestParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpdateClusterRequest) GetParamsOk() (*UpdateClusterRequestParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *UpdateClusterRequest) SetParams(v UpdateClusterRequestParams) {
	o.Params = v
}

// GetAccounts returns the Accounts field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateClusterRequest) GetAccounts() string {
	if o == nil || o.Accounts.Get() == nil {
		var ret string
		return ret
	}

	return *o.Accounts.Get()
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateClusterRequest) GetAccountsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts.Get(), o.Accounts.IsSet()
}

// SetAccounts sets field value
func (o *UpdateClusterRequest) SetAccounts(v string) {
	o.Accounts.Set(&v)
}

// GetRunNotify returns the RunNotify field value
func (o *UpdateClusterRequest) GetRunNotify() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunNotify
}

// GetRunNotifyOk returns a tuple with the RunNotify field value
// and a boolean to check if the value has been set.
func (o *UpdateClusterRequest) GetRunNotifyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunNotify, true
}

// SetRunNotify sets field value
func (o *UpdateClusterRequest) SetRunNotify(v string) {
	o.RunNotify = v
}

func (o UpdateClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateClusterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["concurrency"] = o.Concurrency
	toSerialize["export_unique_results"] = o.ExportUniqueResults
	toSerialize["no_line_breaks"] = o.NoLineBreaks
	toSerialize["to_complete"] = o.ToComplete
	toSerialize["params"] = o.Params
	toSerialize["accounts"] = o.Accounts.Get()
	toSerialize["run_notify"] = o.RunNotify
	return toSerialize, nil
}

type NullableUpdateClusterRequest struct {
	value *UpdateClusterRequest
	isSet bool
}

func (v NullableUpdateClusterRequest) Get() *UpdateClusterRequest {
	return v.value
}

func (v *NullableUpdateClusterRequest) Set(val *UpdateClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateClusterRequest(val *UpdateClusterRequest) *NullableUpdateClusterRequest {
	return &NullableUpdateClusterRequest{value: val, isSet: true}
}

func (v NullableUpdateClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



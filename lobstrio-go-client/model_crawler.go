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

// checks if the Crawler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Crawler{}

// Crawler struct for Crawler
type Crawler struct {
	// unique id
	Id string `json:"id"`
	Object *string `json:"object,omitempty"`
	// date of creation (ISO8601 format)
	CreatedAt string `json:"created_at"`
	// description
	Description *string `json:"description,omitempty"`
	// public status
	IsPublic bool `json:"is_public"`
	// labels
	Labels []interface{} `json:"labels,omitempty"`
	// total number of concurrent threads allowed within one cluster
	MaxConcurrency float32 `json:"max_concurrency"`
	// public name
	Name string `json:"name"`
	// position within the list of crawlers
	Rank float32 `json:"rank"`
	// unique slug
	Slug string `json:"slug"`
	// date of update (ISO8601 format)
	UpdatedAt string `json:"updated_at"`
}

// NewCrawler instantiates a new Crawler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrawler(id string, createdAt string, isPublic bool, maxConcurrency float32, name string, rank float32, slug string, updatedAt string) *Crawler {
	this := Crawler{}
	this.Id = id
	this.CreatedAt = createdAt
	this.IsPublic = isPublic
	this.MaxConcurrency = maxConcurrency
	this.Name = name
	this.Rank = rank
	this.Slug = slug
	this.UpdatedAt = updatedAt
	return &this
}

// NewCrawlerWithDefaults instantiates a new Crawler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrawlerWithDefaults() *Crawler {
	this := Crawler{}
	return &this
}

// GetId returns the Id field value
func (o *Crawler) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Crawler) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Crawler) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Crawler) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Crawler) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Crawler) SetObject(v string) {
	o.Object = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Crawler) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Crawler) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Crawler) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Crawler) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Crawler) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Crawler) SetDescription(v string) {
	o.Description = &v
}

// GetIsPublic returns the IsPublic field value
func (o *Crawler) GetIsPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
func (o *Crawler) GetIsPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublic, true
}

// SetIsPublic sets field value
func (o *Crawler) SetIsPublic(v bool) {
	o.IsPublic = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Crawler) GetLabels() []interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret []interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetLabelsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Crawler) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []interface{} and assigns it to the Labels field.
func (o *Crawler) SetLabels(v []interface{}) {
	o.Labels = v
}

// GetMaxConcurrency returns the MaxConcurrency field value
func (o *Crawler) GetMaxConcurrency() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxConcurrency
}

// GetMaxConcurrencyOk returns a tuple with the MaxConcurrency field value
// and a boolean to check if the value has been set.
func (o *Crawler) GetMaxConcurrencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxConcurrency, true
}

// SetMaxConcurrency sets field value
func (o *Crawler) SetMaxConcurrency(v float32) {
	o.MaxConcurrency = v
}

// GetName returns the Name field value
func (o *Crawler) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Crawler) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Crawler) SetName(v string) {
	o.Name = v
}

// GetRank returns the Rank field value
func (o *Crawler) GetRank() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *Crawler) GetRankOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *Crawler) SetRank(v float32) {
	o.Rank = v
}

// GetSlug returns the Slug field value
func (o *Crawler) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Crawler) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Crawler) SetSlug(v string) {
	o.Slug = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Crawler) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Crawler) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Crawler) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o Crawler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Crawler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["is_public"] = o.IsPublic
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["max_concurrency"] = o.MaxConcurrency
	toSerialize["name"] = o.Name
	toSerialize["rank"] = o.Rank
	toSerialize["slug"] = o.Slug
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableCrawler struct {
	value *Crawler
	isSet bool
}

func (v NullableCrawler) Get() *Crawler {
	return v.value
}

func (v *NullableCrawler) Set(val *Crawler) {
	v.value = val
	v.isSet = true
}

func (v NullableCrawler) IsSet() bool {
	return v.isSet
}

func (v *NullableCrawler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrawler(val *Crawler) *NullableCrawler {
	return &NullableCrawler{value: val, isSet: true}
}

func (v NullableCrawler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrawler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


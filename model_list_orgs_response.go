/*
Steampipe Cloud

Steampipe Cloud is a hosted version of Steampipe (https://steampipe.io), an open source tool to instantly query your cloud services (e.g. AWS, Azure, GCP and more) with SQL. No DB required.

API version: {{OPEN_API_VERSION}}
Contact: help@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package steampipecloud

import (
	"encoding/json"
)

// ListOrgsResponse struct for ListOrgsResponse
type ListOrgsResponse struct {
	Items     *[]Org  `json:"items,omitempty"`
	NextToken *string `json:"next_token,omitempty"`
}

// NewListOrgsResponse instantiates a new ListOrgsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrgsResponse() *ListOrgsResponse {
	this := ListOrgsResponse{}
	return &this
}

// NewListOrgsResponseWithDefaults instantiates a new ListOrgsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrgsResponseWithDefaults() *ListOrgsResponse {
	this := ListOrgsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListOrgsResponse) GetItems() []Org {
	if o == nil || o.Items == nil {
		var ret []Org
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrgsResponse) GetItemsOk() (*[]Org, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListOrgsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Org and assigns it to the Items field.
func (o *ListOrgsResponse) SetItems(v []Org) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListOrgsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrgsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListOrgsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListOrgsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListOrgsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListOrgsResponse struct {
	value *ListOrgsResponse
	isSet bool
}

func (v NullableListOrgsResponse) Get() *ListOrgsResponse {
	return v.value
}

func (v *NullableListOrgsResponse) Set(val *ListOrgsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrgsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrgsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrgsResponse(val *ListOrgsResponse) *NullableListOrgsResponse {
	return &NullableListOrgsResponse{value: val, isSet: true}
}

func (v NullableListOrgsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrgsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

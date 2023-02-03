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

// ListPipelinesResponse struct for ListPipelinesResponse
type ListPipelinesResponse struct {
	Items     *[]Pipeline `json:"items,omitempty"`
	NextToken *string     `json:"next_token,omitempty"`
}

// NewListPipelinesResponse instantiates a new ListPipelinesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPipelinesResponse() *ListPipelinesResponse {
	this := ListPipelinesResponse{}
	return &this
}

// NewListPipelinesResponseWithDefaults instantiates a new ListPipelinesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPipelinesResponseWithDefaults() *ListPipelinesResponse {
	this := ListPipelinesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListPipelinesResponse) GetItems() []Pipeline {
	if o == nil || o.Items == nil {
		var ret []Pipeline
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPipelinesResponse) GetItemsOk() (*[]Pipeline, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListPipelinesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Pipeline and assigns it to the Items field.
func (o *ListPipelinesResponse) SetItems(v []Pipeline) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListPipelinesResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPipelinesResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListPipelinesResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListPipelinesResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListPipelinesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListPipelinesResponse struct {
	value *ListPipelinesResponse
	isSet bool
}

func (v NullableListPipelinesResponse) Get() *ListPipelinesResponse {
	return v.value
}

func (v *NullableListPipelinesResponse) Set(val *ListPipelinesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPipelinesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPipelinesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPipelinesResponse(val *ListPipelinesResponse) *NullableListPipelinesResponse {
	return &NullableListPipelinesResponse{value: val, isSet: true}
}

func (v NullableListPipelinesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPipelinesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
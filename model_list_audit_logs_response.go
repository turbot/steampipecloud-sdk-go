/*
Steampipe Cloud

Steampipe Cloud is a hosted version of Steampipe (https://steampipe.io), an open source tool to instantly query your cloud services (e.g. AWS, Azure, GCP and more) with SQL. No DB required.

API version: 1.0
Contact: help@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package steampipecloud

import (
	"encoding/json"
)

// ListAuditLogsResponse struct for ListAuditLogsResponse
type ListAuditLogsResponse struct {
	Items     *[]AuditRecord `json:"items,omitempty"`
	NextToken *string        `json:"next_token,omitempty"`
}

// NewListAuditLogsResponse instantiates a new ListAuditLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAuditLogsResponse() *ListAuditLogsResponse {
	this := ListAuditLogsResponse{}
	return &this
}

// NewListAuditLogsResponseWithDefaults instantiates a new ListAuditLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAuditLogsResponseWithDefaults() *ListAuditLogsResponse {
	this := ListAuditLogsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListAuditLogsResponse) GetItems() []AuditRecord {
	if o == nil || o.Items == nil {
		var ret []AuditRecord
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuditLogsResponse) GetItemsOk() (*[]AuditRecord, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListAuditLogsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AuditRecord and assigns it to the Items field.
func (o *ListAuditLogsResponse) SetItems(v []AuditRecord) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListAuditLogsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuditLogsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListAuditLogsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListAuditLogsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListAuditLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListAuditLogsResponse struct {
	value *ListAuditLogsResponse
	isSet bool
}

func (v NullableListAuditLogsResponse) Get() *ListAuditLogsResponse {
	return v.value
}

func (v *NullableListAuditLogsResponse) Set(val *ListAuditLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAuditLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAuditLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAuditLogsResponse(val *ListAuditLogsResponse) *NullableListAuditLogsResponse {
	return &NullableListAuditLogsResponse{value: val, isSet: true}
}

func (v NullableListAuditLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAuditLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

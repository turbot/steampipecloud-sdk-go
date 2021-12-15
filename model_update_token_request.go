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

// UpdateTokenRequest struct for UpdateTokenRequest
type UpdateTokenRequest struct {
	Status string `json:"status"`
}

// NewUpdateTokenRequest instantiates a new UpdateTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTokenRequest(status string) *UpdateTokenRequest {
	this := UpdateTokenRequest{}
	this.Status = status
	return &this
}

// NewUpdateTokenRequestWithDefaults instantiates a new UpdateTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTokenRequestWithDefaults() *UpdateTokenRequest {
	this := UpdateTokenRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *UpdateTokenRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateTokenRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateTokenRequest) SetStatus(v string) {
	o.Status = v
}

func (o UpdateTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTokenRequest struct {
	value *UpdateTokenRequest
	isSet bool
}

func (v NullableUpdateTokenRequest) Get() *UpdateTokenRequest {
	return v.value
}

func (v *NullableUpdateTokenRequest) Set(val *UpdateTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTokenRequest(val *UpdateTokenRequest) *NullableUpdateTokenRequest {
	return &NullableUpdateTokenRequest{value: val, isSet: true}
}

func (v NullableUpdateTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

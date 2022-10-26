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

// CreateWorkspaceRequest struct for CreateWorkspaceRequest
type CreateWorkspaceRequest struct {
	Handle string `json:"handle"`
}

// NewCreateWorkspaceRequest instantiates a new CreateWorkspaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceRequest(handle string) *CreateWorkspaceRequest {
	this := CreateWorkspaceRequest{}
	this.Handle = handle
	return &this
}

// NewCreateWorkspaceRequestWithDefaults instantiates a new CreateWorkspaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceRequestWithDefaults() *CreateWorkspaceRequest {
	this := CreateWorkspaceRequest{}
	return &this
}

// GetHandle returns the Handle field value
func (o *CreateWorkspaceRequest) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequest) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *CreateWorkspaceRequest) SetHandle(v string) {
	o.Handle = v
}

func (o CreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handle"] = o.Handle
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWorkspaceRequest struct {
	value *CreateWorkspaceRequest
	isSet bool
}

func (v NullableCreateWorkspaceRequest) Get() *CreateWorkspaceRequest {
	return v.value
}

func (v *NullableCreateWorkspaceRequest) Set(val *CreateWorkspaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceRequest(val *CreateWorkspaceRequest) *NullableCreateWorkspaceRequest {
	return &NullableCreateWorkspaceRequest{value: val, isSet: true}
}

func (v NullableCreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

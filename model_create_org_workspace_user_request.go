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

// CreateOrgWorkspaceUserRequest struct for CreateOrgWorkspaceUserRequest
type CreateOrgWorkspaceUserRequest struct {
	Handle string `json:"handle"`
	Role   string `json:"role"`
}

// NewCreateOrgWorkspaceUserRequest instantiates a new CreateOrgWorkspaceUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrgWorkspaceUserRequest(handle string, role string) *CreateOrgWorkspaceUserRequest {
	this := CreateOrgWorkspaceUserRequest{}
	this.Handle = handle
	this.Role = role
	return &this
}

// NewCreateOrgWorkspaceUserRequestWithDefaults instantiates a new CreateOrgWorkspaceUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrgWorkspaceUserRequestWithDefaults() *CreateOrgWorkspaceUserRequest {
	this := CreateOrgWorkspaceUserRequest{}
	return &this
}

// GetHandle returns the Handle field value
func (o *CreateOrgWorkspaceUserRequest) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *CreateOrgWorkspaceUserRequest) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *CreateOrgWorkspaceUserRequest) SetHandle(v string) {
	o.Handle = v
}

// GetRole returns the Role field value
func (o *CreateOrgWorkspaceUserRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *CreateOrgWorkspaceUserRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *CreateOrgWorkspaceUserRequest) SetRole(v string) {
	o.Role = v
}

func (o CreateOrgWorkspaceUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrgWorkspaceUserRequest struct {
	value *CreateOrgWorkspaceUserRequest
	isSet bool
}

func (v NullableCreateOrgWorkspaceUserRequest) Get() *CreateOrgWorkspaceUserRequest {
	return v.value
}

func (v *NullableCreateOrgWorkspaceUserRequest) Set(val *CreateOrgWorkspaceUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrgWorkspaceUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrgWorkspaceUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrgWorkspaceUserRequest(val *CreateOrgWorkspaceUserRequest) *NullableCreateOrgWorkspaceUserRequest {
	return &NullableCreateOrgWorkspaceUserRequest{value: val, isSet: true}
}

func (v NullableCreateOrgWorkspaceUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrgWorkspaceUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

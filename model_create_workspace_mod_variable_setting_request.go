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

// CreateWorkspaceModVariableSettingRequest struct for CreateWorkspaceModVariableSettingRequest
type CreateWorkspaceModVariableSettingRequest struct {
	Name    string `json:"name"`
	Setting string `json:"setting"`
}

// NewCreateWorkspaceModVariableSettingRequest instantiates a new CreateWorkspaceModVariableSettingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceModVariableSettingRequest(name string, setting string) *CreateWorkspaceModVariableSettingRequest {
	this := CreateWorkspaceModVariableSettingRequest{}
	this.Name = name
	this.Setting = setting
	return &this
}

// NewCreateWorkspaceModVariableSettingRequestWithDefaults instantiates a new CreateWorkspaceModVariableSettingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceModVariableSettingRequestWithDefaults() *CreateWorkspaceModVariableSettingRequest {
	this := CreateWorkspaceModVariableSettingRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateWorkspaceModVariableSettingRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceModVariableSettingRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWorkspaceModVariableSettingRequest) SetName(v string) {
	o.Name = v
}

// GetSetting returns the Setting field value
func (o *CreateWorkspaceModVariableSettingRequest) GetSetting() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Setting
}

// GetSettingOk returns a tuple with the Setting field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceModVariableSettingRequest) GetSettingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Setting, true
}

// SetSetting sets field value
func (o *CreateWorkspaceModVariableSettingRequest) SetSetting(v string) {
	o.Setting = v
}

func (o CreateWorkspaceModVariableSettingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["setting"] = o.Setting
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWorkspaceModVariableSettingRequest struct {
	value *CreateWorkspaceModVariableSettingRequest
	isSet bool
}

func (v NullableCreateWorkspaceModVariableSettingRequest) Get() *CreateWorkspaceModVariableSettingRequest {
	return v.value
}

func (v *NullableCreateWorkspaceModVariableSettingRequest) Set(val *CreateWorkspaceModVariableSettingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceModVariableSettingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceModVariableSettingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceModVariableSettingRequest(val *CreateWorkspaceModVariableSettingRequest) *NullableCreateWorkspaceModVariableSettingRequest {
	return &NullableCreateWorkspaceModVariableSettingRequest{value: val, isSet: true}
}

func (v NullableCreateWorkspaceModVariableSettingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceModVariableSettingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

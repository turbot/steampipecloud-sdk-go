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

// UpdateWorkspaceModVariableSettingRequest struct for UpdateWorkspaceModVariableSettingRequest
type UpdateWorkspaceModVariableSettingRequest struct {
	Setting string `json:"setting"`
}

// NewUpdateWorkspaceModVariableSettingRequest instantiates a new UpdateWorkspaceModVariableSettingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWorkspaceModVariableSettingRequest(setting string) *UpdateWorkspaceModVariableSettingRequest {
	this := UpdateWorkspaceModVariableSettingRequest{}
	this.Setting = setting
	return &this
}

// NewUpdateWorkspaceModVariableSettingRequestWithDefaults instantiates a new UpdateWorkspaceModVariableSettingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWorkspaceModVariableSettingRequestWithDefaults() *UpdateWorkspaceModVariableSettingRequest {
	this := UpdateWorkspaceModVariableSettingRequest{}
	return &this
}

// GetSetting returns the Setting field value
func (o *UpdateWorkspaceModVariableSettingRequest) GetSetting() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Setting
}

// GetSettingOk returns a tuple with the Setting field value
// and a boolean to check if the value has been set.
func (o *UpdateWorkspaceModVariableSettingRequest) GetSettingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Setting, true
}

// SetSetting sets field value
func (o *UpdateWorkspaceModVariableSettingRequest) SetSetting(v string) {
	o.Setting = v
}

func (o UpdateWorkspaceModVariableSettingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["setting"] = o.Setting
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWorkspaceModVariableSettingRequest struct {
	value *UpdateWorkspaceModVariableSettingRequest
	isSet bool
}

func (v NullableUpdateWorkspaceModVariableSettingRequest) Get() *UpdateWorkspaceModVariableSettingRequest {
	return v.value
}

func (v *NullableUpdateWorkspaceModVariableSettingRequest) Set(val *UpdateWorkspaceModVariableSettingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWorkspaceModVariableSettingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWorkspaceModVariableSettingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWorkspaceModVariableSettingRequest(val *UpdateWorkspaceModVariableSettingRequest) *NullableUpdateWorkspaceModVariableSettingRequest {
	return &NullableUpdateWorkspaceModVariableSettingRequest{value: val, isSet: true}
}

func (v NullableUpdateWorkspaceModVariableSettingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWorkspaceModVariableSettingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

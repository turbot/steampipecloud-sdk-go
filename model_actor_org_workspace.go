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

// ActorOrgWorkspace struct for ActorOrgWorkspace
type ActorOrgWorkspace struct {
	// The role of the actor in the workspace.
	Role      string     `json:"role"`
	Workspace *Workspace `json:"workspace,omitempty"`
}

// NewActorOrgWorkspace instantiates a new ActorOrgWorkspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActorOrgWorkspace(role string) *ActorOrgWorkspace {
	this := ActorOrgWorkspace{}
	this.Role = role
	return &this
}

// NewActorOrgWorkspaceWithDefaults instantiates a new ActorOrgWorkspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActorOrgWorkspaceWithDefaults() *ActorOrgWorkspace {
	this := ActorOrgWorkspace{}
	return &this
}

// GetRole returns the Role field value
func (o *ActorOrgWorkspace) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ActorOrgWorkspace) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ActorOrgWorkspace) SetRole(v string) {
	o.Role = v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *ActorOrgWorkspace) GetWorkspace() Workspace {
	if o == nil || o.Workspace == nil {
		var ret Workspace
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorOrgWorkspace) GetWorkspaceOk() (*Workspace, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *ActorOrgWorkspace) HasWorkspace() bool {
	if o != nil && o.Workspace != nil {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given Workspace and assigns it to the Workspace field.
func (o *ActorOrgWorkspace) SetWorkspace(v Workspace) {
	o.Workspace = &v
}

func (o ActorOrgWorkspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role"] = o.Role
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	return json.Marshal(toSerialize)
}

type NullableActorOrgWorkspace struct {
	value *ActorOrgWorkspace
	isSet bool
}

func (v NullableActorOrgWorkspace) Get() *ActorOrgWorkspace {
	return v.value
}

func (v *NullableActorOrgWorkspace) Set(val *ActorOrgWorkspace) {
	v.value = val
	v.isSet = true
}

func (v NullableActorOrgWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullableActorOrgWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActorOrgWorkspace(val *ActorOrgWorkspace) *NullableActorOrgWorkspace {
	return &NullableActorOrgWorkspace{value: val, isSet: true}
}

func (v NullableActorOrgWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActorOrgWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

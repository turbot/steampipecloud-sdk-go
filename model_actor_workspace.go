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

// ActorWorkspace struct for ActorWorkspace
type ActorWorkspace struct {
	Identity    *Identity  `json:"identity,omitempty"`
	IdentityId  *string    `json:"identity_id,omitempty"`
	Role        *string    `json:"role,omitempty"`
	Workspace   *Workspace `json:"workspace,omitempty"`
	WorkspaceId *string    `json:"workspace_id,omitempty"`
}

// NewActorWorkspace instantiates a new ActorWorkspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActorWorkspace() *ActorWorkspace {
	this := ActorWorkspace{}
	return &this
}

// NewActorWorkspaceWithDefaults instantiates a new ActorWorkspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActorWorkspaceWithDefaults() *ActorWorkspace {
	this := ActorWorkspace{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *ActorWorkspace) GetIdentity() Identity {
	if o == nil || o.Identity == nil {
		var ret Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorWorkspace) GetIdentityOk() (*Identity, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *ActorWorkspace) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given Identity and assigns it to the Identity field.
func (o *ActorWorkspace) SetIdentity(v Identity) {
	o.Identity = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *ActorWorkspace) GetIdentityId() string {
	if o == nil || o.IdentityId == nil {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorWorkspace) GetIdentityIdOk() (*string, bool) {
	if o == nil || o.IdentityId == nil {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *ActorWorkspace) HasIdentityId() bool {
	if o != nil && o.IdentityId != nil {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *ActorWorkspace) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ActorWorkspace) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorWorkspace) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ActorWorkspace) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ActorWorkspace) SetRole(v string) {
	o.Role = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *ActorWorkspace) GetWorkspace() Workspace {
	if o == nil || o.Workspace == nil {
		var ret Workspace
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorWorkspace) GetWorkspaceOk() (*Workspace, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *ActorWorkspace) HasWorkspace() bool {
	if o != nil && o.Workspace != nil {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given Workspace and assigns it to the Workspace field.
func (o *ActorWorkspace) SetWorkspace(v Workspace) {
	o.Workspace = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *ActorWorkspace) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorWorkspace) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *ActorWorkspace) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *ActorWorkspace) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o ActorWorkspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.IdentityId != nil {
		toSerialize["identity_id"] = o.IdentityId
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableActorWorkspace struct {
	value *ActorWorkspace
	isSet bool
}

func (v NullableActorWorkspace) Get() *ActorWorkspace {
	return v.value
}

func (v *NullableActorWorkspace) Set(val *ActorWorkspace) {
	v.value = val
	v.isSet = true
}

func (v NullableActorWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullableActorWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActorWorkspace(val *ActorWorkspace) *NullableActorWorkspace {
	return &NullableActorWorkspace{value: val, isSet: true}
}

func (v NullableActorWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActorWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
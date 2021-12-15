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

// Workspace struct for Workspace
type Workspace struct {
	// The workspace created time.
	CreatedAt string `json:"created_at"`
	// The name of the database.
	DatabaseName *string `json:"database_name,omitempty"`
	// The handle name for the workspace.
	Handle string `json:"handle"`
	// The database hive for this workspace.
	Hive *string `json:"hive,omitempty"`
	Host *string `json:"host,omitempty"`
	// The unique identifier for the workspace.
	Id       string    `json:"id"`
	Identity *Identity `json:"identity,omitempty"`
	// The unique identifier for an identity where the workspace is created.
	IdentityId string  `json:"identity_id"`
	PublicKey  *string `json:"public_key,omitempty"`
	// The workspace updated time.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// The current version ID for the workspace.
	VersionId int32 `json:"version_id"`
	// The current state of the workspace.
	WorkspaceState *string `json:"workspace_state,omitempty"`
}

// NewWorkspace instantiates a new Workspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspace(createdAt string, handle string, id string, identityId string, versionId int32) *Workspace {
	this := Workspace{}
	this.CreatedAt = createdAt
	this.Handle = handle
	this.Id = id
	this.IdentityId = identityId
	this.VersionId = versionId
	return &this
}

// NewWorkspaceWithDefaults instantiates a new Workspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceWithDefaults() *Workspace {
	this := Workspace{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Workspace) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Workspace) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *Workspace) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *Workspace) HasDatabaseName() bool {
	if o != nil && o.DatabaseName != nil {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *Workspace) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetHandle returns the Handle field value
func (o *Workspace) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *Workspace) SetHandle(v string) {
	o.Handle = v
}

// GetHive returns the Hive field value if set, zero value otherwise.
func (o *Workspace) GetHive() string {
	if o == nil || o.Hive == nil {
		var ret string
		return ret
	}
	return *o.Hive
}

// GetHiveOk returns a tuple with the Hive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetHiveOk() (*string, bool) {
	if o == nil || o.Hive == nil {
		return nil, false
	}
	return o.Hive, true
}

// HasHive returns a boolean if a field has been set.
func (o *Workspace) HasHive() bool {
	if o != nil && o.Hive != nil {
		return true
	}

	return false
}

// SetHive gets a reference to the given string and assigns it to the Hive field.
func (o *Workspace) SetHive(v string) {
	o.Hive = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Workspace) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Workspace) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Workspace) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value
func (o *Workspace) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Workspace) SetId(v string) {
	o.Id = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *Workspace) GetIdentity() Identity {
	if o == nil || o.Identity == nil {
		var ret Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetIdentityOk() (*Identity, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *Workspace) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given Identity and assigns it to the Identity field.
func (o *Workspace) SetIdentity(v Identity) {
	o.Identity = &v
}

// GetIdentityId returns the IdentityId field value
func (o *Workspace) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *Workspace) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *Workspace) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *Workspace) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *Workspace) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Workspace) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Workspace) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Workspace) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetVersionId returns the VersionId field value
func (o *Workspace) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *Workspace) SetVersionId(v int32) {
	o.VersionId = v
}

// GetWorkspaceState returns the WorkspaceState field value if set, zero value otherwise.
func (o *Workspace) GetWorkspaceState() string {
	if o == nil || o.WorkspaceState == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceState
}

// GetWorkspaceStateOk returns a tuple with the WorkspaceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetWorkspaceStateOk() (*string, bool) {
	if o == nil || o.WorkspaceState == nil {
		return nil, false
	}
	return o.WorkspaceState, true
}

// HasWorkspaceState returns a boolean if a field has been set.
func (o *Workspace) HasWorkspaceState() bool {
	if o != nil && o.WorkspaceState != nil {
		return true
	}

	return false
}

// SetWorkspaceState gets a reference to the given string and assigns it to the WorkspaceState field.
func (o *Workspace) SetWorkspaceState(v string) {
	o.WorkspaceState = &v
}

func (o Workspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DatabaseName != nil {
		toSerialize["database_name"] = o.DatabaseName
	}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if o.Hive != nil {
		toSerialize["hive"] = o.Hive
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if true {
		toSerialize["identity_id"] = o.IdentityId
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["version_id"] = o.VersionId
	}
	if o.WorkspaceState != nil {
		toSerialize["workspace_state"] = o.WorkspaceState
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspace struct {
	value *Workspace
	isSet bool
}

func (v NullableWorkspace) Get() *Workspace {
	return v.value
}

func (v *NullableWorkspace) Set(val *Workspace) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspace(val *Workspace) *NullableWorkspace {
	return &NullableWorkspace{value: val, isSet: true}
}

func (v NullableWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

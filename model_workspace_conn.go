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

// WorkspaceConn struct for WorkspaceConn
type WorkspaceConn struct {
	Connection *Connection `json:"connection,omitempty"`
	// The unique identifier for the connection.
	ConnectionId string `json:"connection_id"`
	// The time of creation in ISO 8601 UTC.
	CreatedAt string `json:"created_at"`
	CreatedBy *User  `json:"created_by,omitempty"`
	// The ID of the user that created this.
	CreatedById string `json:"created_by_id"`
	// The unique identifier for the workspace connection association.
	Id string `json:"id"`
	// The identity ID where the association exists.
	IdentityId string `json:"identity_id"`
	// The time of the last update in ISO 8601 UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *User   `json:"updated_by,omitempty"`
	// The ID of the user that performed the last update.
	UpdatedById string `json:"updated_by_id"`
	// The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item.
	VersionId int32 `json:"version_id"`
	// The unique identifier for the wokspace.
	WorkspaceId string `json:"workspace_id"`
}

// NewWorkspaceConn instantiates a new WorkspaceConn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceConn(connectionId string, createdAt string, createdById string, id string, identityId string, updatedById string, versionId int32, workspaceId string) *WorkspaceConn {
	this := WorkspaceConn{}
	this.ConnectionId = connectionId
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.Id = id
	this.IdentityId = identityId
	this.UpdatedById = updatedById
	this.VersionId = versionId
	this.WorkspaceId = workspaceId
	return &this
}

// NewWorkspaceConnWithDefaults instantiates a new WorkspaceConn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceConnWithDefaults() *WorkspaceConn {
	this := WorkspaceConn{}
	return &this
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *WorkspaceConn) GetConnection() Connection {
	if o == nil || o.Connection == nil {
		var ret Connection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetConnectionOk() (*Connection, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *WorkspaceConn) HasConnection() bool {
	if o != nil && o.Connection != nil {
		return true
	}

	return false
}

// SetConnection gets a reference to the given Connection and assigns it to the Connection field.
func (o *WorkspaceConn) SetConnection(v Connection) {
	o.Connection = &v
}

// GetConnectionId returns the ConnectionId field value
func (o *WorkspaceConn) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *WorkspaceConn) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkspaceConn) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkspaceConn) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WorkspaceConn) GetCreatedBy() User {
	if o == nil || o.CreatedBy == nil {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetCreatedByOk() (*User, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WorkspaceConn) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *WorkspaceConn) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value
func (o *WorkspaceConn) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *WorkspaceConn) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetId returns the Id field value
func (o *WorkspaceConn) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkspaceConn) SetId(v string) {
	o.Id = v
}

// GetIdentityId returns the IdentityId field value
func (o *WorkspaceConn) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *WorkspaceConn) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkspaceConn) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkspaceConn) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *WorkspaceConn) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *WorkspaceConn) GetUpdatedBy() User {
	if o == nil || o.UpdatedBy == nil {
		var ret User
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetUpdatedByOk() (*User, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *WorkspaceConn) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given User and assigns it to the UpdatedBy field.
func (o *WorkspaceConn) SetUpdatedBy(v User) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value
func (o *WorkspaceConn) GetUpdatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetUpdatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedById, true
}

// SetUpdatedById sets field value
func (o *WorkspaceConn) SetUpdatedById(v string) {
	o.UpdatedById = v
}

// GetVersionId returns the VersionId field value
func (o *WorkspaceConn) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *WorkspaceConn) SetVersionId(v int32) {
	o.VersionId = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *WorkspaceConn) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConn) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *WorkspaceConn) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

func (o WorkspaceConn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	if true {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if true {
		toSerialize["created_by_id"] = o.CreatedById
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["identity_id"] = o.IdentityId
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if true {
		toSerialize["updated_by_id"] = o.UpdatedById
	}
	if true {
		toSerialize["version_id"] = o.VersionId
	}
	if true {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceConn struct {
	value *WorkspaceConn
	isSet bool
}

func (v NullableWorkspaceConn) Get() *WorkspaceConn {
	return v.value
}

func (v *NullableWorkspaceConn) Set(val *WorkspaceConn) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceConn) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceConn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceConn(val *WorkspaceConn) *NullableWorkspaceConn {
	return &NullableWorkspaceConn{value: val, isSet: true}
}

func (v NullableWorkspaceConn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceConn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

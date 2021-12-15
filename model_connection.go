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

// Connection struct for Connection
type Connection struct {
	Config *map[string]interface{} `json:"config,omitempty"`
	// The connection created time.
	CreatedAt string `json:"created_at"`
	// The handle name of the  connection.
	Handle string `json:"handle"`
	// The unique identifier for the connection.
	Id       string    `json:"id"`
	Identity *Identity `json:"identity,omitempty"`
	// The unique identifier for an identity where the connection has been created.
	IdentityId string `json:"identity_id"`
	// The plugin name for the connection.
	Plugin *string `json:"plugin,omitempty"`
	// Type of connection i.e aggregator or connection.
	Type *string `json:"type,omitempty"`
	// The connection updated time.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// The current version ID for the connection.
	VersionId int32 `json:"version_id"`
}

// NewConnection instantiates a new Connection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnection(createdAt string, handle string, id string, identityId string, versionId int32) *Connection {
	this := Connection{}
	this.CreatedAt = createdAt
	this.Handle = handle
	this.Id = id
	this.IdentityId = identityId
	this.VersionId = versionId
	return &this
}

// NewConnectionWithDefaults instantiates a new Connection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWithDefaults() *Connection {
	this := Connection{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Connection) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Connection) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *Connection) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Connection) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Connection) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Connection) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetHandle returns the Handle field value
func (o *Connection) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *Connection) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *Connection) SetHandle(v string) {
	o.Handle = v
}

// GetId returns the Id field value
func (o *Connection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Connection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Connection) SetId(v string) {
	o.Id = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *Connection) GetIdentity() Identity {
	if o == nil || o.Identity == nil {
		var ret Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetIdentityOk() (*Identity, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *Connection) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given Identity and assigns it to the Identity field.
func (o *Connection) SetIdentity(v Identity) {
	o.Identity = &v
}

// GetIdentityId returns the IdentityId field value
func (o *Connection) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *Connection) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *Connection) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *Connection) GetPlugin() string {
	if o == nil || o.Plugin == nil {
		var ret string
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetPluginOk() (*string, bool) {
	if o == nil || o.Plugin == nil {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *Connection) HasPlugin() bool {
	if o != nil && o.Plugin != nil {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given string and assigns it to the Plugin field.
func (o *Connection) SetPlugin(v string) {
	o.Plugin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Connection) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Connection) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Connection) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Connection) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Connection) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Connection) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetVersionId returns the VersionId field value
func (o *Connection) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *Connection) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *Connection) SetVersionId(v int32) {
	o.VersionId = v
}

func (o Connection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["handle"] = o.Handle
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
	if o.Plugin != nil {
		toSerialize["plugin"] = o.Plugin
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["version_id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableConnection struct {
	value *Connection
	isSet bool
}

func (v NullableConnection) Get() *Connection {
	return v.value
}

func (v *NullableConnection) Set(val *Connection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnection(val *Connection) *NullableConnection {
	return &NullableConnection{value: val, isSet: true}
}

func (v NullableConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

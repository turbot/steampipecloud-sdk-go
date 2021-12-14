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

// WorkspaceSchema struct for WorkspaceSchema
type WorkspaceSchema struct {
	Duration int32        `json:"duration"`
	Schemas  []SchemaInfo `json:"schemas"`
}

// NewWorkspaceSchema instantiates a new WorkspaceSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceSchema(duration int32, schemas []SchemaInfo) *WorkspaceSchema {
	this := WorkspaceSchema{}
	this.Duration = duration
	this.Schemas = schemas
	return &this
}

// NewWorkspaceSchemaWithDefaults instantiates a new WorkspaceSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceSchemaWithDefaults() *WorkspaceSchema {
	this := WorkspaceSchema{}
	return &this
}

// GetDuration returns the Duration field value
func (o *WorkspaceSchema) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSchema) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *WorkspaceSchema) SetDuration(v int32) {
	o.Duration = v
}

// GetSchemas returns the Schemas field value
func (o *WorkspaceSchema) GetSchemas() []SchemaInfo {
	if o == nil {
		var ret []SchemaInfo
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSchema) GetSchemasOk() (*[]SchemaInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schemas, true
}

// SetSchemas sets field value
func (o *WorkspaceSchema) SetSchemas(v []SchemaInfo) {
	o.Schemas = v
}

func (o WorkspaceSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceSchema struct {
	value *WorkspaceSchema
	isSet bool
}

func (v NullableWorkspaceSchema) Get() *WorkspaceSchema {
	return v.value
}

func (v *NullableWorkspaceSchema) Set(val *WorkspaceSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceSchema(val *WorkspaceSchema) *NullableWorkspaceSchema {
	return &NullableWorkspaceSchema{value: val, isSet: true}
}

func (v NullableWorkspaceSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

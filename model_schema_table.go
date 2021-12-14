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

// SchemaTable struct for SchemaTable
type SchemaTable struct {
	Columns     []SchemaTableColumn `json:"columns"`
	Description *string             `json:"description,omitempty"`
	Name        string              `json:"name"`
}

// NewSchemaTable instantiates a new SchemaTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaTable(columns []SchemaTableColumn, name string) *SchemaTable {
	this := SchemaTable{}
	this.Columns = columns
	this.Name = name
	return &this
}

// NewSchemaTableWithDefaults instantiates a new SchemaTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaTableWithDefaults() *SchemaTable {
	this := SchemaTable{}
	return &this
}

// GetColumns returns the Columns field value
func (o *SchemaTable) GetColumns() []SchemaTableColumn {
	if o == nil {
		var ret []SchemaTableColumn
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *SchemaTable) GetColumnsOk() (*[]SchemaTableColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value
func (o *SchemaTable) SetColumns(v []SchemaTableColumn) {
	o.Columns = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemaTable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaTable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemaTable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemaTable) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *SchemaTable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemaTable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemaTable) SetName(v string) {
	o.Name = v
}

func (o SchemaTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["columns"] = o.Columns
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaTable struct {
	value *SchemaTable
	isSet bool
}

func (v NullableSchemaTable) Get() *SchemaTable {
	return v.value
}

func (v *NullableSchemaTable) Set(val *SchemaTable) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaTable) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaTable(val *SchemaTable) *NullableSchemaTable {
	return &NullableSchemaTable{value: val, isSet: true}
}

func (v NullableSchemaTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

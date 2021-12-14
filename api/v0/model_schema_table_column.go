/*
Steampipe Cloud

Interrogate your CloudOps data with the simplicity and power of SQL, then share your discoveries using Steampipe Cloud.

API version: 1.0
Contact: help@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package steampipecloud

import (
	"encoding/json"
)

// SchemaTableColumn struct for SchemaTableColumn
type SchemaTableColumn struct {
	DataType string `json:"data_type"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
}

// NewSchemaTableColumn instantiates a new SchemaTableColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaTableColumn(dataType string, name string) *SchemaTableColumn {
	this := SchemaTableColumn{}
	this.DataType = dataType
	this.Name = name
	return &this
}

// NewSchemaTableColumnWithDefaults instantiates a new SchemaTableColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaTableColumnWithDefaults() *SchemaTableColumn {
	this := SchemaTableColumn{}
	return &this
}

// GetDataType returns the DataType field value
func (o *SchemaTableColumn) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *SchemaTableColumn) GetDataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *SchemaTableColumn) SetDataType(v string) {
	o.DataType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemaTableColumn) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaTableColumn) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemaTableColumn) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemaTableColumn) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *SchemaTableColumn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemaTableColumn) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemaTableColumn) SetName(v string) {
	o.Name = v
}

func (o SchemaTableColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data_type"] = o.DataType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaTableColumn struct {
	value *SchemaTableColumn
	isSet bool
}

func (v NullableSchemaTableColumn) Get() *SchemaTableColumn {
	return v.value
}

func (v *NullableSchemaTableColumn) Set(val *SchemaTableColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaTableColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaTableColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaTableColumn(val *SchemaTableColumn) *NullableSchemaTableColumn {
	return &NullableSchemaTableColumn{value: val, isSet: true}
}

func (v NullableSchemaTableColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaTableColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



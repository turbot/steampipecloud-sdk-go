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

// CreateConnectionRequest struct for CreateConnectionRequest
type CreateConnectionRequest struct {
	Config *map[string]interface{} `json:"config,omitempty"`
	Handle string                  `json:"handle"`
	// Type   string                 `json:\"type\" binding:\"required,steampipe_connection_type\"`
	Plugin string `json:"plugin"`
}

// NewCreateConnectionRequest instantiates a new CreateConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConnectionRequest(handle string, plugin string) *CreateConnectionRequest {
	this := CreateConnectionRequest{}
	this.Handle = handle
	this.Plugin = plugin
	return &this
}

// NewCreateConnectionRequestWithDefaults instantiates a new CreateConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConnectionRequestWithDefaults() *CreateConnectionRequest {
	this := CreateConnectionRequest{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateConnectionRequest) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequest) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateConnectionRequest) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *CreateConnectionRequest) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetHandle returns the Handle field value
func (o *CreateConnectionRequest) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequest) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *CreateConnectionRequest) SetHandle(v string) {
	o.Handle = v
}

// GetPlugin returns the Plugin field value
func (o *CreateConnectionRequest) GetPlugin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequest) GetPluginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plugin, true
}

// SetPlugin sets field value
func (o *CreateConnectionRequest) SetPlugin(v string) {
	o.Plugin = v
}

func (o CreateConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if true {
		toSerialize["plugin"] = o.Plugin
	}
	return json.Marshal(toSerialize)
}

type NullableCreateConnectionRequest struct {
	value *CreateConnectionRequest
	isSet bool
}

func (v NullableCreateConnectionRequest) Get() *CreateConnectionRequest {
	return v.value
}

func (v *NullableCreateConnectionRequest) Set(val *CreateConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConnectionRequest(val *CreateConnectionRequest) *NullableCreateConnectionRequest {
	return &NullableCreateConnectionRequest{value: val, isSet: true}
}

func (v NullableCreateConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

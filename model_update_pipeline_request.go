/*
Steampipe Cloud

Steampipe Cloud is a hosted version of Steampipe (https://steampipe.io), an open source tool to instantly query your cloud services (e.g. AWS, Azure, GCP and more) with SQL. No DB required.

API version: {{OPEN_API_VERSION}}
Contact: help@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package steampipecloud

import (
	"encoding/json"
)

// UpdatePipelineRequest struct for UpdatePipelineRequest
type UpdatePipelineRequest struct {
	// A map of arguments to be passed to be pipeline.
	Args *map[string]interface{} `json:"args,omitempty"`
	// The frequency at which the pipeline will run.
	Frequency *map[string]interface{} `json:"frequency,omitempty"`
	// The tags for this pipeline.
	Tags *map[string]interface{} `json:"tags,omitempty"`
	// The title of the pipeline.
	Title *string `json:"title,omitempty"`
}

// NewUpdatePipelineRequest instantiates a new UpdatePipelineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePipelineRequest() *UpdatePipelineRequest {
	this := UpdatePipelineRequest{}
	return &this
}

// NewUpdatePipelineRequestWithDefaults instantiates a new UpdatePipelineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePipelineRequestWithDefaults() *UpdatePipelineRequest {
	this := UpdatePipelineRequest{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *UpdatePipelineRequest) GetArgs() map[string]interface{} {
	if o == nil || o.Args == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineRequest) GetArgsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *UpdatePipelineRequest) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given map[string]interface{} and assigns it to the Args field.
func (o *UpdatePipelineRequest) SetArgs(v map[string]interface{}) {
	o.Args = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *UpdatePipelineRequest) GetFrequency() map[string]interface{} {
	if o == nil || o.Frequency == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineRequest) GetFrequencyOk() (*map[string]interface{}, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *UpdatePipelineRequest) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given map[string]interface{} and assigns it to the Frequency field.
func (o *UpdatePipelineRequest) SetFrequency(v map[string]interface{}) {
	o.Frequency = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdatePipelineRequest) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineRequest) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdatePipelineRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *UpdatePipelineRequest) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdatePipelineRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdatePipelineRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdatePipelineRequest) SetTitle(v string) {
	o.Title = &v
}

func (o UpdatePipelineRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePipelineRequest struct {
	value *UpdatePipelineRequest
	isSet bool
}

func (v NullableUpdatePipelineRequest) Get() *UpdatePipelineRequest {
	return v.value
}

func (v *NullableUpdatePipelineRequest) Set(val *UpdatePipelineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePipelineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePipelineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePipelineRequest(val *UpdatePipelineRequest) *NullableUpdatePipelineRequest {
	return &NullableUpdatePipelineRequest{value: val, isSet: true}
}

func (v NullableUpdatePipelineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePipelineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

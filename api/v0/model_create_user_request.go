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

// CreateUserRequest struct for CreateUserRequest
type CreateUserRequest struct {
	DisplayName *string `json:"display_name,omitempty"`
	Email string `json:"email"`
	Handle string `json:"handle"`
	Url *string `json:"url,omitempty"`
}

// NewCreateUserRequest instantiates a new CreateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserRequest(email string, handle string) *CreateUserRequest {
	this := CreateUserRequest{}
	this.Email = email
	this.Handle = handle
	return &this
}

// NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserRequestWithDefaults() *CreateUserRequest {
	this := CreateUserRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateUserRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateUserRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateUserRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value
func (o *CreateUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetHandle returns the Handle field value
func (o *CreateUserRequest) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetHandleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *CreateUserRequest) SetHandle(v string) {
	o.Handle = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateUserRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateUserRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateUserRequest) SetUrl(v string) {
	o.Url = &v
}

func (o CreateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUserRequest struct {
	value *CreateUserRequest
	isSet bool
}

func (v NullableCreateUserRequest) Get() *CreateUserRequest {
	return v.value
}

func (v *NullableCreateUserRequest) Set(val *CreateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserRequest(val *CreateUserRequest) *NullableCreateUserRequest {
	return &NullableCreateUserRequest{value: val, isSet: true}
}

func (v NullableCreateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// UserQuota struct for UserQuota
type UserQuota struct {
	Association  map[string]Quota `json:"association"`
	Conn         Quota            `json:"conn"`
	Mod          map[string]Quota `json:"mod"`
	Organization Quota            `json:"organization"`
	Token        Quota            `json:"token"`
	Workspace    Quota            `json:"workspace"`
}

// NewUserQuota instantiates a new UserQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserQuota(association map[string]Quota, conn Quota, mod map[string]Quota, organization Quota, token Quota, workspace Quota) *UserQuota {
	this := UserQuota{}
	this.Association = association
	this.Conn = conn
	this.Mod = mod
	this.Organization = organization
	this.Token = token
	this.Workspace = workspace
	return &this
}

// NewUserQuotaWithDefaults instantiates a new UserQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserQuotaWithDefaults() *UserQuota {
	this := UserQuota{}
	return &this
}

// GetAssociation returns the Association field value
func (o *UserQuota) GetAssociation() map[string]Quota {
	if o == nil {
		var ret map[string]Quota
		return ret
	}

	return o.Association
}

// GetAssociationOk returns a tuple with the Association field value
// and a boolean to check if the value has been set.
func (o *UserQuota) GetAssociationOk() (*map[string]Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Association, true
}

// SetAssociation sets field value
func (o *UserQuota) SetAssociation(v map[string]Quota) {
	o.Association = v
}

// GetConn returns the Conn field value
func (o *UserQuota) GetConn() Quota {
	if o == nil {
		var ret Quota
		return ret
	}

	return o.Conn
}

// GetConnOk returns a tuple with the Conn field value
// and a boolean to check if the value has been set.
func (o *UserQuota) GetConnOk() (*Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conn, true
}

// SetConn sets field value
func (o *UserQuota) SetConn(v Quota) {
	o.Conn = v
}

// GetMod returns the Mod field value
func (o *UserQuota) GetMod() map[string]Quota {
	if o == nil {
		var ret map[string]Quota
		return ret
	}

	return o.Mod
}

// GetModOk returns a tuple with the Mod field value
// and a boolean to check if the value has been set.
func (o *UserQuota) GetModOk() (*map[string]Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mod, true
}

// SetMod sets field value
func (o *UserQuota) SetMod(v map[string]Quota) {
	o.Mod = v
}

// GetOrganization returns the Organization field value
func (o *UserQuota) GetOrganization() Quota {
	if o == nil {
		var ret Quota
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *UserQuota) GetOrganizationOk() (*Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *UserQuota) SetOrganization(v Quota) {
	o.Organization = v
}

// GetToken returns the Token field value
func (o *UserQuota) GetToken() Quota {
	if o == nil {
		var ret Quota
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserQuota) GetTokenOk() (*Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserQuota) SetToken(v Quota) {
	o.Token = v
}

// GetWorkspace returns the Workspace field value
func (o *UserQuota) GetWorkspace() Quota {
	if o == nil {
		var ret Quota
		return ret
	}

	return o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value
// and a boolean to check if the value has been set.
func (o *UserQuota) GetWorkspaceOk() (*Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workspace, true
}

// SetWorkspace sets field value
func (o *UserQuota) SetWorkspace(v Quota) {
	o.Workspace = v
}

func (o UserQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["association"] = o.Association
	}
	if true {
		toSerialize["conn"] = o.Conn
	}
	if true {
		toSerialize["mod"] = o.Mod
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["workspace"] = o.Workspace
	}
	return json.Marshal(toSerialize)
}

type NullableUserQuota struct {
	value *UserQuota
	isSet bool
}

func (v NullableUserQuota) Get() *UserQuota {
	return v.value
}

func (v *NullableUserQuota) Set(val *UserQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableUserQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableUserQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserQuota(val *UserQuota) *NullableUserQuota {
	return &NullableUserQuota{value: val, isSet: true}
}

func (v NullableUserQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

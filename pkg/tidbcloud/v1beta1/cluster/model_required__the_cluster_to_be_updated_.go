/*
TiDB Cloud Serverless Open API

TiDB Cloud Serverless Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cluster

import (
	"encoding/json"
)

// checks if the RequiredTheClusterToBeUpdated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequiredTheClusterToBeUpdated{}

// RequiredTheClusterToBeUpdated Required. The cluster to be updated.
type RequiredTheClusterToBeUpdated struct {
	// Required. User friendly display name of the cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// Optional. The spending limit for the cluster.
	SpendingLimit *ClusterSpendingLimit `json:"spendingLimit,omitempty"`
	// Optional. Automated backup policy to set on the cluster.
	AutomatedBackupPolicy *V1beta1ClusterAutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`
	// Optional. The endpoints for connecting to the cluster.
	Endpoints *V1beta1ClusterEndpoints `json:"endpoints,omitempty"`
	// Optional. The labels for the cluster. tidb.cloud/organization. The label for the cluster organization id. tidb.cloud/project. The label for the cluster project id.
	Labels               *map[string]string `json:"labels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequiredTheClusterToBeUpdated RequiredTheClusterToBeUpdated

// NewRequiredTheClusterToBeUpdated instantiates a new RequiredTheClusterToBeUpdated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequiredTheClusterToBeUpdated() *RequiredTheClusterToBeUpdated {
	this := RequiredTheClusterToBeUpdated{}
	return &this
}

// NewRequiredTheClusterToBeUpdatedWithDefaults instantiates a new RequiredTheClusterToBeUpdated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequiredTheClusterToBeUpdatedWithDefaults() *RequiredTheClusterToBeUpdated {
	this := RequiredTheClusterToBeUpdated{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *RequiredTheClusterToBeUpdated) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredTheClusterToBeUpdated) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RequiredTheClusterToBeUpdated) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RequiredTheClusterToBeUpdated) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSpendingLimit returns the SpendingLimit field value if set, zero value otherwise.
func (o *RequiredTheClusterToBeUpdated) GetSpendingLimit() ClusterSpendingLimit {
	if o == nil || IsNil(o.SpendingLimit) {
		var ret ClusterSpendingLimit
		return ret
	}
	return *o.SpendingLimit
}

// GetSpendingLimitOk returns a tuple with the SpendingLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredTheClusterToBeUpdated) GetSpendingLimitOk() (*ClusterSpendingLimit, bool) {
	if o == nil || IsNil(o.SpendingLimit) {
		return nil, false
	}
	return o.SpendingLimit, true
}

// HasSpendingLimit returns a boolean if a field has been set.
func (o *RequiredTheClusterToBeUpdated) HasSpendingLimit() bool {
	if o != nil && !IsNil(o.SpendingLimit) {
		return true
	}

	return false
}

// SetSpendingLimit gets a reference to the given ClusterSpendingLimit and assigns it to the SpendingLimit field.
func (o *RequiredTheClusterToBeUpdated) SetSpendingLimit(v ClusterSpendingLimit) {
	o.SpendingLimit = &v
}

// GetAutomatedBackupPolicy returns the AutomatedBackupPolicy field value if set, zero value otherwise.
func (o *RequiredTheClusterToBeUpdated) GetAutomatedBackupPolicy() V1beta1ClusterAutomatedBackupPolicy {
	if o == nil || IsNil(o.AutomatedBackupPolicy) {
		var ret V1beta1ClusterAutomatedBackupPolicy
		return ret
	}
	return *o.AutomatedBackupPolicy
}

// GetAutomatedBackupPolicyOk returns a tuple with the AutomatedBackupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredTheClusterToBeUpdated) GetAutomatedBackupPolicyOk() (*V1beta1ClusterAutomatedBackupPolicy, bool) {
	if o == nil || IsNil(o.AutomatedBackupPolicy) {
		return nil, false
	}
	return o.AutomatedBackupPolicy, true
}

// HasAutomatedBackupPolicy returns a boolean if a field has been set.
func (o *RequiredTheClusterToBeUpdated) HasAutomatedBackupPolicy() bool {
	if o != nil && !IsNil(o.AutomatedBackupPolicy) {
		return true
	}

	return false
}

// SetAutomatedBackupPolicy gets a reference to the given V1beta1ClusterAutomatedBackupPolicy and assigns it to the AutomatedBackupPolicy field.
func (o *RequiredTheClusterToBeUpdated) SetAutomatedBackupPolicy(v V1beta1ClusterAutomatedBackupPolicy) {
	o.AutomatedBackupPolicy = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *RequiredTheClusterToBeUpdated) GetEndpoints() V1beta1ClusterEndpoints {
	if o == nil || IsNil(o.Endpoints) {
		var ret V1beta1ClusterEndpoints
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredTheClusterToBeUpdated) GetEndpointsOk() (*V1beta1ClusterEndpoints, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *RequiredTheClusterToBeUpdated) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given V1beta1ClusterEndpoints and assigns it to the Endpoints field.
func (o *RequiredTheClusterToBeUpdated) SetEndpoints(v V1beta1ClusterEndpoints) {
	o.Endpoints = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *RequiredTheClusterToBeUpdated) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredTheClusterToBeUpdated) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *RequiredTheClusterToBeUpdated) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *RequiredTheClusterToBeUpdated) SetLabels(v map[string]string) {
	o.Labels = &v
}

func (o RequiredTheClusterToBeUpdated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequiredTheClusterToBeUpdated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.SpendingLimit) {
		toSerialize["spendingLimit"] = o.SpendingLimit
	}
	if !IsNil(o.AutomatedBackupPolicy) {
		toSerialize["automatedBackupPolicy"] = o.AutomatedBackupPolicy
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequiredTheClusterToBeUpdated) UnmarshalJSON(data []byte) (err error) {
	varRequiredTheClusterToBeUpdated := _RequiredTheClusterToBeUpdated{}

	err = json.Unmarshal(data, &varRequiredTheClusterToBeUpdated)

	if err != nil {
		return err
	}

	*o = RequiredTheClusterToBeUpdated(varRequiredTheClusterToBeUpdated)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "spendingLimit")
		delete(additionalProperties, "automatedBackupPolicy")
		delete(additionalProperties, "endpoints")
		delete(additionalProperties, "labels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequiredTheClusterToBeUpdated struct {
	value *RequiredTheClusterToBeUpdated
	isSet bool
}

func (v NullableRequiredTheClusterToBeUpdated) Get() *RequiredTheClusterToBeUpdated {
	return v.value
}

func (v *NullableRequiredTheClusterToBeUpdated) Set(val *RequiredTheClusterToBeUpdated) {
	v.value = val
	v.isSet = true
}

func (v NullableRequiredTheClusterToBeUpdated) IsSet() bool {
	return v.isSet
}

func (v *NullableRequiredTheClusterToBeUpdated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequiredTheClusterToBeUpdated(val *RequiredTheClusterToBeUpdated) *NullableRequiredTheClusterToBeUpdated {
	return &NullableRequiredTheClusterToBeUpdated{value: val, isSet: true}
}

func (v NullableRequiredTheClusterToBeUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequiredTheClusterToBeUpdated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
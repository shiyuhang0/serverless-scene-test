/*
TiDB Cloud Serverless Open API

TiDB Cloud Serverless Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package branch

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Branch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Branch{}

// Branch struct for Branch
type Branch struct {
	// Output Only. The name of the resource.
	Name *string `json:"name,omitempty"`
	// Output only. The system-generated ID of the resource.
	BranchId *string `json:"branchId,omitempty"`
	// Required. User-settable and human-readable display name for the branch.
	DisplayName string `json:"displayName"`
	// Output only. The cluster ID of this branch.
	ClusterId *string `json:"clusterId,omitempty"`
	// Optional. The parent ID of this branch.
	ParentId *string `json:"parentId,omitempty"`
	// Output only. The creator of the branch.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Output only. The state of this branch.
	State *BranchState `json:"state,omitempty"`
	// Optional. The endpoints of this branch.
	Endpoints *BranchEndpoints `json:"endpoints,omitempty"`
	// Output only. User name prefix of this branch. For each TiDB Serverless branch, TiDB Cloud generates a unique prefix to distinguish it from other branches. Whenever you use or set a database user name, you must include the prefix in the user name.
	UserPrefix NullableString `json:"userPrefix,omitempty"`
	// Output only. Usage metrics of this branch. Only display in FULL view.
	Usage      *BranchUsage `json:"usage,omitempty"`
	CreateTime *time.Time   `json:"createTime,omitempty"`
	UpdateTime *time.Time   `json:"updateTime,omitempty"`
	// Optional. The annotations of this branch..
	Annotations *map[string]string `json:"annotations,omitempty"`
	// Output only. The parent display name of this branch.
	ParentDisplayName *string `json:"parentDisplayName,omitempty"`
	// Optional. The point in time on the parent branch the branch will be created from.
	ParentTimestamp      NullableTime `json:"parentTimestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Branch Branch

// NewBranch instantiates a new Branch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranch(displayName string) *Branch {
	this := Branch{}
	this.DisplayName = displayName
	return &this
}

// NewBranchWithDefaults instantiates a new Branch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchWithDefaults() *Branch {
	this := Branch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Branch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Branch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Branch) SetName(v string) {
	o.Name = &v
}

// GetBranchId returns the BranchId field value if set, zero value otherwise.
func (o *Branch) GetBranchId() string {
	if o == nil || IsNil(o.BranchId) {
		var ret string
		return ret
	}
	return *o.BranchId
}

// GetBranchIdOk returns a tuple with the BranchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetBranchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BranchId) {
		return nil, false
	}
	return o.BranchId, true
}

// HasBranchId returns a boolean if a field has been set.
func (o *Branch) HasBranchId() bool {
	if o != nil && !IsNil(o.BranchId) {
		return true
	}

	return false
}

// SetBranchId gets a reference to the given string and assigns it to the BranchId field.
func (o *Branch) SetBranchId(v string) {
	o.BranchId = &v
}

// GetDisplayName returns the DisplayName field value
func (o *Branch) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Branch) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Branch) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *Branch) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *Branch) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *Branch) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Branch) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Branch) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *Branch) SetParentId(v string) {
	o.ParentId = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Branch) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Branch) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Branch) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Branch) GetState() BranchState {
	if o == nil || IsNil(o.State) {
		var ret BranchState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetStateOk() (*BranchState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Branch) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given BranchState and assigns it to the State field.
func (o *Branch) SetState(v BranchState) {
	o.State = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *Branch) GetEndpoints() BranchEndpoints {
	if o == nil || IsNil(o.Endpoints) {
		var ret BranchEndpoints
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetEndpointsOk() (*BranchEndpoints, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *Branch) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given BranchEndpoints and assigns it to the Endpoints field.
func (o *Branch) SetEndpoints(v BranchEndpoints) {
	o.Endpoints = &v
}

// GetUserPrefix returns the UserPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Branch) GetUserPrefix() string {
	if o == nil || IsNil(o.UserPrefix.Get()) {
		var ret string
		return ret
	}
	return *o.UserPrefix.Get()
}

// GetUserPrefixOk returns a tuple with the UserPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Branch) GetUserPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserPrefix.Get(), o.UserPrefix.IsSet()
}

// HasUserPrefix returns a boolean if a field has been set.
func (o *Branch) HasUserPrefix() bool {
	if o != nil && o.UserPrefix.IsSet() {
		return true
	}

	return false
}

// SetUserPrefix gets a reference to the given NullableString and assigns it to the UserPrefix field.
func (o *Branch) SetUserPrefix(v string) {
	o.UserPrefix.Set(&v)
}

// SetUserPrefixNil sets the value for UserPrefix to be an explicit nil
func (o *Branch) SetUserPrefixNil() {
	o.UserPrefix.Set(nil)
}

// UnsetUserPrefix ensures that no value is present for UserPrefix, not even an explicit nil
func (o *Branch) UnsetUserPrefix() {
	o.UserPrefix.Unset()
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *Branch) GetUsage() BranchUsage {
	if o == nil || IsNil(o.Usage) {
		var ret BranchUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetUsageOk() (*BranchUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *Branch) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given BranchUsage and assigns it to the Usage field.
func (o *Branch) SetUsage(v BranchUsage) {
	o.Usage = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Branch) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Branch) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Branch) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Branch) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Branch) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *Branch) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *Branch) GetAnnotations() map[string]string {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *Branch) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *Branch) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetParentDisplayName returns the ParentDisplayName field value if set, zero value otherwise.
func (o *Branch) GetParentDisplayName() string {
	if o == nil || IsNil(o.ParentDisplayName) {
		var ret string
		return ret
	}
	return *o.ParentDisplayName
}

// GetParentDisplayNameOk returns a tuple with the ParentDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetParentDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentDisplayName) {
		return nil, false
	}
	return o.ParentDisplayName, true
}

// HasParentDisplayName returns a boolean if a field has been set.
func (o *Branch) HasParentDisplayName() bool {
	if o != nil && !IsNil(o.ParentDisplayName) {
		return true
	}

	return false
}

// SetParentDisplayName gets a reference to the given string and assigns it to the ParentDisplayName field.
func (o *Branch) SetParentDisplayName(v string) {
	o.ParentDisplayName = &v
}

// GetParentTimestamp returns the ParentTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Branch) GetParentTimestamp() time.Time {
	if o == nil || IsNil(o.ParentTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ParentTimestamp.Get()
}

// GetParentTimestampOk returns a tuple with the ParentTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Branch) GetParentTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentTimestamp.Get(), o.ParentTimestamp.IsSet()
}

// HasParentTimestamp returns a boolean if a field has been set.
func (o *Branch) HasParentTimestamp() bool {
	if o != nil && o.ParentTimestamp.IsSet() {
		return true
	}

	return false
}

// SetParentTimestamp gets a reference to the given NullableTime and assigns it to the ParentTimestamp field.
func (o *Branch) SetParentTimestamp(v time.Time) {
	o.ParentTimestamp.Set(&v)
}

// SetParentTimestampNil sets the value for ParentTimestamp to be an explicit nil
func (o *Branch) SetParentTimestampNil() {
	o.ParentTimestamp.Set(nil)
}

// UnsetParentTimestamp ensures that no value is present for ParentTimestamp, not even an explicit nil
func (o *Branch) UnsetParentTimestamp() {
	o.ParentTimestamp.Unset()
}

func (o Branch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Branch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.BranchId) {
		toSerialize["branchId"] = o.BranchId
	}
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.ClusterId) {
		toSerialize["clusterId"] = o.ClusterId
	}
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if o.UserPrefix.IsSet() {
		toSerialize["userPrefix"] = o.UserPrefix.Get()
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.ParentDisplayName) {
		toSerialize["parentDisplayName"] = o.ParentDisplayName
	}
	if o.ParentTimestamp.IsSet() {
		toSerialize["parentTimestamp"] = o.ParentTimestamp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Branch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBranch := _Branch{}

	err = json.Unmarshal(data, &varBranch)

	if err != nil {
		return err
	}

	*o = Branch(varBranch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "branchId")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "clusterId")
		delete(additionalProperties, "parentId")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "state")
		delete(additionalProperties, "endpoints")
		delete(additionalProperties, "userPrefix")
		delete(additionalProperties, "usage")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "annotations")
		delete(additionalProperties, "parentDisplayName")
		delete(additionalProperties, "parentTimestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBranch struct {
	value *Branch
	isSet bool
}

func (v NullableBranch) Get() *Branch {
	return v.value
}

func (v *NullableBranch) Set(val *Branch) {
	v.value = val
	v.isSet = true
}

func (v NullableBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranch(val *Branch) *NullableBranch {
	return &NullableBranch{value: val, isSet: true}
}

func (v NullableBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
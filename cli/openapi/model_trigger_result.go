/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TriggerResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerResult{}

// TriggerResult struct for TriggerResult
type TriggerResult struct {
	TriggerType   *string                     `json:"triggerType,omitempty"`
	TriggerResult *TriggerResultTriggerResult `json:"triggerResult,omitempty"`
}

// NewTriggerResult instantiates a new TriggerResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerResult() *TriggerResult {
	this := TriggerResult{}
	return &this
}

// NewTriggerResultWithDefaults instantiates a new TriggerResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerResultWithDefaults() *TriggerResult {
	this := TriggerResult{}
	return &this
}

// GetTriggerType returns the TriggerType field value if set, zero value otherwise.
func (o *TriggerResult) GetTriggerType() string {
	if o == nil || isNil(o.TriggerType) {
		var ret string
		return ret
	}
	return *o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerResult) GetTriggerTypeOk() (*string, bool) {
	if o == nil || isNil(o.TriggerType) {
		return nil, false
	}
	return o.TriggerType, true
}

// HasTriggerType returns a boolean if a field has been set.
func (o *TriggerResult) HasTriggerType() bool {
	if o != nil && !isNil(o.TriggerType) {
		return true
	}

	return false
}

// SetTriggerType gets a reference to the given string and assigns it to the TriggerType field.
func (o *TriggerResult) SetTriggerType(v string) {
	o.TriggerType = &v
}

// GetTriggerResult returns the TriggerResult field value if set, zero value otherwise.
func (o *TriggerResult) GetTriggerResult() TriggerResultTriggerResult {
	if o == nil || isNil(o.TriggerResult) {
		var ret TriggerResultTriggerResult
		return ret
	}
	return *o.TriggerResult
}

// GetTriggerResultOk returns a tuple with the TriggerResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerResult) GetTriggerResultOk() (*TriggerResultTriggerResult, bool) {
	if o == nil || isNil(o.TriggerResult) {
		return nil, false
	}
	return o.TriggerResult, true
}

// HasTriggerResult returns a boolean if a field has been set.
func (o *TriggerResult) HasTriggerResult() bool {
	if o != nil && !isNil(o.TriggerResult) {
		return true
	}

	return false
}

// SetTriggerResult gets a reference to the given TriggerResultTriggerResult and assigns it to the TriggerResult field.
func (o *TriggerResult) SetTriggerResult(v TriggerResultTriggerResult) {
	o.TriggerResult = &v
}

func (o TriggerResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TriggerType) {
		toSerialize["triggerType"] = o.TriggerType
	}
	if !isNil(o.TriggerResult) {
		toSerialize["triggerResult"] = o.TriggerResult
	}
	return toSerialize, nil
}

type NullableTriggerResult struct {
	value *TriggerResult
	isSet bool
}

func (v NullableTriggerResult) Get() *TriggerResult {
	return v.value
}

func (v *NullableTriggerResult) Set(val *TriggerResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerResult(val *TriggerResult) *NullableTriggerResult {
	return &NullableTriggerResult{value: val, isSet: true}
}

func (v NullableTriggerResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

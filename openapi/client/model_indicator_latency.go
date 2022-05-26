/*
Pyrra

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IndicatorLatency struct for IndicatorLatency
type IndicatorLatency struct {
	Success  Query    `json:"success"`
	Total    Query    `json:"total"`
	Grouping []string `json:"grouping,omitempty"`
}

// NewIndicatorLatency instantiates a new IndicatorLatency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorLatency(success Query, total Query) *IndicatorLatency {
	this := IndicatorLatency{}
	this.Success = success
	this.Total = total
	return &this
}

// NewIndicatorLatencyWithDefaults instantiates a new IndicatorLatency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorLatencyWithDefaults() *IndicatorLatency {
	this := IndicatorLatency{}
	return &this
}

// GetSuccess returns the Success field value
func (o *IndicatorLatency) GetSuccess() Query {
	if o == nil {
		var ret Query
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *IndicatorLatency) GetSuccessOk() (*Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *IndicatorLatency) SetSuccess(v Query) {
	o.Success = v
}

// GetTotal returns the Total field value
func (o *IndicatorLatency) GetTotal() Query {
	if o == nil {
		var ret Query
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *IndicatorLatency) GetTotalOk() (*Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *IndicatorLatency) SetTotal(v Query) {
	o.Total = v
}

// GetGrouping returns the Grouping field value if set, zero value otherwise.
func (o *IndicatorLatency) GetGrouping() []string {
	if o == nil || o.Grouping == nil {
		var ret []string
		return ret
	}
	return o.Grouping
}

// GetGroupingOk returns a tuple with the Grouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorLatency) GetGroupingOk() ([]string, bool) {
	if o == nil || o.Grouping == nil {
		return nil, false
	}
	return o.Grouping, true
}

// HasGrouping returns a boolean if a field has been set.
func (o *IndicatorLatency) HasGrouping() bool {
	if o != nil && o.Grouping != nil {
		return true
	}

	return false
}

// SetGrouping gets a reference to the given []string and assigns it to the Grouping field.
func (o *IndicatorLatency) SetGrouping(v []string) {
	o.Grouping = v
}

func (o IndicatorLatency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["success"] = o.Success
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if o.Grouping != nil {
		toSerialize["grouping"] = o.Grouping
	}
	return json.Marshal(toSerialize)
}

type NullableIndicatorLatency struct {
	value *IndicatorLatency
	isSet bool
}

func (v NullableIndicatorLatency) Get() *IndicatorLatency {
	return v.value
}

func (v *NullableIndicatorLatency) Set(val *IndicatorLatency) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorLatency) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorLatency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorLatency(val *IndicatorLatency) *NullableIndicatorLatency {
	return &NullableIndicatorLatency{value: val, isSet: true}
}

func (v NullableIndicatorLatency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorLatency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// IndicatorRatio struct for IndicatorRatio
type IndicatorRatio struct {
	Errors   Query    `json:"errors"`
	Total    Query    `json:"total"`
	Grouping []string `json:"grouping,omitempty"`
}

// NewIndicatorRatio instantiates a new IndicatorRatio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorRatio(errors Query, total Query) *IndicatorRatio {
	this := IndicatorRatio{}
	this.Errors = errors
	this.Total = total
	return &this
}

// NewIndicatorRatioWithDefaults instantiates a new IndicatorRatio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorRatioWithDefaults() *IndicatorRatio {
	this := IndicatorRatio{}
	return &this
}

// GetErrors returns the Errors field value
func (o *IndicatorRatio) GetErrors() Query {
	if o == nil {
		var ret Query
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *IndicatorRatio) GetErrorsOk() (*Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *IndicatorRatio) SetErrors(v Query) {
	o.Errors = v
}

// GetTotal returns the Total field value
func (o *IndicatorRatio) GetTotal() Query {
	if o == nil {
		var ret Query
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *IndicatorRatio) GetTotalOk() (*Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *IndicatorRatio) SetTotal(v Query) {
	o.Total = v
}

// GetGrouping returns the Grouping field value if set, zero value otherwise.
func (o *IndicatorRatio) GetGrouping() []string {
	if o == nil || o.Grouping == nil {
		var ret []string
		return ret
	}
	return o.Grouping
}

// GetGroupingOk returns a tuple with the Grouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorRatio) GetGroupingOk() ([]string, bool) {
	if o == nil || o.Grouping == nil {
		return nil, false
	}
	return o.Grouping, true
}

// HasGrouping returns a boolean if a field has been set.
func (o *IndicatorRatio) HasGrouping() bool {
	if o != nil && o.Grouping != nil {
		return true
	}

	return false
}

// SetGrouping gets a reference to the given []string and assigns it to the Grouping field.
func (o *IndicatorRatio) SetGrouping(v []string) {
	o.Grouping = v
}

func (o IndicatorRatio) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if o.Grouping != nil {
		toSerialize["grouping"] = o.Grouping
	}
	return json.Marshal(toSerialize)
}

type NullableIndicatorRatio struct {
	value *IndicatorRatio
	isSet bool
}

func (v NullableIndicatorRatio) Get() *IndicatorRatio {
	return v.value
}

func (v *NullableIndicatorRatio) Set(val *IndicatorRatio) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorRatio) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorRatio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorRatio(val *IndicatorRatio) *NullableIndicatorRatio {
	return &NullableIndicatorRatio{value: val, isSet: true}
}

func (v NullableIndicatorRatio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorRatio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

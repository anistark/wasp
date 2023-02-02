/*
Wasp API

REST API for the Wasp node

API version: 123
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the BurnLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BurnLog{}

// BurnLog struct for BurnLog
type BurnLog struct {
	Records []BurnRecord `json:"records"`
}

// NewBurnLog instantiates a new BurnLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnLog(records []BurnRecord) *BurnLog {
	this := BurnLog{}
	this.Records = records
	return &this
}

// NewBurnLogWithDefaults instantiates a new BurnLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnLogWithDefaults() *BurnLog {
	this := BurnLog{}
	return &this
}

// GetRecords returns the Records field value
func (o *BurnLog) GetRecords() []BurnRecord {
	if o == nil {
		var ret []BurnRecord
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *BurnLog) GetRecordsOk() ([]BurnRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.Records, true
}

// SetRecords sets field value
func (o *BurnLog) SetRecords(v []BurnRecord) {
	o.Records = v
}

func (o BurnLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurnLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["records"] = o.Records
	return toSerialize, nil
}

type NullableBurnLog struct {
	value *BurnLog
	isSet bool
}

func (v NullableBurnLog) Get() *BurnLog {
	return v.value
}

func (v *NullableBurnLog) Set(val *BurnLog) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnLog) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnLog(val *BurnLog) *NullableBurnLog {
	return &NullableBurnLog{value: val, isSet: true}
}

func (v NullableBurnLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



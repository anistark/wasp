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

// checks if the FungibleTokens type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FungibleTokens{}

// FungibleTokens struct for FungibleTokens
type FungibleTokens struct {
	BaseTokens int64 `json:"baseTokens"`
	NativeTokens []NativeToken `json:"nativeTokens"`
}

// NewFungibleTokens instantiates a new FungibleTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFungibleTokens(baseTokens int64, nativeTokens []NativeToken) *FungibleTokens {
	this := FungibleTokens{}
	this.BaseTokens = baseTokens
	this.NativeTokens = nativeTokens
	return &this
}

// NewFungibleTokensWithDefaults instantiates a new FungibleTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFungibleTokensWithDefaults() *FungibleTokens {
	this := FungibleTokens{}
	return &this
}

// GetBaseTokens returns the BaseTokens field value
func (o *FungibleTokens) GetBaseTokens() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BaseTokens
}

// GetBaseTokensOk returns a tuple with the BaseTokens field value
// and a boolean to check if the value has been set.
func (o *FungibleTokens) GetBaseTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseTokens, true
}

// SetBaseTokens sets field value
func (o *FungibleTokens) SetBaseTokens(v int64) {
	o.BaseTokens = v
}

// GetNativeTokens returns the NativeTokens field value
func (o *FungibleTokens) GetNativeTokens() []NativeToken {
	if o == nil {
		var ret []NativeToken
		return ret
	}

	return o.NativeTokens
}

// GetNativeTokensOk returns a tuple with the NativeTokens field value
// and a boolean to check if the value has been set.
func (o *FungibleTokens) GetNativeTokensOk() ([]NativeToken, bool) {
	if o == nil {
		return nil, false
	}
	return o.NativeTokens, true
}

// SetNativeTokens sets field value
func (o *FungibleTokens) SetNativeTokens(v []NativeToken) {
	o.NativeTokens = v
}

func (o FungibleTokens) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FungibleTokens) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["baseTokens"] = o.BaseTokens
	toSerialize["nativeTokens"] = o.NativeTokens
	return toSerialize, nil
}

type NullableFungibleTokens struct {
	value *FungibleTokens
	isSet bool
}

func (v NullableFungibleTokens) Get() *FungibleTokens {
	return v.value
}

func (v *NullableFungibleTokens) Set(val *FungibleTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableFungibleTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableFungibleTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFungibleTokens(val *FungibleTokens) *NullableFungibleTokens {
	return &NullableFungibleTokens{value: val, isSet: true}
}

func (v NullableFungibleTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFungibleTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableAddressMapOfAddressArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddressMapOfAddressArrayAppendParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s ImmutableAddressMapOfAddressArrayAppendParams) ValueAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamValueAddr))
}

type MutableAddressMapOfAddressArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddressMapOfAddressArrayAppendParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s MutableAddressMapOfAddressArrayAppendParams) ValueAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamValueAddr))
}

type ImmutableAddressMapOfAddressArrayClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddressMapOfAddressArrayClearParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableAddressMapOfAddressArrayClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddressMapOfAddressArrayClearParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableAddressMapOfAddressArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddressMapOfAddressArraySetParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableAddressMapOfAddressArraySetParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s ImmutableAddressMapOfAddressArraySetParams) ValueAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamValueAddr))
}

type MutableAddressMapOfAddressArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddressMapOfAddressArraySetParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableAddressMapOfAddressArraySetParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s MutableAddressMapOfAddressArraySetParams) ValueAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamValueAddr))
}

type ImmutableAddressMapOfAddressMapClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddressMapOfAddressMapClearParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableAddressMapOfAddressMapClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddressMapOfAddressMapClearParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableAddressMapOfAddressMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddressMapOfAddressMapSetParams) KeyAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s ImmutableAddressMapOfAddressMapSetParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s ImmutableAddressMapOfAddressMapSetParams) ValueAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamValueAddr))
}

type MutableAddressMapOfAddressMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddressMapOfAddressMapSetParams) KeyAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s MutableAddressMapOfAddressMapSetParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s MutableAddressMapOfAddressMapSetParams) ValueAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamValueAddr))
}

type ImmutableArrayOfAddressArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfAddressArrayAppendParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfAddressArrayAppendParams) ValueAddr() ArrayOfImmutableAddress {
	return ArrayOfImmutableAddress{proxy: s.proxy.Root(ParamValueAddr)}
}

type MutableArrayOfAddressArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfAddressArrayAppendParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfAddressArrayAppendParams) ValueAddr() ArrayOfMutableAddress {
	return ArrayOfMutableAddress{proxy: s.proxy.Root(ParamValueAddr)}
}

type ImmutableArrayOfAddressArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfAddressArraySetParams) Index0() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex0))
}

func (s ImmutableArrayOfAddressArraySetParams) Index1() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex1))
}

func (s ImmutableArrayOfAddressArraySetParams) ValueAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamValueAddr))
}

type MutableArrayOfAddressArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfAddressArraySetParams) Index0() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex0))
}

func (s MutableArrayOfAddressArraySetParams) Index1() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex1))
}

func (s MutableArrayOfAddressArraySetParams) ValueAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamValueAddr))
}

type ImmutableArrayOfAddressMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfAddressMapSetParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfAddressMapSetParams) KeyAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s ImmutableArrayOfAddressMapSetParams) ValueAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamValueAddr))
}

type MutableArrayOfAddressMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfAddressMapSetParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfAddressMapSetParams) KeyAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s MutableArrayOfAddressMapSetParams) ValueAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamValueAddr))
}

type ImmutableArrayOfStringArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfStringArrayAppendParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfStringArrayAppendParams) Value() ArrayOfImmutableString {
	return ArrayOfImmutableString{proxy: s.proxy.Root(ParamValue)}
}

type MutableArrayOfStringArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfStringArrayAppendParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfStringArrayAppendParams) Value() ArrayOfMutableString {
	return ArrayOfMutableString{proxy: s.proxy.Root(ParamValue)}
}

type ImmutableArrayOfStringArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfStringArraySetParams) Index0() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex0))
}

func (s ImmutableArrayOfStringArraySetParams) Index1() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex1))
}

func (s ImmutableArrayOfStringArraySetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableArrayOfStringArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfStringArraySetParams) Index0() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex0))
}

func (s MutableArrayOfStringArraySetParams) Index1() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex1))
}

func (s MutableArrayOfStringArraySetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableArrayOfStringMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfStringMapSetParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfStringMapSetParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

func (s ImmutableArrayOfStringMapSetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableArrayOfStringMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfStringMapSetParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfStringMapSetParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

func (s MutableArrayOfStringMapSetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type MapStringToImmutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToImmutableBytes) GetBytes(key string) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type ImmutableParamTypesParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableParamTypesParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

func (s ImmutableParamTypesParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s ImmutableParamTypesParams) Bool() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.proxy.Root(ParamBool))
}

func (s ImmutableParamTypesParams) Bytes() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ParamBytes))
}

func (s ImmutableParamTypesParams) ChainID() wasmtypes.ScImmutableChainID {
	return wasmtypes.NewScImmutableChainID(s.proxy.Root(ParamChainID))
}

func (s ImmutableParamTypesParams) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamHash))
}

func (s ImmutableParamTypesParams) Hname() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHname))
}

func (s ImmutableParamTypesParams) Int16() wasmtypes.ScImmutableInt16 {
	return wasmtypes.NewScImmutableInt16(s.proxy.Root(ParamInt16))
}

func (s ImmutableParamTypesParams) Int32() wasmtypes.ScImmutableInt32 {
	return wasmtypes.NewScImmutableInt32(s.proxy.Root(ParamInt32))
}

func (s ImmutableParamTypesParams) Int64() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamInt64))
}

func (s ImmutableParamTypesParams) Int8() wasmtypes.ScImmutableInt8 {
	return wasmtypes.NewScImmutableInt8(s.proxy.Root(ParamInt8))
}

func (s ImmutableParamTypesParams) NftID() wasmtypes.ScImmutableNftID {
	return wasmtypes.NewScImmutableNftID(s.proxy.Root(ParamNftID))
}

// special hook to be able to pass key/values as raw bytes
func (s ImmutableParamTypesParams) Param() MapStringToImmutableBytes {
	//nolint:gosimple
	return MapStringToImmutableBytes{proxy: s.proxy}
}

func (s ImmutableParamTypesParams) RequestID() wasmtypes.ScImmutableRequestID {
	return wasmtypes.NewScImmutableRequestID(s.proxy.Root(ParamRequestID))
}

func (s ImmutableParamTypesParams) String() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamString))
}

func (s ImmutableParamTypesParams) TokenID() wasmtypes.ScImmutableTokenID {
	return wasmtypes.NewScImmutableTokenID(s.proxy.Root(ParamTokenID))
}

func (s ImmutableParamTypesParams) Uint16() wasmtypes.ScImmutableUint16 {
	return wasmtypes.NewScImmutableUint16(s.proxy.Root(ParamUint16))
}

func (s ImmutableParamTypesParams) Uint32() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamUint32))
}

func (s ImmutableParamTypesParams) Uint64() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamUint64))
}

func (s ImmutableParamTypesParams) Uint8() wasmtypes.ScImmutableUint8 {
	return wasmtypes.NewScImmutableUint8(s.proxy.Root(ParamUint8))
}

type MapStringToMutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToMutableBytes) Clear() {
	m.proxy.ClearMap()
}

func (m MapStringToMutableBytes) GetBytes(key string) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type MutableParamTypesParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableParamTypesParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}

func (s MutableParamTypesParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s MutableParamTypesParams) Bool() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.proxy.Root(ParamBool))
}

func (s MutableParamTypesParams) Bytes() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ParamBytes))
}

func (s MutableParamTypesParams) ChainID() wasmtypes.ScMutableChainID {
	return wasmtypes.NewScMutableChainID(s.proxy.Root(ParamChainID))
}

func (s MutableParamTypesParams) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamHash))
}

func (s MutableParamTypesParams) Hname() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHname))
}

func (s MutableParamTypesParams) Int16() wasmtypes.ScMutableInt16 {
	return wasmtypes.NewScMutableInt16(s.proxy.Root(ParamInt16))
}

func (s MutableParamTypesParams) Int32() wasmtypes.ScMutableInt32 {
	return wasmtypes.NewScMutableInt32(s.proxy.Root(ParamInt32))
}

func (s MutableParamTypesParams) Int64() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamInt64))
}

func (s MutableParamTypesParams) Int8() wasmtypes.ScMutableInt8 {
	return wasmtypes.NewScMutableInt8(s.proxy.Root(ParamInt8))
}

func (s MutableParamTypesParams) NftID() wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(s.proxy.Root(ParamNftID))
}

// special hook to be able to pass key/values as raw bytes
func (s MutableParamTypesParams) Param() MapStringToMutableBytes {
	//nolint:gosimple
	return MapStringToMutableBytes{proxy: s.proxy}
}

func (s MutableParamTypesParams) RequestID() wasmtypes.ScMutableRequestID {
	return wasmtypes.NewScMutableRequestID(s.proxy.Root(ParamRequestID))
}

func (s MutableParamTypesParams) String() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamString))
}

func (s MutableParamTypesParams) TokenID() wasmtypes.ScMutableTokenID {
	return wasmtypes.NewScMutableTokenID(s.proxy.Root(ParamTokenID))
}

func (s MutableParamTypesParams) Uint16() wasmtypes.ScMutableUint16 {
	return wasmtypes.NewScMutableUint16(s.proxy.Root(ParamUint16))
}

func (s MutableParamTypesParams) Uint32() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamUint32))
}

func (s MutableParamTypesParams) Uint64() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamUint64))
}

func (s MutableParamTypesParams) Uint8() wasmtypes.ScMutableUint8 {
	return wasmtypes.NewScMutableUint8(s.proxy.Root(ParamUint8))
}

type ImmutableStringMapOfStringArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringArrayAppendParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableStringMapOfStringArrayAppendParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableStringMapOfStringArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringArrayAppendParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableStringMapOfStringArrayAppendParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableStringMapOfStringArrayClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringArrayClearParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableStringMapOfStringArrayClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringArrayClearParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableStringMapOfStringArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringArraySetParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableStringMapOfStringArraySetParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableStringMapOfStringArraySetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableStringMapOfStringArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringArraySetParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableStringMapOfStringArraySetParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableStringMapOfStringArraySetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableStringMapOfStringMapClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringMapClearParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableStringMapOfStringMapClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringMapClearParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableStringMapOfStringMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringMapSetParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

func (s ImmutableStringMapOfStringMapSetParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableStringMapOfStringMapSetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableStringMapOfStringMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringMapSetParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

func (s MutableStringMapOfStringMapSetParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableStringMapOfStringMapSetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableTriggerEventParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTriggerEventParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

func (s ImmutableTriggerEventParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableTriggerEventParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTriggerEventParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}

func (s MutableTriggerEventParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableAddressMapOfAddressArrayLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddressMapOfAddressArrayLengthParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableAddressMapOfAddressArrayLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddressMapOfAddressArrayLengthParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableAddressMapOfAddressArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddressMapOfAddressArrayValueParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableAddressMapOfAddressArrayValueParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableAddressMapOfAddressArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddressMapOfAddressArrayValueParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableAddressMapOfAddressArrayValueParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableAddressMapOfAddressMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddressMapOfAddressMapValueParams) KeyAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s ImmutableAddressMapOfAddressMapValueParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableAddressMapOfAddressMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddressMapOfAddressMapValueParams) KeyAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s MutableAddressMapOfAddressMapValueParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableArrayOfAddressArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfAddressArrayValueParams) Index0() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex0))
}

func (s ImmutableArrayOfAddressArrayValueParams) Index1() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex1))
}

type MutableArrayOfAddressArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfAddressArrayValueParams) Index0() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex0))
}

func (s MutableArrayOfAddressArrayValueParams) Index1() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex1))
}

type ImmutableArrayOfAddressMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfAddressMapValueParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfAddressMapValueParams) KeyAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamKeyAddr))
}

type MutableArrayOfAddressMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfAddressMapValueParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfAddressMapValueParams) KeyAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamKeyAddr))
}

type ImmutableArrayOfStringArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfStringArrayValueParams) Index0() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex0))
}

func (s ImmutableArrayOfStringArrayValueParams) Index1() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex1))
}

type MutableArrayOfStringArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfStringArrayValueParams) Index0() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex0))
}

func (s MutableArrayOfStringArrayValueParams) Index1() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex1))
}

type ImmutableArrayOfStringMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfStringMapValueParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfStringMapValueParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

type MutableArrayOfStringMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfStringMapValueParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfStringMapValueParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

type ImmutableBigIntAddParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBigIntAddParams) Lhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamLhs))
}

func (s ImmutableBigIntAddParams) Rhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamRhs))
}

type MutableBigIntAddParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBigIntAddParams) Lhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamLhs))
}

func (s MutableBigIntAddParams) Rhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamRhs))
}

type ImmutableBigIntDivParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBigIntDivParams) Lhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamLhs))
}

func (s ImmutableBigIntDivParams) Rhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamRhs))
}

type MutableBigIntDivParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBigIntDivParams) Lhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamLhs))
}

func (s MutableBigIntDivParams) Rhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamRhs))
}

type ImmutableBigIntModParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBigIntModParams) Lhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamLhs))
}

func (s ImmutableBigIntModParams) Rhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamRhs))
}

type MutableBigIntModParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBigIntModParams) Lhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamLhs))
}

func (s MutableBigIntModParams) Rhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamRhs))
}

type ImmutableBigIntMulParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBigIntMulParams) Lhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamLhs))
}

func (s ImmutableBigIntMulParams) Rhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamRhs))
}

type MutableBigIntMulParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBigIntMulParams) Lhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamLhs))
}

func (s MutableBigIntMulParams) Rhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamRhs))
}

type ImmutableBigIntShlParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBigIntShlParams) Lhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamLhs))
}

func (s ImmutableBigIntShlParams) Shift() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamShift))
}

type MutableBigIntShlParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBigIntShlParams) Lhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamLhs))
}

func (s MutableBigIntShlParams) Shift() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamShift))
}

type ImmutableBigIntShrParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBigIntShrParams) Lhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamLhs))
}

func (s ImmutableBigIntShrParams) Shift() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamShift))
}

type MutableBigIntShrParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBigIntShrParams) Lhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamLhs))
}

func (s MutableBigIntShrParams) Shift() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamShift))
}

type ImmutableBigIntSubParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBigIntSubParams) Lhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamLhs))
}

func (s ImmutableBigIntSubParams) Rhs() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.proxy.Root(ParamRhs))
}

type MutableBigIntSubParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBigIntSubParams) Lhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamLhs))
}

func (s MutableBigIntSubParams) Rhs() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.proxy.Root(ParamRhs))
}

type ImmutableBlockRecordParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBlockRecordParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

func (s ImmutableBlockRecordParams) RecordIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamRecordIndex))
}

type MutableBlockRecordParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBlockRecordParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

func (s MutableBlockRecordParams) RecordIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamRecordIndex))
}

type ImmutableBlockRecordsParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBlockRecordsParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

type MutableBlockRecordsParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBlockRecordsParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

type ImmutableStringMapOfStringArrayLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringArrayLengthParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableStringMapOfStringArrayLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringArrayLengthParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableStringMapOfStringArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringArrayValueParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableStringMapOfStringArrayValueParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableStringMapOfStringArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringArrayValueParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableStringMapOfStringArrayValueParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableStringMapOfStringMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringMapValueParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

func (s ImmutableStringMapOfStringMapValueParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableStringMapOfStringMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringMapValueParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

func (s MutableStringMapOfStringMapValueParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

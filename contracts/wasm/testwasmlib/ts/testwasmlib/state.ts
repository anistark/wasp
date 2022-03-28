// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ArrayOfImmutableAddressArray extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getAddressArray(index: u32): sc.ImmutableAddressArray {
		return new sc.ImmutableAddressArray(this.proxy.index(index));
	}
}

export class MapAddressToImmutableAddressArray extends wasmtypes.ScProxy {

	getAddressArray(key: wasmtypes.ScAddress): sc.ImmutableAddressArray {
		return new sc.ImmutableAddressArray(this.proxy.key(wasmtypes.addressToBytes(key)));
	}
}

export class ArrayOfImmutableAddressMap extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getAddressMap(index: u32): sc.ImmutableAddressMap {
		return new sc.ImmutableAddressMap(this.proxy.index(index));
	}
}

export class MapAddressToImmutableAddressMap extends wasmtypes.ScProxy {

	getAddressMap(key: wasmtypes.ScAddress): sc.ImmutableAddressMap {
		return new sc.ImmutableAddressMap(this.proxy.key(wasmtypes.addressToBytes(key)));
	}
}

export class MapInt32ToImmutableLongitude extends wasmtypes.ScProxy {

	getLongitude(key: i32): sc.ImmutableLongitude {
		return new sc.ImmutableLongitude(this.proxy.key(wasmtypes.int32ToBytes(key)));
	}
}

export class ArrayOfImmutableStringArray extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getStringArray(index: u32): sc.ImmutableStringArray {
		return new sc.ImmutableStringArray(this.proxy.index(index));
	}
}

export class ArrayOfImmutableStringMap extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getStringMap(index: u32): sc.ImmutableStringMap {
		return new sc.ImmutableStringMap(this.proxy.index(index));
	}
}

export class MapStringToImmutableStringArray extends wasmtypes.ScProxy {

	getStringArray(key: string): sc.ImmutableStringArray {
		return new sc.ImmutableStringArray(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class MapStringToImmutableStringMap extends wasmtypes.ScProxy {

	getStringMap(key: string): sc.ImmutableStringMap {
		return new sc.ImmutableStringMap(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class ImmutableTestWasmLibState extends wasmtypes.ScProxy {
	addressArrayArrays(): sc.ArrayOfImmutableAddressArray {
		return new sc.ArrayOfImmutableAddressArray(this.proxy.root(sc.StateAddressArrayArrays));
	}

	addressArrays(): sc.MapAddressToImmutableAddressArray {
		return new sc.MapAddressToImmutableAddressArray(this.proxy.root(sc.StateAddressArrays));
	}

	addressMapArrays(): sc.ArrayOfImmutableAddressMap {
		return new sc.ArrayOfImmutableAddressMap(this.proxy.root(sc.StateAddressMapArrays));
	}

	addressMaps(): sc.MapAddressToImmutableAddressMap {
		return new sc.MapAddressToImmutableAddressMap(this.proxy.root(sc.StateAddressMaps));
	}

	latLong(): sc.MapInt32ToImmutableLongitude {
		return new sc.MapInt32ToImmutableLongitude(this.proxy.root(sc.StateLatLong));
	}

	random(): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.StateRandom));
	}

	stringArrayOfArrays(): sc.ArrayOfImmutableStringArray {
		return new sc.ArrayOfImmutableStringArray(this.proxy.root(sc.StateStringArrayOfArrays));
	}

	stringArrayOfMaps(): sc.ArrayOfImmutableStringMap {
		return new sc.ArrayOfImmutableStringMap(this.proxy.root(sc.StateStringArrayOfMaps));
	}

	stringMapOfArrays(): sc.MapStringToImmutableStringArray {
		return new sc.MapStringToImmutableStringArray(this.proxy.root(sc.StateStringMapOfArrays));
	}

	stringMapOfMaps(): sc.MapStringToImmutableStringMap {
		return new sc.MapStringToImmutableStringMap(this.proxy.root(sc.StateStringMapOfMaps));
	}
}

export class ArrayOfMutableAddressArray extends wasmtypes.ScProxy {

	appendAddressArray(): sc.MutableAddressArray {
		return new sc.MutableAddressArray(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getAddressArray(index: u32): sc.MutableAddressArray {
		return new sc.MutableAddressArray(this.proxy.index(index));
	}
}

export class MapAddressToMutableAddressArray extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getAddressArray(key: wasmtypes.ScAddress): sc.MutableAddressArray {
		return new sc.MutableAddressArray(this.proxy.key(wasmtypes.addressToBytes(key)));
	}
}

export class ArrayOfMutableAddressMap extends wasmtypes.ScProxy {

	appendAddressMap(): sc.MutableAddressMap {
		return new sc.MutableAddressMap(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getAddressMap(index: u32): sc.MutableAddressMap {
		return new sc.MutableAddressMap(this.proxy.index(index));
	}
}

export class MapAddressToMutableAddressMap extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getAddressMap(key: wasmtypes.ScAddress): sc.MutableAddressMap {
		return new sc.MutableAddressMap(this.proxy.key(wasmtypes.addressToBytes(key)));
	}
}

export class MapInt32ToMutableLongitude extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getLongitude(key: i32): sc.MutableLongitude {
		return new sc.MutableLongitude(this.proxy.key(wasmtypes.int32ToBytes(key)));
	}
}

export class ArrayOfMutableStringArray extends wasmtypes.ScProxy {

	appendStringArray(): sc.MutableStringArray {
		return new sc.MutableStringArray(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getStringArray(index: u32): sc.MutableStringArray {
		return new sc.MutableStringArray(this.proxy.index(index));
	}
}

export class ArrayOfMutableStringMap extends wasmtypes.ScProxy {

	appendStringMap(): sc.MutableStringMap {
		return new sc.MutableStringMap(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getStringMap(index: u32): sc.MutableStringMap {
		return new sc.MutableStringMap(this.proxy.index(index));
	}
}

export class MapStringToMutableStringArray extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getStringArray(key: string): sc.MutableStringArray {
		return new sc.MutableStringArray(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class MapStringToMutableStringMap extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getStringMap(key: string): sc.MutableStringMap {
		return new sc.MutableStringMap(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class MutableTestWasmLibState extends wasmtypes.ScProxy {
	asImmutable(): sc.ImmutableTestWasmLibState {
		return new sc.ImmutableTestWasmLibState(this.proxy);
	}

	addressArrayArrays(): sc.ArrayOfMutableAddressArray {
		return new sc.ArrayOfMutableAddressArray(this.proxy.root(sc.StateAddressArrayArrays));
	}

	addressArrays(): sc.MapAddressToMutableAddressArray {
		return new sc.MapAddressToMutableAddressArray(this.proxy.root(sc.StateAddressArrays));
	}

	addressMapArrays(): sc.ArrayOfMutableAddressMap {
		return new sc.ArrayOfMutableAddressMap(this.proxy.root(sc.StateAddressMapArrays));
	}

	addressMaps(): sc.MapAddressToMutableAddressMap {
		return new sc.MapAddressToMutableAddressMap(this.proxy.root(sc.StateAddressMaps));
	}

	latLong(): sc.MapInt32ToMutableLongitude {
		return new sc.MapInt32ToMutableLongitude(this.proxy.root(sc.StateLatLong));
	}

	random(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.root(sc.StateRandom));
	}

	stringArrayOfArrays(): sc.ArrayOfMutableStringArray {
		return new sc.ArrayOfMutableStringArray(this.proxy.root(sc.StateStringArrayOfArrays));
	}

	stringArrayOfMaps(): sc.ArrayOfMutableStringMap {
		return new sc.ArrayOfMutableStringMap(this.proxy.root(sc.StateStringArrayOfMaps));
	}

	stringMapOfArrays(): sc.MapStringToMutableStringArray {
		return new sc.MapStringToMutableStringArray(this.proxy.root(sc.StateStringMapOfArrays));
	}

	stringMapOfMaps(): sc.MapStringToMutableStringMap {
		return new sc.MapStringToMutableStringMap(this.proxy.root(sc.StateStringMapOfMaps));
	}
}

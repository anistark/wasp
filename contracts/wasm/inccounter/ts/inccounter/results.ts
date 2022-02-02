// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ImmutableGetCounterResults extends wasmlib.ScMapID {
    counter(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultCounter));
	}
}

export class MutableGetCounterResults extends wasmlib.ScMapID {
    counter(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, wasmlib.Key32.fromString(sc.ResultCounter));
	}
}

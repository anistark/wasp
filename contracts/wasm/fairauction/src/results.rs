// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;
use crate::structs::*;
use crate::typedefs::*;

#[derive(Clone, Copy)]
pub struct ImmutableGetInfoResults {
    pub(crate) id: i32,
}

impl ImmutableGetInfoResults {
    pub fn bidders(&self) -> ScImmutableInt32 {
		ScImmutableInt32::new(self.id, RESULT_BIDDERS.get_key_id())
	}

    pub fn color(&self) -> ScImmutableColor {
		ScImmutableColor::new(self.id, RESULT_COLOR.get_key_id())
	}

    pub fn creator(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.id, RESULT_CREATOR.get_key_id())
	}

    pub fn deposit(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.id, RESULT_DEPOSIT.get_key_id())
	}

    pub fn description(&self) -> ScImmutableString {
		ScImmutableString::new(self.id, RESULT_DESCRIPTION.get_key_id())
	}

    pub fn duration(&self) -> ScImmutableInt32 {
		ScImmutableInt32::new(self.id, RESULT_DURATION.get_key_id())
	}

    pub fn highest_bid(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.id, RESULT_HIGHEST_BID.get_key_id())
	}

    pub fn highest_bidder(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.id, RESULT_HIGHEST_BIDDER.get_key_id())
	}

    pub fn minimum_bid(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.id, RESULT_MINIMUM_BID.get_key_id())
	}

    pub fn num_tokens(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.id, RESULT_NUM_TOKENS.get_key_id())
	}

    pub fn owner_margin(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.id, RESULT_OWNER_MARGIN.get_key_id())
	}

    pub fn when_started(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.id, RESULT_WHEN_STARTED.get_key_id())
	}
}

#[derive(Clone, Copy)]
pub struct MutableGetInfoResults {
    pub(crate) id: i32,
}

impl MutableGetInfoResults {
    pub fn bidders(&self) -> ScMutableInt32 {
		ScMutableInt32::new(self.id, RESULT_BIDDERS.get_key_id())
	}

    pub fn color(&self) -> ScMutableColor {
		ScMutableColor::new(self.id, RESULT_COLOR.get_key_id())
	}

    pub fn creator(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.id, RESULT_CREATOR.get_key_id())
	}

    pub fn deposit(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.id, RESULT_DEPOSIT.get_key_id())
	}

    pub fn description(&self) -> ScMutableString {
		ScMutableString::new(self.id, RESULT_DESCRIPTION.get_key_id())
	}

    pub fn duration(&self) -> ScMutableInt32 {
		ScMutableInt32::new(self.id, RESULT_DURATION.get_key_id())
	}

    pub fn highest_bid(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.id, RESULT_HIGHEST_BID.get_key_id())
	}

    pub fn highest_bidder(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.id, RESULT_HIGHEST_BIDDER.get_key_id())
	}

    pub fn minimum_bid(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.id, RESULT_MINIMUM_BID.get_key_id())
	}

    pub fn num_tokens(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.id, RESULT_NUM_TOKENS.get_key_id())
	}

    pub fn owner_margin(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.id, RESULT_OWNER_MARGIN.get_key_id())
	}

    pub fn when_started(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.id, RESULT_WHEN_STARTED.get_key_id())
	}
}

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ImmutableApproveParams struct {
	id int32
}

func (s ImmutableApproveParams) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, wasmlib.KeyID(ParamAmount))
}

func (s ImmutableApproveParams) Delegation() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamDelegation))
}

type MutableApproveParams struct {
	id int32
}

func (s MutableApproveParams) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, wasmlib.KeyID(ParamAmount))
}

func (s MutableApproveParams) Delegation() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamDelegation))
}

type ImmutableInitParams struct {
	id int32
}

func (s ImmutableInitParams) Creator() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamCreator])
}

func (s ImmutableInitParams) Supply() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamSupply])
}

type MutableInitParams struct {
	id int32
}

func (s MutableInitParams) Creator() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamCreator])
}

func (s MutableInitParams) Supply() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamSupply])
}

type ImmutableTransferParams struct {
	id int32
}

func (s ImmutableTransferParams) Account() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamAccount))
}

func (s ImmutableTransferParams) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, wasmlib.KeyID(ParamAmount))
}

type MutableTransferParams struct {
	id int32
}

func (s MutableTransferParams) Account() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamAccount))
}

func (s MutableTransferParams) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, wasmlib.KeyID(ParamAmount))
}

type ImmutableTransferFromParams struct {
	id int32
}

func (s ImmutableTransferFromParams) Account() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamAccount))
}

func (s ImmutableTransferFromParams) Amount() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, wasmlib.KeyID(ParamAmount))
}

func (s ImmutableTransferFromParams) Recipient() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamRecipient))
}

type MutableTransferFromParams struct {
	id int32
}

func (s MutableTransferFromParams) Account() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamAccount))
}

func (s MutableTransferFromParams) Amount() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, wasmlib.KeyID(ParamAmount))
}

func (s MutableTransferFromParams) Recipient() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamRecipient))
}

type ImmutableAllowanceParams struct {
	id int32
}

func (s ImmutableAllowanceParams) Account() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamAccount))
}

func (s ImmutableAllowanceParams) Delegation() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamDelegation))
}

type MutableAllowanceParams struct {
	id int32
}

func (s MutableAllowanceParams) Account() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamAccount))
}

func (s MutableAllowanceParams) Delegation() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamDelegation))
}

type ImmutableBalanceOfParams struct {
	id int32
}

func (s ImmutableBalanceOfParams) Account() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, wasmlib.KeyID(ParamAccount))
}

type MutableBalanceOfParams struct {
	id int32
}

func (s MutableBalanceOfParams) Account() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, wasmlib.KeyID(ParamAccount))
}

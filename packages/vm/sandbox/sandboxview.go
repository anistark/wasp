// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package sandbox

import (
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/vm/execution"
)

type sandboxView struct {
	SandboxBase
}

func NewSandboxView(ctx execution.WaspContext) isc.SandboxView {
	ret := &sandboxView{}
	ret.Ctx = ctx
	return ret
}

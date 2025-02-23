package testutil

import (
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/transaction"
	"github.com/iotaledger/wasp/packages/vm/core/migrations"
	"github.com/iotaledger/wasp/packages/vm/gas"
)

func DummyStateMetadata(commitment *state.L1Commitment) *transaction.StateMetadata {
	return transaction.NewStateMetadata(
		commitment,
		gas.DefaultFeePolicy(),
		migrations.BaseSchemaVersion+uint32(len(migrations.Migrations)),
		"",
	)
}

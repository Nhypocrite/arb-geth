package eth

import (
	"context"

	"github.com/Nhypocrite/arb-geth/core"
	"github.com/Nhypocrite/arb-geth/core/state"
	"github.com/Nhypocrite/arb-geth/core/types"
	"github.com/Nhypocrite/arb-geth/core/vm"
	"github.com/Nhypocrite/arb-geth/eth/tracers"
	"github.com/Nhypocrite/arb-geth/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}

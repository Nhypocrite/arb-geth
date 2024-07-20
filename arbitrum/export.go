package arbitrum

import (
	"context"

	"github.com/Nhypocrite/arb-geth/common/hexutil"
	"github.com/Nhypocrite/arb-geth/core"
	"github.com/Nhypocrite/arb-geth/internal/ethapi"
	"github.com/Nhypocrite/arb-geth/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *ethapi.StateOverride, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, overrides, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}

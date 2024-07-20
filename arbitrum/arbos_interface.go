package arbitrum

import (
	"context"

	"github.com/Nhypocrite/arb-geth/arbitrum_types"
	"github.com/Nhypocrite/arb-geth/core"
	"github.com/Nhypocrite/arb-geth/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}

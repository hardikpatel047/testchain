package keeper

import (
	"github.com/hardikpatel047/testchain/x/testchain/types"
)

var _ types.QueryServer = Keeper{}

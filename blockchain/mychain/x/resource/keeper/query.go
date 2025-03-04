package keeper

import (
	"mychain/x/resource/types"
)

var _ types.QueryServer = Keeper{}

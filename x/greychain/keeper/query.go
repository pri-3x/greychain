package keeper

import (
	"greychain/x/greychain/types"
)

var _ types.QueryServer = Keeper{}

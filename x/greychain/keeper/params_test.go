package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "greychain/testutil/keeper"
	"greychain/x/greychain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.GreychainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

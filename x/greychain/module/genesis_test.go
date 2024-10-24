package greychain_test

import (
	"testing"

	keepertest "greychain/testutil/keeper"
	"greychain/testutil/nullify"
	greychain "greychain/x/greychain/module"
	"greychain/x/greychain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GreychainKeeper(t)
	greychain.InitGenesis(ctx, k, genesisState)
	got := greychain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

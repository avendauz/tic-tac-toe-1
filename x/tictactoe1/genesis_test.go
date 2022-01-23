package tictactoe1_test

import (
	"testing"

	keepertest "github.com/avendauz/tic-tac-toe-1/testutil/keeper"
	"github.com/avendauz/tic-tac-toe-1/testutil/nullify"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Tictactoe1Keeper(t)
	tictactoe1.InitGenesis(ctx, *k, genesisState)
	got := tictactoe1.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

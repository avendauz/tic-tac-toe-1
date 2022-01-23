package keeper_test

import (
	"testing"

	testkeeper "github.com/avendauz/tic-tac-toe-1/testutil/keeper"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Tictactoe1Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

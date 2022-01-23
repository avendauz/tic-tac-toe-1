package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/avendauz/tic-tac-toe-1/testutil/keeper"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/keeper"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Tictactoe1Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

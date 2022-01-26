package keeper

import (
	"context"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) OpenGame (goCtx context.Context, msg *types.MsgOpenGame) (*types.MsgOpenGameResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	k.AddOGame(ctx, msg.Creator, msg.Uuid);

	return &types.MsgOpenGameResponse{

	}, nil
}
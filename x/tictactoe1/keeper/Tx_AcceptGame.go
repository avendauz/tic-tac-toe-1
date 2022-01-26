package keeper

import (
	"context"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

func (k Keeper) AcceptGame (goCtx context.Context, req *types.MsgAcceptGame) (*types.MsgAcceptGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasOpenGame(ctx, req.Creator, req.Uuid) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	k.RemoveOpenGame(ctx, req.Creator, req.Uuid)

	k.AddCurrGame(ctx, req.Creator, req.Uuid, req.Challenger)

	return &types.MsgAcceptGameResponse{}, nil

}
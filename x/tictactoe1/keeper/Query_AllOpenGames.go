package keeper

import (
	"context"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AllOpenGames (goCtx context.Context, req *types.QueryAllOpenGamesRequest) (*types.QueryAllOpenGamesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	resp := types.QueryAllOpenGamesResponse{
		Games: k.GetAllOGames(ctx),
	}

	return &resp, nil
}
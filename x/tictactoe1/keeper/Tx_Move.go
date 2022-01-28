package keeper

import (
	"context"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

func (k Keeper) Move (goCtx context.Context, req *types.MsgMove) (*types.MsgMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasCurrGame(ctx, req.Creator, req.Uuid) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	game, _ := k.GetCurrGame(ctx, req.Creator, req.Uuid)		// TODO handle error

	if req.Player != game.Turn {
		return nil, sdkerrors.New("tictactoe1", 2, "Not your turn")
	}

	var turn string

	if req.Player == game.X {

		if game.Board[req.Cell] != 0 {
			return nil, sdkerrors.New("tictactoe1", 2, "Invalid move")
		}
		game.Board[req.Cell] = 2
		turn = game.O
	}

	if req.Player == game.O {
		if game.Board[req.Cell] != 0 {
			return nil, sdkerrors.New("tictactoe1", 2, "Invalid move")
		}
		game.Board[req.Cell] = 1
		turn = game.X
	}

	updatedGame := types.CurrGame{
		X:     game.X,
		O:     game.O,
		Uuid:  game.Uuid,
		Board: game.Board,
		Turn: turn ,
	}

	k.SetCurrGame(ctx, req.Creator, req.Uuid, updatedGame)

	if k.CheckEndgame(updatedGame.Board) {
		k.RemoveCurrGame(ctx, req.Creator, req.Uuid)
	}

	return &types.MsgMoveResponse{}, nil

}
package keeper

import (
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/types"
)

var _ types.QueryServer = Keeper{}

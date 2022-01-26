package keeper

import (
	"crypto/sha1"
	"github.com/avendauz/tic-tac-toe-1/x/tictactoe1/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func HashOpponents (initiator string, challenger string) []byte {
	hasher := sha1.New();
	hasher.Write([]byte(initiator + challenger))
	return hasher.Sum(nil)
}

func CreateOpenGameKey (initiator string, uuid string) []byte {
	return []byte(initiator + "\x00" + uuid)
}

func (k Keeper) AddOGame (
	ctx sdk.Context,
	initiator string,
	uuid string,
	) {

	//var newGame = types.OpenGame{
	//	Initiator: initiator,
	//	Uuid: uuid,
	//}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OpenGameKey))
	//v := k.cdc.MustMarshal(&newGame)
	store.Set(CreateOpenGameKey(initiator, uuid), make([]byte, 0))
}

func (k Keeper) HasOpenGame (
	ctx sdk.Context,
	initiator string,
	uuid string,
	) bool {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OpenGameKey))

	return store.Has(CreateOpenGameKey(initiator, uuid))
}

func (k Keeper) RemoveOpenGame (
	ctx sdk.Context,
	initiator string,
	uuid string,
	) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OpenGameKey))
	store.Delete(CreateOpenGameKey(initiator, uuid))
}

func (k Keeper) GetAllOGames (ctx sdk.Context) []string {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OpenGameKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	var keys []string

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		keys = append(keys, key)
	}

	return keys
}

func (k Keeper) GetAllInitiatorsOGames (ctx sdk.Context, initiator string) []string {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OpenGamePrefix(types.OpenGameKey, initiator + "\x00"))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	var keys []string

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		keys = append(keys, key)
	}

	return keys
}

func (k Keeper) AddCurrGame (ctx sdk.Context, initiator string, uuid string, challenger string) {
	hashed := HashOpponents(initiator, challenger)

	board := make([]byte, 9)		// initialize board, 9 cells

	game := types.CurrGame{
		Initiator:  initiator,
		Challenger: challenger,
		Uuid:       uuid,
		Board:      board,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OpenGamePrefix(types.CurrGameKey, string(hashed)))

	store.Set([]byte(uuid), k.cdc.MustMarshal(&game))

}
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

func CreateGameKey(initiator string, uuid string) []byte {
	return []byte(initiator + "\x00" + uuid)
}

func (k Keeper) AddOGame (
	ctx sdk.Context,
	initiator string,
	uuid string,
	) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OpenGameKey))

	store.Set(CreateGameKey(initiator, uuid), make([]byte, 0))
}

func (k Keeper) HasOpenGame (
	ctx sdk.Context,
	initiator string,
	uuid string,
	) bool {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OpenGameKey))

	return store.Has(CreateGameKey(initiator, uuid))
}

func (k Keeper) RemoveOpenGame (
	ctx sdk.Context,
	initiator string,
	uuid string,
	) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OpenGameKey))
	store.Delete(CreateGameKey(initiator, uuid))
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

func FirstBitIsZero (hash []byte) bool {
	return hash[0] != 0
}

func (k Keeper) AddCurrGame (ctx sdk.Context, initiator string, uuid string, challenger string) {
	hashed := HashOpponents(initiator, challenger)

	x := initiator
	o := challenger

	if FirstBitIsZero(hashed) {
		x = challenger
		o = initiator
	}

	board := make([]byte, 9)		// initialize board, 9 cells

	game := types.CurrGame{
		X:  x,
		O: o,
		Uuid: uuid,
		Board: board,
		Turn: x,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrGameKey))

	store.Set(CreateGameKey(initiator, uuid), k.cdc.MustMarshal(&game))
}

func (k Keeper) SetCurrGame (ctx sdk.Context, initiator string, uuid string, game types.CurrGame) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrGameKey))

	store.Set(CreateGameKey(initiator, uuid), k.cdc.MustMarshal(&game))
}


func (k Keeper) HasCurrGame (
	ctx sdk.Context,
	initiator string,
	uuid string,
) bool {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrGameKey))

	return store.Has(CreateGameKey(initiator, uuid))
}


func (k Keeper) GetCurrGame (ctx sdk.Context, initiator string, uuid string) (types.CurrGame, error) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrGameKey))

	var currGame types.CurrGame
	v := store.Get(CreateGameKey(initiator, uuid))

	err := k.cdc.Unmarshal(v, &currGame)

	if err!= nil {
		return types.CurrGame{

		}, err
	}

	return currGame, nil
}
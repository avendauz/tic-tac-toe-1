package types

const (
	// ModuleName defines the module name
	ModuleName = "tictactoe1"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tictactoe1"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func OpenGamePrefix(p string, initiator string) []byte {
	return append(KeyPrefix(p), []byte(initiator)...)
}

const (
	OpenGameKey  = "Open-"
	CurrGameKey = "Curr-"
	DoneGameKey = "Done-"
)

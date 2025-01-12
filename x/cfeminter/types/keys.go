package types

// MinterKey is the key to use for the keeper store.
var IsGenesisKey = []byte{0x01}
var MinterStateKey = []byte{0x02}
var MinterStateHistoryKeyPrefix = []byte{0x03}

const (
	// ModuleName defines the module name
	ModuleName = "cfeminter"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cfeminter"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

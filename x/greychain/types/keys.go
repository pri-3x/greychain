package types

const (
	// ModuleName defines the module name
	ModuleName = "greychain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_greychain"
)

var (
	ParamsKey = []byte("p_greychain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

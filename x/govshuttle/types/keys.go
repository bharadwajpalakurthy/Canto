package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName defines the module name
	ModuleName = "shuttle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_govshuttle"
)

var (
	ModuleAddress common.Address
	PortKey = []byte("Port")
)

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress("govshuttle").Bytes())
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}

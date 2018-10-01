package vertcoin

import (
	"blockbook/bchain/coins/btc"

	"github.com/btcsuite/btcd/wire"
	"github.com/jakm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xdab5bffb
	TestnetMagic wire.BitcoinNet = 0x74726576 // "vert" word
	RegtestMagic wire.BitcoinNet = 0xdab5bffc
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func initParams() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{71}
	MainNetParams.ScriptHashAddrID = []byte{5}
	MainNetParams.Bech32HRPSegwit = "vtc"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{74}
	TestNetParams.ScriptHashAddrID = []byte{196}
	TestNetParams.Bech32HRPSegwit = "tvtc"

	err := chaincfg.Register(&MainNetParams)
	if err == nil {
		err = chaincfg.Register(&TestNetParams)
	}
	if err != nil {
		panic(err)
	}
}

// VertcoinParser handle
type VertcoinParser struct {
	*btc.BitcoinParser
}

// NewVertcoinParser returns new VertcoinParser instance
func NewVertcoinParser(params *chaincfg.Params, c *btc.Configuration) *VertcoinParser {
	return &VertcoinParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Vertcoin network,
// and the test Vertcoin network
func GetChainParams(chain string) *chaincfg.Params {
	if MainNetParams.Name == "" {
		initParams()
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
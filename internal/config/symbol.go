package config

var SymbolToken = map[string]uint8{
	"ETH":  0,
	"PEPE": 1,
	"USDT": 2,
}

var TokenSymbol = map[uint8]string{
	0: "ETH",
	1: "PEPE",
	2: "USDT",
}

var TokenDecimals = map[string]uint8{
	"ETH":  18,
	"USDT": 6,
	"PEPE": 18,
}

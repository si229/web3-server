package config

// var SymbolToken = map[string]string{
// 	"ETH":  "0x0000000000000000000000000000000000000000",
// 	"PEPE": "0x948B3c65b89DF0B4894ABE91E6D02FE579834F8F",
// 	"USDT": "0x8464135c8F25Da09e49BC8782676a84730C318bC",
// }

// var TokenSymbol = map[string]string{
// 	"0x0000000000000000000000000000000000000000": "ETH",
// 	"0x948B3c65b89DF0B4894ABE91E6D02FE579834F8F": "PEPE",
// 	"0x8464135c8F25Da09e49BC8782676a84730C318bC": "USDT",
// }

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

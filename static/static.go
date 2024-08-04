package static

var (
	TigrisIdToAssetInfo map[int]AssetInfo
)

type AssetInfo struct {
	Name            string
	ChainlinkFeed   string
	MinLeverage     string
	MaxLeverage     string
	FeeMultiplier   string
	BaseFundingRate string
	PythnetId       string
}

func InitStatic() {
	tigrisIdDetail := map[int]AssetInfo{}

	tigrisIdDetail[0] = AssetInfo{
		Name:            "BTC/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "5000000000",
		PythnetId:       "0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43",
	}

	tigrisIdDetail[1] = AssetInfo{
		Name:            "ETH/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "5000000000",
		PythnetId:       "0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace",
	}

	tigrisIdDetail[2] = AssetInfo{
		Name:            "XAU/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "5000000000",
		PythnetId:       "0x765d2ba906dbc32ca17cc11f5310a89e9ee1f6420508c63861f2f8ba4ee34bb2",
	}

	tigrisIdDetail[3] = AssetInfo{
		Name:            "MATIC/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "5000000000",
		PythnetId:       "0x5de33a9112c2b700b8d30b8a3402c103578ccfa2765696471cc672bd5cf6ac52",
	}

	tigrisIdDetail[4] = AssetInfo{
		Name:            "LINK/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0x8ac0c70fff57e9aefdf5edf44b51d62c2d433653cbb2cf5cc06bb115af04d221",
	}

	tigrisIdDetail[5] = AssetInfo{
		Name:            "EUR/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "4000000000000000000",
		MaxLeverage:     "500000000000000000000",
		FeeMultiplier:   "2000000000",
		BaseFundingRate: "500000000",
		PythnetId:       "0xa995d00bb36a63cef7fd2c287dc105fc8f3d93779f062f09551b0af3e81ec30b",
	}

	tigrisIdDetail[6] = AssetInfo{
		Name:            "GBP/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "4000000000000000000",
		MaxLeverage:     "500000000000000000000",
		FeeMultiplier:   "2000000000",
		BaseFundingRate: "500000000",
		PythnetId:       "0x84c2dde9633d93d1bcad84e7dc41c9d56578b7ec52fabedc1f335d673df0a7c1",
	}

	tigrisIdDetail[7] = AssetInfo{
		Name:            "USD/JPY",
		ChainlinkFeed:   "0",
		MinLeverage:     "4000000000000000000",
		MaxLeverage:     "500000000000000000000",
		FeeMultiplier:   "2000000000",
		BaseFundingRate: "500000000",
		PythnetId:       "0xef2c98c804ba503c6a707e38be4dfbb16683775f195b091252bf24693042fd52",
	}

	tigrisIdDetail[8] = AssetInfo{
		Name:            "",
		ChainlinkFeed:   "0",
		MinLeverage:     "0",
		MaxLeverage:     "0",
		FeeMultiplier:   "0",
		BaseFundingRate: "0",
		PythnetId:       "0x",
	}

	tigrisIdDetail[9] = AssetInfo{
		Name:            "",
		ChainlinkFeed:   "0",
		MinLeverage:     "0",
		MaxLeverage:     "0",
		FeeMultiplier:   "0",
		BaseFundingRate: "0",
		PythnetId:       "0x",
	}

	tigrisIdDetail[10] = AssetInfo{
		Name:            "USD/CAD",
		ChainlinkFeed:   "0",
		MinLeverage:     "4000000000000000000",
		MaxLeverage:     "500000000000000000000",
		FeeMultiplier:   "2000000000",
		BaseFundingRate: "500000000",
		PythnetId:       "0x3112b03a41c910ed446852aacf67118cb1bec67b2cd0b9a214c58cc0eaa2ecca",
	}

	tigrisIdDetail[11] = AssetInfo{
		Name:            "ETH/BTC",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "5000000000",
		PythnetId:       "0xc96458d393fe9deb7a7d63a0ac41e2898a67a7750dbd166673279e06c868df0a",
	}

	tigrisIdDetail[12] = AssetInfo{
		Name:            "",
		ChainlinkFeed:   "0",
		MinLeverage:     "0",
		MaxLeverage:     "0",
		FeeMultiplier:   "0",
		BaseFundingRate: "0",
		PythnetId:       "0x",
	}

	tigrisIdDetail[13] = AssetInfo{
		Name:            "BNB/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0x2f95862b045670cd22bee3114c39763a4a08beeb663b145d283c31d7d1101c4f",
	}

	tigrisIdDetail[14] = AssetInfo{
		Name:            "ADA/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "5000000000",
		PythnetId:       "0x2a01deaec9e51a579277b34b122399984d0bbf57e2458a7e42fecd2829867a0d",
	}

	tigrisIdDetail[15] = AssetInfo{
		Name:            "ATOM/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0xb00b60f88b03a6a625a8d1c048c3f66653edf217439983d037e7222c4e612819",
	}

	tigrisIdDetail[16] = AssetInfo{
		Name:            "",
		ChainlinkFeed:   "0",
		MinLeverage:     "0",
		MaxLeverage:     "0",
		FeeMultiplier:   "0",
		BaseFundingRate: "0",
		PythnetId:       "0x",
	}

	tigrisIdDetail[17] = AssetInfo{
		Name:            "",
		ChainlinkFeed:   "0",
		MinLeverage:     "0",
		MaxLeverage:     "0",
		FeeMultiplier:   "0",
		BaseFundingRate: "0",
		PythnetId:       "0x",
	}

	tigrisIdDetail[18] = AssetInfo{
		Name:            "SOL/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "5000000000",
		PythnetId:       "0xef0d8b6fda2ceba41da15d4095d1da392a0d2f8ed0c6c7bc0f4cfac8c280b56d",
	}

	tigrisIdDetail[19] = AssetInfo{
		Name:            "DOGE/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0xdcef50dd0a4cd2dcc17e45df1676dcb336a11a61c69df7a0299b0150c672d25c",
	}

	tigrisIdDetail[20] = AssetInfo{
		Name:            "LTC/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0x6e3f3fa8253588df9326580180233eb791e03b443a3ba7a1d892e73874e19a54",
	}

	tigrisIdDetail[21] = AssetInfo{
		Name:            "BCH/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "30000000000",
		PythnetId:       "0x3dd2b63686a450ec7290df3a1e0b583c0481f651351edfa7636f39aed55cf8a3",
	}

	tigrisIdDetail[22] = AssetInfo{
		Name:            "",
		ChainlinkFeed:   "0",
		MinLeverage:     "0",
		MaxLeverage:     "0",
		FeeMultiplier:   "0",
		BaseFundingRate: "0",
		PythnetId:       "0x",
	}

	tigrisIdDetail[23] = AssetInfo{
		Name:            "DOT/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0xca3eed9b267293f6595901c734c7525ce8ef49adafe8284606ceb307afa2ca5b",
	}

	tigrisIdDetail[24] = AssetInfo{
		Name:            "XMR/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "30000000000",
		PythnetId:       "0x46b8cc9347f04391764a0361e0b17c3ba394b001e7c304f7650f6376e37c321d",
	}

	tigrisIdDetail[25] = AssetInfo{
		Name:            "SHIB/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0xf0d57deca57b3da2fe63a493f4c25925fdfd8edf834b20f93e1f84dbd1504d4a",
	}

	tigrisIdDetail[26] = AssetInfo{
		Name:            "AVAX/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0x93da3352f9f1d105fdfe4971cfa80e9dd777bfc5d0f683ebb6e1294b92137bb7",
	}

	tigrisIdDetail[27] = AssetInfo{
		Name:            "UNI/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0x78d185a741d07edb3412b09008b7c5cfb9bbbd7d568bf00ba737b456ba171501",
	}

	tigrisIdDetail[28] = AssetInfo{
		Name:            "",
		ChainlinkFeed:   "0",
		MinLeverage:     "0",
		MaxLeverage:     "0",
		FeeMultiplier:   "0",
		BaseFundingRate: "0",
		PythnetId:       "0x",
	}

	tigrisIdDetail[29] = AssetInfo{
		Name:            "NEAR/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "5000000000",
		PythnetId:       "0xc415de8d2eba7db216527dff4b60e8f3a5311c740dadb233e13e12547e226750",
	}

	tigrisIdDetail[30] = AssetInfo{
		Name:            "ALGO/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "30000000000",
		PythnetId:       "0xfa17ceaf30d19ba51112fdcc750cc83454776f47fb0112e4af07f15f4bb1ebc0",
	}

	tigrisIdDetail[31] = AssetInfo{
		Name:            "",
		ChainlinkFeed:   "0",
		MinLeverage:     "0",
		MaxLeverage:     "0",
		FeeMultiplier:   "0",
		BaseFundingRate: "0",
		PythnetId:       "0x",
	}

	tigrisIdDetail[32] = AssetInfo{
		Name:            "XAG/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "5000000000",
		PythnetId:       "0xf2fb02c32b055c805e7238d628e5e9dadef274376114eb1f012337cabe93871e",
	}

	tigrisIdDetail[33] = AssetInfo{
		Name:            "LINK/BTC",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0x",
	}

	tigrisIdDetail[34] = AssetInfo{
		Name:            "XMR/BTC",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "30000000000",
		PythnetId:       "0x",
	}

	tigrisIdDetail[35] = AssetInfo{
		Name:            "ARB/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "10000000000",
		BaseFundingRate: "30000000000",
		PythnetId:       "0x3fa4252848f9f0a1480be62745a4629d9eb1322aebab8a791e344b3b9c1adcf5",
	}

	tigrisIdDetail[36] = AssetInfo{
		Name:            "PEPE/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "20000000000",
		BaseFundingRate: "60000000000",
		PythnetId:       "0xd69731a2e74ac1ce884fc3890f7ee324b6deb66147055249568869ed700882e4",
	}

	tigrisIdDetail[37] = AssetInfo{
		Name:            "GMX/USD",
		ChainlinkFeed:   "1253659165728396721884295642544774691906214818779",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "15000000000",
		BaseFundingRate: "30000000000",
		PythnetId:       "0xb962539d0fcb272a494d65ea56f94851c2bcf8823935da05bd628916e2e9edbf",
	}

	tigrisIdDetail[38] = AssetInfo{
		Name:            "XRP/USD",
		ChainlinkFeed:   "0",
		MinLeverage:     "2000000000000000000",
		MaxLeverage:     "100000000000000000000",
		FeeMultiplier:   "15000000000",
		BaseFundingRate: "20000000000",
		PythnetId:       "0xec5d399846a9209f3fe5881d70aae9268c94339ff9817e8d18ff19fa05eea1c8",
	}

	tigrisIdDetail[39] = AssetInfo{
		Name:            "USD/CHF",
		ChainlinkFeed:   "0",
		MinLeverage:     "4000000000000000000",
		MaxLeverage:     "500000000000000000000",
		FeeMultiplier:   "2000000000",
		BaseFundingRate: "500000000",
		PythnetId:       "0x0b1e3297e69f162877b577b0d6a47a0d63b2392bc8499e6540da4187a63e28f8",
	}

	tigrisIdDetail[40] = AssetInfo{
		Name:            "USD/CNH",
		ChainlinkFeed:   "0",
		MinLeverage:     "4000000000000000000",
		MaxLeverage:     "500000000000000000000",
		FeeMultiplier:   "2000000000",
		BaseFundingRate: "500000000",
		PythnetId:       "0xeef52e09c878ad41f6a81803e3640fe04dceea727de894edd4ea117e2e332e66",
	}

	tigrisIdDetail[41] = AssetInfo{
		Name:            "USD/MXN",
		ChainlinkFeed:   "0",
		MinLeverage:     "4000000000000000000",
		MaxLeverage:     "500000000000000000000",
		FeeMultiplier:   "2000000000",
		BaseFundingRate: "500000000",
		PythnetId:       "0xe13b1c1ffb32f34e1be9545583f01ef385fde7f42ee66049d30570dc866b77ca",
	}

	TigrisIdToAssetInfo = tigrisIdDetail
}

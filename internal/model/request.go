package request

type VerifySignatureReq struct {
	Address   string `json:"address"`
	Once      string `json:"once"`
	Signature string `json:"signature"`
}

type GenVerifyParamReq struct {
	Once string `json:"once"`
}

type VerifyValidate struct {
	Address string `json:"address"`
}

type GetBalanceReq struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
}

type GetContractBalanceReq struct {
	Address string `json:"address"`
}

type DepositReq struct {
	PrivateKey string  `json:"private_key"`
	Symbol     string  `json:"symbol"`
	Amount     float64 `json:"amount"`
}

type WithdrawReq struct {
	PrivateKey string  `json:"private_key"`
	Symbol     string  `json:"symbol"`
	Amount     float64 `json:"amount"`
}

type DepositDataReq struct {
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount"`
}

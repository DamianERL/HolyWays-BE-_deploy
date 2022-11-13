package transactiondto

type RequestTransaction struct {
	UserFundID int    `json:"userFund_id"`
	FundID        int    `json:"fund_id"`
	Donate        int    `json:"donate"`
	// Status        string `json:"status"`
}
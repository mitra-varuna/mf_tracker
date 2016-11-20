package mf_tracker

import (
	"math/big"
	"time"
)


type MutualFund struct {
	ISIN string
	CurrentNAV big.Float
	RePurchasePrice big.Float
	SalePrice big.Float
}

type MutualFundInvestment struct {
	Fund MutualFund
	Units int
	InvestmentDate time.Time
}

func (m MutualFundInvestment) CurrentValue() big.Float{
	return m.Fund.CurrentNAV * m.Units
}
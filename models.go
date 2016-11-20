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
	Units float32
	InvestmentDate time.Time
}

type PortFolio struct {
	Investments []MutualFundInvestment
}

func (m MutualFundInvestment) CurrentValue() big.Float{
	return m.Fund.CurrentNAV * m.Units
}

func (p PortFolio) TotalValue() big.Float{
	sum := 0
	for _, investment := range p.Investments {
		sum += investment.CurrentValue()
	}
	return sum
}

func (p PortFolio) NewInvestment(m MutualFund, amount big.Float) {
	var units = amount/m.CurrentNAV
	investment := MutualFundInvestment{Fund:m, Units:units, InvestmentDate:time.Now()}
	append(investment, p.Investments)
}
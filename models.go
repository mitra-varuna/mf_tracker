package mf_tracker

import (
	"math/big"
	"time"
)

//A Mutual fund model
type MutualFund struct {
	ISIN string
	CurrentNAV big.Float
	RePurchasePrice big.Float
	SalePrice big.Float
}

//A mutual fund investment model
type MutualFundInvestment struct {
	Fund MutualFund
	Units float32
	InvestmentDate time.Time
}

//Portfolio holds a number of mutual funds
type PortFolio struct {
	Investments []MutualFundInvestment
}

//Get Current value of the mutual fund investment
func (m MutualFundInvestment) CurrentValue() big.Float{
	return m.Fund.CurrentNAV * m.Units
}

//Get Total Value of the Investments
func (p PortFolio) TotalValue() big.Float{
	sum := 0
	for _, investment := range p.Investments {
		sum += investment.CurrentValue()
	}
	return sum
}

//Add a new mutual fund investment to the portfolio
func (p PortFolio) NewInvestment(m MutualFund, amount big.Float) {
	var units = amount/m.CurrentNAV
	investment := MutualFundInvestment{Fund:m, Units:units, InvestmentDate:time.Now()}
	append(investment, p.Investments)
}
package mf_tracker

import (
	"time"
)

//A Mutual fund model
type MutualFund struct {
	ISIN string
	CurrentNAV float64
	RePurchasePrice float64
	SalePrice float64
}

//A mutual fund investment model
type MutualFundInvestment struct {
	Fund MutualFund
	Units float64
	InvestmentDate time.Time
}

//Portfolio holds a number of mutual funds
type PortFolio struct {
	Investments []MutualFundInvestment
}

//Get Current value of the mutual fund investment
func (m MutualFundInvestment) CurrentValue() float64{
	return m.Fund.CurrentNAV * m.Units
}

//Get Total Value of the Investments
func (p PortFolio) TotalValue() float64{
	sum := 0.0
	for _, investment := range p.Investments {
		sum += investment.CurrentValue()
	}
	return sum
}

//Add a new mutual fund investment to the portfolio
func (p *PortFolio) NewInvestment(m MutualFund, amount float64) {
	var units = amount/m.CurrentNAV
	investment := MutualFundInvestment{Fund:m, Units:units, InvestmentDate:time.Now()}
	p.Investments = append(p.Investments, investment)
}
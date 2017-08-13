package mftracker

import (
	"fmt"
	"testing"
	"time"
)

func TestMutualFundInvestment_CurrentValue(t *testing.T) {
	fund := MutualFund{ISIN: "test", CurrentNAV: 12.0, RePurchasePrice: 32.0, SalePrice: 14.0}
	investment := MutualFundInvestment{Fund: fund, InvestmentDate: time.Now(), Units: 12}
	if investment.CurrentValue() != 12*12.0 {
		t.Error("Current Value is not evaluated correctly")
		t.Fail()
	}
}

func TestPortFolio_NewInvestment(t *testing.T) {
	fund1 := MutualFund{ISIN: "test", CurrentNAV: 12.0, RePurchasePrice: 32.0, SalePrice: 14.0}
	fund2 := MutualFund{ISIN: "test", CurrentNAV: 12.0, RePurchasePrice: 32.0, SalePrice: 14.0}
	fund3 := MutualFund{ISIN: "test", CurrentNAV: 12.0, RePurchasePrice: 32.0, SalePrice: 14.0}
	fund4 := MutualFund{ISIN: "test", CurrentNAV: 12.0, RePurchasePrice: 32.0, SalePrice: 14.0}
	folio := PortFolio{}
	folio.NewInvestment(fund1, 25.0)
	folio.NewInvestment(fund2, 34.0)
	folio.NewInvestment(fund3, 32.0)
	folio.NewInvestment(fund4, 12.0)
	if len(folio.Investments) != 4 {
		fmt.Println(len(folio.Investments))
		t.Fail()
	}
}

func TestPortFolio_TotalValue(t *testing.T) {
	fund1 := MutualFund{ISIN: "test", CurrentNAV: 12.0, RePurchasePrice: 32.0, SalePrice: 14.0}
	fund2 := MutualFund{ISIN: "test", CurrentNAV: 12.0, RePurchasePrice: 32.0, SalePrice: 14.0}
	fund3 := MutualFund{ISIN: "test", CurrentNAV: 12.0, RePurchasePrice: 32.0, SalePrice: 14.0}
	fund4 := MutualFund{ISIN: "test", CurrentNAV: 12.0, RePurchasePrice: 32.0, SalePrice: 14.0}
	folio := PortFolio{}
	folio.NewInvestment(fund1, 25.0)
	folio.NewInvestment(fund2, 34.0)
	folio.NewInvestment(fund3, 32.0)
	folio.NewInvestment(fund4, 12.0)
	if folio.TotalValue() != 103 {
		t.Error(folio.TotalValue())
		t.Fail()
	}
}
